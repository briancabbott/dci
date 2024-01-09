//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_List.json
func ExampleMoveResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMoveResourcesClient().NewListPager("rg1", "movecollection1", &armresourcemover.MoveResourcesClientListOptions{Filter: nil})
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
		// page.MoveResourceCollection = armresourcemover.MoveResourceCollection{
		// 	Value: []*armresourcemover.MoveResource{
		// 		{
		// 			Name: to.Ptr("moveresourcename1"),
		// 			Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1"),
		// 			Properties: &armresourcemover.MoveResourceProperties{
		// 				DependsOn: []*armresourcemover.MoveResourceDependency{
		// 					{
		// 						AutomaticResolution: &armresourcemover.AutomaticResolutionProperties{
		// 							MoveResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource2"),
		// 						},
		// 						ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastus/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
		// 						ResolutionStatus: to.Ptr("Resolved"),
		// 						ResolutionType: to.Ptr(armresourcemover.ResolutionTypeAutomatic),
		// 					},
		// 					{
		// 						AutomaticResolution: &armresourcemover.AutomaticResolutionProperties{
		// 							MoveResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource3"),
		// 						},
		// 						ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/EASTUS/providers/Microsoft.Compute/disks/eastusvm1_OsDisk_1_a88a2f0e8cf44fb3af24aa0f27101f83"),
		// 						ResolutionStatus: to.Ptr("Resolved"),
		// 						ResolutionType: to.Ptr(armresourcemover.ResolutionTypeAutomatic),
		// 				}},
		// 				ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
		// 					ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 					TargetResourceGroupName: to.Ptr("rg2"),
		// 					TargetResourceName: to.Ptr("eastusvm1"),
		// 					UserManagedIdentities: []*string{
		// 						to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
		// 					},
		// 					SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastus/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
		// 					SourceResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
		// 						ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 						TargetResourceGroupName: to.Ptr("rg1"),
		// 						TargetResourceName: to.Ptr("eastusvm1"),
		// 						TargetAvailabilityZone: to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
		// 						TargetVMSize: to.Ptr("Standard_B2s"),
		// 						UserManagedIdentities: []*string{
		// 							to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
		// 						},
		// 					},
		// 					SystemData: &armresourcemover.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.275Z"); return t}()),
		// 						CreatedBy: to.Ptr("user@microsoft.com"),
		// 						CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.364Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 						LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("moveresourcename2"),
		// 					Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource2"),
		// 					Properties: &armresourcemover.MoveResourceProperties{
		// 						DependsOn: []*armresourcemover.MoveResourceDependency{
		// 							{
		// 								ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastus/providers/Microsoft.Network/virtualNetworks/eastus-vnet"),
		// 								ManualResolution: &armresourcemover.ManualResolutionProperties{
		// 									TargetID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westus/providers/Microsoft.Network/virtualNetworks/westus-vnet"),
		// 								},
		// 								ResolutionStatus: to.Ptr("Resolved"),
		// 								ResolutionType: to.Ptr(armresourcemover.ResolutionTypeManual),
		// 						}},
		// 						SourceID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastus/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
		// 					},
		// 					SystemData: &armresourcemover.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-02T15:10:54.275Z"); return t}()),
		// 						CreatedBy: to.Ptr("user@microsoft.com"),
		// 						CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-02T17:16:24.364Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 						LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("moveresourcename3"),
		// 					Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource3"),
		// 					Properties: &armresourcemover.MoveResourceProperties{
		// 						SourceID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/EASTUS/providers/Microsoft.Compute/disks/eastusvm1_OsDisk_1_a88a2f0e8cf44fb3af24aa0f27101f83"),
		// 					},
		// 					SystemData: &armresourcemover.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-02T18:05:54.275Z"); return t}()),
		// 						CreatedBy: to.Ptr("user@microsoft.com"),
		// 						CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-02T20:16:24.364Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 						LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Create.json
func ExampleMoveResourcesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveResourcesClient().BeginCreate(ctx, "rg1", "movecollection1", "moveresourcename1", &armresourcemover.MoveResourcesClientBeginCreateOptions{Body: &armresourcemover.MoveResource{
		Properties: &armresourcemover.MoveResourceProperties{
			DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
				{
					ID:       to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
					TargetID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
				}},
			ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
				ResourceType:            to.Ptr("Microsoft.Compute/virtualMachines"),
				TargetResourceName:      to.Ptr("westusvm1"),
				TargetAvailabilitySetID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1"),
				TargetAvailabilityZone:  to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
				UserManagedIdentities: []*string{
					to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
			},
			SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
		},
	},
	})
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
	// res.MoveResource = armresourcemover.MoveResource{
	// 	Name: to.Ptr("moveresourcename1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1"),
	// 	Properties: &armresourcemover.MoveResourceProperties{
	// 		DependsOn: []*armresourcemover.MoveResourceDependency{
	// 			{
	// 				ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 				ResolutionStatus: to.Ptr("Unresolved"),
	// 				ResolutionType: to.Ptr(armresourcemover.ResolutionTypeAutomatic),
	// 		}},
	// 		DependsOnOverrides: []*armresourcemover.MoveResourceDependencyOverride{
	// 			{
	// 				ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 				TargetID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/westusRG/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
	// 			ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			TargetResourceGroupName: to.Ptr("rg1"),
	// 			TargetResourceName: to.Ptr("westusvm1"),
	// 			TargetAvailabilitySetID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/availabilitySets/avset1"),
	// 			TargetAvailabilityZone: to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
	// 			UserManagedIdentities: []*string{
	// 				to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
	// 			},
	// 			SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
	// 		},
	// 		SystemData: &armresourcemover.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.275Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@microsoft.com"),
	// 			CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.364Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Delete.json
func ExampleMoveResourcesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveResourcesClient().BeginDelete(ctx, "rg1", "movecollection1", "moveresourcename1", nil)
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveResources_Get.json
func ExampleMoveResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveResourcesClient().Get(ctx, "rg1", "movecollection1", "moveresourcename1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveResource = armresourcemover.MoveResource{
	// 	Name: to.Ptr("moveresourcename1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections/MoveResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1"),
	// 	Properties: &armresourcemover.MoveResourceProperties{
	// 		DependsOn: []*armresourcemover.MoveResourceDependency{
	// 			{
	// 				ID: to.Ptr("/subscriptions/c4488a3f-a7f7-4ad4-aa72-0e1f4d9c0756/resourceGroups/eastus/providers/Microsoft.Network/networkInterfaces/eastusvm140"),
	// 				IsOptional: to.Ptr("false"),
	// 				ResolutionStatus: to.Ptr("Unresolved"),
	// 				ResolutionType: to.Ptr(armresourcemover.ResolutionTypeAutomatic),
	// 		}},
	// 		ResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
	// 			ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 			TargetResourceGroupName: to.Ptr("rg2"),
	// 			TargetResourceName: to.Ptr("eastusvm1"),
	// 			UserManagedIdentities: []*string{
	// 				to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
	// 			},
	// 			SourceID: to.Ptr("/subscriptions/subid/resourceGroups/eastus/providers/Microsoft.Compute/virtualMachines/eastusvm1"),
	// 			SourceResourceSettings: &armresourcemover.VirtualMachineResourceSettings{
	// 				ResourceType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 				TargetResourceGroupName: to.Ptr("rg1"),
	// 				TargetResourceName: to.Ptr("eastusvm1"),
	// 				TargetAvailabilityZone: to.Ptr(armresourcemover.TargetAvailabilityZoneTwo),
	// 				TargetVMSize: to.Ptr("Standard_B2s"),
	// 				UserManagedIdentities: []*string{
	// 					to.Ptr("/subscriptions/subid/resourceGroups/eastusRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/umi1")},
	// 				},
	// 			},
	// 			SystemData: &armresourcemover.SystemData{
	// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.275Z"); return t}()),
	// 				CreatedBy: to.Ptr("user@microsoft.com"),
	// 				CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.364Z"); return t}()),
	// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 				LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 			},
	// 		}
}
