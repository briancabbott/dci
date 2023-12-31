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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerGroupsServer is a fake server for instances of the armcontainerinstance.ContainerGroupsClient type.
type ContainerGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method ContainerGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup armcontainerinstance.ContainerGroup, options *armcontainerinstance.ContainerGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ContainerGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ContainerGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientGetOptions) (resp azfake.Responder[armcontainerinstance.ContainerGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// GetOutboundNetworkDependenciesEndpoints is the fake for method ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints
	// HTTP status codes to indicate success: http.StatusOK
	GetOutboundNetworkDependenciesEndpoints func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions) (resp azfake.Responder[armcontainerinstance.ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ContainerGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcontainerinstance.ContainerGroupsClientListOptions) (resp azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListResponse])

	// NewListByResourceGroupPager is the fake for method ContainerGroupsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcontainerinstance.ContainerGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListByResourceGroupResponse])

	// BeginRestart is the fake for method ContainerGroupsClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusNoContent
	BeginRestart func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientBeginRestartOptions) (resp azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method ContainerGroupsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginStart func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientBeginStartOptions) (resp azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientStartResponse], errResp azfake.ErrorResponder)

	// Stop is the fake for method ContainerGroupsClient.Stop
	// HTTP status codes to indicate success: http.StatusNoContent
	Stop func(ctx context.Context, resourceGroupName string, containerGroupName string, options *armcontainerinstance.ContainerGroupsClientStopOptions) (resp azfake.Responder[armcontainerinstance.ContainerGroupsClientStopResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ContainerGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, containerGroupName string, resource armcontainerinstance.Resource, options *armcontainerinstance.ContainerGroupsClientUpdateOptions) (resp azfake.Responder[armcontainerinstance.ContainerGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewContainerGroupsServerTransport creates a new instance of ContainerGroupsServerTransport with the provided implementation.
// The returned ContainerGroupsServerTransport instance is connected to an instance of armcontainerinstance.ContainerGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerGroupsServerTransport(srv *ContainerGroupsServer) *ContainerGroupsServerTransport {
	return &ContainerGroupsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListByResourceGroupResponse]](),
		beginRestart:                newTracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientRestartResponse]](),
		beginStart:                  newTracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientStartResponse]](),
	}
}

// ContainerGroupsServerTransport connects instances of armcontainerinstance.ContainerGroupsClient to instances of ContainerGroupsServer.
// Don't use this type directly, use NewContainerGroupsServerTransport instead.
type ContainerGroupsServerTransport struct {
	srv                         *ContainerGroupsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armcontainerinstance.ContainerGroupsClientListByResourceGroupResponse]]
	beginRestart                *tracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientRestartResponse]]
	beginStart                  *tracker[azfake.PollerResponder[armcontainerinstance.ContainerGroupsClientStartResponse]]
}

// Do implements the policy.Transporter interface for ContainerGroupsServerTransport.
func (c *ContainerGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerGroupsClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "ContainerGroupsClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ContainerGroupsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints":
		resp, err = c.dispatchGetOutboundNetworkDependenciesEndpoints(req)
	case "ContainerGroupsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "ContainerGroupsClient.NewListByResourceGroupPager":
		resp, err = c.dispatchNewListByResourceGroupPager(req)
	case "ContainerGroupsClient.BeginRestart":
		resp, err = c.dispatchBeginRestart(req)
	case "ContainerGroupsClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "ContainerGroupsClient.Stop":
		resp, err = c.dispatchStop(req)
	case "ContainerGroupsClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerinstance.ContainerGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, containerGroupNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
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

func (c *ContainerGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchGetOutboundNetworkDependenciesEndpoints(req *http.Request) (*http.Response, error) {
	if c.srv.GetOutboundNetworkDependenciesEndpoints == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOutboundNetworkDependenciesEndpoints not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundNetworkDependenciesEndpoints`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetOutboundNetworkDependenciesEndpoints(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListPager(nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcontainerinstance.ContainerGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcontainerinstance.ContainerGroupsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := c.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginRestart(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		c.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusNoContent}, resp.StatusCode) {
		c.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		c.beginRestart.remove(req)
	}

	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := c.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		c.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		c.beginStart.remove(req)
	}

	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if c.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Stop(req.Context(), resourceGroupNameParam, containerGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroups/(?P<containerGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerinstance.Resource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, containerGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
