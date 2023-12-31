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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
	"net/http"
	"net/url"
	"regexp"
)

// HybridConnectionsServer is a fake server for instances of the armrelay.HybridConnectionsClient type.
type HybridConnectionsServer struct {
	// CreateOrUpdate is the fake for method HybridConnectionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, parameters armrelay.HybridConnection, options *armrelay.HybridConnectionsClientCreateOrUpdateOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAuthorizationRule is the fake for method HybridConnectionsClient.CreateOrUpdateAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters armrelay.AuthorizationRule, options *armrelay.HybridConnectionsClientCreateOrUpdateAuthorizationRuleOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientCreateOrUpdateAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HybridConnectionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *armrelay.HybridConnectionsClientDeleteOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAuthorizationRule is the fake for method HybridConnectionsClient.DeleteAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *armrelay.HybridConnectionsClientDeleteAuthorizationRuleOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientDeleteAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HybridConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, options *armrelay.HybridConnectionsClientGetOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthorizationRule is the fake for method HybridConnectionsClient.GetAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *armrelay.HybridConnectionsClientGetAuthorizationRuleOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientGetAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// NewListAuthorizationRulesPager is the fake for method HybridConnectionsClient.NewListAuthorizationRulesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAuthorizationRulesPager func(resourceGroupName string, namespaceName string, hybridConnectionName string, options *armrelay.HybridConnectionsClientListAuthorizationRulesOptions) (resp azfake.PagerResponder[armrelay.HybridConnectionsClientListAuthorizationRulesResponse])

	// NewListByNamespacePager is the fake for method HybridConnectionsClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armrelay.HybridConnectionsClientListByNamespaceOptions) (resp azfake.PagerResponder[armrelay.HybridConnectionsClientListByNamespaceResponse])

	// ListKeys is the fake for method HybridConnectionsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, options *armrelay.HybridConnectionsClientListKeysOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method HybridConnectionsClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, namespaceName string, hybridConnectionName string, authorizationRuleName string, parameters armrelay.RegenerateAccessKeyParameters, options *armrelay.HybridConnectionsClientRegenerateKeysOptions) (resp azfake.Responder[armrelay.HybridConnectionsClientRegenerateKeysResponse], errResp azfake.ErrorResponder)
}

// NewHybridConnectionsServerTransport creates a new instance of HybridConnectionsServerTransport with the provided implementation.
// The returned HybridConnectionsServerTransport instance is connected to an instance of armrelay.HybridConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHybridConnectionsServerTransport(srv *HybridConnectionsServer) *HybridConnectionsServerTransport {
	return &HybridConnectionsServerTransport{
		srv:                            srv,
		newListAuthorizationRulesPager: newTracker[azfake.PagerResponder[armrelay.HybridConnectionsClientListAuthorizationRulesResponse]](),
		newListByNamespacePager:        newTracker[azfake.PagerResponder[armrelay.HybridConnectionsClientListByNamespaceResponse]](),
	}
}

// HybridConnectionsServerTransport connects instances of armrelay.HybridConnectionsClient to instances of HybridConnectionsServer.
// Don't use this type directly, use NewHybridConnectionsServerTransport instead.
type HybridConnectionsServerTransport struct {
	srv                            *HybridConnectionsServer
	newListAuthorizationRulesPager *tracker[azfake.PagerResponder[armrelay.HybridConnectionsClientListAuthorizationRulesResponse]]
	newListByNamespacePager        *tracker[azfake.PagerResponder[armrelay.HybridConnectionsClientListByNamespaceResponse]]
}

// Do implements the policy.Transporter interface for HybridConnectionsServerTransport.
func (h *HybridConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HybridConnectionsClient.CreateOrUpdate":
		resp, err = h.dispatchCreateOrUpdate(req)
	case "HybridConnectionsClient.CreateOrUpdateAuthorizationRule":
		resp, err = h.dispatchCreateOrUpdateAuthorizationRule(req)
	case "HybridConnectionsClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HybridConnectionsClient.DeleteAuthorizationRule":
		resp, err = h.dispatchDeleteAuthorizationRule(req)
	case "HybridConnectionsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HybridConnectionsClient.GetAuthorizationRule":
		resp, err = h.dispatchGetAuthorizationRule(req)
	case "HybridConnectionsClient.NewListAuthorizationRulesPager":
		resp, err = h.dispatchNewListAuthorizationRulesPager(req)
	case "HybridConnectionsClient.NewListByNamespacePager":
		resp, err = h.dispatchNewListByNamespacePager(req)
	case "HybridConnectionsClient.ListKeys":
		resp, err = h.dispatchListKeys(req)
	case "HybridConnectionsClient.RegenerateKeys":
		resp, err = h.dispatchRegenerateKeys(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrelay.HybridConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HybridConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchCreateOrUpdateAuthorizationRule(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdateAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrelay.AuthorizationRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdateAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchDeleteAuthorizationRule(req *http.Request) (*http.Response, error) {
	if h.srv.DeleteAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.DeleteAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HybridConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchGetAuthorizationRule(req *http.Request) (*http.Response, error) {
	if h.srv.GetAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.GetAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchNewListAuthorizationRulesPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListAuthorizationRulesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAuthorizationRulesPager not implemented")}
	}
	newListAuthorizationRulesPager := h.newListAuthorizationRulesPager.get(req)
	if newListAuthorizationRulesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListAuthorizationRulesPager(resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, nil)
		newListAuthorizationRulesPager = &resp
		h.newListAuthorizationRulesPager.add(req, newListAuthorizationRulesPager)
		server.PagerResponderInjectNextLinks(newListAuthorizationRulesPager, req, func(page *armrelay.HybridConnectionsClientListAuthorizationRulesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAuthorizationRulesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListAuthorizationRulesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAuthorizationRulesPager) {
		h.newListAuthorizationRulesPager.remove(req)
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := h.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, nil)
		newListByNamespacePager = &resp
		h.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armrelay.HybridConnectionsClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		h.newListByNamespacePager.remove(req)
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if h.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.ListKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HybridConnectionsServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if h.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Relay/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hybridConnections/(?P<hybridConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrelay.RegenerateAccessKeyParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	hybridConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hybridConnectionName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, hybridConnectionNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
