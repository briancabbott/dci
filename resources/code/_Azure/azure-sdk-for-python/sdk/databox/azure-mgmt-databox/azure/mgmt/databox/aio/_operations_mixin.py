# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------
from .._serialization import Serializer, Deserializer
from typing import Any, IO, Optional, Union

from .. import models as _models


class DataBoxManagementClientOperationsMixin(object):

    async def mitigate(  # pylint: disable=inconsistent-return-statements
        self,
        job_name: str,
        resource_group_name: str,
        mitigate_job_request: Union[_models.MitigateJobRequest, IO],
        **kwargs: Any
    ) -> None:
        """Request to mitigate for a given job.

        :param job_name: The name of the job Resource within the specified resource group. job names
         must be between 3 and 24 characters in length and use any alphanumeric and underscore only.
         Required.
        :type job_name: str
        :param resource_group_name: The Resource Group Name. Required.
        :type resource_group_name: str
        :param mitigate_job_request: Mitigation Request. Is either a MitigateJobRequest type or a IO
         type. Required.
        :type mitigate_job_request: ~azure.mgmt.databox.v2022_12_01.models.MitigateJobRequest or IO
        :keyword content_type: Body Parameter content-type. Known values are: 'application/json'.
         Default value is None.
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: None or the result of cls(response)
        :rtype: None
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        api_version = self._get_api_version('mitigate')
        if api_version == '2021-03-01':
            from ..v2021_03_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2021-05-01':
            from ..v2021_05_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2021-08-01-preview':
            from ..v2021_08_01_preview.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2021-12-01':
            from ..v2021_12_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2022-02-01':
            from ..v2022_02_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2022-09-01':
            from ..v2022_09_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2022-10-01':
            from ..v2022_10_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        elif api_version == '2022-12-01':
            from ..v2022_12_01.aio.operations import DataBoxManagementClientOperationsMixin as OperationClass
        else:
            raise ValueError("API version {} does not have operation 'mitigate'".format(api_version))
        mixin_instance = OperationClass()
        mixin_instance._client = self._client
        mixin_instance._config = self._config
        mixin_instance._config.api_version = api_version
        mixin_instance._serialize = Serializer(self._models_dict(api_version))
        mixin_instance._serialize.client_side_validation = False
        mixin_instance._deserialize = Deserializer(self._models_dict(api_version))
        return await mixin_instance.mitigate(job_name, resource_group_name, mitigate_job_request, **kwargs)
