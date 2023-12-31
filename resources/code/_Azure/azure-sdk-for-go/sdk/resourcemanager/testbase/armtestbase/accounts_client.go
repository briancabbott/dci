//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// AccountsClient contains the methods for the TestBaseAccounts group.
// Don't use this type directly, use NewAccountsClient() instead.
type AccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccountsClient creates a new instance of AccountsClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckPackageNameAvailability - Checks that the Test Base Package name and version is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - parameters - Parameters supplied to the Test Base Package CheckNameAvailability operation.
//   - options - AccountsClientCheckPackageNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckPackageNameAvailability
//     method.
func (client *AccountsClient) CheckPackageNameAvailability(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters PackageCheckNameAvailabilityParameters, options *AccountsClientCheckPackageNameAvailabilityOptions) (AccountsClientCheckPackageNameAvailabilityResponse, error) {
	var err error
	const operationName = "AccountsClient.CheckPackageNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkPackageNameAvailabilityCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
	if err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	resp, err := client.checkPackageNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkPackageNameAvailabilityCreateRequest creates the CheckPackageNameAvailability request.
func (client *AccountsClient) checkPackageNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters PackageCheckNameAvailabilityParameters, options *AccountsClientCheckPackageNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/checkPackageNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkPackageNameAvailabilityHandleResponse handles the CheckPackageNameAvailability response.
func (client *AccountsClient) checkPackageNameAvailabilityHandleResponse(resp *http.Response) (AccountsClientCheckPackageNameAvailabilityResponse, error) {
	result := AccountsClientCheckPackageNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return AccountsClientCheckPackageNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Create or replace (overwrite/recreate, with potential downtime) a Test Base Account in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - parameters - Parameters supplied to create a Test Base Account.
//   - options - AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
func (client *AccountsClient) BeginCreate(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*runtime.Poller[AccountsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, testBaseAccountName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccountsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccountsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create or replace (overwrite/recreate, with potential downtime) a Test Base Account in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) create(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccountsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
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
func (client *AccountsClient) createCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountResource, options *AccountsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Restore != nil {
		reqQP.Set("restore", strconv.FormatBool(*options.Restore))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
func (client *AccountsClient) BeginDelete(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*runtime.Poller[AccountsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, testBaseAccountName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccountsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccountsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) deleteOperation(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AccountsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
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
func (client *AccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
func (client *AccountsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetOptions) (AccountsClientGetResponse, error) {
	var err error
	const operationName = "AccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountsClient) getHandleResponse(resp *http.Response) (AccountsClientGetResponse, error) {
	result := AccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountResource); err != nil {
		return AccountsClientGetResponse{}, err
	}
	return result, nil
}

// GetFileUploadURL - Gets the file upload URL of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - AccountsClientGetFileUploadURLOptions contains the optional parameters for the AccountsClient.GetFileUploadURL
//     method.
func (client *AccountsClient) GetFileUploadURL(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetFileUploadURLOptions) (AccountsClientGetFileUploadURLResponse, error) {
	var err error
	const operationName = "AccountsClient.GetFileUploadURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFileUploadURLCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
	if err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	resp, err := client.getFileUploadURLHandleResponse(httpResp)
	return resp, err
}

// getFileUploadURLCreateRequest creates the GetFileUploadURL request.
func (client *AccountsClient) getFileUploadURLCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientGetFileUploadURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/getFileUploadUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
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

// getFileUploadURLHandleResponse handles the GetFileUploadURL response.
func (client *AccountsClient) getFileUploadURLHandleResponse(resp *http.Response) (AccountsClientGetFileUploadURLResponse, error) {
	result := AccountsClientGetFileUploadURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FileUploadURLResponse); err != nil {
		return AccountsClientGetFileUploadURLResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the Test Base Accounts in a resource group.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - options - AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
//     method.
func (client *AccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *AccountsClientListByResourceGroupOptions) *runtime.Pager[AccountsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListByResourceGroupResponse]{
		More: func(page AccountsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListByResourceGroupResponse) (AccountsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccountsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AccountsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts"
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
	if options != nil && options.GetDeleted != nil {
		reqQP.Set("getDeleted", strconv.FormatBool(*options.GetDeleted))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (AccountsClientListByResourceGroupResponse, error) {
	result := AccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the Test Base Accounts in a subscription. This API is required by ARM guidelines.
//
// Generated from API version 2020-12-16-preview
//   - options - AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.NewListBySubscriptionPager
//     method.
func (client *AccountsClient) NewListBySubscriptionPager(options *AccountsClientListBySubscriptionOptions) *runtime.Pager[AccountsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccountsClientListBySubscriptionResponse]{
		More: func(page AccountsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccountsClientListBySubscriptionResponse) (AccountsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccountsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AccountsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.TestBase/testBaseAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.GetDeleted != nil {
		reqQP.Set("getDeleted", strconv.FormatBool(*options.GetDeleted))
	}
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (AccountsClientListBySubscriptionResponse, error) {
	result := AccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountListResult); err != nil {
		return AccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginOffboard - Offboard a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - AccountsClientBeginOffboardOptions contains the optional parameters for the AccountsClient.BeginOffboard method.
func (client *AccountsClient) BeginOffboard(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*runtime.Poller[AccountsClientOffboardResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.offboard(ctx, resourceGroupName, testBaseAccountName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccountsClientOffboardResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccountsClientOffboardResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Offboard - Offboard a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) offboard(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*http.Response, error) {
	var err error
	const operationName = "AccountsClient.BeginOffboard"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.offboardCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
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

// offboardCreateRequest creates the Offboard request.
func (client *AccountsClient) offboardCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *AccountsClientBeginOffboardOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/offboard"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update an existing Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - parameters - Parameters supplied to update a Test Base Account.
//   - options - AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
func (client *AccountsClient) BeginUpdate(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*runtime.Poller[AccountsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, testBaseAccountName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccountsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccountsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an existing Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-16-preview
func (client *AccountsClient) update(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccountsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, testBaseAccountName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, parameters AccountUpdateParameters, options *AccountsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-12-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
