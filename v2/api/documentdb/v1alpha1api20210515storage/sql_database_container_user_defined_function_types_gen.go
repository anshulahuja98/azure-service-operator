// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontaineruserdefinedfunctions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontaineruserdefinedfunctions/status,sqldatabasecontaineruserdefinedfunctions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI:
///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}
type SqlDatabaseContainerUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC `json:"spec,omitempty"`
	Status            SqlUserDefinedFunctionCreateUpdateParameters_Status             `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerUserDefinedFunction{}

// GetConditions returns the conditions of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetConditions() conditions.Conditions {
	return function.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (function *SqlDatabaseContainerUserDefinedFunction) SetConditions(conditions conditions.Conditions) {
	function.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerUserDefinedFunction{}

// AzureName returns the Azure name of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) AzureName() string {
	return function.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (function SqlDatabaseContainerUserDefinedFunction) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetSpec() genruntime.ConvertibleSpec {
	return &function.Spec
}

// GetStatus returns the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetStatus() genruntime.ConvertibleStatus {
	return &function.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (function *SqlDatabaseContainerUserDefinedFunction) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (function *SqlDatabaseContainerUserDefinedFunction) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlUserDefinedFunctionCreateUpdateParameters_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (function *SqlDatabaseContainerUserDefinedFunction) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(function.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  function.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlUserDefinedFunctionCreateUpdateParameters_Status); ok {
		function.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlUserDefinedFunctionCreateUpdateParameters_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	function.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerUserDefinedFunction is the hub type for conversion
func (function *SqlDatabaseContainerUserDefinedFunction) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (function *SqlDatabaseContainerUserDefinedFunction) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: function.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerUserDefinedFunction",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI:
///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/userDefinedFunctions/{userDefinedFunctionName}
type SqlDatabaseContainerUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerUserDefinedFunction `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC
type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string                    `json:"azureName"`
	Location        *string                   `json:"location,omitempty"`
	Options         *CreateUpdateOptions_Spec `json:"options,omitempty"`
	OriginalVersion string                    `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference    `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionResource_Spec `json:"resource,omitempty"`
	Tags        map[string]string                    `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC from the provided source
func (spec *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC
func (spec *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionCreateUpdateParameters_Status
type SqlUserDefinedFunctionCreateUpdateParameters_Status struct {
	Conditions  []conditions.Condition                 `json:"conditions,omitempty"`
	Id          *string                                `json:"id,omitempty"`
	Location    *string                                `json:"location,omitempty"`
	Name        *string                                `json:"name,omitempty"`
	Options     *CreateUpdateOptions_Status            `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                 `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionResource_Status `json:"resource,omitempty"`
	Tags        map[string]string                      `json:"tags,omitempty"`
	Type        *string                                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlUserDefinedFunctionCreateUpdateParameters_Status{}

// ConvertStatusFrom populates our SqlUserDefinedFunctionCreateUpdateParameters_Status from the provided source
func (parameters *SqlUserDefinedFunctionCreateUpdateParameters_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(parameters)
}

// ConvertStatusTo populates the provided destination from our SqlUserDefinedFunctionCreateUpdateParameters_Status
func (parameters *SqlUserDefinedFunctionCreateUpdateParameters_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == parameters {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(parameters)
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionResource_Spec
type SqlUserDefinedFunctionResource_Spec struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionResource_Status
type SqlUserDefinedFunctionResource_Status struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerUserDefinedFunction{}, &SqlDatabaseContainerUserDefinedFunctionList{})
}
