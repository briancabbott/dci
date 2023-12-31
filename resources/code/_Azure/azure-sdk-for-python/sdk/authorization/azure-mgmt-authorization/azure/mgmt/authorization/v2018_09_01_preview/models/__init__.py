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
from ._models_py3 import RoleAssignment
from ._models_py3 import RoleAssignmentCreateParameters
from ._models_py3 import RoleAssignmentFilter
from ._models_py3 import RoleAssignmentListResult

from ._authorization_management_client_enums import PrincipalType
from ._patch import __all__ as _patch_all
from ._patch import *  # pylint: disable=unused-wildcard-import
from ._patch import patch_sdk as _patch_sdk

__all__ = [
    "ErrorAdditionalInfo",
    "ErrorDetail",
    "ErrorResponse",
    "RoleAssignment",
    "RoleAssignmentCreateParameters",
    "RoleAssignmentFilter",
    "RoleAssignmentListResult",
    "PrincipalType",
]
__all__.extend([p for p in _patch_all if p not in __all__])
_patch_sdk()
