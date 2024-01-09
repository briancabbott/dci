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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DscNodeServer is a fake server for instances of the armautomation.DscNodeClient type.
type DscNodeServer struct {
	// Delete is the fake for method DscNodeClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *armautomation.DscNodeClientDeleteOptions) (resp azfake.Responder[armautomation.DscNodeClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DscNodeClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *armautomation.DscNodeClientGetOptions) (resp azfake.Responder[armautomation.DscNodeClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAutomationAccountPager is the fake for method DscNodeClient.NewListByAutomationAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAutomationAccountPager func(resourceGroupName string, automationAccountName string, options *armautomation.DscNodeClientListByAutomationAccountOptions) (resp azfake.PagerResponder[armautomation.DscNodeClientListByAutomationAccountResponse])

	// Update is the fake for method DscNodeClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters armautomation.DscNodeUpdateParameters, options *armautomation.DscNodeClientUpdateOptions) (resp azfake.Responder[armautomation.DscNodeClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDscNodeServerTransport creates a new instance of DscNodeServerTransport with the provided implementation.
// The returned DscNodeServerTransport instance is connected to an instance of armautomation.DscNodeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDscNodeServerTransport(srv *DscNodeServer) *DscNodeServerTransport {
	return &DscNodeServerTransport{
		srv:                             srv,
		newListByAutomationAccountPager: newTracker[azfake.PagerResponder[armautomation.DscNodeClientListByAutomationAccountResponse]](),
	}
}

// DscNodeServerTransport connects instances of armautomation.DscNodeClient to instances of DscNodeServer.
// Don't use this type directly, use NewDscNodeServerTransport instead.
type DscNodeServerTransport struct {
	srv                             *DscNodeServer
	newListByAutomationAccountPager *tracker[azfake.PagerResponder[armautomation.DscNodeClientListByAutomationAccountResponse]]
}

// Do implements the policy.Transporter interface for DscNodeServerTransport.
func (d *DscNodeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DscNodeClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DscNodeClient.Get":
		resp, err = d.dispatchGet(req)
	case "DscNodeClient.NewListByAutomationAccountPager":
		resp, err = d.dispatchNewListByAutomationAccountPager(req)
	case "DscNodeClient.Update":
		resp, err = d.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DscNodeServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodes/(?P<nodeId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	nodeIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("nodeId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, automationAccountNameParam, nodeIDParam, nil)
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

func (d *DscNodeServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodes/(?P<nodeId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	nodeIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("nodeId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, nodeIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DscNode, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DscNodeServerTransport) dispatchNewListByAutomationAccountPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByAutomationAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAutomationAccountPager not implemented")}
	}
	newListByAutomationAccountPager := d.newListByAutomationAccountPager.get(req)
	if newListByAutomationAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodes`
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
		automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
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
		inlinecountUnescaped, err := url.QueryUnescape(qp.Get("$inlinecount"))
		if err != nil {
			return nil, err
		}
		inlinecountParam := getOptional(inlinecountUnescaped)
		var options *armautomation.DscNodeClientListByAutomationAccountOptions
		if filterParam != nil || skipParam != nil || topParam != nil || inlinecountParam != nil {
			options = &armautomation.DscNodeClientListByAutomationAccountOptions{
				Filter:      filterParam,
				Skip:        skipParam,
				Top:         topParam,
				Inlinecount: inlinecountParam,
			}
		}
		resp := d.srv.NewListByAutomationAccountPager(resourceGroupNameParam, automationAccountNameParam, options)
		newListByAutomationAccountPager = &resp
		d.newListByAutomationAccountPager.add(req, newListByAutomationAccountPager)
		server.PagerResponderInjectNextLinks(newListByAutomationAccountPager, req, func(page *armautomation.DscNodeClientListByAutomationAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAutomationAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByAutomationAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAutomationAccountPager) {
		d.newListByAutomationAccountPager.remove(req)
	}
	return resp, nil
}

func (d *DscNodeServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodes/(?P<nodeId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomation.DscNodeUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	nodeIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("nodeId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameParam, automationAccountNameParam, nodeIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DscNode, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
