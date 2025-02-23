// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	dbformysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501"
	"github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersDatabaseExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersDatabaseExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&dbformysql.FlexibleServersDatabase{},
		&v1alpha1api20210501storage.FlexibleServersDatabase{}}
}
