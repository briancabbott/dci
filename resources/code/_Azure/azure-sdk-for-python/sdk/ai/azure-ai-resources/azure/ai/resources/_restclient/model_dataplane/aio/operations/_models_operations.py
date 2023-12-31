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
from ...operations._models_operations import build_batch_get_resolved_uris_request, build_batch_query_request, build_create_unregistered_input_model_request, build_create_unregistered_output_model_request, build_delete_request, build_deployment_settings_request, build_list_query_post_request, build_list_request, build_patch_request, build_query_by_id_request, build_register_request
T = TypeVar('T')
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, AsyncHttpResponse], T, Dict[str, Any]], Any]]

class ModelsOperations:
    """ModelsOperations async operations.

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
    async def register(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: "_models.Model",
        auto_version: Optional[bool] = True,
        **kwargs: Any
    ) -> "_models.Model":
        """register.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.Model
        :param auto_version:
        :type auto_version: bool
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Model, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Model
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Model"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        _json = self._serialize.body(body, 'Model')

        request = build_register_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            auto_version=auto_version,
            template_url=self.register.metadata['url'],
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

        deserialized = self._deserialize('Model', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    register.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models"}  # type: ignore


    @distributed_trace_async
    async def list(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        name: Optional[str] = None,
        tag: Optional[str] = None,
        version: Optional[str] = None,
        framework: Optional[str] = None,
        description: Optional[str] = None,
        count: Optional[int] = None,
        offset: Optional[int] = None,
        skip_token: Optional[str] = None,
        tags: Optional[str] = None,
        properties: Optional[str] = None,
        run_id: Optional[str] = None,
        dataset_id: Optional[str] = None,
        order_by: Optional[str] = None,
        latest_version_only: Optional[bool] = False,
        feed: Optional[str] = None,
        list_view_type: Optional[Union[str, "_models.ListViewType"]] = None,
        **kwargs: Any
    ) -> "_models.ModelPagedResponse":
        """list.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param name:
        :type name: str
        :param tag:
        :type tag: str
        :param version:
        :type version: str
        :param framework:
        :type framework: str
        :param description:
        :type description: str
        :param count:
        :type count: int
        :param offset:
        :type offset: int
        :param skip_token:
        :type skip_token: str
        :param tags:
        :type tags: str
        :param properties:
        :type properties: str
        :param run_id:
        :type run_id: str
        :param dataset_id:
        :type dataset_id: str
        :param order_by:
        :type order_by: str
        :param latest_version_only:
        :type latest_version_only: bool
        :param feed:
        :type feed: str
        :param list_view_type:
        :type list_view_type: str or ~azure.mgmt.machinelearningservices.models.ListViewType
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ModelPagedResponse, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.ModelPagedResponse
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.ModelPagedResponse"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        
        request = build_list_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            name=name,
            tag=tag,
            version=version,
            framework=framework,
            description=description,
            count=count,
            offset=offset,
            skip_token=skip_token,
            tags=tags,
            properties=properties,
            run_id=run_id,
            dataset_id=dataset_id,
            order_by=order_by,
            latest_version_only=latest_version_only,
            feed=feed,
            list_view_type=list_view_type,
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

        deserialized = self._deserialize('ModelPagedResponse', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    list.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models"}  # type: ignore


    @distributed_trace_async
    async def create_unregistered_input_model(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: "_models.CreateUnregisteredInputModelDto",
        **kwargs: Any
    ) -> "_models.Model":
        """create_unregistered_input_model.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.CreateUnregisteredInputModelDto
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Model, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Model
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Model"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        _json = self._serialize.body(body, 'CreateUnregisteredInputModelDto')

        request = build_create_unregistered_input_model_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.create_unregistered_input_model.metadata['url'],
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

        deserialized = self._deserialize('Model', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    create_unregistered_input_model.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/createUnregisteredInput"}  # type: ignore


    @distributed_trace_async
    async def create_unregistered_output_model(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: "_models.CreateUnregisteredOutputModelDto",
        **kwargs: Any
    ) -> "_models.Model":
        """create_unregistered_output_model.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.CreateUnregisteredOutputModelDto
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Model, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Model
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Model"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        _json = self._serialize.body(body, 'CreateUnregisteredOutputModelDto')

        request = build_create_unregistered_output_model_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.create_unregistered_output_model.metadata['url'],
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

        deserialized = self._deserialize('Model', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    create_unregistered_output_model.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/createUnregisteredOutput"}  # type: ignore


    @distributed_trace_async
    async def batch_get_resolved_uris(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: Optional["_models.BatchGetResolvedUrisDto"] = None,
        **kwargs: Any
    ) -> "_models.BatchModelPathResponseDto":
        """batch_get_resolved_uris.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.BatchGetResolvedUrisDto
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: BatchModelPathResponseDto, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.BatchModelPathResponseDto
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.BatchModelPathResponseDto"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        if body is not None:
            _json = self._serialize.body(body, 'BatchGetResolvedUrisDto')
        else:
            _json = None

        request = build_batch_get_resolved_uris_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.batch_get_resolved_uris.metadata['url'],
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

        deserialized = self._deserialize('BatchModelPathResponseDto', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    batch_get_resolved_uris.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/batchGetResolvedUris"}  # type: ignore


    @distributed_trace_async
    async def query_by_id(
        self,
        id: str,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        include_deployment_settings: Optional[bool] = False,
        **kwargs: Any
    ) -> "_models.Model":
        """query_by_id.

        :param id:
        :type id: str
        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param include_deployment_settings:
        :type include_deployment_settings: bool
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Model, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Model
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Model"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        
        request = build_query_by_id_request(
            id=id,
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            include_deployment_settings=include_deployment_settings,
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

        deserialized = self._deserialize('Model', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    query_by_id.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/{id}"}  # type: ignore


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

    delete.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/{id}"}  # type: ignore


    @distributed_trace_async
    async def patch(
        self,
        id: str,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: List["_models.Operation"],
        **kwargs: Any
    ) -> "_models.Model":
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
        :return: Model, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.Model
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Model"]
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

        deserialized = self._deserialize('Model', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    patch.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/{id}"}  # type: ignore


    @distributed_trace_async
    async def list_query_post(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: Optional["_models.ListModelsRequest"] = None,
        **kwargs: Any
    ) -> "_models.ModelListModelsRequestPagedResponse":
        """list_query_post.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.ListModelsRequest
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ModelListModelsRequestPagedResponse, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.ModelListModelsRequestPagedResponse
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.ModelListModelsRequestPagedResponse"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        if body is not None:
            _json = self._serialize.body(body, 'ListModelsRequest')
        else:
            _json = None

        request = build_list_query_post_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.list_query_post.metadata['url'],
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

        deserialized = self._deserialize('ModelListModelsRequestPagedResponse', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    list_query_post.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/list"}  # type: ignore


    @distributed_trace_async
    async def batch_query(
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: Optional["_models.ModelBatchDto"] = None,
        **kwargs: Any
    ) -> "_models.ModelBatchResponseDto":
        """batch_query.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.ModelBatchDto
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: ModelBatchResponseDto, or the result of cls(response)
        :rtype: ~azure.mgmt.machinelearningservices.models.ModelBatchResponseDto
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.ModelBatchResponseDto"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        if body is not None:
            _json = self._serialize.body(body, 'ModelBatchDto')
        else:
            _json = None

        request = build_batch_query_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.batch_query.metadata['url'],
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

        deserialized = self._deserialize('ModelBatchResponseDto', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    batch_query.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/querybatch"}  # type: ignore


    @distributed_trace_async
    async def deployment_settings(  # pylint: disable=inconsistent-return-statements
        self,
        subscription_id: str,
        resource_group_name: str,
        workspace_name: str,
        body: Optional["_models.ModelSettingsIdentifiers"] = None,
        **kwargs: Any
    ) -> None:
        """deployment_settings.

        :param subscription_id:
        :type subscription_id: str
        :param resource_group_name:
        :type resource_group_name: str
        :param workspace_name:
        :type workspace_name: str
        :param body:
        :type body: ~azure.mgmt.machinelearningservices.models.ModelSettingsIdentifiers
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

        content_type = kwargs.pop('content_type', "application/json-patch+json")  # type: Optional[str]

        if body is not None:
            _json = self._serialize.body(body, 'ModelSettingsIdentifiers')
        else:
            _json = None

        request = build_deployment_settings_request(
            subscription_id=subscription_id,
            resource_group_name=resource_group_name,
            workspace_name=workspace_name,
            content_type=content_type,
            json=_json,
            template_url=self.deployment_settings.metadata['url'],
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

        if cls:
            return cls(pipeline_response, None, {})

    deployment_settings.metadata = {'url': "/modelregistry/v1.0/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/models/deploymentSettings"}  # type: ignore

