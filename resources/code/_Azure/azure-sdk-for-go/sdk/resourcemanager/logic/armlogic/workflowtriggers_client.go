//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// WorkflowTriggersClient contains the methods for the WorkflowTriggers group.
// Don't use this type directly, use NewWorkflowTriggersClient() instead.
type WorkflowTriggersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkflowTriggersClient creates a new instance of WorkflowTriggersClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkflowTriggersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkflowTriggersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkflowTriggersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a workflow trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - options - WorkflowTriggersClientGetOptions contains the optional parameters for the WorkflowTriggersClient.Get method.
func (client *WorkflowTriggersClient) Get(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientGetOptions) (WorkflowTriggersClientGetResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
	if err != nil {
		return WorkflowTriggersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkflowTriggersClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowTriggersClient) getHandleResponse(resp *http.Response) (WorkflowTriggersClientGetResponse, error) {
	result := WorkflowTriggersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTrigger); err != nil {
		return WorkflowTriggersClientGetResponse{}, err
	}
	return result, nil
}

// GetSchemaJSON - Get the trigger schema as JSON.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - options - WorkflowTriggersClientGetSchemaJSONOptions contains the optional parameters for the WorkflowTriggersClient.GetSchemaJSON
//     method.
func (client *WorkflowTriggersClient) GetSchemaJSON(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientGetSchemaJSONOptions) (WorkflowTriggersClientGetSchemaJSONResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.GetSchemaJSON"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getSchemaJSONCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
	if err != nil {
		return WorkflowTriggersClientGetSchemaJSONResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientGetSchemaJSONResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientGetSchemaJSONResponse{}, err
	}
	resp, err := client.getSchemaJSONHandleResponse(httpResp)
	return resp, err
}

// getSchemaJSONCreateRequest creates the GetSchemaJSON request.
func (client *WorkflowTriggersClient) getSchemaJSONCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientGetSchemaJSONOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/schemas/json"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSchemaJSONHandleResponse handles the GetSchemaJSON response.
func (client *WorkflowTriggersClient) getSchemaJSONHandleResponse(resp *http.Response) (WorkflowTriggersClientGetSchemaJSONResponse, error) {
	result := WorkflowTriggersClientGetSchemaJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JSONSchema); err != nil {
		return WorkflowTriggersClientGetSchemaJSONResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of workflow triggers.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - options - WorkflowTriggersClientListOptions contains the optional parameters for the WorkflowTriggersClient.NewListPager
//     method.
func (client *WorkflowTriggersClient) NewListPager(resourceGroupName string, workflowName string, options *WorkflowTriggersClientListOptions) *runtime.Pager[WorkflowTriggersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkflowTriggersClientListResponse]{
		More: func(page WorkflowTriggersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkflowTriggersClientListResponse) (WorkflowTriggersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkflowTriggersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workflowName, options)
			}, nil)
			if err != nil {
				return WorkflowTriggersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkflowTriggersClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowTriggersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowTriggersClient) listHandleResponse(resp *http.Response) (WorkflowTriggersClientListResponse, error) {
	result := WorkflowTriggersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerListResult); err != nil {
		return WorkflowTriggersClientListResponse{}, err
	}
	return result, nil
}

// ListCallbackURL - Get the callback URL for a workflow trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - options - WorkflowTriggersClientListCallbackURLOptions contains the optional parameters for the WorkflowTriggersClient.ListCallbackURL
//     method.
func (client *WorkflowTriggersClient) ListCallbackURL(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientListCallbackURLOptions) (WorkflowTriggersClientListCallbackURLResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.ListCallbackURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCallbackURLCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
	if err != nil {
		return WorkflowTriggersClientListCallbackURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientListCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientListCallbackURLResponse{}, err
	}
	resp, err := client.listCallbackURLHandleResponse(httpResp)
	return resp, err
}

// listCallbackURLCreateRequest creates the ListCallbackURL request.
func (client *WorkflowTriggersClient) listCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientListCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/listCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCallbackURLHandleResponse handles the ListCallbackURL response.
func (client *WorkflowTriggersClient) listCallbackURLHandleResponse(resp *http.Response) (WorkflowTriggersClientListCallbackURLResponse, error) {
	result := WorkflowTriggersClientListCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return WorkflowTriggersClientListCallbackURLResponse{}, err
	}
	return result, nil
}

// Reset - Resets a workflow trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - options - WorkflowTriggersClientResetOptions contains the optional parameters for the WorkflowTriggersClient.Reset method.
func (client *WorkflowTriggersClient) Reset(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientResetOptions) (WorkflowTriggersClientResetResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.Reset"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
	if err != nil {
		return WorkflowTriggersClientResetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientResetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientResetResponse{}, err
	}
	return WorkflowTriggersClientResetResponse{}, nil
}

// resetCreateRequest creates the Reset request.
func (client *WorkflowTriggersClient) resetCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientResetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/reset"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Run - Runs a workflow trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - options - WorkflowTriggersClientRunOptions contains the optional parameters for the WorkflowTriggersClient.Run method.
func (client *WorkflowTriggersClient) Run(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientRunOptions) (WorkflowTriggersClientRunResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.Run"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.runCreateRequest(ctx, resourceGroupName, workflowName, triggerName, options)
	if err != nil {
		return WorkflowTriggersClientRunResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientRunResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientRunResponse{}, err
	}
	return WorkflowTriggersClientRunResponse{}, nil
}

// runCreateRequest creates the Run request.
func (client *WorkflowTriggersClient) runCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, options *WorkflowTriggersClientRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/run"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// SetState - Sets the state of a workflow trigger.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The resource group name.
//   - workflowName - The workflow name.
//   - triggerName - The workflow trigger name.
//   - setState - The workflow trigger state.
//   - options - WorkflowTriggersClientSetStateOptions contains the optional parameters for the WorkflowTriggersClient.SetState
//     method.
func (client *WorkflowTriggersClient) SetState(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState SetTriggerStateActionDefinition, options *WorkflowTriggersClientSetStateOptions) (WorkflowTriggersClientSetStateResponse, error) {
	var err error
	const operationName = "WorkflowTriggersClient.SetState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.setStateCreateRequest(ctx, resourceGroupName, workflowName, triggerName, setState, options)
	if err != nil {
		return WorkflowTriggersClientSetStateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkflowTriggersClientSetStateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkflowTriggersClientSetStateResponse{}, err
	}
	return WorkflowTriggersClientSetStateResponse{}, nil
}

// setStateCreateRequest creates the SetState request.
func (client *WorkflowTriggersClient) setStateCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, triggerName string, setState SetTriggerStateActionDefinition, options *WorkflowTriggersClientSetStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/triggers/{triggerName}/setState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, setState); err != nil {
		return nil, err
	}
	return req, nil
}
