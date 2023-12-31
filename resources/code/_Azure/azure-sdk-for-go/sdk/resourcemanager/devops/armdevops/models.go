//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevops

// Authorization info used to access a resource (like code repository).
type Authorization struct {
	// REQUIRED; Type of authorization.
	AuthorizationType *AuthorizationType

	// Authorization parameters corresponding to the authorization type.
	Parameters map[string]*string
}

// BootstrapConfiguration - Configuration used to bootstrap a Pipeline.
type BootstrapConfiguration struct {
	// REQUIRED; Template used to bootstrap the pipeline.
	Template *PipelineTemplate

	// Repository containing the source code for the pipeline.
	Repository *CodeRepository
}

// CodeRepository - Repository containing the source code for a pipeline.
type CodeRepository struct {
	// REQUIRED; Default branch used to configure Continuous Integration (CI) in the pipeline.
	DefaultBranch *string

	// REQUIRED; Unique immutable identifier of the code repository.
	ID *string

	// REQUIRED; Type of code repository.
	RepositoryType *CodeRepositoryType

	// Authorization info to access the code repository.
	Authorization *Authorization

	// Repository-specific properties.
	Properties map[string]*string
}

// InputDescriptor - Representation of a pipeline template input parameter.
type InputDescriptor struct {
	// REQUIRED; Identifier of the input parameter.
	ID *string

	// REQUIRED; Data type of the value of the input parameter.
	Type *InputDataType

	// Description of the input parameter.
	Description *string

	// List of possible values for the input parameter.
	PossibleValues []*InputValue
}

// InputValue - Representation of a pipeline template input parameter value.
type InputValue struct {
	// Description of the input parameter value.
	DisplayValue *string

	// Value of an input parameter.
	Value *string
}

// Operation - Properties of an Operation.
type Operation struct {
	// Display information of the operation.
	Display *OperationDisplayValue

	// Indicates whether the operation applies to data-plane.
	IsDataAction *string

	// READ-ONLY; Name of the operation.
	Name *string
}

// OperationDisplayValue - Display information of an operation.
type OperationDisplayValue struct {
	// READ-ONLY; Friendly description of the operation.
	Description *string

	// READ-ONLY; Friendly name of the operation.
	Operation *string

	// READ-ONLY; Friendly name of the resource provider.
	Provider *string

	// READ-ONLY; Friendly name of the resource type the operation applies to.
	Resource *string
}

// OperationListResult - Result of a request to list all operations supported by Microsoft.DevOps resource provider.
type OperationListResult struct {
	// The URL to get the next set of operations, if there are any.
	NextLink *string

	// READ-ONLY; List of operations supported by Microsoft.DevOps resource provider.
	Value []*Operation
}

// OrganizationReference - Reference to an Azure DevOps Organization.
type OrganizationReference struct {
	// REQUIRED; Name of the Azure DevOps Organization.
	Name *string

	// READ-ONLY; Unique immutable identifier for the Azure DevOps Organization.
	ID *string
}

// Pipeline - Azure DevOps Pipeline used to configure Continuous Integration (CI) & Continuous Delivery (CD) for Azure resources.
type Pipeline struct {
	// REQUIRED; Custom properties of the Pipeline.
	Properties *PipelineProperties

	// Resource Location
	Location *string

	// Resource Tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource Name
	Name *string

	// READ-ONLY; Resource Type
	Type *string
}

// PipelineListResult - Result of a request to list all Azure Pipelines under a given scope.
type PipelineListResult struct {
	// URL to get the next set of Pipelines, if there are any.
	NextLink *string

	// List of pipelines.
	Value []*Pipeline
}

// PipelineProperties - Custom properties of a Pipeline.
type PipelineProperties struct {
	// REQUIRED; Configuration used to bootstrap the Pipeline.
	BootstrapConfiguration *BootstrapConfiguration

	// REQUIRED; Reference to the Azure DevOps Organization containing the Pipeline.
	Organization *OrganizationReference

	// REQUIRED; Reference to the Azure DevOps Project containing the Pipeline.
	Project *ProjectReference

	// READ-ONLY; Unique identifier of the Azure Pipeline within the Azure DevOps Project.
	PipelineID *int32
}

// PipelineTemplate - Template used to bootstrap the pipeline.
type PipelineTemplate struct {
	// REQUIRED; Unique identifier of the pipeline template.
	ID *string

	// Dictionary of input parameters used in the pipeline template.
	Parameters map[string]*string
}

// PipelineTemplateDefinition - Definition of a pipeline template.
type PipelineTemplateDefinition struct {
	// REQUIRED; Unique identifier of the pipeline template.
	ID *string

	// Description of the pipeline enabled by the template.
	Description *string

	// List of input parameters required by the template to create a pipeline.
	Inputs []*InputDescriptor
}

// PipelineTemplateDefinitionListResult - Result of a request to list all pipeline template definitions.
type PipelineTemplateDefinitionListResult struct {
	// The URL to get the next set of pipeline template definitions, if there are any.
	NextLink *string

	// List of pipeline template definitions.
	Value []*PipelineTemplateDefinition
}

// PipelineUpdateParameters - Request payload used to update an existing Azure Pipeline.
type PipelineUpdateParameters struct {
	// Dictionary of key-value pairs to be set as tags on the Azure Pipeline. This will overwrite any existing tags.
	Tags map[string]*string
}

// ProjectReference - Reference to an Azure DevOps Project.
type ProjectReference struct {
	// REQUIRED; Name of the Azure DevOps Project.
	Name *string

	// READ-ONLY; Unique immutable identifier of the Azure DevOps Project.
	ID *string
}

// Resource - An Azure Resource Manager (ARM) resource.
type Resource struct {
	// Resource Location
	Location *string

	// Resource Tags
	Tags map[string]*string

	// READ-ONLY; Resource Id
	ID *string

	// READ-ONLY; Resource Name
	Name *string

	// READ-ONLY; Resource Type
	Type *string
}
