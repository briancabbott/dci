# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

try:
    from ._models_py3 import CheckNameAvailabilityParameters
    from ._models_py3 import DicomService
    from ._models_py3 import DicomServiceAuthenticationConfiguration
    from ._models_py3 import DicomServiceCollection
    from ._models_py3 import DicomServicePatchResource
    from ._models_py3 import Error
    from ._models_py3 import ErrorDetails
    from ._models_py3 import ErrorDetailsInternal
    from ._models_py3 import FhirService
    from ._models_py3 import FhirServiceAccessPolicyEntry
    from ._models_py3 import FhirServiceAcrConfiguration
    from ._models_py3 import FhirServiceAuthenticationConfiguration
    from ._models_py3 import FhirServiceCollection
    from ._models_py3 import FhirServiceCorsConfiguration
    from ._models_py3 import FhirServiceExportConfiguration
    from ._models_py3 import FhirServicePatchResource
    from ._models_py3 import IotConnector
    from ._models_py3 import IotConnectorCollection
    from ._models_py3 import IotConnectorPatchResource
    from ._models_py3 import IotDestinationProperties
    from ._models_py3 import IotEventHubIngestionEndpointConfiguration
    from ._models_py3 import IotFhirDestination
    from ._models_py3 import IotFhirDestinationCollection
    from ._models_py3 import IotFhirDestinationProperties
    from ._models_py3 import IotMappingProperties
    from ._models_py3 import ListOperations
    from ._models_py3 import LocationBasedResource
    from ._models_py3 import LogSpecification
    from ._models_py3 import MetricDimension
    from ._models_py3 import MetricSpecification
    from ._models_py3 import OperationDetail
    from ._models_py3 import OperationDisplay
    from ._models_py3 import OperationProperties
    from ._models_py3 import OperationResultsDescription
    from ._models_py3 import PrivateEndpoint
    from ._models_py3 import PrivateEndpointConnection
    from ._models_py3 import PrivateEndpointConnectionDescription
    from ._models_py3 import PrivateEndpointConnectionListResult
    from ._models_py3 import PrivateEndpointConnectionListResultDescription
    from ._models_py3 import PrivateLinkResource
    from ._models_py3 import PrivateLinkResourceDescription
    from ._models_py3 import PrivateLinkResourceListResultDescription
    from ._models_py3 import PrivateLinkServiceConnectionState
    from ._models_py3 import Resource
    from ._models_py3 import ResourceCore
    from ._models_py3 import ResourceTags
    from ._models_py3 import ResourceVersionPolicyConfiguration
    from ._models_py3 import ServiceAccessPolicyEntry
    from ._models_py3 import ServiceAcrConfigurationInfo
    from ._models_py3 import ServiceAuthenticationConfigurationInfo
    from ._models_py3 import ServiceCorsConfigurationInfo
    from ._models_py3 import ServiceCosmosDbConfigurationInfo
    from ._models_py3 import ServiceExportConfigurationInfo
    from ._models_py3 import ServiceManagedIdentity
    from ._models_py3 import ServiceManagedIdentityautogenerated
    from ._models_py3 import ServiceOciArtifactEntry
    from ._models_py3 import ServiceSpecification
    from ._models_py3 import ServicesDescription
    from ._models_py3 import ServicesDescriptionListResult
    from ._models_py3 import ServicesNameAvailabilityInfo
    from ._models_py3 import ServicesPatchDescription
    from ._models_py3 import ServicesProperties
    from ._models_py3 import ServicesResource
    from ._models_py3 import ServicesResourceIdentity
    from ._models_py3 import SystemData
    from ._models_py3 import TaggedResource
    from ._models_py3 import UserAssignedIdentity
    from ._models_py3 import Workspace
    from ._models_py3 import WorkspaceList
    from ._models_py3 import WorkspacePatchResource
    from ._models_py3 import WorkspaceProperties
except (SyntaxError, ImportError):
    from ._models import CheckNameAvailabilityParameters  # type: ignore
    from ._models import DicomService  # type: ignore
    from ._models import DicomServiceAuthenticationConfiguration  # type: ignore
    from ._models import DicomServiceCollection  # type: ignore
    from ._models import DicomServicePatchResource  # type: ignore
    from ._models import Error  # type: ignore
    from ._models import ErrorDetails  # type: ignore
    from ._models import ErrorDetailsInternal  # type: ignore
    from ._models import FhirService  # type: ignore
    from ._models import FhirServiceAccessPolicyEntry  # type: ignore
    from ._models import FhirServiceAcrConfiguration  # type: ignore
    from ._models import FhirServiceAuthenticationConfiguration  # type: ignore
    from ._models import FhirServiceCollection  # type: ignore
    from ._models import FhirServiceCorsConfiguration  # type: ignore
    from ._models import FhirServiceExportConfiguration  # type: ignore
    from ._models import FhirServicePatchResource  # type: ignore
    from ._models import IotConnector  # type: ignore
    from ._models import IotConnectorCollection  # type: ignore
    from ._models import IotConnectorPatchResource  # type: ignore
    from ._models import IotDestinationProperties  # type: ignore
    from ._models import IotEventHubIngestionEndpointConfiguration  # type: ignore
    from ._models import IotFhirDestination  # type: ignore
    from ._models import IotFhirDestinationCollection  # type: ignore
    from ._models import IotFhirDestinationProperties  # type: ignore
    from ._models import IotMappingProperties  # type: ignore
    from ._models import ListOperations  # type: ignore
    from ._models import LocationBasedResource  # type: ignore
    from ._models import LogSpecification  # type: ignore
    from ._models import MetricDimension  # type: ignore
    from ._models import MetricSpecification  # type: ignore
    from ._models import OperationDetail  # type: ignore
    from ._models import OperationDisplay  # type: ignore
    from ._models import OperationProperties  # type: ignore
    from ._models import OperationResultsDescription  # type: ignore
    from ._models import PrivateEndpoint  # type: ignore
    from ._models import PrivateEndpointConnection  # type: ignore
    from ._models import PrivateEndpointConnectionDescription  # type: ignore
    from ._models import PrivateEndpointConnectionListResult  # type: ignore
    from ._models import PrivateEndpointConnectionListResultDescription  # type: ignore
    from ._models import PrivateLinkResource  # type: ignore
    from ._models import PrivateLinkResourceDescription  # type: ignore
    from ._models import PrivateLinkResourceListResultDescription  # type: ignore
    from ._models import PrivateLinkServiceConnectionState  # type: ignore
    from ._models import Resource  # type: ignore
    from ._models import ResourceCore  # type: ignore
    from ._models import ResourceTags  # type: ignore
    from ._models import ResourceVersionPolicyConfiguration  # type: ignore
    from ._models import ServiceAccessPolicyEntry  # type: ignore
    from ._models import ServiceAcrConfigurationInfo  # type: ignore
    from ._models import ServiceAuthenticationConfigurationInfo  # type: ignore
    from ._models import ServiceCorsConfigurationInfo  # type: ignore
    from ._models import ServiceCosmosDbConfigurationInfo  # type: ignore
    from ._models import ServiceExportConfigurationInfo  # type: ignore
    from ._models import ServiceManagedIdentity  # type: ignore
    from ._models import ServiceManagedIdentityautogenerated  # type: ignore
    from ._models import ServiceOciArtifactEntry  # type: ignore
    from ._models import ServiceSpecification  # type: ignore
    from ._models import ServicesDescription  # type: ignore
    from ._models import ServicesDescriptionListResult  # type: ignore
    from ._models import ServicesNameAvailabilityInfo  # type: ignore
    from ._models import ServicesPatchDescription  # type: ignore
    from ._models import ServicesProperties  # type: ignore
    from ._models import ServicesResource  # type: ignore
    from ._models import ServicesResourceIdentity  # type: ignore
    from ._models import SystemData  # type: ignore
    from ._models import TaggedResource  # type: ignore
    from ._models import UserAssignedIdentity  # type: ignore
    from ._models import Workspace  # type: ignore
    from ._models import WorkspaceList  # type: ignore
    from ._models import WorkspacePatchResource  # type: ignore
    from ._models import WorkspaceProperties  # type: ignore

from ._healthcare_apis_management_client_enums import (
    ActionType,
    CreatedByType,
    FhirResourceVersionPolicy,
    FhirServiceKind,
    IotIdentityResolutionType,
    Kind,
    ManagedServiceIdentityType,
    OperationResultStatus,
    PrivateEndpointConnectionProvisioningState,
    PrivateEndpointServiceConnectionStatus,
    ProvisioningState,
    PublicNetworkAccess,
    ServiceEventState,
    ServiceManagedIdentityType,
    ServiceNameUnavailabilityReason,
)

__all__ = [
    'CheckNameAvailabilityParameters',
    'DicomService',
    'DicomServiceAuthenticationConfiguration',
    'DicomServiceCollection',
    'DicomServicePatchResource',
    'Error',
    'ErrorDetails',
    'ErrorDetailsInternal',
    'FhirService',
    'FhirServiceAccessPolicyEntry',
    'FhirServiceAcrConfiguration',
    'FhirServiceAuthenticationConfiguration',
    'FhirServiceCollection',
    'FhirServiceCorsConfiguration',
    'FhirServiceExportConfiguration',
    'FhirServicePatchResource',
    'IotConnector',
    'IotConnectorCollection',
    'IotConnectorPatchResource',
    'IotDestinationProperties',
    'IotEventHubIngestionEndpointConfiguration',
    'IotFhirDestination',
    'IotFhirDestinationCollection',
    'IotFhirDestinationProperties',
    'IotMappingProperties',
    'ListOperations',
    'LocationBasedResource',
    'LogSpecification',
    'MetricDimension',
    'MetricSpecification',
    'OperationDetail',
    'OperationDisplay',
    'OperationProperties',
    'OperationResultsDescription',
    'PrivateEndpoint',
    'PrivateEndpointConnection',
    'PrivateEndpointConnectionDescription',
    'PrivateEndpointConnectionListResult',
    'PrivateEndpointConnectionListResultDescription',
    'PrivateLinkResource',
    'PrivateLinkResourceDescription',
    'PrivateLinkResourceListResultDescription',
    'PrivateLinkServiceConnectionState',
    'Resource',
    'ResourceCore',
    'ResourceTags',
    'ResourceVersionPolicyConfiguration',
    'ServiceAccessPolicyEntry',
    'ServiceAcrConfigurationInfo',
    'ServiceAuthenticationConfigurationInfo',
    'ServiceCorsConfigurationInfo',
    'ServiceCosmosDbConfigurationInfo',
    'ServiceExportConfigurationInfo',
    'ServiceManagedIdentity',
    'ServiceManagedIdentityautogenerated',
    'ServiceOciArtifactEntry',
    'ServiceSpecification',
    'ServicesDescription',
    'ServicesDescriptionListResult',
    'ServicesNameAvailabilityInfo',
    'ServicesPatchDescription',
    'ServicesProperties',
    'ServicesResource',
    'ServicesResourceIdentity',
    'SystemData',
    'TaggedResource',
    'UserAssignedIdentity',
    'Workspace',
    'WorkspaceList',
    'WorkspacePatchResource',
    'WorkspaceProperties',
    'ActionType',
    'CreatedByType',
    'FhirResourceVersionPolicy',
    'FhirServiceKind',
    'IotIdentityResolutionType',
    'Kind',
    'ManagedServiceIdentityType',
    'OperationResultStatus',
    'PrivateEndpointConnectionProvisioningState',
    'PrivateEndpointServiceConnectionStatus',
    'ProvisioningState',
    'PublicNetworkAccess',
    'ServiceEventState',
    'ServiceManagedIdentityType',
    'ServiceNameUnavailabilityReason',
]
