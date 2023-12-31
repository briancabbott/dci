# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
from typing import TYPE_CHECKING
import warnings

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError, ResourceExistsError, ResourceNotFoundError, map_error
from azure.core.paging import ItemPaged
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import HttpRequest, HttpResponse
from azure.mgmt.core.exceptions import ARMErrorFormat

from .. import models as _models

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from typing import Any, Callable, Dict, Generic, Iterable, Optional, TypeVar

    T = TypeVar('T')
    ClsType = Optional[Callable[[PipelineResponse[HttpRequest, HttpResponse], T, Dict[str, Any]], Any]]

class KeysOperations(object):
    """KeysOperations operations.

    You should not instantiate this class directly. Instead, you should create a Client instance that
    instantiates it for you and attaches it as an attribute.

    :ivar models: Alias to model classes used in this operation group.
    :type models: ~azure.mgmt.keyvault.v2021_12_01_preview.models
    :param client: Client for service requests.
    :param config: Configuration of service client.
    :param serializer: An object model serializer.
    :param deserializer: An object model deserializer.
    """

    models = _models

    def __init__(self, client, config, serializer, deserializer):
        self._client = client
        self._serialize = serializer
        self._deserialize = deserializer
        self._config = config

    def create_if_not_exist(
        self,
        resource_group_name,  # type: str
        vault_name,  # type: str
        key_name,  # type: str
        parameters,  # type: "_models.KeyCreateParameters"
        **kwargs  # type: Any
    ):
        # type: (...) -> "_models.Key"
        """Creates the first version of a new key if it does not exist. If it already exists, then the
        existing key is returned without any write operations being performed. This API does not create
        subsequent versions, and does not update existing keys.

        :param resource_group_name: The name of the resource group which contains the specified key
         vault.
        :type resource_group_name: str
        :param vault_name: The name of the key vault which contains the key to be created.
        :type vault_name: str
        :param key_name: The name of the key to be created.
        :type key_name: str
        :param parameters: The parameters used to create the specified key.
        :type parameters: ~azure.mgmt.keyvault.v2021_12_01_preview.models.KeyCreateParameters
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Key, or the result of cls(response)
        :rtype: ~azure.mgmt.keyvault.v2021_12_01_preview.models.Key
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Key"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))
        api_version = "2021-12-01-preview"
        content_type = kwargs.pop("content_type", "application/json")
        accept = "application/json"

        # Construct URL
        url = self.create_if_not_exist.metadata['url']  # type: ignore
        path_format_arguments = {
            'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str'),
            'resourceGroupName': self._serialize.url("resource_group_name", resource_group_name, 'str'),
            'vaultName': self._serialize.url("vault_name", vault_name, 'str', pattern=r'^[a-zA-Z0-9-]{3,24}$'),
            'keyName': self._serialize.url("key_name", key_name, 'str', pattern=r'^[a-zA-Z0-9-]{1,127}$'),
        }
        url = self._client.format_url(url, **path_format_arguments)

        # Construct parameters
        query_parameters = {}  # type: Dict[str, Any]
        query_parameters['api-version'] = self._serialize.query("api_version", api_version, 'str')

        # Construct headers
        header_parameters = {}  # type: Dict[str, Any]
        header_parameters['Content-Type'] = self._serialize.header("content_type", content_type, 'str')
        header_parameters['Accept'] = self._serialize.header("accept", accept, 'str')

        body_content_kwargs = {}  # type: Dict[str, Any]
        body_content = self._serialize.body(parameters, 'KeyCreateParameters')
        body_content_kwargs['content'] = body_content
        request = self._client.put(url, query_parameters, header_parameters, **body_content_kwargs)
        pipeline_response = self._client._pipeline.run(request, stream=False, **kwargs)
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Key', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized
    create_if_not_exist.metadata = {'url': '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}'}  # type: ignore

    def get(
        self,
        resource_group_name,  # type: str
        vault_name,  # type: str
        key_name,  # type: str
        **kwargs  # type: Any
    ):
        # type: (...) -> "_models.Key"
        """Gets the current version of the specified key from the specified key vault.

        :param resource_group_name: The name of the resource group which contains the specified key
         vault.
        :type resource_group_name: str
        :param vault_name: The name of the vault which contains the key to be retrieved.
        :type vault_name: str
        :param key_name: The name of the key to be retrieved.
        :type key_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Key, or the result of cls(response)
        :rtype: ~azure.mgmt.keyvault.v2021_12_01_preview.models.Key
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Key"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))
        api_version = "2021-12-01-preview"
        accept = "application/json"

        # Construct URL
        url = self.get.metadata['url']  # type: ignore
        path_format_arguments = {
            'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str'),
            'resourceGroupName': self._serialize.url("resource_group_name", resource_group_name, 'str'),
            'vaultName': self._serialize.url("vault_name", vault_name, 'str', pattern=r'^[a-zA-Z0-9-]{3,24}$'),
            'keyName': self._serialize.url("key_name", key_name, 'str', pattern=r'^[a-zA-Z0-9-]{1,127}$'),
        }
        url = self._client.format_url(url, **path_format_arguments)

        # Construct parameters
        query_parameters = {}  # type: Dict[str, Any]
        query_parameters['api-version'] = self._serialize.query("api_version", api_version, 'str')

        # Construct headers
        header_parameters = {}  # type: Dict[str, Any]
        header_parameters['Accept'] = self._serialize.header("accept", accept, 'str')

        request = self._client.get(url, query_parameters, header_parameters)
        pipeline_response = self._client._pipeline.run(request, stream=False, **kwargs)
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Key', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized
    get.metadata = {'url': '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}'}  # type: ignore

    def list(
        self,
        resource_group_name,  # type: str
        vault_name,  # type: str
        **kwargs  # type: Any
    ):
        # type: (...) -> Iterable["_models.KeyListResult"]
        """Lists the keys in the specified key vault.

        :param resource_group_name: The name of the resource group which contains the specified key
         vault.
        :type resource_group_name: str
        :param vault_name: The name of the vault which contains the keys to be retrieved.
        :type vault_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: An iterator like instance of either KeyListResult or the result of cls(response)
        :rtype: ~azure.core.paging.ItemPaged[~azure.mgmt.keyvault.v2021_12_01_preview.models.KeyListResult]
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.KeyListResult"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))
        api_version = "2021-12-01-preview"
        accept = "application/json"

        def prepare_request(next_link=None):
            # Construct headers
            header_parameters = {}  # type: Dict[str, Any]
            header_parameters['Accept'] = self._serialize.header("accept", accept, 'str')

            if not next_link:
                # Construct URL
                url = self.list.metadata['url']  # type: ignore
                path_format_arguments = {
                    'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str'),
                    'resourceGroupName': self._serialize.url("resource_group_name", resource_group_name, 'str'),
                    'vaultName': self._serialize.url("vault_name", vault_name, 'str', pattern=r'^[a-zA-Z0-9-]{3,24}$'),
                }
                url = self._client.format_url(url, **path_format_arguments)
                # Construct parameters
                query_parameters = {}  # type: Dict[str, Any]
                query_parameters['api-version'] = self._serialize.query("api_version", api_version, 'str')

                request = self._client.get(url, query_parameters, header_parameters)
            else:
                url = next_link
                query_parameters = {}  # type: Dict[str, Any]
                request = self._client.get(url, query_parameters, header_parameters)
            return request

        def extract_data(pipeline_response):
            deserialized = self._deserialize('KeyListResult', pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)
            return deserialized.next_link or None, iter(list_of_elem)

        def get_next(next_link=None):
            request = prepare_request(next_link)

            pipeline_response = self._client._pipeline.run(request, stream=False, **kwargs)
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return ItemPaged(
            get_next, extract_data
        )
    list.metadata = {'url': '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys'}  # type: ignore

    def get_version(
        self,
        resource_group_name,  # type: str
        vault_name,  # type: str
        key_name,  # type: str
        key_version,  # type: str
        **kwargs  # type: Any
    ):
        # type: (...) -> "_models.Key"
        """Gets the specified version of the specified key in the specified key vault.

        :param resource_group_name: The name of the resource group which contains the specified key
         vault.
        :type resource_group_name: str
        :param vault_name: The name of the vault which contains the key version to be retrieved.
        :type vault_name: str
        :param key_name: The name of the key version to be retrieved.
        :type key_name: str
        :param key_version: The version of the key to be retrieved.
        :type key_version: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: Key, or the result of cls(response)
        :rtype: ~azure.mgmt.keyvault.v2021_12_01_preview.models.Key
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.Key"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))
        api_version = "2021-12-01-preview"
        accept = "application/json"

        # Construct URL
        url = self.get_version.metadata['url']  # type: ignore
        path_format_arguments = {
            'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str'),
            'resourceGroupName': self._serialize.url("resource_group_name", resource_group_name, 'str'),
            'vaultName': self._serialize.url("vault_name", vault_name, 'str', pattern=r'^[a-zA-Z0-9-]{3,24}$'),
            'keyName': self._serialize.url("key_name", key_name, 'str', pattern=r'^[a-zA-Z0-9-]{1,127}$'),
            'keyVersion': self._serialize.url("key_version", key_version, 'str', pattern=r'^[a-fA-F0-9]{32}$'),
        }
        url = self._client.format_url(url, **path_format_arguments)

        # Construct parameters
        query_parameters = {}  # type: Dict[str, Any]
        query_parameters['api-version'] = self._serialize.query("api_version", api_version, 'str')

        # Construct headers
        header_parameters = {}  # type: Dict[str, Any]
        header_parameters['Accept'] = self._serialize.header("accept", accept, 'str')

        request = self._client.get(url, query_parameters, header_parameters)
        pipeline_response = self._client._pipeline.run(request, stream=False, **kwargs)
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('Key', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized
    get_version.metadata = {'url': '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}/versions/{keyVersion}'}  # type: ignore

    def list_versions(
        self,
        resource_group_name,  # type: str
        vault_name,  # type: str
        key_name,  # type: str
        **kwargs  # type: Any
    ):
        # type: (...) -> Iterable["_models.KeyListResult"]
        """Lists the versions of the specified key in the specified key vault.

        :param resource_group_name: The name of the resource group which contains the specified key
         vault.
        :type resource_group_name: str
        :param vault_name: The name of the vault which contains the key versions to be retrieved.
        :type vault_name: str
        :param key_name: The name of the key versions to be retrieved.
        :type key_name: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: An iterator like instance of either KeyListResult or the result of cls(response)
        :rtype: ~azure.core.paging.ItemPaged[~azure.mgmt.keyvault.v2021_12_01_preview.models.KeyListResult]
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        cls = kwargs.pop('cls', None)  # type: ClsType["_models.KeyListResult"]
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}))
        api_version = "2021-12-01-preview"
        accept = "application/json"

        def prepare_request(next_link=None):
            # Construct headers
            header_parameters = {}  # type: Dict[str, Any]
            header_parameters['Accept'] = self._serialize.header("accept", accept, 'str')

            if not next_link:
                # Construct URL
                url = self.list_versions.metadata['url']  # type: ignore
                path_format_arguments = {
                    'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str'),
                    'resourceGroupName': self._serialize.url("resource_group_name", resource_group_name, 'str'),
                    'vaultName': self._serialize.url("vault_name", vault_name, 'str', pattern=r'^[a-zA-Z0-9-]{3,24}$'),
                    'keyName': self._serialize.url("key_name", key_name, 'str', pattern=r'^[a-zA-Z0-9-]{1,127}$'),
                }
                url = self._client.format_url(url, **path_format_arguments)
                # Construct parameters
                query_parameters = {}  # type: Dict[str, Any]
                query_parameters['api-version'] = self._serialize.query("api_version", api_version, 'str')

                request = self._client.get(url, query_parameters, header_parameters)
            else:
                url = next_link
                query_parameters = {}  # type: Dict[str, Any]
                request = self._client.get(url, query_parameters, header_parameters)
            return request

        def extract_data(pipeline_response):
            deserialized = self._deserialize('KeyListResult', pipeline_response)
            list_of_elem = deserialized.value
            if cls:
                list_of_elem = cls(list_of_elem)
            return deserialized.next_link or None, iter(list_of_elem)

        def get_next(next_link=None):
            request = prepare_request(next_link)

            pipeline_response = self._client._pipeline.run(request, stream=False, **kwargs)
            response = pipeline_response.http_response

            if response.status_code not in [200]:
                map_error(status_code=response.status_code, response=response, error_map=error_map)
                raise HttpResponseError(response=response, error_format=ARMErrorFormat)

            return pipeline_response

        return ItemPaged(
            get_next, extract_data
        )
    list_versions.metadata = {'url': '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys/{keyName}/versions'}  # type: ignore
