//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

// AccountsClientAddRootCollectionAdminOptions contains the optional parameters for the AccountsClient.AddRootCollectionAdmin
// method.
type AccountsClientAddRootCollectionAdminOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the AccountsClient.BeginCreateOrUpdate method.
type AccountsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
type AccountsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
// method.
type AccountsClientListByResourceGroupOptions struct {
	// The skip token.
	SkipToken *string
}

// AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.NewListBySubscriptionPager
// method.
type AccountsClientListBySubscriptionOptions struct {
	// The skip token.
	SkipToken *string
}

// AccountsClientListKeysOptions contains the optional parameters for the AccountsClient.ListKeys method.
type AccountsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// DefaultAccountsClientGetOptions contains the optional parameters for the DefaultAccountsClient.Get method.
type DefaultAccountsClientGetOptions struct {
	// The Id of the scope object, for example if the scope is "Subscription" then it is the ID of that subscription.
	Scope *string
}

// DefaultAccountsClientRemoveOptions contains the optional parameters for the DefaultAccountsClient.Remove method.
type DefaultAccountsClientRemoveOptions struct {
	// The Id of the scope object, for example if the scope is "Subscription" then it is the ID of that subscription.
	Scope *string
}

// DefaultAccountsClientSetOptions contains the optional parameters for the DefaultAccountsClient.Set method.
type DefaultAccountsClientSetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
type PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
type PrivateEndpointConnectionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByAccountOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByAccountPager
// method.
type PrivateEndpointConnectionsClientListByAccountOptions struct {
	// The skip token.
	SkipToken *string
}

// PrivateLinkResourcesClientGetByGroupIDOptions contains the optional parameters for the PrivateLinkResourcesClient.GetByGroupID
// method.
type PrivateLinkResourcesClientGetByGroupIDOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByAccountOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByAccountPager
// method.
type PrivateLinkResourcesClientListByAccountOptions struct {
	// placeholder for future optional parameters
}
