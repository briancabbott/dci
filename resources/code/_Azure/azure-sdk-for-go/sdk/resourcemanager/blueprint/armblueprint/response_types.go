//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

// ArtifactsClientCreateOrUpdateResponse contains the response from method ArtifactsClient.CreateOrUpdate.
type ArtifactsClientCreateOrUpdateResponse struct {
	// Represents a blueprint artifact.
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsClientCreateOrUpdateResponse.
func (a *ArtifactsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsClientDeleteResponse contains the response from method ArtifactsClient.Delete.
type ArtifactsClientDeleteResponse struct {
	// Represents a blueprint artifact.
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsClientDeleteResponse.
func (a *ArtifactsClientDeleteResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsClientGetResponse contains the response from method ArtifactsClient.Get.
type ArtifactsClientGetResponse struct {
	// Represents a blueprint artifact.
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsClientGetResponse.
func (a *ArtifactsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsClientListResponse contains the response from method ArtifactsClient.NewListPager.
type ArtifactsClientListResponse struct {
	// List of blueprint artifacts.
	ArtifactList
}

// AssignmentOperationsClientGetResponse contains the response from method AssignmentOperationsClient.Get.
type AssignmentOperationsClientGetResponse struct {
	// Represents underlying deployment detail for each update to the blueprint assignment.
	AssignmentOperation
}

// AssignmentOperationsClientListResponse contains the response from method AssignmentOperationsClient.NewListPager.
type AssignmentOperationsClientListResponse struct {
	// List of AssignmentOperation.
	AssignmentOperationList
}

// AssignmentsClientCreateOrUpdateResponse contains the response from method AssignmentsClient.CreateOrUpdate.
type AssignmentsClientCreateOrUpdateResponse struct {
	// Represents a blueprint assignment.
	Assignment
}

// AssignmentsClientDeleteResponse contains the response from method AssignmentsClient.Delete.
type AssignmentsClientDeleteResponse struct {
	// Represents a blueprint assignment.
	Assignment
}

// AssignmentsClientGetResponse contains the response from method AssignmentsClient.Get.
type AssignmentsClientGetResponse struct {
	// Represents a blueprint assignment.
	Assignment
}

// AssignmentsClientListResponse contains the response from method AssignmentsClient.NewListPager.
type AssignmentsClientListResponse struct {
	// List of blueprint assignments
	AssignmentList
}

// AssignmentsClientWhoIsBlueprintResponse contains the response from method AssignmentsClient.WhoIsBlueprint.
type AssignmentsClientWhoIsBlueprintResponse struct {
	// Response schema for querying the Azure Blueprints service principal in the tenant.
	WhoIsBlueprintContract
}

// BlueprintsClientCreateOrUpdateResponse contains the response from method BlueprintsClient.CreateOrUpdate.
type BlueprintsClientCreateOrUpdateResponse struct {
	// Represents a Blueprint definition.
	Blueprint
}

// BlueprintsClientDeleteResponse contains the response from method BlueprintsClient.Delete.
type BlueprintsClientDeleteResponse struct {
	// Represents a Blueprint definition.
	Blueprint
}

// BlueprintsClientGetResponse contains the response from method BlueprintsClient.Get.
type BlueprintsClientGetResponse struct {
	// Represents a Blueprint definition.
	Blueprint
}

// BlueprintsClientListResponse contains the response from method BlueprintsClient.NewListPager.
type BlueprintsClientListResponse struct {
	// List of blueprint definitions.
	List
}

// PublishedArtifactsClientGetResponse contains the response from method PublishedArtifactsClient.Get.
type PublishedArtifactsClientGetResponse struct {
	// Represents a blueprint artifact.
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublishedArtifactsClientGetResponse.
func (p *PublishedArtifactsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	p.ArtifactClassification = res
	return nil
}

// PublishedArtifactsClientListResponse contains the response from method PublishedArtifactsClient.NewListPager.
type PublishedArtifactsClientListResponse struct {
	// List of blueprint artifacts.
	ArtifactList
}

// PublishedBlueprintsClientCreateResponse contains the response from method PublishedBlueprintsClient.Create.
type PublishedBlueprintsClientCreateResponse struct {
	// Represents a published blueprint.
	PublishedBlueprint
}

// PublishedBlueprintsClientDeleteResponse contains the response from method PublishedBlueprintsClient.Delete.
type PublishedBlueprintsClientDeleteResponse struct {
	// Represents a published blueprint.
	PublishedBlueprint
}

// PublishedBlueprintsClientGetResponse contains the response from method PublishedBlueprintsClient.Get.
type PublishedBlueprintsClientGetResponse struct {
	// Represents a published blueprint.
	PublishedBlueprint
}

// PublishedBlueprintsClientListResponse contains the response from method PublishedBlueprintsClient.NewListPager.
type PublishedBlueprintsClientListResponse struct {
	// List of published blueprint definitions.
	PublishedBlueprintList
}
