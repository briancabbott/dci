# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
import datetime
import sys
from typing import Any, AsyncIterable, Callable, Dict, Optional, TypeVar
import urllib.parse

from azure.core.async_paging import AsyncItemPaged, AsyncList
from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ResourceNotModifiedError,
    map_error,
)
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import AsyncHttpResponse
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator import distributed_trace
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat

from ... import models as _models
from ..._vendor import _convert_request
from ...operations._health_monitors_operations import (
    build_get_request,
    build_get_state_change_request,
    build_list_request,
    build_list_state_changes_request,
)

if sys.version_info >= (3, 8):
    from typing import Literal  # pylint: disable=no-name-in-module, ungrouped-imports
else:
    from typing_extensions import Literal  # type: ignore  # pylint: disable=ungrouped-imports
T = TypeVar("T")
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, AsyncHttpResponse], T, Dict[str, Any]], Any]]


class HealthMonitorsOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.workloadmonitor.aio.WorkloadMonitorAPI`'s
        :attr:`health_monitors` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs) -> None:
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")

    @distributed_trace
    def list(
        self,
        subscription_id: str,
        resource_group_name: str,
        provider_name: str,
        resource_collection_name: str,
        resource_name: str,
        filter: Optional[str] = None,
        expand: Optional[str] = None,
        **kwargs: Any
    ) -> AsyncIterable["_models.HealthMonitor"]:
        """Get the current health status of all monitors of a virtual machine. Optional parameters:
        $expand (retrieve the monitor's evidence and configuration) and $filter (filter by monitor
        name).

        Get the current health status of all monitors of a virtual machine. Optional parameters:
        $expand (retrieve the monitor's evidence and configuration) and $filter (filter by monitor
        name).

        :param subscription_id: The subscription Id of the virtual machine. Required.
        :type subscription_id: str
        :param resource_group_name: The resource group of the virtual machine. Required.
        :type resource_group_name: str
        :param provider_name: The provider name (ex: Microsoft.Compute for virtual machines). Required.
        :type provider_name: str
        :param resource_collection_name: The resource collection name (ex: virtualMachines for virtual
         machines). Required.
        :type resource_collection_name: str
        :param resource_name: The name of the virtual machine. Required.
        :type resource_name: str
        :param filter: Optionally filter by monitor name. Example: $filter=monitorName eq
         'logical-disks|C:|disk-free-space-mb.'. Default value is None.
        :type filter: str
        :param expand: Optionally expand the monitor’s evidence and/or configuration. Example:
         $expand=evidence,configuration. Default value is None.
        :type expand: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: An iterator like instance of either HealthMonitor or the result of cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.workloadmonitor.models.HealthMonitor]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop(
            "api_version", _params.pop("api-version", self._config.api_version)
        )  # type: Literal["2020-01-13-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.HealthMonitorList]

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                request = build_list_request(
                    subscription_id=subscription_id,
                    resource_group_name=resource_group_name,
                    provider_name=provider_name,
                    resource_collection_name=resource_collection_name,
                    resource_name=resource_name,
                    filter=filter,
                    expand=expand,
                    api_version=api_version,
                    template_url=self.list.metadata["url"],
                    headers=_headers,
                    params=_params,
                )
                request = _convert_request(request)
                request.url = self._client.format_url(request.url)  # type: ignore

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._config.api_version
                request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                request = _convert_request(request)
                request.url = self._client.format_url(request.url)  # type: ignore
                request.method = "GET"
            return request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("HealthMonitorList", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)
            return deserialized.next_link or None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            request = prepare_request(next_link)

            pipeline_response = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
                request, stream=False, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
                raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)

    list.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors"}  # type: ignore

    @distributed_trace_async
    async def get(
        self,
        subscription_id: str,
        resource_group_name: str,
        provider_name: str,
        resource_collection_name: str,
        resource_name: str,
        monitor_id: str,
        expand: Optional[str] = None,
        **kwargs: Any
    ) -> _models.HealthMonitor:
        """Get the current health status of a monitor of a virtual machine. Optional parameter: $expand
        (retrieve the monitor's evidence and configuration).

        Get the current health status of a monitor of a virtual machine. Optional parameter: $expand
        (retrieve the monitor's evidence and configuration).

        :param subscription_id: The subscription Id of the virtual machine. Required.
        :type subscription_id: str
        :param resource_group_name: The resource group of the virtual machine. Required.
        :type resource_group_name: str
        :param provider_name: The provider name (ex: Microsoft.Compute for virtual machines). Required.
        :type provider_name: str
        :param resource_collection_name: The resource collection name (ex: virtualMachines for virtual
         machines). Required.
        :type resource_collection_name: str
        :param resource_name: The name of the virtual machine. Required.
        :type resource_name: str
        :param monitor_id: The monitor Id of the virtual machine. Required.
        :type monitor_id: str
        :param expand: Optionally expand the monitor’s evidence and/or configuration. Example:
         $expand=evidence,configuration. Default value is None.
        :type expand: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: HealthMonitor or the result of cls(response)
        :rtype: ~azure.mgmt.workloadmonitor.models.HealthMonitor
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop(
            "api_version", _params.pop("api-version", self._config.api_version)
        )  # type: Literal["2020-01-13-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.HealthMonitor]

        request = build_get_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            provider_name=provider_name,
            resource_collection_name=resource_collection_name,
            resource_name=resource_name,
            monitor_id=monitor_id,
            expand=expand,
            api_version=api_version,
            template_url=self.get.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=False, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("HealthMonitor", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}"}  # type: ignore

    @distributed_trace
    def list_state_changes(
        self,
        subscription_id: str,
        resource_group_name: str,
        provider_name: str,
        resource_collection_name: str,
        resource_name: str,
        monitor_id: str,
        filter: Optional[str] = None,
        expand: Optional[str] = None,
        start_timestamp_utc: Optional[datetime.datetime] = None,
        end_timestamp_utc: Optional[datetime.datetime] = None,
        **kwargs: Any
    ) -> AsyncIterable["_models.HealthMonitorStateChange"]:
        """Get the health state changes of a monitor of a virtual machine within the provided time window
        (default is the last 24 hours). Optional parameters: $expand (retrieve the monitor's evidence
        and configuration) and $filter (filter by heartbeat condition).

        Get the health state changes of a monitor of a virtual machine within the provided time window
        (default is the last 24 hours). Optional parameters: $expand (retrieve the monitor's evidence
        and configuration) and $filter (filter by heartbeat condition).

        :param subscription_id: The subscription Id of the virtual machine. Required.
        :type subscription_id: str
        :param resource_group_name: The resource group of the virtual machine. Required.
        :type resource_group_name: str
        :param provider_name: The provider name (ex: Microsoft.Compute for virtual machines). Required.
        :type provider_name: str
        :param resource_collection_name: The resource collection name (ex: virtualMachines for virtual
         machines). Required.
        :type resource_collection_name: str
        :param resource_name: The name of the virtual machine. Required.
        :type resource_name: str
        :param monitor_id: The monitor Id of the virtual machine. Required.
        :type monitor_id: str
        :param filter: Optionally filter by heartbeat condition. Example: $filter=isHeartbeat eq false.
         Default value is None.
        :type filter: str
        :param expand: Optionally expand the monitor’s evidence and/or configuration. Example:
         $expand=evidence,configuration. Default value is None.
        :type expand: str
        :param start_timestamp_utc: The start of the time window. Default value is None.
        :type start_timestamp_utc: ~datetime.datetime
        :param end_timestamp_utc: The end of the time window. Default value is None.
        :type end_timestamp_utc: ~datetime.datetime
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: An iterator like instance of either HealthMonitorStateChange or the result of
         cls(response)
        :rtype:
         ~azure.core.async_paging.AsyncItemPaged[~azure.mgmt.workloadmonitor.models.HealthMonitorStateChange]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop(
            "api_version", _params.pop("api-version", self._config.api_version)
        )  # type: Literal["2020-01-13-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.HealthMonitorStateChangeList]

        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        def prepare_request(next_link=None):
            if not next_link:

                request = build_list_state_changes_request(
                    subscription_id=subscription_id,
                    resource_group_name=resource_group_name,
                    provider_name=provider_name,
                    resource_collection_name=resource_collection_name,
                    resource_name=resource_name,
                    monitor_id=monitor_id,
                    filter=filter,
                    expand=expand,
                    start_timestamp_utc=start_timestamp_utc,
                    end_timestamp_utc=end_timestamp_utc,
                    api_version=api_version,
                    template_url=self.list_state_changes.metadata["url"],
                    headers=_headers,
                    params=_params,
                )
                request = _convert_request(request)
                request.url = self._client.format_url(request.url)  # type: ignore

            else:
                # make call to next link with the client's api-version
                _parsed_next_link = urllib.parse.urlparse(next_link)
                _next_request_params = case_insensitive_dict(
                    {
                        key: [urllib.parse.quote(v) for v in value]
                        for key, value in urllib.parse.parse_qs(_parsed_next_link.query).items()
                    }
                )
                _next_request_params["api-version"] = self._config.api_version
                request = HttpRequest(
                    "GET", urllib.parse.urljoin(next_link, _parsed_next_link.path), params=_next_request_params
                )
                request = _convert_request(request)
                request.url = self._client.format_url(request.url)  # type: ignore
                request.method = "GET"
            return request

        async def extract_data(pipeline_response):
            deserialized = self._deserialize("HealthMonitorStateChangeList", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)
            return deserialized.next_link or None, AsyncList(list_of_elem)

        async def get_next(next_link=None):
            request = prepare_request(next_link)

            pipeline_response = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
                request, stream=False, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
                raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

            return pipeline_response

        return AsyncItemPaged(get_next, extract_data)

    list_state_changes.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history"}  # type: ignore

    @distributed_trace_async
    async def get_state_change(
        self,
        subscription_id: str,
        resource_group_name: str,
        provider_name: str,
        resource_collection_name: str,
        resource_name: str,
        monitor_id: str,
        timestamp_unix: str,
        expand: Optional[str] = None,
        **kwargs: Any
    ) -> _models.HealthMonitorStateChange:
        """Get the health state change of a monitor of a virtual machine at the provided timestamp.
        Optional parameter: $expand (retrieve the monitor's evidence and configuration).

        Get the health state change of a monitor of a virtual machine at the provided timestamp.
        Optional parameter: $expand (retrieve the monitor's evidence and configuration).

        :param subscription_id: The subscription Id of the virtual machine. Required.
        :type subscription_id: str
        :param resource_group_name: The resource group of the virtual machine. Required.
        :type resource_group_name: str
        :param provider_name: The provider name (ex: Microsoft.Compute for virtual machines). Required.
        :type provider_name: str
        :param resource_collection_name: The resource collection name (ex: virtualMachines for virtual
         machines). Required.
        :type resource_collection_name: str
        :param resource_name: The name of the virtual machine. Required.
        :type resource_name: str
        :param monitor_id: The monitor Id of the virtual machine. Required.
        :type monitor_id: str
        :param timestamp_unix: The timestamp of the state change (unix format). Required.
        :type timestamp_unix: str
        :param expand: Optionally expand the monitor’s evidence and/or configuration. Example:
         $expand=evidence,configuration. Default value is None.
        :type expand: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: HealthMonitorStateChange or the result of cls(response)
        :rtype: ~azure.mgmt.workloadmonitor.models.HealthMonitorStateChange
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop(
            "api_version", _params.pop("api-version", self._config.api_version)
        )  # type: Literal["2020-01-13-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.HealthMonitorStateChange]

        request = build_get_state_change_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            provider_name=provider_name,
            resource_collection_name=resource_collection_name,
            resource_name=resource_name,
            monitor_id=monitor_id,
            timestamp_unix=timestamp_unix,
            expand=expand,
            api_version=api_version,
            template_url=self.get_state_change.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=False, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("HealthMonitorStateChange", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get_state_change.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history/{timestampUnix}"}  # type: ignore
