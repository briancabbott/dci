//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourceconnector

// AppliancesClientCreateOrUpdateResponse contains the response from method AppliancesClient.BeginCreateOrUpdate.
type AppliancesClientCreateOrUpdateResponse struct {
	// Appliances definition.
	Appliance
}

// AppliancesClientDeleteResponse contains the response from method AppliancesClient.BeginDelete.
type AppliancesClientDeleteResponse struct {
	// placeholder for future response values
}

// AppliancesClientGetResponse contains the response from method AppliancesClient.Get.
type AppliancesClientGetResponse struct {
	// Appliances definition.
	Appliance
}

// AppliancesClientGetTelemetryConfigResponse contains the response from method AppliancesClient.GetTelemetryConfig.
type AppliancesClientGetTelemetryConfigResponse struct {
	// The Get Telemetry Config Result appliance.
	ApplianceGetTelemetryConfigResult
}

// AppliancesClientGetUpgradeGraphResponse contains the response from method AppliancesClient.GetUpgradeGraph.
type AppliancesClientGetUpgradeGraphResponse struct {
	// The Upgrade Graph for appliance.
	UpgradeGraph
}

// AppliancesClientListByResourceGroupResponse contains the response from method AppliancesClient.NewListByResourceGroupPager.
type AppliancesClientListByResourceGroupResponse struct {
	// The List Appliances operation response.
	ApplianceListResult
}

// AppliancesClientListBySubscriptionResponse contains the response from method AppliancesClient.NewListBySubscriptionPager.
type AppliancesClientListBySubscriptionResponse struct {
	// The List Appliances operation response.
	ApplianceListResult
}

// AppliancesClientListClusterUserCredentialResponse contains the response from method AppliancesClient.ListClusterUserCredential.
type AppliancesClientListClusterUserCredentialResponse struct {
	// The List Cluster User Credential appliance.
	ApplianceListCredentialResults
}

// AppliancesClientListKeysResponse contains the response from method AppliancesClient.ListKeys.
type AppliancesClientListKeysResponse struct {
	// The List Cluster Keys Results appliance.
	ApplianceListKeysResults
}

// AppliancesClientListOperationsResponse contains the response from method AppliancesClient.NewListOperationsPager.
type AppliancesClientListOperationsResponse struct {
	// Lists of Appliances operations.
	ApplianceOperationsList
}

// AppliancesClientUpdateResponse contains the response from method AppliancesClient.Update.
type AppliancesClientUpdateResponse struct {
	// Appliances definition.
	Appliance
}
