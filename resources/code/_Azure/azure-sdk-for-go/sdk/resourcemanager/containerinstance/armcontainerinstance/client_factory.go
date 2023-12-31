//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

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
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
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

// NewContainerGroupsClient creates a new instance of ContainerGroupsClient.
func (c *ClientFactory) NewContainerGroupsClient() *ContainerGroupsClient {
	subClient, _ := NewContainerGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewContainersClient creates a new instance of ContainersClient.
func (c *ClientFactory) NewContainersClient() *ContainersClient {
	subClient, _ := NewContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLocationClient creates a new instance of LocationClient.
func (c *ClientFactory) NewLocationClient() *LocationClient {
	subClient, _ := NewLocationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewSubnetServiceAssociationLinkClient creates a new instance of SubnetServiceAssociationLinkClient.
func (c *ClientFactory) NewSubnetServiceAssociationLinkClient() *SubnetServiceAssociationLinkClient {
	subClient, _ := NewSubnetServiceAssociationLinkClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
