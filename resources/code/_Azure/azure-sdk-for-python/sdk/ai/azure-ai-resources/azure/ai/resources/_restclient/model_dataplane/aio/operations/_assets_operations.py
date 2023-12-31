# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
from typing import Any, Callable, Dict, List, Optional, TypeVar, Union

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError, ResourceExistsError, ResourceNotFoundError, map_error
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import AsyncHttpResponse
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.mgmt.core.exceptions import ARMErrorFormat

from ... import models as _models
from ..._vendor import _convert_request
from ...operations._assets_operations import build_create_request, build_delete_request, build_list_request, build_patch_request, build_query_by_id_request
T = TypeVar('T')
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, AsyncHttpResponse], T, Dict[str, Any]], Any]]

class AssetsOperations:
    """AssetsOperations async operations.

    You should not instantiate this class directly. Instead, you should create a Client instance that
    instantiates it for you and attaches it as an attribute.

    :ivar models: Alias to model classes used in this operation group.
    :type models: ~azure.mgmt.machinelearningservices.models
    :param client: Client for service requests.
    :param config: Configuration of service client.
    :param serializer: An object model serializer.
    :param deserializer: An object model deserializer.
    """

    models = _models

    def __init__(self, client, config, serializer, deserializer) -> None:
        self._client = client
        self._serialize = serializer
        self._deserialize = deserializer
        self._config = config

    @distributed_trace_async
    async def create(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: Optional["_models.Asset"] = None,
        **kwargs: Any
    ) -> "_models.Asset":
        """create.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.Asset
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Asset, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Asset
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Asset"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        if body is not None:
            _json = self._serialize.body(body, 'Asset')
        else:
            _json = None

        request = build_create_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.create.metadata['url'],
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        pipeline_response = await self._client._pipeline.run(  # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Asset', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    create.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/assets"}  # type: ignore


    @distributed_trace_async
    async def list(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        run_id: Optional[str] = None,
        project_id: Optional[str] = None,
        name: Optional[str] = None,
        tag: Optional[str] = None,
        count: Optional[int] = None,
        skip_token: Optional[str] = None,
        tags: Optional[str] = None,
        properties: Optional[str] = None,
        type: Optional[str] = None,
        orderby: Optional[Union[str, "_models.OrderString"]] = None,
        **kwargs: Any
    ) -> "_models.AssetPaginatedResult":
        """list.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param run_id:
        :type run_id: str
        :param project_id:
        :type project_id: str
        :param name:
        :type name: str
        :param tag:
        :type tag: str
        :param count:
        :type count: int
        :param skip_token:
        :type skip_token: str
        :param tags:
        :type tags: str
        :param properties:
        :type properties: str
        :param type:
        :type type: str
        :param orderby:
        :type orderby: str or ~azure.mgmt.machinelearningservices.models.OrderString
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: AssetPaginatedResult, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.AssetPaginatedResult
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.AssetPaginatedResult"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        
        request = build_list_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            run_id=run_id,
            project_id=project_id,
            name=name,
            tag=tag,
            count=count,
            skip_token=skip_token,
            tags=tags,
            properties=properties,
            type=type,
            orderby=orderby,
            template_url=self.list.metadata['url'],
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        pipeline_response = await self._client._pipeline.run(  # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('AssetPaginatedResult', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    list.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/assets"}  # type: ignore


    @distributed_trace_async
    async def patch(
        self,
        id: str,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: List["_models.Operation"],
        **kwargs: Any
    ) -> "_models.Asset":
        """patch.

        :param id:
        :type id: str
        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: list[~azure.mgmt.machinelearningservices.models.Operation]
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Asset, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Asset
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Asset"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        _json = self._serialize.body(body, '[Operation]')

        request = build_patch_request(
            id=id,
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.patch.metadata['url'],
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        pipeline_response = await self._client._pipeline.run(  # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Asset', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    patch.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/assets/{id}"}  # type: ignore


    @distributed_trace_async
    async def delete(  # pylint: disable=inconsistent-return-statements
        self,
        id: str,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        **kwargs: Any
    ) -> None:
        """delete.

        :param id:
        :type id: str
        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: None, or the result of cls(response)
        :rtype: None
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType[None]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        
        request = build_delete_request(
            id=id,
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            template_url=self.delete.metadata['url'],
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        pipeline_response = await self._client._pipeline.run(  # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200, 204]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        if cls:
            return cls(pipeline_response, None, {})

    delete.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/assets/{id}"}  # type: ignore


    @distributed_trace_async
    async def query_by_id(
        self,
        id: str,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        **kwargs: Any
    ) -> "_models.Asset":
        """query_by_id.

        :param id:
        :type id: str
        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Asset, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Asset
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Asset"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        
        request = build_query_by_id_request(
            id=id,
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            template_url=self.query_by_id.metadata['url'],
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        pipeline_response = await self._client._pipeline.run(  # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Asset', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    query_by_id.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/assets/{id}"}  # type: ignore

