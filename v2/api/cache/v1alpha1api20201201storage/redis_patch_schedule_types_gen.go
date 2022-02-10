// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redispatchschedules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redispatchschedules/status,redispatchschedules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201201.RedisPatchSchedule
//Generator information:
//- Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_SPEC  `json:"spec,omitempty"`
	Status            RedisPatchSchedule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (schedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return schedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (schedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	schedule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource
func (schedule *RedisPatchSchedule) AzureName() string {
	return schedule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (schedule *RedisPatchSchedule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (schedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &schedule.Spec
}

// GetStatus returns the status of this resource
func (schedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &schedule.Status
}

// GetType returns the ARM Type of the resource. This is always ""
func (schedule *RedisPatchSchedule) GetType() string {
	return ""
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  schedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (schedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_Status); ok {
		schedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	schedule.Status = st
	return nil
}

// Hub marks that this RedisPatchSchedule is the hub type for conversion
func (schedule *RedisPatchSchedule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion,
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201201.RedisPatchSchedule
//Generator information:
//- Generated from: /redis/resource-manager/Microsoft.Cache/stable/2020-12-01/redis.json
//- ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

//Storage version of v1alpha1api20201201.RedisPatchSchedule_Status
type RedisPatchSchedule_Status struct {
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_Status `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_Status{}

// ConvertStatusFrom populates our RedisPatchSchedule_Status from the provided source
func (schedule *RedisPatchSchedule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(schedule)
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(schedule)
}

//Storage version of v1alpha1api20201201.RedisPatchSchedules_SPEC
type RedisPatchSchedules_SPEC struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string `json:"azureName"`
	OriginalVersion string `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner           genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_Spec              `json:"scheduleEntries,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_SPEC{}

// ConvertSpecFrom populates our RedisPatchSchedules_SPEC from the provided source
func (spec *RedisPatchSchedules_SPEC) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(spec)
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_SPEC
func (spec *RedisPatchSchedules_SPEC) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == spec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(spec)
}

//Storage version of v1alpha1api20201201.ScheduleEntry_Spec
type ScheduleEntry_Spec struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

//Storage version of v1alpha1api20201201.ScheduleEntry_Status
type ScheduleEntry_Status struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
