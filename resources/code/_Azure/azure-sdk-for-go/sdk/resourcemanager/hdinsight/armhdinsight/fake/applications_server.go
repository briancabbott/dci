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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
	"net/http"
	"net/url"
	"regexp"
)

// ApplicationsServer is a fake server for instances of the armhdinsight.ApplicationsClient type.
type ApplicationsServer struct {
	// BeginCreate is the fake for method ApplicationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK
	BeginCreate func(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, parameters armhdinsight.Application, options *armhdinsight.ApplicationsClientBeginCreateOptions) (resp azfake.PollerResponder[armhdinsight.ApplicationsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ApplicationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, options *armhdinsight.ApplicationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhdinsight.ApplicationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, options *armhdinsight.ApplicationsClientGetOptions) (resp azfake.Responder[armhdinsight.ApplicationsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAzureAsyncOperationStatus is the fake for method ApplicationsClient.GetAzureAsyncOperationStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetAzureAsyncOperationStatus func(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, operationID string, options *armhdinsight.ApplicationsClientGetAzureAsyncOperationStatusOptions) (resp azfake.Responder[armhdinsight.ApplicationsClientGetAzureAsyncOperationStatusResponse], errResp azfake.ErrorResponder)

	// NewListByClusterPager is the fake for method ApplicationsClient.NewListByClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByClusterPager func(resourceGroupName string, clusterName string, options *armhdinsight.ApplicationsClientListByClusterOptions) (resp azfake.PagerResponder[armhdinsight.ApplicationsClientListByClusterResponse])
}

// NewApplicationsServerTransport creates a new instance of ApplicationsServerTransport with the provided implementation.
// The returned ApplicationsServerTransport instance is connected to an instance of armhdinsight.ApplicationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationsServerTransport(srv *ApplicationsServer) *ApplicationsServerTransport {
	return &ApplicationsServerTransport{
		srv:                   srv,
		beginCreate:           newTracker[azfake.PollerResponder[armhdinsight.ApplicationsClientCreateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armhdinsight.ApplicationsClientDeleteResponse]](),
		newListByClusterPager: newTracker[azfake.PagerResponder[armhdinsight.ApplicationsClientListByClusterResponse]](),
	}
}

// ApplicationsServerTransport connects instances of armhdinsight.ApplicationsClient to instances of ApplicationsServer.
// Don't use this type directly, use NewApplicationsServerTransport instead.
type ApplicationsServerTransport struct {
	srv                   *ApplicationsServer
	beginCreate           *tracker[azfake.PollerResponder[armhdinsight.ApplicationsClientCreateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armhdinsight.ApplicationsClientDeleteResponse]]
	newListByClusterPager *tracker[azfake.PagerResponder[armhdinsight.ApplicationsClientListByClusterResponse]]
}

// Do implements the policy.Transporter interface for ApplicationsServerTransport.
func (a *ApplicationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "ApplicationsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "ApplicationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationsClient.GetAzureAsyncOperationStatus":
		resp, err = a.dispatchGetAzureAsyncOperationStatus(req)
	case "ApplicationsClient.NewListByClusterPager":
		resp, err = a.dispatchNewListByClusterPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := a.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhdinsight.Application](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreate(req.Context(), resourceGroupNameParam, clusterNameParam, applicationNameParam, body, nil)
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

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		a.beginCreate.remove(req)
	}

	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, applicationNameParam, nil)
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

func (a *ApplicationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, applicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Application, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchGetAzureAsyncOperationStatus(req *http.Request) (*http.Response, error) {
	if a.srv.GetAzureAsyncOperationStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAzureAsyncOperationStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/azureasyncoperations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetAzureAsyncOperationStatus(req.Context(), resourceGroupNameParam, clusterNameParam, applicationNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AsyncOperationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchNewListByClusterPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByClusterPager not implemented")}
	}
	newListByClusterPager := a.newListByClusterPager.get(req)
	if newListByClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HDInsight/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByClusterPager(resourceGroupNameParam, clusterNameParam, nil)
		newListByClusterPager = &resp
		a.newListByClusterPager.add(req, newListByClusterPager)
		server.PagerResponderInjectNextLinks(newListByClusterPager, req, func(page *armhdinsight.ApplicationsClientListByClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByClusterPager) {
		a.newListByClusterPager.remove(req)
	}
	return resp, nil
}
