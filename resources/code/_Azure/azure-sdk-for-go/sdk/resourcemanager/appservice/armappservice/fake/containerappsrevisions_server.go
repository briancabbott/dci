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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsRevisionsServer is a fake server for instances of the armappservice.ContainerAppsRevisionsClient type.
type ContainerAppsRevisionsServer struct {
	// ActivateRevision is the fake for method ContainerAppsRevisionsClient.ActivateRevision
	// HTTP status codes to indicate success: http.StatusOK
	ActivateRevision func(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *armappservice.ContainerAppsRevisionsClientActivateRevisionOptions) (resp azfake.Responder[armappservice.ContainerAppsRevisionsClientActivateRevisionResponse], errResp azfake.ErrorResponder)

	// DeactivateRevision is the fake for method ContainerAppsRevisionsClient.DeactivateRevision
	// HTTP status codes to indicate success: http.StatusOK
	DeactivateRevision func(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *armappservice.ContainerAppsRevisionsClientDeactivateRevisionOptions) (resp azfake.Responder[armappservice.ContainerAppsRevisionsClientDeactivateRevisionResponse], errResp azfake.ErrorResponder)

	// GetRevision is the fake for method ContainerAppsRevisionsClient.GetRevision
	// HTTP status codes to indicate success: http.StatusOK
	GetRevision func(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *armappservice.ContainerAppsRevisionsClientGetRevisionOptions) (resp azfake.Responder[armappservice.ContainerAppsRevisionsClientGetRevisionResponse], errResp azfake.ErrorResponder)

	// NewListRevisionsPager is the fake for method ContainerAppsRevisionsClient.NewListRevisionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListRevisionsPager func(resourceGroupName string, containerAppName string, options *armappservice.ContainerAppsRevisionsClientListRevisionsOptions) (resp azfake.PagerResponder[armappservice.ContainerAppsRevisionsClientListRevisionsResponse])

	// RestartRevision is the fake for method ContainerAppsRevisionsClient.RestartRevision
	// HTTP status codes to indicate success: http.StatusOK
	RestartRevision func(ctx context.Context, resourceGroupName string, containerAppName string, name string, options *armappservice.ContainerAppsRevisionsClientRestartRevisionOptions) (resp azfake.Responder[armappservice.ContainerAppsRevisionsClientRestartRevisionResponse], errResp azfake.ErrorResponder)
}

// NewContainerAppsRevisionsServerTransport creates a new instance of ContainerAppsRevisionsServerTransport with the provided implementation.
// The returned ContainerAppsRevisionsServerTransport instance is connected to an instance of armappservice.ContainerAppsRevisionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsRevisionsServerTransport(srv *ContainerAppsRevisionsServer) *ContainerAppsRevisionsServerTransport {
	return &ContainerAppsRevisionsServerTransport{
		srv:                   srv,
		newListRevisionsPager: newTracker[azfake.PagerResponder[armappservice.ContainerAppsRevisionsClientListRevisionsResponse]](),
	}
}

// ContainerAppsRevisionsServerTransport connects instances of armappservice.ContainerAppsRevisionsClient to instances of ContainerAppsRevisionsServer.
// Don't use this type directly, use NewContainerAppsRevisionsServerTransport instead.
type ContainerAppsRevisionsServerTransport struct {
	srv                   *ContainerAppsRevisionsServer
	newListRevisionsPager *tracker[azfake.PagerResponder[armappservice.ContainerAppsRevisionsClientListRevisionsResponse]]
}

// Do implements the policy.Transporter interface for ContainerAppsRevisionsServerTransport.
func (c *ContainerAppsRevisionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerAppsRevisionsClient.ActivateRevision":
		resp, err = c.dispatchActivateRevision(req)
	case "ContainerAppsRevisionsClient.DeactivateRevision":
		resp, err = c.dispatchDeactivateRevision(req)
	case "ContainerAppsRevisionsClient.GetRevision":
		resp, err = c.dispatchGetRevision(req)
	case "ContainerAppsRevisionsClient.NewListRevisionsPager":
		resp, err = c.dispatchNewListRevisionsPager(req)
	case "ContainerAppsRevisionsClient.RestartRevision":
		resp, err = c.dispatchRestartRevision(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerAppsRevisionsServerTransport) dispatchActivateRevision(req *http.Request) (*http.Response, error) {
	if c.srv.ActivateRevision == nil {
		return nil, &nonRetriableError{errors.New("fake for method ActivateRevision not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/activate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ActivateRevision(req.Context(), resourceGroupNameParam, containerAppNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsRevisionsServerTransport) dispatchDeactivateRevision(req *http.Request) (*http.Response, error) {
	if c.srv.DeactivateRevision == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeactivateRevision not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deactivate`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.DeactivateRevision(req.Context(), resourceGroupNameParam, containerAppNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsRevisionsServerTransport) dispatchGetRevision(req *http.Request) (*http.Response, error) {
	if c.srv.GetRevision == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRevision not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetRevision(req.Context(), resourceGroupNameParam, containerAppNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Revision, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsRevisionsServerTransport) dispatchNewListRevisionsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListRevisionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListRevisionsPager not implemented")}
	}
	newListRevisionsPager := c.newListRevisionsPager.get(req)
	if newListRevisionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListRevisionsPager(resourceGroupNameParam, containerAppNameParam, nil)
		newListRevisionsPager = &resp
		c.newListRevisionsPager.add(req, newListRevisionsPager)
		server.PagerResponderInjectNextLinks(newListRevisionsPager, req, func(page *armappservice.ContainerAppsRevisionsClientListRevisionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListRevisionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListRevisionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListRevisionsPager) {
		c.newListRevisionsPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerAppsRevisionsServerTransport) dispatchRestartRevision(req *http.Request) (*http.Response, error) {
	if c.srv.RestartRevision == nil {
		return nil, &nonRetriableError{errors.New("fake for method RestartRevision not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.RestartRevision(req.Context(), resourceGroupNameParam, containerAppNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
