// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

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

func Test_DatabaseAccountsMongodbDatabasesCollections_SPECARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsMongodbDatabasesCollections_SPECARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollections_SPECARM, DatabaseAccountsMongodbDatabasesCollections_SPECARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollections_SPECARM runs a test to see if a specific instance of DatabaseAccountsMongodbDatabasesCollections_SPECARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsMongodbDatabasesCollections_SPECARM(subject DatabaseAccountsMongodbDatabasesCollections_SPECARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsMongodbDatabasesCollections_SPECARM
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

// Generator of DatabaseAccountsMongodbDatabasesCollections_SPECARM instances for property testing - lazily instantiated
//by DatabaseAccountsMongodbDatabasesCollections_SPECARMGenerator()
var databaseAccountsMongodbDatabasesCollections_specarmGenerator gopter.Gen

// DatabaseAccountsMongodbDatabasesCollections_SPECARMGenerator returns a generator of DatabaseAccountsMongodbDatabasesCollections_SPECARM instances for property testing.
// We first initialize databaseAccountsMongodbDatabasesCollections_specarmGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsMongodbDatabasesCollections_SPECARMGenerator() gopter.Gen {
	if databaseAccountsMongodbDatabasesCollections_specarmGenerator != nil {
		return databaseAccountsMongodbDatabasesCollections_specarmGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM(generators)
	databaseAccountsMongodbDatabasesCollections_specarmGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SPECARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM(generators)
	databaseAccountsMongodbDatabasesCollections_specarmGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsMongodbDatabasesCollections_SPECARM{}), generators)

	return databaseAccountsMongodbDatabasesCollections_specarmGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsMongodbDatabasesCollections_SPECARM(gens map[string]gopter.Gen) {
	gens["Properties"] = MongoDBCollectionProperties_SpecARMGenerator()
}

func Test_MongoDBCollectionProperties_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionProperties_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionProperties_SpecARM, MongoDBCollectionProperties_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionProperties_SpecARM runs a test to see if a specific instance of MongoDBCollectionProperties_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionProperties_SpecARM(subject MongoDBCollectionProperties_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionProperties_SpecARM
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

// Generator of MongoDBCollectionProperties_SpecARM instances for property testing - lazily instantiated by
//MongoDBCollectionProperties_SpecARMGenerator()
var mongoDBCollectionProperties_specARMGenerator gopter.Gen

// MongoDBCollectionProperties_SpecARMGenerator returns a generator of MongoDBCollectionProperties_SpecARM instances for property testing.
func MongoDBCollectionProperties_SpecARMGenerator() gopter.Gen {
	if mongoDBCollectionProperties_specARMGenerator != nil {
		return mongoDBCollectionProperties_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoDBCollectionProperties_SpecARM(generators)
	mongoDBCollectionProperties_specARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionProperties_SpecARM{}), generators)

	return mongoDBCollectionProperties_specARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionProperties_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionProperties_SpecARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptions_SpecARMGenerator())
	gens["Resource"] = MongoDBCollectionResource_SpecARMGenerator()
}

func Test_CreateUpdateOptions_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions_SpecARM, CreateUpdateOptions_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions_SpecARM runs a test to see if a specific instance of CreateUpdateOptions_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions_SpecARM(subject CreateUpdateOptions_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions_SpecARM
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

// Generator of CreateUpdateOptions_SpecARM instances for property testing - lazily instantiated by
//CreateUpdateOptions_SpecARMGenerator()
var createUpdateOptions_specARMGenerator gopter.Gen

// CreateUpdateOptions_SpecARMGenerator returns a generator of CreateUpdateOptions_SpecARM instances for property testing.
// We first initialize createUpdateOptions_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptions_SpecARMGenerator() gopter.Gen {
	if createUpdateOptions_specARMGenerator != nil {
		return createUpdateOptions_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_SpecARM(generators)
	createUpdateOptions_specARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions_SpecARM(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions_SpecARM(generators)
	createUpdateOptions_specARMGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions_SpecARM{}), generators)

	return createUpdateOptions_specARMGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions_SpecARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions_SpecARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_SpecARMGenerator())
}

func Test_MongoDBCollectionResource_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBCollectionResource_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBCollectionResource_SpecARM, MongoDBCollectionResource_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBCollectionResource_SpecARM runs a test to see if a specific instance of MongoDBCollectionResource_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBCollectionResource_SpecARM(subject MongoDBCollectionResource_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBCollectionResource_SpecARM
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

// Generator of MongoDBCollectionResource_SpecARM instances for property testing - lazily instantiated by
//MongoDBCollectionResource_SpecARMGenerator()
var mongoDBCollectionResource_specARMGenerator gopter.Gen

// MongoDBCollectionResource_SpecARMGenerator returns a generator of MongoDBCollectionResource_SpecARM instances for property testing.
// We first initialize mongoDBCollectionResource_specARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBCollectionResource_SpecARMGenerator() gopter.Gen {
	if mongoDBCollectionResource_specARMGenerator != nil {
		return mongoDBCollectionResource_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_SpecARM(generators)
	mongoDBCollectionResource_specARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBCollectionResource_SpecARM(generators)
	AddRelatedPropertyGeneratorsForMongoDBCollectionResource_SpecARM(generators)
	mongoDBCollectionResource_specARMGenerator = gen.Struct(reflect.TypeOf(MongoDBCollectionResource_SpecARM{}), generators)

	return mongoDBCollectionResource_specARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBCollectionResource_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBCollectionResource_SpecARM(gens map[string]gopter.Gen) {
	gens["AnalyticalStorageTtl"] = gen.PtrOf(gen.Int())
	gens["Id"] = gen.AlphaString()
	gens["ShardKey"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBCollectionResource_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBCollectionResource_SpecARM(gens map[string]gopter.Gen) {
	gens["Indexes"] = gen.SliceOf(MongoIndex_SpecARMGenerator())
}

func Test_AutoscaleSettings_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_SpecARM, AutoscaleSettings_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_SpecARM runs a test to see if a specific instance of AutoscaleSettings_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_SpecARM(subject AutoscaleSettings_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_SpecARM
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

// Generator of AutoscaleSettings_SpecARM instances for property testing - lazily instantiated by
//AutoscaleSettings_SpecARMGenerator()
var autoscaleSettings_specARMGenerator gopter.Gen

// AutoscaleSettings_SpecARMGenerator returns a generator of AutoscaleSettings_SpecARM instances for property testing.
func AutoscaleSettings_SpecARMGenerator() gopter.Gen {
	if autoscaleSettings_specARMGenerator != nil {
		return autoscaleSettings_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_SpecARM(generators)
	autoscaleSettings_specARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_SpecARM{}), generators)

	return autoscaleSettings_specARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_SpecARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_MongoIndex_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndex_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndex_SpecARM, MongoIndex_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndex_SpecARM runs a test to see if a specific instance of MongoIndex_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndex_SpecARM(subject MongoIndex_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndex_SpecARM
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

// Generator of MongoIndex_SpecARM instances for property testing - lazily instantiated by MongoIndex_SpecARMGenerator()
var mongoIndex_specARMGenerator gopter.Gen

// MongoIndex_SpecARMGenerator returns a generator of MongoIndex_SpecARM instances for property testing.
func MongoIndex_SpecARMGenerator() gopter.Gen {
	if mongoIndex_specARMGenerator != nil {
		return mongoIndex_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongoIndex_SpecARM(generators)
	mongoIndex_specARMGenerator = gen.Struct(reflect.TypeOf(MongoIndex_SpecARM{}), generators)

	return mongoIndex_specARMGenerator
}

// AddRelatedPropertyGeneratorsForMongoIndex_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoIndex_SpecARM(gens map[string]gopter.Gen) {
	gens["Key"] = gen.PtrOf(MongoIndexKeys_SpecARMGenerator())
	gens["Options"] = gen.PtrOf(MongoIndexOptions_SpecARMGenerator())
}

func Test_MongoIndexKeys_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexKeys_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexKeys_SpecARM, MongoIndexKeys_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexKeys_SpecARM runs a test to see if a specific instance of MongoIndexKeys_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexKeys_SpecARM(subject MongoIndexKeys_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexKeys_SpecARM
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

// Generator of MongoIndexKeys_SpecARM instances for property testing - lazily instantiated by
//MongoIndexKeys_SpecARMGenerator()
var mongoIndexKeys_specARMGenerator gopter.Gen

// MongoIndexKeys_SpecARMGenerator returns a generator of MongoIndexKeys_SpecARM instances for property testing.
func MongoIndexKeys_SpecARMGenerator() gopter.Gen {
	if mongoIndexKeys_specARMGenerator != nil {
		return mongoIndexKeys_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexKeys_SpecARM(generators)
	mongoIndexKeys_specARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexKeys_SpecARM{}), generators)

	return mongoIndexKeys_specARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexKeys_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexKeys_SpecARM(gens map[string]gopter.Gen) {
	gens["Keys"] = gen.SliceOf(gen.AlphaString())
}

func Test_MongoIndexOptions_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoIndexOptions_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoIndexOptions_SpecARM, MongoIndexOptions_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoIndexOptions_SpecARM runs a test to see if a specific instance of MongoIndexOptions_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoIndexOptions_SpecARM(subject MongoIndexOptions_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoIndexOptions_SpecARM
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

// Generator of MongoIndexOptions_SpecARM instances for property testing - lazily instantiated by
//MongoIndexOptions_SpecARMGenerator()
var mongoIndexOptions_specARMGenerator gopter.Gen

// MongoIndexOptions_SpecARMGenerator returns a generator of MongoIndexOptions_SpecARM instances for property testing.
func MongoIndexOptions_SpecARMGenerator() gopter.Gen {
	if mongoIndexOptions_specARMGenerator != nil {
		return mongoIndexOptions_specARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoIndexOptions_SpecARM(generators)
	mongoIndexOptions_specARMGenerator = gen.Struct(reflect.TypeOf(MongoIndexOptions_SpecARM{}), generators)

	return mongoIndexOptions_specARMGenerator
}

// AddIndependentPropertyGeneratorsForMongoIndexOptions_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoIndexOptions_SpecARM(gens map[string]gopter.Gen) {
	gens["ExpireAfterSeconds"] = gen.PtrOf(gen.Int())
	gens["Unique"] = gen.PtrOf(gen.Bool())
}
