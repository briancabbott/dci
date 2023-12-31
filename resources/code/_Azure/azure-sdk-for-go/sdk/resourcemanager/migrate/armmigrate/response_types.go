//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

// AssessedMachinesClientGetResponse contains the response from method AssessedMachinesClient.Get.
type AssessedMachinesClientGetResponse struct {
	// A machine evaluated as part of an assessment.
	AssessedMachine

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessedMachinesClientListByAssessmentResponse contains the response from method AssessedMachinesClient.NewListByAssessmentPager.
type AssessedMachinesClientListByAssessmentResponse struct {
	// List of assessed machines.
	AssessedMachineResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientCreateResponse contains the response from method AssessmentsClient.Create.
type AssessmentsClientCreateResponse struct {
	// An assessment created for a group in the Migration project.
	Assessment

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientDeleteResponse contains the response from method AssessmentsClient.Delete.
type AssessmentsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientGetReportDownloadURLResponse contains the response from method AssessmentsClient.GetReportDownloadURL.
type AssessmentsClientGetReportDownloadURLResponse struct {
	// Download URL for assessment report.
	DownloadURL

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientGetResponse contains the response from method AssessmentsClient.Get.
type AssessmentsClientGetResponse struct {
	// An assessment created for a group in the Migration project.
	Assessment

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientListByGroupResponse contains the response from method AssessmentsClient.NewListByGroupPager.
type AssessmentsClientListByGroupResponse struct {
	// List of assessments.
	AssessmentResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// AssessmentsClientListByProjectResponse contains the response from method AssessmentsClient.NewListByProjectPager.
type AssessmentsClientListByProjectResponse struct {
	// List of assessments.
	AssessmentResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GroupsClientCreateResponse contains the response from method GroupsClient.Create.
type GroupsClientCreateResponse struct {
	// A group created in a Migration project.
	Group

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GroupsClientDeleteResponse contains the response from method GroupsClient.Delete.
type GroupsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GroupsClientGetResponse contains the response from method GroupsClient.Get.
type GroupsClientGetResponse struct {
	// A group created in a Migration project.
	Group

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GroupsClientListByProjectResponse contains the response from method GroupsClient.NewListByProjectPager.
type GroupsClientListByProjectResponse struct {
	// List of groups.
	GroupResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GroupsClientUpdateMachinesResponse contains the response from method GroupsClient.UpdateMachines.
type GroupsClientUpdateMachinesResponse struct {
	// A group created in a Migration project.
	Group

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// HyperVCollectorsClientCreateResponse contains the response from method HyperVCollectorsClient.Create.
type HyperVCollectorsClientCreateResponse struct {
	HyperVCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// HyperVCollectorsClientDeleteResponse contains the response from method HyperVCollectorsClient.Delete.
type HyperVCollectorsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// HyperVCollectorsClientGetResponse contains the response from method HyperVCollectorsClient.Get.
type HyperVCollectorsClientGetResponse struct {
	HyperVCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// HyperVCollectorsClientListByProjectResponse contains the response from method HyperVCollectorsClient.NewListByProjectPager.
type HyperVCollectorsClientListByProjectResponse struct {
	// List of Hyper-V collectors.
	HyperVCollectorList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ImportCollectorsClientCreateResponse contains the response from method ImportCollectorsClient.Create.
type ImportCollectorsClientCreateResponse struct {
	ImportCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ImportCollectorsClientDeleteResponse contains the response from method ImportCollectorsClient.Delete.
type ImportCollectorsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ImportCollectorsClientGetResponse contains the response from method ImportCollectorsClient.Get.
type ImportCollectorsClientGetResponse struct {
	ImportCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ImportCollectorsClientListByProjectResponse contains the response from method ImportCollectorsClient.NewListByProjectPager.
type ImportCollectorsClientListByProjectResponse struct {
	// List of Import collectors.
	ImportCollectorList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// MachinesClientGetResponse contains the response from method MachinesClient.Get.
type MachinesClientGetResponse struct {
	// A machine in a migration project.
	Machine

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// MachinesClientListByProjectResponse contains the response from method MachinesClient.NewListByProjectPager.
type MachinesClientListByProjectResponse struct {
	// List of machines.
	MachineResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// List of API operations.
	OperationResultList
}

// PrivateEndpointConnectionClientDeleteResponse contains the response from method PrivateEndpointConnectionClient.Delete.
type PrivateEndpointConnectionClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateEndpointConnectionClientGetResponse contains the response from method PrivateEndpointConnectionClient.Get.
type PrivateEndpointConnectionClientGetResponse struct {
	// A private endpoint connection for a project.
	PrivateEndpointConnection

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateEndpointConnectionClientListByProjectResponse contains the response from method PrivateEndpointConnectionClient.ListByProject.
type PrivateEndpointConnectionClientListByProjectResponse struct {
	// A collection of private endpoint connections for a project.
	PrivateEndpointConnectionCollection

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateEndpointConnectionClientUpdateResponse contains the response from method PrivateEndpointConnectionClient.Update.
type PrivateEndpointConnectionClientUpdateResponse struct {
	// A private endpoint connection for a project.
	PrivateEndpointConnection

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateLinkResourceClientGetResponse contains the response from method PrivateLinkResourceClient.Get.
type PrivateLinkResourceClientGetResponse struct {
	// A private link resource for a project for which a private endpoint can be created.
	PrivateLinkResource

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// PrivateLinkResourceClientListByProjectResponse contains the response from method PrivateLinkResourceClient.ListByProject.
type PrivateLinkResourceClientListByProjectResponse struct {
	// A list of private link resources
	PrivateLinkResourceCollection

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientAssessmentOptionsListResponse contains the response from method ProjectsClient.NewAssessmentOptionsListPager.
type ProjectsClientAssessmentOptionsListResponse struct {
	// List of API operations.
	AssessmentOptionsResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientAssessmentOptionsResponse contains the response from method ProjectsClient.AssessmentOptions.
type ProjectsClientAssessmentOptionsResponse struct {
	// Assessment options.
	AssessmentOptions

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientCreateResponse contains the response from method ProjectsClient.Create.
type ProjectsClientCreateResponse struct {
	// Azure Migrate Project.
	Project

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.Delete.
type ProjectsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	// Azure Migrate Project.
	Project

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientListBySubscriptionResponse contains the response from method ProjectsClient.NewListBySubscriptionPager.
type ProjectsClientListBySubscriptionResponse struct {
	// List of projects.
	ProjectResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientListResponse contains the response from method ProjectsClient.NewListPager.
type ProjectsClientListResponse struct {
	// List of projects.
	ProjectResultList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.Update.
type ProjectsClientUpdateResponse struct {
	// Azure Migrate Project.
	Project

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerCollectorsClientCreateResponse contains the response from method ServerCollectorsClient.Create.
type ServerCollectorsClientCreateResponse struct {
	ServerCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerCollectorsClientDeleteResponse contains the response from method ServerCollectorsClient.Delete.
type ServerCollectorsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerCollectorsClientGetResponse contains the response from method ServerCollectorsClient.Get.
type ServerCollectorsClientGetResponse struct {
	ServerCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// ServerCollectorsClientListByProjectResponse contains the response from method ServerCollectorsClient.NewListByProjectPager.
type ServerCollectorsClientListByProjectResponse struct {
	// List of Server collectors.
	ServerCollectorList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// VMwareCollectorsClientCreateResponse contains the response from method VMwareCollectorsClient.Create.
type VMwareCollectorsClientCreateResponse struct {
	VMwareCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// VMwareCollectorsClientDeleteResponse contains the response from method VMwareCollectorsClient.Delete.
type VMwareCollectorsClientDeleteResponse struct {
	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// VMwareCollectorsClientGetResponse contains the response from method VMwareCollectorsClient.Get.
type VMwareCollectorsClientGetResponse struct {
	VMwareCollector

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// VMwareCollectorsClientListByProjectResponse contains the response from method VMwareCollectorsClient.NewListByProjectPager.
type VMwareCollectorsClientListByProjectResponse struct {
	// List of VMware collectors.
	VMwareCollectorList

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}
