// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

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

func Test_SqlDatabaseContainerThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseContainerThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting, SqlDatabaseContainerThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting runs a test to see if a specific instance of SqlDatabaseContainerThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseContainerThroughputSetting(subject SqlDatabaseContainerThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseContainerThroughputSetting
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

// Generator of SqlDatabaseContainerThroughputSetting instances for property testing - lazily instantiated by
//SqlDatabaseContainerThroughputSettingGenerator()
var sqlDatabaseContainerThroughputSettingGenerator gopter.Gen

// SqlDatabaseContainerThroughputSettingGenerator returns a generator of SqlDatabaseContainerThroughputSetting instances for property testing.
func SqlDatabaseContainerThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseContainerThroughputSettingGenerator != nil {
		return sqlDatabaseContainerThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(generators)
	sqlDatabaseContainerThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseContainerThroughputSetting{}), generators)

	return sqlDatabaseContainerThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseContainerThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPECGenerator()
	gens["Status"] = ThroughputSettings_StatusGenerator()
}

func Test_DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC, DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPECGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC runs a test to see if a specific instance of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(subject DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC
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

// Generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC instances for property testing - lazily
//instantiated by DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPECGenerator()
var databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPECGenerator returns a generator of DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC instances for property testing.
// We first initialize databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPECGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator != nil {
		return databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(generators)
	databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC{}), generators)

	return databaseAccountsSqlDatabasesContainersThroughputSettings_specGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesContainersThroughputSettings_SPEC(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResource_SpecGenerator())
}
