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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// GrantsServer is a fake server for instances of the armeducation.GrantsClient type.
type GrantsServer struct {
	// Get is the fake for method GrantsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, billingAccountName string, billingProfileName string, options *armeducation.GrantsClientGetOptions) (resp azfake.Responder[armeducation.GrantsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method GrantsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(billingAccountName string, billingProfileName string, options *armeducation.GrantsClientListOptions) (resp azfake.PagerResponder[armeducation.GrantsClientListResponse])

	// NewListAllPager is the fake for method GrantsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armeducation.GrantsClientListAllOptions) (resp azfake.PagerResponder[armeducation.GrantsClientListAllResponse])
}

// NewGrantsServerTransport creates a new instance of GrantsServerTransport with the provided implementation.
// The returned GrantsServerTransport instance is connected to an instance of armeducation.GrantsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGrantsServerTransport(srv *GrantsServer) *GrantsServerTransport {
	return &GrantsServerTransport{
		srv:             srv,
		newListPager:    newTracker[azfake.PagerResponder[armeducation.GrantsClientListResponse]](),
		newListAllPager: newTracker[azfake.PagerResponder[armeducation.GrantsClientListAllResponse]](),
	}
}

// GrantsServerTransport connects instances of armeducation.GrantsClient to instances of GrantsServer.
// Don't use this type directly, use NewGrantsServerTransport instead.
type GrantsServerTransport struct {
	srv             *GrantsServer
	newListPager    *tracker[azfake.PagerResponder[armeducation.GrantsClientListResponse]]
	newListAllPager *tracker[azfake.PagerResponder[armeducation.GrantsClientListAllResponse]]
}

// Do implements the policy.Transporter interface for GrantsServerTransport.
func (g *GrantsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GrantsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GrantsClient.NewListPager":
		resp, err = g.dispatchNewListPager(req)
	case "GrantsClient.NewListAllPager":
		resp, err = g.dispatchNewListAllPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GrantsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Education/grants/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
	if err != nil {
		return nil, err
	}
	billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
	if err != nil {
		return nil, err
	}
	includeAllocatedBudgetUnescaped, err := url.QueryUnescape(qp.Get("includeAllocatedBudget"))
	if err != nil {
		return nil, err
	}
	includeAllocatedBudgetParam, err := parseOptional(includeAllocatedBudgetUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armeducation.GrantsClientGetOptions
	if includeAllocatedBudgetParam != nil {
		options = &armeducation.GrantsClientGetOptions{
			IncludeAllocatedBudget: includeAllocatedBudgetParam,
		}
	}
	respr, errRespr := g.srv.Get(req.Context(), billingAccountNameParam, billingProfileNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GrantDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GrantsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingProfiles/(?P<billingProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Education/grants`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		billingAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountName")])
		if err != nil {
			return nil, err
		}
		billingProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingProfileName")])
		if err != nil {
			return nil, err
		}
		includeAllocatedBudgetUnescaped, err := url.QueryUnescape(qp.Get("includeAllocatedBudget"))
		if err != nil {
			return nil, err
		}
		includeAllocatedBudgetParam, err := parseOptional(includeAllocatedBudgetUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armeducation.GrantsClientListOptions
		if includeAllocatedBudgetParam != nil {
			options = &armeducation.GrantsClientListOptions{
				IncludeAllocatedBudget: includeAllocatedBudgetParam,
			}
		}
		resp := g.srv.NewListPager(billingAccountNameParam, billingProfileNameParam, options)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armeducation.GrantsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}

func (g *GrantsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := g.newListAllPager.get(req)
	if newListAllPager == nil {
		qp := req.URL.Query()
		includeAllocatedBudgetUnescaped, err := url.QueryUnescape(qp.Get("includeAllocatedBudget"))
		if err != nil {
			return nil, err
		}
		includeAllocatedBudgetParam, err := parseOptional(includeAllocatedBudgetUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armeducation.GrantsClientListAllOptions
		if includeAllocatedBudgetParam != nil {
			options = &armeducation.GrantsClientListAllOptions{
				IncludeAllocatedBudget: includeAllocatedBudgetParam,
			}
		}
		resp := g.srv.NewListAllPager(options)
		newListAllPager = &resp
		g.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armeducation.GrantsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		g.newListAllPager.remove(req)
	}
	return resp, nil
}
