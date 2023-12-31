//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

// DiskPoolZonesClientListResponse contains the response from method DiskPoolZonesClient.NewListPager.
type DiskPoolZonesClientListResponse struct {
	// List Disk Pool skus operation response.
	DiskPoolZoneListResult
}

// DiskPoolsClientCreateOrUpdateResponse contains the response from method DiskPoolsClient.BeginCreateOrUpdate.
type DiskPoolsClientCreateOrUpdateResponse struct {
	// Response for Disk Pool request.
	DiskPool
}

// DiskPoolsClientDeallocateResponse contains the response from method DiskPoolsClient.BeginDeallocate.
type DiskPoolsClientDeallocateResponse struct {
	// placeholder for future response values
}

// DiskPoolsClientDeleteResponse contains the response from method DiskPoolsClient.BeginDelete.
type DiskPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// DiskPoolsClientGetResponse contains the response from method DiskPoolsClient.Get.
type DiskPoolsClientGetResponse struct {
	// Response for Disk Pool request.
	DiskPool
}

// DiskPoolsClientListByResourceGroupResponse contains the response from method DiskPoolsClient.NewListByResourceGroupPager.
type DiskPoolsClientListByResourceGroupResponse struct {
	// List of Disk Pools
	DiskPoolListResult
}

// DiskPoolsClientListBySubscriptionResponse contains the response from method DiskPoolsClient.NewListBySubscriptionPager.
type DiskPoolsClientListBySubscriptionResponse struct {
	// List of Disk Pools
	DiskPoolListResult
}

// DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method DiskPoolsClient.NewListOutboundNetworkDependenciesEndpointsPager.
type DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse struct {
	// Collection of Outbound Environment Endpoints
	OutboundEnvironmentEndpointList
}

// DiskPoolsClientStartResponse contains the response from method DiskPoolsClient.BeginStart.
type DiskPoolsClientStartResponse struct {
	// placeholder for future response values
}

// DiskPoolsClientUpdateResponse contains the response from method DiskPoolsClient.BeginUpdate.
type DiskPoolsClientUpdateResponse struct {
	// Response for Disk Pool request.
	DiskPool
}

// DiskPoolsClientUpgradeResponse contains the response from method DiskPoolsClient.BeginUpgrade.
type DiskPoolsClientUpgradeResponse struct {
	// placeholder for future response values
}

// IscsiTargetsClientCreateOrUpdateResponse contains the response from method IscsiTargetsClient.BeginCreateOrUpdate.
type IscsiTargetsClientCreateOrUpdateResponse struct {
	// Response for iSCSI Target requests.
	IscsiTarget
}

// IscsiTargetsClientDeleteResponse contains the response from method IscsiTargetsClient.BeginDelete.
type IscsiTargetsClientDeleteResponse struct {
	// placeholder for future response values
}

// IscsiTargetsClientGetResponse contains the response from method IscsiTargetsClient.Get.
type IscsiTargetsClientGetResponse struct {
	// Response for iSCSI Target requests.
	IscsiTarget
}

// IscsiTargetsClientListByDiskPoolResponse contains the response from method IscsiTargetsClient.NewListByDiskPoolPager.
type IscsiTargetsClientListByDiskPoolResponse struct {
	// List of iSCSI Targets.
	IscsiTargetList
}

// IscsiTargetsClientUpdateResponse contains the response from method IscsiTargetsClient.BeginUpdate.
type IscsiTargetsClientUpdateResponse struct {
	// Response for iSCSI Target requests.
	IscsiTarget
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// List of operations supported by the RP.
	OperationListResult
}

// ResourceSKUsClientListResponse contains the response from method ResourceSKUsClient.NewListPager.
type ResourceSKUsClientListResponse struct {
	// List Disk Pool skus operation response.
	ResourceSKUListResult
}
