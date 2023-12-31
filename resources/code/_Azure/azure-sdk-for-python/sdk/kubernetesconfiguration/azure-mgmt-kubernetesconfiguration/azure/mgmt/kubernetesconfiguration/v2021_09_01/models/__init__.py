# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from ._models_py3 import ErrorAdditionalInfo
from ._models_py3 import ErrorDetail
from ._models_py3 import ErrorResponse
from ._models_py3 import Extension
from ._models_py3 import ExtensionPropertiesAksAssignedIdentity
from ._models_py3 import ExtensionStatus
from ._models_py3 import ExtensionsList
from ._models_py3 import Identity
from ._models_py3 import OperationStatusList
from ._models_py3 import OperationStatusResult
from ._models_py3 import PatchExtension
from ._models_py3 import ProxyResource
from ._models_py3 import Resource
from ._models_py3 import ResourceProviderOperation
from ._models_py3 import ResourceProviderOperationDisplay
from ._models_py3 import ResourceProviderOperationList
from ._models_py3 import Scope
from ._models_py3 import ScopeCluster
from ._models_py3 import ScopeNamespace
from ._models_py3 import SystemData

from ._source_control_configuration_client_enums import CreatedByType
from ._source_control_configuration_client_enums import Enum0
from ._source_control_configuration_client_enums import Enum1
from ._source_control_configuration_client_enums import LevelType
from ._source_control_configuration_client_enums import ProvisioningState
from ._patch import __all__ as _patch_all
from ._patch import *  # pylint: disable=unused-wildcard-import
from ._patch import patch_sdk as _patch_sdk

__all__ = [
    "ErrorAdditionalInfo",
    "ErrorDetail",
    "ErrorResponse",
    "Extension",
    "ExtensionPropertiesAksAssignedIdentity",
    "ExtensionStatus",
    "ExtensionsList",
    "Identity",
    "OperationStatusList",
    "OperationStatusResult",
    "PatchExtension",
    "ProxyResource",
    "Resource",
    "ResourceProviderOperation",
    "ResourceProviderOperationDisplay",
    "ResourceProviderOperationList",
    "Scope",
    "ScopeCluster",
    "ScopeNamespace",
    "SystemData",
    "CreatedByType",
    "Enum0",
    "Enum1",
    "LevelType",
    "ProvisioningState",
]
__all__.extend([p for p in _patch_all if p not in __all__])
_patch_sdk()
