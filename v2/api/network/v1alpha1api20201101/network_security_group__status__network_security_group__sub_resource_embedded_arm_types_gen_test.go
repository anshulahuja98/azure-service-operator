// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

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

func Test_NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM, NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(subject NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
//lazily instantiated by NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormatStatusARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatStatusARM, NetworkSecurityGroupPropertiesFormatStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatStatusARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatStatusARM(subject NetworkSecurityGroupPropertiesFormat_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_StatusARM
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

// Generator of NetworkSecurityGroupPropertiesFormat_StatusARM instances for property testing - lazily instantiated by
//NetworkSecurityGroupPropertiesFormatStatusARMGenerator()
var networkSecurityGroupPropertiesFormatStatusARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormatStatusARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_StatusARM instances for property testing.
// We first initialize networkSecurityGroupPropertiesFormatStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupPropertiesFormatStatusARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormatStatusARMGenerator != nil {
		return networkSecurityGroupPropertiesFormatStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM(generators)
	networkSecurityGroupPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM(generators)
	networkSecurityGroupPropertiesFormatStatusARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_StatusARM{}), generators)

	return networkSecurityGroupPropertiesFormatStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatStatusARM(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLogStatusSubResourceEmbeddedARMGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["Subnets"] = gen.SliceOf(SubnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
}

func Test_FlowLog_Status_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_Status_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLogStatusSubResourceEmbeddedARM, FlowLogStatusSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLogStatusSubResourceEmbeddedARM runs a test to see if a specific instance of FlowLog_Status_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLogStatusSubResourceEmbeddedARM(subject FlowLog_Status_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_Status_SubResourceEmbeddedARM
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

// Generator of FlowLog_Status_SubResourceEmbeddedARM instances for property testing - lazily instantiated by
//FlowLogStatusSubResourceEmbeddedARMGenerator()
var flowLogStatusSubResourceEmbeddedARMGenerator gopter.Gen

// FlowLogStatusSubResourceEmbeddedARMGenerator returns a generator of FlowLog_Status_SubResourceEmbeddedARM instances for property testing.
func FlowLogStatusSubResourceEmbeddedARMGenerator() gopter.Gen {
	if flowLogStatusSubResourceEmbeddedARMGenerator != nil {
		return flowLogStatusSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbeddedARM(generators)
	flowLogStatusSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(FlowLog_Status_SubResourceEmbeddedARM{}), generators)

	return flowLogStatusSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLogStatusSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM, NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(subject NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
//lazily instantiated by NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfaceStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusARMGenerator())
}

func Test_SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM, SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM(subject SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
//instantiated by SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return securityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM, SubnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM(subject Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM
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

// Generator of Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
//instantiated by SubnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var subnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// SubnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SubnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if subnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return subnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	subnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(Subnet_Status_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return subnetStatusNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetStatusNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
