//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationCheckNameAvailability.json
func ExampleAttachedDatabaseConfigurationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttachedDatabaseConfigurationsClient().CheckNameAvailability(ctx, "kustorptest", "kustoCluster", armkusto.AttachedDatabaseConfigurationsCheckNameRequest{
		Name: to.Ptr("adc1"),
		Type: to.Ptr("Microsoft.Kusto/clusters/attachedDatabaseConfigurations"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armkusto.CheckNameResult{
	// 	Name: to.Ptr("adc1"),
	// 	Message: to.Ptr("Name 'adc1' is already taken. Please specify a different name"),
	// 	NameAvailable: to.Ptr(false),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsListByCluster.json
func ExampleAttachedDatabaseConfigurationsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAttachedDatabaseConfigurationsClient().NewListByClusterPager("kustorptest", "kustoCluster2", nil)
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
		// page.AttachedDatabaseConfigurationListResult = armkusto.AttachedDatabaseConfigurationListResult{
		// 	Value: []*armkusto.AttachedDatabaseConfiguration{
		// 		{
		// 			Name: to.Ptr("kustoCluster2/KustoDatabase8"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/AttachedDatabaseConfigurations/KustoDatabase8"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster2/KustoDatabase9"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/AttachedDatabaseConfigurations/KustoDatabase9"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterLeader"),
		// 				DatabaseName: to.Ptr("db1"),
		// 				DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
		// 				ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
		// 				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
		// 					ExternalTablesToExclude: []*string{
		// 						to.Ptr("ExternalTable2")},
		// 						ExternalTablesToInclude: []*string{
		// 							to.Ptr("ExternalTable1")},
		// 							FunctionsToExclude: []*string{
		// 								to.Ptr("functionsToExclude2")},
		// 								FunctionsToInclude: []*string{
		// 									to.Ptr("functionsToInclude1")},
		// 									MaterializedViewsToExclude: []*string{
		// 										to.Ptr("MaterializedViewTable2")},
		// 										MaterializedViewsToInclude: []*string{
		// 											to.Ptr("MaterializedViewTable1")},
		// 											TablesToExclude: []*string{
		// 												to.Ptr("Table2")},
		// 												TablesToInclude: []*string{
		// 													to.Ptr("Table1")},
		// 												},
		// 											},
		// 									}},
		// 								}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsGet.json
func ExampleAttachedDatabaseConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAttachedDatabaseConfigurationsClient().Get(ctx, "kustorptest", "kustoCluster2", "attachedDatabaseConfigurationsTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AttachedDatabaseConfiguration = armkusto.AttachedDatabaseConfiguration{
	// 	Name: to.Ptr("kustoCluster2/attachedDatabaseConfigurationsTest"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/AttachedDatabaseConfigurations/attachedDatabaseConfigurationsTest"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkusto.AttachedDatabaseConfigurationProperties{
	// 		ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
	// 		DatabaseName: to.Ptr("*"),
	// 		DatabaseNamePrefix: to.Ptr("prefix"),
	// 		DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAttachedDatabaseConfigurationsClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster2", "attachedDatabaseConfigurationsTest", armkusto.AttachedDatabaseConfiguration{
		Location: to.Ptr("westus"),
		Properties: &armkusto.AttachedDatabaseConfigurationProperties{
			ClusterResourceID:                 to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
			DatabaseName:                      to.Ptr("kustodatabase"),
			DatabaseNameOverride:              to.Ptr("overridekustodatabase"),
			DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
			TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
				ExternalTablesToExclude: []*string{
					to.Ptr("ExternalTable2")},
				ExternalTablesToInclude: []*string{
					to.Ptr("ExternalTable1")},
				MaterializedViewsToExclude: []*string{
					to.Ptr("MaterializedViewTable2")},
				MaterializedViewsToInclude: []*string{
					to.Ptr("MaterializedViewTable1")},
				TablesToExclude: []*string{
					to.Ptr("Table2")},
				TablesToInclude: []*string{
					to.Ptr("Table1")},
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
	// res.AttachedDatabaseConfiguration = armkusto.AttachedDatabaseConfiguration{
	// 	Name: to.Ptr("kustoCluster2/attachedDatabaseConfigurationsTest"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/AttachedDatabaseConfigurations"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2/attachedDatabaseConfigurations/attachedDatabaseConfigurationsTest"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armkusto.AttachedDatabaseConfigurationProperties{
	// 		ClusterResourceID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2"),
	// 		DatabaseName: to.Ptr("db1"),
	// 		DatabaseNameOverride: to.Ptr("overridekustodatabase"),
	// 		DefaultPrincipalsModificationKind: to.Ptr(armkusto.DefaultPrincipalsModificationKindUnion),
	// 		ProvisioningState: to.Ptr(armkusto.ProvisioningStateSucceeded),
	// 		TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
	// 			ExternalTablesToExclude: []*string{
	// 				to.Ptr("ExternalTable2")},
	// 				ExternalTablesToInclude: []*string{
	// 					to.Ptr("ExternalTable1")},
	// 					MaterializedViewsToExclude: []*string{
	// 						to.Ptr("MaterializedViewTable2")},
	// 						MaterializedViewsToInclude: []*string{
	// 							to.Ptr("MaterializedViewTable1")},
	// 							TablesToExclude: []*string{
	// 								to.Ptr("Table2")},
	// 								TablesToInclude: []*string{
	// 									to.Ptr("Table1")},
	// 								},
	// 							},
	// 						}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsDelete.json
func ExampleAttachedDatabaseConfigurationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAttachedDatabaseConfigurationsClient().BeginDelete(ctx, "kustorptest", "kustoCluster", "attachedDatabaseConfigurationsTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
