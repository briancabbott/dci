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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ManagedDatabaseMoveOperationsServer is a fake server for instances of the armsql.ManagedDatabaseMoveOperationsClient type.
type ManagedDatabaseMoveOperationsServer struct {
	// Get is the fake for method ManagedDatabaseMoveOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, locationName string, operationID string, options *armsql.ManagedDatabaseMoveOperationsClientGetOptions) (resp azfake.Responder[armsql.ManagedDatabaseMoveOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByLocationPager is the fake for method ManagedDatabaseMoveOperationsClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(resourceGroupName string, locationName string, options *armsql.ManagedDatabaseMoveOperationsClientListByLocationOptions) (resp azfake.PagerResponder[armsql.ManagedDatabaseMoveOperationsClientListByLocationResponse])
}

// NewManagedDatabaseMoveOperationsServerTransport creates a new instance of ManagedDatabaseMoveOperationsServerTransport with the provided implementation.
// The returned ManagedDatabaseMoveOperationsServerTransport instance is connected to an instance of armsql.ManagedDatabaseMoveOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedDatabaseMoveOperationsServerTransport(srv *ManagedDatabaseMoveOperationsServer) *ManagedDatabaseMoveOperationsServerTransport {
	return &ManagedDatabaseMoveOperationsServerTransport{
		srv:                    srv,
		newListByLocationPager: newTracker[azfake.PagerResponder[armsql.ManagedDatabaseMoveOperationsClientListByLocationResponse]](),
	}
}

// ManagedDatabaseMoveOperationsServerTransport connects instances of armsql.ManagedDatabaseMoveOperationsClient to instances of ManagedDatabaseMoveOperationsServer.
// Don't use this type directly, use NewManagedDatabaseMoveOperationsServerTransport instead.
type ManagedDatabaseMoveOperationsServerTransport struct {
	srv                    *ManagedDatabaseMoveOperationsServer
	newListByLocationPager *tracker[azfake.PagerResponder[armsql.ManagedDatabaseMoveOperationsClientListByLocationResponse]]
}

// Do implements the policy.Transporter interface for ManagedDatabaseMoveOperationsServerTransport.
func (m *ManagedDatabaseMoveOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedDatabaseMoveOperationsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedDatabaseMoveOperationsClient.NewListByLocationPager":
		resp, err = m.dispatchNewListByLocationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedDatabaseMoveOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedDatabaseMoveOperationResults/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, locationNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedDatabaseMoveOperationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedDatabaseMoveOperationsServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := m.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/managedDatabaseMoveOperationResults`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		onlyLatestPerDatabaseUnescaped, err := url.QueryUnescape(qp.Get("onlyLatestPerDatabase"))
		if err != nil {
			return nil, err
		}
		onlyLatestPerDatabaseParam, err := parseOptional(onlyLatestPerDatabaseUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsql.ManagedDatabaseMoveOperationsClientListByLocationOptions
		if onlyLatestPerDatabaseParam != nil || filterParam != nil {
			options = &armsql.ManagedDatabaseMoveOperationsClientListByLocationOptions{
				OnlyLatestPerDatabase: onlyLatestPerDatabaseParam,
				Filter:                filterParam,
			}
		}
		resp := m.srv.NewListByLocationPager(resourceGroupNameParam, locationNameParam, options)
		newListByLocationPager = &resp
		m.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armsql.ManagedDatabaseMoveOperationsClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		m.newListByLocationPager.remove(req)
	}
	return resp, nil
}
