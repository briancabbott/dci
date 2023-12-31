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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationGroupServer is a fake server for instances of the armeventhub.ApplicationGroupClient type.
type ApplicationGroupServer struct {
	// CreateOrUpdateApplicationGroup is the fake for method ApplicationGroupClient.CreateOrUpdateApplicationGroup
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateApplicationGroup func(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string, parameters armeventhub.ApplicationGroup, options *armeventhub.ApplicationGroupClientCreateOrUpdateApplicationGroupOptions) (resp azfake.Responder[armeventhub.ApplicationGroupClientCreateOrUpdateApplicationGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ApplicationGroupClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string, options *armeventhub.ApplicationGroupClientDeleteOptions) (resp azfake.Responder[armeventhub.ApplicationGroupClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationGroupClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, applicationGroupName string, options *armeventhub.ApplicationGroupClientGetOptions) (resp azfake.Responder[armeventhub.ApplicationGroupClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNamespacePager is the fake for method ApplicationGroupClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armeventhub.ApplicationGroupClientListByNamespaceOptions) (resp azfake.PagerResponder[armeventhub.ApplicationGroupClientListByNamespaceResponse])
}

// NewApplicationGroupServerTransport creates a new instance of ApplicationGroupServerTransport with the provided implementation.
// The returned ApplicationGroupServerTransport instance is connected to an instance of armeventhub.ApplicationGroupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationGroupServerTransport(srv *ApplicationGroupServer) *ApplicationGroupServerTransport {
	return &ApplicationGroupServerTransport{
		srv:                     srv,
		newListByNamespacePager: newTracker[azfake.PagerResponder[armeventhub.ApplicationGroupClientListByNamespaceResponse]](),
	}
}

// ApplicationGroupServerTransport connects instances of armeventhub.ApplicationGroupClient to instances of ApplicationGroupServer.
// Don't use this type directly, use NewApplicationGroupServerTransport instead.
type ApplicationGroupServerTransport struct {
	srv                     *ApplicationGroupServer
	newListByNamespacePager *tracker[azfake.PagerResponder[armeventhub.ApplicationGroupClientListByNamespaceResponse]]
}

// Do implements the policy.Transporter interface for ApplicationGroupServerTransport.
func (a *ApplicationGroupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationGroupClient.CreateOrUpdateApplicationGroup":
		resp, err = a.dispatchCreateOrUpdateApplicationGroup(req)
	case "ApplicationGroupClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ApplicationGroupClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationGroupClient.NewListByNamespacePager":
		resp, err = a.dispatchNewListByNamespacePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationGroupServerTransport) dispatchCreateOrUpdateApplicationGroup(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdateApplicationGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateApplicationGroup not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGroups/(?P<applicationGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.ApplicationGroup](req)
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
	applicationGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdateApplicationGroup(req.Context(), resourceGroupNameParam, namespaceNameParam, applicationGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGroupServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGroups/(?P<applicationGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	applicationGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, applicationGroupNameParam, nil)
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

func (a *ApplicationGroupServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGroups/(?P<applicationGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	applicationGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, applicationGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ApplicationGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationGroupServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := a.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationGroups`
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
		resp := a.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, nil)
		newListByNamespacePager = &resp
		a.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armeventhub.ApplicationGroupClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		a.newListByNamespacePager.remove(req)
	}
	return resp, nil
}
