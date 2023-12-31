//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

// CdnPeeringPrefixesClientListResponse contains the response from method CdnPeeringPrefixesClient.NewListPager.
type CdnPeeringPrefixesClientListResponse struct {
	// The paginated list of CDN peering prefixes.
	CdnPeeringPrefixListResult
}

// ConnectionMonitorTestsClientCreateOrUpdateResponse contains the response from method ConnectionMonitorTestsClient.CreateOrUpdate.
type ConnectionMonitorTestsClientCreateOrUpdateResponse struct {
	// The Connection Monitor Test class.
	ConnectionMonitorTest
}

// ConnectionMonitorTestsClientDeleteResponse contains the response from method ConnectionMonitorTestsClient.Delete.
type ConnectionMonitorTestsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConnectionMonitorTestsClientGetResponse contains the response from method ConnectionMonitorTestsClient.Get.
type ConnectionMonitorTestsClientGetResponse struct {
	// The Connection Monitor Test class.
	ConnectionMonitorTest
}

// ConnectionMonitorTestsClientListByPeeringServiceResponse contains the response from method ConnectionMonitorTestsClient.NewListByPeeringServicePager.
type ConnectionMonitorTestsClientListByPeeringServiceResponse struct {
	// The paginated list of [T].
	ConnectionMonitorTestListResult
}

// LegacyPeeringsClientListResponse contains the response from method LegacyPeeringsClient.NewListPager.
type LegacyPeeringsClientListResponse struct {
	// The paginated list of peerings.
	ListResult
}

// LocationsClientListResponse contains the response from method LocationsClient.NewListPager.
type LocationsClientListResponse struct {
	// The paginated list of peering locations.
	LocationListResult
}

// LookingGlassClientInvokeResponse contains the response from method LookingGlassClient.Invoke.
type LookingGlassClientInvokeResponse struct {
	// Looking glass output model
	LookingGlassOutput
}

// ManagementClientCheckServiceProviderAvailabilityResponse contains the response from method ManagementClient.CheckServiceProviderAvailability.
type ManagementClientCheckServiceProviderAvailabilityResponse struct {
	Value *Enum0
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The paginated list of peering API operations.
	OperationListResult
}

// PeerAsnsClientCreateOrUpdateResponse contains the response from method PeerAsnsClient.CreateOrUpdate.
type PeerAsnsClientCreateOrUpdateResponse struct {
	// The essential information related to the peer's ASN.
	PeerAsn
}

// PeerAsnsClientDeleteResponse contains the response from method PeerAsnsClient.Delete.
type PeerAsnsClientDeleteResponse struct {
	// placeholder for future response values
}

// PeerAsnsClientGetResponse contains the response from method PeerAsnsClient.Get.
type PeerAsnsClientGetResponse struct {
	// The essential information related to the peer's ASN.
	PeerAsn
}

// PeerAsnsClientListBySubscriptionResponse contains the response from method PeerAsnsClient.NewListBySubscriptionPager.
type PeerAsnsClientListBySubscriptionResponse struct {
	// The paginated list of peer ASNs.
	PeerAsnListResult
}

// PeeringsClientCreateOrUpdateResponse contains the response from method PeeringsClient.CreateOrUpdate.
type PeeringsClientCreateOrUpdateResponse struct {
	// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
	Peering
}

// PeeringsClientDeleteResponse contains the response from method PeeringsClient.Delete.
type PeeringsClientDeleteResponse struct {
	// placeholder for future response values
}

// PeeringsClientGetResponse contains the response from method PeeringsClient.Get.
type PeeringsClientGetResponse struct {
	// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
	Peering
}

// PeeringsClientListByResourceGroupResponse contains the response from method PeeringsClient.NewListByResourceGroupPager.
type PeeringsClientListByResourceGroupResponse struct {
	// The paginated list of peerings.
	ListResult
}

// PeeringsClientListBySubscriptionResponse contains the response from method PeeringsClient.NewListBySubscriptionPager.
type PeeringsClientListBySubscriptionResponse struct {
	// The paginated list of peerings.
	ListResult
}

// PeeringsClientUpdateResponse contains the response from method PeeringsClient.Update.
type PeeringsClientUpdateResponse struct {
	// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
	Peering
}

// PrefixesClientCreateOrUpdateResponse contains the response from method PrefixesClient.CreateOrUpdate.
type PrefixesClientCreateOrUpdateResponse struct {
	// The peering service prefix class.
	ServicePrefix
}

// PrefixesClientDeleteResponse contains the response from method PrefixesClient.Delete.
type PrefixesClientDeleteResponse struct {
	// placeholder for future response values
}

// PrefixesClientGetResponse contains the response from method PrefixesClient.Get.
type PrefixesClientGetResponse struct {
	// The peering service prefix class.
	ServicePrefix
}

// PrefixesClientListByPeeringServiceResponse contains the response from method PrefixesClient.NewListByPeeringServicePager.
type PrefixesClientListByPeeringServiceResponse struct {
	// The paginated list of peering service prefixes.
	ServicePrefixListResult
}

// ReceivedRoutesClientListByPeeringResponse contains the response from method ReceivedRoutesClient.NewListByPeeringPager.
type ReceivedRoutesClientListByPeeringResponse struct {
	// The paginated list of received routes for the peering.
	ReceivedRouteListResult
}

// RegisteredAsnsClientCreateOrUpdateResponse contains the response from method RegisteredAsnsClient.CreateOrUpdate.
type RegisteredAsnsClientCreateOrUpdateResponse struct {
	// The customer's ASN that is registered by the peering service provider.
	RegisteredAsn
}

// RegisteredAsnsClientDeleteResponse contains the response from method RegisteredAsnsClient.Delete.
type RegisteredAsnsClientDeleteResponse struct {
	// placeholder for future response values
}

// RegisteredAsnsClientGetResponse contains the response from method RegisteredAsnsClient.Get.
type RegisteredAsnsClientGetResponse struct {
	// The customer's ASN that is registered by the peering service provider.
	RegisteredAsn
}

// RegisteredAsnsClientListByPeeringResponse contains the response from method RegisteredAsnsClient.NewListByPeeringPager.
type RegisteredAsnsClientListByPeeringResponse struct {
	// The paginated list of peering registered ASNs.
	RegisteredAsnListResult
}

// RegisteredPrefixesClientCreateOrUpdateResponse contains the response from method RegisteredPrefixesClient.CreateOrUpdate.
type RegisteredPrefixesClientCreateOrUpdateResponse struct {
	// The customer's prefix that is registered by the peering service provider.
	RegisteredPrefix
}

// RegisteredPrefixesClientDeleteResponse contains the response from method RegisteredPrefixesClient.Delete.
type RegisteredPrefixesClientDeleteResponse struct {
	// placeholder for future response values
}

// RegisteredPrefixesClientGetResponse contains the response from method RegisteredPrefixesClient.Get.
type RegisteredPrefixesClientGetResponse struct {
	// The customer's prefix that is registered by the peering service provider.
	RegisteredPrefix
}

// RegisteredPrefixesClientListByPeeringResponse contains the response from method RegisteredPrefixesClient.NewListByPeeringPager.
type RegisteredPrefixesClientListByPeeringResponse struct {
	// The paginated list of peering registered prefixes.
	RegisteredPrefixListResult
}

// ServiceCountriesClientListResponse contains the response from method ServiceCountriesClient.NewListPager.
type ServiceCountriesClientListResponse struct {
	// The paginated list of peering service countries.
	ServiceCountryListResult
}

// ServiceLocationsClientListResponse contains the response from method ServiceLocationsClient.NewListPager.
type ServiceLocationsClientListResponse struct {
	// The paginated list of peering service locations.
	ServiceLocationListResult
}

// ServiceProvidersClientListResponse contains the response from method ServiceProvidersClient.NewListPager.
type ServiceProvidersClientListResponse struct {
	// The paginated list of peering service providers.
	ServiceProviderListResult
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// Peering Service
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// Peering Service
	Service
}

// ServicesClientInitializeConnectionMonitorResponse contains the response from method ServicesClient.InitializeConnectionMonitor.
type ServicesClientInitializeConnectionMonitorResponse struct {
	// placeholder for future response values
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// The paginated list of peering services.
	ServiceListResult
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.NewListBySubscriptionPager.
type ServicesClientListBySubscriptionResponse struct {
	// The paginated list of peering services.
	ServiceListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	// Peering Service
	Service
}
