//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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

// NotificationRegistrationsClient contains the methods for the NotificationRegistrations group.
// Don't use this type directly, use NewNotificationRegistrationsClient() instead.
type NotificationRegistrationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNotificationRegistrationsClient creates a new instance of NotificationRegistrationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNotificationRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NotificationRegistrationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NotificationRegistrationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a notification registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - notificationRegistrationName - The notification registration.
//   - properties - The required body parameters supplied to the notification registration operation.
//   - options - NotificationRegistrationsClientCreateOrUpdateOptions contains the optional parameters for the NotificationRegistrationsClient.CreateOrUpdate
//     method.
func (client *NotificationRegistrationsClient) CreateOrUpdate(ctx context.Context, providerNamespace string, notificationRegistrationName string, properties NotificationRegistration, options *NotificationRegistrationsClientCreateOrUpdateOptions) (NotificationRegistrationsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "NotificationRegistrationsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, notificationRegistrationName, properties, options)
	if err != nil {
		return NotificationRegistrationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRegistrationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRegistrationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationRegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, notificationRegistrationName string, properties NotificationRegistration, options *NotificationRegistrationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/notificationRegistrations/{notificationRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if notificationRegistrationName == "" {
		return nil, errors.New("parameter notificationRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationRegistrationName}", url.PathEscape(notificationRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationRegistrationsClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationRegistrationsClientCreateOrUpdateResponse, error) {
	result := NotificationRegistrationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationRegistration); err != nil {
		return NotificationRegistrationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a notification registration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - notificationRegistrationName - The notification registration.
//   - options - NotificationRegistrationsClientDeleteOptions contains the optional parameters for the NotificationRegistrationsClient.Delete
//     method.
func (client *NotificationRegistrationsClient) Delete(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *NotificationRegistrationsClientDeleteOptions) (NotificationRegistrationsClientDeleteResponse, error) {
	var err error
	const operationName = "NotificationRegistrationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, providerNamespace, notificationRegistrationName, options)
	if err != nil {
		return NotificationRegistrationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRegistrationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRegistrationsClientDeleteResponse{}, err
	}
	return NotificationRegistrationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationRegistrationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *NotificationRegistrationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/notificationRegistrations/{notificationRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if notificationRegistrationName == "" {
		return nil, errors.New("parameter notificationRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationRegistrationName}", url.PathEscape(notificationRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the notification registration details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - notificationRegistrationName - The notification registration.
//   - options - NotificationRegistrationsClientGetOptions contains the optional parameters for the NotificationRegistrationsClient.Get
//     method.
func (client *NotificationRegistrationsClient) Get(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *NotificationRegistrationsClientGetOptions) (NotificationRegistrationsClientGetResponse, error) {
	var err error
	const operationName = "NotificationRegistrationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, providerNamespace, notificationRegistrationName, options)
	if err != nil {
		return NotificationRegistrationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NotificationRegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NotificationRegistrationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NotificationRegistrationsClient) getCreateRequest(ctx context.Context, providerNamespace string, notificationRegistrationName string, options *NotificationRegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/notificationRegistrations/{notificationRegistrationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if notificationRegistrationName == "" {
		return nil, errors.New("parameter notificationRegistrationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notificationRegistrationName}", url.PathEscape(notificationRegistrationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NotificationRegistrationsClient) getHandleResponse(resp *http.Response) (NotificationRegistrationsClientGetResponse, error) {
	result := NotificationRegistrationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationRegistration); err != nil {
		return NotificationRegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProviderRegistrationPager - Gets the list of the notification registrations for the given provider.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - options - NotificationRegistrationsClientListByProviderRegistrationOptions contains the optional parameters for the NotificationRegistrationsClient.NewListByProviderRegistrationPager
//     method.
func (client *NotificationRegistrationsClient) NewListByProviderRegistrationPager(providerNamespace string, options *NotificationRegistrationsClientListByProviderRegistrationOptions) *runtime.Pager[NotificationRegistrationsClientListByProviderRegistrationResponse] {
	return runtime.NewPager(runtime.PagingHandler[NotificationRegistrationsClientListByProviderRegistrationResponse]{
		More: func(page NotificationRegistrationsClientListByProviderRegistrationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NotificationRegistrationsClientListByProviderRegistrationResponse) (NotificationRegistrationsClientListByProviderRegistrationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NotificationRegistrationsClient.NewListByProviderRegistrationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
			}, nil)
			if err != nil {
				return NotificationRegistrationsClientListByProviderRegistrationResponse{}, err
			}
			return client.listByProviderRegistrationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *NotificationRegistrationsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, options *NotificationRegistrationsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/notificationRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *NotificationRegistrationsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (NotificationRegistrationsClientListByProviderRegistrationResponse, error) {
	result := NotificationRegistrationsClientListByProviderRegistrationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationRegistrationArrayResponseWithContinuation); err != nil {
		return NotificationRegistrationsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}
