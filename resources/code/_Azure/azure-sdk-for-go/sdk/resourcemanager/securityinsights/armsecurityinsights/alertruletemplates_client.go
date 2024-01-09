//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// AlertRuleTemplatesClient contains the methods for the AlertRuleTemplates group.
// Don't use this type directly, use NewAlertRuleTemplatesClient() instead.
type AlertRuleTemplatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAlertRuleTemplatesClient creates a new instance of AlertRuleTemplatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertRuleTemplatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertRuleTemplatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertRuleTemplatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the alert rule template.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - alertRuleTemplateID - Alert rule template ID
//   - options - AlertRuleTemplatesClientGetOptions contains the optional parameters for the AlertRuleTemplatesClient.Get method.
func (client *AlertRuleTemplatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string, options *AlertRuleTemplatesClientGetOptions) (AlertRuleTemplatesClientGetResponse, error) {
	var err error
	const operationName = "AlertRuleTemplatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, alertRuleTemplateID, options)
	if err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AlertRuleTemplatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string, options *AlertRuleTemplatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRuleTemplates/{alertRuleTemplateId}"
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
	if alertRuleTemplateID == "" {
		return nil, errors.New("parameter alertRuleTemplateID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{alertRuleTemplateId}", url.PathEscape(alertRuleTemplateID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertRuleTemplatesClient) getHandleResponse(resp *http.Response) (AlertRuleTemplatesClientGetResponse, error) {
	result := AlertRuleTemplatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return AlertRuleTemplatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all alert rule templates.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - AlertRuleTemplatesClientListOptions contains the optional parameters for the AlertRuleTemplatesClient.NewListPager
//     method.
func (client *AlertRuleTemplatesClient) NewListPager(resourceGroupName string, workspaceName string, options *AlertRuleTemplatesClientListOptions) *runtime.Pager[AlertRuleTemplatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AlertRuleTemplatesClientListResponse]{
		More: func(page AlertRuleTemplatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertRuleTemplatesClientListResponse) (AlertRuleTemplatesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AlertRuleTemplatesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return AlertRuleTemplatesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AlertRuleTemplatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AlertRuleTemplatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRuleTemplates"
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
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertRuleTemplatesClient) listHandleResponse(resp *http.Response) (AlertRuleTemplatesClientListResponse, error) {
	result := AlertRuleTemplatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertRuleTemplatesList); err != nil {
		return AlertRuleTemplatesClientListResponse{}, err
	}
	return result, nil
}
