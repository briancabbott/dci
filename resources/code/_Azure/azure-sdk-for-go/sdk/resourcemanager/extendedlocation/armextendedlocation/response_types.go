//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextendedlocation

// CustomLocationsClientCreateOrUpdateResponse contains the response from method CustomLocationsClient.BeginCreateOrUpdate.
type CustomLocationsClientCreateOrUpdateResponse struct {
	// Custom Locations definition.
	CustomLocation
}

// CustomLocationsClientDeleteResponse contains the response from method CustomLocationsClient.BeginDelete.
type CustomLocationsClientDeleteResponse struct {
	// placeholder for future response values
}

// CustomLocationsClientFindTargetResourceGroupResponse contains the response from method CustomLocationsClient.FindTargetResourceGroup.
type CustomLocationsClientFindTargetResourceGroupResponse struct {
	// The Find Target Resource Group operation response.
	CustomLocationFindTargetResourceGroupResult
}

// CustomLocationsClientGetResponse contains the response from method CustomLocationsClient.Get.
type CustomLocationsClientGetResponse struct {
	// Custom Locations definition.
	CustomLocation
}

// CustomLocationsClientListByResourceGroupResponse contains the response from method CustomLocationsClient.NewListByResourceGroupPager.
type CustomLocationsClientListByResourceGroupResponse struct {
	// The List Custom Locations operation response.
	CustomLocationListResult
}

// CustomLocationsClientListBySubscriptionResponse contains the response from method CustomLocationsClient.NewListBySubscriptionPager.
type CustomLocationsClientListBySubscriptionResponse struct {
	// The List Custom Locations operation response.
	CustomLocationListResult
}

// CustomLocationsClientListEnabledResourceTypesResponse contains the response from method CustomLocationsClient.NewListEnabledResourceTypesPager.
type CustomLocationsClientListEnabledResourceTypesResponse struct {
	// List of EnabledResourceTypes definition.
	EnabledResourceTypesListResult
}

// CustomLocationsClientListOperationsResponse contains the response from method CustomLocationsClient.NewListOperationsPager.
type CustomLocationsClientListOperationsResponse struct {
	// Lists of Custom Locations operations.
	CustomLocationOperationsList
}

// CustomLocationsClientUpdateResponse contains the response from method CustomLocationsClient.Update.
type CustomLocationsClientUpdateResponse struct {
	// Custom Locations definition.
	CustomLocation
}

// ResourceSyncRulesClientCreateOrUpdateResponse contains the response from method ResourceSyncRulesClient.BeginCreateOrUpdate.
type ResourceSyncRulesClientCreateOrUpdateResponse struct {
	// Resource Sync Rules definition.
	ResourceSyncRule
}

// ResourceSyncRulesClientDeleteResponse contains the response from method ResourceSyncRulesClient.Delete.
type ResourceSyncRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceSyncRulesClientGetResponse contains the response from method ResourceSyncRulesClient.Get.
type ResourceSyncRulesClientGetResponse struct {
	// Resource Sync Rules definition.
	ResourceSyncRule
}

// ResourceSyncRulesClientListByCustomLocationIDResponse contains the response from method ResourceSyncRulesClient.NewListByCustomLocationIDPager.
type ResourceSyncRulesClientListByCustomLocationIDResponse struct {
	// The List Resource Sync Rules operation response.
	ResourceSyncRuleListResult
}

// ResourceSyncRulesClientUpdateResponse contains the response from method ResourceSyncRulesClient.BeginUpdate.
type ResourceSyncRulesClientUpdateResponse struct {
	// Resource Sync Rules definition.
	ResourceSyncRule
}
