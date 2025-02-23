// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Topic_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Topic_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicStatusARM, TopicStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicStatusARM runs a test to see if a specific instance of Topic_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicStatusARM(subject Topic_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Topic_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Topic_StatusARM instances for property testing - lazily instantiated by TopicStatusARMGenerator()
var topicStatusARMGenerator gopter.Gen

// TopicStatusARMGenerator returns a generator of Topic_StatusARM instances for property testing.
// We first initialize topicStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TopicStatusARMGenerator() gopter.Gen {
	if topicStatusARMGenerator != nil {
		return topicStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicStatusARM(generators)
	topicStatusARMGenerator = gen.Struct(reflect.TypeOf(Topic_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicStatusARM(generators)
	AddRelatedPropertyGeneratorsForTopicStatusARM(generators)
	topicStatusARMGenerator = gen.Struct(reflect.TypeOf(Topic_StatusARM{}), generators)

	return topicStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForTopicStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForTopicStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(TopicPropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_TopicProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of TopicProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForTopicPropertiesStatusARM, TopicPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForTopicPropertiesStatusARM runs a test to see if a specific instance of TopicProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForTopicPropertiesStatusARM(subject TopicProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual TopicProperties_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of TopicProperties_StatusARM instances for property testing - lazily instantiated by
//TopicPropertiesStatusARMGenerator()
var topicPropertiesStatusARMGenerator gopter.Gen

// TopicPropertiesStatusARMGenerator returns a generator of TopicProperties_StatusARM instances for property testing.
// We first initialize topicPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func TopicPropertiesStatusARMGenerator() gopter.Gen {
	if topicPropertiesStatusARMGenerator != nil {
		return topicPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicPropertiesStatusARM(generators)
	topicPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(TopicProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForTopicPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForTopicPropertiesStatusARM(generators)
	topicPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(TopicProperties_StatusARM{}), generators)

	return topicPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForTopicPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForTopicPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Endpoint"] = gen.PtrOf(gen.AlphaString())
	gens["InputSchema"] = gen.PtrOf(gen.OneConstOf(TopicPropertiesStatusInputSchemaCloudEventSchemaV10, TopicPropertiesStatusInputSchemaCustomEventSchema, TopicPropertiesStatusInputSchemaEventGridSchema))
	gens["MetricResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		TopicPropertiesStatusProvisioningStateCanceled,
		TopicPropertiesStatusProvisioningStateCreating,
		TopicPropertiesStatusProvisioningStateDeleting,
		TopicPropertiesStatusProvisioningStateFailed,
		TopicPropertiesStatusProvisioningStateSucceeded,
		TopicPropertiesStatusProvisioningStateUpdating))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(TopicPropertiesStatusPublicNetworkAccessDisabled, TopicPropertiesStatusPublicNetworkAccessEnabled))
}

// AddRelatedPropertyGeneratorsForTopicPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForTopicPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["InboundIpRules"] = gen.SliceOf(InboundIpRuleStatusARMGenerator())
	gens["InputSchemaMapping"] = gen.PtrOf(InputSchemaMappingStatusARMGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator())
}

func Test_PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM, PrivateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM runs a test to see if a specific instance of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM(subject PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM instances for property testing - lazily
//instantiated by PrivateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator()
var privateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator gopter.Gen

// PrivateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator returns a generator of PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM instances for property testing.
func PrivateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator() gopter.Gen {
	if privateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator != nil {
		return privateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM(generators)
	privateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM{}), generators)

	return privateEndpointConnectionStatusTopicSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionStatusTopicSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
