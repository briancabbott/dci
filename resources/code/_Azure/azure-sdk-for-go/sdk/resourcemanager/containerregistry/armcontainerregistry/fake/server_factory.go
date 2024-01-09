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

// ServerFactory is a fake server for instances of the armcontainerregistry.ClientFactory type.
type ServerFactory struct {
	AgentPoolsServer                 AgentPoolsServer
	ArchiveVersionsServer            ArchiveVersionsServer
	ArchivesServer                   ArchivesServer
	CacheRulesServer                 CacheRulesServer
	ConnectedRegistriesServer        ConnectedRegistriesServer
	CredentialSetsServer             CredentialSetsServer
	ExportPipelinesServer            ExportPipelinesServer
	ImportPipelinesServer            ImportPipelinesServer
	OperationsServer                 OperationsServer
	PipelineRunsServer               PipelineRunsServer
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer
	RegistriesServer                 RegistriesServer
	ReplicationsServer               ReplicationsServer
	RunsServer                       RunsServer
	ScopeMapsServer                  ScopeMapsServer
	TaskRunsServer                   TaskRunsServer
	TasksServer                      TasksServer
	TokensServer                     TokensServer
	WebhooksServer                   WebhooksServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armcontainerregistry.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armcontainerregistry.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trAgentPoolsServer                 *AgentPoolsServerTransport
	trArchiveVersionsServer            *ArchiveVersionsServerTransport
	trArchivesServer                   *ArchivesServerTransport
	trCacheRulesServer                 *CacheRulesServerTransport
	trConnectedRegistriesServer        *ConnectedRegistriesServerTransport
	trCredentialSetsServer             *CredentialSetsServerTransport
	trExportPipelinesServer            *ExportPipelinesServerTransport
	trImportPipelinesServer            *ImportPipelinesServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPipelineRunsServer               *PipelineRunsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trRegistriesServer                 *RegistriesServerTransport
	trReplicationsServer               *ReplicationsServerTransport
	trRunsServer                       *RunsServerTransport
	trScopeMapsServer                  *ScopeMapsServerTransport
	trTaskRunsServer                   *TaskRunsServerTransport
	trTasksServer                      *TasksServerTransport
	trTokensServer                     *TokensServerTransport
	trWebhooksServer                   *WebhooksServerTransport
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
	case "AgentPoolsClient":
		initServer(s, &s.trAgentPoolsServer, func() *AgentPoolsServerTransport { return NewAgentPoolsServerTransport(&s.srv.AgentPoolsServer) })
		resp, err = s.trAgentPoolsServer.Do(req)
	case "ArchiveVersionsClient":
		initServer(s, &s.trArchiveVersionsServer, func() *ArchiveVersionsServerTransport {
			return NewArchiveVersionsServerTransport(&s.srv.ArchiveVersionsServer)
		})
		resp, err = s.trArchiveVersionsServer.Do(req)
	case "ArchivesClient":
		initServer(s, &s.trArchivesServer, func() *ArchivesServerTransport { return NewArchivesServerTransport(&s.srv.ArchivesServer) })
		resp, err = s.trArchivesServer.Do(req)
	case "CacheRulesClient":
		initServer(s, &s.trCacheRulesServer, func() *CacheRulesServerTransport { return NewCacheRulesServerTransport(&s.srv.CacheRulesServer) })
		resp, err = s.trCacheRulesServer.Do(req)
	case "ConnectedRegistriesClient":
		initServer(s, &s.trConnectedRegistriesServer, func() *ConnectedRegistriesServerTransport {
			return NewConnectedRegistriesServerTransport(&s.srv.ConnectedRegistriesServer)
		})
		resp, err = s.trConnectedRegistriesServer.Do(req)
	case "CredentialSetsClient":
		initServer(s, &s.trCredentialSetsServer, func() *CredentialSetsServerTransport {
			return NewCredentialSetsServerTransport(&s.srv.CredentialSetsServer)
		})
		resp, err = s.trCredentialSetsServer.Do(req)
	case "ExportPipelinesClient":
		initServer(s, &s.trExportPipelinesServer, func() *ExportPipelinesServerTransport {
			return NewExportPipelinesServerTransport(&s.srv.ExportPipelinesServer)
		})
		resp, err = s.trExportPipelinesServer.Do(req)
	case "ImportPipelinesClient":
		initServer(s, &s.trImportPipelinesServer, func() *ImportPipelinesServerTransport {
			return NewImportPipelinesServerTransport(&s.srv.ImportPipelinesServer)
		})
		resp, err = s.trImportPipelinesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PipelineRunsClient":
		initServer(s, &s.trPipelineRunsServer, func() *PipelineRunsServerTransport { return NewPipelineRunsServerTransport(&s.srv.PipelineRunsServer) })
		resp, err = s.trPipelineRunsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "RegistriesClient":
		initServer(s, &s.trRegistriesServer, func() *RegistriesServerTransport { return NewRegistriesServerTransport(&s.srv.RegistriesServer) })
		resp, err = s.trRegistriesServer.Do(req)
	case "ReplicationsClient":
		initServer(s, &s.trReplicationsServer, func() *ReplicationsServerTransport { return NewReplicationsServerTransport(&s.srv.ReplicationsServer) })
		resp, err = s.trReplicationsServer.Do(req)
	case "RunsClient":
		initServer(s, &s.trRunsServer, func() *RunsServerTransport { return NewRunsServerTransport(&s.srv.RunsServer) })
		resp, err = s.trRunsServer.Do(req)
	case "ScopeMapsClient":
		initServer(s, &s.trScopeMapsServer, func() *ScopeMapsServerTransport { return NewScopeMapsServerTransport(&s.srv.ScopeMapsServer) })
		resp, err = s.trScopeMapsServer.Do(req)
	case "TaskRunsClient":
		initServer(s, &s.trTaskRunsServer, func() *TaskRunsServerTransport { return NewTaskRunsServerTransport(&s.srv.TaskRunsServer) })
		resp, err = s.trTaskRunsServer.Do(req)
	case "TasksClient":
		initServer(s, &s.trTasksServer, func() *TasksServerTransport { return NewTasksServerTransport(&s.srv.TasksServer) })
		resp, err = s.trTasksServer.Do(req)
	case "TokensClient":
		initServer(s, &s.trTokensServer, func() *TokensServerTransport { return NewTokensServerTransport(&s.srv.TokensServer) })
		resp, err = s.trTokensServer.Do(req)
	case "WebhooksClient":
		initServer(s, &s.trWebhooksServer, func() *WebhooksServerTransport { return NewWebhooksServerTransport(&s.srv.WebhooksServer) })
		resp, err = s.trWebhooksServer.Do(req)
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
