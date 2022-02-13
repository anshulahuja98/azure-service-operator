// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties to update Azure Cosmos DB resource throughput.
	Properties ThroughputSettingsProperties_SpecARM `json:"properties"`
	Tags       map[string]string                    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (specarm DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SPECARM) GetType() string {
	return ""
}

type ThroughputSettingsProperties_SpecARM struct {
	//Resource: The standard JSON format of a resource throughput
	Resource ThroughputSettingsResource_SpecARM `json:"resource"`
}

type ThroughputSettingsResource_SpecARM struct {
	//AutoscaleSettings: Cosmos DB resource for autoscale settings. Either throughput
	//is required or autoscaleSettings is required, but not both.
	AutoscaleSettings *AutoscaleSettingsResource_SpecARM `json:"autoscaleSettings,omitempty"`

	//Throughput: Value of the Cosmos DB resource throughput. Either throughput is
	//required or autoscaleSettings is required, but not both.
	Throughput *int `json:"throughput,omitempty"`
}

type AutoscaleSettingsResource_SpecARM struct {
	//AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy
	AutoUpgradePolicy *AutoUpgradePolicyResource_SpecARM `json:"autoUpgradePolicy,omitempty"`

	//MaxThroughput: Represents maximum throughput container can scale up to.
	MaxThroughput int `json:"maxThroughput"`
}

type AutoUpgradePolicyResource_SpecARM struct {
	//ThroughputPolicy: Represents throughput policy which service must adhere to for
	//auto-upgrade
	ThroughputPolicy *ThroughputPolicyResource_SpecARM `json:"throughputPolicy,omitempty"`
}

type ThroughputPolicyResource_SpecARM struct {
	//IncrementPercent: Represents the percentage by which throughput can increase
	//every time throughput policy kicks in.
	IncrementPercent *int `json:"incrementPercent,omitempty"`

	//IsEnabled: Determines whether the ThroughputPolicy is active or not
	IsEnabled *bool `json:"isEnabled,omitempty"`
}
