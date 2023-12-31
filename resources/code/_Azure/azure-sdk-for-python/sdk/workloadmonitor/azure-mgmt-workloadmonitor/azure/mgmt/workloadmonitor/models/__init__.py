# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from ._models_py3 import ErrorDetails
from ._models_py3 import ErrorResponse
from ._models_py3 import ErrorResponseError
from ._models_py3 import HealthMonitor
from ._models_py3 import HealthMonitorList
from ._models_py3 import HealthMonitorStateChange
from ._models_py3 import HealthMonitorStateChangeList
from ._models_py3 import Operation
from ._models_py3 import OperationDisplay
from ._models_py3 import OperationList
from ._models_py3 import Resource

from ._workload_monitor_api_enums import HealthState
from ._patch import __all__ as _patch_all
from ._patch import *  # type: ignore # pylint: disable=unused-wildcard-import
from ._patch import patch_sdk as _patch_sdk

__all__ = [
    "ErrorDetails",
    "ErrorResponse",
    "ErrorResponseError",
    "HealthMonitor",
    "HealthMonitorList",
    "HealthMonitorStateChange",
    "HealthMonitorStateChangeList",
    "Operation",
    "OperationDisplay",
    "OperationList",
    "Resource",
    "HealthState",
]
__all__.extend([p for p in _patch_all if p not in __all__])
_patch_sdk()
