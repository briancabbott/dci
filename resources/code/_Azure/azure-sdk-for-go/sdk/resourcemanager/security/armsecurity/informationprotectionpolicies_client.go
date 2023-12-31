//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// InformationProtectionPoliciesClient contains the methods for the InformationProtectionPolicies group.
// Don't use this type directly, use NewInformationProtectionPoliciesClient() instead.
type InformationProtectionPoliciesClient struct {
	internal *arm.Client
}

// NewInformationProtectionPoliciesClient creates a new instance of InformationProtectionPoliciesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInformationProtectionPoliciesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*InformationProtectionPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InformationProtectionPoliciesClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Details of the information protection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - informationProtectionPolicyName - Name of the information protection policy.
//   - informationProtectionPolicy - Information protection policy.
//   - options - InformationProtectionPoliciesClientCreateOrUpdateOptions contains the optional parameters for the InformationProtectionPoliciesClient.CreateOrUpdate
//     method.
func (client *InformationProtectionPoliciesClient) CreateOrUpdate(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesClientCreateOrUpdateOptions) (InformationProtectionPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "InformationProtectionPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, scope, informationProtectionPolicyName, informationProtectionPolicy, options)
	if err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InformationProtectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, informationProtectionPolicy InformationProtectionPolicy, options *InformationProtectionPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, informationProtectionPolicy); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *InformationProtectionPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientCreateOrUpdateResponse, error) {
	result := InformationProtectionPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Details of the information protection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - informationProtectionPolicyName - Name of the information protection policy.
//   - options - InformationProtectionPoliciesClientGetOptions contains the optional parameters for the InformationProtectionPoliciesClient.Get
//     method.
func (client *InformationProtectionPoliciesClient) Get(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, options *InformationProtectionPoliciesClientGetOptions) (InformationProtectionPoliciesClientGetResponse, error) {
	var err error
	const operationName = "InformationProtectionPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, informationProtectionPolicyName, options)
	if err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InformationProtectionPoliciesClient) getCreateRequest(ctx context.Context, scope string, informationProtectionPolicyName InformationProtectionPolicyName, options *InformationProtectionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies/{informationProtectionPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if informationProtectionPolicyName == "" {
		return nil, errors.New("parameter informationProtectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{informationProtectionPolicyName}", url.PathEscape(string(informationProtectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InformationProtectionPoliciesClient) getHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientGetResponse, error) {
	result := InformationProtectionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicy); err != nil {
		return InformationProtectionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Information protection policies of a specific management group.
//
// Generated from API version 2017-08-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - options - InformationProtectionPoliciesClientListOptions contains the optional parameters for the InformationProtectionPoliciesClient.NewListPager
//     method.
func (client *InformationProtectionPoliciesClient) NewListPager(scope string, options *InformationProtectionPoliciesClientListOptions) *runtime.Pager[InformationProtectionPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InformationProtectionPoliciesClientListResponse]{
		More: func(page InformationProtectionPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InformationProtectionPoliciesClientListResponse) (InformationProtectionPoliciesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "InformationProtectionPoliciesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return InformationProtectionPoliciesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *InformationProtectionPoliciesClient) listCreateRequest(ctx context.Context, scope string, options *InformationProtectionPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/informationProtectionPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InformationProtectionPoliciesClient) listHandleResponse(resp *http.Response) (InformationProtectionPoliciesClientListResponse, error) {
	result := InformationProtectionPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformationProtectionPolicyList); err != nil {
		return InformationProtectionPoliciesClientListResponse{}, err
	}
	return result, nil
}
