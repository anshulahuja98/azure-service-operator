// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/storage/v1alpha1api20210401storage"
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
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_queueServices_queues
type StorageAccountsQueueServicesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsQueueServicesQueues_Spec `json:"spec,omitempty"`
	Status            StorageQueue_Status                     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsQueueServicesQueue{}

// GetConditions returns the conditions of the resource
func (queue *StorageAccountsQueueServicesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *StorageAccountsQueueServicesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsQueueServicesQueue{}

// ConvertFrom populates our StorageAccountsQueueServicesQueue from the provided hub StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20210401storage.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage:storage/v1alpha1api20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesFromStorageAccountsQueueServicesQueue(source)
}

// ConvertTo populates the provided hub StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210401storage.StorageAccountsQueueServicesQueue)
	if !ok {
		return fmt.Errorf("expected storage:storage/v1alpha1api20210401storage/StorageAccountsQueueServicesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesToStorageAccountsQueueServicesQueue(destination)
}

// +kubebuilder:webhook:path=/mutate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservicesqueue,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservicesqueues,verbs=create;update,versions=v1alpha1api20210401,name=default.v1alpha1api20210401.storageaccountsqueueservicesqueues.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &StorageAccountsQueueServicesQueue{}

// Default applies defaults to the StorageAccountsQueueServicesQueue resource
func (queue *StorageAccountsQueueServicesQueue) Default() {
	queue.defaultImpl()
	var temp interface{} = queue
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (queue *StorageAccountsQueueServicesQueue) defaultAzureName() {
	if queue.Spec.AzureName == "" {
		queue.Spec.AzureName = queue.Name
	}
}

// defaultImpl applies the code generated defaults to the StorageAccountsQueueServicesQueue resource
func (queue *StorageAccountsQueueServicesQueue) defaultImpl() { queue.defaultAzureName() }

var _ genruntime.KubernetesResource = &StorageAccountsQueueServicesQueue{}

// AzureName returns the Azure name of the resource
func (queue *StorageAccountsQueueServicesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (queue StorageAccountsQueueServicesQueue) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceKind returns the kind of the resource
func (queue *StorageAccountsQueueServicesQueue) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (queue *StorageAccountsQueueServicesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *StorageAccountsQueueServicesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices/queues"
func (queue *StorageAccountsQueueServicesQueue) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *StorageAccountsQueueServicesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageQueue_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (queue *StorageAccountsQueueServicesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (queue *StorageAccountsQueueServicesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageQueue_Status); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageQueue_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-storage-azure-com-v1alpha1api20210401-storageaccountsqueueservicesqueue,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=storage.azure.com,resources=storageaccountsqueueservicesqueues,verbs=create;update,versions=v1alpha1api20210401,name=validate.v1alpha1api20210401.storageaccountsqueueservicesqueues.storage.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &StorageAccountsQueueServicesQueue{}

// ValidateCreate validates the creation of the resource
func (queue *StorageAccountsQueueServicesQueue) ValidateCreate() error {
	validations := queue.createValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) ValidateDelete() error {
	validations := queue.deleteValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) ValidateUpdate(old runtime.Object) error {
	validations := queue.updateValidations()
	var temp interface{} = queue
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
func (queue *StorageAccountsQueueServicesQueue) createValidations() []func() error {
	return []func() error{queue.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (queue *StorageAccountsQueueServicesQueue) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (queue *StorageAccountsQueueServicesQueue) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return queue.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (queue *StorageAccountsQueueServicesQueue) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&queue.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromStorageAccountsQueueServicesQueue populates our StorageAccountsQueueServicesQueue from the provided source StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesFromStorageAccountsQueueServicesQueue(source *v1alpha1api20210401storage.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccountsQueueServicesQueues_Spec
	err := spec.AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status StorageQueue_Status
	err = status.AssignPropertiesFromStorageQueueStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromStorageQueueStatus() to populate field Status")
	}
	queue.Status = status

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueue populates the provided destination StorageAccountsQueueServicesQueue from our StorageAccountsQueueServicesQueue
func (queue *StorageAccountsQueueServicesQueue) AssignPropertiesToStorageAccountsQueueServicesQueue(destination *v1alpha1api20210401storage.StorageAccountsQueueServicesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec
	err := queue.Spec.AssignPropertiesToStorageAccountsQueueServicesQueuesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageAccountsQueueServicesQueuesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210401storage.StorageQueue_Status
	err = queue.Status.AssignPropertiesToStorageQueueStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToStorageQueueStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *StorageAccountsQueueServicesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion(),
		Kind:    "StorageAccountsQueueServicesQueue",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_queueServices_queues
type StorageAccountsQueueServicesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsQueueServicesQueue `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-04-01"}
type StorageAccountsQueueServicesQueuesSpecAPIVersion string

const StorageAccountsQueueServicesQueuesSpecAPIVersion20210401 = StorageAccountsQueueServicesQueuesSpecAPIVersion("2021-04-01")

type StorageAccountsQueueServicesQueues_Spec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:MinLength=3
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner" kind:"StorageAccountsQueueService"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &StorageAccountsQueueServicesQueues_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (queues *StorageAccountsQueueServicesQueues_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if queues == nil {
		return nil, nil
	}
	var result StorageAccountsQueueServicesQueues_SpecARM

	// Set property ‘Location’:
	if queues.Location != nil {
		location := *queues.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if queues.Metadata != nil {
		result.Properties = &QueuePropertiesARM{}
	}
	if queues.Metadata != nil {
		result.Properties.Metadata = make(map[string]string)
		for key, value := range queues.Metadata {
			result.Properties.Metadata[key] = value
		}
	}

	// Set property ‘Tags’:
	if queues.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range queues.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (queues *StorageAccountsQueueServicesQueues_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageAccountsQueueServicesQueues_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (queues *StorageAccountsQueueServicesQueues_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageAccountsQueueServicesQueues_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageAccountsQueueServicesQueues_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	queues.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		queues.Location = &location
	}

	// Set property ‘Metadata’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Metadata != nil {
			queues.Metadata = make(map[string]string)
			for key, value := range typedInput.Properties.Metadata {
				queues.Metadata[key] = value
			}
		}
	}

	// Set property ‘Owner’:
	queues.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		queues.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			queues.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &StorageAccountsQueueServicesQueues_Spec{}

// ConvertSpecFrom populates our StorageAccountsQueueServicesQueues_Spec from the provided source
func (queues *StorageAccountsQueueServicesQueues_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec)
	if ok {
		// Populate our instance from source
		return queues.AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queues.AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccountsQueueServicesQueues_Spec
func (queues *StorageAccountsQueueServicesQueues_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec)
	if ok {
		// Populate destination from our instance
		return queues.AssignPropertiesToStorageAccountsQueueServicesQueuesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec{}
	err := queues.AssignPropertiesToStorageAccountsQueueServicesQueuesSpec(dst)
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

// AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec populates our StorageAccountsQueueServicesQueues_Spec from the provided source StorageAccountsQueueServicesQueues_Spec
func (queues *StorageAccountsQueueServicesQueues_Spec) AssignPropertiesFromStorageAccountsQueueServicesQueuesSpec(source *v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec) error {

	// AzureName
	queues.AzureName = source.AzureName

	// Location
	queues.Location = genruntime.ClonePointerToString(source.Location)

	// Metadata
	queues.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Owner
	queues.Owner = source.Owner.Copy()

	// Tags
	queues.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToStorageAccountsQueueServicesQueuesSpec populates the provided destination StorageAccountsQueueServicesQueues_Spec from our StorageAccountsQueueServicesQueues_Spec
func (queues *StorageAccountsQueueServicesQueues_Spec) AssignPropertiesToStorageAccountsQueueServicesQueuesSpec(destination *v1alpha1api20210401storage.StorageAccountsQueueServicesQueues_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = queues.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(queues.Location)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queues.Metadata)

	// OriginalVersion
	destination.OriginalVersion = queues.OriginalVersion()

	// Owner
	destination.Owner = queues.Owner.Copy()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(queues.Tags)

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
func (queues *StorageAccountsQueueServicesQueues_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (queues *StorageAccountsQueueServicesQueues_Spec) SetAzureName(azureName string) {
	queues.AzureName = azureName
}

type StorageQueue_Status struct {
	//ApproximateMessageCount: Integer indicating an approximate number of messages in the queue. This number is not lower
	//than the actual number of messages in the queue, but could be higher.
	ApproximateMessageCount *int `json:"approximateMessageCount,omitempty"`

	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Metadata: A name-value pair that represents queue metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageQueue_Status{}

// ConvertStatusFrom populates our StorageQueue_Status from the provided source
func (queue *StorageQueue_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210401storage.StorageQueue_Status)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromStorageQueueStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210401storage.StorageQueue_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromStorageQueueStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageQueue_Status
func (queue *StorageQueue_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210401storage.StorageQueue_Status)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToStorageQueueStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210401storage.StorageQueue_Status{}
	err := queue.AssignPropertiesToStorageQueueStatus(dst)
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

var _ genruntime.FromARMConverter = &StorageQueue_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (queue *StorageQueue_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &StorageQueue_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (queue *StorageQueue_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(StorageQueue_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected StorageQueue_StatusARM, got %T", armInput)
	}

	// Set property ‘ApproximateMessageCount’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ApproximateMessageCount != nil {
			approximateMessageCount := *typedInput.Properties.ApproximateMessageCount
			queue.ApproximateMessageCount = &approximateMessageCount
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		queue.Id = &id
	}

	// Set property ‘Metadata’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Metadata != nil {
			queue.Metadata = make(map[string]string)
			for key, value := range typedInput.Properties.Metadata {
				queue.Metadata[key] = value
			}
		}
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		queue.Name = &name
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		queue.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromStorageQueueStatus populates our StorageQueue_Status from the provided source StorageQueue_Status
func (queue *StorageQueue_Status) AssignPropertiesFromStorageQueueStatus(source *v1alpha1api20210401storage.StorageQueue_Status) error {

	// ApproximateMessageCount
	queue.ApproximateMessageCount = genruntime.ClonePointerToInt(source.ApproximateMessageCount)

	// Conditions
	queue.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// Metadata
	queue.Metadata = genruntime.CloneMapOfStringToString(source.Metadata)

	// Name
	queue.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToStorageQueueStatus populates the provided destination StorageQueue_Status from our StorageQueue_Status
func (queue *StorageQueue_Status) AssignPropertiesToStorageQueueStatus(destination *v1alpha1api20210401storage.StorageQueue_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ApproximateMessageCount
	destination.ApproximateMessageCount = genruntime.ClonePointerToInt(queue.ApproximateMessageCount)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queue.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// Metadata
	destination.Metadata = genruntime.CloneMapOfStringToString(queue.Metadata)

	// Name
	destination.Name = genruntime.ClonePointerToString(queue.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&StorageAccountsQueueServicesQueue{}, &StorageAccountsQueueServicesQueueList{})
}
