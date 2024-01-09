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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsListByDatabase.json
func ExampleScriptsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScriptsClient().NewListByDatabasePager("kustorptest", "kustoCluster", "Kustodatabase8", nil)
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
		// page.ScriptListResult = armkusto.ScriptListResult{
		// 	Value: []*armkusto.Script{
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript1"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript1"),
		// 			Properties: &armkusto.ScriptProperties{
		// 				ContinueOnErrors: to.Ptr(true),
		// 				ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
		// 				ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
		// 			},
		// 			SystemData: &armkusto.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript2"),
		// 			Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript2"),
		// 			Properties: &armkusto.ScriptProperties{
		// 				ContinueOnErrors: to.Ptr(true),
		// 				ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
		// 				ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script2.txt"),
		// 			},
		// 			SystemData: &armkusto.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-30T16:07:55.283Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsGet.json
func ExampleScriptsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptsClient().Get(ctx, "kustorptest", "kustoCluster", "Kustodatabase8", "kustoScript", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Script = armkusto.Script{
	// 	Name: to.Ptr("kustoCluster/Kustodatabase8/kustoScript"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/Kustodatabase8/Scripts/kustoScript"),
	// 	Properties: &armkusto.ScriptProperties{
	// 		ContinueOnErrors: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
	// 		ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
	// 	},
	// 	SystemData: &armkusto.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsCreateOrUpdate.json
func ExampleScriptsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScriptsClient().BeginCreateOrUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", armkusto.Script{
		Properties: &armkusto.ScriptProperties{
			ContinueOnErrors:  to.Ptr(true),
			ForceUpdateTag:    to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
			ScriptURL:         to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
			ScriptURLSasToken: to.Ptr("?sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=********************************"),
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
	// res.Script = armkusto.Script{
	// 	Name: to.Ptr("kustoCluster/KustoDatabase8/kustoScript"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/Scripts/kustoScript"),
	// 	Properties: &armkusto.ScriptProperties{
	// 		ContinueOnErrors: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
	// 		ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
	// 	},
	// 	SystemData: &armkusto.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T17:16:24.364Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsUpdate.json
func ExampleScriptsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScriptsClient().BeginUpdate(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", armkusto.Script{
		Properties: &armkusto.ScriptProperties{
			ContinueOnErrors:  to.Ptr(true),
			ForceUpdateTag:    to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
			ScriptURL:         to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
			ScriptURLSasToken: to.Ptr("?sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=********************************"),
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
	// res.Script = armkusto.Script{
	// 	Name: to.Ptr("kustoCluster/KustoDatabase8/kustoScript"),
	// 	Type: to.Ptr("Microsoft.Kusto/Clusters/Databases/Scripts"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster/Databases/KustoDatabase8/Scripts/kustoScript"),
	// 	Properties: &armkusto.ScriptProperties{
	// 		ContinueOnErrors: to.Ptr(true),
	// 		ForceUpdateTag: to.Ptr("2bcf3c21-ffd1-4444-b9dd-e52e00ee53fe"),
	// 		ScriptURL: to.Ptr("https://mysa.blob.core.windows.net/container/script.txt"),
	// 	},
	// 	SystemData: &armkusto.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-29T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T17:16:24.364Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armkusto.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsDelete.json
func ExampleScriptsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewScriptsClient().BeginDelete(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", "kustoScript", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoScriptsCheckNameAvailability.json
func ExampleScriptsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptsClient().CheckNameAvailability(ctx, "kustorptest", "kustoCluster", "db", armkusto.ScriptCheckNameRequest{
		Name: to.Ptr("kustoScriptName1"),
		Type: to.Ptr("Microsoft.Kusto/clusters/databases/scripts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armkusto.CheckNameResult{
	// 	Name: to.Ptr("kustoScriptName1"),
	// 	Message: to.Ptr("Name 'kustoScriptName1' is already taken. Please specify a different name"),
	// 	NameAvailable: to.Ptr(false),
	// }
}
