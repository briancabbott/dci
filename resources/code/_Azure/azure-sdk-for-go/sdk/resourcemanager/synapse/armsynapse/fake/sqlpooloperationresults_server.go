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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// SQLPoolOperationResultsServer is a fake server for instances of the armsynapse.SQLPoolOperationResultsClient type.
type SQLPoolOperationResultsServer struct {
	// BeginGetLocationHeaderResult is the fake for method SQLPoolOperationResultsClient.BeginGetLocationHeaderResult
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetLocationHeaderResult func(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, operationID string, options *armsynapse.SQLPoolOperationResultsClientBeginGetLocationHeaderResultOptions) (resp azfake.PollerResponder[armsynapse.SQLPoolOperationResultsClientGetLocationHeaderResultResponse], errResp azfake.ErrorResponder)
}

// NewSQLPoolOperationResultsServerTransport creates a new instance of SQLPoolOperationResultsServerTransport with the provided implementation.
// The returned SQLPoolOperationResultsServerTransport instance is connected to an instance of armsynapse.SQLPoolOperationResultsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSQLPoolOperationResultsServerTransport(srv *SQLPoolOperationResultsServer) *SQLPoolOperationResultsServerTransport {
	return &SQLPoolOperationResultsServerTransport{
		srv:                          srv,
		beginGetLocationHeaderResult: newTracker[azfake.PollerResponder[armsynapse.SQLPoolOperationResultsClientGetLocationHeaderResultResponse]](),
	}
}

// SQLPoolOperationResultsServerTransport connects instances of armsynapse.SQLPoolOperationResultsClient to instances of SQLPoolOperationResultsServer.
// Don't use this type directly, use NewSQLPoolOperationResultsServerTransport instead.
type SQLPoolOperationResultsServerTransport struct {
	srv                          *SQLPoolOperationResultsServer
	beginGetLocationHeaderResult *tracker[azfake.PollerResponder[armsynapse.SQLPoolOperationResultsClientGetLocationHeaderResultResponse]]
}

// Do implements the policy.Transporter interface for SQLPoolOperationResultsServerTransport.
func (s *SQLPoolOperationResultsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SQLPoolOperationResultsClient.BeginGetLocationHeaderResult":
		resp, err = s.dispatchBeginGetLocationHeaderResult(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SQLPoolOperationResultsServerTransport) dispatchBeginGetLocationHeaderResult(req *http.Request) (*http.Response, error) {
	if s.srv.BeginGetLocationHeaderResult == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetLocationHeaderResult not implemented")}
	}
	beginGetLocationHeaderResult := s.beginGetLocationHeaderResult.get(req)
	if beginGetLocationHeaderResult == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sqlPools/(?P<sqlPoolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		sqlPoolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlPoolName")])
		if err != nil {
			return nil, err
		}
		operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginGetLocationHeaderResult(req.Context(), resourceGroupNameParam, workspaceNameParam, sqlPoolNameParam, operationIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetLocationHeaderResult = &respr
		s.beginGetLocationHeaderResult.add(req, beginGetLocationHeaderResult)
	}

	resp, err := server.PollerResponderNext(beginGetLocationHeaderResult, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginGetLocationHeaderResult.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetLocationHeaderResult) {
		s.beginGetLocationHeaderResult.remove(req)
	}

	return resp, nil
}
