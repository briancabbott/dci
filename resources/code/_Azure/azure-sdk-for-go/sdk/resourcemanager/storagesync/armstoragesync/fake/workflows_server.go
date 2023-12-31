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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
	"net/http"
	"net/url"
	"regexp"
)

// WorkflowsServer is a fake server for instances of the armstoragesync.WorkflowsClient type.
type WorkflowsServer struct {
	// Abort is the fake for method WorkflowsClient.Abort
	// HTTP status codes to indicate success: http.StatusOK
	Abort func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *armstoragesync.WorkflowsClientAbortOptions) (resp azfake.Responder[armstoragesync.WorkflowsClientAbortResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkflowsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, storageSyncServiceName string, workflowID string, options *armstoragesync.WorkflowsClientGetOptions) (resp azfake.Responder[armstoragesync.WorkflowsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByStorageSyncServicePager is the fake for method WorkflowsClient.NewListByStorageSyncServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByStorageSyncServicePager func(resourceGroupName string, storageSyncServiceName string, options *armstoragesync.WorkflowsClientListByStorageSyncServiceOptions) (resp azfake.PagerResponder[armstoragesync.WorkflowsClientListByStorageSyncServiceResponse])
}

// NewWorkflowsServerTransport creates a new instance of WorkflowsServerTransport with the provided implementation.
// The returned WorkflowsServerTransport instance is connected to an instance of armstoragesync.WorkflowsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowsServerTransport(srv *WorkflowsServer) *WorkflowsServerTransport {
	return &WorkflowsServerTransport{
		srv:                              srv,
		newListByStorageSyncServicePager: newTracker[azfake.PagerResponder[armstoragesync.WorkflowsClientListByStorageSyncServiceResponse]](),
	}
}

// WorkflowsServerTransport connects instances of armstoragesync.WorkflowsClient to instances of WorkflowsServer.
// Don't use this type directly, use NewWorkflowsServerTransport instead.
type WorkflowsServerTransport struct {
	srv                              *WorkflowsServer
	newListByStorageSyncServicePager *tracker[azfake.PagerResponder[armstoragesync.WorkflowsClientListByStorageSyncServiceResponse]]
}

// Do implements the policy.Transporter interface for WorkflowsServerTransport.
func (w *WorkflowsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowsClient.Abort":
		resp, err = w.dispatchAbort(req)
	case "WorkflowsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowsClient.NewListByStorageSyncServicePager":
		resp, err = w.dispatchNewListByStorageSyncServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowsServerTransport) dispatchAbort(req *http.Request) (*http.Response, error) {
	if w.srv.Abort == nil {
		return nil, &nonRetriableError{errors.New("fake for method Abort not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workflows/(?P<workflowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/abort`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
	if err != nil {
		return nil, err
	}
	workflowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Abort(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, workflowIDParam, nil)
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
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (w *WorkflowsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workflows/(?P<workflowId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
	if err != nil {
		return nil, err
	}
	workflowIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, storageSyncServiceNameParam, workflowIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Workflow, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (w *WorkflowsServerTransport) dispatchNewListByStorageSyncServicePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByStorageSyncServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByStorageSyncServicePager not implemented")}
	}
	newListByStorageSyncServicePager := w.newListByStorageSyncServicePager.get(req)
	if newListByStorageSyncServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageSync/storageSyncServices/(?P<storageSyncServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workflows`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		storageSyncServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageSyncServiceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByStorageSyncServicePager(resourceGroupNameParam, storageSyncServiceNameParam, nil)
		newListByStorageSyncServicePager = &resp
		w.newListByStorageSyncServicePager.add(req, newListByStorageSyncServicePager)
	}
	resp, err := server.PagerResponderNext(newListByStorageSyncServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByStorageSyncServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByStorageSyncServicePager) {
		w.newListByStorageSyncServicePager.remove(req)
	}
	return resp, nil
}
