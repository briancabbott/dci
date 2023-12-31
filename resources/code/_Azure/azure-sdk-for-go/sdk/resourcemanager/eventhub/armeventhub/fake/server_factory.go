//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armeventhub.ClientFactory type.
type ServerFactory struct {
	ApplicationGroupServer                       ApplicationGroupServer
	ClustersServer                               ClustersServer
	ConfigurationServer                          ConfigurationServer
	ConsumerGroupsServer                         ConsumerGroupsServer
	DisasterRecoveryConfigsServer                DisasterRecoveryConfigsServer
	EventHubsServer                              EventHubsServer
	NamespacesServer                             NamespacesServer
	NetworkSecurityPerimeterConfigurationServer  NetworkSecurityPerimeterConfigurationServer
	NetworkSecurityPerimeterConfigurationsServer NetworkSecurityPerimeterConfigurationsServer
	OperationsServer                             OperationsServer
	PrivateEndpointConnectionsServer             PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer                   PrivateLinkResourcesServer
	SchemaRegistryServer                         SchemaRegistryServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armeventhub.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armeventhub.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                            *ServerFactory
	trMu                                           sync.Mutex
	trApplicationGroupServer                       *ApplicationGroupServerTransport
	trClustersServer                               *ClustersServerTransport
	trConfigurationServer                          *ConfigurationServerTransport
	trConsumerGroupsServer                         *ConsumerGroupsServerTransport
	trDisasterRecoveryConfigsServer                *DisasterRecoveryConfigsServerTransport
	trEventHubsServer                              *EventHubsServerTransport
	trNamespacesServer                             *NamespacesServerTransport
	trNetworkSecurityPerimeterConfigurationServer  *NetworkSecurityPerimeterConfigurationServerTransport
	trNetworkSecurityPerimeterConfigurationsServer *NetworkSecurityPerimeterConfigurationsServerTransport
	trOperationsServer                             *OperationsServerTransport
	trPrivateEndpointConnectionsServer             *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer                   *PrivateLinkResourcesServerTransport
	trSchemaRegistryServer                         *SchemaRegistryServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "ApplicationGroupClient":
		initServer(s, &s.trApplicationGroupServer, func() *ApplicationGroupServerTransport {
			return NewApplicationGroupServerTransport(&s.srv.ApplicationGroupServer)
		})
		resp, err = s.trApplicationGroupServer.Do(req)
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "ConfigurationClient":
		initServer(s, &s.trConfigurationServer, func() *ConfigurationServerTransport {
			return NewConfigurationServerTransport(&s.srv.ConfigurationServer)
		})
		resp, err = s.trConfigurationServer.Do(req)
	case "ConsumerGroupsClient":
		initServer(s, &s.trConsumerGroupsServer, func() *ConsumerGroupsServerTransport {
			return NewConsumerGroupsServerTransport(&s.srv.ConsumerGroupsServer)
		})
		resp, err = s.trConsumerGroupsServer.Do(req)
	case "DisasterRecoveryConfigsClient":
		initServer(s, &s.trDisasterRecoveryConfigsServer, func() *DisasterRecoveryConfigsServerTransport {
			return NewDisasterRecoveryConfigsServerTransport(&s.srv.DisasterRecoveryConfigsServer)
		})
		resp, err = s.trDisasterRecoveryConfigsServer.Do(req)
	case "EventHubsClient":
		initServer(s, &s.trEventHubsServer, func() *EventHubsServerTransport { return NewEventHubsServerTransport(&s.srv.EventHubsServer) })
		resp, err = s.trEventHubsServer.Do(req)
	case "NamespacesClient":
		initServer(s, &s.trNamespacesServer, func() *NamespacesServerTransport { return NewNamespacesServerTransport(&s.srv.NamespacesServer) })
		resp, err = s.trNamespacesServer.Do(req)
	case "NetworkSecurityPerimeterConfigurationClient":
		initServer(s, &s.trNetworkSecurityPerimeterConfigurationServer, func() *NetworkSecurityPerimeterConfigurationServerTransport {
			return NewNetworkSecurityPerimeterConfigurationServerTransport(&s.srv.NetworkSecurityPerimeterConfigurationServer)
		})
		resp, err = s.trNetworkSecurityPerimeterConfigurationServer.Do(req)
	case "NetworkSecurityPerimeterConfigurationsClient":
		initServer(s, &s.trNetworkSecurityPerimeterConfigurationsServer, func() *NetworkSecurityPerimeterConfigurationsServerTransport {
			return NewNetworkSecurityPerimeterConfigurationsServerTransport(&s.srv.NetworkSecurityPerimeterConfigurationsServer)
		})
		resp, err = s.trNetworkSecurityPerimeterConfigurationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "SchemaRegistryClient":
		initServer(s, &s.trSchemaRegistryServer, func() *SchemaRegistryServerTransport {
			return NewSchemaRegistryServerTransport(&s.srv.SchemaRegistryServer)
		})
		resp, err = s.trSchemaRegistryServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
