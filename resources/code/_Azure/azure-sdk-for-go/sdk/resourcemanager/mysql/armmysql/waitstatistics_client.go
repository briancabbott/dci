//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// WaitStatisticsClient contains the methods for the WaitStatistics group.
// Don't use this type directly, use NewWaitStatisticsClient() instead.
type WaitStatisticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWaitStatisticsClient creates a new instance of WaitStatisticsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWaitStatisticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WaitStatisticsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WaitStatisticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve wait statistics for specified identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - waitStatisticsID - The Wait Statistic identifier.
//   - options - WaitStatisticsClientGetOptions contains the optional parameters for the WaitStatisticsClient.Get method.
func (client *WaitStatisticsClient) Get(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsClientGetOptions) (WaitStatisticsClientGetResponse, error) {
	var err error
	const operationName = "WaitStatisticsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, waitStatisticsID, options)
	if err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WaitStatisticsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WaitStatisticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, waitStatisticsID string, options *WaitStatisticsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics/{waitStatisticsId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if waitStatisticsID == "" {
		return nil, errors.New("parameter waitStatisticsID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{waitStatisticsId}", url.PathEscape(waitStatisticsID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WaitStatisticsClient) getHandleResponse(resp *http.Response) (WaitStatisticsClientGetResponse, error) {
	result := WaitStatisticsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatistic); err != nil {
		return WaitStatisticsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Retrieve wait statistics for specified aggregation window.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - parameters - The required parameters for retrieving wait statistics.
//   - options - WaitStatisticsClientListByServerOptions contains the optional parameters for the WaitStatisticsClient.NewListByServerPager
//     method.
func (client *WaitStatisticsClient) NewListByServerPager(resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsClientListByServerOptions) *runtime.Pager[WaitStatisticsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[WaitStatisticsClientListByServerResponse]{
		More: func(page WaitStatisticsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WaitStatisticsClientListByServerResponse) (WaitStatisticsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WaitStatisticsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
			}, nil)
			if err != nil {
				return WaitStatisticsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *WaitStatisticsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters WaitStatisticsInput, options *WaitStatisticsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/waitStatistics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *WaitStatisticsClient) listByServerHandleResponse(resp *http.Response) (WaitStatisticsClientListByServerResponse, error) {
	result := WaitStatisticsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WaitStatisticsResultList); err != nil {
		return WaitStatisticsClientListByServerResponse{}, err
	}
	return result, nil
}
