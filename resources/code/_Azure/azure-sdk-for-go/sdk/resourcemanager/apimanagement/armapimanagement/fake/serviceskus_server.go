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

// ServiceSKUsServer is a fake server for instances of the armapimanagement.ServiceSKUsClient type.
type ServiceSKUsServer struct {
	// NewListAvailableServiceSKUsPager is the fake for method ServiceSKUsClient.NewListAvailableServiceSKUsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAvailableServiceSKUsPager func(resourceGroupName string, serviceName string, options *armapimanagement.ServiceSKUsClientListAvailableServiceSKUsOptions) (resp azfake.PagerResponder[armapimanagement.ServiceSKUsClientListAvailableServiceSKUsResponse])
}

// NewServiceSKUsServerTransport creates a new instance of ServiceSKUsServerTransport with the provided implementation.
// The returned ServiceSKUsServerTransport instance is connected to an instance of armapimanagement.ServiceSKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceSKUsServerTransport(srv *ServiceSKUsServer) *ServiceSKUsServerTransport {
	return &ServiceSKUsServerTransport{
		srv:                              srv,
		newListAvailableServiceSKUsPager: newTracker[azfake.PagerResponder[armapimanagement.ServiceSKUsClientListAvailableServiceSKUsResponse]](),
	}
}

// ServiceSKUsServerTransport connects instances of armapimanagement.ServiceSKUsClient to instances of ServiceSKUsServer.
// Don't use this type directly, use NewServiceSKUsServerTransport instead.
type ServiceSKUsServerTransport struct {
	srv                              *ServiceSKUsServer
	newListAvailableServiceSKUsPager *tracker[azfake.PagerResponder[armapimanagement.ServiceSKUsClientListAvailableServiceSKUsResponse]]
}

// Do implements the policy.Transporter interface for ServiceSKUsServerTransport.
func (s *ServiceSKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceSKUsClient.NewListAvailableServiceSKUsPager":
		resp, err = s.dispatchNewListAvailableServiceSKUsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceSKUsServerTransport) dispatchNewListAvailableServiceSKUsPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListAvailableServiceSKUsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAvailableServiceSKUsPager not implemented")}
	}
	newListAvailableServiceSKUsPager := s.newListAvailableServiceSKUsPager.get(req)
	if newListAvailableServiceSKUsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
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
		resp := s.srv.NewListAvailableServiceSKUsPager(resourceGroupNameParam, serviceNameParam, nil)
		newListAvailableServiceSKUsPager = &resp
		s.newListAvailableServiceSKUsPager.add(req, newListAvailableServiceSKUsPager)
		server.PagerResponderInjectNextLinks(newListAvailableServiceSKUsPager, req, func(page *armapimanagement.ServiceSKUsClientListAvailableServiceSKUsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAvailableServiceSKUsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListAvailableServiceSKUsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAvailableServiceSKUsPager) {
		s.newListAvailableServiceSKUsPager.remove(req)
	}
	return resp, nil
}
