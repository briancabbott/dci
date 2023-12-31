//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

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

// ConfigurationProfileHCIAssignmentsClient contains the methods for the ConfigurationProfileHCIAssignments group.
// Don't use this type directly, use NewConfigurationProfileHCIAssignmentsClient() instead.
type ConfigurationProfileHCIAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationProfileHCIAssignmentsClient creates a new instance of ConfigurationProfileHCIAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationProfileHCIAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationProfileHCIAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationProfileHCIAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates an association between a AzureStackHCI cluster and Automanage configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Arc machine.
//   - configurationProfileAssignmentName - Name of the configuration profile assignment. Only default is supported.
//   - parameters - Parameters supplied to the create or update configuration profile assignment.
//   - options - ConfigurationProfileHCIAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationProfileHCIAssignmentsClient.CreateOrUpdate
//     method.
func (client *ConfigurationProfileHCIAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileHCIAssignmentsClientCreateOrUpdateOptions) (ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConfigurationProfileHCIAssignmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, parameters, options)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationProfileHCIAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileHCIAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationProfileHCIAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse, error) {
	result := ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Arc machine.
//   - configurationProfileAssignmentName - Name of the configuration profile assignment
//   - options - ConfigurationProfileHCIAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationProfileHCIAssignmentsClient.Delete
//     method.
func (client *ConfigurationProfileHCIAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *ConfigurationProfileHCIAssignmentsClientDeleteOptions) (ConfigurationProfileHCIAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "ConfigurationProfileHCIAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, options)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileHCIAssignmentsClientDeleteResponse{}, err
	}
	return ConfigurationProfileHCIAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationProfileHCIAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *ConfigurationProfileHCIAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get information about a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Arc machine.
//   - configurationProfileAssignmentName - The configuration profile assignment name.
//   - options - ConfigurationProfileHCIAssignmentsClientGetOptions contains the optional parameters for the ConfigurationProfileHCIAssignmentsClient.Get
//     method.
func (client *ConfigurationProfileHCIAssignmentsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *ConfigurationProfileHCIAssignmentsClientGetOptions) (ConfigurationProfileHCIAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "ConfigurationProfileHCIAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, configurationProfileAssignmentName, options)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileHCIAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileHCIAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationProfileHCIAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *ConfigurationProfileHCIAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationProfileHCIAssignmentsClient) getHandleResponse(resp *http.Response) (ConfigurationProfileHCIAssignmentsClientGetResponse, error) {
	result := ConfigurationProfileHCIAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileHCIAssignmentsClientGetResponse{}, err
	}
	return result, nil
}
