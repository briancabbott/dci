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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
	"net/http"
	"net/url"
	"regexp"
)

// AccessPolicyServer is a fake server for instances of the armredis.AccessPolicyClient type.
type AccessPolicyServer struct {
	// BeginCreateUpdate is the fake for method AccessPolicyClient.BeginCreateUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateUpdate func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, parameters armredis.CacheAccessPolicy, options *armredis.AccessPolicyClientBeginCreateUpdateOptions) (resp azfake.PollerResponder[armredis.AccessPolicyClientCreateUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AccessPolicyClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *armredis.AccessPolicyClientBeginDeleteOptions) (resp azfake.PollerResponder[armredis.AccessPolicyClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AccessPolicyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *armredis.AccessPolicyClientGetOptions) (resp azfake.Responder[armredis.AccessPolicyClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AccessPolicyClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, cacheName string, options *armredis.AccessPolicyClientListOptions) (resp azfake.PagerResponder[armredis.AccessPolicyClientListResponse])
}

// NewAccessPolicyServerTransport creates a new instance of AccessPolicyServerTransport with the provided implementation.
// The returned AccessPolicyServerTransport instance is connected to an instance of armredis.AccessPolicyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAccessPolicyServerTransport(srv *AccessPolicyServer) *AccessPolicyServerTransport {
	return &AccessPolicyServerTransport{
		srv:               srv,
		beginCreateUpdate: newTracker[azfake.PollerResponder[armredis.AccessPolicyClientCreateUpdateResponse]](),
		beginDelete:       newTracker[azfake.PollerResponder[armredis.AccessPolicyClientDeleteResponse]](),
		newListPager:      newTracker[azfake.PagerResponder[armredis.AccessPolicyClientListResponse]](),
	}
}

// AccessPolicyServerTransport connects instances of armredis.AccessPolicyClient to instances of AccessPolicyServer.
// Don't use this type directly, use NewAccessPolicyServerTransport instead.
type AccessPolicyServerTransport struct {
	srv               *AccessPolicyServer
	beginCreateUpdate *tracker[azfake.PollerResponder[armredis.AccessPolicyClientCreateUpdateResponse]]
	beginDelete       *tracker[azfake.PollerResponder[armredis.AccessPolicyClientDeleteResponse]]
	newListPager      *tracker[azfake.PagerResponder[armredis.AccessPolicyClientListResponse]]
}

// Do implements the policy.Transporter interface for AccessPolicyServerTransport.
func (a *AccessPolicyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AccessPolicyClient.BeginCreateUpdate":
		resp, err = a.dispatchBeginCreateUpdate(req)
	case "AccessPolicyClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AccessPolicyClient.Get":
		resp, err = a.dispatchGet(req)
	case "AccessPolicyClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AccessPolicyServerTransport) dispatchBeginCreateUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateUpdate not implemented")}
	}
	beginCreateUpdate := a.beginCreateUpdate.get(req)
	if beginCreateUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicies/(?P<accessPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armredis.CacheAccessPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		accessPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateUpdate(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateUpdate = &respr
		a.beginCreateUpdate.add(req, beginCreateUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateUpdate) {
		a.beginCreateUpdate.remove(req)
	}

	return resp, nil
}

func (a *AccessPolicyServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicies/(?P<accessPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		accessPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AccessPolicyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicies/(?P<accessPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
	if err != nil {
		return nil, err
	}
	accessPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, cacheNameParam, accessPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CacheAccessPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AccessPolicyServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cache/redis/(?P<cacheName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cacheNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cacheName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, cacheNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armredis.AccessPolicyClientListResponse, createLink func() string) {
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
