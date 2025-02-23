// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
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

func Test_NamespacesAuthorizationRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesAuthorizationRule tests if a specific instance of NamespacesAuthorizationRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1alpha1api20211101storage.NamespacesAuthorizationRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesAuthorizationRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule to NamespacesAuthorizationRule via AssignPropertiesToNamespacesAuthorizationRule & AssignPropertiesFromNamespacesAuthorizationRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRule tests if a specific instance of NamespacesAuthorizationRule can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20211101storage.NamespacesAuthorizationRule
	err := copied.AssignPropertiesToNamespacesAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRule
	err = actual.AssignPropertiesFromNamespacesAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule runs a test to see if a specific instance of NamespacesAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule
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

// Generator of NamespacesAuthorizationRule instances for property testing - lazily instantiated by
//NamespacesAuthorizationRuleGenerator()
var namespacesAuthorizationRuleGenerator gopter.Gen

// NamespacesAuthorizationRuleGenerator returns a generator of NamespacesAuthorizationRule instances for property testing.
func NamespacesAuthorizationRuleGenerator() gopter.Gen {
	if namespacesAuthorizationRuleGenerator != nil {
		return namespacesAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(generators)
	namespacesAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule{}), generators)

	return namespacesAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesAuthorizationRulesSpecGenerator()
	gens["Status"] = AuthorizationRuleStatusGenerator()
}

func Test_AuthorizationRule_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationRule_Status to AuthorizationRule_Status via AssignPropertiesToAuthorizationRuleStatus & AssignPropertiesFromAuthorizationRuleStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationRuleStatus, AuthorizationRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationRuleStatus tests if a specific instance of AuthorizationRule_Status can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationRuleStatus(subject AuthorizationRule_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20211101storage.AuthorizationRule_Status
	err := copied.AssignPropertiesToAuthorizationRuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationRule_Status
	err = actual.AssignPropertiesFromAuthorizationRuleStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationRule_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRule_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRuleStatus, AuthorizationRuleStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRuleStatus runs a test to see if a specific instance of AuthorizationRule_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRuleStatus(subject AuthorizationRule_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRule_Status
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

// Generator of AuthorizationRule_Status instances for property testing - lazily instantiated by
//AuthorizationRuleStatusGenerator()
var authorizationRuleStatusGenerator gopter.Gen

// AuthorizationRuleStatusGenerator returns a generator of AuthorizationRule_Status instances for property testing.
// We first initialize authorizationRuleStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationRuleStatusGenerator() gopter.Gen {
	if authorizationRuleStatusGenerator != nil {
		return authorizationRuleStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleStatus(generators)
	authorizationRuleStatusGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleStatus(generators)
	AddRelatedPropertyGeneratorsForAuthorizationRuleStatus(generators)
	authorizationRuleStatusGenerator = gen.Struct(reflect.TypeOf(AuthorizationRule_Status{}), generators)

	return authorizationRuleStatusGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRuleStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRuleStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRuleStatusPropertiesRightsListen, AuthorizationRuleStatusPropertiesRightsManage, AuthorizationRuleStatusPropertiesRightsSend))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationRuleStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationRuleStatus(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_NamespacesAuthorizationRules_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRules_Spec to NamespacesAuthorizationRules_Spec via AssignPropertiesToNamespacesAuthorizationRulesSpec & AssignPropertiesFromNamespacesAuthorizationRulesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRulesSpec, NamespacesAuthorizationRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRulesSpec tests if a specific instance of NamespacesAuthorizationRules_Spec can be assigned to v1alpha1api20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRulesSpec(subject NamespacesAuthorizationRules_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20211101storage.NamespacesAuthorizationRules_Spec
	err := copied.AssignPropertiesToNamespacesAuthorizationRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRules_Spec
	err = actual.AssignPropertiesFromNamespacesAuthorizationRulesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRules_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRules_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRulesSpec, NamespacesAuthorizationRulesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRulesSpec runs a test to see if a specific instance of NamespacesAuthorizationRules_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRulesSpec(subject NamespacesAuthorizationRules_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRules_Spec
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

// Generator of NamespacesAuthorizationRules_Spec instances for property testing - lazily instantiated by
//NamespacesAuthorizationRulesSpecGenerator()
var namespacesAuthorizationRulesSpecGenerator gopter.Gen

// NamespacesAuthorizationRulesSpecGenerator returns a generator of NamespacesAuthorizationRules_Spec instances for property testing.
func NamespacesAuthorizationRulesSpecGenerator() gopter.Gen {
	if namespacesAuthorizationRulesSpecGenerator != nil {
		return namespacesAuthorizationRulesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpec(generators)
	namespacesAuthorizationRulesSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRules_Spec{}), generators)

	return namespacesAuthorizationRulesSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRulesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRulePropertiesRightsListen, AuthorizationRulePropertiesRightsManage, AuthorizationRulePropertiesRightsSend))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
