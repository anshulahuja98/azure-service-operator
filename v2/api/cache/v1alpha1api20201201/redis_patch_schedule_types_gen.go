// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_Spec  `json:"spec,omitempty"`
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

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20201201storage.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesFromRedisPatchSchedule(source)
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20201201storage.RedisPatchSchedule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisPatchSchedule but received %T instead", hub)
	}

	return schedule.AssignPropertiesToRedisPatchSchedule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisPatchSchedule{}

// Default applies defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) Default() {
	schedule.defaultImpl()
	var temp interface{} = schedule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) defaultImpl() {}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (schedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return "2020-12-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
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

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisPatchSchedule{}

// ValidateCreate validates the creation of the resource
func (schedule *RedisPatchSchedule) ValidateCreate() error {
	validations := schedule.createValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (schedule *RedisPatchSchedule) ValidateDelete() error {
	validations := schedule.deleteValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (schedule *RedisPatchSchedule) ValidateUpdate(old runtime.Object) error {
	validations := schedule.updateValidations()
	var temp interface{} = schedule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (schedule *RedisPatchSchedule) createValidations() []func() error {
	return []func() error{schedule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (schedule *RedisPatchSchedule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (schedule *RedisPatchSchedule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return schedule.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (schedule *RedisPatchSchedule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&schedule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesFromRedisPatchSchedule(source *v1alpha1api20201201storage.RedisPatchSchedule) error {

	// ObjectMeta
	schedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisPatchSchedules_Spec
	err := spec.AssignPropertiesFromRedisPatchSchedulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedulesSpec() to populate field Spec")
	}
	schedule.Spec = spec

	// Status
	var status RedisPatchSchedule_Status
	err = status.AssignPropertiesFromRedisPatchScheduleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchScheduleStatus() to populate field Status")
	}
	schedule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesToRedisPatchSchedule(destination *v1alpha1api20201201storage.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *schedule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20201201storage.RedisPatchSchedules_Spec
	err := schedule.Spec.AssignPropertiesToRedisPatchSchedulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201201storage.RedisPatchSchedule_Status
	err = schedule.Status.AssignPropertiesToRedisPatchScheduleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchScheduleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion(),
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

type RedisPatchSchedule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry_Status `json:"scheduleEntries,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_Status{}

// ConvertStatusFrom populates our RedisPatchSchedule_Status from the provided source
func (schedule *RedisPatchSchedule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisPatchSchedule_Status)
	if ok {
		// Populate our instance from source
		return schedule.AssignPropertiesFromRedisPatchScheduleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisPatchSchedule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = schedule.AssignPropertiesFromRedisPatchScheduleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisPatchSchedule_Status)
	if ok {
		// Populate destination from our instance
		return schedule.AssignPropertiesToRedisPatchScheduleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisPatchSchedule_Status{}
	err := schedule.AssignPropertiesToRedisPatchScheduleStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &RedisPatchSchedule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedule *RedisPatchSchedule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedule *RedisPatchSchedule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		schedule.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		schedule.Name = &name
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry_Status
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			schedule.ScheduleEntries = append(schedule.ScheduleEntries, item1)
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		schedule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisPatchScheduleStatus populates our RedisPatchSchedule_Status from the provided source RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) AssignPropertiesFromRedisPatchScheduleStatus(source *v1alpha1api20201201storage.RedisPatchSchedule_Status) error {

	// Conditions
	schedule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	schedule.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_Status, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_Status
			err := scheduleEntry.AssignPropertiesFromScheduleEntryStatus(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchScheduleStatus populates the provided destination RedisPatchSchedule_Status from our RedisPatchSchedule_Status
func (schedule *RedisPatchSchedule_Status) AssignPropertiesToRedisPatchScheduleStatus(destination *v1alpha1api20201201storage.RedisPatchSchedule_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(schedule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(schedule.Name)

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]v1alpha1api20201201storage.ScheduleEntry_Status, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v1alpha1api20201201storage.ScheduleEntry_Status
			err := scheduleEntryItem.AssignPropertiesToScheduleEntryStatus(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntryStatus() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2020-12-01"}
type RedisPatchSchedulesSpecAPIVersion string

const RedisPatchSchedulesSpecAPIVersion20201201 = RedisPatchSchedulesSpecAPIVersion("2020-12-01")

type RedisPatchSchedules_Spec struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`

	// +kubebuilder:validation:Required
	//ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `json:"scheduleEntries"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisPatchSchedules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (schedules *RedisPatchSchedules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if schedules == nil {
		return nil, nil
	}
	var result RedisPatchSchedules_SpecARM

	// Set property ‘Location’:
	if schedules.Location != nil {
		location := *schedules.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	for _, item := range schedules.ScheduleEntries {
		itemARM, err := item.ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		result.Properties.ScheduleEntries = append(result.Properties.ScheduleEntries, itemARM.(ScheduleEntryARM))
	}

	// Set property ‘Tags’:
	if schedules.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range schedules.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedules *RedisPatchSchedules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedules *RedisPatchSchedules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedules_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		schedules.Location = &location
	}

	// Set property ‘Owner’:
	schedules.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	for _, item := range typedInput.Properties.ScheduleEntries {
		var item1 ScheduleEntry
		err := item1.PopulateFromARM(owner, item)
		if err != nil {
			return err
		}
		schedules.ScheduleEntries = append(schedules.ScheduleEntries, item1)
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		schedules.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			schedules.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedules_Spec from the provided source
func (schedules *RedisPatchSchedules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisPatchSchedules_Spec)
	if ok {
		// Populate our instance from source
		return schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisPatchSchedules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisPatchSchedules_Spec)
	if ok {
		// Populate destination from our instance
		return schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisPatchSchedules_Spec{}
	err := schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromRedisPatchSchedulesSpec populates our RedisPatchSchedules_Spec from the provided source RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesFromRedisPatchSchedulesSpec(source *v1alpha1api20201201storage.RedisPatchSchedules_Spec) error {

	// Location
	schedules.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	schedules.Owner = source.Owner.Copy()

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignPropertiesFromScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedules.ScheduleEntries = scheduleEntryList
	} else {
		schedules.ScheduleEntries = nil
	}

	// Tags
	schedules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedulesSpec populates the provided destination RedisPatchSchedules_Spec from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesToRedisPatchSchedulesSpec(destination *v1alpha1api20201201storage.RedisPatchSchedules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(schedules.Location)

	// OriginalVersion
	destination.OriginalVersion = schedules.OriginalVersion()

	// Owner
	destination.Owner = schedules.Owner.Copy()

	// ScheduleEntries
	if schedules.ScheduleEntries != nil {
		scheduleEntryList := make([]v1alpha1api20201201storage.ScheduleEntry, len(schedules.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedules.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry v1alpha1api20201201storage.ScheduleEntry
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(schedules.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (schedules *RedisPatchSchedules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry
type ScheduleEntry struct {
	// +kubebuilder:validation:Required
	//DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek ScheduleEntryDayOfWeek `json:"dayOfWeek"`

	//MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	//StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc int `json:"startHourUtc"`
}

var _ genruntime.ARMTransformer = &ScheduleEntry{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (entry *ScheduleEntry) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if entry == nil {
		return nil, nil
	}
	var result ScheduleEntryARM

	// Set property ‘DayOfWeek’:
	result.DayOfWeek = entry.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		result.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	result.StartHourUtc = entry.StartHourUtc
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntryARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntryARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntryARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	entry.DayOfWeek = typedInput.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	entry.StartHourUtc = typedInput.StartHourUtc

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesFromScheduleEntry(source *v1alpha1api20201201storage.ScheduleEntry) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		entry.DayOfWeek = ScheduleEntryDayOfWeek(*source.DayOfWeek)
	} else {
		entry.DayOfWeek = ""
	}

	// MaintenanceWindow
	if source.MaintenanceWindow != nil {
		maintenanceWindow := *source.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	} else {
		entry.MaintenanceWindow = nil
	}

	// StartHourUtc
	entry.StartHourUtc = genruntime.GetOptionalIntValue(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesToScheduleEntry(destination *v1alpha1api20201201storage.ScheduleEntry) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	dayOfWeek := string(entry.DayOfWeek)
	destination.DayOfWeek = &dayOfWeek

	// MaintenanceWindow
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		destination.MaintenanceWindow = &maintenanceWindow
	} else {
		destination.MaintenanceWindow = nil
	}

	// StartHourUtc
	startHourUtc := entry.StartHourUtc
	destination.StartHourUtc = &startHourUtc

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type ScheduleEntry_Status struct {
	// +kubebuilder:validation:Required
	//DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek ScheduleEntryStatusDayOfWeek `json:"dayOfWeek"`

	//MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	//StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc int `json:"startHourUtc"`
}

var _ genruntime.FromARMConverter = &ScheduleEntry_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntry_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntry_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntry_StatusARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	entry.DayOfWeek = typedInput.DayOfWeek

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	entry.StartHourUtc = typedInput.StartHourUtc

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntryStatus populates our ScheduleEntry_Status from the provided source ScheduleEntry_Status
func (entry *ScheduleEntry_Status) AssignPropertiesFromScheduleEntryStatus(source *v1alpha1api20201201storage.ScheduleEntry_Status) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		entry.DayOfWeek = ScheduleEntryStatusDayOfWeek(*source.DayOfWeek)
	} else {
		entry.DayOfWeek = ""
	}

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.GetOptionalIntValue(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntryStatus populates the provided destination ScheduleEntry_Status from our ScheduleEntry_Status
func (entry *ScheduleEntry_Status) AssignPropertiesToScheduleEntryStatus(destination *v1alpha1api20201201storage.ScheduleEntry_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	dayOfWeek := string(entry.DayOfWeek)
	destination.DayOfWeek = &dayOfWeek

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	startHourUtc := entry.StartHourUtc
	destination.StartHourUtc = &startHourUtc

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Everyday","Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday","Weekend"}
type ScheduleEntryDayOfWeek string

const (
	ScheduleEntryDayOfWeekEveryday  = ScheduleEntryDayOfWeek("Everyday")
	ScheduleEntryDayOfWeekFriday    = ScheduleEntryDayOfWeek("Friday")
	ScheduleEntryDayOfWeekMonday    = ScheduleEntryDayOfWeek("Monday")
	ScheduleEntryDayOfWeekSaturday  = ScheduleEntryDayOfWeek("Saturday")
	ScheduleEntryDayOfWeekSunday    = ScheduleEntryDayOfWeek("Sunday")
	ScheduleEntryDayOfWeekThursday  = ScheduleEntryDayOfWeek("Thursday")
	ScheduleEntryDayOfWeekTuesday   = ScheduleEntryDayOfWeek("Tuesday")
	ScheduleEntryDayOfWeekWednesday = ScheduleEntryDayOfWeek("Wednesday")
	ScheduleEntryDayOfWeekWeekend   = ScheduleEntryDayOfWeek("Weekend")
)

type ScheduleEntryStatusDayOfWeek string

const (
	ScheduleEntryStatusDayOfWeekEveryday  = ScheduleEntryStatusDayOfWeek("Everyday")
	ScheduleEntryStatusDayOfWeekFriday    = ScheduleEntryStatusDayOfWeek("Friday")
	ScheduleEntryStatusDayOfWeekMonday    = ScheduleEntryStatusDayOfWeek("Monday")
	ScheduleEntryStatusDayOfWeekSaturday  = ScheduleEntryStatusDayOfWeek("Saturday")
	ScheduleEntryStatusDayOfWeekSunday    = ScheduleEntryStatusDayOfWeek("Sunday")
	ScheduleEntryStatusDayOfWeekThursday  = ScheduleEntryStatusDayOfWeek("Thursday")
	ScheduleEntryStatusDayOfWeekTuesday   = ScheduleEntryStatusDayOfWeek("Tuesday")
	ScheduleEntryStatusDayOfWeekWednesday = ScheduleEntryStatusDayOfWeek("Wednesday")
	ScheduleEntryStatusDayOfWeekWeekend   = ScheduleEntryStatusDayOfWeek("Weekend")
)

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
