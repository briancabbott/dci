//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

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

// MetadataClient contains the methods for the Metadata group.
// Don't use this type directly, use NewMetadataClient() instead.
type MetadataClient struct {
	internal *arm.Client
}

// NewMetadataClient creates a new instance of MetadataClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MetadataClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetadataClient{
		internal: cl,
	}
	return client, nil
}

// GetEntity - Gets the list of metadata entities.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - name - Name of metadata entity.
//   - options - MetadataClientGetEntityOptions contains the optional parameters for the MetadataClient.GetEntity method.
func (client *MetadataClient) GetEntity(ctx context.Context, name string, options *MetadataClientGetEntityOptions) (MetadataClientGetEntityResponse, error) {
	var err error
	const operationName = "MetadataClient.GetEntity"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityCreateRequest(ctx, name, options)
	if err != nil {
		return MetadataClientGetEntityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataClientGetEntityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetadataClientGetEntityResponse{}, err
	}
	resp, err := client.getEntityHandleResponse(httpResp)
	return resp, err
}

// getEntityCreateRequest creates the GetEntity request.
func (client *MetadataClient) getEntityCreateRequest(ctx context.Context, name string, options *MetadataClientGetEntityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/metadata/{name}"
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityHandleResponse handles the GetEntity response.
func (client *MetadataClient) getEntityHandleResponse(resp *http.Response) (MetadataClientGetEntityResponse, error) {
	result := MetadataClientGetEntityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataEntity); err != nil {
		return MetadataClientGetEntityResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of metadata entities.
//
// Generated from API version 2023-10-01-preview
//   - options - MetadataClientListOptions contains the optional parameters for the MetadataClient.NewListPager method.
func (client *MetadataClient) NewListPager(options *MetadataClientListOptions) *runtime.Pager[MetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetadataClientListResponse]{
		More: func(page MetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MetadataClientListResponse) (MetadataClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetadataClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MetadataClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MetadataClient) listCreateRequest(ctx context.Context, options *MetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/metadata"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetadataClient) listHandleResponse(resp *http.Response) (MetadataClientListResponse, error) {
	result := MetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataEntityListResult); err != nil {
		return MetadataClientListResponse{}, err
	}
	return result, nil
}
