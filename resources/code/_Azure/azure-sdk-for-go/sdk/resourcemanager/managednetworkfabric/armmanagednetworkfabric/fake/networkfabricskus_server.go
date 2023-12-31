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

// NetworkFabricSKUsServer is a fake server for instances of the armmanagednetworkfabric.NetworkFabricSKUsClient type.
type NetworkFabricSKUsServer struct {
	// Get is the fake for method NetworkFabricSKUsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, networkFabricSKUName string, options *armmanagednetworkfabric.NetworkFabricSKUsClientGetOptions) (resp azfake.Responder[armmanagednetworkfabric.NetworkFabricSKUsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method NetworkFabricSKUsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmanagednetworkfabric.NetworkFabricSKUsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmanagednetworkfabric.NetworkFabricSKUsClientListBySubscriptionResponse])
}

// NewNetworkFabricSKUsServerTransport creates a new instance of NetworkFabricSKUsServerTransport with the provided implementation.
// The returned NetworkFabricSKUsServerTransport instance is connected to an instance of armmanagednetworkfabric.NetworkFabricSKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkFabricSKUsServerTransport(srv *NetworkFabricSKUsServer) *NetworkFabricSKUsServerTransport {
	return &NetworkFabricSKUsServerTransport{
		srv:                        srv,
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkFabricSKUsClientListBySubscriptionResponse]](),
	}
}

// NetworkFabricSKUsServerTransport connects instances of armmanagednetworkfabric.NetworkFabricSKUsClient to instances of NetworkFabricSKUsServer.
// Don't use this type directly, use NewNetworkFabricSKUsServerTransport instead.
type NetworkFabricSKUsServerTransport struct {
	srv                        *NetworkFabricSKUsServer
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armmanagednetworkfabric.NetworkFabricSKUsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for NetworkFabricSKUsServerTransport.
func (n *NetworkFabricSKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NetworkFabricSKUsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NetworkFabricSKUsClient.NewListBySubscriptionPager":
		resp, err = n.dispatchNewListBySubscriptionPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NetworkFabricSKUsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabricSkus/(?P<networkFabricSkuName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	networkFabricSKUNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkFabricSkuName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), networkFabricSKUNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkFabricSKU, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkFabricSKUsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := n.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedNetworkFabric/networkFabricSkus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := n.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		n.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmanagednetworkfabric.NetworkFabricSKUsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		n.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}
