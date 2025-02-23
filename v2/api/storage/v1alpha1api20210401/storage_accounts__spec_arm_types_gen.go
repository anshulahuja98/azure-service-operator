// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccounts_SpecARM struct {
	//ExtendedLocation: The complex type of the extended location.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Identity: Identity for the resource.
	Identity *IdentityARM `json:"identity,omitempty"`

	//Kind: Required. Indicates the type of storage account.
	Kind StorageAccountsSpecKind `json:"kind"`

	//Location: Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure
	//Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is
	//created, but if an identical geo region is specified on update, the request will succeed.
	Location string `json:"location,omitempty"`

	//Name: The name of the storage account within the specified resource group. Storage account names must be between 3 and
	//24 characters in length and use numbers and lower-case letters only.
	Name string `json:"name"`

	//Properties: The parameters used to create the storage account.
	Properties *StorageAccountPropertiesCreateParametersARM `json:"properties,omitempty"`

	//Sku: The SKU of the storage account.
	Sku SkuARM `json:"sku"`

	//Tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping
	//this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key
	//with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (accounts StorageAccounts_SpecARM) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (accounts StorageAccounts_SpecARM) GetName() string {
	return accounts.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts"
func (accounts StorageAccounts_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts"
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ExtendedLocation
type ExtendedLocationARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Identity
type IdentityARM struct {
	//Type: The identity type.
	Type IdentityType `json:"type"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Sku
type SkuARM struct {
	Name SkuName  `json:"name"`
	Tier *SkuTier `json:"tier,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/StorageAccountPropertiesCreateParameters
type StorageAccountPropertiesCreateParametersARM struct {
	//AccessTier: Required for storage accounts where kind = BlobStorage. The access tier used for billing.
	AccessTier *StorageAccountPropertiesCreateParametersAccessTier `json:"accessTier,omitempty"`

	//AllowBlobPublicAccess: Allow or disallow public access to all blobs or containers in the storage account. The default
	//interpretation is true for this property.
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty"`

	//AllowCrossTenantReplication: Allow or disallow cross AAD tenant object replication. The default interpretation is true
	//for this property.
	AllowCrossTenantReplication *bool `json:"allowCrossTenantReplication,omitempty"`

	//AllowSharedKeyAccess: Indicates whether the storage account permits requests to be authorized with the account access
	//key via Shared Key. If false, then all requests, including shared access signatures, must be authorized with Azure
	//Active Directory (Azure AD). The default value is null, which is equivalent to true.
	AllowSharedKeyAccess *bool `json:"allowSharedKeyAccess,omitempty"`

	//AzureFilesIdentityBasedAuthentication: Settings for Azure Files identity based authentication.
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthenticationARM `json:"azureFilesIdentityBasedAuthentication,omitempty"`

	//CustomDomain: The custom domain assigned to this storage account. This can be set via Update.
	CustomDomain *CustomDomainARM `json:"customDomain,omitempty"`

	//Encryption: The encryption settings on the storage account.
	Encryption *EncryptionARM `json:"encryption,omitempty"`

	//IsHnsEnabled: Account HierarchicalNamespace enabled if sets to true.
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty"`

	//IsNfsV3Enabled: NFS 3.0 protocol support enabled if set to true.
	IsNfsV3Enabled *bool `json:"isNfsV3Enabled,omitempty"`

	//KeyPolicy: KeyPolicy assigned to the storage account.
	KeyPolicy *KeyPolicyARM `json:"keyPolicy,omitempty"`

	//LargeFileSharesState: Allow large file shares if sets to Enabled. It cannot be disabled once it is enabled.
	LargeFileSharesState *StorageAccountPropertiesCreateParametersLargeFileSharesState `json:"largeFileSharesState,omitempty"`

	//MinimumTlsVersion: Set the minimum TLS version to be permitted on requests to storage. The default interpretation is TLS
	//1.0 for this property.
	MinimumTlsVersion *StorageAccountPropertiesCreateParametersMinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	//NetworkAcls: Network rule set
	NetworkAcls *NetworkRuleSetARM `json:"networkAcls,omitempty"`

	//RoutingPreference: Routing preference defines the type of network, either microsoft or internet routing to be used to
	//deliver the user data, the default option is microsoft routing
	RoutingPreference *RoutingPreferenceARM `json:"routingPreference,omitempty"`

	//SasPolicy: SasPolicy assigned to the storage account.
	SasPolicy *SasPolicyARM `json:"sasPolicy,omitempty"`

	//SupportsHttpsTrafficOnly: Allows https traffic only to storage service if sets to true. The default value is true since
	//API version 2019-04-01.
	SupportsHttpsTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// +kubebuilder:validation:Enum={"BlobStorage","BlockBlobStorage","FileStorage","Storage","StorageV2"}
type StorageAccountsSpecKind string

const (
	StorageAccountsSpecKindBlobStorage      = StorageAccountsSpecKind("BlobStorage")
	StorageAccountsSpecKindBlockBlobStorage = StorageAccountsSpecKind("BlockBlobStorage")
	StorageAccountsSpecKindFileStorage      = StorageAccountsSpecKind("FileStorage")
	StorageAccountsSpecKindStorage          = StorageAccountsSpecKind("Storage")
	StorageAccountsSpecKindStorageV2        = StorageAccountsSpecKind("StorageV2")
)

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/AzureFilesIdentityBasedAuthentication
type AzureFilesIdentityBasedAuthenticationARM struct {
	//ActiveDirectoryProperties: Settings properties for Active Directory (AD).
	ActiveDirectoryProperties *ActiveDirectoryPropertiesARM `json:"activeDirectoryProperties,omitempty"`

	//DefaultSharePermission: Default share permission for users using Kerberos authentication if RBAC role is not assigned.
	DefaultSharePermission *AzureFilesIdentityBasedAuthenticationDefaultSharePermission `json:"defaultSharePermission,omitempty"`

	//DirectoryServiceOptions: Indicates the directory service used.
	DirectoryServiceOptions AzureFilesIdentityBasedAuthenticationDirectoryServiceOptions `json:"directoryServiceOptions"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CustomDomain
type CustomDomainARM struct {
	//Name: Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name string `json:"name"`

	//UseSubDomainName: Indicates whether indirect CName validation is enabled. Default value is false. This should only be
	//set on updates.
	UseSubDomainName *bool `json:"useSubDomainName,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/Encryption
type EncryptionARM struct {
	//Identity: Encryption identity for the storage account.
	Identity *EncryptionIdentityARM `json:"identity,omitempty"`

	//KeySource: The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage,
	//Microsoft.Keyvault.
	KeySource EncryptionKeySource `json:"keySource"`

	//Keyvaultproperties: Properties of key vault.
	Keyvaultproperties *KeyVaultPropertiesARM `json:"keyvaultproperties,omitempty"`

	//RequireInfrastructureEncryption: A boolean indicating whether or not the service applies a secondary layer of encryption
	//with platform managed keys for data at rest.
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`

	//Services: A list of services that support encryption.
	Services *EncryptionServicesARM `json:"services,omitempty"`
}

// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationTypeEdgeZone = ExtendedLocationType("EdgeZone")

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type IdentityType string

const (
	IdentityTypeNone                       = IdentityType("None")
	IdentityTypeSystemAssigned             = IdentityType("SystemAssigned")
	IdentityTypeSystemAssignedUserAssigned = IdentityType("SystemAssigned,UserAssigned")
	IdentityTypeUserAssigned               = IdentityType("UserAssigned")
)

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/KeyPolicy
type KeyPolicyARM struct {
	//KeyExpirationPeriodInDays: The key expiration period in days.
	KeyExpirationPeriodInDays int `json:"keyExpirationPeriodInDays"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/NetworkRuleSet
type NetworkRuleSetARM struct {
	//Bypass: Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of
	//Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	Bypass *NetworkRuleSetBypass `json:"bypass,omitempty"`

	//DefaultAction: Specifies the default action of allow or deny when no other rules match.
	DefaultAction NetworkRuleSetDefaultAction `json:"defaultAction"`

	//IpRules: Sets the IP ACL rules
	IpRules []IPRuleARM `json:"ipRules,omitempty"`

	//ResourceAccessRules: Sets the resource access rules
	ResourceAccessRules []ResourceAccessRuleARM `json:"resourceAccessRules,omitempty"`

	//VirtualNetworkRules: Sets the virtual network rules
	VirtualNetworkRules []VirtualNetworkRuleARM `json:"virtualNetworkRules,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/RoutingPreference
type RoutingPreferenceARM struct {
	//PublishInternetEndpoints: A boolean flag which indicates whether internet routing storage endpoints are to be published
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty"`

	//PublishMicrosoftEndpoints: A boolean flag which indicates whether microsoft routing storage endpoints are to be published
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty"`

	//RoutingChoice: Routing Choice defines the kind of network routing opted by the user.
	RoutingChoice *RoutingPreferenceRoutingChoice `json:"routingChoice,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/SasPolicy
type SasPolicyARM struct {
	//ExpirationAction: The SAS expiration action. Can only be Log.
	ExpirationAction SasPolicyExpirationAction `json:"expirationAction"`

	//SasExpirationPeriod: The SAS expiration period, DD.HH:MM:SS.
	SasExpirationPeriod string `json:"sasExpirationPeriod"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Premium_ZRS","Standard_GRS","Standard_GZRS","Standard_LRS","Standard_RAGRS","Standard_RAGZRS","Standard_ZRS"}
type SkuName string

const (
	SkuNamePremiumLRS     = SkuName("Premium_LRS")
	SkuNamePremiumZRS     = SkuName("Premium_ZRS")
	SkuNameStandardGRS    = SkuName("Standard_GRS")
	SkuNameStandardGZRS   = SkuName("Standard_GZRS")
	SkuNameStandardLRS    = SkuName("Standard_LRS")
	SkuNameStandardRAGRS  = SkuName("Standard_RAGRS")
	SkuNameStandardRAGZRS = SkuName("Standard_RAGZRS")
	SkuNameStandardZRS    = SkuName("Standard_ZRS")
)

// +kubebuilder:validation:Enum={"Premium","Standard"}
type SkuTier string

const (
	SkuTierPremium  = SkuTier("Premium")
	SkuTierStandard = SkuTier("Standard")
)

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ActiveDirectoryProperties
type ActiveDirectoryPropertiesARM struct {
	//AzureStorageSid: Specifies the security identifier (SID) for Azure Storage.
	AzureStorageSid string `json:"azureStorageSid"`

	//DomainGuid: Specifies the domain GUID.
	DomainGuid string `json:"domainGuid"`

	//DomainName: Specifies the primary domain that the AD DNS server is authoritative for.
	DomainName string `json:"domainName"`

	//DomainSid: Specifies the security identifier (SID).
	DomainSid string `json:"domainSid"`

	//ForestName: Specifies the Active Directory forest to get.
	ForestName string `json:"forestName"`

	//NetBiosDomainName: Specifies the NetBIOS domain name.
	NetBiosDomainName string `json:"netBiosDomainName"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionIdentity
type EncryptionIdentityARM struct {
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionServices
type EncryptionServicesARM struct {
	//Blob: A service that allows server-side encryption to be used.
	Blob *EncryptionServiceARM `json:"blob,omitempty"`

	//File: A service that allows server-side encryption to be used.
	File *EncryptionServiceARM `json:"file,omitempty"`

	//Queue: A service that allows server-side encryption to be used.
	Queue *EncryptionServiceARM `json:"queue,omitempty"`

	//Table: A service that allows server-side encryption to be used.
	Table *EncryptionServiceARM `json:"table,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/IPRule
type IPRuleARM struct {
	//Action: The action of IP ACL rule.
	Action *IPRuleAction `json:"action,omitempty"`

	//Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	Value string `json:"value"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/KeyVaultProperties
type KeyVaultPropertiesARM struct {
	//Keyname: The name of KeyVault key.
	Keyname *string `json:"keyname,omitempty"`

	//Keyvaulturi: The Uri of KeyVault.
	Keyvaulturi *string `json:"keyvaulturi,omitempty"`

	//Keyversion: The version of KeyVault key.
	Keyversion *string `json:"keyversion,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ResourceAccessRule
type ResourceAccessRuleARM struct {
	ResourceId *string `json:"resourceId,omitempty"`

	//TenantId: Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/VirtualNetworkRule
type VirtualNetworkRuleARM struct {
	//Action: The action of virtual network rule.
	Action *VirtualNetworkRuleAction `json:"action,omitempty"`
	Id     string                    `json:"id"`

	//State: Gets the state of virtual network rule.
	State *VirtualNetworkRuleState `json:"state,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/EncryptionService
type EncryptionServiceARM struct {
	//Enabled: A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled *bool `json:"enabled,omitempty"`

	//KeyType: Encryption key type to be used for the encryption service. 'Account' key type implies that an account-scoped
	//encryption key will be used. 'Service' key type implies that a default service key is used.
	KeyType *EncryptionServiceKeyType `json:"keyType,omitempty"`
}
