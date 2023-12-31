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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"
	"net/http"
	"net/url"
	"regexp"
)

// ManagementAssociationsServer is a fake server for instances of the armoperationsmanagement.ManagementAssociationsClient type.
type ManagementAssociationsServer struct {
	// CreateOrUpdate is the fake for method ManagementAssociationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, parameters armoperationsmanagement.ManagementAssociation, options *armoperationsmanagement.ManagementAssociationsClientCreateOrUpdateOptions) (resp azfake.Responder[armoperationsmanagement.ManagementAssociationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagementAssociationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *armoperationsmanagement.ManagementAssociationsClientDeleteOptions) (resp azfake.Responder[armoperationsmanagement.ManagementAssociationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagementAssociationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *armoperationsmanagement.ManagementAssociationsClientGetOptions) (resp azfake.Responder[armoperationsmanagement.ManagementAssociationsClientGetResponse], errResp azfake.ErrorResponder)

	// ListBySubscription is the fake for method ManagementAssociationsClient.ListBySubscription
	// HTTP status codes to indicate success: http.StatusOK
	ListBySubscription func(ctx context.Context, options *armoperationsmanagement.ManagementAssociationsClientListBySubscriptionOptions) (resp azfake.Responder[armoperationsmanagement.ManagementAssociationsClientListBySubscriptionResponse], errResp azfake.ErrorResponder)
}

// NewManagementAssociationsServerTransport creates a new instance of ManagementAssociationsServerTransport with the provided implementation.
// The returned ManagementAssociationsServerTransport instance is connected to an instance of armoperationsmanagement.ManagementAssociationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagementAssociationsServerTransport(srv *ManagementAssociationsServer) *ManagementAssociationsServerTransport {
	return &ManagementAssociationsServerTransport{srv: srv}
}

// ManagementAssociationsServerTransport connects instances of armoperationsmanagement.ManagementAssociationsClient to instances of ManagementAssociationsServer.
// Don't use this type directly, use NewManagementAssociationsServerTransport instead.
type ManagementAssociationsServerTransport struct {
	srv *ManagementAssociationsServer
}

// Do implements the policy.Transporter interface for ManagementAssociationsServerTransport.
func (m *ManagementAssociationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagementAssociationsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagementAssociationsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagementAssociationsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagementAssociationsClient.ListBySubscription":
		resp, err = m.dispatchListBySubscription(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagementAssociationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationsManagement/ManagementAssociations/(?P<managementAssociationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationsmanagement.ManagementAssociation](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
	if err != nil {
		return nil, err
	}
	resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	managementAssociationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementAssociationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, providerNameParam, resourceTypeParam, resourceNameParam, managementAssociationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagementAssociation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementAssociationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationsManagement/ManagementAssociations/(?P<managementAssociationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
	if err != nil {
		return nil, err
	}
	resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	managementAssociationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementAssociationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, providerNameParam, resourceTypeParam, resourceNameParam, managementAssociationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementAssociationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<providerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationsManagement/ManagementAssociations/(?P<managementAssociationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	providerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("providerName")])
	if err != nil {
		return nil, err
	}
	resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	managementAssociationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementAssociationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, providerNameParam, resourceTypeParam, resourceNameParam, managementAssociationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagementAssociation, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagementAssociationsServerTransport) dispatchListBySubscription(req *http.Request) (*http.Response, error) {
	if m.srv.ListBySubscription == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListBySubscription not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationsManagement/ManagementAssociations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := m.srv.ListBySubscription(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagementAssociationPropertiesList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
