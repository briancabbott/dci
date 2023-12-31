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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RegionServer is a fake server for instances of the armapimanagement.RegionClient type.
type RegionServer struct {
	// NewListByServicePager is the fake for method RegionClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.RegionClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.RegionClientListByServiceResponse])
}

// NewRegionServerTransport creates a new instance of RegionServerTransport with the provided implementation.
// The returned RegionServerTransport instance is connected to an instance of armapimanagement.RegionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegionServerTransport(srv *RegionServer) *RegionServerTransport {
	return &RegionServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.RegionClientListByServiceResponse]](),
	}
}

// RegionServerTransport connects instances of armapimanagement.RegionClient to instances of RegionServer.
// Don't use this type directly, use NewRegionServerTransport instead.
type RegionServerTransport struct {
	srv                   *RegionServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.RegionClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for RegionServerTransport.
func (r *RegionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegionClient.NewListByServicePager":
		resp, err = r.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegionServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := r.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, nil)
		newListByServicePager = &resp
		r.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.RegionClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		r.newListByServicePager.remove(req)
	}
	return resp, nil
}
