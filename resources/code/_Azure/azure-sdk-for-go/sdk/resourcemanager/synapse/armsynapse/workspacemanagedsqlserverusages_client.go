//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspaceManagedSQLServerUsagesClient contains the methods for the WorkspaceManagedSQLServerUsages group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerUsagesClient() instead.
type WorkspaceManagedSQLServerUsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagedSQLServerUsagesClient creates a new instance of WorkspaceManagedSQLServerUsagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerUsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedSQLServerUsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get list of server usages metric for workspace managed sql server.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedSQLServerUsagesClientListOptions contains the optional parameters for the WorkspaceManagedSQLServerUsagesClient.NewListPager
//     method.
func (client *WorkspaceManagedSQLServerUsagesClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerUsagesClientListOptions) *runtime.Pager[WorkspaceManagedSQLServerUsagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerUsagesClientListResponse]{
		More: func(page WorkspaceManagedSQLServerUsagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerUsagesClientListResponse) (WorkspaceManagedSQLServerUsagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceManagedSQLServerUsagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return WorkspaceManagedSQLServerUsagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerUsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerUsagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlUsages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagedSQLServerUsagesClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerUsagesClientListResponse, error) {
	result := WorkspaceManagedSQLServerUsagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerUsageListResult); err != nil {
		return WorkspaceManagedSQLServerUsagesClientListResponse{}, err
	}
	return result, nil
}
