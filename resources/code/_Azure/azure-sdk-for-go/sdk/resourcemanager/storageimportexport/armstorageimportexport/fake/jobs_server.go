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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// JobsServer is a fake server for instances of the armstorageimportexport.JobsClient type.
type JobsServer struct {
	// Create is the fake for method JobsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, jobName string, resourceGroupName string, body armstorageimportexport.PutJobParameters, options *armstorageimportexport.JobsClientCreateOptions) (resp azfake.Responder[armstorageimportexport.JobsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method JobsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, jobName string, resourceGroupName string, options *armstorageimportexport.JobsClientDeleteOptions) (resp azfake.Responder[armstorageimportexport.JobsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, jobName string, resourceGroupName string, options *armstorageimportexport.JobsClientGetOptions) (resp azfake.Responder[armstorageimportexport.JobsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method JobsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armstorageimportexport.JobsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armstorageimportexport.JobsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method JobsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armstorageimportexport.JobsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armstorageimportexport.JobsClientListBySubscriptionResponse])

	// Update is the fake for method JobsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, jobName string, resourceGroupName string, body armstorageimportexport.UpdateJobParameters, options *armstorageimportexport.JobsClientUpdateOptions) (resp azfake.Responder[armstorageimportexport.JobsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewJobsServerTransport creates a new instance of JobsServerTransport with the provided implementation.
// The returned JobsServerTransport instance is connected to an instance of armstorageimportexport.JobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobsServerTransport(srv *JobsServer) *JobsServerTransport {
	return &JobsServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armstorageimportexport.JobsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armstorageimportexport.JobsClientListBySubscriptionResponse]](),
	}
}

// JobsServerTransport connects instances of armstorageimportexport.JobsClient to instances of JobsServer.
// Don't use this type directly, use NewJobsServerTransport instead.
type JobsServerTransport struct {
	srv                         *JobsServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armstorageimportexport.JobsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armstorageimportexport.JobsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for JobsServerTransport.
func (j *JobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobsClient.Create":
		resp, err = j.dispatchCreate(req)
	case "JobsClient.Delete":
		resp, err = j.dispatchDelete(req)
	case "JobsClient.Get":
		resp, err = j.dispatchGet(req)
	case "JobsClient.NewListByResourceGroupPager":
		resp, err = j.dispatchNewListByResourceGroupPager(req)
	case "JobsClient.NewListBySubscriptionPager":
		resp, err = j.dispatchNewListBySubscriptionPager(req)
	case "JobsClient.Update":
		resp, err = j.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if j.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorageimportexport.PutJobParameters](req)
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
	clientTenantIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-tenant-id"))
	var options *armstorageimportexport.JobsClientCreateOptions
	if acceptLanguageParam != nil || clientTenantIDParam != nil {
		options = &armstorageimportexport.JobsClientCreateOptions{
			AcceptLanguage: acceptLanguageParam,
			ClientTenantID: clientTenantIDParam,
		}
	}
	respr, errRespr := j.srv.Create(req.Context(), jobNameParam, resourceGroupNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if j.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
	var options *armstorageimportexport.JobsClientDeleteOptions
	if acceptLanguageParam != nil {
		options = &armstorageimportexport.JobsClientDeleteOptions{
			AcceptLanguage: acceptLanguageParam,
		}
	}
	respr, errRespr := j.srv.Delete(req.Context(), jobNameParam, resourceGroupNameParam, options)
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

func (j *JobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
	var options *armstorageimportexport.JobsClientGetOptions
	if acceptLanguageParam != nil {
		options = &armstorageimportexport.JobsClientGetOptions{
			AcceptLanguage: acceptLanguageParam,
		}
	}
	respr, errRespr := j.srv.Get(req.Context(), jobNameParam, resourceGroupNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := j.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
		var options *armstorageimportexport.JobsClientListByResourceGroupOptions
		if topParam != nil || filterParam != nil || acceptLanguageParam != nil {
			options = &armstorageimportexport.JobsClientListByResourceGroupOptions{
				Top:            topParam,
				Filter:         filterParam,
				AcceptLanguage: acceptLanguageParam,
			}
		}
		resp := j.srv.NewListByResourceGroupPager(resourceGroupNameParam, options)
		newListByResourceGroupPager = &resp
		j.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armstorageimportexport.JobsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		j.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := j.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
		var options *armstorageimportexport.JobsClientListBySubscriptionOptions
		if topParam != nil || filterParam != nil || acceptLanguageParam != nil {
			options = &armstorageimportexport.JobsClientListBySubscriptionOptions{
				Top:            topParam,
				Filter:         filterParam,
				AcceptLanguage: acceptLanguageParam,
			}
		}
		resp := j.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		j.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armstorageimportexport.JobsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		j.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstorageimportexport.UpdateJobParameters](req)
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
	var options *armstorageimportexport.JobsClientUpdateOptions
	if acceptLanguageParam != nil {
		options = &armstorageimportexport.JobsClientUpdateOptions{
			AcceptLanguage: acceptLanguageParam,
		}
	}
	respr, errRespr := j.srv.Update(req.Context(), jobNameParam, resourceGroupNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
