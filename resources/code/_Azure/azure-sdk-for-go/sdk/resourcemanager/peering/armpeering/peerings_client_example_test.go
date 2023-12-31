//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/GetPeering.json
func ExamplePeeringsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().Get(ctx, "rgName", "peeringName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindExchange),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Exchange: &armpeering.PropertiesExchange{
	// 			Connections: []*armpeering.ExchangeConnection{
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::1"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 				},
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::2"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 			}},
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Basic_Exchange_Free"),
	// 		Family: to.Ptr(armpeering.FamilyExchange),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierBasic),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreateDirectPeering.json
func ExamplePeeringsClient_CreateOrUpdate_createADirectPeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().CreateOrUpdate(ctx, "rgName", "peeringName", armpeering.Peering{
		Kind:     to.Ptr(armpeering.KindDirect),
		Location: to.Ptr("eastus"),
		Properties: &armpeering.Properties{
			Direct: &armpeering.PropertiesDirect{
				Connections: []*armpeering.DirectConnection{
					{
						BandwidthInMbps: to.Ptr[int32](10000),
						BgpSession: &armpeering.BgpSession{
							MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
							MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
							MD5AuthenticationKey:    to.Ptr("test-md5-auth-key"),
							SessionPrefixV4:         to.Ptr("192.168.0.0/31"),
							SessionPrefixV6:         to.Ptr("fd00::0/127"),
						},
						ConnectionIdentifier:   to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
						PeeringDBFacilityID:    to.Ptr[int32](99999),
						SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
						UseForPeeringService:   to.Ptr(false),
					},
					{
						BandwidthInMbps:        to.Ptr[int32](10000),
						ConnectionIdentifier:   to.Ptr("8AB00818-D533-4504-A25A-03A17F61201C"),
						PeeringDBFacilityID:    to.Ptr[int32](99999),
						SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderMicrosoft),
						UseForPeeringService:   to.Ptr(true),
					}},
				DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
				PeerAsn: &armpeering.SubResource{
					ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
				},
			},
			PeeringLocation: to.Ptr("peeringLocation0"),
		},
		SKU: &armpeering.SKU{
			Name: to.Ptr("Basic_Direct_Free"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindDirect),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Direct: &armpeering.PropertiesDirect{
	// 			Connections: []*armpeering.DirectConnection{
	// 				{
	// 					BandwidthInMbps: to.Ptr[int32](10000),
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						SessionPrefixV4: to.Ptr("192.168.0.0/31"),
	// 						SessionPrefixV6: to.Ptr("fd00::0/127"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateProvisioningFailed),
	// 					ErrorMessage: to.Ptr("IPv4 address is already configured with a different ASN"),
	// 					MicrosoftTrackingID: to.Ptr("test-microsoft-reference-id-1"),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 					ProvisionedBandwidthInMbps: to.Ptr[int32](10000),
	// 					SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
	// 					UseForPeeringService: to.Ptr(false),
	// 				},
	// 				{
	// 					BandwidthInMbps: to.Ptr[int32](10000),
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						SessionPrefixV4: to.Ptr("192.168.1.0/31"),
	// 						SessionPrefixV6: to.Ptr("fd00::2/127"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("8AB00818-D533-4504-A25A-03A17F61201C"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					MicrosoftTrackingID: to.Ptr("test-microsoft-reference-id-2"),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 					ProvisionedBandwidthInMbps: to.Ptr[int32](10000),
	// 					SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderMicrosoft),
	// 					UseForPeeringService: to.Ptr(true),
	// 			}},
	// 			DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeEdge),
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 			UseForPeeringService: to.Ptr(true),
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Basic_Direct_Free"),
	// 		Family: to.Ptr(armpeering.FamilyDirect),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierBasic),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreatePeeringWithExchangeRouteServer.json
func ExamplePeeringsClient_CreateOrUpdate_createAPeeringWithExchangeRouteServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().CreateOrUpdate(ctx, "rgName", "peeringName", armpeering.Peering{
		Kind:     to.Ptr(armpeering.KindDirect),
		Location: to.Ptr("eastus"),
		Properties: &armpeering.Properties{
			Direct: &armpeering.PropertiesDirect{
				Connections: []*armpeering.DirectConnection{
					{
						BandwidthInMbps: to.Ptr[int32](10000),
						BgpSession: &armpeering.BgpSession{
							MaxPrefixesAdvertisedV4:     to.Ptr[int32](1000),
							MaxPrefixesAdvertisedV6:     to.Ptr[int32](100),
							MicrosoftSessionIPv4Address: to.Ptr("192.168.0.123"),
							PeerSessionIPv4Address:      to.Ptr("192.168.0.234"),
							SessionPrefixV4:             to.Ptr("192.168.0.0/24"),
						},
						ConnectionIdentifier:   to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
						PeeringDBFacilityID:    to.Ptr[int32](99999),
						SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
						UseForPeeringService:   to.Ptr(true),
					}},
				DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeIxRs),
				PeerAsn: &armpeering.SubResource{
					ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
				},
			},
			PeeringLocation: to.Ptr("peeringLocation0"),
		},
		SKU: &armpeering.SKU{
			Name: to.Ptr("Premium_Direct_Free"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindDirect),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Direct: &armpeering.PropertiesDirect{
	// 			Connections: []*armpeering.DirectConnection{
	// 				{
	// 					BandwidthInMbps: to.Ptr[int32](10000),
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.0.123"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.0.234"),
	// 						SessionPrefixV4: to.Ptr("192.168.0.0/24"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("5F4CB5C7-6B43-4444-9338-9ABC72606C16"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 					ProvisionedBandwidthInMbps: to.Ptr[int32](10000),
	// 					SessionAddressProvider: to.Ptr(armpeering.SessionAddressProviderPeer),
	// 					UseForPeeringService: to.Ptr(true),
	// 			}},
	// 			DirectPeeringType: to.Ptr(armpeering.DirectPeeringTypeIxRs),
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 			UseForPeeringService: to.Ptr(true),
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Premium_Direct_Free"),
	// 		Family: to.Ptr(armpeering.FamilyDirect),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierPremium),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/CreateExchangePeering.json
func ExamplePeeringsClient_CreateOrUpdate_createAnExchangePeering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().CreateOrUpdate(ctx, "rgName", "peeringName", armpeering.Peering{
		Kind:     to.Ptr(armpeering.KindExchange),
		Location: to.Ptr("eastus"),
		Properties: &armpeering.Properties{
			Exchange: &armpeering.PropertiesExchange{
				Connections: []*armpeering.ExchangeConnection{
					{
						BgpSession: &armpeering.BgpSession{
							MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
							MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
							MD5AuthenticationKey:    to.Ptr("test-md5-auth-key"),
							PeerSessionIPv4Address:  to.Ptr("192.168.2.1"),
							PeerSessionIPv6Address:  to.Ptr("fd00::1"),
						},
						ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
						PeeringDBFacilityID:  to.Ptr[int32](99999),
					},
					{
						BgpSession: &armpeering.BgpSession{
							MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
							MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
							MD5AuthenticationKey:    to.Ptr("test-md5-auth-key"),
							PeerSessionIPv4Address:  to.Ptr("192.168.2.2"),
							PeerSessionIPv6Address:  to.Ptr("fd00::2"),
						},
						ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
						PeeringDBFacilityID:  to.Ptr[int32](99999),
					}},
				PeerAsn: &armpeering.SubResource{
					ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
				},
			},
			PeeringLocation: to.Ptr("peeringLocation0"),
		},
		SKU: &armpeering.SKU{
			Name: to.Ptr("Basic_Exchange_Free"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindExchange),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Exchange: &armpeering.PropertiesExchange{
	// 			Connections: []*armpeering.ExchangeConnection{
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::1"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateProvisioningFailed),
	// 					ErrorMessage: to.Ptr("IPv4 address is already configured with a different ASN"),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 				},
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::2"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 			}},
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Basic_Exchange_Free"),
	// 		Family: to.Ptr(armpeering.FamilyExchange),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierBasic),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/DeletePeering.json
func ExamplePeeringsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPeeringsClient().Delete(ctx, "rgName", "peeringName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/UpdatePeeringTags.json
func ExamplePeeringsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPeeringsClient().Update(ctx, "rgName", "peeringName", armpeering.ResourceTags{
		Tags: map[string]*string{
			"tag0": to.Ptr("value0"),
			"tag1": to.Ptr("value1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Peering = armpeering.Peering{
	// 	Name: to.Ptr("peeringName"),
	// 	Type: to.Ptr("Microsoft.Peering/peerings"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
	// 	Kind: to.Ptr(armpeering.KindExchange),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armpeering.Properties{
	// 		Exchange: &armpeering.PropertiesExchange{
	// 			Connections: []*armpeering.ExchangeConnection{
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::1"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 				},
	// 				{
	// 					BgpSession: &armpeering.BgpSession{
	// 						MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
	// 						MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
	// 						MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
	// 						MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
	// 						MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
	// 						PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
	// 						PeerSessionIPv6Address: to.Ptr("fd00::2"),
	// 						SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
	// 						SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
	// 					},
	// 					ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
	// 					ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
	// 					PeeringDBFacilityID: to.Ptr[int32](99999),
	// 			}},
	// 			PeerAsn: &armpeering.SubResource{
	// 				ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
	// 			},
	// 		},
	// 		PeeringLocation: to.Ptr("peeringLocation0"),
	// 		ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armpeering.SKU{
	// 		Name: to.Ptr("Basic_Exchange_Free"),
	// 		Family: to.Ptr(armpeering.FamilyExchange),
	// 		Size: to.Ptr(armpeering.SizeFree),
	// 		Tier: to.Ptr(armpeering.TierBasic),
	// 	},
	// 	Tags: map[string]*string{
	// 		"tag0": to.Ptr("value0"),
	// 		"tag1": to.Ptr("value1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListPeeringsByResourceGroup.json
func ExamplePeeringsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPeeringsClient().NewListByResourceGroupPager("rgName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListResult = armpeering.ListResult{
		// 	Value: []*armpeering.Peering{
		// 		{
		// 			Name: to.Ptr("peeringName"),
		// 			Type: to.Ptr("Microsoft.Peering/peerings"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
		// 			Kind: to.Ptr(armpeering.KindExchange),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armpeering.Properties{
		// 				Exchange: &armpeering.PropertiesExchange{
		// 					Connections: []*armpeering.ExchangeConnection{
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::1"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 						},
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::2"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 					}},
		// 					PeerAsn: &armpeering.SubResource{
		// 						ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
		// 					},
		// 				},
		// 				PeeringLocation: to.Ptr("peeringLocation0"),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armpeering.SKU{
		// 				Name: to.Ptr("Basic_Exchange_Free"),
		// 				Family: to.Ptr(armpeering.FamilyExchange),
		// 				Size: to.Ptr(armpeering.SizeFree),
		// 				Tier: to.Ptr(armpeering.TierBasic),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/peering/resource-manager/Microsoft.Peering/stable/2022-01-01/examples/ListPeeringsBySubscription.json
func ExamplePeeringsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpeering.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPeeringsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ListResult = armpeering.ListResult{
		// 	Value: []*armpeering.Peering{
		// 		{
		// 			Name: to.Ptr("peeringName"),
		// 			Type: to.Ptr("Microsoft.Peering/peerings"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.Peering/peerings/peeringName"),
		// 			Kind: to.Ptr(armpeering.KindExchange),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armpeering.Properties{
		// 				Exchange: &armpeering.PropertiesExchange{
		// 					Connections: []*armpeering.ExchangeConnection{
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.1"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:1"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.1"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::1"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CE495334-0E94-4E51-8164-8116D6CD284D"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 						},
		// 						{
		// 							BgpSession: &armpeering.BgpSession{
		// 								MaxPrefixesAdvertisedV4: to.Ptr[int32](1000),
		// 								MaxPrefixesAdvertisedV6: to.Ptr[int32](100),
		// 								MD5AuthenticationKey: to.Ptr("test-md5-auth-key"),
		// 								MicrosoftSessionIPv4Address: to.Ptr("192.168.3.2"),
		// 								MicrosoftSessionIPv6Address: to.Ptr("fd00::1:2"),
		// 								PeerSessionIPv4Address: to.Ptr("192.168.2.2"),
		// 								PeerSessionIPv6Address: to.Ptr("fd00::2"),
		// 								SessionStateV4: to.Ptr(armpeering.SessionStateV4Established),
		// 								SessionStateV6: to.Ptr(armpeering.SessionStateV6Established),
		// 							},
		// 							ConnectionIdentifier: to.Ptr("CDD8E673-CB07-47E6-84DE-3739F778762B"),
		// 							ConnectionState: to.Ptr(armpeering.ConnectionStateActive),
		// 							PeeringDBFacilityID: to.Ptr[int32](99999),
		// 					}},
		// 					PeerAsn: &armpeering.SubResource{
		// 						ID: to.Ptr("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
		// 					},
		// 				},
		// 				PeeringLocation: to.Ptr("peeringLocation0"),
		// 				ProvisioningState: to.Ptr(armpeering.ProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armpeering.SKU{
		// 				Name: to.Ptr("Basic_Exchange_Free"),
		// 				Family: to.Ptr(armpeering.FamilyExchange),
		// 				Size: to.Ptr(armpeering.SizeFree),
		// 				Tier: to.Ptr(armpeering.TierBasic),
		// 			},
		// 	}},
		// }
	}
}
