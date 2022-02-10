// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=batch.azure.com,resources=batchaccounts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch.azure.com,resources={batchaccounts/status,batchaccounts/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210101.BatchAccount
//Generator information:
//- Generated from: /batch/resource-manager/Microsoft.Batch/stable/2021-01-01/BatchManagement.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}
type BatchAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccounts_SPEC  `json:"spec,omitempty"`
	Status            BatchAccount_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &BatchAccount{}

// GetConditions returns the conditions of the resource
func (account *BatchAccount) GetConditions() conditions.Conditions {
	return account.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (account *BatchAccount) SetConditions(conditions conditions.Conditions) {
	account.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &BatchAccount{}

// AzureName returns the Azure name of the resource
func (account *BatchAccount) AzureName() string {
	return account.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01"
func (account BatchAccount) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (account *BatchAccount) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (account *BatchAccount) GetSpec() genruntime.ConvertibleSpec {
	return &account.Spec
}

// GetStatus returns the status of this resource
func (account *BatchAccount) GetStatus() genruntime.ConvertibleStatus {
	return &account.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (account *BatchAccount) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (account *BatchAccount) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BatchAccount_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (account *BatchAccount) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(account.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  account.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (account *BatchAccount) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BatchAccount_Status); ok {
		account.Status = *st
		return nil
	}

	// Convert status to required version
	var st BatchAccount_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	account.Status = st
	return nil
}

// Hub marks that this BatchAccount is the hub type for conversion
func (account *BatchAccount) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (account *BatchAccount) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: account.Spec.OriginalVersion,
		Kind:    "BatchAccount",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210101.BatchAccount
//Generator information:
//- Generated from: /batch/resource-manager/Microsoft.Batch/stable/2021-01-01/BatchManagement.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}
type BatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchAccount `json:"items"`
}

//Storage version of v1alpha1api20210101.BatchAccount_Status
type BatchAccount_Status struct {
	AutoStorage         *AutoStorageBaseProperties_Status `json:"autoStorage,omitempty"`
	Conditions          []conditions.Condition            `json:"conditions,omitempty"`
	Encryption          *EncryptionProperties_Status      `json:"encryption,omitempty"`
	Identity            *BatchAccountIdentity_Status      `json:"identity,omitempty"`
	KeyVaultReference   *KeyVaultReference_Status         `json:"keyVaultReference,omitempty"`
	Location            *string                           `json:"location,omitempty"`
	PoolAllocationMode  *string                           `json:"poolAllocationMode,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                           `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BatchAccount_Status{}

// ConvertStatusFrom populates our BatchAccount_Status from the provided source
func (account *BatchAccount_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(account)
}

// ConvertStatusTo populates the provided destination from our BatchAccount_Status
func (account *BatchAccount_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == account {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(account)
}

//Storage version of v1alpha1api20210101.BatchAccounts_SPEC
type BatchAccounts_SPEC struct {
	AutoStorage *AutoStorageBaseProperties_Spec `json:"autoStorage,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName         string                     `json:"azureName"`
	Encryption        *EncryptionProperties_Spec `json:"encryption,omitempty"`
	Identity          *BatchAccountIdentity_Spec `json:"identity,omitempty"`
	KeyVaultReference *KeyVaultReference_Spec    `json:"keyVaultReference,omitempty"`
	Location          *string                    `json:"location,omitempty"`
	OriginalVersion   string                     `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner               genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PoolAllocationMode  *string                           `json:"poolAllocationMode,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	PublicNetworkAccess *string                           `json:"publicNetworkAccess,omitempty"`
	Tags                map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &BatchAccounts_SPEC{}

// ConvertSpecFrom populates our BatchAccounts_SPEC from the provided source
func (spec *BatchAccounts_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our BatchAccounts_SPEC
func (spec *BatchAccounts_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20210101.AutoStorageBaseProperties_Spec
type AutoStorageBaseProperties_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//StorageAccountReference: The resource ID of the storage account to be used for
	//auto-storage account.
	StorageAccountReference genruntime.ResourceReference `armReference:"StorageAccountId" json:"storageAccountReference"`
}

//Storage version of v1alpha1api20210101.AutoStorageBaseProperties_Status
type AutoStorageBaseProperties_Status struct {
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageAccountId *string                `json:"storageAccountId,omitempty"`
}

//Storage version of v1alpha1api20210101.BatchAccountIdentity_Spec
type BatchAccountIdentity_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210101.BatchAccountIdentity_Status
type BatchAccountIdentity_Status struct {
	PrincipalId            *string                                                       `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                        `json:"$propertyBag,omitempty"`
	TenantId               *string                                                       `json:"tenantId,omitempty"`
	Type                   *string                                                       `json:"type,omitempty"`
	UserAssignedIdentities map[string]BatchAccountIdentity_Status_UserAssignedIdentities `json:"userAssignedIdentities,omitempty"`
}

//Storage version of v1alpha1api20210101.EncryptionProperties_Spec
type EncryptionProperties_Spec struct {
	KeySource          *string                  `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties_Spec `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101.EncryptionProperties_Status
type EncryptionProperties_Status struct {
	KeySource          *string                    `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties_Status `json:"keyVaultProperties,omitempty"`
	PropertyBag        genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101.KeyVaultReference_Spec
type KeyVaultReference_Spec struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//Reference: The resource ID of the Azure key vault associated with the Batch
	//account.
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
	Url       *string                      `json:"url,omitempty"`
}

//Storage version of v1alpha1api20210101.KeyVaultReference_Status
type KeyVaultReference_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Url         *string                `json:"url,omitempty"`
}

//Storage version of v1alpha1api20210101.BatchAccountIdentity_Status_UserAssignedIdentities
type BatchAccountIdentity_Status_UserAssignedIdentities struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101.KeyVaultProperties_Spec
type KeyVaultProperties_Spec struct {
	KeyIdentifier *string                `json:"keyIdentifier,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210101.KeyVaultProperties_Status
type KeyVaultProperties_Status struct {
	KeyIdentifier *string                `json:"keyIdentifier,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&BatchAccount{}, &BatchAccountList{})
}
