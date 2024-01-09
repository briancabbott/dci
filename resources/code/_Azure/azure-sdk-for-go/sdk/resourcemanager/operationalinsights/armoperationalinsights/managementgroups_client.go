//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// ManagementGroupsClient contains the methods for the ManagementGroups group.
// Don't use this type directly, use NewManagementGroupsClient() instead.
type ManagementGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementGroupsClient creates a new instance of ManagementGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets a list of management groups connected to a workspace.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - ManagementGroupsClientListOptions contains the optional parameters for the ManagementGroupsClient.NewListPager
//     method.
func (client *ManagementGroupsClient) NewListPager(resourceGroupName string, workspaceName string, options *ManagementGroupsClientListOptions) *runtime.Pager[ManagementGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagementGroupsClientListResponse]{
		More: func(page ManagementGroupsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ManagementGroupsClientListResponse) (ManagementGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagementGroupsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			if err != nil {
				return ManagementGroupsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagementGroupsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagementGroupsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagementGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ManagementGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/managementGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagementGroupsClient) listHandleResponse(resp *http.Response) (ManagementGroupsClientListResponse, error) {
	result := ManagementGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListManagementGroupsResult); err != nil {
		return ManagementGroupsClientListResponse{}, err
	}
	return result, nil
}
