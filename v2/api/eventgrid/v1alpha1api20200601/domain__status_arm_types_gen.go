// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

type Domain_StatusARM struct {
	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	//Location: Location of the resource.
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the domain.
	Properties *DomainProperties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system metadata relating to Domain resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Tags: Tags of the resource.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type DomainProperties_StatusARM struct {
	//Endpoint: Endpoint for the domain.
	Endpoint *string `json:"endpoint,omitempty"`

	//InboundIpRules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered
	//only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRule_StatusARM `json:"inboundIpRules,omitempty"`

	//InputSchema: This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *DomainPropertiesStatusInputSchema `json:"inputSchema,omitempty"`

	//InputSchemaMapping: Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *InputSchemaMapping_StatusARM `json:"inputSchemaMapping,omitempty"`

	//MetricResourceId: Metric resource id for the domain.
	MetricResourceId *string `json:"metricResourceId,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	//ProvisioningState: Provisioning state of the domain.
	ProvisioningState *DomainPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//PublicNetworkAccess: This determines if traffic is allowed over public network. By default it is enabled.
	//You can further restrict to specific IPs by configuring <seealso
	//cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />
	PublicNetworkAccess *DomainPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

type InboundIpRule_StatusARM struct {
	//Action: Action to perform based on the match or no match of the IpMask.
	Action *InboundIpRuleStatusAction `json:"action,omitempty"`

	//IpMask: IP Address in CIDR notation e.g., 10.0.0.0/8.
	IpMask *string `json:"ipMask,omitempty"`
}

type InputSchemaMapping_StatusARM struct {
	//InputSchemaMappingType: Type of the custom mapping
	InputSchemaMappingType InputSchemaMappingStatusInputSchemaMappingType `json:"inputSchemaMappingType"`
}

type PrivateEndpointConnection_Status_Domain_SubResourceEmbeddedARM struct {
	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`
}
