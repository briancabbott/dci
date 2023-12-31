//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

// DedicatedHsmClientBeginCreateOrUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginCreateOrUpdate
// method.
type DedicatedHsmClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientBeginDeleteOptions contains the optional parameters for the DedicatedHsmClient.BeginDelete method.
type DedicatedHsmClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientBeginUpdateOptions contains the optional parameters for the DedicatedHsmClient.BeginUpdate method.
type DedicatedHsmClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DedicatedHsmClientGetOptions contains the optional parameters for the DedicatedHsmClient.Get method.
type DedicatedHsmClientGetOptions struct {
	// placeholder for future optional parameters
}

// DedicatedHsmClientListByResourceGroupOptions contains the optional parameters for the DedicatedHsmClient.NewListByResourceGroupPager
// method.
type DedicatedHsmClientListByResourceGroupOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// DedicatedHsmClientListBySubscriptionOptions contains the optional parameters for the DedicatedHsmClient.NewListBySubscriptionPager
// method.
type DedicatedHsmClientListBySubscriptionOptions struct {
	// Maximum number of results to return.
	Top *int32
}

// DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DedicatedHsmClient.NewListOutboundNetworkDependenciesEndpointsPager
// method.
type DedicatedHsmClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}
