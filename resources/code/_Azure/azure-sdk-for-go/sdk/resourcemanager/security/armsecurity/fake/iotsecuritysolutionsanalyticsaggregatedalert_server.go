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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// IotSecuritySolutionsAnalyticsAggregatedAlertServer is a fake server for instances of the armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClient type.
type IotSecuritySolutionsAnalyticsAggregatedAlertServer struct {
	// Dismiss is the fake for method IotSecuritySolutionsAnalyticsAggregatedAlertClient.Dismiss
	// HTTP status codes to indicate success: http.StatusOK
	Dismiss func(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissOptions) (resp azfake.Responder[armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientDismissResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IotSecuritySolutionsAnalyticsAggregatedAlertClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, solutionName string, aggregatedAlertName string, options *armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientGetOptions) (resp azfake.Responder[armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IotSecuritySolutionsAnalyticsAggregatedAlertClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, solutionName string, options *armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions) (resp azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse])
}

// NewIotSecuritySolutionsAnalyticsAggregatedAlertServerTransport creates a new instance of IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport with the provided implementation.
// The returned IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport instance is connected to an instance of armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIotSecuritySolutionsAnalyticsAggregatedAlertServerTransport(srv *IotSecuritySolutionsAnalyticsAggregatedAlertServer) *IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport {
	return &IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse]](),
	}
}

// IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport connects instances of armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClient to instances of IotSecuritySolutionsAnalyticsAggregatedAlertServer.
// Don't use this type directly, use NewIotSecuritySolutionsAnalyticsAggregatedAlertServerTransport instead.
type IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport struct {
	srv          *IotSecuritySolutionsAnalyticsAggregatedAlertServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse]]
}

// Do implements the policy.Transporter interface for IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport.
func (i *IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IotSecuritySolutionsAnalyticsAggregatedAlertClient.Dismiss":
		resp, err = i.dispatchDismiss(req)
	case "IotSecuritySolutionsAnalyticsAggregatedAlertClient.Get":
		resp, err = i.dispatchGet(req)
	case "IotSecuritySolutionsAnalyticsAggregatedAlertClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport) dispatchDismiss(req *http.Request) (*http.Response, error) {
	if i.srv.Dismiss == nil {
		return nil, &nonRetriableError{errors.New("fake for method Dismiss not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/iotSecuritySolutions/(?P<solutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analyticsModels/default/aggregatedAlerts/(?P<aggregatedAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dismiss`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	solutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionName")])
	if err != nil {
		return nil, err
	}
	aggregatedAlertNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("aggregatedAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Dismiss(req.Context(), resourceGroupNameParam, solutionNameParam, aggregatedAlertNameParam, nil)
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

func (i *IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/iotSecuritySolutions/(?P<solutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analyticsModels/default/aggregatedAlerts/(?P<aggregatedAlertName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	solutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionName")])
	if err != nil {
		return nil, err
	}
	aggregatedAlertNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("aggregatedAlertName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, solutionNameParam, aggregatedAlertNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IoTSecurityAggregatedAlert, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IotSecuritySolutionsAnalyticsAggregatedAlertServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/iotSecuritySolutions/(?P<solutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/analyticsModels/default/aggregatedAlerts`
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
		solutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("solutionName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions
		if topParam != nil {
			options = &armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListOptions{
				Top: topParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, solutionNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.IotSecuritySolutionsAnalyticsAggregatedAlertClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}
