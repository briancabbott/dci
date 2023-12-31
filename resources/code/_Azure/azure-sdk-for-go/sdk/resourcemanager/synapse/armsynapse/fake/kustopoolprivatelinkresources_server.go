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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// KustoPoolPrivateLinkResourcesServer is a fake server for instances of the armsynapse.KustoPoolPrivateLinkResourcesClient type.
type KustoPoolPrivateLinkResourcesServer struct {
	// NewListPager is the fake for method KustoPoolPrivateLinkResourcesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, kustoPoolName string, options *armsynapse.KustoPoolPrivateLinkResourcesClientListOptions) (resp azfake.PagerResponder[armsynapse.KustoPoolPrivateLinkResourcesClientListResponse])
}

// NewKustoPoolPrivateLinkResourcesServerTransport creates a new instance of KustoPoolPrivateLinkResourcesServerTransport with the provided implementation.
// The returned KustoPoolPrivateLinkResourcesServerTransport instance is connected to an instance of armsynapse.KustoPoolPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewKustoPoolPrivateLinkResourcesServerTransport(srv *KustoPoolPrivateLinkResourcesServer) *KustoPoolPrivateLinkResourcesServerTransport {
	return &KustoPoolPrivateLinkResourcesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsynapse.KustoPoolPrivateLinkResourcesClientListResponse]](),
	}
}

// KustoPoolPrivateLinkResourcesServerTransport connects instances of armsynapse.KustoPoolPrivateLinkResourcesClient to instances of KustoPoolPrivateLinkResourcesServer.
// Don't use this type directly, use NewKustoPoolPrivateLinkResourcesServerTransport instead.
type KustoPoolPrivateLinkResourcesServerTransport struct {
	srv          *KustoPoolPrivateLinkResourcesServer
	newListPager *tracker[azfake.PagerResponder[armsynapse.KustoPoolPrivateLinkResourcesClientListResponse]]
}

// Do implements the policy.Transporter interface for KustoPoolPrivateLinkResourcesServerTransport.
func (k *KustoPoolPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "KustoPoolPrivateLinkResourcesClient.NewListPager":
		resp, err = k.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (k *KustoPoolPrivateLinkResourcesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if k.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := k.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/kustoPools/(?P<kustoPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		kustoPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("kustoPoolName")])
		if err != nil {
			return nil, err
		}
		resp := k.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, kustoPoolNameParam, nil)
		newListPager = &resp
		k.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		k.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		k.newListPager.remove(req)
	}
	return resp, nil
}
