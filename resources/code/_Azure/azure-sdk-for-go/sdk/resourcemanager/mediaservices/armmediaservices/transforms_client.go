//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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

// TransformsClient contains the methods for the Transforms group.
// Don't use this type directly, use NewTransformsClient() instead.
type TransformsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTransformsClient creates a new instance of TransformsClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTransformsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TransformsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TransformsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a new Transform.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - transformName - The Transform name.
//   - parameters - The request parameters
//   - options - TransformsClientCreateOrUpdateOptions contains the optional parameters for the TransformsClient.CreateOrUpdate
//     method.
func (client *TransformsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters Transform, options *TransformsClientCreateOrUpdateOptions) (TransformsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TransformsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, transformName, parameters, options)
	if err != nil {
		return TransformsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TransformsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TransformsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters Transform, options *TransformsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}"
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
	if transformName == "" {
		return nil, errors.New("parameter transformName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformName}", url.PathEscape(transformName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TransformsClient) createOrUpdateHandleResponse(resp *http.Response) (TransformsClientCreateOrUpdateResponse, error) {
	result := TransformsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transform); err != nil {
		return TransformsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Transform.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - transformName - The Transform name.
//   - options - TransformsClientDeleteOptions contains the optional parameters for the TransformsClient.Delete method.
func (client *TransformsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, transformName string, options *TransformsClientDeleteOptions) (TransformsClientDeleteResponse, error) {
	var err error
	const operationName = "TransformsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, transformName, options)
	if err != nil {
		return TransformsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TransformsClientDeleteResponse{}, err
	}
	return TransformsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TransformsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, transformName string, options *TransformsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}"
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
	if transformName == "" {
		return nil, errors.New("parameter transformName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformName}", url.PathEscape(transformName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Transform.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - transformName - The Transform name.
//   - options - TransformsClientGetOptions contains the optional parameters for the TransformsClient.Get method.
func (client *TransformsClient) Get(ctx context.Context, resourceGroupName string, accountName string, transformName string, options *TransformsClientGetOptions) (TransformsClientGetResponse, error) {
	var err error
	const operationName = "TransformsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, transformName, options)
	if err != nil {
		return TransformsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TransformsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TransformsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, transformName string, options *TransformsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}"
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
	if transformName == "" {
		return nil, errors.New("parameter transformName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformName}", url.PathEscape(transformName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TransformsClient) getHandleResponse(resp *http.Response) (TransformsClientGetResponse, error) {
	result := TransformsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transform); err != nil {
		return TransformsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Transforms in the account.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - options - TransformsClientListOptions contains the optional parameters for the TransformsClient.NewListPager method.
func (client *TransformsClient) NewListPager(resourceGroupName string, accountName string, options *TransformsClientListOptions) *runtime.Pager[TransformsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TransformsClientListResponse]{
		More: func(page TransformsClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TransformsClientListResponse) (TransformsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TransformsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ODataNextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return TransformsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TransformsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TransformsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms"
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
	reqQP.Set("api-version", "2022-07-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TransformsClient) listHandleResponse(resp *http.Response) (TransformsClientListResponse, error) {
	result := TransformsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransformCollection); err != nil {
		return TransformsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates a Transform.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - transformName - The Transform name.
//   - parameters - The request parameters
//   - options - TransformsClientUpdateOptions contains the optional parameters for the TransformsClient.Update method.
func (client *TransformsClient) Update(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters Transform, options *TransformsClientUpdateOptions) (TransformsClientUpdateResponse, error) {
	var err error
	const operationName = "TransformsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, transformName, parameters, options)
	if err != nil {
		return TransformsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TransformsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TransformsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TransformsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, transformName string, parameters Transform, options *TransformsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/transforms/{transformName}"
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
	if transformName == "" {
		return nil, errors.New("parameter transformName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{transformName}", url.PathEscape(transformName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TransformsClient) updateHandleResponse(resp *http.Response) (TransformsClientUpdateResponse, error) {
	result := TransformsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Transform); err != nil {
		return TransformsClientUpdateResponse{}, err
	}
	return result, nil
}
