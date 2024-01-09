//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// AllowedConnectionsServer is a fake server for instances of the armsecurity.AllowedConnectionsClient type.
type AllowedConnectionsServer struct {
	// Get is the fake for method AllowedConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ascLocation string, connectionType armsecurity.ConnectionType, options *armsecurity.AllowedConnectionsClientGetOptions) (resp azfake.Responder[armsecurity.AllowedConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AllowedConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsecurity.AllowedConnectionsClientListOptions) (resp azfake.PagerResponder[armsecurity.AllowedConnectionsClientListResponse])

	// NewListByHomeRegionPager is the fake for method AllowedConnectionsClient.NewListByHomeRegionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHomeRegionPager func(ascLocation string, options *armsecurity.AllowedConnectionsClientListByHomeRegionOptions) (resp azfake.PagerResponder[armsecurity.AllowedConnectionsClientListByHomeRegionResponse])
}

// NewAllowedConnectionsServerTransport creates a new instance of AllowedConnectionsServerTransport with the provided implementation.
// The returned AllowedConnectionsServerTransport instance is connected to an instance of armsecurity.AllowedConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAllowedConnectionsServerTransport(srv *AllowedConnectionsServer) *AllowedConnectionsServerTransport {
	return &AllowedConnectionsServerTransport{
		srv:                      srv,
		newListPager:             newTracker[azfake.PagerResponder[armsecurity.AllowedConnectionsClientListResponse]](),
		newListByHomeRegionPager: newTracker[azfake.PagerResponder[armsecurity.AllowedConnectionsClientListByHomeRegionResponse]](),
	}
}

// AllowedConnectionsServerTransport connects instances of armsecurity.AllowedConnectionsClient to instances of AllowedConnectionsServer.
// Don't use this type directly, use NewAllowedConnectionsServerTransport instead.
type AllowedConnectionsServerTransport struct {
	srv                      *AllowedConnectionsServer
	newListPager             *tracker[azfake.PagerResponder[armsecurity.AllowedConnectionsClientListResponse]]
	newListByHomeRegionPager *tracker[azfake.PagerResponder[armsecurity.AllowedConnectionsClientListByHomeRegionResponse]]
}

// Do implements the policy.Transporter interface for AllowedConnectionsServerTransport.
func (a *AllowedConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AllowedConnectionsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AllowedConnectionsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AllowedConnectionsClient.NewListByHomeRegionPager":
		resp, err = a.dispatchNewListByHomeRegionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AllowedConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/locations/(?P<ascLocation>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/allowedConnections/(?P<connectionType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	ascLocationParam, err := url.PathUnescape(matches[regex.SubexpIndex("ascLocation")])
	if err != nil {
		return nil, err
	}
	connectionTypeParam, err := parseWithCast(matches[regex.SubexpIndex("connectionType")], func(v string) (armsecurity.ConnectionType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armsecurity.ConnectionType(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, ascLocationParam, connectionTypeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AllowedConnectionsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AllowedConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/allowedConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.AllowedConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AllowedConnectionsServerTransport) dispatchNewListByHomeRegionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByHomeRegionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHomeRegionPager not implemented")}
	}
	newListByHomeRegionPager := a.newListByHomeRegionPager.get(req)
	if newListByHomeRegionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/locations/(?P<ascLocation>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/allowedConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		ascLocationParam, err := url.PathUnescape(matches[regex.SubexpIndex("ascLocation")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByHomeRegionPager(ascLocationParam, nil)
		newListByHomeRegionPager = &resp
		a.newListByHomeRegionPager.add(req, newListByHomeRegionPager)
		server.PagerResponderInjectNextLinks(newListByHomeRegionPager, req, func(page *armsecurity.AllowedConnectionsClientListByHomeRegionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHomeRegionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByHomeRegionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHomeRegionPager) {
		a.newListByHomeRegionPager.remove(req)
	}
	return resp, nil
}
