// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

type RedisLinkedServer_StatusARM struct {
	//Properties: Properties required to create a linked server.
	Properties *RedisLinkedServerProperties_StatusARM `json:"properties,omitempty"`
}

type RedisLinkedServerProperties_StatusARM struct {
	//LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.
	LinkedRedisCacheId string `json:"linkedRedisCacheId"`

	//LinkedRedisCacheLocation: Location of the linked redis cache.
	LinkedRedisCacheLocation string `json:"linkedRedisCacheLocation"`

	//ServerRole: Role of the linked server.
	ServerRole RedisLinkedServerProperties_ServerRole_Status `json:"serverRole"`
}
