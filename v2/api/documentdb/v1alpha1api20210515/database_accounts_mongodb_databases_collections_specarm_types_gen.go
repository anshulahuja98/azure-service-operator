// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsMongodbDatabasesCollections_SPECARM struct {
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB MongoDB collection.
	Properties MongoDBCollectionProperties_SpecARM `json:"properties"`
	Tags       map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabasesCollections_SPECARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (specarm DatabaseAccountsMongodbDatabasesCollections_SPECARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (specarm DatabaseAccountsMongodbDatabasesCollections_SPECARM) GetName() string {
	return specarm.Name
}

// GetType returns the ARM Type of the resource. This is always ""
func (specarm DatabaseAccountsMongodbDatabasesCollections_SPECARM) GetType() string {
	return ""
}

type MongoDBCollectionProperties_SpecARM struct {
	//Options: A key-value pair of options to be applied for the request. This
	//corresponds to the headers sent with the request.
	Options *CreateUpdateOptions_SpecARM `json:"options,omitempty"`

	//Resource: The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResource_SpecARM `json:"resource"`
}

type CreateUpdateOptions_SpecARM struct {
	//AutoscaleSettings: Specifies the Autoscale settings.
	AutoscaleSettings *AutoscaleSettings_SpecARM `json:"autoscaleSettings,omitempty"`

	//Throughput: Request Units per second. For example, "throughput": 10000.
	Throughput *int `json:"throughput,omitempty"`
}

type MongoDBCollectionResource_SpecARM struct {
	//AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	//Id: Name of the Cosmos DB MongoDB collection
	Id string `json:"id"`

	//Indexes: List of index keys
	Indexes []MongoIndex_SpecARM `json:"indexes,omitempty"`

	//ShardKey: A key-value pair of shard keys to be applied for the request.
	ShardKey map[string]string `json:"shardKey,omitempty"`
}

type AutoscaleSettings_SpecARM struct {
	//MaxThroughput: Represents maximum throughput, the resource can scale up to.
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

type MongoIndex_SpecARM struct {
	//Key: Cosmos DB MongoDB collection index keys
	Key *MongoIndexKeys_SpecARM `json:"key,omitempty"`

	//Options: Cosmos DB MongoDB collection index key options
	Options *MongoIndexOptions_SpecARM `json:"options,omitempty"`
}

type MongoIndexKeys_SpecARM struct {
	//Keys: List of keys for each MongoDB collection in the Azure Cosmos DB service
	Keys []string `json:"keys,omitempty"`
}

type MongoIndexOptions_SpecARM struct {
	//ExpireAfterSeconds: Expire after seconds
	ExpireAfterSeconds *int `json:"expireAfterSeconds,omitempty"`

	//Unique: Is unique or not
	Unique *bool `json:"unique,omitempty"`
}
