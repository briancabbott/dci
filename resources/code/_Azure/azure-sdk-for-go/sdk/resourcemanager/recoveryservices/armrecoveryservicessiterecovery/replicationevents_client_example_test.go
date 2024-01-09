//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationEvents_List.json
func ExampleReplicationEventsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationEventsClient().NewListPager("vault1", "resourceGroupPS1", &armrecoveryservicessiterecovery.ReplicationEventsClientListOptions{Filter: nil})
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
		// page.EventCollection = armrecoveryservicessiterecovery.EventCollection{
		// 	Value: []*armrecoveryservicessiterecovery.Event{
		// 		{
		// 			Name: to.Ptr("JobStatusMonitoringEvent;9091989892524070155_4ed6f1a6-9b6d-4048-9079-1307dd24b814"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationEvents"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationEvents/JobStatusMonitoringEvent;9091989892524070155_4ed6f1a6-9b6d-4048-9079-1307dd24b814"),
		// 			Properties: &armrecoveryservicessiterecovery.EventProperties{
		// 				Description: to.Ptr("TestFailover - Failed"),
		// 				AffectedObjectFriendlyName: to.Ptr("vm1"),
		// 				EventCode: to.Ptr("d32574f6-f59e-4545-b5ac-bc88d545f089"),
		// 				EventSpecificDetails: &armrecoveryservicessiterecovery.JobStatusEventDetails{
		// 					InstanceType: to.Ptr("JobStatus"),
		// 				},
		// 				EventType: to.Ptr("JobStatus"),
		// 				FabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
		// 				HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
		// 					{
		// 						ErrorCode: to.Ptr("499"),
		// 						ErrorMessage: to.Ptr("An internal error occurred."),
		// 						PossibleCauses: to.Ptr("The operation failed due to an internal error."),
		// 						RecommendedAction: to.Ptr("Retry the last action. If the issue persists, contact Support."),
		// 				}},
		// 				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureEventDetails{
		// 					InstanceType: to.Ptr("HyperVReplicaAzure"),
		// 				},
		// 				Severity: to.Ptr("Critical"),
		// 				TimeOfOccurrence: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T16:00:33.070Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("VmMonitoringEvent;9091989947769704276_516de684-0079-48f7-b44b-882923268654"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationEvents"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationEvents/VmMonitoringEvent;9091989947769704276_516de684-0079-48f7-b44b-882923268654"),
		// 			Properties: &armrecoveryservicessiterecovery.EventProperties{
		// 				Description: to.Ptr("Virtual machine health is OK"),
		// 				AffectedObjectFriendlyName: to.Ptr("vm1"),
		// 				EventCode: to.Ptr("d9a07b07-c7b5-49ca-ab6c-6926596dfe47"),
		// 				EventSpecificDetails: &armrecoveryservicessiterecovery.JobStatusEventDetails{
		// 					InstanceType: to.Ptr("JobStatus"),
		// 				},
		// 				EventType: to.Ptr("VmHealth"),
		// 				FabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
		// 				HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
		// 				},
		// 				ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureEventDetails{
		// 					InstanceType: to.Ptr("HyperVReplicaAzure"),
		// 				},
		// 				Severity: to.Ptr("OK"),
		// 				TimeOfOccurrence: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T14:28:28.507Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationEvents_Get.json
func ExampleReplicationEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationEventsClient().Get(ctx, "vault1", "resourceGroupPS1", "654b71d0-b2ce-4e6e-a861-98528d4bd375", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Event = armrecoveryservicessiterecovery.Event{
	// 	Name: to.Ptr("VmMonitoringEvent;9091989947769704276_516de684-0079-48f7-b44b-882923268654"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationEvents"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationEvents/VmMonitoringEvent;9091989947769704276_516de684-0079-48f7-b44b-882923268654"),
	// 	Properties: &armrecoveryservicessiterecovery.EventProperties{
	// 		Description: to.Ptr("Virtual machine health is OK"),
	// 		AffectedObjectFriendlyName: to.Ptr("vm1"),
	// 		EventCode: to.Ptr("d9a07b07-c7b5-49ca-ab6c-6926596dfe47"),
	// 		EventSpecificDetails: &armrecoveryservicessiterecovery.JobStatusEventDetails{
	// 			InstanceType: to.Ptr("JobStatus"),
	// 		},
	// 		EventType: to.Ptr("VmHealth"),
	// 		FabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 		HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.HyperVReplicaAzureEventDetails{
	// 			InstanceType: to.Ptr("HyperVReplicaAzure"),
	// 		},
	// 		Severity: to.Ptr("OK"),
	// 		TimeOfOccurrence: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-02T14:28:28.507Z"); return t}()),
	// 	},
	// }
}
