// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServers_SpecARM struct {
	//Location: The geo-location where the resource lives
	Location string `json:"location,omitempty"`

	//Name: The name of the server.
	Name string `json:"name"`

	//Properties: The properties of a server.
	Properties ServerPropertiesARM `json:"properties"`

	//Sku: Sku information related properties of a server.
	Sku *SkuARM `json:"sku,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (servers FlexibleServers_SpecARM) GetAPIVersion() string {
	return "2021-06-01"
}

// GetName returns the Name of the resource
func (servers FlexibleServers_SpecARM) GetName() string {
	return servers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers"
func (servers FlexibleServers_SpecARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers"
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ServerProperties
type ServerPropertiesARM struct {
	//AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	//(and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	//AdministratorLoginPassword: The administrator login password (required for server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	//AvailabilityZone: availability zone information of the server.
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	//Backup: Backup properties of a server
	Backup *BackupARM `json:"backup,omitempty"`

	//CreateMode: The mode to create a new PostgreSQL server.
	CreateMode *ServerPropertiesCreateMode `json:"createMode,omitempty"`

	//HighAvailability: High availability properties of a server
	HighAvailability *HighAvailabilityARM `json:"highAvailability,omitempty"`

	//MaintenanceWindow: Maintenance window properties of a server.
	MaintenanceWindow *MaintenanceWindowARM `json:"maintenanceWindow,omitempty"`

	//Network: Network properties of a server
	Network *NetworkARM `json:"network,omitempty"`

	//PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It's required when
	//'createMode' is 'PointInTimeRestore'.
	PointInTimeUTC         *string `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId *string `json:"sourceServerResourceId,omitempty"`

	//Storage: Storage properties of a server
	Storage *StorageARM `json:"storage,omitempty"`

	//Version: PostgreSQL Server version.
	Version *ServerPropertiesVersion `json:"version,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku
type SkuARM struct {
	//Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.
	Name string `json:"name"`

	//Tier: The tier of the particular SKU, e.g. Burstable.
	Tier SkuTier `json:"tier"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup
type BackupARM struct {
	//BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	//GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.
	GeoRedundantBackup *BackupGeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability
type HighAvailabilityARM struct {
	//Mode: The HA mode for the server.
	Mode *HighAvailabilityMode `json:"mode,omitempty"`

	//StandbyAvailabilityZone: availability zone information of the standby.
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow
type MaintenanceWindowARM struct {
	//CustomWindow: indicates whether custom window is enabled or disabled
	CustomWindow *string `json:"customWindow,omitempty"`

	//DayOfWeek: day of week for maintenance window
	DayOfWeek *int `json:"dayOfWeek,omitempty"`

	//StartHour: start hour for maintenance window
	StartHour *int `json:"startHour,omitempty"`

	//StartMinute: start minute for maintenance window
	StartMinute *int `json:"startMinute,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network
type NetworkARM struct {
	DelegatedSubnetResourceId   *string `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string `json:"privateDnsZoneArmResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"Burstable","GeneralPurpose","MemoryOptimized"}
type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage
type StorageARM struct {
	//StorageSizeGB: Max storage allowed for a server.
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}
