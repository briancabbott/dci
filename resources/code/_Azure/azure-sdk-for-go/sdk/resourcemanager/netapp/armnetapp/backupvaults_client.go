//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// BackupVaultsClient contains the methods for the BackupVaults group.
// Don't use this type directly, use NewBackupVaultsClient() instead.
type BackupVaultsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupVaultsClient creates a new instance of BackupVaultsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupVaultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupVaultsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupVaultsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the specified Backup Vault in the NetApp account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - body - BackupVault object supplied in the body of the operation.
//   - options - BackupVaultsClientBeginCreateOrUpdateOptions contains the optional parameters for the BackupVaultsClient.BeginCreateOrUpdate
//     method.
func (client *BackupVaultsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVault, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*runtime.Poller[BackupVaultsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, accountName, backupVaultName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupVaultsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupVaultsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update the specified Backup Vault in the NetApp account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *BackupVaultsClient) createOrUpdate(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVault, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupVaultsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, body, options)
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
func (client *BackupVaultsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVault, options *BackupVaultsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}"
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
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the specified Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - options - BackupVaultsClientBeginDeleteOptions contains the optional parameters for the BackupVaultsClient.BeginDelete
//     method.
func (client *BackupVaultsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupVaultsClientBeginDeleteOptions) (*runtime.Poller[BackupVaultsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, backupVaultName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupVaultsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupVaultsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the specified Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *BackupVaultsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupVaultsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupVaultsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, options)
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
func (client *BackupVaultsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupVaultsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}"
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
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - options - BackupVaultsClientGetOptions contains the optional parameters for the BackupVaultsClient.Get method.
func (client *BackupVaultsClient) Get(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupVaultsClientGetOptions) (BackupVaultsClientGetResponse, error) {
	var err error
	const operationName = "BackupVaultsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, options)
	if err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupVaultsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BackupVaultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, options *BackupVaultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}"
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
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupVaultsClient) getHandleResponse(resp *http.Response) (BackupVaultsClientGetResponse, error) {
	result := BackupVaultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupVault); err != nil {
		return BackupVaultsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNetAppAccountPager - List and describe all Backup Vaults in the NetApp account.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - options - BackupVaultsClientListByNetAppAccountOptions contains the optional parameters for the BackupVaultsClient.NewListByNetAppAccountPager
//     method.
func (client *BackupVaultsClient) NewListByNetAppAccountPager(resourceGroupName string, accountName string, options *BackupVaultsClientListByNetAppAccountOptions) *runtime.Pager[BackupVaultsClientListByNetAppAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupVaultsClientListByNetAppAccountResponse]{
		More: func(page BackupVaultsClientListByNetAppAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupVaultsClientListByNetAppAccountResponse) (BackupVaultsClientListByNetAppAccountResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupVaultsClient.NewListByNetAppAccountPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNetAppAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return BackupVaultsClientListByNetAppAccountResponse{}, err
			}
			return client.listByNetAppAccountHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNetAppAccountCreateRequest creates the ListByNetAppAccount request.
func (client *BackupVaultsClient) listByNetAppAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BackupVaultsClientListByNetAppAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNetAppAccountHandleResponse handles the ListByNetAppAccount response.
func (client *BackupVaultsClient) listByNetAppAccountHandleResponse(resp *http.Response) (BackupVaultsClientListByNetAppAccountResponse, error) {
	result := BackupVaultsClientListByNetAppAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupVaultsList); err != nil {
		return BackupVaultsClientListByNetAppAccountResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch the specified NetApp Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - backupVaultName - The name of the Backup Vault
//   - body - Backup Vault object supplied in the body of the operation.
//   - options - BackupVaultsClientBeginUpdateOptions contains the optional parameters for the BackupVaultsClient.BeginUpdate
//     method.
func (client *BackupVaultsClient) BeginUpdate(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVaultPatch, options *BackupVaultsClientBeginUpdateOptions) (*runtime.Poller[BackupVaultsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, accountName, backupVaultName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupVaultsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupVaultsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch the specified NetApp Backup Vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *BackupVaultsClient) update(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVaultPatch, options *BackupVaultsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupVaultsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, backupVaultName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *BackupVaultsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, body BackupVaultPatch, options *BackupVaultsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/backupVaults/{backupVaultName}"
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
	if backupVaultName == "" {
		return nil, errors.New("parameter backupVaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupVaultName}", url.PathEscape(backupVaultName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
