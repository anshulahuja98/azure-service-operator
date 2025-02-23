// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsQueueServicesQueues_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: A queue name must be unique within a storage account and must be between 3 and 63 characters.The name must
	//comprise of lowercase alphanumeric and dash(-) characters only, it should begin and end with an alphanumeric character
	//and it cannot have two consecutive dash(-) characters.
	Name       string              `json:"name"`
	Properties *QueuePropertiesARM `json:"properties,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueServicesQueues_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queues StorageAccountsQueueServicesQueues_SpecARM) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (queues StorageAccountsQueueServicesQueues_SpecARM) GetName() string {
	return queues.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queues StorageAccountsQueueServicesQueues_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/QueueProperties
type QueuePropertiesARM struct {
	//Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
}
