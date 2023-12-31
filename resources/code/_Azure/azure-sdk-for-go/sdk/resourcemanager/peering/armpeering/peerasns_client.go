//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// PeerAsnsClient contains the methods for the PeerAsns group.
// Don't use this type directly, use NewPeerAsnsClient() instead.
type PeerAsnsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPeerAsnsClient creates a new instance of PeerAsnsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPeerAsnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PeerAsnsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PeerAsnsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new peer ASN or updates an existing peer ASN with the specified name under the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - peerAsnName - The peer ASN name.
//   - peerAsn - The peer ASN.
//   - options - PeerAsnsClientCreateOrUpdateOptions contains the optional parameters for the PeerAsnsClient.CreateOrUpdate method.
func (client *PeerAsnsClient) CreateOrUpdate(ctx context.Context, peerAsnName string, peerAsn PeerAsn, options *PeerAsnsClientCreateOrUpdateOptions) (PeerAsnsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PeerAsnsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, peerAsnName, peerAsn, options)
	if err != nil {
		return PeerAsnsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PeerAsnsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PeerAsnsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PeerAsnsClient) createOrUpdateCreateRequest(ctx context.Context, peerAsnName string, peerAsn PeerAsn, options *PeerAsnsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}"
	if peerAsnName == "" {
		return nil, errors.New("parameter peerAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peerAsnName}", url.PathEscape(peerAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, peerAsn); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PeerAsnsClient) createOrUpdateHandleResponse(resp *http.Response) (PeerAsnsClientCreateOrUpdateResponse, error) {
	result := PeerAsnsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeerAsn); err != nil {
		return PeerAsnsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing peer ASN with the specified name under the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - peerAsnName - The peer ASN name.
//   - options - PeerAsnsClientDeleteOptions contains the optional parameters for the PeerAsnsClient.Delete method.
func (client *PeerAsnsClient) Delete(ctx context.Context, peerAsnName string, options *PeerAsnsClientDeleteOptions) (PeerAsnsClientDeleteResponse, error) {
	var err error
	const operationName = "PeerAsnsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, peerAsnName, options)
	if err != nil {
		return PeerAsnsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PeerAsnsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PeerAsnsClientDeleteResponse{}, err
	}
	return PeerAsnsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PeerAsnsClient) deleteCreateRequest(ctx context.Context, peerAsnName string, options *PeerAsnsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}"
	if peerAsnName == "" {
		return nil, errors.New("parameter peerAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peerAsnName}", url.PathEscape(peerAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the peer ASN with the specified name under the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - peerAsnName - The peer ASN name.
//   - options - PeerAsnsClientGetOptions contains the optional parameters for the PeerAsnsClient.Get method.
func (client *PeerAsnsClient) Get(ctx context.Context, peerAsnName string, options *PeerAsnsClientGetOptions) (PeerAsnsClientGetResponse, error) {
	var err error
	const operationName = "PeerAsnsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, peerAsnName, options)
	if err != nil {
		return PeerAsnsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PeerAsnsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PeerAsnsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PeerAsnsClient) getCreateRequest(ctx context.Context, peerAsnName string, options *PeerAsnsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns/{peerAsnName}"
	if peerAsnName == "" {
		return nil, errors.New("parameter peerAsnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peerAsnName}", url.PathEscape(peerAsnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeerAsnsClient) getHandleResponse(resp *http.Response) (PeerAsnsClientGetResponse, error) {
	result := PeerAsnsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeerAsn); err != nil {
		return PeerAsnsClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the peer ASNs under the given subscription.
//
// Generated from API version 2022-01-01
//   - options - PeerAsnsClientListBySubscriptionOptions contains the optional parameters for the PeerAsnsClient.NewListBySubscriptionPager
//     method.
func (client *PeerAsnsClient) NewListBySubscriptionPager(options *PeerAsnsClientListBySubscriptionOptions) *runtime.Pager[PeerAsnsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PeerAsnsClientListBySubscriptionResponse]{
		More: func(page PeerAsnsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PeerAsnsClientListBySubscriptionResponse) (PeerAsnsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PeerAsnsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PeerAsnsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PeerAsnsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PeerAsnsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerAsns"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PeerAsnsClient) listBySubscriptionHandleResponse(resp *http.Response) (PeerAsnsClientListBySubscriptionResponse, error) {
	result := PeerAsnsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeerAsnListResult); err != nil {
		return PeerAsnsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
