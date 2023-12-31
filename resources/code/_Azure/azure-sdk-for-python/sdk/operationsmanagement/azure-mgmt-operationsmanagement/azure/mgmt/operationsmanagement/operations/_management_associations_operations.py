# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
import sys
from typing import Any, Callable, Dict, IO, Optional, TypeVar, Union, overload

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ResourceExistsError,
    ResourceNotFoundError,
    ResourceNotModifiedError,
    map_error,
)
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import HttpResponse
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator import distributed_trace
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat

from .. import models as _models
from .._serialization import Serializer
from .._vendor import _convert_request, _format_url_section

if sys.version_info >= (3, 8):
    from typing import Literal  # pylint: disable=no-name-in-module, ungrouped-imports
else:
    from typing_extensions import Literal  # type: ignore  # pylint: disable=ungrouped-imports
T = TypeVar("T")
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, HttpResponse], T, Dict[str, Any]], Any]]

_SERIALIZER = Serializer()
_SERIALIZER.client_side_validation = False


def build_list_by_subscription_request(subscription_id: str, **kwargs: Any) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop(
        "api_version", _params.pop("api-version", "2015-11-01-preview")
    )  # type: Literal["2015-11-01-preview"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str"),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="GET", url=_url, params=_params, headers=_headers, **kwargs)


def build_create_or_update_request(
    resource_group_name: str,
    provider_name: str,
    resource_type: str,
    resource_name: str,
    management_association_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop(
        "api_version", _params.pop("api-version", "2015-11-01-preview")
    )  # type: Literal["2015-11-01-preview"]
    content_type = kwargs.pop("content_type", _headers.pop("Content-Type", None))  # type: Optional[str]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str"),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1, pattern=r"^[-\w\._\(\)]+$"
        ),
        "providerName": _SERIALIZER.url("provider_name", provider_name, "str"),
        "resourceType": _SERIALIZER.url("resource_type", resource_type, "str"),
        "resourceName": _SERIALIZER.url("resource_name", resource_name, "str"),
        "managementAssociationName": _SERIALIZER.url("management_association_name", management_association_name, "str"),
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
    provider_name: str,
    resource_type: str,
    resource_name: str,
    management_association_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop(
        "api_version", _params.pop("api-version", "2015-11-01-preview")
    )  # type: Literal["2015-11-01-preview"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str"),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1, pattern=r"^[-\w\._\(\)]+$"
        ),
        "providerName": _SERIALIZER.url("provider_name", provider_name, "str"),
        "resourceType": _SERIALIZER.url("resource_type", resource_type, "str"),
        "resourceName": _SERIALIZER.url("resource_name", resource_name, "str"),
        "managementAssociationName": _SERIALIZER.url("management_association_name", management_association_name, "str"),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="DELETE", url=_url, params=_params, headers=_headers, **kwargs)


def build_get_request(
    resource_group_name: str,
    provider_name: str,
    resource_type: str,
    resource_name: str,
    management_association_name: str,
    subscription_id: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop(
        "api_version", _params.pop("api-version", "2015-11-01-preview")
    )  # type: Literal["2015-11-01-preview"]
    accept = _headers.pop("Accept", "application/json")

    # Construct URL
    _url = kwargs.pop(
        "template_url",
        "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}",
    )  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, "str"),
        "resourceGroupName": _SERIALIZER.url(
            "resource_group_name", resource_group_name, "str", max_length=90, min_length=1, pattern=r"^[-\w\._\(\)]+$"
        ),
        "providerName": _SERIALIZER.url("provider_name", provider_name, "str"),
        "resourceType": _SERIALIZER.url("resource_type", resource_type, "str"),
        "resourceName": _SERIALIZER.url("resource_name", resource_name, "str"),
        "managementAssociationName": _SERIALIZER.url("management_association_name", management_association_name, "str"),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params["api-version"] = _SERIALIZER.query("api_version", api_version, "str")

    # Construct headers
    _headers["Accept"] = _SERIALIZER.header("accept", accept, "str")

    return HttpRequest(method="GET", url=_url, params=_params, headers=_headers, **kwargs)


class ManagementAssociationsOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.operationsmanagement.OperationsManagementClient`'s
        :attr:`management_associations` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs):
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")

    @distributed_trace
    def list_by_subscription(self, **kwargs: Any) -> _models.ManagementAssociationPropertiesList:
        """Retrieves the ManagementAssociations list for the subscription.

        Retrieves the ManagementAssociations list.

        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ManagementAssociationPropertiesList or the result of cls(response)
        :rtype: ~azure.mgmt.operationsmanagement.models.ManagementAssociationPropertiesList
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
        )  # type: Literal["2015-11-01-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.ManagementAssociationPropertiesList]

        request = build_list_by_subscription_request(
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            template_url=self.list_by_subscription.metadata["url"],
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
            error = self._deserialize.failsafe_deserialize(_models.CodeMessageError, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("ManagementAssociationPropertiesList", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    list_by_subscription.metadata = {"url": "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations"}  # type: ignore

    @overload
    def create_or_update(
        self,
        resource_group_name: str,
        provider_name: str,
        resource_type: str,
        resource_name: str,
        management_association_name: str,
        parameters: _models.ManagementAssociation,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.ManagementAssociation:
        """Create/Update ManagementAssociation.

        Creates or updates the ManagementAssociation.

        :param resource_group_name: The name of the resource group to get. The name is case
         insensitive. Required.
        :type resource_group_name: str
        :param provider_name: Provider name for the parent resource. Required.
        :type provider_name: str
        :param resource_type: Resource type for the parent resource. Required.
        :type resource_type: str
        :param resource_name: Parent resource name. Required.
        :type resource_name: str
        :param management_association_name: User ManagementAssociation Name. Required.
        :type management_association_name: str
        :param parameters: The parameters required to create ManagementAssociation extension. Required.
        :type parameters: ~azure.mgmt.operationsmanagement.models.ManagementAssociation
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ManagementAssociation or the result of cls(response)
        :rtype: ~azure.mgmt.operationsmanagement.models.ManagementAssociation
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    def create_or_update(
        self,
        resource_group_name: str,
        provider_name: str,
        resource_type: str,
        resource_name: str,
        management_association_name: str,
        parameters: IO,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.ManagementAssociation:
        """Create/Update ManagementAssociation.

        Creates or updates the ManagementAssociation.

        :param resource_group_name: The name of the resource group to get. The name is case
         insensitive. Required.
        :type resource_group_name: str
        :param provider_name: Provider name for the parent resource. Required.
        :type provider_name: str
        :param resource_type: Resource type for the parent resource. Required.
        :type resource_type: str
        :param resource_name: Parent resource name. Required.
        :type resource_name: str
        :param management_association_name: User ManagementAssociation Name. Required.
        :type management_association_name: str
        :param parameters: The parameters required to create ManagementAssociation extension. Required.
        :type parameters: IO
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ManagementAssociation or the result of cls(response)
        :rtype: ~azure.mgmt.operationsmanagement.models.ManagementAssociation
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace
    def create_or_update(
        self,
        resource_group_name: str,
        provider_name: str,
        resource_type: str,
        resource_name: str,
        management_association_name: str,
        parameters: Union[_models.ManagementAssociation, IO],
        **kwargs: Any
    ) -> _models.ManagementAssociation:
        """Create/Update ManagementAssociation.

        Creates or updates the ManagementAssociation.

        :param resource_group_name: The name of the resource group to get. The name is case
         insensitive. Required.
        :type resource_group_name: str
        :param provider_name: Provider name for the parent resource. Required.
        :type provider_name: str
        :param resource_type: Resource type for the parent resource. Required.
        :type resource_type: str
        :param resource_name: Parent resource name. Required.
        :type resource_name: str
        :param management_association_name: User ManagementAssociation Name. Required.
        :type management_association_name: str
        :param parameters: The parameters required to create ManagementAssociation extension. Is either
         a model type or a IO type. Required.
        :type parameters: ~azure.mgmt.operationsmanagement.models.ManagementAssociation or IO
        :keyword content_type: Body Parameter content-type. Known values are: 'application/json'.
         Default value is None.
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ManagementAssociation or the result of cls(response)
        :rtype: ~azure.mgmt.operationsmanagement.models.ManagementAssociation
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

        api_version = kwargs.pop(
            "api_version", _params.pop("api-version", self._config.api_version)
        )  # type: Literal["2015-11-01-preview"]
        content_type = kwargs.pop("content_type", _headers.pop("Content-Type", None))  # type: Optional[str]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.ManagementAssociation]

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(parameters, (IO, bytes)):
            _content = parameters
        else:
            _json = self._serialize.body(parameters, "ManagementAssociation")

        request = build_create_or_update_request(
            resource_group_name=resource_group_name,
            provider_name=provider_name,
            resource_type=resource_type,
            resource_name=resource_name,
            management_association_name=management_association_name,
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

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.CodeMessageError, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("ManagementAssociation", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    create_or_update.metadata = {"url": "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"}  # type: ignore

    @distributed_trace
    def delete(  # pylint: disable=inconsistent-return-statements
        self,
        resource_group_name: str,
        provider_name: str,
        resource_type: str,
        resource_name: str,
        management_association_name: str,
        **kwargs: Any
    ) -> None:
        """Deletes the ManagementAssociation.

        Deletes the ManagementAssociation in the subscription.

        :param resource_group_name: The name of the resource group to get. The name is case
         insensitive. Required.
        :type resource_group_name: str
        :param provider_name: Provider name for the parent resource. Required.
        :type provider_name: str
        :param resource_type: Resource type for the parent resource. Required.
        :type resource_type: str
        :param resource_name: Parent resource name. Required.
        :type resource_name: str
        :param management_association_name: User ManagementAssociation Name. Required.
        :type management_association_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: None or the result of cls(response)
        :rtype: None
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
        )  # type: Literal["2015-11-01-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[None]

        request = build_delete_request(
            resource_group_name=resource_group_name,
            provider_name=provider_name,
            resource_type=resource_type,
            resource_name=resource_name,
            management_association_name=management_association_name,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            template_url=self.delete.metadata["url"],
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
            error = self._deserialize.failsafe_deserialize(_models.CodeMessageError, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        if cls:
            return cls(pipeline_response, None, {})

    delete.metadata = {"url": "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"}  # type: ignore

    @distributed_trace
    def get(
        self,
        resource_group_name: str,
        provider_name: str,
        resource_type: str,
        resource_name: str,
        management_association_name: str,
        **kwargs: Any
    ) -> _models.ManagementAssociation:
        """Retrieve ManagementAssociation.

        Retrieves the user ManagementAssociation.

        :param resource_group_name: The name of the resource group to get. The name is case
         insensitive. Required.
        :type resource_group_name: str
        :param provider_name: Provider name for the parent resource. Required.
        :type provider_name: str
        :param resource_type: Resource type for the parent resource. Required.
        :type resource_type: str
        :param resource_name: Parent resource name. Required.
        :type resource_name: str
        :param management_association_name: User ManagementAssociation Name. Required.
        :type management_association_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ManagementAssociation or the result of cls(response)
        :rtype: ~azure.mgmt.operationsmanagement.models.ManagementAssociation
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
        )  # type: Literal["2015-11-01-preview"]
        cls = kwargs.pop("cls", None)  # type: ClsType[_models.ManagementAssociation]

        request = build_get_request(
            resource_group_name=resource_group_name,
            provider_name=provider_name,
            resource_type=resource_type,
            resource_name=resource_name,
            management_association_name=management_association_name,
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
            error = self._deserialize.failsafe_deserialize(_models.CodeMessageError, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("ManagementAssociation", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get.metadata = {"url": "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"}  # type: ignore
