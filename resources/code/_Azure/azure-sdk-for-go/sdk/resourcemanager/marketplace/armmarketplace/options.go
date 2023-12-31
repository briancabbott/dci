//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplace

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientAcknowledgeOfferNotificationOptions contains the optional parameters for the PrivateStoreClient.AcknowledgeOfferNotification
// method.
type PrivateStoreClientAcknowledgeOfferNotificationOptions struct {
	Payload *AcknowledgeOfferNotificationProperties
}

// PrivateStoreClientAdminRequestApprovalsListOptions contains the optional parameters for the PrivateStoreClient.AdminRequestApprovalsList
// method.
type PrivateStoreClientAdminRequestApprovalsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientBillingAccountsOptions contains the optional parameters for the PrivateStoreClient.BillingAccounts method.
type PrivateStoreClientBillingAccountsOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientBulkCollectionsActionOptions contains the optional parameters for the PrivateStoreClient.BulkCollectionsAction
// method.
type PrivateStoreClientBulkCollectionsActionOptions struct {
	Payload *BulkCollectionsPayload
}

// PrivateStoreClientCollectionsToSubscriptionsMappingOptions contains the optional parameters for the PrivateStoreClient.CollectionsToSubscriptionsMapping
// method.
type PrivateStoreClientCollectionsToSubscriptionsMappingOptions struct {
	Payload *CollectionsToSubscriptionsMappingPayload
}

// PrivateStoreClientCreateApprovalRequestOptions contains the optional parameters for the PrivateStoreClient.CreateApprovalRequest
// method.
type PrivateStoreClientCreateApprovalRequestOptions struct {
	Payload *RequestApprovalResource
}

// PrivateStoreClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreClient.CreateOrUpdate method.
type PrivateStoreClientCreateOrUpdateOptions struct {
	Payload *PrivateStore
}

// PrivateStoreClientDeleteOptions contains the optional parameters for the PrivateStoreClient.Delete method.
type PrivateStoreClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientFetchAllSubscriptionsInTenantOptions contains the optional parameters for the PrivateStoreClient.FetchAllSubscriptionsInTenant
// method.
type PrivateStoreClientFetchAllSubscriptionsInTenantOptions struct {
	// The skip token to get the next page.
	NextPageToken *string
}

// PrivateStoreClientGetAdminRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.GetAdminRequestApproval
// method.
type PrivateStoreClientGetAdminRequestApprovalOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetApprovalRequestsListOptions contains the optional parameters for the PrivateStoreClient.GetApprovalRequestsList
// method.
type PrivateStoreClientGetApprovalRequestsListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetOptions contains the optional parameters for the PrivateStoreClient.Get method.
type PrivateStoreClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientGetRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.GetRequestApproval
// method.
type PrivateStoreClientGetRequestApprovalOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientListNewPlansNotificationsOptions contains the optional parameters for the PrivateStoreClient.ListNewPlansNotifications
// method.
type PrivateStoreClientListNewPlansNotificationsOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientListOptions contains the optional parameters for the PrivateStoreClient.NewListPager method.
type PrivateStoreClientListOptions struct {
	// Determines if to use cache or DB for serving this request
	UseCache *string
}

// PrivateStoreClientListStopSellOffersPlansNotificationsOptions contains the optional parameters for the PrivateStoreClient.ListStopSellOffersPlansNotifications
// method.
type PrivateStoreClientListStopSellOffersPlansNotificationsOptions struct {
	StopSellSubscriptions *StopSellSubscriptions
}

// PrivateStoreClientListSubscriptionsContextOptions contains the optional parameters for the PrivateStoreClient.ListSubscriptionsContext
// method.
type PrivateStoreClientListSubscriptionsContextOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientQueryApprovedPlansOptions contains the optional parameters for the PrivateStoreClient.QueryApprovedPlans
// method.
type PrivateStoreClientQueryApprovedPlansOptions struct {
	Payload *QueryApprovedPlansPayload
}

// PrivateStoreClientQueryNotificationsStateOptions contains the optional parameters for the PrivateStoreClient.QueryNotificationsState
// method.
type PrivateStoreClientQueryNotificationsStateOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientQueryOffersOptions contains the optional parameters for the PrivateStoreClient.QueryOffers method.
type PrivateStoreClientQueryOffersOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreClientQueryRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.QueryRequestApproval
// method.
type PrivateStoreClientQueryRequestApprovalOptions struct {
	Payload *QueryRequestApprovalProperties
}

// PrivateStoreClientUpdateAdminRequestApprovalOptions contains the optional parameters for the PrivateStoreClient.UpdateAdminRequestApproval
// method.
type PrivateStoreClientUpdateAdminRequestApprovalOptions struct {
	Payload *AdminRequestApprovalsResource
}

// PrivateStoreClientWithdrawPlanOptions contains the optional parameters for the PrivateStoreClient.WithdrawPlan method.
type PrivateStoreClientWithdrawPlanOptions struct {
	Payload *WithdrawProperties
}

// PrivateStoreCollectionClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreCollectionClient.CreateOrUpdate
// method.
type PrivateStoreCollectionClientCreateOrUpdateOptions struct {
	Payload *Collection
}

// PrivateStoreCollectionClientDeleteOptions contains the optional parameters for the PrivateStoreCollectionClient.Delete
// method.
type PrivateStoreCollectionClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientGetOptions contains the optional parameters for the PrivateStoreCollectionClient.Get method.
type PrivateStoreCollectionClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientListOptions contains the optional parameters for the PrivateStoreCollectionClient.List method.
type PrivateStoreCollectionClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionClientPostOptions contains the optional parameters for the PrivateStoreCollectionClient.Post method.
type PrivateStoreCollectionClientPostOptions struct {
	Payload *Operation
}

// PrivateStoreCollectionClientTransferOffersOptions contains the optional parameters for the PrivateStoreCollectionClient.TransferOffers
// method.
type PrivateStoreCollectionClientTransferOffersOptions struct {
	Payload *TransferOffersProperties
}

// PrivateStoreCollectionOfferClientCreateOrUpdateOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.CreateOrUpdate
// method.
type PrivateStoreCollectionOfferClientCreateOrUpdateOptions struct {
	Payload *Offer
}

// PrivateStoreCollectionOfferClientDeleteOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Delete
// method.
type PrivateStoreCollectionOfferClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientGetOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Get
// method.
type PrivateStoreCollectionOfferClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientListOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.NewListPager
// method.
type PrivateStoreCollectionOfferClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateStoreCollectionOfferClientPostOptions contains the optional parameters for the PrivateStoreCollectionOfferClient.Post
// method.
type PrivateStoreCollectionOfferClientPostOptions struct {
	Payload *Operation
}
