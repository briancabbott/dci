# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from ._models_py3 import AddressResponse
from ._models_py3 import ApiDefinitionInfo
from ._models_py3 import AppServiceEnvironmentCollection
from ._models_py3 import AppServiceEnvironmentPatchResource
from ._models_py3 import AppServiceEnvironmentResource
from ._models_py3 import AppServicePlan
from ._models_py3 import AppServicePlanCollection
from ._models_py3 import AppServicePlanPatchResource
from ._models_py3 import AutoHealActions
from ._models_py3 import AutoHealCustomAction
from ._models_py3 import AutoHealRules
from ._models_py3 import AutoHealTriggers
from ._models_py3 import Capability
from ._models_py3 import CloningInfo
from ._models_py3 import ConnStringInfo
from ._models_py3 import CorsSettings
from ._models_py3 import CsmUsageQuota
from ._models_py3 import CsmUsageQuotaCollection
from ._models_py3 import ErrorEntity
from ._models_py3 import Experiments
from ._models_py3 import HandlerMapping
from ._models_py3 import HostNameSslState
from ._models_py3 import HostingEnvironmentDiagnostics
from ._models_py3 import HostingEnvironmentProfile
from ._models_py3 import HybridConnection
from ._models_py3 import HybridConnectionCollection
from ._models_py3 import HybridConnectionKey
from ._models_py3 import HybridConnectionLimits
from ._models_py3 import IpSecurityRestriction
from ._models_py3 import LocalizableString
from ._models_py3 import ManagedServiceIdentity
from ._models_py3 import MetricAvailabilily
from ._models_py3 import MetricDefinition
from ._models_py3 import NameValuePair
from ._models_py3 import NetworkAccessControlEntry
from ._models_py3 import Operation
from ._models_py3 import ProxyOnlyResource
from ._models_py3 import PushSettings
from ._models_py3 import RampUpRule
from ._models_py3 import RequestsBasedTrigger
from ._models_py3 import Resource
from ._models_py3 import ResourceCollection
from ._models_py3 import ResourceMetric
from ._models_py3 import ResourceMetricAvailability
from ._models_py3 import ResourceMetricCollection
from ._models_py3 import ResourceMetricDefinition
from ._models_py3 import ResourceMetricDefinitionCollection
from ._models_py3 import ResourceMetricName
from ._models_py3 import ResourceMetricProperty
from ._models_py3 import ResourceMetricValue
from ._models_py3 import Site
from ._models_py3 import SiteConfig
from ._models_py3 import SiteLimits
from ._models_py3 import SiteMachineKey
from ._models_py3 import SkuCapacity
from ._models_py3 import SkuDescription
from ._models_py3 import SkuInfo
from ._models_py3 import SkuInfoCollection
from ._models_py3 import SlotSwapStatus
from ._models_py3 import SlowRequestsBasedTrigger
from ._models_py3 import SnapshotRecoveryRequest
from ._models_py3 import SnapshotRecoveryTarget
from ._models_py3 import StampCapacity
from ._models_py3 import StampCapacityCollection
from ._models_py3 import StatusCodesBasedTrigger
from ._models_py3 import Usage
from ._models_py3 import UsageCollection
from ._models_py3 import VirtualApplication
from ._models_py3 import VirtualDirectory
from ._models_py3 import VirtualIPMapping
from ._models_py3 import VirtualNetworkProfile
from ._models_py3 import VnetGateway
from ._models_py3 import VnetInfo
from ._models_py3 import VnetRoute
from ._models_py3 import WebAppCollection
from ._models_py3 import WorkerPool
from ._models_py3 import WorkerPoolCollection
from ._models_py3 import WorkerPoolResource

from ._web_site_management_client_enums import AccessControlEntryAction
from ._web_site_management_client_enums import AutoHealActionType
from ._web_site_management_client_enums import ComputeModeOptions
from ._web_site_management_client_enums import ConnectionStringType
from ._web_site_management_client_enums import HostType
from ._web_site_management_client_enums import HostingEnvironmentStatus
from ._web_site_management_client_enums import InternalLoadBalancingMode
from ._web_site_management_client_enums import ManagedPipelineMode
from ._web_site_management_client_enums import ManagedServiceIdentityType
from ._web_site_management_client_enums import OperationStatus
from ._web_site_management_client_enums import ProvisioningState
from ._web_site_management_client_enums import RouteType
from ._web_site_management_client_enums import ScmType
from ._web_site_management_client_enums import SiteAvailabilityState
from ._web_site_management_client_enums import SiteLoadBalancing
from ._web_site_management_client_enums import SslState
from ._web_site_management_client_enums import StatusOptions
from ._web_site_management_client_enums import SupportedTlsVersions
from ._web_site_management_client_enums import UsageState
from ._web_site_management_client_enums import WorkerSizeOptions
from ._patch import __all__ as _patch_all
from ._patch import *  # pylint: disable=unused-wildcard-import
from ._patch import patch_sdk as _patch_sdk

__all__ = [
    "AddressResponse",
    "ApiDefinitionInfo",
    "AppServiceEnvironmentCollection",
    "AppServiceEnvironmentPatchResource",
    "AppServiceEnvironmentResource",
    "AppServicePlan",
    "AppServicePlanCollection",
    "AppServicePlanPatchResource",
    "AutoHealActions",
    "AutoHealCustomAction",
    "AutoHealRules",
    "AutoHealTriggers",
    "Capability",
    "CloningInfo",
    "ConnStringInfo",
    "CorsSettings",
    "CsmUsageQuota",
    "CsmUsageQuotaCollection",
    "ErrorEntity",
    "Experiments",
    "HandlerMapping",
    "HostNameSslState",
    "HostingEnvironmentDiagnostics",
    "HostingEnvironmentProfile",
    "HybridConnection",
    "HybridConnectionCollection",
    "HybridConnectionKey",
    "HybridConnectionLimits",
    "IpSecurityRestriction",
    "LocalizableString",
    "ManagedServiceIdentity",
    "MetricAvailabilily",
    "MetricDefinition",
    "NameValuePair",
    "NetworkAccessControlEntry",
    "Operation",
    "ProxyOnlyResource",
    "PushSettings",
    "RampUpRule",
    "RequestsBasedTrigger",
    "Resource",
    "ResourceCollection",
    "ResourceMetric",
    "ResourceMetricAvailability",
    "ResourceMetricCollection",
    "ResourceMetricDefinition",
    "ResourceMetricDefinitionCollection",
    "ResourceMetricName",
    "ResourceMetricProperty",
    "ResourceMetricValue",
    "Site",
    "SiteConfig",
    "SiteLimits",
    "SiteMachineKey",
    "SkuCapacity",
    "SkuDescription",
    "SkuInfo",
    "SkuInfoCollection",
    "SlotSwapStatus",
    "SlowRequestsBasedTrigger",
    "SnapshotRecoveryRequest",
    "SnapshotRecoveryTarget",
    "StampCapacity",
    "StampCapacityCollection",
    "StatusCodesBasedTrigger",
    "Usage",
    "UsageCollection",
    "VirtualApplication",
    "VirtualDirectory",
    "VirtualIPMapping",
    "VirtualNetworkProfile",
    "VnetGateway",
    "VnetInfo",
    "VnetRoute",
    "WebAppCollection",
    "WorkerPool",
    "WorkerPoolCollection",
    "WorkerPoolResource",
    "AccessControlEntryAction",
    "AutoHealActionType",
    "ComputeModeOptions",
    "ConnectionStringType",
    "HostType",
    "HostingEnvironmentStatus",
    "InternalLoadBalancingMode",
    "ManagedPipelineMode",
    "ManagedServiceIdentityType",
    "OperationStatus",
    "ProvisioningState",
    "RouteType",
    "ScmType",
    "SiteAvailabilityState",
    "SiteLoadBalancing",
    "SslState",
    "StatusOptions",
    "SupportedTlsVersions",
    "UsageState",
    "WorkerSizeOptions",
]
__all__.extend([p for p in _patch_all if p not in __all__])
_patch_sdk()
