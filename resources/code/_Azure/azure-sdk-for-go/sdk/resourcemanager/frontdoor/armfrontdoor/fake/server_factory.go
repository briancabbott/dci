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

// ServerFactory is a fake server for instances of the armfrontdoor.ClientFactory type.
type ServerFactory struct {
	EndpointsServer                        EndpointsServer
	ExperimentsServer                      ExperimentsServer
	FrontDoorsServer                       FrontDoorsServer
	FrontendEndpointsServer                FrontendEndpointsServer
	ManagedRuleSetsServer                  ManagedRuleSetsServer
	NameAvailabilityServer                 NameAvailabilityServer
	NameAvailabilityWithSubscriptionServer NameAvailabilityWithSubscriptionServer
	NetworkExperimentProfilesServer        NetworkExperimentProfilesServer
	PoliciesServer                         PoliciesServer
	PreconfiguredEndpointsServer           PreconfiguredEndpointsServer
	ReportsServer                          ReportsServer
	RulesEnginesServer                     RulesEnginesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armfrontdoor.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armfrontdoor.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                      *ServerFactory
	trMu                                     sync.Mutex
	trEndpointsServer                        *EndpointsServerTransport
	trExperimentsServer                      *ExperimentsServerTransport
	trFrontDoorsServer                       *FrontDoorsServerTransport
	trFrontendEndpointsServer                *FrontendEndpointsServerTransport
	trManagedRuleSetsServer                  *ManagedRuleSetsServerTransport
	trNameAvailabilityServer                 *NameAvailabilityServerTransport
	trNameAvailabilityWithSubscriptionServer *NameAvailabilityWithSubscriptionServerTransport
	trNetworkExperimentProfilesServer        *NetworkExperimentProfilesServerTransport
	trPoliciesServer                         *PoliciesServerTransport
	trPreconfiguredEndpointsServer           *PreconfiguredEndpointsServerTransport
	trReportsServer                          *ReportsServerTransport
	trRulesEnginesServer                     *RulesEnginesServerTransport
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
	case "EndpointsClient":
		initServer(s, &s.trEndpointsServer, func() *EndpointsServerTransport { return NewEndpointsServerTransport(&s.srv.EndpointsServer) })
		resp, err = s.trEndpointsServer.Do(req)
	case "ExperimentsClient":
		initServer(s, &s.trExperimentsServer, func() *ExperimentsServerTransport { return NewExperimentsServerTransport(&s.srv.ExperimentsServer) })
		resp, err = s.trExperimentsServer.Do(req)
	case "FrontDoorsClient":
		initServer(s, &s.trFrontDoorsServer, func() *FrontDoorsServerTransport { return NewFrontDoorsServerTransport(&s.srv.FrontDoorsServer) })
		resp, err = s.trFrontDoorsServer.Do(req)
	case "FrontendEndpointsClient":
		initServer(s, &s.trFrontendEndpointsServer, func() *FrontendEndpointsServerTransport {
			return NewFrontendEndpointsServerTransport(&s.srv.FrontendEndpointsServer)
		})
		resp, err = s.trFrontendEndpointsServer.Do(req)
	case "ManagedRuleSetsClient":
		initServer(s, &s.trManagedRuleSetsServer, func() *ManagedRuleSetsServerTransport {
			return NewManagedRuleSetsServerTransport(&s.srv.ManagedRuleSetsServer)
		})
		resp, err = s.trManagedRuleSetsServer.Do(req)
	case "NameAvailabilityClient":
		initServer(s, &s.trNameAvailabilityServer, func() *NameAvailabilityServerTransport {
			return NewNameAvailabilityServerTransport(&s.srv.NameAvailabilityServer)
		})
		resp, err = s.trNameAvailabilityServer.Do(req)
	case "NameAvailabilityWithSubscriptionClient":
		initServer(s, &s.trNameAvailabilityWithSubscriptionServer, func() *NameAvailabilityWithSubscriptionServerTransport {
			return NewNameAvailabilityWithSubscriptionServerTransport(&s.srv.NameAvailabilityWithSubscriptionServer)
		})
		resp, err = s.trNameAvailabilityWithSubscriptionServer.Do(req)
	case "NetworkExperimentProfilesClient":
		initServer(s, &s.trNetworkExperimentProfilesServer, func() *NetworkExperimentProfilesServerTransport {
			return NewNetworkExperimentProfilesServerTransport(&s.srv.NetworkExperimentProfilesServer)
		})
		resp, err = s.trNetworkExperimentProfilesServer.Do(req)
	case "PoliciesClient":
		initServer(s, &s.trPoliciesServer, func() *PoliciesServerTransport { return NewPoliciesServerTransport(&s.srv.PoliciesServer) })
		resp, err = s.trPoliciesServer.Do(req)
	case "PreconfiguredEndpointsClient":
		initServer(s, &s.trPreconfiguredEndpointsServer, func() *PreconfiguredEndpointsServerTransport {
			return NewPreconfiguredEndpointsServerTransport(&s.srv.PreconfiguredEndpointsServer)
		})
		resp, err = s.trPreconfiguredEndpointsServer.Do(req)
	case "ReportsClient":
		initServer(s, &s.trReportsServer, func() *ReportsServerTransport { return NewReportsServerTransport(&s.srv.ReportsServer) })
		resp, err = s.trReportsServer.Do(req)
	case "RulesEnginesClient":
		initServer(s, &s.trRulesEnginesServer, func() *RulesEnginesServerTransport { return NewRulesEnginesServerTransport(&s.srv.RulesEnginesServer) })
		resp, err = s.trRulesEnginesServer.Do(req)
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
