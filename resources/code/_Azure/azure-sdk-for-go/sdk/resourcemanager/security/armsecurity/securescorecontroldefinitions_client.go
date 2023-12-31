//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// SecureScoreControlDefinitionsClient contains the methods for the SecureScoreControlDefinitions group.
// Don't use this type directly, use NewSecureScoreControlDefinitionsClient() instead.
type SecureScoreControlDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSecureScoreControlDefinitionsClient creates a new instance of SecureScoreControlDefinitionsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSecureScoreControlDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SecureScoreControlDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SecureScoreControlDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List the available security controls, their assessments, and the max score
//
// Generated from API version 2020-01-01
//   - options - SecureScoreControlDefinitionsClientListOptions contains the optional parameters for the SecureScoreControlDefinitionsClient.NewListPager
//     method.
func (client *SecureScoreControlDefinitionsClient) NewListPager(options *SecureScoreControlDefinitionsClientListOptions) *runtime.Pager[SecureScoreControlDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecureScoreControlDefinitionsClientListResponse]{
		More: func(page SecureScoreControlDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecureScoreControlDefinitionsClientListResponse) (SecureScoreControlDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecureScoreControlDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SecureScoreControlDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SecureScoreControlDefinitionsClient) listCreateRequest(ctx context.Context, options *SecureScoreControlDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Security/secureScoreControlDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoreControlDefinitionsClient) listHandleResponse(resp *http.Response) (SecureScoreControlDefinitionsClientListResponse, error) {
	result := SecureScoreControlDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlDefinitionList); err != nil {
		return SecureScoreControlDefinitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - For a specified subscription, list the available security controls, their assessments, and
// the max score
//
// Generated from API version 2020-01-01
//   - options - SecureScoreControlDefinitionsClientListBySubscriptionOptions contains the optional parameters for the SecureScoreControlDefinitionsClient.NewListBySubscriptionPager
//     method.
func (client *SecureScoreControlDefinitionsClient) NewListBySubscriptionPager(options *SecureScoreControlDefinitionsClientListBySubscriptionOptions) *runtime.Pager[SecureScoreControlDefinitionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SecureScoreControlDefinitionsClientListBySubscriptionResponse]{
		More: func(page SecureScoreControlDefinitionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SecureScoreControlDefinitionsClientListBySubscriptionResponse) (SecureScoreControlDefinitionsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SecureScoreControlDefinitionsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SecureScoreControlDefinitionsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SecureScoreControlDefinitionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SecureScoreControlDefinitionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControlDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SecureScoreControlDefinitionsClient) listBySubscriptionHandleResponse(resp *http.Response) (SecureScoreControlDefinitionsClientListBySubscriptionResponse, error) {
	result := SecureScoreControlDefinitionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlDefinitionList); err != nil {
		return SecureScoreControlDefinitionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
