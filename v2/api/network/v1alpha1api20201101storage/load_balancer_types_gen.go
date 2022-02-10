// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=loadbalancers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={loadbalancers/status,loadbalancers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.LoadBalancer
//Generator information:
//- Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/loadBalancer.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancers_SPEC  `json:"spec,omitempty"`
	Status            LoadBalancer_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &LoadBalancer{}

// GetConditions returns the conditions of the resource
func (balancer *LoadBalancer) GetConditions() conditions.Conditions {
	return balancer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (balancer *LoadBalancer) SetConditions(conditions conditions.Conditions) {
	balancer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &LoadBalancer{}

// AzureName returns the Azure name of the resource
func (balancer *LoadBalancer) AzureName() string {
	return balancer.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (balancer LoadBalancer) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (balancer *LoadBalancer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (balancer *LoadBalancer) GetSpec() genruntime.ConvertibleSpec {
	return &balancer.Spec
}

// GetStatus returns the status of this resource
func (balancer *LoadBalancer) GetStatus() genruntime.ConvertibleStatus {
	return &balancer.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (balancer *LoadBalancer) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (balancer *LoadBalancer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &LoadBalancer_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (balancer *LoadBalancer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(balancer.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  balancer.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (balancer *LoadBalancer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*LoadBalancer_Status); ok {
		balancer.Status = *st
		return nil
	}

	// Convert status to required version
	var st LoadBalancer_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	balancer.Status = st
	return nil
}

// Hub marks that this LoadBalancer is the hub type for conversion
func (balancer *LoadBalancer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (balancer *LoadBalancer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: balancer.Spec.OriginalVersion,
		Kind:    "LoadBalancer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.LoadBalancer
//Generator information:
//- Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/loadBalancer.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

//Storage version of v1alpha1api20201101.LoadBalancer_Status
type LoadBalancer_Status struct {
	BackendAddressPools      []BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	Conditions               []conditions.Condition                                            `json:"conditions,omitempty"`
	Etag                     *string                                                           `json:"etag,omitempty"`
	ExtendedLocation         *ExtendedLocation_Status                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                                                           `json:"id,omitempty"`
	InboundNatPools          []InboundNatPool_Status                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_Status_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_Status                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                           `json:"location,omitempty"`
	Name                     *string                                                           `json:"name,omitempty"`
	OutboundRules            []OutboundRule_Status                                             `json:"outboundRules,omitempty"`
	Probes                   []Probe_Status                                                    `json:"probes,omitempty"`
	PropertyBag              genruntime.PropertyBag                                            `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                                           `json:"provisioningState,omitempty"`
	ResourceGuid             *string                                                           `json:"resourceGuid,omitempty"`
	Sku                      *LoadBalancerSku_Status                                           `json:"sku,omitempty"`
	Tags                     map[string]string                                                 `json:"tags,omitempty"`
	Type                     *string                                                           `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &LoadBalancer_Status{}

// ConvertStatusFrom populates our LoadBalancer_Status from the provided source
func (balancer *LoadBalancer_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(balancer)
}

// ConvertStatusTo populates the provided destination from our LoadBalancer_Status
func (balancer *LoadBalancer_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(balancer)
}

//Storage version of v1alpha1api20201101.LoadBalancers_SPEC
type LoadBalancers_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                string                                                          `json:"azureName"`
	BackendAddressPools      []BackendAddressPool_Spec_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	ExtendedLocation         *ExtendedLocation_Spec                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_Spec_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []InboundNatPool_Spec                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_Spec_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_Spec                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                         `json:"location,omitempty"`
	OriginalVersion          string                                                          `json:"originalVersion"`
	OutboundRules            []OutboundRule_Spec                                             `json:"outboundRules,omitempty"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	Probes      []Probe_Spec                      `json:"probes,omitempty"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	Sku       *LoadBalancerSku_Spec         `json:"sku,omitempty"`
	Tags      map[string]string             `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &LoadBalancers_SPEC{}

// ConvertSpecFrom populates our LoadBalancers_SPEC from the provided source
func (spec *LoadBalancers_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our LoadBalancers_SPEC
func (spec *LoadBalancers_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20201101.BackendAddressPool_Spec_LoadBalancer_SubResourceEmbedded
type BackendAddressPool_Spec_LoadBalancer_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded
type BackendAddressPool_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ExtendedLocation_Spec
type ExtendedLocation_Spec struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.ExtendedLocation_Status
type ExtendedLocation_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.FrontendIPConfiguration_Spec_LoadBalancer_SubResourceEmbedded
type FrontendIPConfiguration_Spec_LoadBalancer_SubResourceEmbedded struct {
	Name                      *string                `json:"name,omitempty"`
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress           *PublicIPAddress_Spec  `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_Spec      `json:"publicIPPrefix,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference                 `armReference:"Id" json:"reference,omitempty"`
	Subnet    *Subnet_Spec_LoadBalancer_SubResourceEmbedded `json:"subnet,omitempty"`
	Zones     []string                                      `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded
type FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbedded struct {
	Etag                      *string                                                  `json:"etag,omitempty"`
	Id                        *string                                                  `json:"id,omitempty"`
	InboundNatPools           []SubResource_Status                                     `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_Status                                     `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_Status                                     `json:"loadBalancingRules,omitempty"`
	Name                      *string                                                  `json:"name,omitempty"`
	OutboundRules             []SubResource_Status                                     `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                  `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                                  `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                                  `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                   `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                  `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_Status                                      `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_Status_LoadBalancer_SubResourceEmbedded          `json:"subnet,omitempty"`
	Type                      *string                                                  `json:"type,omitempty"`
	Zones                     []string                                                 `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatPool_Spec
type InboundNatPool_Spec struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource_Spec      `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatPool_Status
type InboundNatPool_Status struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_Status    `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatRule_Spec_LoadBalancer_SubResourceEmbedded
type InboundNatRule_Spec_LoadBalancer_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatRule_Status_LoadBalancer_SubResourceEmbedded
type InboundNatRule_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancerSku_Spec
type LoadBalancerSku_Spec struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancerSku_Status
type LoadBalancerSku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancingRule_Spec
type LoadBalancingRule_Spec struct {
	BackendAddressPool      *SubResource_Spec      `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource_Spec      `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource_Spec      `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.LoadBalancingRule_Status
type LoadBalancingRule_Status struct {
	BackendAddressPool      *SubResource_Status    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_Status    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource_Status    `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.OutboundRule_Spec
type OutboundRule_Spec struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_Spec      `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResource_Spec     `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.OutboundRule_Status
type OutboundRule_Status struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_Status    `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	Etag                     *string                `json:"etag,omitempty"`
	FrontendIPConfigurations []SubResource_Status   `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.Probe_Spec
type Probe_Spec struct {
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NumberOfProbes    *int                   `json:"numberOfProbes,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol          *string                `json:"protocol,omitempty"`

	//Reference: Resource ID.
	Reference   *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	RequestPath *string                       `json:"requestPath,omitempty"`
}

//Storage version of v1alpha1api20201101.Probe_Status
type Probe_Status struct {
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IntervalInSeconds  *int                   `json:"intervalInSeconds,omitempty"`
	LoadBalancingRules []SubResource_Status   `json:"loadBalancingRules,omitempty"`
	Name               *string                `json:"name,omitempty"`
	NumberOfProbes     *int                   `json:"numberOfProbes,omitempty"`
	Port               *int                   `json:"port,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol           *string                `json:"protocol,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	RequestPath        *string                `json:"requestPath,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Spec
type PublicIPAddress_Spec struct {
	DdosSettings             *DdosSettings_Spec               `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettings_Spec `json:"dnsSettings,omitempty"`
	ExtendedLocation         *ExtendedLocation_Spec           `json:"extendedLocation,omitempty"`
	IdleTimeoutInMinutes     *int                             `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                          `json:"ipAddress,omitempty"`
	IpTags                   []IpTag_Spec                     `json:"ipTags,omitempty"`
	LinkedPublicIPAddress    *PublicIPAddress_Spec            `json:"linkedPublicIPAddress,omitempty"`
	Location                 *string                          `json:"location,omitempty"`
	MigrationPhase           *string                          `json:"migrationPhase,omitempty"`
	NatGateway               *NatGateway_Spec                 `json:"natGateway,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	PublicIPAddressVersion   *string                          `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *string                          `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_Spec                `json:"publicIPPrefix,omitempty"`

	//Reference: Resource ID.
	Reference              *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	ServicePublicIPAddress *PublicIPAddress_Spec         `json:"servicePublicIPAddress,omitempty"`
	Sku                    *PublicIPAddressSku_Spec      `json:"sku,omitempty"`
	Tags                   map[string]string             `json:"tags,omitempty"`
	Zones                  []string                      `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded
type PublicIPAddress_Status_LoadBalancer_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_Status `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Spec_LoadBalancer_SubResourceEmbedded
type Subnet_Spec_LoadBalancer_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_LoadBalancer_SubResourceEmbedded
type Subnet_Status_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
