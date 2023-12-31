//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolicyClient contains the methods for the SubscriptionPolicy group.
// Don't use this type directly, use NewPolicyClient() instead.
type PolicyClient struct {
	internal *arm.Client
}

// NewPolicyClient creates a new instance of PolicyClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicyClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PolicyClient{
		internal: cl,
	}
	return client, nil
}

// AddUpdatePolicyForTenant - Create or Update Subscription tenant policy for user's tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - options - PolicyClientAddUpdatePolicyForTenantOptions contains the optional parameters for the PolicyClient.AddUpdatePolicyForTenant
//     method.
func (client *PolicyClient) AddUpdatePolicyForTenant(ctx context.Context, body PutTenantPolicyRequestProperties, options *PolicyClientAddUpdatePolicyForTenantOptions) (PolicyClientAddUpdatePolicyForTenantResponse, error) {
	var err error
	const operationName = "PolicyClient.AddUpdatePolicyForTenant"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addUpdatePolicyForTenantCreateRequest(ctx, body, options)
	if err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	resp, err := client.addUpdatePolicyForTenantHandleResponse(httpResp)
	return resp, err
}

// addUpdatePolicyForTenantCreateRequest creates the AddUpdatePolicyForTenant request.
func (client *PolicyClient) addUpdatePolicyForTenantCreateRequest(ctx context.Context, body PutTenantPolicyRequestProperties, options *PolicyClientAddUpdatePolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies/default"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// addUpdatePolicyForTenantHandleResponse handles the AddUpdatePolicyForTenant response.
func (client *PolicyClient) addUpdatePolicyForTenantHandleResponse(resp *http.Response) (PolicyClientAddUpdatePolicyForTenantResponse, error) {
	result := PolicyClientAddUpdatePolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyResponse); err != nil {
		return PolicyClientAddUpdatePolicyForTenantResponse{}, err
	}
	return result, nil
}

// GetPolicyForTenant - Get the subscription tenant policy for the user's tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - options - PolicyClientGetPolicyForTenantOptions contains the optional parameters for the PolicyClient.GetPolicyForTenant
//     method.
func (client *PolicyClient) GetPolicyForTenant(ctx context.Context, options *PolicyClientGetPolicyForTenantOptions) (PolicyClientGetPolicyForTenantResponse, error) {
	var err error
	const operationName = "PolicyClient.GetPolicyForTenant"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getPolicyForTenantCreateRequest(ctx, options)
	if err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	resp, err := client.getPolicyForTenantHandleResponse(httpResp)
	return resp, err
}

// getPolicyForTenantCreateRequest creates the GetPolicyForTenant request.
func (client *PolicyClient) getPolicyForTenantCreateRequest(ctx context.Context, options *PolicyClientGetPolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPolicyForTenantHandleResponse handles the GetPolicyForTenant response.
func (client *PolicyClient) getPolicyForTenantHandleResponse(resp *http.Response) (PolicyClientGetPolicyForTenantResponse, error) {
	result := PolicyClientGetPolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyResponse); err != nil {
		return PolicyClientGetPolicyForTenantResponse{}, err
	}
	return result, nil
}

// NewListPolicyForTenantPager - Get the subscription tenant policy for the user's tenant.
//
// Generated from API version 2021-10-01
//   - options - PolicyClientListPolicyForTenantOptions contains the optional parameters for the PolicyClient.NewListPolicyForTenantPager
//     method.
func (client *PolicyClient) NewListPolicyForTenantPager(options *PolicyClientListPolicyForTenantOptions) *runtime.Pager[PolicyClientListPolicyForTenantResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicyClientListPolicyForTenantResponse]{
		More: func(page PolicyClientListPolicyForTenantResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicyClientListPolicyForTenantResponse) (PolicyClientListPolicyForTenantResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PolicyClient.NewListPolicyForTenantPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listPolicyForTenantCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PolicyClientListPolicyForTenantResponse{}, err
			}
			return client.listPolicyForTenantHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listPolicyForTenantCreateRequest creates the ListPolicyForTenant request.
func (client *PolicyClient) listPolicyForTenantCreateRequest(ctx context.Context, options *PolicyClientListPolicyForTenantOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Subscription/policies"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPolicyForTenantHandleResponse handles the ListPolicyForTenant response.
func (client *PolicyClient) listPolicyForTenantHandleResponse(resp *http.Response) (PolicyClientListPolicyForTenantResponse, error) {
	result := PolicyClientListPolicyForTenantResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GetTenantPolicyListResponse); err != nil {
		return PolicyClientListPolicyForTenantResponse{}, err
	}
	return result, nil
}
