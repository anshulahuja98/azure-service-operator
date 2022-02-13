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

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=mongodbdatabasecollections,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={mongodbdatabasecollections/status,mongodbdatabasecollections/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.MongodbDatabaseCollection
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesCollections_SPEC `json:"spec,omitempty"`
	Status            MongoDBCollection_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollection{}

// GetConditions returns the conditions of the resource
func (collection *MongodbDatabaseCollection) GetConditions() conditions.Conditions {
	return collection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (collection *MongodbDatabaseCollection) SetConditions(conditions conditions.Conditions) {
	collection.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollection{}

// AzureName returns the Azure name of the resource
func (collection *MongodbDatabaseCollection) AzureName() string {
	return collection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (collection MongodbDatabaseCollection) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (collection *MongodbDatabaseCollection) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (collection *MongodbDatabaseCollection) GetSpec() genruntime.ConvertibleSpec {
	return &collection.Spec
}

// GetStatus returns the status of this resource
func (collection *MongodbDatabaseCollection) GetStatus() genruntime.ConvertibleStatus {
	return &collection.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (collection *MongodbDatabaseCollection) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (collection *MongodbDatabaseCollection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongoDBCollection_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (collection *MongodbDatabaseCollection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(collection.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  collection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (collection *MongodbDatabaseCollection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongoDBCollection_Status); ok {
		collection.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongoDBCollection_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	collection.Status = st
	return nil
}

// Hub marks that this MongodbDatabaseCollection is the hub type for conversion
func (collection *MongodbDatabaseCollection) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (collection *MongodbDatabaseCollection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: collection.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollection",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabaseCollection
//Generator information:
//- Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}
type MongodbDatabaseCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollection `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabasesCollections_SPEC
type DatabaseAccountsMongodbDatabasesCollections_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string                    `json:"azureName"`
	Location        *string                   `json:"location,omitempty"`
	Options         *CreateUpdateOptions_Spec `json:"options,omitempty"`
	OriginalVersion string                    `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionResource_Spec   `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesCollections_SPEC{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesCollections_SPEC from the provided source
func (spec *DatabaseAccountsMongodbDatabasesCollections_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollections_SPEC
func (spec *DatabaseAccountsMongodbDatabasesCollections_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210515.MongoDBCollection_Status
type MongoDBCollection_Status struct {
	Conditions  []conditions.Condition            `json:"conditions,omitempty"`
	Id          *string                           `json:"id,omitempty"`
	Location    *string                           `json:"location,omitempty"`
	Name        *string                           `json:"name,omitempty"`
	Options     *CreateUpdateOptions_Status       `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionResource_Status `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
	Type        *string                           `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongoDBCollection_Status{}

// ConvertStatusFrom populates our MongoDBCollection_Status from the provided source
func (collection *MongoDBCollection_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(collection)
}

// ConvertStatusTo populates the provided destination from our MongoDBCollection_Status
func (collection *MongoDBCollection_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == collection {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(collection)
}

//Storage version of v1alpha1api20210515.MongoDBCollectionResource_Spec
type MongoDBCollectionResource_Spec struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex_Spec      `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBCollectionResource_Status
type MongoDBCollectionResource_Status struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex_Status    `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndex_Spec
type MongoIndex_Spec struct {
	Key         *MongoIndexKeys_Spec    `json:"key,omitempty"`
	Options     *MongoIndexOptions_Spec `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndex_Status
type MongoIndex_Status struct {
	Key         *MongoIndexKeys_Status    `json:"key,omitempty"`
	Options     *MongoIndexOptions_Status `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexKeys_Spec
type MongoIndexKeys_Spec struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexKeys_Status
type MongoIndexKeys_Status struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexOptions_Spec
type MongoIndexOptions_Spec struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexOptions_Status
type MongoIndexOptions_Status struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollection{}, &MongodbDatabaseCollectionList{})
}
