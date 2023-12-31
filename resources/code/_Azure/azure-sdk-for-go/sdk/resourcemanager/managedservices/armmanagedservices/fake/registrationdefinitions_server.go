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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
	"net/http"
	"net/url"
	"regexp"
)

// RegistrationDefinitionsServer is a fake server for instances of the armmanagedservices.RegistrationDefinitionsClient type.
type RegistrationDefinitionsServer struct {
	// BeginCreateOrUpdate is the fake for method RegistrationDefinitionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, registrationDefinitionID string, scope string, requestBody armmanagedservices.RegistrationDefinition, options *armmanagedservices.RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmanagedservices.RegistrationDefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method RegistrationDefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, registrationDefinitionID string, scope string, options *armmanagedservices.RegistrationDefinitionsClientDeleteOptions) (resp azfake.Responder[armmanagedservices.RegistrationDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegistrationDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, registrationDefinitionID string, options *armmanagedservices.RegistrationDefinitionsClientGetOptions) (resp azfake.Responder[armmanagedservices.RegistrationDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RegistrationDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armmanagedservices.RegistrationDefinitionsClientListOptions) (resp azfake.PagerResponder[armmanagedservices.RegistrationDefinitionsClientListResponse])
}

// NewRegistrationDefinitionsServerTransport creates a new instance of RegistrationDefinitionsServerTransport with the provided implementation.
// The returned RegistrationDefinitionsServerTransport instance is connected to an instance of armmanagedservices.RegistrationDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistrationDefinitionsServerTransport(srv *RegistrationDefinitionsServer) *RegistrationDefinitionsServerTransport {
	return &RegistrationDefinitionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmanagedservices.RegistrationDefinitionsClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmanagedservices.RegistrationDefinitionsClientListResponse]](),
	}
}

// RegistrationDefinitionsServerTransport connects instances of armmanagedservices.RegistrationDefinitionsClient to instances of RegistrationDefinitionsServer.
// Don't use this type directly, use NewRegistrationDefinitionsServerTransport instead.
type RegistrationDefinitionsServerTransport struct {
	srv                 *RegistrationDefinitionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmanagedservices.RegistrationDefinitionsClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmanagedservices.RegistrationDefinitionsClientListResponse]]
}

// Do implements the policy.Transporter interface for RegistrationDefinitionsServerTransport.
func (r *RegistrationDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegistrationDefinitionsClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RegistrationDefinitionsClient.Delete":
		resp, err = r.dispatchDelete(req)
	case "RegistrationDefinitionsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RegistrationDefinitionsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistrationDefinitionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationDefinitions/(?P<registrationDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmanagedservices.RegistrationDefinition](req)
		if err != nil {
			return nil, err
		}
		registrationDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationDefinitionId")])
		if err != nil {
			return nil, err
		}
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), registrationDefinitionIDParam, scopeParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RegistrationDefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationDefinitions/(?P<registrationDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	registrationDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationDefinitionId")])
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), registrationDefinitionIDParam, scopeParam, nil)
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

func (r *RegistrationDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationDefinitions/(?P<registrationDefinitionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	registrationDefinitionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("registrationDefinitionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), scopeParam, registrationDefinitionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RegistrationDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistrationDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/registrationDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armmanagedservices.RegistrationDefinitionsClientListOptions
		if filterParam != nil {
			options = &armmanagedservices.RegistrationDefinitionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := r.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmanagedservices.RegistrationDefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}
