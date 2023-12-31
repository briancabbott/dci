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

// ServerFactory is a fake server for instances of the armiotsecurity.ClientFactory type.
type ServerFactory struct {
	DefenderSettingsServer DefenderSettingsServer
	DeviceGroupsServer     DeviceGroupsServer
	DevicesServer          DevicesServer
	LocationsServer        LocationsServer
	OnPremiseSensorsServer OnPremiseSensorsServer
	OperationsServer       OperationsServer
	SensorsServer          SensorsServer
	SitesServer            SitesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armiotsecurity.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armiotsecurity.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                      *ServerFactory
	trMu                     sync.Mutex
	trDefenderSettingsServer *DefenderSettingsServerTransport
	trDeviceGroupsServer     *DeviceGroupsServerTransport
	trDevicesServer          *DevicesServerTransport
	trLocationsServer        *LocationsServerTransport
	trOnPremiseSensorsServer *OnPremiseSensorsServerTransport
	trOperationsServer       *OperationsServerTransport
	trSensorsServer          *SensorsServerTransport
	trSitesServer            *SitesServerTransport
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
	case "DefenderSettingsClient":
		initServer(s, &s.trDefenderSettingsServer, func() *DefenderSettingsServerTransport {
			return NewDefenderSettingsServerTransport(&s.srv.DefenderSettingsServer)
		})
		resp, err = s.trDefenderSettingsServer.Do(req)
	case "DeviceGroupsClient":
		initServer(s, &s.trDeviceGroupsServer, func() *DeviceGroupsServerTransport { return NewDeviceGroupsServerTransport(&s.srv.DeviceGroupsServer) })
		resp, err = s.trDeviceGroupsServer.Do(req)
	case "DevicesClient":
		initServer(s, &s.trDevicesServer, func() *DevicesServerTransport { return NewDevicesServerTransport(&s.srv.DevicesServer) })
		resp, err = s.trDevicesServer.Do(req)
	case "LocationsClient":
		initServer(s, &s.trLocationsServer, func() *LocationsServerTransport { return NewLocationsServerTransport(&s.srv.LocationsServer) })
		resp, err = s.trLocationsServer.Do(req)
	case "OnPremiseSensorsClient":
		initServer(s, &s.trOnPremiseSensorsServer, func() *OnPremiseSensorsServerTransport {
			return NewOnPremiseSensorsServerTransport(&s.srv.OnPremiseSensorsServer)
		})
		resp, err = s.trOnPremiseSensorsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "SensorsClient":
		initServer(s, &s.trSensorsServer, func() *SensorsServerTransport { return NewSensorsServerTransport(&s.srv.SensorsServer) })
		resp, err = s.trSensorsServer.Do(req)
	case "SitesClient":
		initServer(s, &s.trSitesServer, func() *SitesServerTransport { return NewSitesServerTransport(&s.srv.SitesServer) })
		resp, err = s.trSitesServer.Do(req)
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
