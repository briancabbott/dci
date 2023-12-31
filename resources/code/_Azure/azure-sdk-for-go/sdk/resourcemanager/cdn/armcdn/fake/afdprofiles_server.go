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

// AFDProfilesServer is a fake server for instances of the armcdn.AFDProfilesClient type.
type AFDProfilesServer struct {
	// CheckEndpointNameAvailability is the fake for method AFDProfilesClient.CheckEndpointNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckEndpointNameAvailability func(ctx context.Context, resourceGroupName string, profileName string, checkEndpointNameAvailabilityInput armcdn.CheckEndpointNameAvailabilityInput, options *armcdn.AFDProfilesClientCheckEndpointNameAvailabilityOptions) (resp azfake.Responder[armcdn.AFDProfilesClientCheckEndpointNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// CheckHostNameAvailability is the fake for method AFDProfilesClient.CheckHostNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckHostNameAvailability func(ctx context.Context, resourceGroupName string, profileName string, checkHostNameAvailabilityInput armcdn.CheckHostNameAvailabilityInput, options *armcdn.AFDProfilesClientCheckHostNameAvailabilityOptions) (resp azfake.Responder[armcdn.AFDProfilesClientCheckHostNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// NewListResourceUsagePager is the fake for method AFDProfilesClient.NewListResourceUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListResourceUsagePager func(resourceGroupName string, profileName string, options *armcdn.AFDProfilesClientListResourceUsageOptions) (resp azfake.PagerResponder[armcdn.AFDProfilesClientListResourceUsageResponse])

	// BeginUpgrade is the fake for method AFDProfilesClient.BeginUpgrade
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpgrade func(ctx context.Context, resourceGroupName string, profileName string, profileUpgradeParameters armcdn.ProfileUpgradeParameters, options *armcdn.AFDProfilesClientBeginUpgradeOptions) (resp azfake.PollerResponder[armcdn.AFDProfilesClientUpgradeResponse], errResp azfake.ErrorResponder)

	// ValidateSecret is the fake for method AFDProfilesClient.ValidateSecret
	// HTTP status codes to indicate success: http.StatusOK
	ValidateSecret func(ctx context.Context, resourceGroupName string, profileName string, validateSecretInput armcdn.ValidateSecretInput, options *armcdn.AFDProfilesClientValidateSecretOptions) (resp azfake.Responder[armcdn.AFDProfilesClientValidateSecretResponse], errResp azfake.ErrorResponder)
}

// NewAFDProfilesServerTransport creates a new instance of AFDProfilesServerTransport with the provided implementation.
// The returned AFDProfilesServerTransport instance is connected to an instance of armcdn.AFDProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAFDProfilesServerTransport(srv *AFDProfilesServer) *AFDProfilesServerTransport {
	return &AFDProfilesServerTransport{
		srv:                       srv,
		newListResourceUsagePager: newTracker[azfake.PagerResponder[armcdn.AFDProfilesClientListResourceUsageResponse]](),
		beginUpgrade:              newTracker[azfake.PollerResponder[armcdn.AFDProfilesClientUpgradeResponse]](),
	}
}

// AFDProfilesServerTransport connects instances of armcdn.AFDProfilesClient to instances of AFDProfilesServer.
// Don't use this type directly, use NewAFDProfilesServerTransport instead.
type AFDProfilesServerTransport struct {
	srv                       *AFDProfilesServer
	newListResourceUsagePager *tracker[azfake.PagerResponder[armcdn.AFDProfilesClientListResourceUsageResponse]]
	beginUpgrade              *tracker[azfake.PollerResponder[armcdn.AFDProfilesClientUpgradeResponse]]
}

// Do implements the policy.Transporter interface for AFDProfilesServerTransport.
func (a *AFDProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AFDProfilesClient.CheckEndpointNameAvailability":
		resp, err = a.dispatchCheckEndpointNameAvailability(req)
	case "AFDProfilesClient.CheckHostNameAvailability":
		resp, err = a.dispatchCheckHostNameAvailability(req)
	case "AFDProfilesClient.NewListResourceUsagePager":
		resp, err = a.dispatchNewListResourceUsagePager(req)
	case "AFDProfilesClient.BeginUpgrade":
		resp, err = a.dispatchBeginUpgrade(req)
	case "AFDProfilesClient.ValidateSecret":
		resp, err = a.dispatchValidateSecret(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AFDProfilesServerTransport) dispatchCheckEndpointNameAvailability(req *http.Request) (*http.Response, error) {
	if a.srv.CheckEndpointNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckEndpointNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkEndpointNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcdn.CheckEndpointNameAvailabilityInput](req)
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
	respr, errRespr := a.srv.CheckEndpointNameAvailability(req.Context(), resourceGroupNameParam, profileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckEndpointNameAvailabilityOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AFDProfilesServerTransport) dispatchCheckHostNameAvailability(req *http.Request) (*http.Response, error) {
	if a.srv.CheckHostNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckHostNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkHostNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcdn.CheckHostNameAvailabilityInput](req)
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
	respr, errRespr := a.srv.CheckHostNameAvailability(req.Context(), resourceGroupNameParam, profileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AFDProfilesServerTransport) dispatchNewListResourceUsagePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListResourceUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListResourceUsagePager not implemented")}
	}
	newListResourceUsagePager := a.newListResourceUsagePager.get(req)
	if newListResourceUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
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
		resp := a.srv.NewListResourceUsagePager(resourceGroupNameParam, profileNameParam, nil)
		newListResourceUsagePager = &resp
		a.newListResourceUsagePager.add(req, newListResourceUsagePager)
		server.PagerResponderInjectNextLinks(newListResourceUsagePager, req, func(page *armcdn.AFDProfilesClientListResourceUsageResponse, createLink func() string) {
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

func (a *AFDProfilesServerTransport) dispatchBeginUpgrade(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpgrade == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgrade not implemented")}
	}
	beginUpgrade := a.beginUpgrade.get(req)
	if beginUpgrade == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgrade`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.ProfileUpgradeParameters](req)
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
		respr, errRespr := a.srv.BeginUpgrade(req.Context(), resourceGroupNameParam, profileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgrade = &respr
		a.beginUpgrade.add(req, beginUpgrade)
	}

	resp, err := server.PollerResponderNext(beginUpgrade, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpgrade.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgrade) {
		a.beginUpgrade.remove(req)
	}

	return resp, nil
}

func (a *AFDProfilesServerTransport) dispatchValidateSecret(req *http.Request) (*http.Response, error) {
	if a.srv.ValidateSecret == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateSecret not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateSecret`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcdn.ValidateSecretInput](req)
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
	respr, errRespr := a.srv.ValidateSecret(req.Context(), resourceGroupNameParam, profileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ValidateSecretOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
