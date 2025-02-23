/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/version"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/codegen/pipeline"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

// CodeGenerator is a generator of code
type CodeGenerator struct {
	configuration *config.Configuration
	pipeline      []pipeline.Stage
}

// NewCodeGeneratorFromConfigFile produces a new Generator with the given configuration file
func NewCodeGeneratorFromConfigFile(configurationFile string) (*CodeGenerator, error) {
	configuration, err := config.LoadConfiguration(configurationFile)
	if err != nil {
		return nil, err
	}

	target, err := pipeline.TranslatePipelineToTarget(configuration.Pipeline)
	if err != nil {
		return nil, err
	}

	return NewTargetedCodeGeneratorFromConfig(configuration, astmodel.NewIdentifierFactory(), target)
}

// NewTargetedCodeGeneratorFromConfig produces a new code generator with the given configuration and
// only the stages appropriate for the specified target.
func NewTargetedCodeGeneratorFromConfig(
	configuration *config.Configuration, idFactory astmodel.IdentifierFactory, target pipeline.Target) (*CodeGenerator, error) {

	result, err := NewCodeGeneratorFromConfig(configuration, idFactory)
	if err != nil {
		return nil, errors.Wrapf(err, "creating pipeline targeting %s", target)
	}

	// Filter stages to use only those appropriate for our target
	var stages []pipeline.Stage
	for _, s := range result.pipeline {
		if s.IsUsedFor(target) {
			stages = append(stages, s)
		}
	}

	result.pipeline = stages

	err = result.verifyPipeline()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewCodeGeneratorFromConfig produces a new code generator with the given configuration all available stages
func NewCodeGeneratorFromConfig(configuration *config.Configuration, idFactory astmodel.IdentifierFactory) (*CodeGenerator, error) {
	result := &CodeGenerator{
		configuration: configuration,
		pipeline:      createAllPipelineStages(idFactory, configuration),
	}

	return result, nil
}

func createAllPipelineStages(idFactory astmodel.IdentifierFactory, configuration *config.Configuration) []pipeline.Stage {
	return []pipeline.Stage{

		pipeline.LoadSchemaIntoTypes(idFactory, configuration, pipeline.DefaultSchemaLoader),

		// Import status info from Swagger:
		pipeline.AddStatusFromSwagger(idFactory, configuration),

		// Reduces oneOf/allOf types from schemas to object types:
		pipeline.ConvertAllOfAndOneOfToObjects(idFactory),

		// Flatten out any nested resources created by allOf, etc. we want to do this before naming types or things
		// get named with names like Resource_Spec_Spec_Spec:
		pipeline.FlattenResources(),

		// Copy additional swagger-derived information from status into spec
		pipeline.AugmentSpecWithStatus().RequiresPrerequisiteStages("allof-anyof-objects", "addStatusFromSwagger"),

		pipeline.StripUnreferencedTypeDefinitions(),

		// Name all anonymous object, enum, and validated types (required by controller-gen):
		pipeline.NameTypesForCRD(idFactory),

		// Apply property type rewrites from the config file
		// Must come after NameTypesForCRD ('nameTypes)' and ConvertAllOfAndOneOfToObjects ('allof-anyof-objects') so
		// that objects are all expanded
		pipeline.ApplyPropertyRewrites(configuration).
			RequiresPrerequisiteStages("nameTypes", "allof-anyof-objects"),
		pipeline.RemoveResourceScope(),

		pipeline.MakeStatusPropertiesOptional(),
		pipeline.RemoveStatusValidations(),
		pipeline.UnrollRecursiveTypes(),

		// Figure out resource owners:
		pipeline.DetermineResourceOwnership(configuration),

		// Strip out redundant type aliases:
		pipeline.RemoveTypeAliases(),

		// Collapse cross group references
		pipeline.CollapseCrossGroupReferences(),

		// De-pluralize resource types
		// (Must come after type aliases are resolved)
		pipeline.ImproveResourcePluralization(),

		pipeline.StripUnreferencedTypeDefinitions(),

		pipeline.AssertTypesCollectionValid(),

		pipeline.RemoveEmbeddedResources().UsedFor(pipeline.ARMTarget), // TODO: For now only used for ARM,

		// Apply export filters before generating
		// ARM types for resources etc:
		pipeline.ApplyExportFilters(configuration),

		// TODO: These should be removed if/when we move to Swagger as the single source of truth
		pipeline.RemoveTypeProperty(),
		pipeline.RemoveAPIVersionProperty(),

		pipeline.VerifyNoErroredTypes(),

		pipeline.StripUnreferencedTypeDefinitions(),

		pipeline.ReplaceAnyTypeWithJSON(),

		pipeline.AddCrossResourceReferences(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.AddSecrets(configuration).UsedFor(pipeline.ARMTarget),

		pipeline.ReportOnTypesAndVersions(configuration).UsedFor(pipeline.ARMTarget), // TODO: For now only used for ARM

		pipeline.CreateARMTypes(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ApplyARMConversionInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ApplyKubernetesResourceInterface(idFactory).UsedFor(pipeline.ARMTarget),

		// Effects the "flatten" property of Properties:
		pipeline.FlattenProperties(),

		// Remove types which may not be needed after flattening
		pipeline.StripUnreferencedTypeDefinitions(),

		pipeline.AddStatusConditions(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.AddOperatorSpec(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		// To be added when needed
		// pipeline.AddOperatorStatus(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.AddCrossplaneOwnerProperties(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneForProvider(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneAtProvider(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneEmbeddedResourceSpec(idFactory).UsedFor(pipeline.CrossplaneTarget),
		pipeline.AddCrossplaneEmbeddedResourceStatus(idFactory).UsedFor(pipeline.CrossplaneTarget),

		// Create Storage types
		// TODO: For now only used for ARM
		pipeline.CreateConversionGraph(configuration).UsedFor(pipeline.ARMTarget),
		pipeline.InjectOriginalVersionFunction(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.CreateStorageTypes().UsedFor(pipeline.ARMTarget),
		pipeline.InjectOriginalVersionProperty().UsedFor(pipeline.ARMTarget),
		pipeline.InjectPropertyAssignmentFunctions(configuration, idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleSpecInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleStatusInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectOriginalGVKFunction(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.MarkLatestStorageVariantAsHubVersion().UsedFor(pipeline.ARMTarget),
		pipeline.MarkLatestAPIVersionAsStorageVersion().UsedFor(pipeline.CrossplaneTarget),

		pipeline.InjectHubFunction(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.ImplementConvertibleInterface(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectResourceConversionTestCases(idFactory).UsedFor(pipeline.ARMTarget),

		// Inject test cases
		pipeline.InjectJsonSerializationTests(idFactory).UsedFor(pipeline.ARMTarget),
		pipeline.InjectPropertyAssignmentTests(idFactory).UsedFor(pipeline.ARMTarget),

		pipeline.SimplifyDefinitions(),

		// Create Resource Extensions
		pipeline.CreateResourceExtensions(configuration.LocalPathPrefix(), idFactory).UsedFor(pipeline.ARMTarget),

		// Safety checks at the end:
		pipeline.EnsureDefinitionsDoNotUseAnyTypes(),
		pipeline.EnsureARMTypeExistsForEveryResource().UsedFor(pipeline.ARMTarget),
		pipeline.DetectSkippingProperties().UsedFor(pipeline.ARMTarget),

		pipeline.DeleteGeneratedCode(configuration.FullTypesOutputPath()),
		pipeline.ExportPackages(configuration.FullTypesOutputPath()),

		pipeline.ExportControllerResourceRegistrations(idFactory, configuration.FullTypesRegistrationOutputFilePath()).UsedFor(pipeline.ARMTarget),

		pipeline.ReportResourceVersions(configuration),
	}
}

// Generate produces the Go code corresponding to the configured JSON schema in the given output folder
func (generator *CodeGenerator) Generate(ctx context.Context) error {
	klog.V(1).Infof("Generator version: %s", version.BuildVersion)

	state := pipeline.NewState()
	for i, stage := range generator.pipeline {
		klog.V(0).Infof("%d/%d: %s", i+1, len(generator.pipeline), stage.Description())
		// Defensive copy (in case the pipeline modifies its inputs) so that we can compare types in vs out
		stateOut, err := stage.Run(ctx, state)
		if err != nil {
			return errors.Wrapf(err, "failed during pipeline stage %d/%d: %s", i+1, len(generator.pipeline), stage.Description())
		}

		defsAdded := stateOut.Definitions().Except(state.Definitions())
		defsRemoved := state.Definitions().Except(stateOut.Definitions())

		if len(defsAdded) > 0 && len(defsRemoved) > 0 {
			klog.V(1).Infof("Added %d, removed %d type definitions", len(defsAdded), len(defsRemoved))
		} else if len(defsAdded) > 0 {
			klog.V(1).Infof("Added %d type definitions", len(defsAdded))
		} else if len(defsRemoved) > 0 {
			klog.V(1).Infof("Removed %d type definitions", len(defsRemoved))
		}

		state = stateOut
	}

	klog.Info("Finished")

	return nil
}

func (generator *CodeGenerator) verifyPipeline() error {
	var errs []error

	// Set of stages that we've already seen, used to confirm prerequisites
	stagesSeen := astmodel.MakeStringSet()

	// Set of stages we expect to see, each associated with a slice containing the earlier stages that expected each
	stagesExpected := make(map[string][]string)

	for _, stage := range generator.pipeline {
		err := stage.CheckPrerequisites(stagesSeen)
		if err != nil {
			errs = append(errs, err)
		}

		for _, postreq := range stage.Postrequisites() {
			if stagesSeen.Contains(postreq) {
				errs = append(errs, errors.Errorf("postrequisite %q of stage %q satisfied too early", postreq, stage.Id()))
			} else {
				stagesExpected[postreq] = append(stagesExpected[postreq], stage.Id())
			}
		}

		stagesSeen.Add(stage.Id())
		delete(stagesExpected, stage.Id())
	}

	for required, requiredBy := range stagesExpected {
		for _, stageId := range requiredBy {
			errs = append(errs, errors.Errorf("postrequisite %q of stage %q not satisfied", required, stageId))
		}
	}

	return kerrors.NewAggregate(errs)
}

// RemoveStages will remove all stages from the pipeline with the given ids.
// Only available for test builds.
// Will panic if you specify an unknown id.
func (generator *CodeGenerator) RemoveStages(stageIds ...string) {
	stagesToRemove := make(map[string]bool)
	for _, s := range stageIds {
		stagesToRemove[s] = false
	}

	var stages []pipeline.Stage

	for _, stage := range generator.pipeline {
		if _, ok := stagesToRemove[stage.Id()]; ok {
			stagesToRemove[stage.Id()] = true
			continue
		}

		stages = append(stages, stage)
	}

	for stage, removed := range stagesToRemove {
		if !removed {
			panic(fmt.Sprintf("Expected to remove stage %s from stages, but it wasn't found.", stage))
		}
	}

	generator.pipeline = stages
}

// ReplaceStage replaces all uses of an existing stage with another one.
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) ReplaceStage(existingStage string, stage pipeline.Stage) {
	replaced := false
	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			generator.pipeline[i] = stage
			replaced = true
		}
	}

	if !replaced {
		panic(fmt.Sprintf("Expected to replace stage %s but it wasn't found", existingStage))
	}
}

// InjectStageAfter injects a new stage immediately after the first occurrence of an existing stage
// Only available for test builds.
// Will panic if the existing stage is not found.
func (generator *CodeGenerator) InjectStageAfter(existingStage string, stage pipeline.Stage) {
	injected := false

	for i, s := range generator.pipeline {
		if s.HasId(existingStage) {
			var p []pipeline.Stage
			p = append(p, generator.pipeline[:i+1]...)
			p = append(p, stage)
			p = append(p, generator.pipeline[i+1:]...)
			generator.pipeline = p
			injected = true
			break
		}
	}

	if !injected {
		panic(fmt.Sprintf("Expected to inject stage %s but %s wasn't found", stage.Id(), existingStage))
	}
}

// HasStage tests whether the pipeline has a stage with the given id
// Only available for test builds.
func (generator *CodeGenerator) HasStage(id string) bool {
	return generator.IndexOfStage(id) != -1
}

// IndexOfStage returns the index of the stage, if present, or -1 if not
// Only available for test builds.
func (generator *CodeGenerator) IndexOfStage(id string) int {
	for i, s := range generator.pipeline {
		if s.HasId(id) {
			return i
		}
	}

	return -1
}
