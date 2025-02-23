// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleservers/status,flexibleservers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210601.FlexibleServer
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers
type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Spec `json:"spec,omitempty"`
	Status            Server_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServer{}

// GetConditions returns the conditions of the resource
func (server *FlexibleServer) GetConditions() conditions.Conditions {
	return server.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (server *FlexibleServer) SetConditions(conditions conditions.Conditions) {
	server.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServer{}

// AzureName returns the Azure name of the resource
func (server *FlexibleServer) AzureName() string {
	return server.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (server FlexibleServer) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (server *FlexibleServer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (server *FlexibleServer) GetSpec() genruntime.ConvertibleSpec {
	return &server.Spec
}

// GetStatus returns the status of this resource
func (server *FlexibleServer) GetStatus() genruntime.ConvertibleStatus {
	return &server.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers"
func (server *FlexibleServer) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers"
}

// NewEmptyStatus returns a new empty (blank) status
func (server *FlexibleServer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Server_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (server *FlexibleServer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(server.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  server.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (server *FlexibleServer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Server_Status); ok {
		server.Status = *st
		return nil
	}

	// Convert status to required version
	var st Server_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	server.Status = st
	return nil
}

// Hub marks that this FlexibleServer is the hub type for conversion
func (server *FlexibleServer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (server *FlexibleServer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: server.Spec.OriginalVersion,
		Kind:    "FlexibleServer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210601.FlexibleServer
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServer `json:"items"`
}

//Storage version of v1alpha1api20210601.FlexibleServers_Spec
type FlexibleServers_Spec struct {
	AdministratorLogin         *string                     `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *genruntime.SecretReference `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string                     `json:"availabilityZone,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName         string             `json:"azureName"`
	Backup            *Backup            `json:"backup,omitempty"`
	CreateMode        *string            `json:"createMode,omitempty"`
	HighAvailability  *HighAvailability  `json:"highAvailability,omitempty"`
	Location          *string            `json:"location,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	Network           *Network           `json:"network,omitempty"`
	OriginalVersion   string             `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner          genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PointInTimeUTC *string                           `json:"pointInTimeUTC,omitempty"`
	PropertyBag    genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Sku            *Sku                              `json:"sku,omitempty"`

	//SourceServerResourceReference: The source server resource ID to restore from. It's required when 'createMode' is
	//'PointInTimeRestore'.
	SourceServerResourceReference *genruntime.ResourceReference `armReference:"SourceServerResourceId" json:"sourceServerResourceReference,omitempty"`
	Storage                       *Storage                      `json:"storage,omitempty"`
	Tags                          map[string]string             `json:"tags,omitempty"`
	Version                       *string                       `json:"version,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Spec from the provided source
func (servers *FlexibleServers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == servers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(servers)
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Spec
func (servers *FlexibleServers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == servers {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(servers)
}

//Storage version of v1alpha1api20210601.Server_Status
type Server_Status struct {
	AdministratorLogin       *string                   `json:"administratorLogin,omitempty"`
	AvailabilityZone         *string                   `json:"availabilityZone,omitempty"`
	Backup                   *Backup_Status            `json:"backup,omitempty"`
	Conditions               []conditions.Condition    `json:"conditions,omitempty"`
	CreateMode               *string                   `json:"createMode,omitempty"`
	FullyQualifiedDomainName *string                   `json:"fullyQualifiedDomainName,omitempty"`
	HighAvailability         *HighAvailability_Status  `json:"highAvailability,omitempty"`
	Id                       *string                   `json:"id,omitempty"`
	Location                 *string                   `json:"location,omitempty"`
	MaintenanceWindow        *MaintenanceWindow_Status `json:"maintenanceWindow,omitempty"`
	MinorVersion             *string                   `json:"minorVersion,omitempty"`
	Name                     *string                   `json:"name,omitempty"`
	Network                  *Network_Status           `json:"network,omitempty"`
	PointInTimeUTC           *string                   `json:"pointInTimeUTC,omitempty"`
	PropertyBag              genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Sku                      *Sku_Status               `json:"sku,omitempty"`
	SourceServerResourceId   *string                   `json:"sourceServerResourceId,omitempty"`
	State                    *string                   `json:"state,omitempty"`
	Storage                  *Storage_Status           `json:"storage,omitempty"`
	SystemData               *SystemData_Status        `json:"systemData,omitempty"`
	Tags                     map[string]string         `json:"tags,omitempty"`
	Type                     *string                   `json:"type,omitempty"`
	Version                  *string                   `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Server_Status{}

// ConvertStatusFrom populates our Server_Status from the provided source
func (server *Server_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(server)
}

// ConvertStatusTo populates the provided destination from our Server_Status
func (server *Server_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == server {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(server)
}

//Storage version of v1alpha1api20210601.Backup
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup
type Backup struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.Backup_Status
type Backup_Status struct {
	BackupRetentionDays *int                   `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string                `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *string                `json:"geoRedundantBackup,omitempty"`
	PropertyBag         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.HighAvailability
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability
type HighAvailability struct {
	Mode                    *string                `json:"mode,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
}

//Storage version of v1alpha1api20210601.HighAvailability_Status
type HighAvailability_Status struct {
	Mode                    *string                `json:"mode,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StandbyAvailabilityZone *string                `json:"standbyAvailabilityZone,omitempty"`
	State                   *string                `json:"state,omitempty"`
}

//Storage version of v1alpha1api20210601.MaintenanceWindow
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow
type MaintenanceWindow struct {
	CustomWindow *string                `json:"customWindow,omitempty"`
	DayOfWeek    *int                   `json:"dayOfWeek,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHour    *int                   `json:"startHour,omitempty"`
	StartMinute  *int                   `json:"startMinute,omitempty"`
}

//Storage version of v1alpha1api20210601.MaintenanceWindow_Status
type MaintenanceWindow_Status struct {
	CustomWindow *string                `json:"customWindow,omitempty"`
	DayOfWeek    *int                   `json:"dayOfWeek,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHour    *int                   `json:"startHour,omitempty"`
	StartMinute  *int                   `json:"startMinute,omitempty"`
}

//Storage version of v1alpha1api20210601.Network
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network
type Network struct {
	//DelegatedSubnetResourceReference: delegated subnet arm resource id.
	DelegatedSubnetResourceReference *genruntime.ResourceReference `armReference:"DelegatedSubnetResourceId" json:"delegatedSubnetResourceReference,omitempty"`

	//PrivateDnsZoneArmResourceReference: private dns zone arm resource id.
	PrivateDnsZoneArmResourceReference *genruntime.ResourceReference `armReference:"PrivateDnsZoneArmResourceId" json:"privateDnsZoneArmResourceReference,omitempty"`
	PropertyBag                        genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210601.Network_Status
type Network_Status struct {
	DelegatedSubnetResourceId   *string                `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string                `json:"privateDnsZoneArmResourceId,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetworkAccess         *string                `json:"publicNetworkAccess,omitempty"`
}

//Storage version of v1alpha1api20210601.Sku
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku
type Sku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210601.Sku_Status
type Sku_Status struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

//Storage version of v1alpha1api20210601.Storage
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage
type Storage struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageSizeGB *int                   `json:"storageSizeGB,omitempty"`
}

//Storage version of v1alpha1api20210601.Storage_Status
type Storage_Status struct {
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StorageSizeGB *int                   `json:"storageSizeGB,omitempty"`
}

//Storage version of v1alpha1api20210601.SystemData_Status
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&FlexibleServer{}, &FlexibleServerList{})
}
