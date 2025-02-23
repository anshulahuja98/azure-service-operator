// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkInterfaces_SpecARM struct {
	//ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocationARM `json:"extendedLocation,omitempty"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Properties of the network interface.
	Properties NetworkInterfaces_Spec_PropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkInterfaces_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (interfaces NetworkInterfaces_SpecARM) GetAPIVersion() string {
	return "2020-11-01"
}

// GetName returns the Name of the resource
func (interfaces NetworkInterfaces_SpecARM) GetName() string {
	return interfaces.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkInterfaces"
func (interfaces NetworkInterfaces_SpecARM) GetType() string {
	return "Microsoft.Network/networkInterfaces"
}

type NetworkInterfaces_Spec_PropertiesARM struct {
	//DnsSettings: The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettingsARM `json:"dnsSettings,omitempty"`

	//EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	//EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	//IpConfigurations: A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaces_Spec_Properties_IpConfigurationsARM `json:"ipConfigurations"`

	//NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *SubResourceARM `json:"networkSecurityGroup,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings
type NetworkInterfaceDnsSettingsARM struct {
	//DnsServers: List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution.
	//'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
	DnsServers []string `json:"dnsServers,omitempty"`

	//InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
	//network.
	InternalDnsNameLabel *string `json:"internalDnsNameLabel,omitempty"`
}

type NetworkInterfaces_Spec_Properties_IpConfigurationsARM struct {
	//Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name"`

	//Properties: Network interface IP configuration properties.
	Properties *NetworkInterfaceIPConfigurationPropertiesFormatARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource
type SubResourceARM struct {
	Id string `json:"id"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceIPConfigurationPropertiesFormat
type NetworkInterfaceIPConfigurationPropertiesFormatARM struct {
	//ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.
	ApplicationGatewayBackendAddressPools []SubResourceARM `json:"applicationGatewayBackendAddressPools,omitempty"`

	//ApplicationSecurityGroups: Application security groups in which the IP configuration is included.
	ApplicationSecurityGroups []SubResourceARM `json:"applicationSecurityGroups,omitempty"`

	//LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.
	LoadBalancerBackendAddressPools []SubResourceARM `json:"loadBalancerBackendAddressPools,omitempty"`

	//LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.
	LoadBalancerInboundNatRules []SubResourceARM `json:"loadBalancerInboundNatRules,omitempty"`

	//Primary: Whether this is a primary customer address on the network interface.
	Primary *bool `json:"primary,omitempty"`

	//PrivateIPAddress: Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAddressVersion `json:"privateIPAddressVersion,omitempty"`

	//PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *NetworkInterfaceIPConfigurationPropertiesFormatPrivateIPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	//PublicIPAddress: Public IP address bound to the IP configuration.
	PublicIPAddress *SubResourceARM `json:"publicIPAddress,omitempty"`

	//Subnet: Subnet bound to the IP configuration.
	Subnet *SubResourceARM `json:"subnet,omitempty"`

	//VirtualNetworkTaps: The reference to Virtual Network Taps.
	VirtualNetworkTaps []SubResourceARM `json:"virtualNetworkTaps,omitempty"`
}
