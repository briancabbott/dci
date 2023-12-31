//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

// CreationSupportedClientGetOptions contains the optional parameters for the CreationSupportedClient.Get method.
type CreationSupportedClientGetOptions struct {
	// placeholder for future optional parameters
}

// CreationSupportedClientListOptions contains the optional parameters for the CreationSupportedClient.NewListPager method.
type CreationSupportedClientListOptions struct {
	// placeholder for future optional parameters
}

// MarketplaceAgreementsClientCreateOrUpdateOptions contains the optional parameters for the MarketplaceAgreementsClient.CreateOrUpdate
// method.
type MarketplaceAgreementsClientCreateOrUpdateOptions struct {
	Body *AgreementResource
}

// MarketplaceAgreementsClientListOptions contains the optional parameters for the MarketplaceAgreementsClient.NewListPager
// method.
type MarketplaceAgreementsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitoredSubscriptionsClientBeginCreateorUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginCreateorUpdate
// method.
type MonitoredSubscriptionsClientBeginCreateorUpdateOptions struct {
	Body *MonitoredSubscriptionProperties

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientBeginDeleteOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginDelete
// method.
type MonitoredSubscriptionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientBeginUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginUpdate
// method.
type MonitoredSubscriptionsClientBeginUpdateOptions struct {
	Body *MonitoredSubscriptionProperties

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitoredSubscriptionsClientGetOptions contains the optional parameters for the MonitoredSubscriptionsClient.Get method.
type MonitoredSubscriptionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitoredSubscriptionsClientListOptions contains the optional parameters for the MonitoredSubscriptionsClient.NewListPager
// method.
type MonitoredSubscriptionsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientBeginCreateOptions contains the optional parameters for the MonitorsClient.BeginCreate method.
type MonitorsClientBeginCreateOptions struct {
	Body *MonitorResource

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginUpdateOptions contains the optional parameters for the MonitorsClient.BeginUpdate method.
type MonitorsClientBeginUpdateOptions struct {
	Body *MonitorResourceUpdateParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientGetDefaultKeyOptions contains the optional parameters for the MonitorsClient.GetDefaultKey method.
type MonitorsClientGetDefaultKeyOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListAPIKeysOptions contains the optional parameters for the MonitorsClient.NewListAPIKeysPager method.
type MonitorsClientListAPIKeysOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.NewListByResourceGroupPager
// method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListHostsOptions contains the optional parameters for the MonitorsClient.NewListHostsPager method.
type MonitorsClientListHostsOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListLinkedResourcesOptions contains the optional parameters for the MonitorsClient.NewListLinkedResourcesPager
// method.
type MonitorsClientListLinkedResourcesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListMonitoredResourcesOptions contains the optional parameters for the MonitorsClient.NewListMonitoredResourcesPager
// method.
type MonitorsClientListMonitoredResourcesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListOptions contains the optional parameters for the MonitorsClient.NewListPager method.
type MonitorsClientListOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientRefreshSetPasswordLinkOptions contains the optional parameters for the MonitorsClient.RefreshSetPasswordLink
// method.
type MonitorsClientRefreshSetPasswordLinkOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientSetDefaultKeyOptions contains the optional parameters for the MonitorsClient.SetDefaultKey method.
type MonitorsClientSetDefaultKeyOptions struct {
	Body *APIKey
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOnConfigurationsClient.BeginCreateOrUpdate
// method.
type SingleSignOnConfigurationsClientBeginCreateOrUpdateOptions struct {
	Body *SingleSignOnResource

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SingleSignOnConfigurationsClientGetOptions contains the optional parameters for the SingleSignOnConfigurationsClient.Get
// method.
type SingleSignOnConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SingleSignOnConfigurationsClientListOptions contains the optional parameters for the SingleSignOnConfigurationsClient.NewListPager
// method.
type SingleSignOnConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.CreateOrUpdate method.
type TagRulesClientCreateOrUpdateOptions struct {
	Body *MonitoringTagRules
}

// TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
type TagRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientListOptions contains the optional parameters for the TagRulesClient.NewListPager method.
type TagRulesClientListOptions struct {
	// placeholder for future optional parameters
}
