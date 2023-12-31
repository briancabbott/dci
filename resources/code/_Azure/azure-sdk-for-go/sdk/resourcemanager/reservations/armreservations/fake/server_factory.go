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

// ServerFactory is a fake server for instances of the armreservations.ClientFactory type.
type ServerFactory struct {
	AzureReservationAPIServer AzureReservationAPIServer
	CalculateExchangeServer   CalculateExchangeServer
	CalculateRefundServer     CalculateRefundServer
	ExchangeServer            ExchangeServer
	OperationServer           OperationServer
	QuotaServer               QuotaServer
	QuotaRequestStatusServer  QuotaRequestStatusServer
	ReservationServer         ReservationServer
	ReservationOrderServer    ReservationOrderServer
	ReturnServer              ReturnServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armreservations.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armreservations.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                         *ServerFactory
	trMu                        sync.Mutex
	trAzureReservationAPIServer *AzureReservationAPIServerTransport
	trCalculateExchangeServer   *CalculateExchangeServerTransport
	trCalculateRefundServer     *CalculateRefundServerTransport
	trExchangeServer            *ExchangeServerTransport
	trOperationServer           *OperationServerTransport
	trQuotaServer               *QuotaServerTransport
	trQuotaRequestStatusServer  *QuotaRequestStatusServerTransport
	trReservationServer         *ReservationServerTransport
	trReservationOrderServer    *ReservationOrderServerTransport
	trReturnServer              *ReturnServerTransport
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
	case "AzureReservationAPIClient":
		initServer(s, &s.trAzureReservationAPIServer, func() *AzureReservationAPIServerTransport {
			return NewAzureReservationAPIServerTransport(&s.srv.AzureReservationAPIServer)
		})
		resp, err = s.trAzureReservationAPIServer.Do(req)
	case "CalculateExchangeClient":
		initServer(s, &s.trCalculateExchangeServer, func() *CalculateExchangeServerTransport {
			return NewCalculateExchangeServerTransport(&s.srv.CalculateExchangeServer)
		})
		resp, err = s.trCalculateExchangeServer.Do(req)
	case "CalculateRefundClient":
		initServer(s, &s.trCalculateRefundServer, func() *CalculateRefundServerTransport {
			return NewCalculateRefundServerTransport(&s.srv.CalculateRefundServer)
		})
		resp, err = s.trCalculateRefundServer.Do(req)
	case "ExchangeClient":
		initServer(s, &s.trExchangeServer, func() *ExchangeServerTransport { return NewExchangeServerTransport(&s.srv.ExchangeServer) })
		resp, err = s.trExchangeServer.Do(req)
	case "OperationClient":
		initServer(s, &s.trOperationServer, func() *OperationServerTransport { return NewOperationServerTransport(&s.srv.OperationServer) })
		resp, err = s.trOperationServer.Do(req)
	case "QuotaClient":
		initServer(s, &s.trQuotaServer, func() *QuotaServerTransport { return NewQuotaServerTransport(&s.srv.QuotaServer) })
		resp, err = s.trQuotaServer.Do(req)
	case "QuotaRequestStatusClient":
		initServer(s, &s.trQuotaRequestStatusServer, func() *QuotaRequestStatusServerTransport {
			return NewQuotaRequestStatusServerTransport(&s.srv.QuotaRequestStatusServer)
		})
		resp, err = s.trQuotaRequestStatusServer.Do(req)
	case "ReservationClient":
		initServer(s, &s.trReservationServer, func() *ReservationServerTransport { return NewReservationServerTransport(&s.srv.ReservationServer) })
		resp, err = s.trReservationServer.Do(req)
	case "ReservationOrderClient":
		initServer(s, &s.trReservationOrderServer, func() *ReservationOrderServerTransport {
			return NewReservationOrderServerTransport(&s.srv.ReservationOrderServer)
		})
		resp, err = s.trReservationOrderServer.Do(req)
	case "ReturnClient":
		initServer(s, &s.trReturnServer, func() *ReturnServerTransport { return NewReturnServerTransport(&s.srv.ReturnServer) })
		resp, err = s.trReturnServer.Do(req)
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
