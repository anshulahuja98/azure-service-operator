// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type EventSubscriptions_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should
	//use alphanumeric letters only.
	Name string `json:"name"`

	//Properties: Properties of the Event Subscription.
	Properties EventSubscriptionPropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &EventSubscriptions_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscriptions EventSubscriptions_SpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (subscriptions EventSubscriptions_SpecARM) GetName() string {
	return subscriptions.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscriptions EventSubscriptions_SpecARM) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventSubscriptionProperties
type EventSubscriptionPropertiesARM struct {
	//DeadLetterDestination: Information about the dead letter destination for an event subscription. To configure a
	//deadletter destination, do not directly instantiate an object of this class. Instead, instantiate an object of a derived
	//class. Currently, StorageBlobDeadLetterDestination is the only class that derives from this class.
	DeadLetterDestination *StorageBlobDeadLetterDestinationARM `json:"deadLetterDestination,omitempty"`

	//Destination: Information about the destination for an event subscription.
	Destination *EventSubscriptionDestinationARM `json:"destination,omitempty"`

	//EventDeliverySchema: The event delivery schema for the event subscription.
	EventDeliverySchema *EventSubscriptionPropertiesEventDeliverySchema `json:"eventDeliverySchema,omitempty"`

	//ExpirationTimeUtc: Expiration time of the event subscription.
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty"`

	//Filter: Filter for the Event Subscription.
	Filter *EventSubscriptionFilterARM `json:"filter,omitempty"`

	//Labels: List of user defined labels.
	Labels []string `json:"labels,omitempty"`

	//RetryPolicy: Information about the retry policy for an event subscription.
	RetryPolicy *RetryPolicyARM `json:"retryPolicy,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventSubscriptionDestination
type EventSubscriptionDestinationARM struct {
	//AzureFunction: Mutually exclusive with all other properties
	AzureFunction *AzureFunctionEventSubscriptionDestinationARM `json:"azureFunctionEventSubscriptionDestination,omitempty"`

	//EventHub: Mutually exclusive with all other properties
	EventHub *EventHubEventSubscriptionDestinationARM `json:"eventHubEventSubscriptionDestination,omitempty"`

	//HybridConnection: Mutually exclusive with all other properties
	HybridConnection *HybridConnectionEventSubscriptionDestinationARM `json:"hybridConnectionEventSubscriptionDestination,omitempty"`

	//ServiceBusQueue: Mutually exclusive with all other properties
	ServiceBusQueue *ServiceBusQueueEventSubscriptionDestinationARM `json:"serviceBusQueueEventSubscriptionDestination,omitempty"`

	//ServiceBusTopic: Mutually exclusive with all other properties
	ServiceBusTopic *ServiceBusTopicEventSubscriptionDestinationARM `json:"serviceBusTopicEventSubscriptionDestination,omitempty"`

	//StorageQueue: Mutually exclusive with all other properties
	StorageQueue *StorageQueueEventSubscriptionDestinationARM `json:"storageQueueEventSubscriptionDestination,omitempty"`

	//WebHook: Mutually exclusive with all other properties
	WebHook *WebHookEventSubscriptionDestinationARM `json:"webHookEventSubscriptionDestination,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because EventSubscriptionDestinationARM represents a discriminated union (JSON OneOf)
func (destination EventSubscriptionDestinationARM) MarshalJSON() ([]byte, error) {
	if destination.AzureFunction != nil {
		return json.Marshal(destination.AzureFunction)
	}
	if destination.EventHub != nil {
		return json.Marshal(destination.EventHub)
	}
	if destination.HybridConnection != nil {
		return json.Marshal(destination.HybridConnection)
	}
	if destination.ServiceBusQueue != nil {
		return json.Marshal(destination.ServiceBusQueue)
	}
	if destination.ServiceBusTopic != nil {
		return json.Marshal(destination.ServiceBusTopic)
	}
	if destination.StorageQueue != nil {
		return json.Marshal(destination.StorageQueue)
	}
	if destination.WebHook != nil {
		return json.Marshal(destination.WebHook)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the EventSubscriptionDestinationARM
func (destination *EventSubscriptionDestinationARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "AzureFunction" {
		destination.AzureFunction = &AzureFunctionEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.AzureFunction)
	}
	if discriminator == "EventHub" {
		destination.EventHub = &EventHubEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.EventHub)
	}
	if discriminator == "HybridConnection" {
		destination.HybridConnection = &HybridConnectionEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.HybridConnection)
	}
	if discriminator == "ServiceBusQueue" {
		destination.ServiceBusQueue = &ServiceBusQueueEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.ServiceBusQueue)
	}
	if discriminator == "ServiceBusTopic" {
		destination.ServiceBusTopic = &ServiceBusTopicEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.ServiceBusTopic)
	}
	if discriminator == "StorageQueue" {
		destination.StorageQueue = &StorageQueueEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.StorageQueue)
	}
	if discriminator == "WebHook" {
		destination.WebHook = &WebHookEventSubscriptionDestinationARM{}
		return json.Unmarshal(data, destination.WebHook)
	}

	// No error
	return nil
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventSubscriptionFilter
type EventSubscriptionFilterARM struct {
	//AdvancedFilters: An array of advanced filters that are used for filtering event subscriptions.
	AdvancedFilters []AdvancedFilterARM `json:"advancedFilters,omitempty"`

	//IncludedEventTypes: A list of applicable event types that need to be part of the event subscription. If it is desired to
	//subscribe to all default event types, set the IncludedEventTypes to null.
	IncludedEventTypes []string `json:"includedEventTypes,omitempty"`

	//IsSubjectCaseSensitive: Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	//should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `json:"isSubjectCaseSensitive,omitempty"`

	//SubjectBeginsWith: An optional string to filter events for an event subscription based on a resource path prefix.
	//The format of this depends on the publisher of the events.
	//Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty"`

	//SubjectEndsWith: An optional string to filter events for an event subscription based on a resource path suffix.
	//Wildcard characters are not supported in this path.
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/RetryPolicy
type RetryPolicyARM struct {
	//EventTimeToLiveInMinutes: Time To Live (in minutes) for events.
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`

	//MaxDeliveryAttempts: Maximum number of delivery retry attempts for events.
	MaxDeliveryAttempts *int `json:"maxDeliveryAttempts,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageBlobDeadLetterDestination
type StorageBlobDeadLetterDestinationARM struct {
	EndpointType StorageBlobDeadLetterDestinationEndpointType `json:"endpointType"`

	//Properties: Properties of the storage blob based dead letter destination.
	Properties *StorageBlobDeadLetterDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AdvancedFilter
type AdvancedFilterARM struct {
	//BoolEquals: Mutually exclusive with all other properties
	BoolEquals *AdvancedFilter_BoolEqualsARM `json:"boolEqualsAdvancedFilter,omitempty"`

	//NumberGreaterThan: Mutually exclusive with all other properties
	NumberGreaterThan *AdvancedFilter_NumberGreaterThanARM `json:"numberGreaterThanAdvancedFilter,omitempty"`

	//NumberGreaterThanOrEquals: Mutually exclusive with all other properties
	NumberGreaterThanOrEquals *AdvancedFilter_NumberGreaterThanOrEqualsARM `json:"numberGreaterThanOrEqualsAdvancedFilter,omitempty"`

	//NumberIn: Mutually exclusive with all other properties
	NumberIn *AdvancedFilter_NumberInARM `json:"numberInAdvancedFilter,omitempty"`

	//NumberLessThan: Mutually exclusive with all other properties
	NumberLessThan *AdvancedFilter_NumberLessThanARM `json:"numberLessThanAdvancedFilter,omitempty"`

	//NumberLessThanOrEquals: Mutually exclusive with all other properties
	NumberLessThanOrEquals *AdvancedFilter_NumberLessThanOrEqualsARM `json:"numberLessThanOrEqualsAdvancedFilter,omitempty"`

	//NumberNotIn: Mutually exclusive with all other properties
	NumberNotIn *AdvancedFilter_NumberNotInARM `json:"numberNotInAdvancedFilter,omitempty"`

	//StringBeginsWith: Mutually exclusive with all other properties
	StringBeginsWith *AdvancedFilter_StringBeginsWithARM `json:"stringBeginsWithAdvancedFilter,omitempty"`

	//StringContains: Mutually exclusive with all other properties
	StringContains *AdvancedFilter_StringContainsARM `json:"stringContainsAdvancedFilter,omitempty"`

	//StringEndsWith: Mutually exclusive with all other properties
	StringEndsWith *AdvancedFilter_StringEndsWithARM `json:"stringEndsWithAdvancedFilter,omitempty"`

	//StringIn: Mutually exclusive with all other properties
	StringIn *AdvancedFilter_StringInARM `json:"stringInAdvancedFilter,omitempty"`

	//StringNotIn: Mutually exclusive with all other properties
	StringNotIn *AdvancedFilter_StringNotInARM `json:"stringNotInAdvancedFilter,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AdvancedFilterARM represents a discriminated union (JSON OneOf)
func (filter AdvancedFilterARM) MarshalJSON() ([]byte, error) {
	if filter.BoolEquals != nil {
		return json.Marshal(filter.BoolEquals)
	}
	if filter.NumberGreaterThan != nil {
		return json.Marshal(filter.NumberGreaterThan)
	}
	if filter.NumberGreaterThanOrEquals != nil {
		return json.Marshal(filter.NumberGreaterThanOrEquals)
	}
	if filter.NumberIn != nil {
		return json.Marshal(filter.NumberIn)
	}
	if filter.NumberLessThan != nil {
		return json.Marshal(filter.NumberLessThan)
	}
	if filter.NumberLessThanOrEquals != nil {
		return json.Marshal(filter.NumberLessThanOrEquals)
	}
	if filter.NumberNotIn != nil {
		return json.Marshal(filter.NumberNotIn)
	}
	if filter.StringBeginsWith != nil {
		return json.Marshal(filter.StringBeginsWith)
	}
	if filter.StringContains != nil {
		return json.Marshal(filter.StringContains)
	}
	if filter.StringEndsWith != nil {
		return json.Marshal(filter.StringEndsWith)
	}
	if filter.StringIn != nil {
		return json.Marshal(filter.StringIn)
	}
	if filter.StringNotIn != nil {
		return json.Marshal(filter.StringNotIn)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the AdvancedFilterARM
func (filter *AdvancedFilterARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["operatorType"]
	if discriminator == "BoolEquals" {
		filter.BoolEquals = &AdvancedFilter_BoolEqualsARM{}
		return json.Unmarshal(data, filter.BoolEquals)
	}
	if discriminator == "NumberGreaterThan" {
		filter.NumberGreaterThan = &AdvancedFilter_NumberGreaterThanARM{}
		return json.Unmarshal(data, filter.NumberGreaterThan)
	}
	if discriminator == "NumberGreaterThanOrEquals" {
		filter.NumberGreaterThanOrEquals = &AdvancedFilter_NumberGreaterThanOrEqualsARM{}
		return json.Unmarshal(data, filter.NumberGreaterThanOrEquals)
	}
	if discriminator == "NumberIn" {
		filter.NumberIn = &AdvancedFilter_NumberInARM{}
		return json.Unmarshal(data, filter.NumberIn)
	}
	if discriminator == "NumberLessThan" {
		filter.NumberLessThan = &AdvancedFilter_NumberLessThanARM{}
		return json.Unmarshal(data, filter.NumberLessThan)
	}
	if discriminator == "NumberLessThanOrEquals" {
		filter.NumberLessThanOrEquals = &AdvancedFilter_NumberLessThanOrEqualsARM{}
		return json.Unmarshal(data, filter.NumberLessThanOrEquals)
	}
	if discriminator == "NumberNotIn" {
		filter.NumberNotIn = &AdvancedFilter_NumberNotInARM{}
		return json.Unmarshal(data, filter.NumberNotIn)
	}
	if discriminator == "StringBeginsWith" {
		filter.StringBeginsWith = &AdvancedFilter_StringBeginsWithARM{}
		return json.Unmarshal(data, filter.StringBeginsWith)
	}
	if discriminator == "StringContains" {
		filter.StringContains = &AdvancedFilter_StringContainsARM{}
		return json.Unmarshal(data, filter.StringContains)
	}
	if discriminator == "StringEndsWith" {
		filter.StringEndsWith = &AdvancedFilter_StringEndsWithARM{}
		return json.Unmarshal(data, filter.StringEndsWith)
	}
	if discriminator == "StringIn" {
		filter.StringIn = &AdvancedFilter_StringInARM{}
		return json.Unmarshal(data, filter.StringIn)
	}
	if discriminator == "StringNotIn" {
		filter.StringNotIn = &AdvancedFilter_StringNotInARM{}
		return json.Unmarshal(data, filter.StringNotIn)
	}

	// No error
	return nil
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AzureFunctionEventSubscriptionDestination
type AzureFunctionEventSubscriptionDestinationARM struct {
	EndpointType AzureFunctionEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties that represent the Azure Function destination of an event subscription.
	Properties *AzureFunctionEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventHubEventSubscriptionDestination
type EventHubEventSubscriptionDestinationARM struct {
	EndpointType EventHubEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties for a event hub destination.
	Properties *EventHubEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/HybridConnectionEventSubscriptionDestination
type HybridConnectionEventSubscriptionDestinationARM struct {
	EndpointType HybridConnectionEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties for a hybrid connection destination.
	Properties *HybridConnectionEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusQueueEventSubscriptionDestination
type ServiceBusQueueEventSubscriptionDestinationARM struct {
	EndpointType ServiceBusQueueEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties that represent the Service Bus destination of an event subscription.
	Properties *ServiceBusQueueEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusTopicEventSubscriptionDestination
type ServiceBusTopicEventSubscriptionDestinationARM struct {
	EndpointType ServiceBusTopicEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties that represent the Service Bus Topic destination of an event subscription.
	Properties *ServiceBusTopicEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageBlobDeadLetterDestinationProperties
type StorageBlobDeadLetterDestinationPropertiesARM struct {
	//BlobContainerName: The name of the Storage blob container that is the destination of the deadletter events
	BlobContainerName *string `json:"blobContainerName,omitempty"`
	ResourceId        *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageQueueEventSubscriptionDestination
type StorageQueueEventSubscriptionDestinationARM struct {
	EndpointType StorageQueueEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: The properties for a storage queue destination.
	Properties *StorageQueueEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/WebHookEventSubscriptionDestination
type WebHookEventSubscriptionDestinationARM struct {
	EndpointType WebHookEventSubscriptionDestinationEndpointType `json:"endpointType"`

	//Properties: Information about the webhook destination properties for an event subscription.
	Properties *WebHookEventSubscriptionDestinationPropertiesARM `json:"properties,omitempty"`
}

type AdvancedFilter_BoolEqualsARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                              `json:"key,omitempty"`
	OperatorType AdvancedFilterBoolEqualsOperatorType `json:"operatorType"`

	//Value: The boolean filter value.
	Value *bool `json:"value,omitempty"`
}

type AdvancedFilter_NumberGreaterThanARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                     `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberGreaterThanOperatorType `json:"operatorType"`

	//Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type AdvancedFilter_NumberGreaterThanOrEqualsARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                             `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberGreaterThanOrEqualsOperatorType `json:"operatorType"`

	//Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type AdvancedFilter_NumberInARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                            `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberInOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

type AdvancedFilter_NumberLessThanARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberLessThanOperatorType `json:"operatorType"`

	//Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type AdvancedFilter_NumberLessThanOrEqualsARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                          `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberLessThanOrEqualsOperatorType `json:"operatorType"`

	//Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type AdvancedFilter_NumberNotInARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                               `json:"key,omitempty"`
	OperatorType AdvancedFilterNumberNotInOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

type AdvancedFilter_StringBeginsWithARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                    `json:"key,omitempty"`
	OperatorType AdvancedFilterStringBeginsWithOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type AdvancedFilter_StringContainsARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterStringContainsOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type AdvancedFilter_StringEndsWithARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                                  `json:"key,omitempty"`
	OperatorType AdvancedFilterStringEndsWithOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type AdvancedFilter_StringInARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                            `json:"key,omitempty"`
	OperatorType AdvancedFilterStringInOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type AdvancedFilter_StringNotInARM struct {
	//Key: The field/property in the event based on which you want to filter.
	Key          *string                               `json:"key,omitempty"`
	OperatorType AdvancedFilterStringNotInOperatorType `json:"operatorType"`

	//Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AzureFunctionEventSubscriptionDestinationProperties
type AzureFunctionEventSubscriptionDestinationPropertiesARM struct {
	//MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	//PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
	ResourceId                    *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventHubEventSubscriptionDestinationProperties
type EventHubEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/HybridConnectionEventSubscriptionDestinationProperties
type HybridConnectionEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusQueueEventSubscriptionDestinationProperties
type ServiceBusQueueEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusTopicEventSubscriptionDestinationProperties
type ServiceBusTopicEventSubscriptionDestinationPropertiesARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageQueueEventSubscriptionDestinationProperties
type StorageQueueEventSubscriptionDestinationPropertiesARM struct {
	//QueueName: The name of the Storage queue under a storage account that is the destination of an event subscription.
	QueueName  *string `json:"queueName,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/WebHookEventSubscriptionDestinationProperties
type WebHookEventSubscriptionDestinationPropertiesARM struct {
	//AzureActiveDirectoryApplicationIdOrUri: The Azure Active Directory Application ID or URI to get the access token that
	//will be included as the bearer token in delivery requests.
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`

	//AzureActiveDirectoryTenantId: The Azure Active Directory Tenant ID to get the access token that will be included as the
	//bearer token in delivery requests.
	AzureActiveDirectoryTenantId *string `json:"azureActiveDirectoryTenantId,omitempty"`

	//EndpointUrl: The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	//MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	//PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int `json:"preferredBatchSizeInKilobytes,omitempty"`
}
