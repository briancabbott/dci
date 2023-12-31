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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectorsServer is a fake server for instances of the armcustomerinsights.ConnectorsClient type.
type ConnectorsServer struct {
	// BeginCreateOrUpdate is the fake for method ConnectorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hubName string, connectorName string, parameters armcustomerinsights.ConnectorResourceFormat, options *armcustomerinsights.ConnectorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcustomerinsights.ConnectorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, hubName string, connectorName string, options *armcustomerinsights.ConnectorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcustomerinsights.ConnectorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hubName string, connectorName string, options *armcustomerinsights.ConnectorsClientGetOptions) (resp azfake.Responder[armcustomerinsights.ConnectorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHubPager is the fake for method ConnectorsClient.NewListByHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHubPager func(resourceGroupName string, hubName string, options *armcustomerinsights.ConnectorsClientListByHubOptions) (resp azfake.PagerResponder[armcustomerinsights.ConnectorsClientListByHubResponse])
}

// NewConnectorsServerTransport creates a new instance of ConnectorsServerTransport with the provided implementation.
// The returned ConnectorsServerTransport instance is connected to an instance of armcustomerinsights.ConnectorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectorsServerTransport(srv *ConnectorsServer) *ConnectorsServerTransport {
	return &ConnectorsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcustomerinsights.ConnectorsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcustomerinsights.ConnectorsClientDeleteResponse]](),
		newListByHubPager:   newTracker[azfake.PagerResponder[armcustomerinsights.ConnectorsClientListByHubResponse]](),
	}
}

// ConnectorsServerTransport connects instances of armcustomerinsights.ConnectorsClient to instances of ConnectorsServer.
// Don't use this type directly, use NewConnectorsServerTransport instead.
type ConnectorsServerTransport struct {
	srv                 *ConnectorsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcustomerinsights.ConnectorsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcustomerinsights.ConnectorsClientDeleteResponse]]
	newListByHubPager   *tracker[azfake.PagerResponder[armcustomerinsights.ConnectorsClientListByHubResponse]]
}

// Do implements the policy.Transporter interface for ConnectorsServerTransport.
func (c *ConnectorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectorsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ConnectorsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ConnectorsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectorsClient.NewListByHubPager":
		resp, err = c.dispatchNewListByHubPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectors/(?P<connectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.ConnectorResourceFormat](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		connectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, hubNameParam, connectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ConnectorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectors/(?P<connectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		connectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, hubNameParam, connectorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConnectorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectors/(?P<connectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	connectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, hubNameParam, connectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectorResourceFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectorsServerTransport) dispatchNewListByHubPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHubPager not implemented")}
	}
	newListByHubPager := c.newListByHubPager.get(req)
	if newListByHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByHubPager(resourceGroupNameParam, hubNameParam, nil)
		newListByHubPager = &resp
		c.newListByHubPager.add(req, newListByHubPager)
		server.PagerResponderInjectNextLinks(newListByHubPager, req, func(page *armcustomerinsights.ConnectorsClientListByHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHubPager) {
		c.newListByHubPager.remove(req)
	}
	return resp, nil
}
