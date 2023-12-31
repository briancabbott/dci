//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

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

// NewAttachedDataNetworksClient creates a new instance of AttachedDataNetworksClient.
func (c *ClientFactory) NewAttachedDataNetworksClient() *AttachedDataNetworksClient {
	subClient, _ := NewAttachedDataNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDataNetworksClient creates a new instance of DataNetworksClient.
func (c *ClientFactory) NewDataNetworksClient() *DataNetworksClient {
	subClient, _ := NewDataNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDiagnosticsPackagesClient creates a new instance of DiagnosticsPackagesClient.
func (c *ClientFactory) NewDiagnosticsPackagesClient() *DiagnosticsPackagesClient {
	subClient, _ := NewDiagnosticsPackagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMobileNetworksClient creates a new instance of MobileNetworksClient.
func (c *ClientFactory) NewMobileNetworksClient() *MobileNetworksClient {
	subClient, _ := NewMobileNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient.
func (c *ClientFactory) NewPacketCapturesClient() *PacketCapturesClient {
	subClient, _ := NewPacketCapturesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPacketCoreControlPlaneVersionsClient creates a new instance of PacketCoreControlPlaneVersionsClient.
func (c *ClientFactory) NewPacketCoreControlPlaneVersionsClient() *PacketCoreControlPlaneVersionsClient {
	subClient, _ := NewPacketCoreControlPlaneVersionsClient(c.credential, c.options)
	return subClient
}

// NewPacketCoreControlPlanesClient creates a new instance of PacketCoreControlPlanesClient.
func (c *ClientFactory) NewPacketCoreControlPlanesClient() *PacketCoreControlPlanesClient {
	subClient, _ := NewPacketCoreControlPlanesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPacketCoreDataPlanesClient creates a new instance of PacketCoreDataPlanesClient.
func (c *ClientFactory) NewPacketCoreDataPlanesClient() *PacketCoreDataPlanesClient {
	subClient, _ := NewPacketCoreDataPlanesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSimGroupsClient creates a new instance of SimGroupsClient.
func (c *ClientFactory) NewSimGroupsClient() *SimGroupsClient {
	subClient, _ := NewSimGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSimPoliciesClient creates a new instance of SimPoliciesClient.
func (c *ClientFactory) NewSimPoliciesClient() *SimPoliciesClient {
	subClient, _ := NewSimPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSimsClient creates a new instance of SimsClient.
func (c *ClientFactory) NewSimsClient() *SimsClient {
	subClient, _ := NewSimsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSitesClient creates a new instance of SitesClient.
func (c *ClientFactory) NewSitesClient() *SitesClient {
	subClient, _ := NewSitesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSlicesClient creates a new instance of SlicesClient.
func (c *ClientFactory) NewSlicesClient() *SlicesClient {
	subClient, _ := NewSlicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
