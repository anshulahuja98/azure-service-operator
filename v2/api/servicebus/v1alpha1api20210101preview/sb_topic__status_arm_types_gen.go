// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

type SBTopic_StatusARM struct {
	//Id: Resource Id
	Id *string `json:"id,omitempty"`

	//Name: Resource name
	Name *string `json:"name,omitempty"`

	//Properties: Properties of topic resource.
	Properties *SBTopicProperties_StatusARM `json:"properties,omitempty"`

	//SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	//Type: Resource type
	Type *string `json:"type,omitempty"`
}

type SBTopicProperties_StatusARM struct {
	//AccessedAt: Last time the message was sent, or a request was received, for this topic.
	AccessedAt *string `json:"accessedAt,omitempty"`

	//AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
	//is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	//CountDetails: Message count details
	CountDetails *MessageCountDetails_StatusARM `json:"countDetails,omitempty"`

	//CreatedAt: Exact time the message was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	//DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
	//expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	//set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	//DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
	//history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	//EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	//EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
	//temporarily before writing it to persistent storage.
	EnableExpress *bool `json:"enableExpress,omitempty"`

	//EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty"`

	//MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
	//Default is 1024.
	MaxSizeInMegabytes *int `json:"maxSizeInMegabytes,omitempty"`

	//RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty"`

	//SizeInBytes: Size of the topic, in bytes.
	SizeInBytes *int `json:"sizeInBytes,omitempty"`

	//Status: Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus_Status `json:"status,omitempty"`

	//SubscriptionCount: Number of subscriptions.
	SubscriptionCount *int `json:"subscriptionCount,omitempty"`

	//SupportOrdering: Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `json:"supportOrdering,omitempty"`

	//UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}
