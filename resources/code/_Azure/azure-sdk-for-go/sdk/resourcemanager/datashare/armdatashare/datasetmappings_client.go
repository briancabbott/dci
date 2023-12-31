//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatashare

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

// DataSetMappingsClient contains the methods for the DataSetMappings group.
// Don't use this type directly, use NewDataSetMappingsClient() instead.
type DataSetMappingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataSetMappingsClient creates a new instance of DataSetMappingsClient with the specified values.
//   - subscriptionID - The subscription identifier
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataSetMappingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataSetMappingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataSetMappingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a DataSetMapping
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareSubscriptionName - The name of the share subscription which will hold the data set sink.
//   - dataSetMappingName - The name of the data set mapping to be created.
//   - dataSetMapping - Destination data set configuration details.
//   - options - DataSetMappingsClientCreateOptions contains the optional parameters for the DataSetMappingsClient.Create method.
func (client *DataSetMappingsClient) Create(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping DataSetMappingClassification, options *DataSetMappingsClientCreateOptions) (DataSetMappingsClientCreateResponse, error) {
	var err error
	const operationName = "DataSetMappingsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName, dataSetMapping, options)
	if err != nil {
		return DataSetMappingsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataSetMappingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataSetMappingsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *DataSetMappingsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping DataSetMappingClassification, options *DataSetMappingsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}"
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
	if shareSubscriptionName == "" {
		return nil, errors.New("parameter shareSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareSubscriptionName}", url.PathEscape(shareSubscriptionName))
	if dataSetMappingName == "" {
		return nil, errors.New("parameter dataSetMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSetMappingName}", url.PathEscape(dataSetMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dataSetMapping); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataSetMappingsClient) createHandleResponse(resp *http.Response) (DataSetMappingsClientCreateResponse, error) {
	result := DataSetMappingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataSetMappingsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a DataSetMapping in a shareSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareSubscriptionName - The name of the shareSubscription.
//   - dataSetMappingName - The name of the dataSetMapping.
//   - options - DataSetMappingsClientDeleteOptions contains the optional parameters for the DataSetMappingsClient.Delete method.
func (client *DataSetMappingsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *DataSetMappingsClientDeleteOptions) (DataSetMappingsClientDeleteResponse, error) {
	var err error
	const operationName = "DataSetMappingsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName, options)
	if err != nil {
		return DataSetMappingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataSetMappingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DataSetMappingsClientDeleteResponse{}, err
	}
	return DataSetMappingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataSetMappingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *DataSetMappingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}"
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
	if shareSubscriptionName == "" {
		return nil, errors.New("parameter shareSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareSubscriptionName}", url.PathEscape(shareSubscriptionName))
	if dataSetMappingName == "" {
		return nil, errors.New("parameter dataSetMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSetMappingName}", url.PathEscape(dataSetMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a DataSetMapping in a shareSubscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareSubscriptionName - The name of the shareSubscription.
//   - dataSetMappingName - The name of the dataSetMapping.
//   - options - DataSetMappingsClientGetOptions contains the optional parameters for the DataSetMappingsClient.Get method.
func (client *DataSetMappingsClient) Get(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *DataSetMappingsClientGetOptions) (DataSetMappingsClientGetResponse, error) {
	var err error
	const operationName = "DataSetMappingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, shareSubscriptionName, dataSetMappingName, options)
	if err != nil {
		return DataSetMappingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataSetMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataSetMappingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataSetMappingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *DataSetMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings/{dataSetMappingName}"
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
	if shareSubscriptionName == "" {
		return nil, errors.New("parameter shareSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareSubscriptionName}", url.PathEscape(shareSubscriptionName))
	if dataSetMappingName == "" {
		return nil, errors.New("parameter dataSetMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSetMappingName}", url.PathEscape(dataSetMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataSetMappingsClient) getHandleResponse(resp *http.Response) (DataSetMappingsClientGetResponse, error) {
	result := DataSetMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DataSetMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByShareSubscriptionPager - List DataSetMappings in a share subscription
//
// Generated from API version 2020-09-01
//   - resourceGroupName - The resource group name.
//   - accountName - The name of the share account.
//   - shareSubscriptionName - The name of the share subscription.
//   - options - DataSetMappingsClientListByShareSubscriptionOptions contains the optional parameters for the DataSetMappingsClient.NewListByShareSubscriptionPager
//     method.
func (client *DataSetMappingsClient) NewListByShareSubscriptionPager(resourceGroupName string, accountName string, shareSubscriptionName string, options *DataSetMappingsClientListByShareSubscriptionOptions) *runtime.Pager[DataSetMappingsClientListByShareSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataSetMappingsClientListByShareSubscriptionResponse]{
		More: func(page DataSetMappingsClientListByShareSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataSetMappingsClientListByShareSubscriptionResponse) (DataSetMappingsClientListByShareSubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataSetMappingsClient.NewListByShareSubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByShareSubscriptionCreateRequest(ctx, resourceGroupName, accountName, shareSubscriptionName, options)
			}, nil)
			if err != nil {
				return DataSetMappingsClientListByShareSubscriptionResponse{}, err
			}
			return client.listByShareSubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByShareSubscriptionCreateRequest creates the ListByShareSubscription request.
func (client *DataSetMappingsClient) listByShareSubscriptionCreateRequest(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, options *DataSetMappingsClientListByShareSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataShare/accounts/{accountName}/shareSubscriptions/{shareSubscriptionName}/dataSetMappings"
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
	if shareSubscriptionName == "" {
		return nil, errors.New("parameter shareSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{shareSubscriptionName}", url.PathEscape(shareSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
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

// listByShareSubscriptionHandleResponse handles the ListByShareSubscription response.
func (client *DataSetMappingsClient) listByShareSubscriptionHandleResponse(resp *http.Response) (DataSetMappingsClientListByShareSubscriptionResponse, error) {
	result := DataSetMappingsClientListByShareSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataSetMappingList); err != nil {
		return DataSetMappingsClientListByShareSubscriptionResponse{}, err
	}
	return result, nil
}
