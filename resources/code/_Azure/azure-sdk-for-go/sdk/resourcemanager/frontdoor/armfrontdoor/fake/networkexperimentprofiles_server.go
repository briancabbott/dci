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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkExperimentProfilesServer is a fake server for instances of the armfrontdoor.NetworkExperimentProfilesClient type.
type NetworkExperimentProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method NetworkExperimentProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, profileName string, resourceGroupName string, parameters armfrontdoor.Profile, options *armfrontdoor.NetworkExperimentProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkExperimentProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, options *armfrontdoor.NetworkExperimentProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkExperimentProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, options *armfrontdoor.NetworkExperimentProfilesClientGetOptions) (resp azfake.Responder[armfrontdoor.NetworkExperimentProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NetworkExperimentProfilesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armfrontdoor.NetworkExperimentProfilesClientListOptions) (resp azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListResponse])

	// NewListByResourceGroupPager is the fake for method NetworkExperimentProfilesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armfrontdoor.NetworkExperimentProfilesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method NetworkExperimentProfilesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, profileName string, parameters armfrontdoor.ProfileUpdateModel, options *armfrontdoor.NetworkExperimentProfilesClientBeginUpdateOptions) (resp azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNetworkExperimentProfilesServerTransport creates a new instance of NetworkExperimentProfilesServerTransport with the provided implementation.
// The returned NetworkExperimentProfilesServerTransport instance is connected to an instance of armfrontdoor.NetworkExperimentProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkExperimentProfilesServerTransport(srv *NetworkExperimentProfilesServer) *NetworkExperimentProfilesServerTransport {
	return &NetworkExperimentProfilesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientUpdateResponse]](),
	}
}

// NetworkExperimentProfilesServerTransport connects instances of armfrontdoor.NetworkExperimentProfilesClient to instances of NetworkExperimentProfilesServer.
// Don't use this type directly, use NewNetworkExperimentProfilesServerTransport instead.
type NetworkExperimentProfilesServerTransport struct {
	srv                         *NetworkExperimentProfilesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armfrontdoor.NetworkExperimentProfilesClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armfrontdoor.NetworkExperimentProfilesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for NetworkExperimentProfilesServerTransport.
func (n *NetworkExperimentProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkExperimentProfilesClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NetworkExperimentProfilesClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkExperimentProfilesClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkExperimentProfilesClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	case "NetworkExperimentProfilesClient.NewListByResourceGroupPager":
		resp, err = n.dispatchNewListByResourceGroupPager(req)
	case "NetworkExperimentProfilesClient.BeginUpdate":
		resp, err = n.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.Profile](req)
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), profileNameParam, resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Profile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListPager(nil)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armfrontdoor.NetworkExperimentProfilesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		n.newListPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := n.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		n.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armfrontdoor.NetworkExperimentProfilesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		n.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkExperimentProfilesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := n.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.ProfileUpdateModel](req)
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
		respr, errRespr := n.srv.BeginUpdate(req.Context(), resourceGroupNameParam, profileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		n.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		n.beginUpdate.remove(req)
	}

	return resp, nil
}
