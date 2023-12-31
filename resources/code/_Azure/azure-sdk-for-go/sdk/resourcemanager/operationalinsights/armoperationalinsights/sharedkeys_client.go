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

// SharedKeysClient contains the methods for the SharedKeys group.
// Don't use this type directly, use NewSharedKeysClient() instead.
type SharedKeysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharedKeysClient creates a new instance of SharedKeysClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharedKeysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedKeysClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharedKeysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetSharedKeys - Gets the shared keys for a workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - SharedKeysClientGetSharedKeysOptions contains the optional parameters for the SharedKeysClient.GetSharedKeys
//     method.
func (client *SharedKeysClient) GetSharedKeys(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysClientGetSharedKeysOptions) (SharedKeysClientGetSharedKeysResponse, error) {
	var err error
	const operationName = "SharedKeysClient.GetSharedKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSharedKeysCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SharedKeysClientGetSharedKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharedKeysClientGetSharedKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharedKeysClientGetSharedKeysResponse{}, err
	}
	resp, err := client.getSharedKeysHandleResponse(httpResp)
	return resp, err
}

// getSharedKeysCreateRequest creates the GetSharedKeys request.
func (client *SharedKeysClient) getSharedKeysCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysClientGetSharedKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/sharedKeys"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSharedKeysHandleResponse handles the GetSharedKeys response.
func (client *SharedKeysClient) getSharedKeysHandleResponse(resp *http.Response) (SharedKeysClientGetSharedKeysResponse, error) {
	result := SharedKeysClientGetSharedKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedKeys); err != nil {
		return SharedKeysClientGetSharedKeysResponse{}, err
	}
	return result, nil
}

// Regenerate - Regenerates the shared keys for a Log Analytics Workspace. These keys are used to connect Microsoft Operational
// Insights agents to the workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - SharedKeysClientRegenerateOptions contains the optional parameters for the SharedKeysClient.Regenerate method.
func (client *SharedKeysClient) Regenerate(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysClientRegenerateOptions) (SharedKeysClientRegenerateResponse, error) {
	var err error
	const operationName = "SharedKeysClient.Regenerate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SharedKeysClientRegenerateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharedKeysClientRegenerateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharedKeysClientRegenerateResponse{}, err
	}
	resp, err := client.regenerateHandleResponse(httpResp)
	return resp, err
}

// regenerateCreateRequest creates the Regenerate request.
func (client *SharedKeysClient) regenerateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SharedKeysClientRegenerateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/regenerateSharedKey"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// regenerateHandleResponse handles the Regenerate response.
func (client *SharedKeysClient) regenerateHandleResponse(resp *http.Response) (SharedKeysClientRegenerateResponse, error) {
	result := SharedKeysClientRegenerateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedKeys); err != nil {
		return SharedKeysClientRegenerateResponse{}, err
	}
	return result, nil
}
