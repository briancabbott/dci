//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationGet_example.json
func ExampleCustomAssessmentAutomationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomAssessmentAutomationsClient().Get(ctx, "TestResourceGroup", "MyCustomAssessmentAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomAssessmentAutomation = armsecurity.CustomAssessmentAutomation{
	// 	Name: to.Ptr("MyCustomAssessmentAutomation"),
	// 	Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
	// 	ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation"),
	// 	Properties: &armsecurity.CustomAssessmentAutomationProperties{
	// 		Description: to.Ptr("organization passwords policy"),
	// 		AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
	// 		CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
	// 		DisplayName: to.Ptr("Password Policy"),
	// 		RemediationDescription: to.Ptr("Change password policy to..."),
	// 		Severity: to.Ptr(armsecurity.SeverityEnumLow),
	// 		SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationCreate_example.json
func ExampleCustomAssessmentAutomationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomAssessmentAutomationsClient().Create(ctx, "TestResourceGroup", "MyCustomAssessmentAutomation", armsecurity.CustomAssessmentAutomationRequest{
		Properties: &armsecurity.CustomAssessmentAutomationRequestProperties{
			Description:            to.Ptr("Data should be encrypted"),
			CompressedQuery:        to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
			DisplayName:            to.Ptr("Password Policy"),
			RemediationDescription: to.Ptr("Encrypt store by..."),
			Severity:               to.Ptr(armsecurity.SeverityEnumMedium),
			SupportedCloud:         to.Ptr(armsecurity.SupportedCloudEnumAWS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomAssessmentAutomation = armsecurity.CustomAssessmentAutomation{
	// 	Name: to.Ptr("33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
	// 	ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/33e7cc6e-a139-4723-a0e5-76993aee0771"),
	// 	Properties: &armsecurity.CustomAssessmentAutomationProperties{
	// 		Description: to.Ptr("organization passwords policy"),
	// 		AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
	// 		CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
	// 		DisplayName: to.Ptr("Password Policy"),
	// 		RemediationDescription: to.Ptr("Change password policy to..."),
	// 		Severity: to.Ptr(armsecurity.SeverityEnumMedium),
	// 		SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationDelete_example.json
func ExampleCustomAssessmentAutomationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCustomAssessmentAutomationsClient().Delete(ctx, "TestResourceGroup", "MyCustomAssessmentAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationListByResourceGroup_example.json
func ExampleCustomAssessmentAutomationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomAssessmentAutomationsClient().NewListByResourceGroupPager("TestResourceGroup", nil)
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
		// page.CustomAssessmentAutomationsListResult = armsecurity.CustomAssessmentAutomationsListResult{
		// 	Value: []*armsecurity.CustomAssessmentAutomation{
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation1"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation1"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
		// 				CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumMedium),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation2"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation2"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("fc1dbcd0-502c-4eab-9312-4014cfc8ea56"),
		// 				CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8AUABhAHMAcwB3AG8AcgBkAFAAbwBsAGkAYwB5ACAADQAKAHwAIABlAHgAdABlAG4AZAAgAEgAZQBhAGwAdABoAFMAdABhAHQAdQBzACAAPQAgAGkAZgBmACgAdABvAGkAbgB0ACgAUgBlAGMAbwByAGQALgBNAGkAbgBpAG0AdQBtAFAAYQBzAHMAdwBvAHIAZABMAGUAbgBnAHQAaAApACAAPAAgADgALAAgACcAVQBOAEgARQBBAEwAVABIAFkAJwAsACAAJwBIAEUAQQBMAFQASABZACcAKQANAAoA"),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumLow),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e716082ac474f182e2220e4f38f1d6191e7636cf/specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationListBySubscription_example.json
func ExampleCustomAssessmentAutomationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomAssessmentAutomationsClient().NewListBySubscriptionPager(nil)
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
		// page.CustomAssessmentAutomationsListResult = armsecurity.CustomAssessmentAutomationsListResult{
		// 	Value: []*armsecurity.CustomAssessmentAutomation{
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation1"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation1"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("d5f442f7-7e77-4bcf-a450-a9c1b9a94eeb"),
		// 				CompressedQuery: to.Ptr("DQAKAEkAYQBtAF8ARwByAG8AdQBwAA0ACgB8ACAAZQB4AHQAZQBuAGQAIABIAGUAYQBsAHQAaABTAHQAYQB0AHUAcwAgAD0AIABpAGYAZgAoAHQAbwBzAHQAcgBpAG4AZwAoAFIAZQBjAG8AcgBkAC4AVQBzAGUAcgBOAGEAbQBlACkAIABjAG8AbgB0AGEAaQBuAHMAIAAnAHUAcwBlAHIAJwAsACAAJwBVAE4ASABFAEEATABUAEgAWQAnACwAIAAnAEgARQBBAEwAVABIAFkAJwApAA0ACgA="),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumMedium),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumAWS),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCustomAssessmentAutomation2"),
		// 			Type: to.Ptr("Microsoft.Security/customAssessmentAutomations"),
		// 			ID: to.Ptr("/subscriptions/e5d1b86c-3051-44d5-8802-aa65d45a279b/resourcegroups/TestResourceGroup/providers/Microsoft.Security/customAssessmentAutomations/MyCustomAssessmentAutomation2"),
		// 			Properties: &armsecurity.CustomAssessmentAutomationProperties{
		// 				Description: to.Ptr("organization passwords policy"),
		// 				AssessmentKey: to.Ptr("fc1dbcd0-502c-4eab-9312-4014cfc8ea56"),
		// 				CompressedQuery: to.Ptr("Q29tcHV0ZV9OZXR3b3JrCnwgZXh0ZW5kIEhlYWx0aFN0YXR1cyA9ICdVTkhFQUxUSFkn"),
		// 				DisplayName: to.Ptr("Password Policy"),
		// 				RemediationDescription: to.Ptr("Change password policy to..."),
		// 				Severity: to.Ptr(armsecurity.SeverityEnumLow),
		// 				SupportedCloud: to.Ptr(armsecurity.SupportedCloudEnumGCP),
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:01:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
