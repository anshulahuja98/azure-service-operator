// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

type Database_StatusARM struct {
	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: The properties of a database.
	Properties *DatabaseProperties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system metadata relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type DatabaseProperties_StatusARM struct {
	//Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	//Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}
