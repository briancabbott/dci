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

// GovernanceAssignmentsClient contains the methods for the GovernanceAssignments group.
// Don't use this type directly, use NewGovernanceAssignmentsClient() instead.
type GovernanceAssignmentsClient struct {
	internal *arm.Client
}

// NewGovernanceAssignmentsClient creates a new instance of GovernanceAssignmentsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGovernanceAssignmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GovernanceAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GovernanceAssignmentsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a governance assignment on the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format:
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - assessmentName - The Assessment Key - A unique key for the assessment type
//   - assignmentKey - The governance assignment key - the assessment key of the required governance assignment
//   - governanceAssignment - Governance assignment over a subscription scope
//   - options - GovernanceAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the GovernanceAssignmentsClient.CreateOrUpdate
//     method.
func (client *GovernanceAssignmentsClient) CreateOrUpdate(ctx context.Context, scope string, assessmentName string, assignmentKey string, governanceAssignment GovernanceAssignment, options *GovernanceAssignmentsClientCreateOrUpdateOptions) (GovernanceAssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "GovernanceAssignmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, scope, assessmentName, assignmentKey, governanceAssignment, options)
	if err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GovernanceAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, governanceAssignment GovernanceAssignment, options *GovernanceAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, governanceAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GovernanceAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (GovernanceAssignmentsClientCreateOrUpdateResponse, error) {
	result := GovernanceAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignment); err != nil {
		return GovernanceAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a GovernanceAssignment over a given scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format:
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - assessmentName - The Assessment Key - A unique key for the assessment type
//   - assignmentKey - The governance assignment key - the assessment key of the required governance assignment
//   - options - GovernanceAssignmentsClientDeleteOptions contains the optional parameters for the GovernanceAssignmentsClient.Delete
//     method.
func (client *GovernanceAssignmentsClient) Delete(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientDeleteOptions) (GovernanceAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "GovernanceAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, scope, assessmentName, assignmentKey, options)
	if err != nil {
		return GovernanceAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GovernanceAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GovernanceAssignmentsClientDeleteResponse{}, err
	}
	return GovernanceAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GovernanceAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific governanceAssignment for the requested scope by AssignmentKey
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format:
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - assessmentName - The Assessment Key - A unique key for the assessment type
//   - assignmentKey - The governance assignment key - the assessment key of the required governance assignment
//   - options - GovernanceAssignmentsClientGetOptions contains the optional parameters for the GovernanceAssignmentsClient.Get
//     method.
func (client *GovernanceAssignmentsClient) Get(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientGetOptions) (GovernanceAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "GovernanceAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, assessmentName, assignmentKey, options)
	if err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GovernanceAssignmentsClient) getCreateRequest(ctx context.Context, scope string, assessmentName string, assignmentKey string, options *GovernanceAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments/{assignmentKey}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if assignmentKey == "" {
		return nil, errors.New("parameter assignmentKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assignmentKey}", url.PathEscape(assignmentKey))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GovernanceAssignmentsClient) getHandleResponse(resp *http.Response) (GovernanceAssignmentsClientGetResponse, error) {
	result := GovernanceAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignment); err != nil {
		return GovernanceAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get governance assignments on all of your resources inside a scope
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the Governance assignments. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'),
//     or security connector (format:
//     'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//   - assessmentName - The Assessment Key - A unique key for the assessment type
//   - options - GovernanceAssignmentsClientListOptions contains the optional parameters for the GovernanceAssignmentsClient.NewListPager
//     method.
func (client *GovernanceAssignmentsClient) NewListPager(scope string, assessmentName string, options *GovernanceAssignmentsClientListOptions) *runtime.Pager[GovernanceAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GovernanceAssignmentsClientListResponse]{
		More: func(page GovernanceAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GovernanceAssignmentsClientListResponse) (GovernanceAssignmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GovernanceAssignmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, assessmentName, options)
			}, nil)
			if err != nil {
				return GovernanceAssignmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GovernanceAssignmentsClient) listCreateRequest(ctx context.Context, scope string, assessmentName string, options *GovernanceAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/governanceAssignments"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GovernanceAssignmentsClient) listHandleResponse(resp *http.Response) (GovernanceAssignmentsClientListResponse, error) {
	result := GovernanceAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GovernanceAssignmentsList); err != nil {
		return GovernanceAssignmentsClientListResponse{}, err
	}
	return result, nil
}
