// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200930

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Disks_SpecARM struct {
	//ExtendedLocation: The complex type of the extended location.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//Name: The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
	//characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name string `json:"name"`

	//Properties: Disk resource properties.
	Properties DiskPropertiesARM `json:"properties"`

	//Sku: The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
	Sku *DiskSkuARM `json:"sku,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Zones: The Logical zone list for Disk.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Disks_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-09-30"
func (disks Disks_SpecARM) GetAPIVersion() string {
	return "2020-09-30"
}

// GetName returns the Name of the resource
func (disks Disks_SpecARM) GetName() string {
	return disks.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/disks"
func (disks Disks_SpecARM) GetType() string {
	return "Microsoft.Compute/disks"
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskProperties
type DiskPropertiesARM struct {
	//BurstingEnabled: Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is
	//disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `json:"burstingEnabled,omitempty"`

	//CreationData: Data used when creating a disk.
	CreationData CreationDataARM `json:"creationData"`
	DiskAccessId *string         `json:"diskAccessId,omitempty"`

	//DiskIOPSReadOnly: The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One
	//operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *int `json:"diskIOPSReadOnly,omitempty"`

	//DiskIOPSReadWrite: The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can
	//transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *int `json:"diskIOPSReadWrite,omitempty"`

	//DiskMBpsReadOnly: The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly.
	//MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *int `json:"diskMBpsReadOnly,omitempty"`

	//DiskMBpsReadWrite: The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes
	//per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *int `json:"diskMBpsReadWrite,omitempty"`

	//DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	//create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	//allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	//Encryption: Encryption at rest settings for disk or snapshot
	Encryption *EncryptionARM `json:"encryption,omitempty"`

	//EncryptionSettingsCollection: Encryption settings for disk or snapshot
	EncryptionSettingsCollection *EncryptionSettingsCollectionARM `json:"encryptionSettingsCollection,omitempty"`

	//HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *DiskPropertiesHyperVGeneration `json:"hyperVGeneration,omitempty"`

	//MaxShares: The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a
	//disk that can be mounted on multiple VMs at the same time.
	MaxShares           *int                               `json:"maxShares,omitempty"`
	NetworkAccessPolicy *DiskPropertiesNetworkAccessPolicy `json:"networkAccessPolicy,omitempty"`

	//OsType: The Operating System type.
	OsType *DiskPropertiesOsType `json:"osType,omitempty"`

	//PurchasePlan: Used for establishing the purchase context of any 3rd Party artifact through MarketPlace.
	PurchasePlan *PurchasePlanARM `json:"purchasePlan,omitempty"`

	//Tier: Performance tier of the disk (e.g, P4, S10) as described here:
	//https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `json:"tier,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/DiskSku
type DiskSkuARM struct {
	//Name: The sku name.
	Name *DiskSkuName `json:"name,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ExtendedLocation
type ExtendedLocationARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/CreationData
type CreationDataARM struct {
	//CreateOption: This enumerates the possible sources of a disk's creation.
	CreateOption CreationDataCreateOption `json:"createOption"`

	//GalleryImageReference: The source image used for creating the disk.
	GalleryImageReference *ImageDiskReferenceARM `json:"galleryImageReference,omitempty"`

	//ImageReference: The source image used for creating the disk.
	ImageReference *ImageDiskReferenceARM `json:"imageReference,omitempty"`

	//LogicalSectorSize: Logical sector size in bytes for Ultra disks. Supported values are 512 ad 4096. 4096 is the default.
	LogicalSectorSize *int    `json:"logicalSectorSize,omitempty"`
	SourceResourceId  *string `json:"sourceResourceId,omitempty"`

	//SourceUri: If createOption is Import, this is the URI of a blob to be imported into a managed disk.
	SourceUri *string `json:"sourceUri,omitempty"`

	//StorageAccountId: Required if createOption is Import. The Azure Resource Manager identifier of the storage account
	//containing the blob to import as a disk.
	StorageAccountId *string `json:"storageAccountId,omitempty"`

	//UploadSizeBytes: If createOption is Upload, this is the size of the contents of the upload including the VHD footer.
	//This value should be between 20972032 (20 MiB + 512 bytes for the VHD footer) and 35183298347520 bytes (32 TiB + 512
	//bytes for the VHD footer).
	UploadSizeBytes *int `json:"uploadSizeBytes,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Standard_LRS","StandardSSD_LRS","UltraSSD_LRS"}
type DiskSkuName string

const (
	DiskSkuNamePremiumLRS     = DiskSkuName("Premium_LRS")
	DiskSkuNameStandardLRS    = DiskSkuName("Standard_LRS")
	DiskSkuNameStandardSSDLRS = DiskSkuName("StandardSSD_LRS")
	DiskSkuNameUltraSSDLRS    = DiskSkuName("UltraSSD_LRS")
)

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/Encryption
type EncryptionARM struct {
	DiskEncryptionSetId *string         `json:"diskEncryptionSetId,omitempty"`
	Type                *EncryptionType `json:"type,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsCollection
type EncryptionSettingsCollectionARM struct {
	//Enabled: Set this flag to true and provide DiskEncryptionKey and optional KeyEncryptionKey to enable encryption. Set
	//this flag to false and remove DiskEncryptionKey and KeyEncryptionKey to disable encryption. If EncryptionSettings is
	//null in the request object, the existing settings remain unchanged.
	Enabled bool `json:"enabled"`

	//EncryptionSettings: A collection of encryption settings, one for each disk volume.
	EncryptionSettings []EncryptionSettingsElementARM `json:"encryptionSettings,omitempty"`

	//EncryptionSettingsVersion: Describes what type of encryption is used for the disks. Once this field is set, it cannot be
	//overwritten. '1.0' corresponds to Azure Disk Encryption with AAD app.'1.1' corresponds to Azure Disk Encryption.
	EncryptionSettingsVersion *string `json:"encryptionSettingsVersion,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/PurchasePlan
type PurchasePlanARM struct {
	//Name: The plan ID.
	Name string `json:"name"`

	//Product: Specifies the product of the image from the marketplace. This is the same value as Offer under the
	//imageReference element.
	Product string `json:"product"`

	//PromotionCode: The Offer Promotion Code.
	PromotionCode *string `json:"promotionCode,omitempty"`

	//Publisher: The publisher ID.
	Publisher string `json:"publisher"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/EncryptionSettingsElement
type EncryptionSettingsElementARM struct {
	//DiskEncryptionKey: Key Vault Secret Url and vault id of the encryption key
	DiskEncryptionKey *KeyVaultAndSecretReferenceARM `json:"diskEncryptionKey,omitempty"`

	//KeyEncryptionKey: Key Vault Key Url and vault id of KeK, KeK is optional and when provided is used to unwrap the
	//encryptionKey
	KeyEncryptionKey *KeyVaultAndKeyReferenceARM `json:"keyEncryptionKey,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/ImageDiskReference
type ImageDiskReferenceARM struct {
	Id string `json:"id"`

	//Lun: If the disk is created from an image's data disk, this is an index that indicates which of the data disks in the
	//image to use. For OS disks, this field is null.
	Lun *int `json:"lun,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndKeyReference
type KeyVaultAndKeyReferenceARM struct {
	//KeyUrl: Url pointing to a key or secret in KeyVault
	KeyUrl string `json:"keyUrl"`

	//SourceVault: The vault id is an Azure Resource Manager Resource id in the form
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
	SourceVault SourceVaultARM `json:"sourceVault"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/KeyVaultAndSecretReference
type KeyVaultAndSecretReferenceARM struct {
	//SecretUrl: Url pointing to a key or secret in KeyVault
	SecretUrl string `json:"secretUrl"`

	//SourceVault: The vault id is an Azure Resource Manager Resource id in the form
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
	SourceVault SourceVaultARM `json:"sourceVault"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-09-30/Microsoft.Compute.json#/definitions/SourceVault
type SourceVaultARM struct {
	Id *string `json:"id,omitempty"`
}
