//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

// ApplicationsClientListResponse contains the response from method ApplicationsClient.NewListPager.
type ApplicationsClientListResponse struct {
	// saas app response with continuation.
	AppResponseWithContinuation
}

// ClientCreateResourceResponse contains the response from method Client.BeginCreateResource.
type ClientCreateResourceResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// ClientDeleteResponse contains the response from method Client.BeginDelete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetResourceResponse contains the response from method Client.GetResource.
type ClientGetResourceResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// ClientUpdateResourceResponse contains the response from method Client.BeginUpdateResource.
type ClientUpdateResourceResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// OperationClientGetResponse contains the response from method OperationClient.BeginGet.
type OperationClientGetResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// saas app operation response with continuation.
	AppOperationsResponseWithContinuation
}

// ResourcesClientListAccessTokenResponse contains the response from method ResourcesClient.ListAccessToken.
type ResourcesClientListAccessTokenResponse struct {
	// the ISV access token result response.
	AccessTokenResult
}

// ResourcesClientListResponse contains the response from method ResourcesClient.NewListPager.
type ResourcesClientListResponse struct {
	// saas resources response with continuation.
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientCreateOrUpdateResponse contains the response from method SubscriptionLevelClient.BeginCreateOrUpdate.
type SubscriptionLevelClientCreateOrUpdateResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// SubscriptionLevelClientDeleteResponse contains the response from method SubscriptionLevelClient.BeginDelete.
type SubscriptionLevelClientDeleteResponse struct {
	// placeholder for future response values
}

// SubscriptionLevelClientGetResponse contains the response from method SubscriptionLevelClient.Get.
type SubscriptionLevelClientGetResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// SubscriptionLevelClientListAccessTokenResponse contains the response from method SubscriptionLevelClient.ListAccessToken.
type SubscriptionLevelClientListAccessTokenResponse struct {
	// the ISV access token result response.
	AccessTokenResult
}

// SubscriptionLevelClientListByAzureSubscriptionResponse contains the response from method SubscriptionLevelClient.NewListByAzureSubscriptionPager.
type SubscriptionLevelClientListByAzureSubscriptionResponse struct {
	// saas resources response with continuation.
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientListByResourceGroupResponse contains the response from method SubscriptionLevelClient.NewListByResourceGroupPager.
type SubscriptionLevelClientListByResourceGroupResponse struct {
	// saas resources response with continuation.
	ResourceResponseWithContinuation
}

// SubscriptionLevelClientMoveResourcesResponse contains the response from method SubscriptionLevelClient.BeginMoveResources.
type SubscriptionLevelClientMoveResourcesResponse struct {
	// placeholder for future response values
}

// SubscriptionLevelClientUpdateResponse contains the response from method SubscriptionLevelClient.BeginUpdate.
type SubscriptionLevelClientUpdateResponse struct {
	// SaaS REST API resource definition.
	Resource
}

// SubscriptionLevelClientUpdateToUnsubscribedResponse contains the response from method SubscriptionLevelClient.BeginUpdateToUnsubscribed.
type SubscriptionLevelClientUpdateToUnsubscribedResponse struct {
	// placeholder for future response values
}

// SubscriptionLevelClientValidateMoveResourcesResponse contains the response from method SubscriptionLevelClient.ValidateMoveResources.
type SubscriptionLevelClientValidateMoveResourcesResponse struct {
	// placeholder for future response values
}
