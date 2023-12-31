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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ProductPolicyServer is a fake server for instances of the armapimanagement.ProductPolicyClient type.
type ProductPolicyServer struct {
	// CreateOrUpdate is the fake for method ProductPolicyClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID armapimanagement.PolicyIDName, parameters armapimanagement.PolicyContract, options *armapimanagement.ProductPolicyClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.ProductPolicyClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProductPolicyClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID armapimanagement.PolicyIDName, ifMatch string, options *armapimanagement.ProductPolicyClientDeleteOptions) (resp azfake.Responder[armapimanagement.ProductPolicyClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProductPolicyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID armapimanagement.PolicyIDName, options *armapimanagement.ProductPolicyClientGetOptions) (resp azfake.Responder[armapimanagement.ProductPolicyClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method ProductPolicyClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, productID string, policyID armapimanagement.PolicyIDName, options *armapimanagement.ProductPolicyClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.ProductPolicyClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// ListByProduct is the fake for method ProductPolicyClient.ListByProduct
	// HTTP status codes to indicate success: http.StatusOK
	ListByProduct func(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *armapimanagement.ProductPolicyClientListByProductOptions) (resp azfake.Responder[armapimanagement.ProductPolicyClientListByProductResponse], errResp azfake.ErrorResponder)
}

// NewProductPolicyServerTransport creates a new instance of ProductPolicyServerTransport with the provided implementation.
// The returned ProductPolicyServerTransport instance is connected to an instance of armapimanagement.ProductPolicyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProductPolicyServerTransport(srv *ProductPolicyServer) *ProductPolicyServerTransport {
	return &ProductPolicyServerTransport{srv: srv}
}

// ProductPolicyServerTransport connects instances of armapimanagement.ProductPolicyClient to instances of ProductPolicyServer.
// Don't use this type directly, use NewProductPolicyServerTransport instead.
type ProductPolicyServerTransport struct {
	srv *ProductPolicyServer
}

// Do implements the policy.Transporter interface for ProductPolicyServerTransport.
func (p *ProductPolicyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ProductPolicyClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "ProductPolicyClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "ProductPolicyClient.Get":
		resp, err = p.dispatchGet(req)
	case "ProductPolicyClient.GetEntityTag":
		resp, err = p.dispatchGetEntityTag(req)
	case "ProductPolicyClient.ListByProduct":
		resp, err = p.dispatchListByProduct(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProductPolicyServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.PolicyContract](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.ProductPolicyClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.ProductPolicyClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, policyIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductPolicyServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, policyIDParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (p *ProductPolicyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
	if err != nil {
		return nil, err
	}
	formatParam := getOptional(armapimanagement.PolicyExportFormat(formatUnescaped))
	var options *armapimanagement.ProductPolicyClientGetOptions
	if formatParam != nil {
		options = &armapimanagement.ProductPolicyClientGetOptions{
			Format: formatParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, policyIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductPolicyServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if p.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, policyIDParam, nil)
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
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (p *ProductPolicyServerTransport) dispatchListByProduct(req *http.Request) (*http.Response, error) {
	if p.srv.ListByProduct == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByProduct not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/products/(?P<productId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	productIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("productId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListByProduct(req.Context(), resourceGroupNameParam, serviceNameParam, productIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
