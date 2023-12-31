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
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// SynapseLinkWorkspacesServer is a fake server for instances of the armsql.SynapseLinkWorkspacesClient type.
type SynapseLinkWorkspacesServer struct {
	// NewListByDatabasePager is the fake for method SynapseLinkWorkspacesClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.SynapseLinkWorkspacesClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.SynapseLinkWorkspacesClientListByDatabaseResponse])
}

// NewSynapseLinkWorkspacesServerTransport creates a new instance of SynapseLinkWorkspacesServerTransport with the provided implementation.
// The returned SynapseLinkWorkspacesServerTransport instance is connected to an instance of armsql.SynapseLinkWorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSynapseLinkWorkspacesServerTransport(srv *SynapseLinkWorkspacesServer) *SynapseLinkWorkspacesServerTransport {
	return &SynapseLinkWorkspacesServerTransport{
		srv:                    srv,
		newListByDatabasePager: newTracker[azfake.PagerResponder[armsql.SynapseLinkWorkspacesClientListByDatabaseResponse]](),
	}
}

// SynapseLinkWorkspacesServerTransport connects instances of armsql.SynapseLinkWorkspacesClient to instances of SynapseLinkWorkspacesServer.
// Don't use this type directly, use NewSynapseLinkWorkspacesServerTransport instead.
type SynapseLinkWorkspacesServerTransport struct {
	srv                    *SynapseLinkWorkspacesServer
	newListByDatabasePager *tracker[azfake.PagerResponder[armsql.SynapseLinkWorkspacesClientListByDatabaseResponse]]
}

// Do implements the policy.Transporter interface for SynapseLinkWorkspacesServerTransport.
func (s *SynapseLinkWorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SynapseLinkWorkspacesClient.NewListByDatabasePager":
		resp, err = s.dispatchNewListByDatabasePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SynapseLinkWorkspacesServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := s.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkWorkspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		s.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.SynapseLinkWorkspacesClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		s.newListByDatabasePager.remove(req)
	}
	return resp, nil
}
