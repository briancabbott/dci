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
)

// DscCompilationJobServer is a fake server for instances of the armautomation.DscCompilationJobClient type.
type DscCompilationJobServer struct {
	// BeginCreate is the fake for method DscCompilationJobClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, parameters armautomation.DscCompilationJobCreateParameters, options *armautomation.DscCompilationJobClientBeginCreateOptions) (resp azfake.PollerResponder[armautomation.DscCompilationJobClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DscCompilationJobClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, compilationJobName string, options *armautomation.DscCompilationJobClientGetOptions) (resp azfake.Responder[armautomation.DscCompilationJobClientGetResponse], errResp azfake.ErrorResponder)

	// GetStream is the fake for method DscCompilationJobClient.GetStream
	// HTTP status codes to indicate success: http.StatusOK
	GetStream func(ctx context.Context, resourceGroupName string, automationAccountName string, jobID string, jobStreamID string, options *armautomation.DscCompilationJobClientGetStreamOptions) (resp azfake.Responder[armautomation.DscCompilationJobClientGetStreamResponse], errResp azfake.ErrorResponder)

	// NewListByAutomationAccountPager is the fake for method DscCompilationJobClient.NewListByAutomationAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAutomationAccountPager func(resourceGroupName string, automationAccountName string, options *armautomation.DscCompilationJobClientListByAutomationAccountOptions) (resp azfake.PagerResponder[armautomation.DscCompilationJobClientListByAutomationAccountResponse])
}

// NewDscCompilationJobServerTransport creates a new instance of DscCompilationJobServerTransport with the provided implementation.
// The returned DscCompilationJobServerTransport instance is connected to an instance of armautomation.DscCompilationJobClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDscCompilationJobServerTransport(srv *DscCompilationJobServer) *DscCompilationJobServerTransport {
	return &DscCompilationJobServerTransport{
		srv:                             srv,
		beginCreate:                     newTracker[azfake.PollerResponder[armautomation.DscCompilationJobClientCreateResponse]](),
		newListByAutomationAccountPager: newTracker[azfake.PagerResponder[armautomation.DscCompilationJobClientListByAutomationAccountResponse]](),
	}
}

// DscCompilationJobServerTransport connects instances of armautomation.DscCompilationJobClient to instances of DscCompilationJobServer.
// Don't use this type directly, use NewDscCompilationJobServerTransport instead.
type DscCompilationJobServerTransport struct {
	srv                             *DscCompilationJobServer
	beginCreate                     *tracker[azfake.PollerResponder[armautomation.DscCompilationJobClientCreateResponse]]
	newListByAutomationAccountPager *tracker[azfake.PagerResponder[armautomation.DscCompilationJobClientListByAutomationAccountResponse]]
}

// Do implements the policy.Transporter interface for DscCompilationJobServerTransport.
func (d *DscCompilationJobServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DscCompilationJobClient.BeginCreate":
		resp, err = d.dispatchBeginCreate(req)
	case "DscCompilationJobClient.Get":
		resp, err = d.dispatchGet(req)
	case "DscCompilationJobClient.GetStream":
		resp, err = d.dispatchGetStream(req)
	case "DscCompilationJobClient.NewListByAutomationAccountPager":
		resp, err = d.dispatchNewListByAutomationAccountPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DscCompilationJobServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := d.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/compilationjobs/(?P<compilationJobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armautomation.DscCompilationJobCreateParameters](req)
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
		compilationJobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("compilationJobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreate(req.Context(), resourceGroupNameParam, automationAccountNameParam, compilationJobNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		d.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusCreated}, resp.StatusCode) {
		d.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		d.beginCreate.remove(req)
	}

	return resp, nil
}

func (d *DscCompilationJobServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/compilationjobs/(?P<compilationJobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	compilationJobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("compilationJobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, compilationJobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DscCompilationJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DscCompilationJobServerTransport) dispatchGetStream(req *http.Request) (*http.Response, error) {
	if d.srv.GetStream == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetStream not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/compilationjobs/(?P<jobId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/streams/(?P<jobStreamId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	jobIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobId")])
	if err != nil {
		return nil, err
	}
	jobStreamIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobStreamId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetStream(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobIDParam, jobStreamIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobStream, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DscCompilationJobServerTransport) dispatchNewListByAutomationAccountPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByAutomationAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAutomationAccountPager not implemented")}
	}
	newListByAutomationAccountPager := d.newListByAutomationAccountPager.get(req)
	if newListByAutomationAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/compilationjobs`
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
		var options *armautomation.DscCompilationJobClientListByAutomationAccountOptions
		if filterParam != nil {
			options = &armautomation.DscCompilationJobClientListByAutomationAccountOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListByAutomationAccountPager(resourceGroupNameParam, automationAccountNameParam, options)
		newListByAutomationAccountPager = &resp
		d.newListByAutomationAccountPager.add(req, newListByAutomationAccountPager)
		server.PagerResponderInjectNextLinks(newListByAutomationAccountPager, req, func(page *armautomation.DscCompilationJobClientListByAutomationAccountResponse, createLink func() string) {
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
