//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Get.json
func ExampleReplicationExtensionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationExtensionClient().Get(ctx, "rgrecoveryservicesdatareplication", "4", "g16yjJ", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicationExtensionModel = armrecoveryservicesdatareplication.ReplicationExtensionModel{
	// 	Name: to.Ptr("xvjffbiecsd"),
	// 	Type: to.Ptr("miadbgilpheilzfoonveznybthgdwh"),
	// 	ID: to.Ptr("awu"),
	// 	Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
	// 		CustomProperties: &armrecoveryservicesdatareplication.ReplicationExtensionModelCustomProperties{
	// 			InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
	// 		},
	// 		ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armrecoveryservicesdatareplication.ReplicationExtensionModelSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 		CreatedBy: to.Ptr("lagtinqhksctfdxmfbpf"),
	// 		CreatedByType: to.Ptr("dsqllpglanwztdmisrknjtqz"),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("suwjpejlaya"),
	// 		LastModifiedByType: to.Ptr("nrfjuyghtbivwihr"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Create.json
func ExampleReplicationExtensionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationExtensionClient().BeginCreate(ctx, "rgrecoveryservicesdatareplication", "4", "g16yjJ", armrecoveryservicesdatareplication.ReplicationExtensionModel{
		Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
			CustomProperties: &armrecoveryservicesdatareplication.ReplicationExtensionModelCustomProperties{
				InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
			},
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
	// res.ReplicationExtensionModel = armrecoveryservicesdatareplication.ReplicationExtensionModel{
	// 	Name: to.Ptr("xvjffbiecsd"),
	// 	Type: to.Ptr("miadbgilpheilzfoonveznybthgdwh"),
	// 	ID: to.Ptr("awu"),
	// 	Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
	// 		CustomProperties: &armrecoveryservicesdatareplication.ReplicationExtensionModelCustomProperties{
	// 			InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
	// 		},
	// 		ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armrecoveryservicesdatareplication.ReplicationExtensionModelSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 		CreatedBy: to.Ptr("lagtinqhksctfdxmfbpf"),
	// 		CreatedByType: to.Ptr("dsqllpglanwztdmisrknjtqz"),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("suwjpejlaya"),
	// 		LastModifiedByType: to.Ptr("nrfjuyghtbivwihr"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_Delete.json
func ExampleReplicationExtensionClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationExtensionClient().BeginDelete(ctx, "rgrecoveryservicesdatareplication", "4", "g16yjJ", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ReplicationExtension_List.json
func ExampleReplicationExtensionClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationExtensionClient().NewListPager("rgrecoveryservicesdatareplication", "4", nil)
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
		// page.ReplicationExtensionModelCollection = armrecoveryservicesdatareplication.ReplicationExtensionModelCollection{
		// 	Value: []*armrecoveryservicesdatareplication.ReplicationExtensionModel{
		// 		{
		// 			Name: to.Ptr("xvjffbiecsd"),
		// 			Type: to.Ptr("miadbgilpheilzfoonveznybthgdwh"),
		// 			ID: to.Ptr("awu"),
		// 			Properties: &armrecoveryservicesdatareplication.ReplicationExtensionModelProperties{
		// 				CustomProperties: &armrecoveryservicesdatareplication.ReplicationExtensionModelCustomProperties{
		// 					InstanceType: to.Ptr("ReplicationExtensionModelCustomProperties"),
		// 				},
		// 				ProvisioningState: to.Ptr(armrecoveryservicesdatareplication.ProvisioningStateSucceeded),
		// 			},
		// 			SystemData: &armrecoveryservicesdatareplication.ReplicationExtensionModelSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
		// 				CreatedBy: to.Ptr("lagtinqhksctfdxmfbpf"),
		// 				CreatedByType: to.Ptr("dsqllpglanwztdmisrknjtqz"),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T00:28:56.732Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("suwjpejlaya"),
		// 				LastModifiedByType: to.Ptr("nrfjuyghtbivwihr"),
		// 			},
		// 	}},
		// }
	}
}
