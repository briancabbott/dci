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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting"
	"net/http"
	"net/url"
	"regexp"
)

// LoadTestsServer is a fake server for instances of the armloadtesting.LoadTestsClient type.
type LoadTestsServer struct {
	// BeginCreateOrUpdate is the fake for method LoadTestsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResource armloadtesting.LoadTestResource, options *armloadtesting.LoadTestsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armloadtesting.LoadTestsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LoadTestsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, loadTestName string, options *armloadtesting.LoadTestsClientBeginDeleteOptions) (resp azfake.PollerResponder[armloadtesting.LoadTestsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LoadTestsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, loadTestName string, options *armloadtesting.LoadTestsClientGetOptions) (resp azfake.Responder[armloadtesting.LoadTestsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method LoadTestsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armloadtesting.LoadTestsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armloadtesting.LoadTestsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method LoadTestsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armloadtesting.LoadTestsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armloadtesting.LoadTestsClientListBySubscriptionResponse])

	// NewListOutboundNetworkDependenciesEndpointsPager is the fake for method LoadTestsClient.NewListOutboundNetworkDependenciesEndpointsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOutboundNetworkDependenciesEndpointsPager func(resourceGroupName string, loadTestName string, options *armloadtesting.LoadTestsClientListOutboundNetworkDependenciesEndpointsOptions) (resp azfake.PagerResponder[armloadtesting.LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse])

	// BeginUpdate is the fake for method LoadTestsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, loadTestName string, loadTestResourcePatchRequestBody armloadtesting.LoadTestResourcePatchRequestBody, options *armloadtesting.LoadTestsClientBeginUpdateOptions) (resp azfake.PollerResponder[armloadtesting.LoadTestsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLoadTestsServerTransport creates a new instance of LoadTestsServerTransport with the provided implementation.
// The returned LoadTestsServerTransport instance is connected to an instance of armloadtesting.LoadTestsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLoadTestsServerTransport(srv *LoadTestsServer) *LoadTestsServerTransport {
	return &LoadTestsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armloadtesting.LoadTestsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armloadtesting.LoadTestsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListBySubscriptionResponse]](),
		newListOutboundNetworkDependenciesEndpointsPager: newTracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse]](),
		beginUpdate: newTracker[azfake.PollerResponder[armloadtesting.LoadTestsClientUpdateResponse]](),
	}
}

// LoadTestsServerTransport connects instances of armloadtesting.LoadTestsClient to instances of LoadTestsServer.
// Don't use this type directly, use NewLoadTestsServerTransport instead.
type LoadTestsServerTransport struct {
	srv                                              *LoadTestsServer
	beginCreateOrUpdate                              *tracker[azfake.PollerResponder[armloadtesting.LoadTestsClientCreateOrUpdateResponse]]
	beginDelete                                      *tracker[azfake.PollerResponder[armloadtesting.LoadTestsClientDeleteResponse]]
	newListByResourceGroupPager                      *tracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListByResourceGroupResponse]]
	newListBySubscriptionPager                       *tracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListBySubscriptionResponse]]
	newListOutboundNetworkDependenciesEndpointsPager *tracker[azfake.PagerResponder[armloadtesting.LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse]]
	beginUpdate                                      *tracker[azfake.PollerResponder[armloadtesting.LoadTestsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LoadTestsServerTransport.
func (l *LoadTestsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LoadTestsClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LoadTestsClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LoadTestsClient.Get":
		resp, err = l.dispatchGet(req)
	case "LoadTestsClient.NewListByResourceGroupPager":
		resp, err = l.dispatchNewListByResourceGroupPager(req)
	case "LoadTestsClient.NewListBySubscriptionPager":
		resp, err = l.dispatchNewListBySubscriptionPager(req)
	case "LoadTestsClient.NewListOutboundNetworkDependenciesEndpointsPager":
		resp, err = l.dispatchNewListOutboundNetworkDependenciesEndpointsPager(req)
	case "LoadTestsClient.BeginUpdate":
		resp, err = l.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armloadtesting.LoadTestResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, loadTestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, loadTestNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, loadTestNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LoadTestResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := l.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		l.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armloadtesting.LoadTestsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		l.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := l.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := l.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		l.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armloadtesting.LoadTestsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		l.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchNewListOutboundNetworkDependenciesEndpointsPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListOutboundNetworkDependenciesEndpointsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOutboundNetworkDependenciesEndpointsPager not implemented")}
	}
	newListOutboundNetworkDependenciesEndpointsPager := l.newListOutboundNetworkDependenciesEndpointsPager.get(req)
	if newListOutboundNetworkDependenciesEndpointsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/outboundNetworkDependenciesEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListOutboundNetworkDependenciesEndpointsPager(resourceGroupNameParam, loadTestNameParam, nil)
		newListOutboundNetworkDependenciesEndpointsPager = &resp
		l.newListOutboundNetworkDependenciesEndpointsPager.add(req, newListOutboundNetworkDependenciesEndpointsPager)
		server.PagerResponderInjectNextLinks(newListOutboundNetworkDependenciesEndpointsPager, req, func(page *armloadtesting.LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOutboundNetworkDependenciesEndpointsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListOutboundNetworkDependenciesEndpointsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOutboundNetworkDependenciesEndpointsPager) {
		l.newListOutboundNetworkDependenciesEndpointsPager.remove(req)
	}
	return resp, nil
}

func (l *LoadTestsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LoadTestService/loadTests/(?P<loadTestName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armloadtesting.LoadTestResourcePatchRequestBody](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		loadTestNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("loadTestName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, loadTestNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}
