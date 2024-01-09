//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_ListByVolumeGroup_MaximumSet_Gen.json
func ExampleVolumeSnapshotsClient_NewListByVolumeGroupPager_volumeSnapshotsListByVolumeGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeSnapshotsClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", &armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions{Filter: to.Ptr("volumeName eq <volume name>")})
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
		// page.SnapshotList = armelasticsan.SnapshotList{
		// 	Value: []*armelasticsan.Snapshot{
		// 		{
		// 			Name: to.Ptr("qukfugetqthsufp"),
		// 			Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
		// 			ID: to.Ptr("bbqqgzxagggqgkdgjqq"),
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
		// 				CreatedBy: to.Ptr("f"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 			Properties: &armelasticsan.SnapshotProperties{
		// 				CreationData: &armelasticsan.SnapshotCreationData{
		// 					SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
		// 				},
		// 				ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 				SourceVolumeSizeGiB: to.Ptr[int64](28),
		// 				VolumeName: to.Ptr("volumename"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_ListByVolumeGroup_MinimumSet_Gen.json
func ExampleVolumeSnapshotsClient_NewListByVolumeGroupPager_volumeSnapshotsListByVolumeGroupMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeSnapshotsClient().NewListByVolumeGroupPager("resourcegroupname", "elasticsanname", "volumegroupname", &armelasticsan.VolumeSnapshotsClientListByVolumeGroupOptions{Filter: nil})
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
		// page.SnapshotList = armelasticsan.SnapshotList{
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Create_MaximumSet_Gen.json
func ExampleVolumeSnapshotsClient_BeginCreate_volumeSnapshotsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeSnapshotsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", armelasticsan.Snapshot{
		Properties: &armelasticsan.SnapshotProperties{
			CreationData: &armelasticsan.SnapshotCreationData{
				SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
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
	// res.Snapshot = armelasticsan.Snapshot{
	// 	Name: to.Ptr("qukfugetqthsufp"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/snapshots/{snapshotName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		CreatedBy: to.Ptr("f"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.SnapshotProperties{
	// 		CreationData: &armelasticsan.SnapshotCreationData{
	// 			SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SourceVolumeSizeGiB: to.Ptr[int64](28),
	// 		VolumeName: to.Ptr("volumename"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Create_MinimumSet_Gen.json
func ExampleVolumeSnapshotsClient_BeginCreate_volumeSnapshotsCreateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeSnapshotsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", armelasticsan.Snapshot{
		Properties: &armelasticsan.SnapshotProperties{
			CreationData: &armelasticsan.SnapshotCreationData{
				SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
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
	// res.Snapshot = armelasticsan.Snapshot{
	// 	Name: to.Ptr("qukfugetqthsufp"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/snapshots/{snapshotName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		CreatedBy: to.Ptr("f"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.SnapshotProperties{
	// 		CreationData: &armelasticsan.SnapshotCreationData{
	// 			SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SourceVolumeSizeGiB: to.Ptr[int64](28),
	// 		VolumeName: to.Ptr("volumename"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Delete_MaximumSet_Gen.json
func ExampleVolumeSnapshotsClient_BeginDelete_volumeSnapshotsDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeSnapshotsClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Delete_MinimumSet_Gen.json
func ExampleVolumeSnapshotsClient_BeginDelete_volumeSnapshotsDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVolumeSnapshotsClient().BeginDelete(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Get_MaximumSet_Gen.json
func ExampleVolumeSnapshotsClient_Get_volumeSnapshotsGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeSnapshotsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Snapshot = armelasticsan.Snapshot{
	// 	Name: to.Ptr("qukfugetqthsufp"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
	// 	ID: to.Ptr("bbqqgzxagggqgkdgjqq"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		CreatedBy: to.Ptr("f"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.SnapshotProperties{
	// 		CreationData: &armelasticsan.SnapshotCreationData{
	// 			SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SourceVolumeSizeGiB: to.Ptr[int64](28),
	// 		VolumeName: to.Ptr("volumename"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/VolumeSnapshots_Get_MinimumSet_Gen.json
func ExampleVolumeSnapshotsClient_Get_volumeSnapshotsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVolumeSnapshotsClient().Get(ctx, "resourcegroupname", "elasticsanname", "volumegroupname", "snapshotname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Snapshot = armelasticsan.Snapshot{
	// 	Name: to.Ptr("qukfugetqthsufp"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/volumeGroups/snapshots"),
	// 	ID: to.Ptr("bbqqgzxagggqgkdgjqq"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		CreatedBy: to.Ptr("f"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-06T06:58:45.864Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("wifrlzeszzcckvdzdwxhvservlvarp"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.SnapshotProperties{
	// 		CreationData: &armelasticsan.SnapshotCreationData{
	// 			SourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}/volumes/{volumeName}"),
	// 		},
	// 		ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		SourceVolumeSizeGiB: to.Ptr[int64](28),
	// 		VolumeName: to.Ptr("volumename"),
	// 	},
	// }
}
