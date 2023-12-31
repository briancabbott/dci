//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis

// ChangesClientListChangesByResourceGroupResponse contains the response from method ChangesClient.NewListChangesByResourceGroupPager.
type ChangesClientListChangesByResourceGroupResponse struct {
	// The list of detected changes.
	ChangeList
}

// ChangesClientListChangesBySubscriptionResponse contains the response from method ChangesClient.NewListChangesBySubscriptionPager.
type ChangesClientListChangesBySubscriptionResponse struct {
	// The list of detected changes.
	ChangeList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The resource provider operation list.
	ResourceProviderOperationList
}

// ResourceChangesClientListResponse contains the response from method ResourceChangesClient.NewListPager.
type ResourceChangesClientListResponse struct {
	// The list of detected changes.
	ChangeList
}
