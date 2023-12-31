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

// ServerFactory is a fake server for instances of the armservicebus.ClientFactory type.
type ServerFactory struct {
	DisasterRecoveryConfigsServer    DisasterRecoveryConfigsServer
	MigrationConfigsServer           MigrationConfigsServer
	NamespacesServer                 NamespacesServer
	OperationsServer                 OperationsServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer       PrivateLinkResourcesServer
	QueuesServer                     QueuesServer
	RulesServer                      RulesServer
	SubscriptionsServer              SubscriptionsServer
	TopicsServer                     TopicsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armservicebus.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armservicebus.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trDisasterRecoveryConfigsServer    *DisasterRecoveryConfigsServerTransport
	trMigrationConfigsServer           *MigrationConfigsServerTransport
	trNamespacesServer                 *NamespacesServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trQueuesServer                     *QueuesServerTransport
	trRulesServer                      *RulesServerTransport
	trSubscriptionsServer              *SubscriptionsServerTransport
	trTopicsServer                     *TopicsServerTransport
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
	case "DisasterRecoveryConfigsClient":
		initServer(s, &s.trDisasterRecoveryConfigsServer, func() *DisasterRecoveryConfigsServerTransport {
			return NewDisasterRecoveryConfigsServerTransport(&s.srv.DisasterRecoveryConfigsServer)
		})
		resp, err = s.trDisasterRecoveryConfigsServer.Do(req)
	case "MigrationConfigsClient":
		initServer(s, &s.trMigrationConfigsServer, func() *MigrationConfigsServerTransport {
			return NewMigrationConfigsServerTransport(&s.srv.MigrationConfigsServer)
		})
		resp, err = s.trMigrationConfigsServer.Do(req)
	case "NamespacesClient":
		initServer(s, &s.trNamespacesServer, func() *NamespacesServerTransport { return NewNamespacesServerTransport(&s.srv.NamespacesServer) })
		resp, err = s.trNamespacesServer.Do(req)
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
	case "QueuesClient":
		initServer(s, &s.trQueuesServer, func() *QueuesServerTransport { return NewQueuesServerTransport(&s.srv.QueuesServer) })
		resp, err = s.trQueuesServer.Do(req)
	case "RulesClient":
		initServer(s, &s.trRulesServer, func() *RulesServerTransport { return NewRulesServerTransport(&s.srv.RulesServer) })
		resp, err = s.trRulesServer.Do(req)
	case "SubscriptionsClient":
		initServer(s, &s.trSubscriptionsServer, func() *SubscriptionsServerTransport {
			return NewSubscriptionsServerTransport(&s.srv.SubscriptionsServer)
		})
		resp, err = s.trSubscriptionsServer.Do(req)
	case "TopicsClient":
		initServer(s, &s.trTopicsServer, func() *TopicsServerTransport { return NewTopicsServerTransport(&s.srv.TopicsServer) })
		resp, err = s.trTopicsServer.Do(req)
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
