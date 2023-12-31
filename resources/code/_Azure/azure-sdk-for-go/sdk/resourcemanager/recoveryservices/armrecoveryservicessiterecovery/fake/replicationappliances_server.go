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

// ReplicationAppliancesServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationAppliancesClient type.
type ReplicationAppliancesServer struct {
	// NewListPager is the fake for method ReplicationAppliancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceName string, resourceGroupName string, options *armrecoveryservicessiterecovery.ReplicationAppliancesClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationAppliancesClientListResponse])
}

// NewReplicationAppliancesServerTransport creates a new instance of ReplicationAppliancesServerTransport with the provided implementation.
// The returned ReplicationAppliancesServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationAppliancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationAppliancesServerTransport(srv *ReplicationAppliancesServer) *ReplicationAppliancesServerTransport {
	return &ReplicationAppliancesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationAppliancesClientListResponse]](),
	}
}

// ReplicationAppliancesServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationAppliancesClient to instances of ReplicationAppliancesServer.
// Don't use this type directly, use NewReplicationAppliancesServerTransport instead.
type ReplicationAppliancesServerTransport struct {
	srv          *ReplicationAppliancesServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationAppliancesClientListResponse]]
}

// Do implements the policy.Transporter interface for ReplicationAppliancesServerTransport.
func (r *ReplicationAppliancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationAppliancesClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationAppliancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationAppliances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armrecoveryservicessiterecovery.ReplicationAppliancesClientListOptions
		if filterParam != nil {
			options = &armrecoveryservicessiterecovery.ReplicationAppliancesClientListOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListPager(resourceNameParam, resourceGroupNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationAppliancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}
