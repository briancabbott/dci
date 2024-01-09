// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package main

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"log"
	"os"
)

var (
	subscriptionID    string
	location          = "westus"
	resourceGroupName = "sample-resource-group"
	serviceName       = "sample-api-service"
	versionSetID      = "sample-api-version"
)

var (
	apimanagementClientFactory *armapimanagement.ClientFactory
	resourcesClientFactory     *armresources.ClientFactory
)

var (
	resourceGroupClient *armresources.ResourceGroupsClient
	serviceClient       *armapimanagement.ServiceClient
	apiVersionSetClient *armapimanagement.APIVersionSetClient
)

func main() {
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID is not set.")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	resourcesClientFactory, err = armresources.NewClientFactory(subscriptionID, cred, nil)
	if err != nil {
		log.Fatal(err)
	}
	resourceGroupClient = resourcesClientFactory.NewResourceGroupsClient()

	apimanagementClientFactory, err = armapimanagement.NewClientFactory(subscriptionID, cred, nil)
	if err != nil {
		log.Fatal(err)
	}
	serviceClient = apimanagementClientFactory.NewServiceClient()
	apiVersionSetClient = apimanagementClientFactory.NewAPIVersionSetClient()

	resourceGroup, err := createResourceGroup(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resources group:", *resourceGroup.ID)

	apiManagementService, err := createApiManagementService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("api management service:", *apiManagementService.ID)

	apiVersionSet, err := createApiVersionSet(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("api version set:", *apiVersionSet.ID)

	keepResource := os.Getenv("KEEP_RESOURCE")
	if len(keepResource) == 0 {
		err = cleanup(ctx)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("cleaned up successfully.")
	}
}

func createApiManagementService(ctx context.Context) (*armapimanagement.ServiceResource, error) {

	pollerResp, err := serviceClient.BeginCreateOrUpdate(
		ctx,
		resourceGroupName,
		serviceName,
		armapimanagement.ServiceResource{
			Location: to.Ptr(location),
			Properties: &armapimanagement.ServiceProperties{
				PublisherName:  to.Ptr("sample"),
				PublisherEmail: to.Ptr("xxx@wircesoft.com"),
			},
			SKU: &armapimanagement.ServiceSKUProperties{
				Name:     to.Ptr(armapimanagement.SKUTypeStandard),
				Capacity: to.Ptr[int32](2),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	resp, err := pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &resp.ServiceResource, nil
}

func createApiVersionSet(ctx context.Context) (*armapimanagement.APIVersionSetContract, error) {

	resp, err := apiVersionSetClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serviceName,
		versionSetID,
		armapimanagement.APIVersionSetContract{
			Properties: &armapimanagement.APIVersionSetContractProperties{
				DisplayName:      to.Ptr("sample api version set description"),
				VersioningScheme: to.Ptr(armapimanagement.VersioningSchemeQuery),
				VersionQueryName: to.Ptr("sample-version-query"),
			},
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &resp.APIVersionSetContract, nil
}

func createResourceGroup(ctx context.Context) (*armresources.ResourceGroup, error) {

	resourceGroupResp, err := resourceGroupClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		armresources.ResourceGroup{
			Location: to.Ptr(location),
		},
		nil)
	if err != nil {
		return nil, err
	}
	return &resourceGroupResp.ResourceGroup, nil
}

func cleanup(ctx context.Context) error {

	pollerResp, err := resourceGroupClient.BeginDelete(ctx, resourceGroupName, nil)
	if err != nil {
		return err
	}

	_, err = pollerResp.PollUntilDone(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}
