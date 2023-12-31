//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// NetworkFunctionDefinitionGroupsClient contains the methods for the NetworkFunctionDefinitionGroups group.
// Don't use this type directly, use NewNetworkFunctionDefinitionGroupsClient() instead.
type NetworkFunctionDefinitionGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkFunctionDefinitionGroupsClient creates a new instance of NetworkFunctionDefinitionGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkFunctionDefinitionGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkFunctionDefinitionGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkFunctionDefinitionGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network function definition group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - parameters - Parameters supplied to the create or update publisher network function definition group operation.
//   - options - NetworkFunctionDefinitionGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkFunctionDefinitionGroupsClient.BeginCreateOrUpdate
//     method.
func (client *NetworkFunctionDefinitionGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, parameters NetworkFunctionDefinitionGroup, options *NetworkFunctionDefinitionGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkFunctionDefinitionGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionDefinitionGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionDefinitionGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a network function definition group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkFunctionDefinitionGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, parameters NetworkFunctionDefinitionGroup, options *NetworkFunctionDefinitionGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkFunctionDefinitionGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, parameters NetworkFunctionDefinitionGroup, options *NetworkFunctionDefinitionGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a specified network function definition group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - options - NetworkFunctionDefinitionGroupsClientBeginDeleteOptions contains the optional parameters for the NetworkFunctionDefinitionGroupsClient.BeginDelete
//     method.
func (client *NetworkFunctionDefinitionGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionGroupsClientBeginDeleteOptions) (*runtime.Poller[NetworkFunctionDefinitionGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkFunctionDefinitionGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkFunctionDefinitionGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a specified network function definition group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkFunctionDefinitionGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkFunctionDefinitionGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified networkFunctionDefinition group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - options - NetworkFunctionDefinitionGroupsClientGetOptions contains the optional parameters for the NetworkFunctionDefinitionGroupsClient.Get
//     method.
func (client *NetworkFunctionDefinitionGroupsClient) Get(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionGroupsClientGetOptions) (NetworkFunctionDefinitionGroupsClientGetResponse, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, options)
	if err != nil {
		return NetworkFunctionDefinitionGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionDefinitionGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionDefinitionGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkFunctionDefinitionGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, options *NetworkFunctionDefinitionGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkFunctionDefinitionGroupsClient) getHandleResponse(resp *http.Response) (NetworkFunctionDefinitionGroupsClientGetResponse, error) {
	result := NetworkFunctionDefinitionGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionGroup); err != nil {
		return NetworkFunctionDefinitionGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPublisherPager - Gets information of the network function definition groups under a publisher.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - options - NetworkFunctionDefinitionGroupsClientListByPublisherOptions contains the optional parameters for the NetworkFunctionDefinitionGroupsClient.NewListByPublisherPager
//     method.
func (client *NetworkFunctionDefinitionGroupsClient) NewListByPublisherPager(resourceGroupName string, publisherName string, options *NetworkFunctionDefinitionGroupsClientListByPublisherOptions) *runtime.Pager[NetworkFunctionDefinitionGroupsClientListByPublisherResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkFunctionDefinitionGroupsClientListByPublisherResponse]{
		More: func(page NetworkFunctionDefinitionGroupsClientListByPublisherResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkFunctionDefinitionGroupsClientListByPublisherResponse) (NetworkFunctionDefinitionGroupsClientListByPublisherResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkFunctionDefinitionGroupsClient.NewListByPublisherPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPublisherCreateRequest(ctx, resourceGroupName, publisherName, options)
			}, nil)
			if err != nil {
				return NetworkFunctionDefinitionGroupsClientListByPublisherResponse{}, err
			}
			return client.listByPublisherHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPublisherCreateRequest creates the ListByPublisher request.
func (client *NetworkFunctionDefinitionGroupsClient) listByPublisherCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, options *NetworkFunctionDefinitionGroupsClientListByPublisherOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPublisherHandleResponse handles the ListByPublisher response.
func (client *NetworkFunctionDefinitionGroupsClient) listByPublisherHandleResponse(resp *http.Response) (NetworkFunctionDefinitionGroupsClientListByPublisherResponse, error) {
	result := NetworkFunctionDefinitionGroupsClientListByPublisherResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionGroupListResult); err != nil {
		return NetworkFunctionDefinitionGroupsClientListByPublisherResponse{}, err
	}
	return result, nil
}

// Update - Updates a network function definition group resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkFunctionDefinitionGroupName - The name of the network function definition group.
//   - parameters - Parameters supplied to the create or update publisher network function definition group operation.
//   - options - NetworkFunctionDefinitionGroupsClientUpdateOptions contains the optional parameters for the NetworkFunctionDefinitionGroupsClient.Update
//     method.
func (client *NetworkFunctionDefinitionGroupsClient) Update(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, parameters TagsObject, options *NetworkFunctionDefinitionGroupsClientUpdateOptions) (NetworkFunctionDefinitionGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "NetworkFunctionDefinitionGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, publisherName, networkFunctionDefinitionGroupName, parameters, options)
	if err != nil {
		return NetworkFunctionDefinitionGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkFunctionDefinitionGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkFunctionDefinitionGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *NetworkFunctionDefinitionGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkFunctionDefinitionGroupName string, parameters TagsObject, options *NetworkFunctionDefinitionGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkFunctionDefinitionGroups/{networkFunctionDefinitionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkFunctionDefinitionGroupName == "" {
		return nil, errors.New("parameter networkFunctionDefinitionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkFunctionDefinitionGroupName}", url.PathEscape(networkFunctionDefinitionGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *NetworkFunctionDefinitionGroupsClient) updateHandleResponse(resp *http.Response) (NetworkFunctionDefinitionGroupsClientUpdateResponse, error) {
	result := NetworkFunctionDefinitionGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionDefinitionGroup); err != nil {
		return NetworkFunctionDefinitionGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
