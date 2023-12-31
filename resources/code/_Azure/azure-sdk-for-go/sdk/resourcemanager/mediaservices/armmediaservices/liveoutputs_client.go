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

// LiveOutputsClient contains the methods for the LiveOutputs group.
// Don't use this type directly, use NewLiveOutputsClient() instead.
type LiveOutputsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLiveOutputsClient creates a new instance of LiveOutputsClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLiveOutputsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LiveOutputsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LiveOutputsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AsyncOperation - Get a Live Output operation status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - operationID - The ID of an ongoing async operation.
//   - options - LiveOutputsClientAsyncOperationOptions contains the optional parameters for the LiveOutputsClient.AsyncOperation
//     method.
func (client *LiveOutputsClient) AsyncOperation(ctx context.Context, resourceGroupName string, accountName string, operationID string, options *LiveOutputsClientAsyncOperationOptions) (LiveOutputsClientAsyncOperationResponse, error) {
	var err error
	const operationName = "LiveOutputsClient.AsyncOperation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.asyncOperationCreateRequest(ctx, resourceGroupName, accountName, operationID, options)
	if err != nil {
		return LiveOutputsClientAsyncOperationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LiveOutputsClientAsyncOperationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LiveOutputsClientAsyncOperationResponse{}, err
	}
	resp, err := client.asyncOperationHandleResponse(httpResp)
	return resp, err
}

// asyncOperationCreateRequest creates the AsyncOperation request.
func (client *LiveOutputsClient) asyncOperationCreateRequest(ctx context.Context, resourceGroupName string, accountName string, operationID string, options *LiveOutputsClientAsyncOperationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveOutputOperations/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// asyncOperationHandleResponse handles the AsyncOperation response.
func (client *LiveOutputsClient) asyncOperationHandleResponse(resp *http.Response) (LiveOutputsClientAsyncOperationResponse, error) {
	result := LiveOutputsClientAsyncOperationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AsyncOperationResult); err != nil {
		return LiveOutputsClientAsyncOperationResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates a new live output.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - liveEventName - The name of the live event, maximum length is 32.
//   - liveOutputName - The name of the live output.
//   - parameters - Live Output properties needed for creation.
//   - options - LiveOutputsClientBeginCreateOptions contains the optional parameters for the LiveOutputsClient.BeginCreate method.
func (client *LiveOutputsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*runtime.Poller[LiveOutputsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LiveOutputsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LiveOutputsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a new live output.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *LiveOutputsClient) create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "LiveOutputsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *LiveOutputsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
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
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - liveEventName - The name of the live event, maximum length is 32.
//   - liveOutputName - The name of the live output.
//   - options - LiveOutputsClientBeginDeleteOptions contains the optional parameters for the LiveOutputsClient.BeginDelete method.
func (client *LiveOutputsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*runtime.Poller[LiveOutputsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LiveOutputsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LiveOutputsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *LiveOutputsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LiveOutputsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LiveOutputsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
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
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a live output.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - liveEventName - The name of the live event, maximum length is 32.
//   - liveOutputName - The name of the live output.
//   - options - LiveOutputsClientGetOptions contains the optional parameters for the LiveOutputsClient.Get method.
func (client *LiveOutputsClient) Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientGetOptions) (LiveOutputsClientGetResponse, error) {
	var err error
	const operationName = "LiveOutputsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
	if err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LiveOutputsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LiveOutputsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
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
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LiveOutputsClient) getHandleResponse(resp *http.Response) (LiveOutputsClientGetResponse, error) {
	result := LiveOutputsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveOutput); err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the live outputs of a live event.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - liveEventName - The name of the live event, maximum length is 32.
//   - options - LiveOutputsClientListOptions contains the optional parameters for the LiveOutputsClient.NewListPager method.
func (client *LiveOutputsClient) NewListPager(resourceGroupName string, accountName string, liveEventName string, options *LiveOutputsClientListOptions) *runtime.Pager[LiveOutputsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LiveOutputsClientListResponse]{
		More: func(page LiveOutputsClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LiveOutputsClientListResponse) (LiveOutputsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LiveOutputsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ODataNextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
			}, nil)
			if err != nil {
				return LiveOutputsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LiveOutputsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveOutputsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs"
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
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LiveOutputsClient) listHandleResponse(resp *http.Response) (LiveOutputsClientListResponse, error) {
	result := LiveOutputsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveOutputListResult); err != nil {
		return LiveOutputsClientListResponse{}, err
	}
	return result, nil
}

// OperationLocation - Get a Live Output operation status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - liveEventName - The name of the live event, maximum length is 32.
//   - liveOutputName - The name of the live output.
//   - operationID - The ID of an ongoing async operation.
//   - options - LiveOutputsClientOperationLocationOptions contains the optional parameters for the LiveOutputsClient.OperationLocation
//     method.
func (client *LiveOutputsClient) OperationLocation(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, operationID string, options *LiveOutputsClientOperationLocationOptions) (LiveOutputsClientOperationLocationResponse, error) {
	var err error
	const operationName = "LiveOutputsClient.OperationLocation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.operationLocationCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, operationID, options)
	if err != nil {
		return LiveOutputsClientOperationLocationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LiveOutputsClientOperationLocationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return LiveOutputsClientOperationLocationResponse{}, err
	}
	resp, err := client.operationLocationHandleResponse(httpResp)
	return resp, err
}

// operationLocationCreateRequest creates the OperationLocation request.
func (client *LiveOutputsClient) operationLocationCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, operationID string, options *LiveOutputsClientOperationLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}/operationLocations/{operationId}"
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
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// operationLocationHandleResponse handles the OperationLocation response.
func (client *LiveOutputsClient) operationLocationHandleResponse(resp *http.Response) (LiveOutputsClientOperationLocationResponse, error) {
	result := LiveOutputsClientOperationLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveOutput); err != nil {
		return LiveOutputsClientOperationLocationResponse{}, err
	}
	return result, nil
}
