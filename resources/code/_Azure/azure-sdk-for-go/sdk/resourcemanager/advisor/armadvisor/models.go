//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

import "time"

// ARMErrorResponseBody - ARM error response body.
type ARMErrorResponseBody struct {
	// Gets or sets the string that can be used to programmatically identify the error.
	Code *string

	// Gets or sets the string that describes the error in detail and provides debugging information.
	Message *string
}

type ArmErrorResponse struct {
	// ARM error response body.
	Error *ARMErrorResponseBody
}

// ConfigData - The Advisor configuration data structure.
type ConfigData struct {
	// The Advisor configuration data structure.
	Properties *ConfigDataProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ConfigDataProperties - Configuration data properties
type ConfigDataProperties struct {
	// Advisor digest configuration. Valid only for subscriptions
	Digests []*DigestConfig

	// Exclude the resource from Advisor evaluations. Valid values: False (default) or True.
	Exclude *bool

	// Minimum percentage threshold for Advisor low CPU utilization evaluation. Valid only for subscriptions. Valid values: 5
	// (default), 10, 15 or 20.
	LowCPUThreshold *CPUThreshold
}

// ConfigurationListResult - The list of Advisor configurations.
type ConfigurationListResult struct {
	// The link used to get the next page of configurations.
	NextLink *string

	// The list of configurations.
	Value []*ConfigData
}

// DigestConfig - Advisor Digest configuration entity
type DigestConfig struct {
	// Action group resource id used by digest.
	ActionGroupResourceID *string

	// Categories to send digest for. If categories are not provided, then digest will be sent for all categories.
	Categories []*Category

	// Frequency that digest will be triggered, in days. Value must be between 7 and 30 days inclusive.
	Frequency *int32

	// Language for digest content body. Value must be ISO 639-1 code for one of Azure portal supported languages. Otherwise,
	// it will be converted into one. Default value is English (en).
	Language *string

	// Name of digest configuration. Value is case-insensitive and must be unique within a subscription.
	Name *string

	// State of digest configuration.
	State *DigestConfigState
}

// MetadataEntity - The metadata entity contract.
type MetadataEntity struct {
	// The resource Id of the metadata entity.
	ID *string

	// The name of the metadata entity.
	Name *string

	// The metadata entity properties.
	Properties *MetadataEntityProperties

	// The type of the metadata entity.
	Type *string
}

// MetadataEntityListResult - The list of metadata entities
type MetadataEntityListResult struct {
	// The link used to get the next page of metadata.
	NextLink *string

	// The list of metadata entities.
	Value []*MetadataEntity
}

// MetadataEntityProperties - The metadata entity properties
type MetadataEntityProperties struct {
	// The list of scenarios applicable to this metadata entity.
	ApplicableScenarios []*Scenario

	// The list of keys on which this entity depends on.
	DependsOn []*string

	// The display name.
	DisplayName *string

	// The list of supported values.
	SupportedValues []*MetadataSupportedValueDetail
}

// MetadataSupportedValueDetail - The metadata supported value detail.
type MetadataSupportedValueDetail struct {
	// The display name.
	DisplayName *string

	// The id.
	ID *string
}

// OperationDisplayInfo - The operation supported by Advisor.
type OperationDisplayInfo struct {
	// The description of the operation.
	Description *string

	// The action that users can perform, based on their permission level.
	Operation *string

	// Service provider: Microsoft Advisor.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// OperationEntity - The operation supported by Advisor.
type OperationEntity struct {
	// The operation supported by Advisor.
	Display *OperationDisplayInfo

	// Operation name: {provider}/{resource}/{operation}.
	Name *string
}

// OperationEntityListResult - The list of Advisor operations.
type OperationEntityListResult struct {
	// The link used to get the next page of operations.
	NextLink *string

	// The list of operations.
	Value []*OperationEntity
}

// RecommendationProperties - The properties of the recommendation.
type RecommendationProperties struct {
	// The list of recommended actions to implement recommendation.
	Actions []map[string]any

	// The category of the recommendation.
	Category *Category

	// The detailed description of recommendation.
	Description *string

	// The recommendation metadata properties exposed to customer to provide additional information.
	ExposedMetadataProperties map[string]any

	// Extended properties
	ExtendedProperties map[string]*string

	// The business impact of the recommendation.
	Impact *Impact

	// The resource type identified by Advisor.
	ImpactedField *string

	// The resource identified by Advisor.
	ImpactedValue *string

	// The label of recommendation.
	Label *string

	// The most recent time that Advisor checked the validity of the recommendation.
	LastUpdated *time.Time

	// The link to learn more about recommendation and generation logic.
	LearnMoreLink *string

	// The recommendation metadata.
	Metadata map[string]any

	// The potential benefit of implementing recommendation.
	PotentialBenefits *string

	// The recommendation-type GUID.
	RecommendationTypeID *string

	// The automated way to apply recommendation.
	Remediation map[string]any

	// Metadata of resource that was assessed
	ResourceMetadata *ResourceMetadata

	// The potential risk of not implementing the recommendation.
	Risk *Risk

	// A summary of the recommendation.
	ShortDescription *ShortDescription

	// The list of snoozed and dismissed rules for the recommendation.
	SuppressionIDs []*string
}

// Resource - An Azure resource.
type Resource struct {
	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ResourceMetadata - Recommendation resource metadata
type ResourceMetadata struct {
	// The action to view resource.
	Action map[string]any

	// The plural user friendly name of resource type. eg: virtual machines
	Plural *string

	// Azure resource Id of the assessed resource
	ResourceID *string

	// The singular user friendly name of resource type. eg: virtual machine
	Singular *string

	// Source from which recommendation is generated
	Source *string
}

// ResourceRecommendationBase - Advisor Recommendation.
type ResourceRecommendationBase struct {
	// The properties of the recommendation.
	Properties *RecommendationProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// ResourceRecommendationBaseListResult - The list of Advisor recommendations.
type ResourceRecommendationBaseListResult struct {
	// The link used to get the next page of recommendations.
	NextLink *string

	// The list of recommendations.
	Value []*ResourceRecommendationBase
}

// ShortDescription - A summary of the recommendation.
type ShortDescription struct {
	// The issue or opportunity identified by the recommendation.
	Problem *string

	// The remediation action suggested by the recommendation.
	Solution *string
}

// SuppressionContract - The details of the snoozed or dismissed rule; for example, the duration, name, and GUID associated
// with the rule.
type SuppressionContract struct {
	// The properties of the suppression.
	Properties *SuppressionProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// SuppressionContractListResult - The list of Advisor suppressions.
type SuppressionContractListResult struct {
	// The link used to get the next page of suppressions.
	NextLink *string

	// The list of suppressions.
	Value []*SuppressionContract
}

// SuppressionProperties - The properties of the suppression.
type SuppressionProperties struct {
	// The GUID of the suppression.
	SuppressionID *string

	// The duration for which the suppression is valid.
	TTL *string

	// READ-ONLY; Gets or sets the expiration time stamp.
	ExpirationTimeStamp *time.Time
}
