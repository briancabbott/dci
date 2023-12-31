//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import "time"

// ArmAsyncOperation - Arm async operation class. Ref: https://docs.microsoft.com/en-us/azure/azure-resource-manager/management/async-operations.
type ArmAsyncOperation struct {
	// Status of the async operation.
	Status *string
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// DetailedInformation - Model to capture detailed information for farmBeatsExtensions.
type DetailedInformation struct {
	// List of apiInputParameters.
	APIInputParameters []*string

	// ApiName available for the farmBeatsExtension.
	APIName *string

	// List of customParameters.
	CustomParameters []*string

	// List of platformParameters.
	PlatformParameters []*string

	// Unit systems info for the data provider.
	UnitsSupported *UnitSystemsInfo
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Extension resource.
type Extension struct {
	// Extension resource properties.
	Properties *ExtensionProperties

	// READ-ONLY; The ETag value to implement optimistic concurrency.
	ETag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ExtensionListResponse - Paged response contains list of requested objects and a URL link to get the next set of results.
type ExtensionListResponse struct {
	// List of requested objects.
	Value []*Extension

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string
}

// ExtensionProperties - Extension resource properties.
type ExtensionProperties struct {
	// READ-ONLY; Extension api docs link.
	ExtensionAPIDocsLink *string

	// READ-ONLY; Extension auth link.
	ExtensionAuthLink *string

	// READ-ONLY; Extension category. e.g. weather/sensor/satellite.
	ExtensionCategory *string

	// READ-ONLY; Extension Id.
	ExtensionID *string

	// READ-ONLY; Installed extension version.
	InstalledExtensionVersion *string
}

// FarmBeats ARM Resource.
type FarmBeats struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Identity for the resource.
	Identity *Identity

	// FarmBeats ARM Resource properties.
	Properties *FarmBeatsProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FarmBeatsExtension - FarmBeats extension resource.
type FarmBeatsExtension struct {
	// FarmBeatsExtension properties.
	Properties *FarmBeatsExtensionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FarmBeatsExtensionListResponse - Paged response contains list of requested objects and a URL link to get the next set of
// results.
type FarmBeatsExtensionListResponse struct {
	// List of requested objects.
	Value []*FarmBeatsExtension

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string
}

// FarmBeatsExtensionProperties - FarmBeatsExtension properties.
type FarmBeatsExtensionProperties struct {
	// READ-ONLY; Textual description.
	Description *string

	// READ-ONLY; Detailed information which shows summary of requested data. Used in descriptive get extension metadata call.
	// Information for weather category per api included are apisSupported, customParameters,
	// PlatformParameters and Units supported.
	DetailedInformation []*DetailedInformation

	// READ-ONLY; FarmBeatsExtension api docs link.
	ExtensionAPIDocsLink *string

	// READ-ONLY; FarmBeatsExtension auth link.
	ExtensionAuthLink *string

	// READ-ONLY; Category of the extension. e.g. weather/sensor/satellite.
	ExtensionCategory *string

	// READ-ONLY; FarmBeatsExtension ID.
	FarmBeatsExtensionID *string

	// READ-ONLY; FarmBeatsExtension name.
	FarmBeatsExtensionName *string

	// READ-ONLY; FarmBeatsExtension version.
	FarmBeatsExtensionVersion *string

	// READ-ONLY; Publisher ID.
	PublisherID *string

	// READ-ONLY; Target ResourceType of the farmBeatsExtension.
	TargetResourceType *string
}

// FarmBeatsListResponse - Paged response contains list of requested objects and a URL link to get the next set of results.
type FarmBeatsListResponse struct {
	// List of requested objects.
	Value []*FarmBeats

	// READ-ONLY; Continuation link (absolute URI) to the next page of results in the list.
	NextLink *string
}

// FarmBeatsProperties - FarmBeats ARM Resource properties.
type FarmBeatsProperties struct {
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *PublicNetworkAccess

	// Sensor integration request model.
	SensorIntegration *SensorIntegration

	// READ-ONLY; Uri of the FarmBeats instance.
	InstanceURI *string

	// READ-ONLY; The Private Endpoint Connection resource.
	PrivateEndpointConnections *PrivateEndpointConnection

	// READ-ONLY; FarmBeats instance provisioning state.
	ProvisioningState *ProvisioningState
}

// FarmBeatsUpdateProperties - FarmBeats ARM Resource properties.
type FarmBeatsUpdateProperties struct {
	// Property to allow or block public traffic for an Azure FarmBeats resource.
	PublicNetworkAccess *PublicNetworkAccess

	// Sensor integration request model.
	SensorIntegration *SensorIntegration
}

// FarmBeatsUpdateRequestModel - FarmBeats update request.
type FarmBeatsUpdateRequestModel struct {
	// Identity for the resource.
	Identity *Identity

	// Geo-location where the resource lives.
	Location *string

	// FarmBeats ARM Resource properties.
	Properties *FarmBeatsUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *string

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SensorIntegration - Sensor integration request model.
type SensorIntegration struct {
	// Sensor integration enable state. Allowed values are True, None
	Enabled *string

	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows
	// the OData error response format.).
	ProvisioningInfo *ErrorResponse

	// READ-ONLY; Sensor integration instance provisioning state.
	ProvisioningState *ProvisioningState
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UnitSystemsInfo - Unit systems info for the data provider.
type UnitSystemsInfo struct {
	// REQUIRED; UnitSystem key sent as part of ProviderInput.
	Key *string

	// REQUIRED; List of unit systems supported by this data provider.
	Values []*string
}
