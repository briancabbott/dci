//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

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

// ComputePoliciesClient contains the methods for the ComputePolicies group.
// Don't use this type directly, use NewComputePoliciesClient() instead.
type ComputePoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewComputePoliciesClient creates a new instance of ComputePoliciesClient with the specified values.
//   - subscriptionID - Get subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewComputePoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComputePoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ComputePoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the specified compute policy. During update, the compute policy with the specified
// name will be replaced with this new compute policy. An account supports, at most, 50 policies
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - computePolicyName - The name of the compute policy to create or update.
//   - parameters - Parameters supplied to create or update the compute policy. The max degree of parallelism per job property,
//     min priority per job property, or both must be present.
//   - options - ComputePoliciesClientCreateOrUpdateOptions contains the optional parameters for the ComputePoliciesClient.CreateOrUpdate
//     method.
func (client *ComputePoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters CreateOrUpdateComputePolicyParameters, options *ComputePoliciesClientCreateOrUpdateOptions) (ComputePoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ComputePoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, computePolicyName, parameters, options)
	if err != nil {
		return ComputePoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputePoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputePoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ComputePoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters CreateOrUpdateComputePolicyParameters, options *ComputePoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if computePolicyName == "" {
		return nil, errors.New("parameter computePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computePolicyName}", url.PathEscape(computePolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ComputePoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ComputePoliciesClientCreateOrUpdateResponse, error) {
	result := ComputePoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputePolicy); err != nil {
		return ComputePoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified compute policy from the specified Data Lake Analytics account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - computePolicyName - The name of the compute policy to delete.
//   - options - ComputePoliciesClientDeleteOptions contains the optional parameters for the ComputePoliciesClient.Delete method.
func (client *ComputePoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientDeleteOptions) (ComputePoliciesClientDeleteResponse, error) {
	var err error
	const operationName = "ComputePoliciesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, computePolicyName, options)
	if err != nil {
		return ComputePoliciesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputePoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ComputePoliciesClientDeleteResponse{}, err
	}
	return ComputePoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ComputePoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if computePolicyName == "" {
		return nil, errors.New("parameter computePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computePolicyName}", url.PathEscape(computePolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Data Lake Analytics compute policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - computePolicyName - The name of the compute policy to retrieve.
//   - options - ComputePoliciesClientGetOptions contains the optional parameters for the ComputePoliciesClient.Get method.
func (client *ComputePoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientGetOptions) (ComputePoliciesClientGetResponse, error) {
	var err error
	const operationName = "ComputePoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, computePolicyName, options)
	if err != nil {
		return ComputePoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputePoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputePoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ComputePoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if computePolicyName == "" {
		return nil, errors.New("parameter computePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computePolicyName}", url.PathEscape(computePolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComputePoliciesClient) getHandleResponse(resp *http.Response) (ComputePoliciesClientGetResponse, error) {
	result := ComputePoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputePolicy); err != nil {
		return ComputePoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Lists the Data Lake Analytics compute policies within the specified Data Lake Analytics account.
// An account supports, at most, 50 policies
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - options - ComputePoliciesClientListByAccountOptions contains the optional parameters for the ComputePoliciesClient.NewListByAccountPager
//     method.
func (client *ComputePoliciesClient) NewListByAccountPager(resourceGroupName string, accountName string, options *ComputePoliciesClientListByAccountOptions) *runtime.Pager[ComputePoliciesClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComputePoliciesClientListByAccountResponse]{
		More: func(page ComputePoliciesClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComputePoliciesClientListByAccountResponse) (ComputePoliciesClientListByAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ComputePoliciesClient.NewListByAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return ComputePoliciesClientListByAccountResponse{}, err
			}
			return client.listByAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *ComputePoliciesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ComputePoliciesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *ComputePoliciesClient) listByAccountHandleResponse(resp *http.Response) (ComputePoliciesClientListByAccountResponse, error) {
	result := ComputePoliciesClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputePolicyListResult); err != nil {
		return ComputePoliciesClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates the specified compute policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - computePolicyName - The name of the compute policy to update.
//   - options - ComputePoliciesClientUpdateOptions contains the optional parameters for the ComputePoliciesClient.Update method.
func (client *ComputePoliciesClient) Update(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientUpdateOptions) (ComputePoliciesClientUpdateResponse, error) {
	var err error
	const operationName = "ComputePoliciesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, computePolicyName, options)
	if err != nil {
		return ComputePoliciesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComputePoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComputePoliciesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ComputePoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, options *ComputePoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if computePolicyName == "" {
		return nil, errors.New("parameter computePolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{computePolicyName}", url.PathEscape(computePolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ComputePoliciesClient) updateHandleResponse(resp *http.Response) (ComputePoliciesClientUpdateResponse, error) {
	result := ComputePoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComputePolicy); err != nil {
		return ComputePoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
