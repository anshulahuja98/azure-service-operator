// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
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

func Test_SqlDatabaseContainerStoredProcedure_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseContainerStoredProcedure to SqlDatabaseContainerStoredProcedure via AssignPropertiesToSqlDatabaseContainerStoredProcedure & AssignPropertiesFromSqlDatabaseContainerStoredProcedure returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure tests if a specific instance of SqlDatabaseContainerStoredProcedure can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseContainerStoredProcedure
	err := copied.AssignPropertiesToSqlDatabaseContainerStoredProcedure(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseContainerStoredProcedure
	err = actual.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(&other)
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

func Test_SqlDatabaseContainerStoredProcedure_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerStoredProcedure via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure, SqlDatabaseContainerStoredProcedureGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure runs a test to see if a specific instance of SqlDatabaseContainerStoredProcedure round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerStoredProcedure(subject SqlDatabaseContainerStoredProcedure) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerStoredProcedure
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

// Generator of SqlDatabaseContainerStoredProcedure instances for property testing - lazily instantiated by
//SqlDatabaseContainerStoredProcedureGenerator()
var sqlDatabaseContainerStoredProcedureGenerator gopter.Gen

// SqlDatabaseContainerStoredProcedureGenerator returns a generator of SqlDatabaseContainerStoredProcedure instances for property testing.
func SqlDatabaseContainerStoredProcedureGenerator() gopter.Gen {
	if sqlDatabaseContainerStoredProcedureGenerator != nil {
		return sqlDatabaseContainerStoredProcedureGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(generators)
	sqlDatabaseContainerStoredProcedureGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerStoredProcedure{}), generators)

	return sqlDatabaseContainerStoredProcedureGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerStoredProcedure(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()
	gens["Status"] = SqlStoredProcedureGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec to DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec, DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec tests if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err := copied.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&other)
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

func Test_DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec, DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(subject DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
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

// Generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator()
var databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator != nil {
		return databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(generators)
	databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}), generators)

	return databaseAccountsSqlDatabasesContainersStoredProceduresSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = SqlStoredProcedureResourceGenerator()
}

func Test_SqlStoredProcedureGetResults_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetResults_Status to SqlStoredProcedureGetResults_Status via AssignPropertiesToSqlStoredProcedureGetResultsStatus & AssignPropertiesFromSqlStoredProcedureGetResultsStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetResultsStatus, SqlStoredProcedureGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetResultsStatus tests if a specific instance of SqlStoredProcedureGetResults_Status can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetResultsStatus(subject SqlStoredProcedureGetResults_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlStoredProcedureGetResults_Status
	err := copied.AssignPropertiesToSqlStoredProcedureGetResultsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetResults_Status
	err = actual.AssignPropertiesFromSqlStoredProcedureGetResultsStatus(&other)
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

func Test_SqlStoredProcedureGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetResultsStatus, SqlStoredProcedureGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetResultsStatus runs a test to see if a specific instance of SqlStoredProcedureGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetResultsStatus(subject SqlStoredProcedureGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetResults_Status
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

// Generator of SqlStoredProcedureGetResults_Status instances for property testing - lazily instantiated by
//SqlStoredProcedureGetResultsStatusGenerator()
var sqlStoredProcedureGetResultsStatusGenerator gopter.Gen

// SqlStoredProcedureGetResultsStatusGenerator returns a generator of SqlStoredProcedureGetResults_Status instances for property testing.
// We first initialize sqlStoredProcedureGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlStoredProcedureGetResultsStatusGenerator() gopter.Gen {
	if sqlStoredProcedureGetResultsStatusGenerator != nil {
		return sqlStoredProcedureGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsStatus(generators)
	sqlStoredProcedureGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsStatus(generators)
	sqlStoredProcedureGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetResults_Status{}), generators)

	return sqlStoredProcedureGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlStoredProcedureGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlStoredProcedureGetPropertiesStatusResourceGenerator())
}

func Test_SqlStoredProcedureGetProperties_Status_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureGetProperties_Status_Resource to SqlStoredProcedureGetProperties_Status_Resource via AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource & AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureGetPropertiesStatusResource, SqlStoredProcedureGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureGetPropertiesStatusResource tests if a specific instance of SqlStoredProcedureGetProperties_Status_Resource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureGetPropertiesStatusResource(subject SqlStoredProcedureGetProperties_Status_Resource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlStoredProcedureGetProperties_Status_Resource
	err := copied.AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureGetProperties_Status_Resource
	err = actual.AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource(&other)
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

func Test_SqlStoredProcedureGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureGetPropertiesStatusResource, SqlStoredProcedureGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureGetPropertiesStatusResource runs a test to see if a specific instance of SqlStoredProcedureGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureGetPropertiesStatusResource(subject SqlStoredProcedureGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureGetProperties_Status_Resource
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

// Generator of SqlStoredProcedureGetProperties_Status_Resource instances for property testing - lazily instantiated by
//SqlStoredProcedureGetPropertiesStatusResourceGenerator()
var sqlStoredProcedureGetPropertiesStatusResourceGenerator gopter.Gen

// SqlStoredProcedureGetPropertiesStatusResourceGenerator returns a generator of SqlStoredProcedureGetProperties_Status_Resource instances for property testing.
func SqlStoredProcedureGetPropertiesStatusResourceGenerator() gopter.Gen {
	if sqlStoredProcedureGetPropertiesStatusResourceGenerator != nil {
		return sqlStoredProcedureGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesStatusResource(generators)
	sqlStoredProcedureGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureGetProperties_Status_Resource{}), generators)

	return sqlStoredProcedureGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

func Test_SqlStoredProcedureResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlStoredProcedureResource to SqlStoredProcedureResource via AssignPropertiesToSqlStoredProcedureResource & AssignPropertiesFromSqlStoredProcedureResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlStoredProcedureResource tests if a specific instance of SqlStoredProcedureResource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlStoredProcedureResource
	err := copied.AssignPropertiesToSqlStoredProcedureResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlStoredProcedureResource
	err = actual.AssignPropertiesFromSqlStoredProcedureResource(&other)
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

func Test_SqlStoredProcedureResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlStoredProcedureResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlStoredProcedureResource, SqlStoredProcedureResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlStoredProcedureResource runs a test to see if a specific instance of SqlStoredProcedureResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlStoredProcedureResource(subject SqlStoredProcedureResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlStoredProcedureResource
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

// Generator of SqlStoredProcedureResource instances for property testing - lazily instantiated by
//SqlStoredProcedureResourceGenerator()
var sqlStoredProcedureResourceGenerator gopter.Gen

// SqlStoredProcedureResourceGenerator returns a generator of SqlStoredProcedureResource instances for property testing.
func SqlStoredProcedureResourceGenerator() gopter.Gen {
	if sqlStoredProcedureResourceGenerator != nil {
		return sqlStoredProcedureResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(generators)
	sqlStoredProcedureResourceGenerator = gen.Struct(reflect.TypeOf(SqlStoredProcedureResource{}), generators)

	return sqlStoredProcedureResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlStoredProcedureResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlStoredProcedureResource(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}
