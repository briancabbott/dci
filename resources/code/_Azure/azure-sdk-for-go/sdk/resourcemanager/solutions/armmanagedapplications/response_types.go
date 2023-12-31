//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

// ApplicationClientListOperationsResponse contains the response from method ApplicationClient.NewListOperationsPager.
type ApplicationClientListOperationsResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ApplicationDefinitionsClientCreateOrUpdateByIDResponse contains the response from method ApplicationDefinitionsClient.CreateOrUpdateByID.
type ApplicationDefinitionsClientCreateOrUpdateByIDResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationDefinitionsClientCreateOrUpdateResponse contains the response from method ApplicationDefinitionsClient.CreateOrUpdate.
type ApplicationDefinitionsClientCreateOrUpdateResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationDefinitionsClientDeleteByIDResponse contains the response from method ApplicationDefinitionsClient.DeleteByID.
type ApplicationDefinitionsClientDeleteByIDResponse struct {
	// placeholder for future response values
}

// ApplicationDefinitionsClientDeleteResponse contains the response from method ApplicationDefinitionsClient.Delete.
type ApplicationDefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationDefinitionsClientGetByIDResponse contains the response from method ApplicationDefinitionsClient.GetByID.
type ApplicationDefinitionsClientGetByIDResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationDefinitionsClientGetResponse contains the response from method ApplicationDefinitionsClient.Get.
type ApplicationDefinitionsClientGetResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationDefinitionsClientListByResourceGroupResponse contains the response from method ApplicationDefinitionsClient.NewListByResourceGroupPager.
type ApplicationDefinitionsClientListByResourceGroupResponse struct {
	// List of managed application definitions.
	ApplicationDefinitionListResult
}

// ApplicationDefinitionsClientListBySubscriptionResponse contains the response from method ApplicationDefinitionsClient.NewListBySubscriptionPager.
type ApplicationDefinitionsClientListBySubscriptionResponse struct {
	// List of managed application definitions.
	ApplicationDefinitionListResult
}

// ApplicationDefinitionsClientUpdateByIDResponse contains the response from method ApplicationDefinitionsClient.UpdateByID.
type ApplicationDefinitionsClientUpdateByIDResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationDefinitionsClientUpdateResponse contains the response from method ApplicationDefinitionsClient.Update.
type ApplicationDefinitionsClientUpdateResponse struct {
	// Information about managed application definition.
	ApplicationDefinition
}

// ApplicationsClientCreateOrUpdateByIDResponse contains the response from method ApplicationsClient.BeginCreateOrUpdateByID.
type ApplicationsClientCreateOrUpdateByIDResponse struct {
	// Information about managed application.
	Application
}

// ApplicationsClientCreateOrUpdateResponse contains the response from method ApplicationsClient.BeginCreateOrUpdate.
type ApplicationsClientCreateOrUpdateResponse struct {
	// Information about managed application.
	Application
}

// ApplicationsClientDeleteByIDResponse contains the response from method ApplicationsClient.BeginDeleteByID.
type ApplicationsClientDeleteByIDResponse struct {
	// placeholder for future response values
}

// ApplicationsClientDeleteResponse contains the response from method ApplicationsClient.BeginDelete.
type ApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationsClientGetByIDResponse contains the response from method ApplicationsClient.GetByID.
type ApplicationsClientGetByIDResponse struct {
	// Information about managed application.
	Application
}

// ApplicationsClientGetResponse contains the response from method ApplicationsClient.Get.
type ApplicationsClientGetResponse struct {
	// Information about managed application.
	Application
}

// ApplicationsClientListAllowedUpgradePlansResponse contains the response from method ApplicationsClient.ListAllowedUpgradePlans.
type ApplicationsClientListAllowedUpgradePlansResponse struct {
	// The array of plan.
	AllowedUpgradePlansResult
}

// ApplicationsClientListByResourceGroupResponse contains the response from method ApplicationsClient.NewListByResourceGroupPager.
type ApplicationsClientListByResourceGroupResponse struct {
	// List of managed applications.
	ApplicationListResult
}

// ApplicationsClientListBySubscriptionResponse contains the response from method ApplicationsClient.NewListBySubscriptionPager.
type ApplicationsClientListBySubscriptionResponse struct {
	// List of managed applications.
	ApplicationListResult
}

// ApplicationsClientListTokensResponse contains the response from method ApplicationsClient.ListTokens.
type ApplicationsClientListTokensResponse struct {
	// The array of managed identity tokens.
	ManagedIdentityTokenResult
}

// ApplicationsClientRefreshPermissionsResponse contains the response from method ApplicationsClient.BeginRefreshPermissions.
type ApplicationsClientRefreshPermissionsResponse struct {
	// placeholder for future response values
}

// ApplicationsClientUpdateAccessResponse contains the response from method ApplicationsClient.BeginUpdateAccess.
type ApplicationsClientUpdateAccessResponse struct {
	// Update access request definition.
	UpdateAccessDefinition
}

// ApplicationsClientUpdateByIDResponse contains the response from method ApplicationsClient.BeginUpdateByID.
type ApplicationsClientUpdateByIDResponse struct {
	// Information about managed application.
	ApplicationPatchable
}

// ApplicationsClientUpdateResponse contains the response from method ApplicationsClient.BeginUpdate.
type ApplicationsClientUpdateResponse struct {
	// Information about managed application.
	ApplicationPatchable
}

// JitRequestsClientCreateOrUpdateResponse contains the response from method JitRequestsClient.BeginCreateOrUpdate.
type JitRequestsClientCreateOrUpdateResponse struct {
	// Information about JIT request definition.
	JitRequestDefinition
}

// JitRequestsClientDeleteResponse contains the response from method JitRequestsClient.Delete.
type JitRequestsClientDeleteResponse struct {
	// placeholder for future response values
}

// JitRequestsClientGetResponse contains the response from method JitRequestsClient.Get.
type JitRequestsClientGetResponse struct {
	// Information about JIT request definition.
	JitRequestDefinition
}

// JitRequestsClientListByResourceGroupResponse contains the response from method JitRequestsClient.ListByResourceGroup.
type JitRequestsClientListByResourceGroupResponse struct {
	// List of JIT requests.
	JitRequestDefinitionListResult
}

// JitRequestsClientListBySubscriptionResponse contains the response from method JitRequestsClient.ListBySubscription.
type JitRequestsClientListBySubscriptionResponse struct {
	// List of JIT requests.
	JitRequestDefinitionListResult
}

// JitRequestsClientUpdateResponse contains the response from method JitRequestsClient.Update.
type JitRequestsClientUpdateResponse struct {
	// Information about JIT request definition.
	JitRequestDefinition
}
