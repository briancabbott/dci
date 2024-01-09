//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

// APIKeysClientCreateOptions contains the optional parameters for the APIKeysClient.Create method.
type APIKeysClientCreateOptions struct {
	// placeholder for future optional parameters
}

// APIKeysClientDeleteOptions contains the optional parameters for the APIKeysClient.Delete method.
type APIKeysClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// APIKeysClientGetOptions contains the optional parameters for the APIKeysClient.Get method.
type APIKeysClientGetOptions struct {
	// placeholder for future optional parameters
}

// APIKeysClientListOptions contains the optional parameters for the APIKeysClient.NewListPager method.
type APIKeysClientListOptions struct {
	// placeholder for future optional parameters
}

// AnalyticsItemsClientDeleteOptions contains the optional parameters for the AnalyticsItemsClient.Delete method.
type AnalyticsItemsClientDeleteOptions struct {
	// The Id of a specific item defined in the Application Insights component
	ID *string

	// The name of a specific item defined in the Application Insights component
	Name *string
}

// AnalyticsItemsClientGetOptions contains the optional parameters for the AnalyticsItemsClient.Get method.
type AnalyticsItemsClientGetOptions struct {
	// The Id of a specific item defined in the Application Insights component
	ID *string

	// The name of a specific item defined in the Application Insights component
	Name *string
}

// AnalyticsItemsClientListOptions contains the optional parameters for the AnalyticsItemsClient.List method.
type AnalyticsItemsClientListOptions struct {
	// Flag indicating whether or not to return the content of each applicable item. If false, only return the item information.
	IncludeContent *bool

	// Enum indicating if this item definition is owned by a specific user or is shared between all users with access to the Application
	// Insights component.
	Scope *ItemScope

	// Enum indicating the type of the Analytics item.
	Type *ItemTypeParameter
}

// AnalyticsItemsClientPutOptions contains the optional parameters for the AnalyticsItemsClient.Put method.
type AnalyticsItemsClientPutOptions struct {
	// Flag indicating whether or not to force save an item. This allows overriding an item if it already exists.
	OverrideItem *bool
}

// AnnotationsClientCreateOptions contains the optional parameters for the AnnotationsClient.Create method.
type AnnotationsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// AnnotationsClientDeleteOptions contains the optional parameters for the AnnotationsClient.Delete method.
type AnnotationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AnnotationsClientGetOptions contains the optional parameters for the AnnotationsClient.Get method.
type AnnotationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AnnotationsClientListOptions contains the optional parameters for the AnnotationsClient.NewListPager method.
type AnnotationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ComponentAvailableFeaturesClientGetOptions contains the optional parameters for the ComponentAvailableFeaturesClient.Get
// method.
type ComponentAvailableFeaturesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentCurrentBillingFeaturesClientGetOptions contains the optional parameters for the ComponentCurrentBillingFeaturesClient.Get
// method.
type ComponentCurrentBillingFeaturesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentCurrentBillingFeaturesClientUpdateOptions contains the optional parameters for the ComponentCurrentBillingFeaturesClient.Update
// method.
type ComponentCurrentBillingFeaturesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ComponentFeatureCapabilitiesClientGetOptions contains the optional parameters for the ComponentFeatureCapabilitiesClient.Get
// method.
type ComponentFeatureCapabilitiesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentQuotaStatusClientGetOptions contains the optional parameters for the ComponentQuotaStatusClient.Get method.
type ComponentQuotaStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientCreateOrUpdateOptions contains the optional parameters for the ComponentsClient.CreateOrUpdate method.
type ComponentsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientDeleteOptions contains the optional parameters for the ComponentsClient.Delete method.
type ComponentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientGetOptions contains the optional parameters for the ComponentsClient.Get method.
type ComponentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientGetPurgeStatusOptions contains the optional parameters for the ComponentsClient.GetPurgeStatus method.
type ComponentsClientGetPurgeStatusOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientListByResourceGroupOptions contains the optional parameters for the ComponentsClient.NewListByResourceGroupPager
// method.
type ComponentsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientListOptions contains the optional parameters for the ComponentsClient.NewListPager method.
type ComponentsClientListOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientPurgeOptions contains the optional parameters for the ComponentsClient.Purge method.
type ComponentsClientPurgeOptions struct {
	// placeholder for future optional parameters
}

// ComponentsClientUpdateTagsOptions contains the optional parameters for the ComponentsClient.UpdateTags method.
type ComponentsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// ExportConfigurationsClientCreateOptions contains the optional parameters for the ExportConfigurationsClient.Create method.
type ExportConfigurationsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// ExportConfigurationsClientDeleteOptions contains the optional parameters for the ExportConfigurationsClient.Delete method.
type ExportConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ExportConfigurationsClientGetOptions contains the optional parameters for the ExportConfigurationsClient.Get method.
type ExportConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ExportConfigurationsClientListOptions contains the optional parameters for the ExportConfigurationsClient.List method.
type ExportConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ExportConfigurationsClientUpdateOptions contains the optional parameters for the ExportConfigurationsClient.Update method.
type ExportConfigurationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// FavoritesClientAddOptions contains the optional parameters for the FavoritesClient.Add method.
type FavoritesClientAddOptions struct {
	// placeholder for future optional parameters
}

// FavoritesClientDeleteOptions contains the optional parameters for the FavoritesClient.Delete method.
type FavoritesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FavoritesClientGetOptions contains the optional parameters for the FavoritesClient.Get method.
type FavoritesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FavoritesClientListOptions contains the optional parameters for the FavoritesClient.List method.
type FavoritesClientListOptions struct {
	// Flag indicating whether or not to return the full content for each applicable favorite. If false, only return summary content
	// for favorites.
	CanFetchContent *bool

	// The type of favorite. Value can be either shared or user.
	FavoriteType *FavoriteType

	// Source type of favorite to return. When left out, the source type defaults to 'other' (not present in this enum).
	SourceType *FavoriteSourceType

	// Tags that must be present on each favorite returned.
	Tags []string
}

// FavoritesClientUpdateOptions contains the optional parameters for the FavoritesClient.Update method.
type FavoritesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// MyWorkbooksClientCreateOrUpdateOptions contains the optional parameters for the MyWorkbooksClient.CreateOrUpdate method.
type MyWorkbooksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// MyWorkbooksClientDeleteOptions contains the optional parameters for the MyWorkbooksClient.Delete method.
type MyWorkbooksClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MyWorkbooksClientGetOptions contains the optional parameters for the MyWorkbooksClient.Get method.
type MyWorkbooksClientGetOptions struct {
	// placeholder for future optional parameters
}

// MyWorkbooksClientListByResourceGroupOptions contains the optional parameters for the MyWorkbooksClient.NewListByResourceGroupPager
// method.
type MyWorkbooksClientListByResourceGroupOptions struct {
	// Flag indicating whether or not to return the full content for each applicable workbook. If false, only return summary content
	// for workbooks.
	CanFetchContent *bool

	// Tags presents on each workbook returned.
	Tags []string
}

// MyWorkbooksClientListBySubscriptionOptions contains the optional parameters for the MyWorkbooksClient.NewListBySubscriptionPager
// method.
type MyWorkbooksClientListBySubscriptionOptions struct {
	// Flag indicating whether or not to return the full content for each applicable workbook. If false, only return summary content
	// for workbooks.
	CanFetchContent *bool

	// Tags presents on each workbook returned.
	Tags []string
}

// MyWorkbooksClientUpdateOptions contains the optional parameters for the MyWorkbooksClient.Update method.
type MyWorkbooksClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ProactiveDetectionConfigurationsClientGetOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Get
// method.
type ProactiveDetectionConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ProactiveDetectionConfigurationsClientListOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.List
// method.
type ProactiveDetectionConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ProactiveDetectionConfigurationsClientUpdateOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Update
// method.
type ProactiveDetectionConfigurationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// WebTestLocationsClientListOptions contains the optional parameters for the WebTestLocationsClient.NewListPager method.
type WebTestLocationsClientListOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientCreateOrUpdateOptions contains the optional parameters for the WebTestsClient.CreateOrUpdate method.
type WebTestsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientDeleteOptions contains the optional parameters for the WebTestsClient.Delete method.
type WebTestsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientGetOptions contains the optional parameters for the WebTestsClient.Get method.
type WebTestsClientGetOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientListByComponentOptions contains the optional parameters for the WebTestsClient.NewListByComponentPager method.
type WebTestsClientListByComponentOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientListByResourceGroupOptions contains the optional parameters for the WebTestsClient.NewListByResourceGroupPager
// method.
type WebTestsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientListOptions contains the optional parameters for the WebTestsClient.NewListPager method.
type WebTestsClientListOptions struct {
	// placeholder for future optional parameters
}

// WebTestsClientUpdateTagsOptions contains the optional parameters for the WebTestsClient.UpdateTags method.
type WebTestsClientUpdateTagsOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientCreateOptions contains the optional parameters for the WorkItemConfigurationsClient.Create
// method.
type WorkItemConfigurationsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientDeleteOptions contains the optional parameters for the WorkItemConfigurationsClient.Delete
// method.
type WorkItemConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientGetDefaultOptions contains the optional parameters for the WorkItemConfigurationsClient.GetDefault
// method.
type WorkItemConfigurationsClientGetDefaultOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientGetItemOptions contains the optional parameters for the WorkItemConfigurationsClient.GetItem
// method.
type WorkItemConfigurationsClientGetItemOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientListOptions contains the optional parameters for the WorkItemConfigurationsClient.NewListPager
// method.
type WorkItemConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// WorkItemConfigurationsClientUpdateItemOptions contains the optional parameters for the WorkItemConfigurationsClient.UpdateItem
// method.
type WorkItemConfigurationsClientUpdateItemOptions struct {
	// placeholder for future optional parameters
}

// WorkbooksClientCreateOrUpdateOptions contains the optional parameters for the WorkbooksClient.CreateOrUpdate method.
type WorkbooksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// WorkbooksClientDeleteOptions contains the optional parameters for the WorkbooksClient.Delete method.
type WorkbooksClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkbooksClientGetOptions contains the optional parameters for the WorkbooksClient.Get method.
type WorkbooksClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkbooksClientListByResourceGroupOptions contains the optional parameters for the WorkbooksClient.NewListByResourceGroupPager
// method.
type WorkbooksClientListByResourceGroupOptions struct {
	// Flag indicating whether or not to return the full content for each applicable workbook. If false, only return summary content
	// for workbooks.
	CanFetchContent *bool

	// Tags presents on each workbook returned.
	Tags []string
}

// WorkbooksClientUpdateOptions contains the optional parameters for the WorkbooksClient.Update method.
type WorkbooksClientUpdateOptions struct {
	// placeholder for future optional parameters
}
