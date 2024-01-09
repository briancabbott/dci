//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TopicsClient contains the methods for the Topics group.
// Don't use this type directly, use NewTopicsClient() instead.
type TopicsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTopicsClient creates a new instance of TopicsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTopicsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TopicsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TopicsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates a new topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - topicInfo - Topic information.
//   - options - TopicsClientBeginCreateOrUpdateOptions contains the optional parameters for the TopicsClient.BeginCreateOrUpdate
//     method.
func (client *TopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (*runtime.Poller[TopicsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, topicName, topicInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopicsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Asynchronously creates a new topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *TopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TopicsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, topicName, topicInfo, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, topicName string, topicInfo Topic, options *TopicsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, topicInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete existing topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - options - TopicsClientBeginDeleteOptions contains the optional parameters for the TopicsClient.BeginDelete method.
func (client *TopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (*runtime.Poller[TopicsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, topicName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopicsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete existing topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *TopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "TopicsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, topicName, options)
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
func (client *TopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of a topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - options - TopicsClientGetOptions contains the optional parameters for the TopicsClient.Get method.
func (client *TopicsClient) Get(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientGetOptions) (TopicsClientGetResponse, error) {
	var err error
	const operationName = "TopicsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TopicsClient) getHandleResponse(resp *http.Response) (TopicsClientGetResponse, error) {
	result := TopicsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Topic); err != nil {
		return TopicsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the topics under a resource group.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - options - TopicsClientListByResourceGroupOptions contains the optional parameters for the TopicsClient.NewListByResourceGroupPager
//     method.
func (client *TopicsClient) NewListByResourceGroupPager(resourceGroupName string, options *TopicsClientListByResourceGroupOptions) *runtime.Pager[TopicsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicsClientListByResourceGroupResponse]{
		More: func(page TopicsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopicsClientListByResourceGroupResponse) (TopicsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return TopicsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *TopicsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *TopicsClient) listByResourceGroupHandleResponse(resp *http.Response) (TopicsClientListByResourceGroupResponse, error) {
	result := TopicsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicsListResult); err != nil {
		return TopicsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the topics under an Azure subscription.
//
// Generated from API version 2023-12-15-preview
//   - options - TopicsClientListBySubscriptionOptions contains the optional parameters for the TopicsClient.NewListBySubscriptionPager
//     method.
func (client *TopicsClient) NewListBySubscriptionPager(options *TopicsClientListBySubscriptionOptions) *runtime.Pager[TopicsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicsClientListBySubscriptionResponse]{
		More: func(page TopicsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TopicsClientListBySubscriptionResponse) (TopicsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TopicsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *TopicsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/topics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *TopicsClient) listBySubscriptionHandleResponse(resp *http.Response) (TopicsClientListBySubscriptionResponse, error) {
	result := TopicsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicsListResult); err != nil {
		return TopicsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListEventTypesPager - List event types for a topic.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - providerNamespace - Namespace of the provider of the topic.
//   - resourceTypeName - Name of the topic type.
//   - resourceName - Name of the topic.
//   - options - TopicsClientListEventTypesOptions contains the optional parameters for the TopicsClient.NewListEventTypesPager
//     method.
func (client *TopicsClient) NewListEventTypesPager(resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, options *TopicsClientListEventTypesOptions) *runtime.Pager[TopicsClientListEventTypesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TopicsClientListEventTypesResponse]{
		More: func(page TopicsClientListEventTypesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TopicsClientListEventTypesResponse) (TopicsClientListEventTypesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TopicsClient.NewListEventTypesPager")
			req, err := client.listEventTypesCreateRequest(ctx, resourceGroupName, providerNamespace, resourceTypeName, resourceName, options)
			if err != nil {
				return TopicsClientListEventTypesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TopicsClientListEventTypesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TopicsClientListEventTypesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listEventTypesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listEventTypesCreateRequest creates the ListEventTypes request.
func (client *TopicsClient) listEventTypesCreateRequest(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, options *TopicsClientListEventTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerNamespace}/{resourceTypeName}/{resourceName}/providers/Microsoft.EventGrid/eventTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if resourceTypeName == "" {
		return nil, errors.New("parameter resourceTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceTypeName}", url.PathEscape(resourceTypeName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEventTypesHandleResponse handles the ListEventTypes response.
func (client *TopicsClient) listEventTypesHandleResponse(resp *http.Response) (TopicsClientListEventTypesResponse, error) {
	result := TopicsClientListEventTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventTypesListResult); err != nil {
		return TopicsClientListEventTypesResponse{}, err
	}
	return result, nil
}

// ListSharedAccessKeys - List the two keys used to publish to a topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - options - TopicsClientListSharedAccessKeysOptions contains the optional parameters for the TopicsClient.ListSharedAccessKeys
//     method.
func (client *TopicsClient) ListSharedAccessKeys(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientListSharedAccessKeysOptions) (TopicsClientListSharedAccessKeysResponse, error) {
	var err error
	const operationName = "TopicsClient.ListSharedAccessKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSharedAccessKeysCreateRequest(ctx, resourceGroupName, topicName, options)
	if err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	resp, err := client.listSharedAccessKeysHandleResponse(httpResp)
	return resp, err
}

// listSharedAccessKeysCreateRequest creates the ListSharedAccessKeys request.
func (client *TopicsClient) listSharedAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, topicName string, options *TopicsClientListSharedAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSharedAccessKeysHandleResponse handles the ListSharedAccessKeys response.
func (client *TopicsClient) listSharedAccessKeysHandleResponse(resp *http.Response) (TopicsClientListSharedAccessKeysResponse, error) {
	result := TopicsClientListSharedAccessKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TopicSharedAccessKeys); err != nil {
		return TopicsClientListSharedAccessKeysResponse{}, err
	}
	return result, nil
}

// BeginRegenerateKey - Regenerate a shared access key for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - regenerateKeyRequest - Request body to regenerate key.
//   - options - TopicsClientBeginRegenerateKeyOptions contains the optional parameters for the TopicsClient.BeginRegenerateKey
//     method.
func (client *TopicsClient) BeginRegenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (*runtime.Poller[TopicsClientRegenerateKeyResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.regenerateKey(ctx, resourceGroupName, topicName, regenerateKeyRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicsClientRegenerateKeyResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopicsClientRegenerateKeyResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RegenerateKey - Regenerate a shared access key for a topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *TopicsClient) regenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (*http.Response, error) {
	var err error
	const operationName = "TopicsClient.BeginRegenerateKey"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, topicName, regenerateKeyRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *TopicsClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest TopicRegenerateKeyRequest, options *TopicsClientBeginRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, regenerateKeyRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Asynchronously updates a topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - topicName - Name of the topic.
//   - topicUpdateParameters - Topic update information.
//   - options - TopicsClientBeginUpdateOptions contains the optional parameters for the TopicsClient.BeginUpdate method.
func (client *TopicsClient) BeginUpdate(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (*runtime.Poller[TopicsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, topicName, topicUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TopicsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TopicsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Asynchronously updates a topic with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
func (client *TopicsClient) update(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "TopicsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, topicName, topicUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *TopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters TopicUpdateParameters, options *TopicsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if topicName == "" {
		return nil, errors.New("parameter topicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{topicName}", url.PathEscape(topicName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, topicUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
