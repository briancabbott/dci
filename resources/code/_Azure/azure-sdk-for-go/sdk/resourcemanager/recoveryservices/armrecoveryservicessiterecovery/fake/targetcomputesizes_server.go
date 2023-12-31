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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// TargetComputeSizesServer is a fake server for instances of the armrecoveryservicessiterecovery.TargetComputeSizesClient type.
type TargetComputeSizesServer struct {
	// NewListByReplicationProtectedItemsPager is the fake for method TargetComputeSizesClient.NewListByReplicationProtectedItemsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReplicationProtectedItemsPager func(resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, options *armrecoveryservicessiterecovery.TargetComputeSizesClientListByReplicationProtectedItemsOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.TargetComputeSizesClientListByReplicationProtectedItemsResponse])
}

// NewTargetComputeSizesServerTransport creates a new instance of TargetComputeSizesServerTransport with the provided implementation.
// The returned TargetComputeSizesServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.TargetComputeSizesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTargetComputeSizesServerTransport(srv *TargetComputeSizesServer) *TargetComputeSizesServerTransport {
	return &TargetComputeSizesServerTransport{
		srv:                                     srv,
		newListByReplicationProtectedItemsPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.TargetComputeSizesClientListByReplicationProtectedItemsResponse]](),
	}
}

// TargetComputeSizesServerTransport connects instances of armrecoveryservicessiterecovery.TargetComputeSizesClient to instances of TargetComputeSizesServer.
// Don't use this type directly, use NewTargetComputeSizesServerTransport instead.
type TargetComputeSizesServerTransport struct {
	srv                                     *TargetComputeSizesServer
	newListByReplicationProtectedItemsPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.TargetComputeSizesClientListByReplicationProtectedItemsResponse]]
}

// Do implements the policy.Transporter interface for TargetComputeSizesServerTransport.
func (t *TargetComputeSizesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TargetComputeSizesClient.NewListByReplicationProtectedItemsPager":
		resp, err = t.dispatchNewListByReplicationProtectedItemsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TargetComputeSizesServerTransport) dispatchNewListByReplicationProtectedItemsPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByReplicationProtectedItemsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReplicationProtectedItemsPager not implemented")}
	}
	newListByReplicationProtectedItemsPager := t.newListByReplicationProtectedItemsPager.get(req)
	if newListByReplicationProtectedItemsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionContainers/(?P<protectionContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectedItems/(?P<replicatedProtectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetComputeSizes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		protectionContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectionContainerName")])
		if err != nil {
			return nil, err
		}
		replicatedProtectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicatedProtectedItemName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByReplicationProtectedItemsPager(resourceNameParam, resourceGroupNameParam, fabricNameParam, protectionContainerNameParam, replicatedProtectedItemNameParam, nil)
		newListByReplicationProtectedItemsPager = &resp
		t.newListByReplicationProtectedItemsPager.add(req, newListByReplicationProtectedItemsPager)
		server.PagerResponderInjectNextLinks(newListByReplicationProtectedItemsPager, req, func(page *armrecoveryservicessiterecovery.TargetComputeSizesClientListByReplicationProtectedItemsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReplicationProtectedItemsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByReplicationProtectedItemsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReplicationProtectedItemsPager) {
		t.newListByReplicationProtectedItemsPager.remove(req)
	}
	return resp, nil
}
