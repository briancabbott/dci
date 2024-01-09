//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

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

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCustomCertificatesClient creates a new instance of CustomCertificatesClient.
func (c *ClientFactory) NewCustomCertificatesClient() *CustomCertificatesClient {
	subClient, _ := NewCustomCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCustomDomainsClient creates a new instance of CustomDomainsClient.
func (c *ClientFactory) NewCustomDomainsClient() *CustomDomainsClient {
	subClient, _ := NewCustomDomainsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewHubsClient creates a new instance of HubsClient.
func (c *ClientFactory) NewHubsClient() *HubsClient {
	subClient, _ := NewHubsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewReplicasClient creates a new instance of ReplicasClient.
func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	subClient, _ := NewReplicasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSharedPrivateLinkResourcesClient creates a new instance of SharedPrivateLinkResourcesClient.
func (c *ClientFactory) NewSharedPrivateLinkResourcesClient() *SharedPrivateLinkResourcesClient {
	subClient, _ := NewSharedPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
