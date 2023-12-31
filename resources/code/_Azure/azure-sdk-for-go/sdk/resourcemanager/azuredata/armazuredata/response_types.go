//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredata

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list SQL operations.
	OperationListResult
}

// SQLServerRegistrationsClientCreateOrUpdateResponse contains the response from method SQLServerRegistrationsClient.CreateOrUpdate.
type SQLServerRegistrationsClientCreateOrUpdateResponse struct {
	// A SQL server registration.
	SQLServerRegistration
}

// SQLServerRegistrationsClientDeleteResponse contains the response from method SQLServerRegistrationsClient.Delete.
type SQLServerRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLServerRegistrationsClientGetResponse contains the response from method SQLServerRegistrationsClient.Get.
type SQLServerRegistrationsClientGetResponse struct {
	// A SQL server registration.
	SQLServerRegistration
}

// SQLServerRegistrationsClientListByResourceGroupResponse contains the response from method SQLServerRegistrationsClient.NewListByResourceGroupPager.
type SQLServerRegistrationsClientListByResourceGroupResponse struct {
	// Server
	SQLServerRegistrationListResult
}

// SQLServerRegistrationsClientListResponse contains the response from method SQLServerRegistrationsClient.NewListPager.
type SQLServerRegistrationsClientListResponse struct {
	// Server
	SQLServerRegistrationListResult
}

// SQLServerRegistrationsClientUpdateResponse contains the response from method SQLServerRegistrationsClient.Update.
type SQLServerRegistrationsClientUpdateResponse struct {
	// A SQL server registration.
	SQLServerRegistration
}

// SQLServersClientCreateOrUpdateResponse contains the response from method SQLServersClient.CreateOrUpdate.
type SQLServersClientCreateOrUpdateResponse struct {
	// A SQL server.
	SQLServer
}

// SQLServersClientDeleteResponse contains the response from method SQLServersClient.Delete.
type SQLServersClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLServersClientGetResponse contains the response from method SQLServersClient.Get.
type SQLServersClientGetResponse struct {
	// A SQL server.
	SQLServer
}

// SQLServersClientListByResourceGroupResponse contains the response from method SQLServersClient.NewListByResourceGroupPager.
type SQLServersClientListByResourceGroupResponse struct {
	// A list of SQL servers.
	SQLServerListResult
}
