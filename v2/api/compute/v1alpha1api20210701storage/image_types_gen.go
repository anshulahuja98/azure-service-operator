// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210701storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=compute.azure.com,resources=images,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compute.azure.com,resources={images/status,images/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210701.Image
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/resourceDefinitions/images
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Images_Spec  `json:"spec,omitempty"`
	Status            Image_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Image{}

// GetConditions returns the conditions of the resource
func (image *Image) GetConditions() conditions.Conditions {
	return image.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (image *Image) SetConditions(conditions conditions.Conditions) {
	image.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Image{}

// AzureName returns the Azure name of the resource
func (image *Image) AzureName() string {
	return image.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (image Image) GetAPIVersion() string {
	return "2021-07-01"
}

// GetResourceKind returns the kind of the resource
func (image *Image) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (image *Image) GetSpec() genruntime.ConvertibleSpec {
	return &image.Spec
}

// GetStatus returns the status of this resource
func (image *Image) GetStatus() genruntime.ConvertibleStatus {
	return &image.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/images"
func (image *Image) GetType() string {
	return "Microsoft.Compute/images"
}

// NewEmptyStatus returns a new empty (blank) status
func (image *Image) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Image_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (image *Image) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(image.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  image.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (image *Image) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Image_Status); ok {
		image.Status = *st
		return nil
	}

	// Convert status to required version
	var st Image_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	image.Status = st
	return nil
}

// Hub marks that this Image is the hub type for conversion
func (image *Image) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (image *Image) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: image.Spec.OriginalVersion,
		Kind:    "Image",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210701.Image
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/resourceDefinitions/images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

//Storage version of v1alpha1api20210701.Image_Status
type Image_Status struct {
	Conditions           []conditions.Condition      `json:"conditions,omitempty"`
	ExtendedLocation     *ExtendedLocation_Status    `json:"extendedLocation,omitempty"`
	HyperVGeneration     *string                     `json:"hyperVGeneration,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Location             *string                     `json:"location,omitempty"`
	Name                 *string                     `json:"name,omitempty"`
	PropertyBag          genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	ProvisioningState    *string                     `json:"provisioningState,omitempty"`
	SourceVirtualMachine *SubResource_Status         `json:"sourceVirtualMachine,omitempty"`
	StorageProfile       *ImageStorageProfile_Status `json:"storageProfile,omitempty"`
	Tags                 map[string]string           `json:"tags,omitempty"`
	Type                 *string                     `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Image_Status{}

// ConvertStatusFrom populates our Image_Status from the provided source
func (image *Image_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(image)
}

// ConvertStatusTo populates the provided destination from our Image_Status
func (image *Image_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == image {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(image)
}

//Storage version of v1alpha1api20210701.Images_Spec
type Images_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName        string            `json:"azureName"`
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`
	HyperVGeneration *string           `json:"hyperVGeneration,omitempty"`
	Location         *string           `json:"location,omitempty"`
	OriginalVersion  string            `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag          genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	SourceVirtualMachine *SubResource                      `json:"sourceVirtualMachine,omitempty"`
	StorageProfile       *ImageStorageProfile              `json:"storageProfile,omitempty"`
	Tags                 map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Images_Spec{}

// ConvertSpecFrom populates our Images_Spec from the provided source
func (images *Images_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == images {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(images)
}

// ConvertSpecTo populates the provided destination from our Images_Spec
func (images *Images_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == images {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(images)
}

//Storage version of v1alpha1api20210701.ExtendedLocation
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ExtendedLocation
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210701.ExtendedLocation_Status
type ExtendedLocation_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageStorageProfile
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageStorageProfile
type ImageStorageProfile struct {
	DataDisks     []ImageDataDisk        `json:"dataDisks,omitempty"`
	OsDisk        *ImageOSDisk           `json:"osDisk,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ZoneResilient *bool                  `json:"zoneResilient,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageStorageProfile_Status
type ImageStorageProfile_Status struct {
	DataDisks     []ImageDataDisk_Status `json:"dataDisks,omitempty"`
	OsDisk        *ImageOSDisk_Status    `json:"osDisk,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ZoneResilient *bool                  `json:"zoneResilient,omitempty"`
}

//Storage version of v1alpha1api20210701.SubResource
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/SubResource
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

//Storage version of v1alpha1api20210701.SubResource_Status
type SubResource_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageDataDisk
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageDataDisk
type ImageDataDisk struct {
	BlobUri            *string                      `json:"blobUri,omitempty"`
	Caching            *string                      `json:"caching,omitempty"`
	DiskEncryptionSet  *DiskEncryptionSetParameters `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                         `json:"diskSizeGB,omitempty"`
	Lun                *int                         `json:"lun,omitempty"`
	ManagedDisk        *SubResource                 `json:"managedDisk,omitempty"`
	PropertyBag        genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource                 `json:"snapshot,omitempty"`
	StorageAccountType *string                      `json:"storageAccountType,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageDataDisk_Status
type ImageDataDisk_Status struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_Status    `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	Lun                *int                   `json:"lun,omitempty"`
	ManagedDisk        *SubResource_Status    `json:"managedDisk,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource_Status    `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageOSDisk
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/ImageOSDisk
type ImageOSDisk struct {
	BlobUri            *string                      `json:"blobUri,omitempty"`
	Caching            *string                      `json:"caching,omitempty"`
	DiskEncryptionSet  *DiskEncryptionSetParameters `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                         `json:"diskSizeGB,omitempty"`
	ManagedDisk        *SubResource                 `json:"managedDisk,omitempty"`
	OsState            *string                      `json:"osState,omitempty"`
	OsType             *string                      `json:"osType,omitempty"`
	PropertyBag        genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource                 `json:"snapshot,omitempty"`
	StorageAccountType *string                      `json:"storageAccountType,omitempty"`
}

//Storage version of v1alpha1api20210701.ImageOSDisk_Status
type ImageOSDisk_Status struct {
	BlobUri            *string                `json:"blobUri,omitempty"`
	Caching            *string                `json:"caching,omitempty"`
	DiskEncryptionSet  *SubResource_Status    `json:"diskEncryptionSet,omitempty"`
	DiskSizeGB         *int                   `json:"diskSizeGB,omitempty"`
	ManagedDisk        *SubResource_Status    `json:"managedDisk,omitempty"`
	OsState            *string                `json:"osState,omitempty"`
	OsType             *string                `json:"osType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Snapshot           *SubResource_Status    `json:"snapshot,omitempty"`
	StorageAccountType *string                `json:"storageAccountType,omitempty"`
}

//Storage version of v1alpha1api20210701.DiskEncryptionSetParameters
//Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.Compute.json#/definitions/DiskEncryptionSetParameters
type DiskEncryptionSetParameters struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
