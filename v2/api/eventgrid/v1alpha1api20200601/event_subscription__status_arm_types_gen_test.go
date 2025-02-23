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

func Test_EventSubscription_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscription_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionStatusARM, EventSubscriptionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionStatusARM runs a test to see if a specific instance of EventSubscription_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionStatusARM(subject EventSubscription_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscription_StatusARM
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

// Generator of EventSubscription_StatusARM instances for property testing - lazily instantiated by
//EventSubscriptionStatusARMGenerator()
var eventSubscriptionStatusARMGenerator gopter.Gen

// EventSubscriptionStatusARMGenerator returns a generator of EventSubscription_StatusARM instances for property testing.
// We first initialize eventSubscriptionStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionStatusARMGenerator() gopter.Gen {
	if eventSubscriptionStatusARMGenerator != nil {
		return eventSubscriptionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionStatusARM(generators)
	eventSubscriptionStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscription_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionStatusARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionStatusARM(generators)
	eventSubscriptionStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscription_StatusARM{}), generators)

	return eventSubscriptionStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(EventSubscriptionPropertiesStatusARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_EventSubscriptionProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionPropertiesStatusARM, EventSubscriptionPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionPropertiesStatusARM runs a test to see if a specific instance of EventSubscriptionProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionPropertiesStatusARM(subject EventSubscriptionProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionProperties_StatusARM
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

// Generator of EventSubscriptionProperties_StatusARM instances for property testing - lazily instantiated by
//EventSubscriptionPropertiesStatusARMGenerator()
var eventSubscriptionPropertiesStatusARMGenerator gopter.Gen

// EventSubscriptionPropertiesStatusARMGenerator returns a generator of EventSubscriptionProperties_StatusARM instances for property testing.
// We first initialize eventSubscriptionPropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionPropertiesStatusARMGenerator() gopter.Gen {
	if eventSubscriptionPropertiesStatusARMGenerator != nil {
		return eventSubscriptionPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesStatusARM(generators)
	eventSubscriptionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesStatusARM(generators)
	eventSubscriptionPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionProperties_StatusARM{}), generators)

	return eventSubscriptionPropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["EventDeliverySchema"] = gen.PtrOf(gen.OneConstOf(EventSubscriptionPropertiesStatusEventDeliverySchemaCloudEventSchemaV10, EventSubscriptionPropertiesStatusEventDeliverySchemaCustomInputSchema, EventSubscriptionPropertiesStatusEventDeliverySchemaEventGridSchema))
	gens["ExpirationTimeUtc"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.SliceOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		EventSubscriptionPropertiesStatusProvisioningStateAwaitingManualAction,
		EventSubscriptionPropertiesStatusProvisioningStateCanceled,
		EventSubscriptionPropertiesStatusProvisioningStateCreating,
		EventSubscriptionPropertiesStatusProvisioningStateDeleting,
		EventSubscriptionPropertiesStatusProvisioningStateFailed,
		EventSubscriptionPropertiesStatusProvisioningStateSucceeded,
		EventSubscriptionPropertiesStatusProvisioningStateUpdating))
	gens["Topic"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["DeadLetterDestination"] = gen.PtrOf(DeadLetterDestinationStatusARMGenerator())
	gens["Destination"] = gen.PtrOf(EventSubscriptionDestinationStatusARMGenerator())
	gens["Filter"] = gen.PtrOf(EventSubscriptionFilterStatusARMGenerator())
	gens["RetryPolicy"] = gen.PtrOf(RetryPolicyStatusARMGenerator())
}

func Test_DeadLetterDestination_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DeadLetterDestination_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDeadLetterDestinationStatusARM, DeadLetterDestinationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDeadLetterDestinationStatusARM runs a test to see if a specific instance of DeadLetterDestination_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDeadLetterDestinationStatusARM(subject DeadLetterDestination_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DeadLetterDestination_StatusARM
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

// Generator of DeadLetterDestination_StatusARM instances for property testing - lazily instantiated by
//DeadLetterDestinationStatusARMGenerator()
var deadLetterDestinationStatusARMGenerator gopter.Gen

// DeadLetterDestinationStatusARMGenerator returns a generator of DeadLetterDestination_StatusARM instances for property testing.
func DeadLetterDestinationStatusARMGenerator() gopter.Gen {
	if deadLetterDestinationStatusARMGenerator != nil {
		return deadLetterDestinationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDeadLetterDestinationStatusARM(generators)
	deadLetterDestinationStatusARMGenerator = gen.Struct(reflect.TypeOf(DeadLetterDestination_StatusARM{}), generators)

	return deadLetterDestinationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDeadLetterDestinationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDeadLetterDestinationStatusARM(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.OneConstOf(DeadLetterDestinationStatusEndpointTypeStorageBlob)
}

func Test_EventSubscriptionDestination_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionDestination_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionDestinationStatusARM, EventSubscriptionDestinationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionDestinationStatusARM runs a test to see if a specific instance of EventSubscriptionDestination_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionDestinationStatusARM(subject EventSubscriptionDestination_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionDestination_StatusARM
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

// Generator of EventSubscriptionDestination_StatusARM instances for property testing - lazily instantiated by
//EventSubscriptionDestinationStatusARMGenerator()
var eventSubscriptionDestinationStatusARMGenerator gopter.Gen

// EventSubscriptionDestinationStatusARMGenerator returns a generator of EventSubscriptionDestination_StatusARM instances for property testing.
func EventSubscriptionDestinationStatusARMGenerator() gopter.Gen {
	if eventSubscriptionDestinationStatusARMGenerator != nil {
		return eventSubscriptionDestinationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionDestinationStatusARM(generators)
	eventSubscriptionDestinationStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionDestination_StatusARM{}), generators)

	return eventSubscriptionDestinationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionDestinationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionDestinationStatusARM(gens map[string]gopter.Gen) {
	gens["EndpointType"] = gen.OneConstOf(
		EventSubscriptionDestinationStatusEndpointTypeAzureFunction,
		EventSubscriptionDestinationStatusEndpointTypeEventHub,
		EventSubscriptionDestinationStatusEndpointTypeHybridConnection,
		EventSubscriptionDestinationStatusEndpointTypeServiceBusQueue,
		EventSubscriptionDestinationStatusEndpointTypeServiceBusTopic,
		EventSubscriptionDestinationStatusEndpointTypeStorageQueue,
		EventSubscriptionDestinationStatusEndpointTypeWebHook)
}

func Test_EventSubscriptionFilter_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EventSubscriptionFilter_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventSubscriptionFilterStatusARM, EventSubscriptionFilterStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventSubscriptionFilterStatusARM runs a test to see if a specific instance of EventSubscriptionFilter_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventSubscriptionFilterStatusARM(subject EventSubscriptionFilter_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EventSubscriptionFilter_StatusARM
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

// Generator of EventSubscriptionFilter_StatusARM instances for property testing - lazily instantiated by
//EventSubscriptionFilterStatusARMGenerator()
var eventSubscriptionFilterStatusARMGenerator gopter.Gen

// EventSubscriptionFilterStatusARMGenerator returns a generator of EventSubscriptionFilter_StatusARM instances for property testing.
// We first initialize eventSubscriptionFilterStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventSubscriptionFilterStatusARMGenerator() gopter.Gen {
	if eventSubscriptionFilterStatusARMGenerator != nil {
		return eventSubscriptionFilterStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilterStatusARM(generators)
	eventSubscriptionFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventSubscriptionFilterStatusARM(generators)
	AddRelatedPropertyGeneratorsForEventSubscriptionFilterStatusARM(generators)
	eventSubscriptionFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(EventSubscriptionFilter_StatusARM{}), generators)

	return eventSubscriptionFilterStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventSubscriptionFilterStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventSubscriptionFilterStatusARM(gens map[string]gopter.Gen) {
	gens["IncludedEventTypes"] = gen.SliceOf(gen.AlphaString())
	gens["IsSubjectCaseSensitive"] = gen.PtrOf(gen.Bool())
	gens["SubjectBeginsWith"] = gen.PtrOf(gen.AlphaString())
	gens["SubjectEndsWith"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventSubscriptionFilterStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventSubscriptionFilterStatusARM(gens map[string]gopter.Gen) {
	gens["AdvancedFilters"] = gen.SliceOf(AdvancedFilterStatusARMGenerator())
}

func Test_RetryPolicy_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetryPolicy_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetryPolicyStatusARM, RetryPolicyStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetryPolicyStatusARM runs a test to see if a specific instance of RetryPolicy_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRetryPolicyStatusARM(subject RetryPolicy_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetryPolicy_StatusARM
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

// Generator of RetryPolicy_StatusARM instances for property testing - lazily instantiated by
//RetryPolicyStatusARMGenerator()
var retryPolicyStatusARMGenerator gopter.Gen

// RetryPolicyStatusARMGenerator returns a generator of RetryPolicy_StatusARM instances for property testing.
func RetryPolicyStatusARMGenerator() gopter.Gen {
	if retryPolicyStatusARMGenerator != nil {
		return retryPolicyStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetryPolicyStatusARM(generators)
	retryPolicyStatusARMGenerator = gen.Struct(reflect.TypeOf(RetryPolicy_StatusARM{}), generators)

	return retryPolicyStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForRetryPolicyStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetryPolicyStatusARM(gens map[string]gopter.Gen) {
	gens["EventTimeToLiveInMinutes"] = gen.PtrOf(gen.Int())
	gens["MaxDeliveryAttempts"] = gen.PtrOf(gen.Int())
}

func Test_AdvancedFilter_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AdvancedFilter_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAdvancedFilterStatusARM, AdvancedFilterStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAdvancedFilterStatusARM runs a test to see if a specific instance of AdvancedFilter_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAdvancedFilterStatusARM(subject AdvancedFilter_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AdvancedFilter_StatusARM
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

// Generator of AdvancedFilter_StatusARM instances for property testing - lazily instantiated by
//AdvancedFilterStatusARMGenerator()
var advancedFilterStatusARMGenerator gopter.Gen

// AdvancedFilterStatusARMGenerator returns a generator of AdvancedFilter_StatusARM instances for property testing.
func AdvancedFilterStatusARMGenerator() gopter.Gen {
	if advancedFilterStatusARMGenerator != nil {
		return advancedFilterStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAdvancedFilterStatusARM(generators)
	advancedFilterStatusARMGenerator = gen.Struct(reflect.TypeOf(AdvancedFilter_StatusARM{}), generators)

	return advancedFilterStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAdvancedFilterStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAdvancedFilterStatusARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(gen.AlphaString())
	gens["OperatorType"] = gen.OneConstOf(
		AdvancedFilterStatusOperatorTypeBoolEquals,
		AdvancedFilterStatusOperatorTypeNumberGreaterThan,
		AdvancedFilterStatusOperatorTypeNumberGreaterThanOrEquals,
		AdvancedFilterStatusOperatorTypeNumberIn,
		AdvancedFilterStatusOperatorTypeNumberLessThan,
		AdvancedFilterStatusOperatorTypeNumberLessThanOrEquals,
		AdvancedFilterStatusOperatorTypeNumberNotIn,
		AdvancedFilterStatusOperatorTypeStringBeginsWith,
		AdvancedFilterStatusOperatorTypeStringContains,
		AdvancedFilterStatusOperatorTypeStringEndsWith,
		AdvancedFilterStatusOperatorTypeStringIn,
		AdvancedFilterStatusOperatorTypeStringNotIn)
}
