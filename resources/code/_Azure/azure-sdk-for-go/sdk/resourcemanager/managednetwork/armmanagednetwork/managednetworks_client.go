//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

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

// ManagedNetworksClient contains the methods for the ManagedNetworks group.
// Don't use this type directly, use NewManagedNetworksClient() instead.
type ManagedNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedNetworksClient creates a new instance of ManagedNetworksClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedNetworksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - The Put ManagedNetworks operation creates/updates a Managed Network Resource, specified by resource group
// and Managed Network name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - managedNetwork - Parameters supplied to the create/update a Managed Network Resource
//   - options - ManagedNetworksClientCreateOrUpdateOptions contains the optional parameters for the ManagedNetworksClient.CreateOrUpdate
//     method.
func (client *ManagedNetworksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetwork ManagedNetwork, options *ManagedNetworksClientCreateOrUpdateOptions) (ManagedNetworksClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ManagedNetworksClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedNetworkName, managedNetwork, options)
	if err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, managedNetwork ManagedNetwork, options *ManagedNetworksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, managedNetwork); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedNetworksClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedNetworksClientCreateOrUpdateResponse, error) {
	result := ManagedNetworksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedNetwork); err != nil {
		return ManagedNetworksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - The Delete ManagedNetworks operation deletes a Managed Network Resource, specified by the resource group
// and Managed Network name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - options - ManagedNetworksClientBeginDeleteOptions contains the optional parameters for the ManagedNetworksClient.BeginDelete
//     method.
func (client *ManagedNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*runtime.Poller[ManagedNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, managedNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedNetworksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The Delete ManagedNetworks operation deletes a Managed Network Resource, specified by the resource group and Managed
// Network name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *ManagedNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedNetworksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
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
func (client *ManagedNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The Get ManagedNetworks operation gets a Managed Network Resource, specified by the resource group and Managed Network
// name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - options - ManagedNetworksClientGetOptions contains the optional parameters for the ManagedNetworksClient.Get method.
func (client *ManagedNetworksClient) Get(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientGetOptions) (ManagedNetworksClientGetResponse, error) {
	var err error
	const operationName = "ManagedNetworksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedNetworkName, options)
	if err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, options *ManagedNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedNetworksClient) getHandleResponse(resp *http.Response) (ManagedNetworksClientGetResponse, error) {
	result := ManagedNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedNetwork); err != nil {
		return ManagedNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The ListByResourceGroup ManagedNetwork operation retrieves all the Managed Network resources
// in a resource group in a paginated format.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - options - ManagedNetworksClientListByResourceGroupOptions contains the optional parameters for the ManagedNetworksClient.NewListByResourceGroupPager
//     method.
func (client *ManagedNetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *ManagedNetworksClientListByResourceGroupOptions) *runtime.Pager[ManagedNetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedNetworksClientListByResourceGroupResponse]{
		More: func(page ManagedNetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedNetworksClientListByResourceGroupResponse) (ManagedNetworksClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedNetworksClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ManagedNetworksClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ManagedNetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ManagedNetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks"
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
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ManagedNetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (ManagedNetworksClientListByResourceGroupResponse, error) {
	result := ManagedNetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ManagedNetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - The ListBySubscription ManagedNetwork operation retrieves all the Managed Network Resources
// in the current subscription in a paginated format.
//
// Generated from API version 2019-06-01-preview
//   - options - ManagedNetworksClientListBySubscriptionOptions contains the optional parameters for the ManagedNetworksClient.NewListBySubscriptionPager
//     method.
func (client *ManagedNetworksClient) NewListBySubscriptionPager(options *ManagedNetworksClientListBySubscriptionOptions) *runtime.Pager[ManagedNetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedNetworksClientListBySubscriptionResponse]{
		More: func(page ManagedNetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedNetworksClientListBySubscriptionResponse) (ManagedNetworksClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedNetworksClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ManagedNetworksClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagedNetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagedNetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetwork/managedNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagedNetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagedNetworksClientListBySubscriptionResponse, error) {
	result := ManagedNetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return ManagedNetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the specified Managed Network resource tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group.
//   - managedNetworkName - The name of the Managed Network.
//   - parameters - Parameters supplied to update application gateway tags and/or scope.
//   - options - ManagedNetworksClientBeginUpdateOptions contains the optional parameters for the ManagedNetworksClient.BeginUpdate
//     method.
func (client *ManagedNetworksClient) BeginUpdate(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*runtime.Poller[ManagedNetworksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, managedNetworkName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ManagedNetworksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ManagedNetworksClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the specified Managed Network resource tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *ManagedNetworksClient) update(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ManagedNetworksClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, managedNetworkName, parameters, options)
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
func (client *ManagedNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, managedNetworkName string, parameters Update, options *ManagedNetworksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetwork/managedNetworks/{managedNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedNetworkName == "" {
		return nil, errors.New("parameter managedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedNetworkName}", url.PathEscape(managedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
