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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AFDEndpointsServer is a fake server for instances of the armcdn.AFDEndpointsClient type.
type AFDEndpointsServer struct {
	// BeginCreate is the fake for method AFDEndpointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint armcdn.AFDEndpoint, options *armcdn.AFDEndpointsClientBeginCreateOptions) (resp azfake.PollerResponder[armcdn.AFDEndpointsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AFDEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.AFDEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcdn.AFDEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AFDEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.AFDEndpointsClientGetOptions) (resp azfake.Responder[armcdn.AFDEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProfilePager is the fake for method AFDEndpointsClient.NewListByProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProfilePager func(resourceGroupName string, profileName string, options *armcdn.AFDEndpointsClientListByProfileOptions) (resp azfake.PagerResponder[armcdn.AFDEndpointsClientListByProfileResponse])

	// NewListResourceUsagePager is the fake for method AFDEndpointsClient.NewListResourceUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListResourceUsagePager func(resourceGroupName string, profileName string, endpointName string, options *armcdn.AFDEndpointsClientListResourceUsageOptions) (resp azfake.PagerResponder[armcdn.AFDEndpointsClientListResourceUsageResponse])

	// BeginPurgeContent is the fake for method AFDEndpointsClient.BeginPurgeContent
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPurgeContent func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contents armcdn.AfdPurgeParameters, options *armcdn.AFDEndpointsClientBeginPurgeContentOptions) (resp azfake.PollerResponder[armcdn.AFDEndpointsClientPurgeContentResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method AFDEndpointsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties armcdn.AFDEndpointUpdateParameters, options *armcdn.AFDEndpointsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcdn.AFDEndpointsClientUpdateResponse], errResp azfake.ErrorResponder)

	// ValidateCustomDomain is the fake for method AFDEndpointsClient.ValidateCustomDomain
	// HTTP status codes to indicate success: http.StatusOK
	ValidateCustomDomain func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties armcdn.ValidateCustomDomainInput, options *armcdn.AFDEndpointsClientValidateCustomDomainOptions) (resp azfake.Responder[armcdn.AFDEndpointsClientValidateCustomDomainResponse], errResp azfake.ErrorResponder)
}

// NewAFDEndpointsServerTransport creates a new instance of AFDEndpointsServerTransport with the provided implementation.
// The returned AFDEndpointsServerTransport instance is connected to an instance of armcdn.AFDEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAFDEndpointsServerTransport(srv *AFDEndpointsServer) *AFDEndpointsServerTransport {
	return &AFDEndpointsServerTransport{
		srv:                       srv,
		beginCreate:               newTracker[azfake.PollerResponder[armcdn.AFDEndpointsClientCreateResponse]](),
		beginDelete:               newTracker[azfake.PollerResponder[armcdn.AFDEndpointsClientDeleteResponse]](),
		newListByProfilePager:     newTracker[azfake.PagerResponder[armcdn.AFDEndpointsClientListByProfileResponse]](),
		newListResourceUsagePager: newTracker[azfake.PagerResponder[armcdn.AFDEndpointsClientListResourceUsageResponse]](),
		beginPurgeContent:         newTracker[azfake.PollerResponder[armcdn.AFDEndpointsClientPurgeContentResponse]](),
		beginUpdate:               newTracker[azfake.PollerResponder[armcdn.AFDEndpointsClientUpdateResponse]](),
	}
}

// AFDEndpointsServerTransport connects instances of armcdn.AFDEndpointsClient to instances of AFDEndpointsServer.
// Don't use this type directly, use NewAFDEndpointsServerTransport instead.
type AFDEndpointsServerTransport struct {
	srv                       *AFDEndpointsServer
	beginCreate               *tracker[azfake.PollerResponder[armcdn.AFDEndpointsClientCreateResponse]]
	beginDelete               *tracker[azfake.PollerResponder[armcdn.AFDEndpointsClientDeleteResponse]]
	newListByProfilePager     *tracker[azfake.PagerResponder[armcdn.AFDEndpointsClientListByProfileResponse]]
	newListResourceUsagePager *tracker[azfake.PagerResponder[armcdn.AFDEndpointsClientListResourceUsageResponse]]
	beginPurgeContent         *tracker[azfake.PollerResponder[armcdn.AFDEndpointsClientPurgeContentResponse]]
	beginUpdate               *tracker[azfake.PollerResponder[armcdn.AFDEndpointsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AFDEndpointsServerTransport.
func (a *AFDEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AFDEndpointsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "AFDEndpointsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AFDEndpointsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AFDEndpointsClient.NewListByProfilePager":
		resp, err = a.dispatchNewListByProfilePager(req)
	case "AFDEndpointsClient.NewListResourceUsagePager":
		resp, err = a.dispatchNewListResourceUsagePager(req)
	case "AFDEndpointsClient.BeginPurgeContent":
		resp, err = a.dispatchBeginPurgeContent(req)
	case "AFDEndpointsClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	case "AFDEndpointsClient.ValidateCustomDomain":
		resp, err = a.dispatchValidateCustomDomain(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := a.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.AFDEndpoint](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreate(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		a.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		a.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		a.beginCreate.remove(req)
	}

	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
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

func (a *AFDEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AFDEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchNewListByProfilePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProfilePager not implemented")}
	}
	newListByProfilePager := a.newListByProfilePager.get(req)
	if newListByProfilePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByProfilePager(resourceGroupNameParam, profileNameParam, nil)
		newListByProfilePager = &resp
		a.newListByProfilePager.add(req, newListByProfilePager)
		server.PagerResponderInjectNextLinks(newListByProfilePager, req, func(page *armcdn.AFDEndpointsClientListByProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProfilePager) {
		a.newListByProfilePager.remove(req)
	}
	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchNewListResourceUsagePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListResourceUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListResourceUsagePager not implemented")}
	}
	newListResourceUsagePager := a.newListResourceUsagePager.get(req)
	if newListResourceUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListResourceUsagePager(resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
		newListResourceUsagePager = &resp
		a.newListResourceUsagePager.add(req, newListResourceUsagePager)
		server.PagerResponderInjectNextLinks(newListResourceUsagePager, req, func(page *armcdn.AFDEndpointsClientListResourceUsageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListResourceUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListResourceUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListResourceUsagePager) {
		a.newListResourceUsagePager.remove(req)
	}
	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchBeginPurgeContent(req *http.Request) (*http.Response, error) {
	if a.srv.BeginPurgeContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurgeContent not implemented")}
	}
	beginPurgeContent := a.beginPurgeContent.get(req)
	if beginPurgeContent == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/purge`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.AfdPurgeParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginPurgeContent(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurgeContent = &respr
		a.beginPurgeContent.add(req, beginPurgeContent)
	}

	resp, err := server.PollerResponderNext(beginPurgeContent, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginPurgeContent.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurgeContent) {
		a.beginPurgeContent.remove(req)
	}

	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.AFDEndpointUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}

func (a *AFDEndpointsServerTransport) dispatchValidateCustomDomain(req *http.Request) (*http.Response, error) {
	if a.srv.ValidateCustomDomain == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateCustomDomain not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/afdEndpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateCustomDomain`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcdn.ValidateCustomDomainInput](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ValidateCustomDomain(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ValidateCustomDomainOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
