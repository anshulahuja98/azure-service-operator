// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=topics,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={topics/status,topics/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20200601.Topic
//Generator information:
//- Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Topics_SPEC  `json:"spec,omitempty"`
	Status            Topic_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Topic{}

// GetConditions returns the conditions of the resource
func (topic *Topic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *Topic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Topic{}

// AzureName returns the Azure name of the resource
func (topic *Topic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic Topic) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (topic *Topic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (topic *Topic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *Topic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (topic *Topic) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *Topic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Topic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *Topic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *Topic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Topic_Status); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st Topic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// Hub marks that this Topic is the hub type for conversion
func (topic *Topic) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *Topic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion,
		Kind:    "Topic",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20200601.Topic
//Generator information:
//- Generated from: /eventgrid/resource-manager/Microsoft.EventGrid/stable/2020-06-01/EventGrid.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Topic `json:"items"`
}

//Storage version of v1alpha1api20200601.Topic_Status
type Topic_Status struct {
	Conditions                 []conditions.Condition                                       `json:"conditions,omitempty"`
	Endpoint                   *string                                                      `json:"endpoint,omitempty"`
	Id                         *string                                                      `json:"id,omitempty"`
	InboundIpRules             []InboundIpRule_Status                                       `json:"inboundIpRules,omitempty"`
	InputSchema                *string                                                      `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_Status                                   `json:"inputSchemaMapping,omitempty"`
	Location                   *string                                                      `json:"location,omitempty"`
	MetricResourceId           *string                                                      `json:"metricResourceId,omitempty"`
	Name                       *string                                                      `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_Topic_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                       `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                      `json:"publicNetworkAccess,omitempty"`
	SystemData                 *SystemData_Status                                           `json:"systemData,omitempty"`
	Tags                       map[string]string                                            `json:"tags,omitempty"`
	Type                       *string                                                      `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Topic_Status{}

// ConvertStatusFrom populates our Topic_Status from the provided source
func (topic *Topic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(topic)
}

// ConvertStatusTo populates the provided destination from our Topic_Status
func (topic *Topic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == topic {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(topic)
}

//Storage version of v1alpha1api20200601.Topics_SPEC
type Topics_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Topics_SPEC{}

// ConvertSpecFrom populates our Topics_SPEC from the provided source
func (spec *Topics_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our Topics_SPEC
func (spec *Topics_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20200601.PrivateEndpointConnection_Status_Topic_SubResourceEmbedded
type PrivateEndpointConnection_Status_Topic_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Topic{}, &TopicList{})
}
