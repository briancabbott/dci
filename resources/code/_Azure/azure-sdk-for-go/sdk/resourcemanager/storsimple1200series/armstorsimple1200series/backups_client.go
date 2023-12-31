//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series

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

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
//   - subscriptionID - The subscription id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginClone - Clones the given backup element to a new disk or share with given details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - backupName - The backup name.
//   - elementName - The backup element name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - cloneRequest - The clone request.
//   - options - BackupsClientBeginCloneOptions contains the optional parameters for the BackupsClient.BeginClone method.
func (client *BackupsClient) BeginClone(ctx context.Context, deviceName string, backupName string, elementName string, resourceGroupName string, managerName string, cloneRequest CloneRequest, options *BackupsClientBeginCloneOptions) (*runtime.Poller[BackupsClientCloneResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.clone(ctx, deviceName, backupName, elementName, resourceGroupName, managerName, cloneRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientCloneResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientCloneResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Clone - Clones the given backup element to a new disk or share with given details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *BackupsClient) clone(ctx context.Context, deviceName string, backupName string, elementName string, resourceGroupName string, managerName string, cloneRequest CloneRequest, options *BackupsClientBeginCloneOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginClone"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cloneCreateRequest(ctx, deviceName, backupName, elementName, resourceGroupName, managerName, cloneRequest, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// cloneCreateRequest creates the Clone request.
func (client *BackupsClient) cloneCreateRequest(ctx context.Context, deviceName string, backupName string, elementName string, resourceGroupName string, managerName string, cloneRequest CloneRequest, options *BackupsClientBeginCloneOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}/elements/{elementName}/clone"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	if elementName == "" {
		return nil, errors.New("parameter elementName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{elementName}", url.PathEscape(elementName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, cloneRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the backup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - backupName - The backup name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
func (client *BackupsClient) BeginDelete(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginDeleteOptions) (*runtime.Poller[BackupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, backupName, resourceGroupName, managerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[BackupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[BackupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the backup.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-10-01
func (client *BackupsClient) deleteOperation(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "BackupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, backupName, resourceGroupName, managerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupsClient) deleteCreateRequest(ctx context.Context, deviceName string, backupName string, resourceGroupName string, managerName string, options *BackupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups/{backupName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByDevicePager - Retrieves all the backups in a device. Can be used to get the backups for failover also.
//
// Generated from API version 2016-10-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientListByDeviceOptions contains the optional parameters for the BackupsClient.NewListByDevicePager
//     method.
func (client *BackupsClient) NewListByDevicePager(deviceName string, resourceGroupName string, managerName string, options *BackupsClientListByDeviceOptions) *runtime.Pager[BackupsClientListByDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupsClientListByDeviceResponse]{
		More: func(page BackupsClientListByDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupsClientListByDeviceResponse) (BackupsClientListByDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupsClient.NewListByDevicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDeviceCreateRequest(ctx, deviceName, resourceGroupName, managerName, options)
			}, nil)
			if err != nil {
				return BackupsClientListByDeviceResponse{}, err
			}
			return client.listByDeviceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDeviceCreateRequest creates the ListByDevice request.
func (client *BackupsClient) listByDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, managerName string, options *BackupsClientListByDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/devices/{deviceName}/backups"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.ForFailover != nil {
		reqQP.Set("forFailover", strconv.FormatBool(*options.ForFailover))
	}
	reqQP.Set("api-version", "2016-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDeviceHandleResponse handles the ListByDevice response.
func (client *BackupsClient) listByDeviceHandleResponse(resp *http.Response) (BackupsClientListByDeviceResponse, error) {
	result := BackupsClientListByDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupList); err != nil {
		return BackupsClientListByDeviceResponse{}, err
	}
	return result, nil
}

// NewListByManagerPager - Retrieves all the backups in a manager.
//
// Generated from API version 2016-10-01
//   - resourceGroupName - The resource group name
//   - managerName - The manager name
//   - options - BackupsClientListByManagerOptions contains the optional parameters for the BackupsClient.NewListByManagerPager
//     method.
func (client *BackupsClient) NewListByManagerPager(resourceGroupName string, managerName string, options *BackupsClientListByManagerOptions) *runtime.Pager[BackupsClientListByManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupsClientListByManagerResponse]{
		More: func(page BackupsClientListByManagerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupsClientListByManagerResponse) (BackupsClientListByManagerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupsClient.NewListByManagerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagerCreateRequest(ctx, resourceGroupName, managerName, options)
			}, nil)
			if err != nil {
				return BackupsClientListByManagerResponse{}, err
			}
			return client.listByManagerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagerCreateRequest creates the ListByManager request.
func (client *BackupsClient) listByManagerCreateRequest(ctx context.Context, resourceGroupName string, managerName string, options *BackupsClientListByManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorSimple/managers/{managerName}/backups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managerName == "" {
		return nil, errors.New("parameter managerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managerName}", url.PathEscape(managerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagerHandleResponse handles the ListByManager response.
func (client *BackupsClient) listByManagerHandleResponse(resp *http.Response) (BackupsClientListByManagerResponse, error) {
	result := BackupsClientListByManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupList); err != nil {
		return BackupsClientListByManagerResponse{}, err
	}
	return result, nil
}
