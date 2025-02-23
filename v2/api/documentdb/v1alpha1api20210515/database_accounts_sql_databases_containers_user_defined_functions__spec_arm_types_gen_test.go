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

func Test_DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM, DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(subject DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM
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

// Generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM instances for property testing -
//lazily instantiated by DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator()
var databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator != nil {
		return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(generators)
	databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM{}), generators)

	return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator()
}

func Test_SqlUserDefinedFunctionCreateUpdatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionCreateUpdatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM, SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM runs a test to see if a specific instance of SqlUserDefinedFunctionCreateUpdatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionCreateUpdatePropertiesARM(subject SqlUserDefinedFunctionCreateUpdatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionCreateUpdatePropertiesARM
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

// Generator of SqlUserDefinedFunctionCreateUpdatePropertiesARM instances for property testing - lazily instantiated by
//SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator()
var sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator gopter.Gen

// SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator returns a generator of SqlUserDefinedFunctionCreateUpdatePropertiesARM instances for property testing.
func SqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator != nil {
		return sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM(generators)
	sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionCreateUpdatePropertiesARM{}), generators)

	return sqlUserDefinedFunctionCreateUpdatePropertiesARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlUserDefinedFunctionCreateUpdatePropertiesARM(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsARMGenerator())
	gens["Resource"] = SqlUserDefinedFunctionResourceARMGenerator()
}

func Test_SqlUserDefinedFunctionResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlUserDefinedFunctionResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM, SqlUserDefinedFunctionResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM runs a test to see if a specific instance of SqlUserDefinedFunctionResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlUserDefinedFunctionResourceARM(subject SqlUserDefinedFunctionResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlUserDefinedFunctionResourceARM
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

// Generator of SqlUserDefinedFunctionResourceARM instances for property testing - lazily instantiated by
//SqlUserDefinedFunctionResourceARMGenerator()
var sqlUserDefinedFunctionResourceARMGenerator gopter.Gen

// SqlUserDefinedFunctionResourceARMGenerator returns a generator of SqlUserDefinedFunctionResourceARM instances for property testing.
func SqlUserDefinedFunctionResourceARMGenerator() gopter.Gen {
	if sqlUserDefinedFunctionResourceARMGenerator != nil {
		return sqlUserDefinedFunctionResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM(generators)
	sqlUserDefinedFunctionResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlUserDefinedFunctionResourceARM{}), generators)

	return sqlUserDefinedFunctionResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlUserDefinedFunctionResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
}
