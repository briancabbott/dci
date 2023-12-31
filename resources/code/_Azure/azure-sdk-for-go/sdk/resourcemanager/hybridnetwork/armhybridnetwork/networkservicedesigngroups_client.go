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

// NetworkServiceDesignGroupsClient contains the methods for the NetworkServiceDesignGroups group.
// Don't use this type directly, use NewNetworkServiceDesignGroupsClient() instead.
type NetworkServiceDesignGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkServiceDesignGroupsClient creates a new instance of NetworkServiceDesignGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkServiceDesignGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkServiceDesignGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkServiceDesignGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network service design group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - parameters - Parameters supplied to the create or update publisher network service design group operation.
//   - options - NetworkServiceDesignGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkServiceDesignGroupsClient.BeginCreateOrUpdate
//     method.
func (client *NetworkServiceDesignGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters NetworkServiceDesignGroup, options *NetworkServiceDesignGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkServiceDesignGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkServiceDesignGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkServiceDesignGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a network service design group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkServiceDesignGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters NetworkServiceDesignGroup, options *NetworkServiceDesignGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkServiceDesignGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, parameters, options)
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
func (client *NetworkServiceDesignGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters NetworkServiceDesignGroup, options *NetworkServiceDesignGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
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

// BeginDelete - Deletes a specified network service design group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - options - NetworkServiceDesignGroupsClientBeginDeleteOptions contains the optional parameters for the NetworkServiceDesignGroupsClient.BeginDelete
//     method.
func (client *NetworkServiceDesignGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignGroupsClientBeginDeleteOptions) (*runtime.Poller[NetworkServiceDesignGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkServiceDesignGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkServiceDesignGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a specified network service design group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkServiceDesignGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkServiceDesignGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, options)
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
func (client *NetworkServiceDesignGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
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

// Get - Gets information about the specified networkServiceDesign group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - options - NetworkServiceDesignGroupsClientGetOptions contains the optional parameters for the NetworkServiceDesignGroupsClient.Get
//     method.
func (client *NetworkServiceDesignGroupsClient) Get(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignGroupsClientGetOptions) (NetworkServiceDesignGroupsClientGetResponse, error) {
	var err error
	const operationName = "NetworkServiceDesignGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, options)
	if err != nil {
		return NetworkServiceDesignGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkServiceDesignGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkServiceDesignGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkServiceDesignGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
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
func (client *NetworkServiceDesignGroupsClient) getHandleResponse(resp *http.Response) (NetworkServiceDesignGroupsClientGetResponse, error) {
	result := NetworkServiceDesignGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignGroup); err != nil {
		return NetworkServiceDesignGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPublisherPager - Gets information of the network service design groups under a publisher.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - options - NetworkServiceDesignGroupsClientListByPublisherOptions contains the optional parameters for the NetworkServiceDesignGroupsClient.NewListByPublisherPager
//     method.
func (client *NetworkServiceDesignGroupsClient) NewListByPublisherPager(resourceGroupName string, publisherName string, options *NetworkServiceDesignGroupsClientListByPublisherOptions) *runtime.Pager[NetworkServiceDesignGroupsClientListByPublisherResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkServiceDesignGroupsClientListByPublisherResponse]{
		More: func(page NetworkServiceDesignGroupsClientListByPublisherResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkServiceDesignGroupsClientListByPublisherResponse) (NetworkServiceDesignGroupsClientListByPublisherResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkServiceDesignGroupsClient.NewListByPublisherPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPublisherCreateRequest(ctx, resourceGroupName, publisherName, options)
			}, nil)
			if err != nil {
				return NetworkServiceDesignGroupsClientListByPublisherResponse{}, err
			}
			return client.listByPublisherHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPublisherCreateRequest creates the ListByPublisher request.
func (client *NetworkServiceDesignGroupsClient) listByPublisherCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, options *NetworkServiceDesignGroupsClientListByPublisherOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
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
func (client *NetworkServiceDesignGroupsClient) listByPublisherHandleResponse(resp *http.Response) (NetworkServiceDesignGroupsClientListByPublisherResponse, error) {
	result := NetworkServiceDesignGroupsClientListByPublisherResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignGroupListResult); err != nil {
		return NetworkServiceDesignGroupsClientListByPublisherResponse{}, err
	}
	return result, nil
}

// Update - Updates a network service design groups resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - parameters - Parameters supplied to the create or update publisher network service design group operation.
//   - options - NetworkServiceDesignGroupsClientUpdateOptions contains the optional parameters for the NetworkServiceDesignGroupsClient.Update
//     method.
func (client *NetworkServiceDesignGroupsClient) Update(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters TagsObject, options *NetworkServiceDesignGroupsClientUpdateOptions) (NetworkServiceDesignGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "NetworkServiceDesignGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, parameters, options)
	if err != nil {
		return NetworkServiceDesignGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkServiceDesignGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkServiceDesignGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *NetworkServiceDesignGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, parameters TagsObject, options *NetworkServiceDesignGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
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
func (client *NetworkServiceDesignGroupsClient) updateHandleResponse(resp *http.Response) (NetworkServiceDesignGroupsClientUpdateResponse, error) {
	result := NetworkServiceDesignGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignGroup); err != nil {
		return NetworkServiceDesignGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
