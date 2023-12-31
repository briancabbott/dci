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
	"strconv"
	"strings"
)

// StorageAccountsClient contains the methods for the StorageAccounts group.
// Don't use this type directly, use NewStorageAccountsClient() instead.
type StorageAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageAccountsClient creates a new instance of StorageAccountsClient with the specified values.
//   - subscriptionID - Get subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Add - Updates the specified Data Lake Analytics account to add an Azure Storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure Storage account to add
//   - parameters - The parameters containing the access key and optional suffix for the Azure Storage Account.
//   - options - StorageAccountsClientAddOptions contains the optional parameters for the StorageAccountsClient.Add method.
func (client *StorageAccountsClient) Add(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsClientAddOptions) (StorageAccountsClientAddResponse, error) {
	var err error
	const operationName = "StorageAccountsClient.Add"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, parameters, options)
	if err != nil {
		return StorageAccountsClientAddResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientAddResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountsClientAddResponse{}, err
	}
	return StorageAccountsClientAddResponse{}, nil
}

// addCreateRequest creates the Add request.
func (client *StorageAccountsClient) addCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsClientAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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

// Delete - Updates the specified Data Lake Analytics account to remove an Azure Storage account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure Storage account to remove
//   - options - StorageAccountsClientDeleteOptions contains the optional parameters for the StorageAccountsClient.Delete method.
func (client *StorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientDeleteOptions) (StorageAccountsClientDeleteResponse, error) {
	var err error
	const operationName = "StorageAccountsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountsClientDeleteResponse{}, err
	}
	return StorageAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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

// Get - Gets the specified Azure Storage account linked to the given Data Lake Analytics account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure Storage account for which to retrieve the details.
//   - options - StorageAccountsClientGetOptions contains the optional parameters for the StorageAccountsClient.Get method.
func (client *StorageAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientGetOptions) (StorageAccountsClientGetResponse, error) {
	var err error
	const operationName = "StorageAccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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
func (client *StorageAccountsClient) getHandleResponse(resp *http.Response) (StorageAccountsClientGetResponse, error) {
	result := StorageAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformation); err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	return result, nil
}

// GetStorageContainer - Gets the specified Azure Storage container associated with the given Data Lake Analytics and Azure
// Storage accounts.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure storage account from which to retrieve the blob container.
//   - containerName - The name of the Azure storage container to retrieve
//   - options - StorageAccountsClientGetStorageContainerOptions contains the optional parameters for the StorageAccountsClient.GetStorageContainer
//     method.
func (client *StorageAccountsClient) GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientGetStorageContainerOptions) (StorageAccountsClientGetStorageContainerResponse, error) {
	var err error
	const operationName = "StorageAccountsClient.GetStorageContainer"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStorageContainerCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
	if err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	resp, err := client.getStorageContainerHandleResponse(httpResp)
	return resp, err
}

// getStorageContainerCreateRequest creates the GetStorageContainer request.
func (client *StorageAccountsClient) getStorageContainerCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientGetStorageContainerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
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

// getStorageContainerHandleResponse handles the GetStorageContainer response.
func (client *StorageAccountsClient) getStorageContainerHandleResponse(resp *http.Response) (StorageAccountsClientGetStorageContainerResponse, error) {
	result := StorageAccountsClientGetStorageContainerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainer); err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Gets the first page of Azure Storage accounts, if any, linked to the specified Data Lake Analytics
// account. The response includes a link to the next page, if any.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - options - StorageAccountsClientListByAccountOptions contains the optional parameters for the StorageAccountsClient.NewListByAccountPager
//     method.
func (client *StorageAccountsClient) NewListByAccountPager(resourceGroupName string, accountName string, options *StorageAccountsClientListByAccountOptions) *runtime.Pager[StorageAccountsClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListByAccountResponse]{
		More: func(page StorageAccountsClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListByAccountResponse) (StorageAccountsClientListByAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAccountsClient.NewListByAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return StorageAccountsClientListByAccountResponse{}, err
			}
			return client.listByAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *StorageAccountsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *StorageAccountsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *StorageAccountsClient) listByAccountHandleResponse(resp *http.Response) (StorageAccountsClientListByAccountResponse, error) {
	result := StorageAccountsClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformationListResult); err != nil {
		return StorageAccountsClientListByAccountResponse{}, err
	}
	return result, nil
}

// NewListSasTokensPager - Gets the SAS token associated with the specified Data Lake Analytics and Azure Storage account
// and container combination.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure storage account for which the SAS token is being requested.
//   - containerName - The name of the Azure storage container for which the SAS token is being requested.
//   - options - StorageAccountsClientListSasTokensOptions contains the optional parameters for the StorageAccountsClient.NewListSasTokensPager
//     method.
func (client *StorageAccountsClient) NewListSasTokensPager(resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientListSasTokensOptions) *runtime.Pager[StorageAccountsClientListSasTokensResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListSasTokensResponse]{
		More: func(page StorageAccountsClientListSasTokensResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListSasTokensResponse) (StorageAccountsClientListSasTokensResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAccountsClient.NewListSasTokensPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSasTokensCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
			}, nil)
			if err != nil {
				return StorageAccountsClientListSasTokensResponse{}, err
			}
			return client.listSasTokensHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSasTokensCreateRequest creates the ListSasTokens request.
func (client *StorageAccountsClient) listSasTokensCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientListSasTokensOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}/listSasTokens"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSasTokensHandleResponse handles the ListSasTokens response.
func (client *StorageAccountsClient) listSasTokensHandleResponse(resp *http.Response) (StorageAccountsClientListSasTokensResponse, error) {
	result := StorageAccountsClientListSasTokensResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasTokenInformationListResult); err != nil {
		return StorageAccountsClientListSasTokensResponse{}, err
	}
	return result, nil
}

// NewListStorageContainersPager - Lists the Azure Storage containers, if any, associated with the specified Data Lake Analytics
// and Azure Storage account combination. The response includes a link to the next page of results, if any.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The name of the Azure storage account from which to list blob containers.
//   - options - StorageAccountsClientListStorageContainersOptions contains the optional parameters for the StorageAccountsClient.NewListStorageContainersPager
//     method.
func (client *StorageAccountsClient) NewListStorageContainersPager(resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientListStorageContainersOptions) *runtime.Pager[StorageAccountsClientListStorageContainersResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListStorageContainersResponse]{
		More: func(page StorageAccountsClientListStorageContainersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListStorageContainersResponse) (StorageAccountsClientListStorageContainersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAccountsClient.NewListStorageContainersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listStorageContainersCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
			}, nil)
			if err != nil {
				return StorageAccountsClientListStorageContainersResponse{}, err
			}
			return client.listStorageContainersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listStorageContainersCreateRequest creates the ListStorageContainers request.
func (client *StorageAccountsClient) listStorageContainersCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientListStorageContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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

// listStorageContainersHandleResponse handles the ListStorageContainers response.
func (client *StorageAccountsClient) listStorageContainersHandleResponse(resp *http.Response) (StorageAccountsClientListStorageContainersResponse, error) {
	result := StorageAccountsClientListStorageContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainerListResult); err != nil {
		return StorageAccountsClientListStorageContainersResponse{}, err
	}
	return result, nil
}

// Update - Updates the Data Lake Analytics account to replace Azure Storage blob account details, such as the access key
// and/or suffix.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Analytics account.
//   - storageAccountName - The Azure Storage account to modify
//   - options - StorageAccountsClientUpdateOptions contains the optional parameters for the StorageAccountsClient.Update method.
func (client *StorageAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientUpdateOptions) (StorageAccountsClientUpdateResponse, error) {
	var err error
	const operationName = "StorageAccountsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountsClientUpdateResponse{}, err
	}
	return StorageAccountsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *StorageAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
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
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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
