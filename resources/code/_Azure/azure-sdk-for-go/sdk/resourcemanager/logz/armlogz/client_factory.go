//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

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
//   - subscriptionID - The ID of the target subscription.
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

// NewMonitorClient creates a new instance of MonitorClient.
func (c *ClientFactory) NewMonitorClient() *MonitorClient {
	subClient, _ := NewMonitorClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMonitorsClient creates a new instance of MonitorsClient.
func (c *ClientFactory) NewMonitorsClient() *MonitorsClient {
	subClient, _ := NewMonitorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewSingleSignOnClient creates a new instance of SingleSignOnClient.
func (c *ClientFactory) NewSingleSignOnClient() *SingleSignOnClient {
	subClient, _ := NewSingleSignOnClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSubAccountClient creates a new instance of SubAccountClient.
func (c *ClientFactory) NewSubAccountClient() *SubAccountClient {
	subClient, _ := NewSubAccountClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSubAccountTagRulesClient creates a new instance of SubAccountTagRulesClient.
func (c *ClientFactory) NewSubAccountTagRulesClient() *SubAccountTagRulesClient {
	subClient, _ := NewSubAccountTagRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTagRulesClient creates a new instance of TagRulesClient.
func (c *ClientFactory) NewTagRulesClient() *TagRulesClient {
	subClient, _ := NewTagRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
