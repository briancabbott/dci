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
	"strconv"
	"strings"
)

// IncidentsClient contains the methods for the Incidents group.
// Don't use this type directly, use NewIncidentsClient() instead.
type IncidentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIncidentsClient creates a new instance of IncidentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIncidentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IncidentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IncidentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - incident - The incident
//   - options - IncidentsClientCreateOrUpdateOptions contains the optional parameters for the IncidentsClient.CreateOrUpdate
//     method.
func (client *IncidentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (IncidentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IncidentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incident, options)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident Incident, options *IncidentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, incident); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentsClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentsClientCreateOrUpdateResponse, error) {
	result := IncidentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a given incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientDeleteOptions contains the optional parameters for the IncidentsClient.Delete method.
func (client *IncidentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (IncidentsClientDeleteResponse, error) {
	var err error
	const operationName = "IncidentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientDeleteResponse{}, err
	}
	return IncidentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a given incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientGetOptions contains the optional parameters for the IncidentsClient.Get method.
func (client *IncidentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (IncidentsClientGetResponse, error) {
	var err error
	const operationName = "IncidentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IncidentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
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
func (client *IncidentsClient) getHandleResponse(resp *http.Response) (IncidentsClientGetResponse, error) {
	result := IncidentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Incident); err != nil {
		return IncidentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all incidents.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - IncidentsClientListOptions contains the optional parameters for the IncidentsClient.NewListPager method.
func (client *IncidentsClient) NewListPager(resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) *runtime.Pager[IncidentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IncidentsClientListResponse]{
		More: func(page IncidentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IncidentsClientListResponse) (IncidentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IncidentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return IncidentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IncidentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IncidentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentsClient) listHandleResponse(resp *http.Response) (IncidentsClientListResponse, error) {
	result := IncidentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentList); err != nil {
		return IncidentsClientListResponse{}, err
	}
	return result, nil
}

// ListAlerts - Gets all alerts for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListAlertsOptions contains the optional parameters for the IncidentsClient.ListAlerts method.
func (client *IncidentsClient) ListAlerts(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (IncidentsClientListAlertsResponse, error) {
	var err error
	const operationName = "IncidentsClient.ListAlerts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAlertsCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientListAlertsResponse{}, err
	}
	resp, err := client.listAlertsHandleResponse(httpResp)
	return resp, err
}

// listAlertsCreateRequest creates the ListAlerts request.
func (client *IncidentsClient) listAlertsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListAlertsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/alerts"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAlertsHandleResponse handles the ListAlerts response.
func (client *IncidentsClient) listAlertsHandleResponse(resp *http.Response) (IncidentsClientListAlertsResponse, error) {
	result := IncidentsClientListAlertsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentAlertList); err != nil {
		return IncidentsClientListAlertsResponse{}, err
	}
	return result, nil
}

// ListBookmarks - Gets all bookmarks for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListBookmarksOptions contains the optional parameters for the IncidentsClient.ListBookmarks method.
func (client *IncidentsClient) ListBookmarks(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (IncidentsClientListBookmarksResponse, error) {
	var err error
	const operationName = "IncidentsClient.ListBookmarks"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listBookmarksCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientListBookmarksResponse{}, err
	}
	resp, err := client.listBookmarksHandleResponse(httpResp)
	return resp, err
}

// listBookmarksCreateRequest creates the ListBookmarks request.
func (client *IncidentsClient) listBookmarksCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListBookmarksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/bookmarks"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBookmarksHandleResponse handles the ListBookmarks response.
func (client *IncidentsClient) listBookmarksHandleResponse(resp *http.Response) (IncidentsClientListBookmarksResponse, error) {
	result := IncidentsClientListBookmarksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentBookmarkList); err != nil {
		return IncidentsClientListBookmarksResponse{}, err
	}
	return result, nil
}

// ListEntities - Gets all entities for an incident.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentsClientListEntitiesOptions contains the optional parameters for the IncidentsClient.ListEntities method.
func (client *IncidentsClient) ListEntities(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (IncidentsClientListEntitiesResponse, error) {
	var err error
	const operationName = "IncidentsClient.ListEntities"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listEntitiesCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IncidentsClientListEntitiesResponse{}, err
	}
	resp, err := client.listEntitiesHandleResponse(httpResp)
	return resp, err
}

// listEntitiesCreateRequest creates the ListEntities request.
func (client *IncidentsClient) listEntitiesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentsClientListEntitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/entities"
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
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEntitiesHandleResponse handles the ListEntities response.
func (client *IncidentsClient) listEntitiesHandleResponse(resp *http.Response) (IncidentsClientListEntitiesResponse, error) {
	result := IncidentsClientListEntitiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentEntitiesResponse); err != nil {
		return IncidentsClientListEntitiesResponse{}, err
	}
	return result, nil
}
