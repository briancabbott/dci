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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"net/http"
	"net/url"
	"regexp"
)

// LinksServer is a fake server for instances of the armcustomerinsights.LinksClient type.
type LinksServer struct {
	// BeginCreateOrUpdate is the fake for method LinksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hubName string, linkName string, parameters armcustomerinsights.LinkResourceFormat, options *armcustomerinsights.LinksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcustomerinsights.LinksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method LinksClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Delete func(ctx context.Context, resourceGroupName string, hubName string, linkName string, options *armcustomerinsights.LinksClientDeleteOptions) (resp azfake.Responder[armcustomerinsights.LinksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LinksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hubName string, linkName string, options *armcustomerinsights.LinksClientGetOptions) (resp azfake.Responder[armcustomerinsights.LinksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHubPager is the fake for method LinksClient.NewListByHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHubPager func(resourceGroupName string, hubName string, options *armcustomerinsights.LinksClientListByHubOptions) (resp azfake.PagerResponder[armcustomerinsights.LinksClientListByHubResponse])
}

// NewLinksServerTransport creates a new instance of LinksServerTransport with the provided implementation.
// The returned LinksServerTransport instance is connected to an instance of armcustomerinsights.LinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLinksServerTransport(srv *LinksServer) *LinksServerTransport {
	return &LinksServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcustomerinsights.LinksClientCreateOrUpdateResponse]](),
		newListByHubPager:   newTracker[azfake.PagerResponder[armcustomerinsights.LinksClientListByHubResponse]](),
	}
}

// LinksServerTransport connects instances of armcustomerinsights.LinksClient to instances of LinksServer.
// Don't use this type directly, use NewLinksServerTransport instead.
type LinksServerTransport struct {
	srv                 *LinksServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcustomerinsights.LinksClientCreateOrUpdateResponse]]
	newListByHubPager   *tracker[azfake.PagerResponder[armcustomerinsights.LinksClientListByHubResponse]]
}

// Do implements the policy.Transporter interface for LinksServerTransport.
func (l *LinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LinksClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LinksClient.Delete":
		resp, err = l.dispatchDelete(req)
	case "LinksClient.Get":
		resp, err = l.dispatchGet(req)
	case "LinksClient.NewListByHubPager":
		resp, err = l.dispatchNewListByHubPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LinksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links/(?P<linkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.LinkResourceFormat](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		linkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, hubNameParam, linkNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LinksServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if l.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links/(?P<linkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	linkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Delete(req.Context(), resourceGroupNameParam, hubNameParam, linkNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links/(?P<linkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	linkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, hubNameParam, linkNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LinkResourceFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LinksServerTransport) dispatchNewListByHubPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHubPager not implemented")}
	}
	newListByHubPager := l.newListByHubPager.get(req)
	if newListByHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByHubPager(resourceGroupNameParam, hubNameParam, nil)
		newListByHubPager = &resp
		l.newListByHubPager.add(req, newListByHubPager)
		server.PagerResponderInjectNextLinks(newListByHubPager, req, func(page *armcustomerinsights.LinksClientListByHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHubPager) {
		l.newListByHubPager.remove(req)
	}
	return resp, nil
}
