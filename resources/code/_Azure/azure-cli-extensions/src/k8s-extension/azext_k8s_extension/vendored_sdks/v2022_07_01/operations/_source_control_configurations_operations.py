# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
import sys
from typing import Any, Callable, Dict, IO, Iterable, Optional, TypeVar, Union, cast, overload
import urllib.parse

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ResourceNotModifiedError,
    map_error,
)
from azure.core.paging import ItemPaged
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import HttpResponse
from azure.core.polling import LROPoller, NoPolling, PollingMethod
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator import distributed_trace
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat
from azure.mgmt.core.polling.arm_polling import ARMPolling

from .. import models as _models
from ..._serialization import Serializer
from .._vendor import _convert_request, _format_url_section

if sys.version_info >= (3, 8):
    from typing import Literal  # pylint: disable=no-name-in-module, ungrouped-imports
else:
    from typing_extensions import Literal  # type: ignore  # pylint: disable=ungrouped-imports
T = TypeVar("T")
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, HttpResponse], T, Dict[str, Any]], Any]]

_SERIALIZER = Serializer()
_SERIALIZER.client_side_validation = False


def build_get_request(
    resource_group_name: str,
    cluster_rp: str,
    cluster_resource_name: str,
    cluster_name: str,
    source_control_configuration_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str", min_length=1),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1
        ),
        "clusterRp": _SERIALIZER.url("cluster_rp", cluster_rp, "str"),
        "clusterResourceName": _SERIALIZER.url("cluster_resource_name", cluster_resource_name, "str"),
        "clusterName": _SERIALIZER.url("cluster_name", cluster_name, "str"),
        "sourceControlConfigurationName": _SERIALIZER.url(
            "source_control_configuration_name", source_control_configuration_name, "str"
        ),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="GET", url=_url, params=_params, headers=_headers, **kwargs)


def build_create_or_update_request(
    resource_group_name: str,
    cluster_rp: str,
    cluster_resource_name: str,
    cluster_name: str,
    source_control_configuration_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
    content_type = kwargs.pop("content_type", _headers.pop("Content-Type", None))  # type: Optional[str]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str", min_length=1),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1
        ),
        "clusterRp": _SERIALIZER.url("cluster_rp", cluster_rp, "str"),
        "clusterResourceName": _SERIALIZER.url("cluster_resource_name", cluster_resource_name, "str"),
        "clusterName": _SERIALIZER.url("cluster_name", cluster_name, "str"),
        "sourceControlConfigurationName": _SERIALIZER.url(
            "source_control_configuration_name", source_control_configuration_name, "str"
        ),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    if content_type is not None:
        _headers["Content-Type"] = _SERIALIZER.header("content_type", content_type, "str")
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="PUT", url=_url, params=_params, headers=_headers, **kwargs)


def build_delete_request(
    resource_group_name: str,
    cluster_rp: str,
    cluster_resource_name: str,
    cluster_name: str,
    source_control_configuration_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str", min_length=1),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1
        ),
        "clusterRp": _SERIALIZER.url("cluster_rp", cluster_rp, "str"),
        "clusterResourceName": _SERIALIZER.url("cluster_resource_name", cluster_resource_name, "str"),
        "clusterName": _SERIALIZER.url("cluster_name", cluster_name, "str"),
        "sourceControlConfigurationName": _SERIALIZER.url(
            "source_control_configuration_name", source_control_configuration_name, "str"
        ),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="DELETE", url=_url, params=_params, headers=_headers, **kwargs)


def build_list_request(
    resource_group_name: str,
    cluster_rp: str,
    cluster_resource_name: str,
    cluster_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str", min_length=1),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1
        ),
        "clusterRp": _SERIALIZER.url("cluster_rp", cluster_rp, "str"),
        "clusterResourceName": _SERIALIZER.url("cluster_resource_name", cluster_resource_name, "str"),
        "clusterName": _SERIALIZER.url("cluster_name", cluster_name, "str"),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="GET", url=_url, params=_params, headers=_headers, **kwargs)


class SourceControlConfigurationsOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.kubernetesconfiguration.v2022_07_01.SourceControlConfigurationClient`'s
        :attr:`source_control_configurations` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs):
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")

    @distributed_trace
    def get(
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        **kwargs: Any
    ) -> _models.SourceControlConfiguration:
        """Gets details of the Source Control Configuration.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :param source_control_configuration_name: Name of the Source Control Configuration. Required.
        :type source_control_configuration_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: SourceControlConfiguration or the result of cls(response)
        :rtype: ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration
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

        api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.SourceControlConfiguration]

        request = build_get_request(
            resource_group_name=resource_group_name,
            cluster_rp=cluster_rp,
            cluster_resource_name=cluster_resource_name,
            cluster_name=cluster_name,
            source_control_configuration_name=source_control_configuration_name,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            template_url=self.get.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=False, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("SourceControlConfiguration", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"}  # type: ignore

    @overload
    def create_or_update(
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        source_control_configuration: _models.SourceControlConfiguration,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.SourceControlConfiguration:
        """Create a new Kubernetes Source Control Configuration.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :param source_control_configuration_name: Name of the Source Control Configuration. Required.
        :type source_control_configuration_name: str
        :param source_control_configuration: Properties necessary to Create KubernetesConfiguration.
         Required.
        :type source_control_configuration:
         ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: SourceControlConfiguration or the result of cls(response)
        :rtype: ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    def create_or_update(
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        source_control_configuration: IO,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.SourceControlConfiguration:
        """Create a new Kubernetes Source Control Configuration.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :param source_control_configuration_name: Name of the Source Control Configuration. Required.
        :type source_control_configuration_name: str
        :param source_control_configuration: Properties necessary to Create KubernetesConfiguration.
         Required.
        :type source_control_configuration: IO
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: SourceControlConfiguration or the result of cls(response)
        :rtype: ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace
    def create_or_update(
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        source_control_configuration: Union[_models.SourceControlConfiguration, IO],
        **kwargs: Any
    ) -> _models.SourceControlConfiguration:
        """Create a new Kubernetes Source Control Configuration.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :param source_control_configuration_name: Name of the Source Control Configuration. Required.
        :type source_control_configuration_name: str
        :param source_control_configuration: Properties necessary to Create KubernetesConfiguration. Is
         either a model type or a IO type. Required.
        :type source_control_configuration:
         ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration or IO
        :keyword content_type: Body Parameter content-type. Known values are: 'application/json'.
         Default value is None.
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: SourceControlConfiguration or the result of cls(response)
        :rtype: ~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
        content_type = kwargs.pop("content_type", _headers.pop("Content-Type", None))  # type: Optional[str]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.SourceControlConfiguration]

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(source_control_configuration, (IO, bytes)):
            _content = source_control_configuration
        else:
            _json = self._serialize.body(source_control_configuration, "SourceControlConfiguration")

        request = build_create_or_update_request(
            resource_group_name=resource_group_name,
            cluster_rp=cluster_rp,
            cluster_resource_name=cluster_resource_name,
            cluster_name=cluster_name,
            source_control_configuration_name=source_control_configuration_name,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            content_type=content_type,
            json=_json,
            content=_content,
            template_url=self.create_or_update.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=False, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200, 201]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        if response.status_code == 200:
            deserialized = self._deserialize("SourceControlConfiguration", pipeline_response)

        if response.status_code == 201:
            deserialized = self._deserialize("SourceControlConfiguration", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    create_or_update.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"}  # type: ignore

    def _delete_initial(  # pylint: disable=inconsistent-return-statements
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        **kwargs: Any
    ) -> None:
        error_map = {
            401: ClientAuthenticationError,
            404: ResourceNotFoundError,
            409: ResourceExistsError,
            304: ResourceNotModifiedError,
        }
        error_map.update(kwargs.pop("error_map", {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
        cls = kwargs.pop("cls", None)  # type: ClsType[None]

        request = build_delete_request(
            resource_group_name=resource_group_name,
            cluster_rp=cluster_rp,
            cluster_resource_name=cluster_resource_name,
            cluster_name=cluster_name,
            source_control_configuration_name=source_control_configuration_name,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            template_url=self._delete_initial.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=False, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200, 204]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        if cls:
            return cls(pipeline_response, None, {})

    _delete_initial.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"}  # type: ignore

    @distributed_trace
    def begin_delete(
        self,
        resource_group_name: str,
        cluster_rp: str,
        cluster_resource_name: str,
        cluster_name: str,
        source_control_configuration_name: str,
        **kwargs: Any
    ) -> LROPoller[None]:
        """This will delete the YAML file used to set up the Source control configuration, thus stopping
        future sync from the source repo.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :param source_control_configuration_name: Name of the Source Control Configuration. Required.
        :type source_control_configuration_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :keyword str continuation_token: A continuation token to restart a poller from a saved state.
        :keyword polling: By default, your polling method will be ARMPolling. Pass in False for this
         operation to not poll, or pass in your own initialized polling object for a personal polling
         strategy.
        :paramtype polling: bool or ~azure.core.polling.PollingMethod
        :keyword int polling_interval: Default waiting time between two polls for LRO operations if no
         Retry-After header is present.
        :return: An instance of LROPoller that returns either None or the result of cls(response)
        :rtype: ~azure.core.polling.LROPoller[None]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
        cls = kwargs.pop("cls", None)  # type: ClsType[None]
        polling = kwargs.pop("polling", True)  # type: Union[bool, PollingMethod]
        lro_delay = kwargs.pop("polling_interval", self._config.polling_interval)
        cont_token = kwargs.pop("continuation_token", None)  # type: Optional[str]
        if cont_token is None:
            raw_result = self._delete_initial(  # type: ignore
                resource_group_name=resource_group_name,
                cluster_rp=cluster_rp,
                cluster_resource_name=cluster_resource_name,
                cluster_name=cluster_name,
                source_control_configuration_name=source_control_configuration_name,
                api_version=api_version,
                cls=lambda x, y, z: x,
                headers=_headers,
                params=_params,
                **kwargs
            )
        kwargs.pop("error_map", None)

        def get_long_running_output(pipeline_response):  # pylint: disable=inconsistent-return-statements
            if cls:
                return cls(pipeline_response, None, {})

        if polling is True:
            polling_method = cast(PollingMethod, ARMPolling(lro_delay, **kwargs))  # type: PollingMethod
        elif polling is False:
            polling_method = cast(PollingMethod, NoPolling())
        else:
            polling_method = polling
        if cont_token:
            return LROPoller.from_continuation_token(
                polling_method=polling_method,
                continuation_token=cont_token,
                client=self._client,
                deserialization_callback=get_long_running_output,
            )
        return LROPoller(self._client, raw_result, get_long_running_output, polling_method)

    begin_delete.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/{sourceControlConfigurationName}"}  # type: ignore

    @distributed_trace
    def list(
        self, resource_group_name: str, cluster_rp: str, cluster_resource_name: str, cluster_name: str, **kwargs: Any
    ) -> Iterable["_models.SourceControlConfiguration"]:
        """List all Source Control Configurations.

        :param resource_group_name: The name of the resource group. The name is case insensitive.
         Required.
        :type resource_group_name: str
        :param cluster_rp: The Kubernetes cluster RP - i.e. Microsoft.ContainerService,
         Microsoft.Kubernetes, Microsoft.HybridContainerService. Required.
        :type cluster_rp: str
        :param cluster_resource_name: The Kubernetes cluster resource name - i.e. managedClusters,
         connectedClusters, provisionedClusters. Required.
        :type cluster_resource_name: str
        :param cluster_name: The name of the kubernetes cluster. Required.
        :type cluster_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: An iterator like instance of either SourceControlConfiguration or the result of
         cls(response)
        :rtype:
         ~azure.core.paging.ItemPaged[~azure.mgmt.kubernetesconfiguration.v2022_07_01.models.SourceControlConfiguration]
        :raises ~azure.core.exceptions.HttpResponseError:
        """
        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop("api_version", _params.pop("api-version", "2022-07-01"))  # type: Literal["2022-07-01"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.SourceControlConfigurationList]

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
                    resource_group_name=resource_group_name,
                    cluster_rp=cluster_rp,
                    cluster_resource_name=cluster_resource_name,
                    cluster_name=cluster_name,
                    subscription_id=self._config.subscription_id,
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

        def extract_data(pipeline_response):
            deserialized = self._deserialize("SourceControlConfigurationList", pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)
            return deserialized.next_link or None, iter(list_of_elem)

        def get_next(next_link=None):
            request = prepare_request(next_link)

            pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
                request, stream=False, **kwargs
            )
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
                raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

            return pipeline_response

        return ItemPaged(get_next, extract_data)

    list.metadata = {"url": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations"}  # type: ignore
