// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=operationalinsights.azure.com,resources=workspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operationalinsights.azure.com,resources={workspaces/status,workspaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210601.Workspace
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/resourceDefinitions/workspaces
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Workspaces_Spec  `json:"spec,omitempty"`
	Status            Workspace_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Workspace{}

// GetConditions returns the conditions of the resource
func (workspace *Workspace) GetConditions() conditions.Conditions {
	return workspace.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (workspace *Workspace) SetConditions(conditions conditions.Conditions) {
	workspace.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Workspace{}

// AzureName returns the Azure name of the resource
func (workspace *Workspace) AzureName() string {
	return workspace.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (workspace *Workspace) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (workspace *Workspace) GetSpec() genruntime.ConvertibleSpec {
	return &workspace.Spec
}

// GetStatus returns the status of this resource
func (workspace *Workspace) GetStatus() genruntime.ConvertibleStatus {
	return &workspace.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspace *Workspace) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (workspace *Workspace) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Workspace_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (workspace *Workspace) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(workspace.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  workspace.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (workspace *Workspace) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Workspace_Status); ok {
		workspace.Status = *st
		return nil
	}

	// Convert status to required version
	var st Workspace_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	workspace.Status = st
	return nil
}

// Hub marks that this Workspace is the hub type for conversion
func (workspace *Workspace) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (workspace *Workspace) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: workspace.Spec.OriginalVersion,
		Kind:    "Workspace",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210601.Workspace
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/resourceDefinitions/workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

//Storage version of v1alpha1api20210601.Workspace_Status
type Workspace_Status struct {
	Conditions                      []conditions.Condition             `json:"conditions,omitempty"`
	CreatedDate                     *string                            `json:"createdDate,omitempty"`
	CustomerId                      *string                            `json:"customerId,omitempty"`
	ETag                            *string                            `json:"eTag,omitempty"`
	Features                        *WorkspaceFeatures_Status          `json:"features,omitempty"`
	ForceCmkForQuery                *bool                              `json:"forceCmkForQuery,omitempty"`
	Id                              *string                            `json:"id,omitempty"`
	Location                        *string                            `json:"location,omitempty"`
	ModifiedDate                    *string                            `json:"modifiedDate,omitempty"`
	Name                            *string                            `json:"name,omitempty"`
	PrivateLinkScopedResources      []PrivateLinkScopedResource_Status `json:"privateLinkScopedResources,omitempty"`
	PropertyBag                     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                            `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *string                            `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                            `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                               `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSku_Status               `json:"sku,omitempty"`
	Tags                            map[string]string                  `json:"tags,omitempty"`
	Type                            *string                            `json:"type,omitempty"`
	WorkspaceCapping                *WorkspaceCapping_Status           `json:"workspaceCapping,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Workspace_Status{}

// ConvertStatusFrom populates our Workspace_Status from the provided source
func (workspace *Workspace_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(workspace)
}

// ConvertStatusTo populates the provided destination from our Workspace_Status
func (workspace *Workspace_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == workspace {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(workspace)
}

//Storage version of v1alpha1api20210601.Workspaces_Spec
type Workspaces_Spec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=4
	// +kubebuilder:validation:Pattern="^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$"
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName        string             `json:"azureName"`
	ETag             *string            `json:"eTag,omitempty"`
	Features         *WorkspaceFeatures `json:"features,omitempty"`
	ForceCmkForQuery *bool              `json:"forceCmkForQuery,omitempty"`
	Location         *string            `json:"location,omitempty"`
	OriginalVersion  string             `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                           genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag                     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ProvisioningState               *string                           `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *string                           `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *string                           `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                              `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSku                     `json:"sku,omitempty"`
	Tags                            map[string]string                 `json:"tags,omitempty"`
	WorkspaceCapping                *WorkspaceCapping                 `json:"workspaceCapping,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Workspaces_Spec{}

// ConvertSpecFrom populates our Workspaces_Spec from the provided source
func (workspaces *Workspaces_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == workspaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(workspaces)
}

// ConvertSpecTo populates the provided destination from our Workspaces_Spec
func (workspaces *Workspaces_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == workspaces {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(workspaces)
}

//Storage version of v1alpha1api20210601.PrivateLinkScopedResource_Status
type PrivateLinkScopedResource_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
	ScopeId     *string                `json:"scopeId,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceCapping
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceCapping
type WorkspaceCapping struct {
	DailyQuotaGb *float64               `json:"dailyQuotaGb,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceCapping_Status
type WorkspaceCapping_Status struct {
	DailyQuotaGb        *float64               `json:"dailyQuotaGb,omitempty"`
	DataIngestionStatus *string                `json:"dataIngestionStatus,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QuotaNextResetTime  *string                `json:"quotaNextResetTime,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceFeatures
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceFeatures
type WorkspaceFeatures struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	//ClusterResourceReference: Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceReference                    *genruntime.ResourceReference `armReference:"ClusterResourceId" json:"clusterResourceReference,omitempty"`
	DisableLocalAuth                            *bool                         `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool                         `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool                         `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool                         `json:"immediatePurgeDataOn30Days,omitempty"`
	PropertyBag                                 genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceFeatures_Status
type WorkspaceFeatures_Status struct {
	ClusterResourceId                           *string                `json:"clusterResourceId,omitempty"`
	DisableLocalAuth                            *bool                  `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool                  `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool                  `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool                  `json:"immediatePurgeDataOn30Days,omitempty"`
	PropertyBag                                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceSku
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.OperationalInsights.json#/definitions/WorkspaceSku
type WorkspaceSku struct {
	CapacityReservationLevel *int                   `json:"capacityReservationLevel,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.WorkspaceSku_Status
type WorkspaceSku_Status struct {
	CapacityReservationLevel *int                   `json:"capacityReservationLevel,omitempty"`
	LastSkuUpdate            *string                `json:"lastSkuUpdate,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
