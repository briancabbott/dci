//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Create_MaximumSet_Gen.json
func ExampleNetworkRacksClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkRacksClient().BeginCreate(ctx, "example-rg", "example-rack", armmanagednetworkfabric.NetworkRack{
		Location: to.Ptr("eastuseuap"),
		Tags: map[string]*string{
			"keyID": to.Ptr("keyValue"),
		},
		Properties: &armmanagednetworkfabric.NetworkRackProperties{
			Annotation:      to.Ptr("annotation"),
			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
			NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkRack = armmanagednetworkfabric.NetworkRack{
	// 	Name: to.Ptr("example-rack"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkRacks"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@email.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("keyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkRackProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		NetworkDevices: []*string{
	// 			to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice")},
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
	// 			NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Get_MaximumSet_Gen.json
func ExampleNetworkRacksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkRacksClient().Get(ctx, "example-rg", "example-rack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkRack = armmanagednetworkfabric.NetworkRack{
	// 	Name: to.Ptr("example-rack"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkRacks"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@email.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("keyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkRackProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		NetworkDevices: []*string{
	// 			to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice")},
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
	// 			NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Update_MaximumSet_Gen.json
func ExampleNetworkRacksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkRacksClient().BeginUpdate(ctx, "example-rg", "example-rack", armmanagednetworkfabric.TagsUpdate{
		Tags: map[string]*string{
			"keyID": to.Ptr("keyValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkRack = armmanagednetworkfabric.NetworkRack{
	// 	Name: to.Ptr("example-rack"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/networkRacks"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@mail.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@email.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastuseuap"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("keyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.NetworkRackProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		NetworkDevices: []*string{
	// 			to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice")},
	// 			NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
	// 			NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_Delete_MaximumSet_Gen.json
func ExampleNetworkRacksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkRacksClient().BeginDelete(ctx, "example-rg", "example-rack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_ListByResourceGroup_MaximumSet_Gen.json
func ExampleNetworkRacksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkRacksClient().NewListByResourceGroupPager("example-rg", nil)
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
		// page.NetworkRacksListResult = armmanagednetworkfabric.NetworkRacksListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkRack{
		// 		{
		// 			Name: to.Ptr("example-rack"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkRacks"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@mail.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@email.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastuseuap"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("keyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkRackProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				NetworkDevices: []*string{
		// 					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice")},
		// 					NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
		// 					NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkRacks_ListBySubscription_MaximumSet_Gen.json
func ExampleNetworkRacksClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkRacksClient().NewListBySubscriptionPager(nil)
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
		// page.NetworkRacksListResult = armmanagednetworkfabric.NetworkRacksListResult{
		// 	Value: []*armmanagednetworkfabric.NetworkRack{
		// 		{
		// 			Name: to.Ptr("example-rack"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/networkRacks"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkRacks/example-rack"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@mail.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T06:00:50.441Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@email.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastuseuap"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("keyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.NetworkRackProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				NetworkDevices: []*string{
		// 					to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkDevices/example-networkDevice")},
		// 					NetworkFabricID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabrics/example-networkFabric"),
		// 					NetworkRackType: to.Ptr(armmanagednetworkfabric.NetworkRackTypeAggregate),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
