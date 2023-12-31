//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewAccessClient creates a new instance of AccessClient.
func (c *ClientFactory) NewAccessClient() *AccessClient {
	subClient, _ := NewAccessClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMarketplaceAgreementsClient creates a new instance of MarketplaceAgreementsClient.
func (c *ClientFactory) NewMarketplaceAgreementsClient() *MarketplaceAgreementsClient {
	subClient, _ := NewMarketplaceAgreementsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOrganizationClient creates a new instance of OrganizationClient.
func (c *ClientFactory) NewOrganizationClient() *OrganizationClient {
	subClient, _ := NewOrganizationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOrganizationOperationsClient creates a new instance of OrganizationOperationsClient.
func (c *ClientFactory) NewOrganizationOperationsClient() *OrganizationOperationsClient {
	subClient, _ := NewOrganizationOperationsClient(c.credential, c.options)
	return subClient
}

// NewValidationsClient creates a new instance of ValidationsClient.
func (c *ClientFactory) NewValidationsClient() *ValidationsClient {
	subClient, _ := NewValidationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
