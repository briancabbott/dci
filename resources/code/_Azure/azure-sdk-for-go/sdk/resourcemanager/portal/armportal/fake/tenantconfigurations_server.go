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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
	"net/http"
	"net/url"
	"regexp"
)

// TenantConfigurationsServer is a fake server for instances of the armportal.TenantConfigurationsClient type.
type TenantConfigurationsServer struct {
	// Create is the fake for method TenantConfigurationsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, configurationName armportal.ConfigurationName, tenantConfiguration armportal.Configuration, options *armportal.TenantConfigurationsClientCreateOptions) (resp azfake.Responder[armportal.TenantConfigurationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TenantConfigurationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, configurationName armportal.ConfigurationName, options *armportal.TenantConfigurationsClientDeleteOptions) (resp azfake.Responder[armportal.TenantConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TenantConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotFound
	Get func(ctx context.Context, configurationName armportal.ConfigurationName, options *armportal.TenantConfigurationsClientGetOptions) (resp azfake.Responder[armportal.TenantConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TenantConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armportal.TenantConfigurationsClientListOptions) (resp azfake.PagerResponder[armportal.TenantConfigurationsClientListResponse])
}

// NewTenantConfigurationsServerTransport creates a new instance of TenantConfigurationsServerTransport with the provided implementation.
// The returned TenantConfigurationsServerTransport instance is connected to an instance of armportal.TenantConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTenantConfigurationsServerTransport(srv *TenantConfigurationsServer) *TenantConfigurationsServerTransport {
	return &TenantConfigurationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armportal.TenantConfigurationsClientListResponse]](),
	}
}

// TenantConfigurationsServerTransport connects instances of armportal.TenantConfigurationsClient to instances of TenantConfigurationsServer.
// Don't use this type directly, use NewTenantConfigurationsServerTransport instead.
type TenantConfigurationsServerTransport struct {
	srv          *TenantConfigurationsServer
	newListPager *tracker[azfake.PagerResponder[armportal.TenantConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for TenantConfigurationsServerTransport.
func (t *TenantConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TenantConfigurationsClient.Create":
		resp, err = t.dispatchCreate(req)
	case "TenantConfigurationsClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TenantConfigurationsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TenantConfigurationsClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TenantConfigurationsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if t.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Portal/tenantConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armportal.Configuration](req)
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armportal.ConfigurationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armportal.ConfigurationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Create(req.Context(), configurationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Configuration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantConfigurationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Portal/tenantConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armportal.ConfigurationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armportal.ConfigurationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), configurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Portal/tenantConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	configurationNameParam, err := parseWithCast(matches[regex.SubexpIndex("configurationName")], func(v string) (armportal.ConfigurationName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armportal.ConfigurationName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), configurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Configuration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TenantConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		resp := t.srv.NewListPager(nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armportal.TenantConfigurationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}
