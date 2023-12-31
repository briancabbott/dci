//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

import "time"

// APIKey - API key payload which is exposed in the request/response of the resource provider.
type APIKey struct {
	// Gets or sets the API key name.
	KeyName *string

	// Gets or sets the API key value.
	Value *string
}

// APIKeyCollection - Collection of the API key payload which is exposed in the response of the resource provider.
type APIKeyCollection struct {
	// Gets or sets the collection of API key.
	Keys []*APIKey
}

// Consortium payload
type Consortium struct {
	// Gets or sets the blockchain member name.
	Name *string

	// Gets or sets the protocol for the consortium.
	Protocol *BlockchainProtocol
}

// ConsortiumCollection - Collection of the consortium payload.
type ConsortiumCollection struct {
	// Gets or sets the collection of consortiums.
	Value []*Consortium
}

// ConsortiumMember - Consortium approval
type ConsortiumMember struct {
	// Gets the consortium member modified date.
	DateModified *time.Time

	// Gets the consortium member display name.
	DisplayName *string

	// Gets the consortium member join date.
	JoinDate *time.Time

	// Gets the consortium member name.
	Name *string

	// Gets the consortium member role.
	Role *string

	// Gets the consortium member status.
	Status *string

	// Gets the consortium member subscription id.
	SubscriptionID *string
}

// ConsortiumMemberCollection - Collection of consortium payload.
type ConsortiumMemberCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string

	// Gets or sets the collection of consortiums.
	Value []*ConsortiumMember
}

// FirewallRule - Ip range for firewall rules
type FirewallRule struct {
	// Gets or sets the end IP address of the firewall rule range.
	EndIPAddress *string

	// Gets or sets the name of the firewall rules.
	RuleName *string

	// Gets or sets the start IP address of the firewall rule range.
	StartIPAddress *string
}

// Member - Payload of the blockchain member which is exposed in the request/response of the resource provider.
type Member struct {
	// The GEO location of the blockchain service.
	Location *string

	// Gets or sets the blockchain member properties.
	Properties *MemberProperties

	// Gets or sets the blockchain member Sku.
	SKU *SKU

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string
}

// MemberCollection - Collection of the blockchain member payload which is exposed in the request/response of the resource
// provider.
type MemberCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string

	// Gets or sets the collection of blockchain members.
	Value []*Member
}

// MemberNodesSKU - Payload of the blockchain member nodes Sku for a blockchain member.
type MemberNodesSKU struct {
	// Gets or sets the nodes capacity.
	Capacity *int32
}

// MemberProperties - Payload of the blockchain member properties for a blockchain member.
type MemberProperties struct {
	// Gets or sets the consortium for the blockchain member.
	Consortium *string

	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string

	// Gets the display name of the member in the consortium.
	ConsortiumMemberDisplayName *string

	// Gets the role of the member in the consortium.
	ConsortiumRole *string

	// Gets or sets firewall rules
	FirewallRules []*FirewallRule

	// Sets the basic auth password of the blockchain member.
	Password *string

	// Gets or sets the blockchain protocol.
	Protocol *BlockchainProtocol

	// Gets or sets the blockchain validator nodes Sku.
	ValidatorNodesSKU *MemberNodesSKU

	// READ-ONLY; Gets the managed consortium management account address.
	ConsortiumManagementAccountAddress *string

	// READ-ONLY; Gets the dns endpoint of the blockchain member.
	DNS *string

	// READ-ONLY; Gets or sets the blockchain member provision state.
	ProvisioningState *BlockchainMemberProvisioningState

	// READ-ONLY; Gets the public key of the blockchain member (default transaction node).
	PublicKey *string

	// READ-ONLY; Gets the Ethereum root contract address of the blockchain.
	RootContractAddress *string

	// READ-ONLY; Gets the auth user name of the blockchain member.
	UserName *string
}

// MemberPropertiesUpdate - Update the payload of the blockchain member properties for a blockchain member.
type MemberPropertiesUpdate struct {
	// Sets the managed consortium management account password.
	ConsortiumManagementAccountPassword *string

	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule

	// Sets the transaction node dns endpoint basic auth password.
	Password *string
}

// MemberUpdate - Update the payload of the blockchain member which is exposed in the request/response of the resource provider.
type MemberUpdate struct {
	// Gets or sets the blockchain member update properties.
	Properties *MemberPropertiesUpdate

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string
}

// NameAvailability - Name availability payload which is exposed in the response of the resource provider.
type NameAvailability struct {
	// Gets or sets the message.
	Message *string

	// Gets or sets the value indicating whether the name is available.
	NameAvailable *bool

	// Gets or sets the name availability reason.
	Reason *NameAvailabilityReason
}

// NameAvailabilityRequest - Name availability request payload which is exposed in the request of the resource provider.
type NameAvailabilityRequest struct {
	// Gets or sets the name to check.
	Name *string

	// Gets or sets the type of the resource to check.
	Type *string
}

// OperationResult - Operation result payload which is exposed in the response of the resource provider.
type OperationResult struct {
	// Gets or sets the operation end time.
	EndTime *time.Time

	// Gets or sets the operation name.
	Name *string

	// Gets or sets the operation start time.
	StartTime *time.Time
}

// Resource - The core properties of the resources.
type Resource struct {
	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string
}

// ResourceProviderOperation - Operation payload which is exposed in the response of the resource provider.
type ResourceProviderOperation struct {
	// Gets or sets operation display
	Display *ResourceProviderOperationDisplay

	// Gets or sets a value indicating whether the operation is a data action or not.
	IsDataAction *bool

	// Gets or sets the operation name.
	Name *string

	// Gets or sets the origin.
	Origin *string
}

// ResourceProviderOperationCollection - Collection of operation payload which is exposed in the response of the resource
// provider.
type ResourceProviderOperationCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string

	// Gets or sets the collection of operations.
	Value []*ResourceProviderOperation
}

// ResourceProviderOperationDisplay - Operation display payload which is exposed in the response of the resource provider.
type ResourceProviderOperationDisplay struct {
	// Gets or sets the description of the provider for display purposes.
	Description *string

	// Gets or sets the name of the operation for display purposes.
	Operation *string

	// Gets or sets the name of the provider for display purposes.
	Provider *string

	// Gets or sets the name of the resource type for display purposes.
	Resource *string
}

// ResourceTypeSKU - Resource type Sku.
type ResourceTypeSKU struct {
	// Gets or sets the resource type
	ResourceType *string

	// Gets or sets the Skus
	SKUs []*SKUSetting
}

// ResourceTypeSKUCollection - Collection of the resource type Sku.
type ResourceTypeSKUCollection struct {
	// Gets or sets the collection of resource type Sku.
	Value []*ResourceTypeSKU
}

// SKU - Blockchain member Sku in payload
type SKU struct {
	// Gets or sets Sku name
	Name *string

	// Gets or sets Sku tier
	Tier *string
}

// SKUSetting - Sku Setting.
type SKUSetting struct {
	// Gets or sets the locations.
	Locations []*string

	// Gets or sets the Sku name.
	Name *string

	// Gets or sets the required features.
	RequiredFeatures []*string

	// Gets or sets the Sku tier.
	Tier *string
}

// TrackedResource - The resource model definition for a top level resource.
type TrackedResource struct {
	// The GEO location of the blockchain service.
	Location *string

	// Tags of the service which is a list of key value pairs that describes the resource.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string
}

// TransactionNode - Payload of the transaction node which is the request/response of the resource provider.
type TransactionNode struct {
	// Gets or sets the transaction node location.
	Location *string

	// Gets or sets the blockchain member properties.
	Properties *TransactionNodeProperties

	// READ-ONLY; Fully qualified resource Id of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the service - e.g. "Microsoft.Blockchain"
	Type *string
}

// TransactionNodeCollection - Collection of transaction node payload which is exposed in the request/response of the resource
// provider.
type TransactionNodeCollection struct {
	// Gets or sets the URL, that the client should use to fetch the next page (per server side paging). It's null for now, added
	// for future use.
	NextLink *string

	// Gets or sets the collection of transaction nodes.
	Value []*TransactionNode
}

// TransactionNodeProperties - Payload of transaction node properties payload in the transaction node payload.
type TransactionNodeProperties struct {
	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule

	// Sets the transaction node dns endpoint basic auth password.
	Password *string

	// READ-ONLY; Gets or sets the transaction node dns endpoint.
	DNS *string

	// READ-ONLY; Gets or sets the blockchain member provision state.
	ProvisioningState *NodeProvisioningState

	// READ-ONLY; Gets or sets the transaction node public key.
	PublicKey *string

	// READ-ONLY; Gets or sets the transaction node dns endpoint basic auth user name.
	UserName *string
}

// TransactionNodePropertiesUpdate - Update the payload of the transaction node properties in the transaction node payload.
type TransactionNodePropertiesUpdate struct {
	// Gets or sets the firewall rules.
	FirewallRules []*FirewallRule

	// Sets the transaction node dns endpoint basic auth password.
	Password *string
}

// TransactionNodeUpdate - Update the transaction node payload which is exposed in the request/response of the resource provider.
type TransactionNodeUpdate struct {
	// Gets or sets the transaction node update properties.
	Properties *TransactionNodePropertiesUpdate
}
