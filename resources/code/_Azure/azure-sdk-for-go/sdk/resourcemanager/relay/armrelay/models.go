//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrelay

import "time"

// AccessKeys - Namespace/Relay Connection String
type AccessKeys struct {
	// A string that describes the authorization rule.
	KeyName *string

	// Primary connection string of the created namespace authorization rule.
	PrimaryConnectionString *string

	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string

	// Secondary connection string of the created namespace authorization rule.
	SecondaryConnectionString *string

	// A base64-encoded 256-bit secondary key for signing and validating the SAS token.
	SecondaryKey *string
}

// AuthorizationRule - Single item in a List or Get AuthorizationRule operation
type AuthorizationRule struct {
	// Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRuleProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// AuthorizationRuleListResult - The response from the list namespace operation.
type AuthorizationRuleListResult struct {
	// Link to the next set of results. Not empty if value contains incomplete list of authorization rules.
	NextLink *string

	// Result of the list authorization rules operation.
	Value []*AuthorizationRule
}

// AuthorizationRuleProperties - Properties supplied to create or update AuthorizationRule
type AuthorizationRuleProperties struct {
	// REQUIRED; The rights associated with the rule.
	Rights []*AccessRights
}

// CheckNameAvailability - Description of the check name availability request properties.
type CheckNameAvailability struct {
	// REQUIRED; The namespace name to check for availability. The namespace name can contain only letters, numbers, and hyphens.
	// The namespace must start with a letter, and it must end with a letter or number.
	Name *string
}

// CheckNameAvailabilityResult - Description of the check name availability request properties.
type CheckNameAvailabilityResult struct {
	// Value indicating namespace is available. Returns true if the namespace is available; otherwise, false.
	NameAvailable *bool

	// The reason for unavailability of a namespace.
	Reason *UnavailableReason

	// READ-ONLY; The detailed info regarding the reason associated with the namespace.
	Message *string
}

// ConnectionState information.
type ConnectionState struct {
	// Description of the connection state.
	Description *string

	// Status of the connection.
	Status *PrivateLinkConnectionStatus
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

// HybridConnection - Description of hybrid connection resource.
type HybridConnection struct {
	// Properties of the HybridConnection.
	Properties *HybridConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// HybridConnectionListResult - The response of the list hybrid connection operation.
type HybridConnectionListResult struct {
	// Link to the next set of results. Not empty if value contains incomplete list hybrid connection operation.
	NextLink *string

	// Result of the list hybrid connections.
	Value []*HybridConnection
}

// HybridConnectionProperties - Properties of the HybridConnection.
type HybridConnectionProperties struct {
	// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
	RequiresClientAuthorization *bool

	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it
	// can be used to store descriptive data, such as a list of teams and their contact
	// information. Also, user-defined configuration settings can be stored.
	UserMetadata *string

	// READ-ONLY; The time the hybrid connection was created.
	CreatedAt *time.Time

	// READ-ONLY; The number of listeners for this hybrid connection. Note that min : 1 and max:25 are supported.
	ListenerCount *int32

	// READ-ONLY; The time the namespace was updated.
	UpdatedAt *time.Time
}

// NWRuleSetIPRules - The response from the List namespace operation.
type NWRuleSetIPRules struct {
	// The IP Filter Action
	Action *NetworkRuleIPAction

	// IP Mask
	IPMask *string
}

// Namespace - Description of a namespace resource.
type Namespace struct {
	// REQUIRED; Resource location.
	Location *string

	// Description of Relay namespace
	Properties *NamespaceProperties

	// SKU of the namespace.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// NamespaceListResult - The response from the list namespace operation.
type NamespaceListResult struct {
	// Link to the next set of results. Not empty if value contains incomplete list of namespaces.
	NextLink *string

	// Result of the list namespace operation.
	Value []*Namespace
}

// NamespaceProperties - Properties of the namespace.
type NamespaceProperties struct {
	// List of private endpoint connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// READ-ONLY; The time the namespace was created.
	CreatedAt *time.Time

	// READ-ONLY; Identifier for Azure Insights metrics.
	MetricID *string

	// READ-ONLY; Provisioning state of the Namespace.
	ProvisioningState *string

	// READ-ONLY; Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string

	// READ-ONLY; Status of the Namespace.
	Status *string

	// READ-ONLY; The time the namespace was updated.
	UpdatedAt *time.Time
}

// NetworkRuleSet - Description of topic resource.
type NetworkRuleSet struct {
	// NetworkRuleSet properties
	Properties *NetworkRuleSetProperties

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Resource type.
	Type *string
}

// NetworkRuleSetProperties - NetworkRuleSet properties
type NetworkRuleSetProperties struct {
	// Default Action for Network Rule Set
	DefaultAction *DefaultAction

	// List of IpRules
	IPRules []*NWRuleSetIPRules
}

// Operation - A Relay REST API operation
type Operation struct {
	// Properties of the operation
	Properties any

	// READ-ONLY; Display of the operation
	Display *OperationDisplay

	// READ-ONLY; Indicates whether the operation is a data action
	IsDataAction *bool

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}
	Name *string

	// READ-ONLY; Origin of the operation
	Origin *string
}

// OperationDisplay - Operation display payload
type OperationDisplay struct {
	// READ-ONLY; Localized friendly description for the operation
	Description *string

	// READ-ONLY; Localized friendly name for the operation
	Operation *string

	// READ-ONLY; Resource provider of the operation
	Provider *string

	// READ-ONLY; Resource of the operation
	Resource *string
}

// OperationListResult - Result of the request to list Relay operations. It contains a list of operations and a URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string

	// READ-ONLY; List of Relay operations supported by the Microsoft.EventHub resource provider.
	Value []*Operation
}

// PrivateEndpoint information.
type PrivateEndpoint struct {
	// The ARM identifier for Private Endpoint.
	ID *string
}

// PrivateEndpointConnection - Properties of the PrivateEndpointConnection.
type PrivateEndpointConnection struct {
	// Properties of the PrivateEndpointConnection.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// PrivateEndpointConnectionListResult - Result of the list of all private endpoint connections operation.
type PrivateEndpointConnectionListResult struct {
	// A link for the next page of private endpoint connection resources.
	NextLink *string

	// A collection of private endpoint connection resources.
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection resource.
type PrivateEndpointConnectionProperties struct {
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint

	// Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState

	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *EndPointProvisioningState
}

// PrivateLinkResource - Information of the private link resource.
type PrivateLinkResource struct {
	// Fully qualified identifier of the resource.
	ID *string

	// Name of the resource
	Name *string

	// Properties of the private link resource.
	Properties *PrivateLinkResourceProperties

	// Type of the resource
	Type *string
}

// PrivateLinkResourceProperties - Properties of PrivateLinkResource
type PrivateLinkResourceProperties struct {
	// The private link resource group id.
	GroupID *string

	// The private link resource required member names.
	RequiredMembers []*string

	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string
}

// PrivateLinkResourcesListResult - Result of the List private link resources operation.
type PrivateLinkResourcesListResult struct {
	// A link for the next page of private link resources.
	NextLink *string

	// A collection of private link resources
	Value []*PrivateLinkResource
}

// ProxyResource - Common fields that are returned in the response for all Azure Resource Manager resources
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// RegenerateAccessKeyParameters - Parameters supplied to the regenerate authorization rule operation, specifies which key
// needs to be reset.
type RegenerateAccessKeyParameters struct {
	// REQUIRED; The access key to regenerate.
	KeyType *KeyType

	// Optional. If the key value is provided, this is set to key type, or autogenerated key value set for key type.
	Key *string
}

// Resource - The resource definition.
type Resource struct {
	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// ResourceNamespacePatch - Definition of resource.
type ResourceNamespacePatch struct {
	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// SKU of the namespace.
type SKU struct {
	// REQUIRED; Name of this SKU.
	Name *SKUName

	// The tier of this SKU.
	Tier *SKUTier
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

// TrackedResource - Definition of resource.
type TrackedResource struct {
	// REQUIRED; Resource location.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// UpdateParameters - Description of a namespace resource.
type UpdateParameters struct {
	// Description of Relay namespace.
	Properties *NamespaceProperties

	// SKU of the namespace.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Resource ID.
	ID *string

	// READ-ONLY; Resource name.
	Name *string

	// READ-ONLY; Resource type.
	Type *string
}

// WcfRelay - Description of the WCF relay resource.
type WcfRelay struct {
	// Properties of the WCF relay.
	Properties *WcfRelayProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// WcfRelayProperties - Properties of the WCF relay.
type WcfRelayProperties struct {
	// WCF relay type.
	RelayType *Relaytype

	// Returns true if client authorization is needed for this relay; otherwise, false.
	RequiresClientAuthorization *bool

	// Returns true if transport security is needed for this relay; otherwise, false.
	RequiresTransportSecurity *bool

	// The usermetadata is a placeholder to store user-defined string data for the WCF Relay endpoint. For example, it can be
	// used to store descriptive data, such as list of teams and their contact
	// information. Also, user-defined configuration settings can be stored.
	UserMetadata *string

	// READ-ONLY; The time the WCF relay was created.
	CreatedAt *time.Time

	// READ-ONLY; Returns true if the relay is dynamic; otherwise, false.
	IsDynamic *bool

	// READ-ONLY; The number of listeners for this relay. Note that min :1 and max:25 are supported.
	ListenerCount *int32

	// READ-ONLY; The time the namespace was updated.
	UpdatedAt *time.Time
}

// WcfRelaysListResult - The response of the list WCF relay operation.
type WcfRelaysListResult struct {
	// Link to the next set of results. Not empty if value contains incomplete list of WCF relays.
	NextLink *string

	// Result of the list WCF relay operation.
	Value []*WcfRelay
}
