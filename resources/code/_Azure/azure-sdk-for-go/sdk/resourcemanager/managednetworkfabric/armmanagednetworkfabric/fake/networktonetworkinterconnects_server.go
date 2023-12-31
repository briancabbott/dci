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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkToNetworkInterconnectsServer is a fake server for instances of the armmanagednetworkfabric.NetworkToNetworkInterconnectsClient type.
type NetworkToNetworkInterconnectsServer struct {
	// BeginCreate is the fake for method NetworkToNetworkInterconnectsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body armmanagednetworkfabric.NetworkToNetworkInterconnect, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientBeginCreateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NetworkToNetworkInterconnectsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NetworkToNetworkInterconnectsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientGetOptions) (resp azfake.Responder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNetworkFabricPager is the fake for method NetworkToNetworkInterconnectsClient.NewListByNetworkFabricPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNetworkFabricPager func(resourceGroupName string, networkFabricName string, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientListByNetworkFabricOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientListByNetworkFabricResponse])

	// BeginUpdate is the fake for method NetworkToNetworkInterconnectsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body armmanagednetworkfabric.NetworkToNetworkInterconnectPatch, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientBeginUpdateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateAdministrativeState is the fake for method NetworkToNetworkInterconnectsClient.BeginUpdateAdministrativeState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateAdministrativeState func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body armmanagednetworkfabric.UpdateAdministrativeState, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientBeginUpdateAdministrativeStateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateNpbStaticRouteBfdAdministrativeState is the fake for method NetworkToNetworkInterconnectsClient.BeginUpdateNpbStaticRouteBfdAdministrativeState
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateNpbStaticRouteBfdAdministrativeState func(ctx context.Context, resourceGroupName string, networkFabricName string, networkToNetworkInterconnectName string, body armmanagednetworkfabric.UpdateAdministrativeState, options *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientBeginUpdateNpbStaticRouteBfdAdministrativeStateOptions) (resp azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse], errResp azfake.ErrorResponder)
}

// NewNetworkToNetworkInterconnectsServerTransport creates a new instance of NetworkToNetworkInterconnectsServerTransport with the provided implementation.
// The returned NetworkToNetworkInterconnectsServerTransport instance is connected to an instance of armmanagednetworkfabric.NetworkToNetworkInterconnectsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkToNetworkInterconnectsServerTransport(srv *NetworkToNetworkInterconnectsServer) *NetworkToNetworkInterconnectsServerTransport {
	return &NetworkToNetworkInterconnectsServerTransport{
		srv:                            srv,
		beginCreate:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientCreateResponse]](),
		beginDelete:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientDeleteResponse]](),
		newListByNetworkFabricPager:    newTracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientListByNetworkFabricResponse]](),
		beginUpdate:                    newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateResponse]](),
		beginUpdateAdministrativeState: newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse]](),
		beginUpdateNpbStaticRouteBfdAdministrativeState: newTracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse]](),
	}
}

// NetworkToNetworkInterconnectsServerTransport connects instances of armmanagednetworkfabric.NetworkToNetworkInterconnectsClient to instances of NetworkToNetworkInterconnectsServer.
// Don't use this type directly, use NewNetworkToNetworkInterconnectsServerTransport instead.
type NetworkToNetworkInterconnectsServerTransport struct {
	srv                                             *NetworkToNetworkInterconnectsServer
	beginCreate                                     *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientCreateResponse]]
	beginDelete                                     *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientDeleteResponse]]
	newListByNetworkFabricPager                     *tracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientListByNetworkFabricResponse]]
	beginUpdate                                     *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateResponse]]
	beginUpdateAdministrativeState                  *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateAdministrativeStateResponse]]
	beginUpdateNpbStaticRouteBfdAdministrativeState *tracker[azfake.PollerResponder[armmanagednetworkfabric.NetworkToNetworkInterconnectsClientUpdateNpbStaticRouteBfdAdministrativeStateResponse]]
}

// Do implements the policy.Transporter interface for NetworkToNetworkInterconnectsServerTransport.
func (n *NetworkToNetworkInterconnectsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkToNetworkInterconnectsClient.BeginCreate":
		resp, err = n.dispatchBeginCreate(req)
	case "NetworkToNetworkInterconnectsClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NetworkToNetworkInterconnectsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkToNetworkInterconnectsClient.NewListByNetworkFabricPager":
		resp, err = n.dispatchNewListByNetworkFabricPager(req)
	case "NetworkToNetworkInterconnectsClient.BeginUpdate":
		resp, err = n.dispatchBeginUpdate(req)
	case "NetworkToNetworkInterconnectsClient.BeginUpdateAdministrativeState":
		resp, err = n.dispatchBeginUpdateAdministrativeState(req)
	case "NetworkToNetworkInterconnectsClient.BeginUpdateNpbStaticRouteBfdAdministrativeState":
		resp, err = n.dispatchBeginUpdateNpbStaticRouteBfdAdministrativeState(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := n.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.NetworkToNetworkInterconnect](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreate(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		n.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		n.beginCreate.remove(req)
	}

	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
	if err != nil {
		return nil, err
	}
	networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkToNetworkInterconnect, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchNewListByNetworkFabricPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByNetworkFabricPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNetworkFabricPager not implemented")}
	}
	newListByNetworkFabricPager := n.newListByNetworkFabricPager.get(req)
	if newListByNetworkFabricPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByNetworkFabricPager(resourceGroupNameParam, networkFabricNameParam, nil)
		newListByNetworkFabricPager = &resp
		n.newListByNetworkFabricPager.add(req, newListByNetworkFabricPager)
		server.PagerResponderInjectNextLinks(newListByNetworkFabricPager, req, func(page *armmanagednetworkfabric.NetworkToNetworkInterconnectsClientListByNetworkFabricResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNetworkFabricPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByNetworkFabricPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNetworkFabricPager) {
		n.newListByNetworkFabricPager.remove(req)
	}
	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := n.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.NetworkToNetworkInterconnectPatch](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdate(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, body, nil)
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

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchBeginUpdateAdministrativeState(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdateAdministrativeState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateAdministrativeState not implemented")}
	}
	beginUpdateAdministrativeState := n.beginUpdateAdministrativeState.get(req)
	if beginUpdateAdministrativeState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateAdministrativeState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.UpdateAdministrativeState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdateAdministrativeState(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateAdministrativeState = &respr
		n.beginUpdateAdministrativeState.add(req, beginUpdateAdministrativeState)
	}

	resp, err := server.PollerResponderNext(beginUpdateAdministrativeState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdateAdministrativeState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateAdministrativeState) {
		n.beginUpdateAdministrativeState.remove(req)
	}

	return resp, nil
}

func (n *NetworkToNetworkInterconnectsServerTransport) dispatchBeginUpdateNpbStaticRouteBfdAdministrativeState(req *http.Request) (*http.Response, error) {
	if n.srv.BeginUpdateNpbStaticRouteBfdAdministrativeState == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateNpbStaticRouteBfdAdministrativeState not implemented")}
	}
	beginUpdateNpbStaticRouteBfdAdministrativeState := n.beginUpdateNpbStaticRouteBfdAdministrativeState.get(req)
	if beginUpdateNpbStaticRouteBfdAdministrativeState == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabrics/(?P<networkFabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkToNetworkInterconnects/(?P<networkToNetworkInterconnectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateNpbStaticRouteBfdAdministrativeState`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagednetworkfabric.UpdateAdministrativeState](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkFabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricName")])
		if err != nil {
			return nil, err
		}
		networkToNetworkInterconnectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkToNetworkInterconnectName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginUpdateNpbStaticRouteBfdAdministrativeState(req.Context(), resourceGroupNameParam, networkFabricNameParam, networkToNetworkInterconnectNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateNpbStaticRouteBfdAdministrativeState = &respr
		n.beginUpdateNpbStaticRouteBfdAdministrativeState.add(req, beginUpdateNpbStaticRouteBfdAdministrativeState)
	}

	resp, err := server.PollerResponderNext(beginUpdateNpbStaticRouteBfdAdministrativeState, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginUpdateNpbStaticRouteBfdAdministrativeState.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateNpbStaticRouteBfdAdministrativeState) {
		n.beginUpdateNpbStaticRouteBfdAdministrativeState.remove(req)
	}

	return resp, nil
}
