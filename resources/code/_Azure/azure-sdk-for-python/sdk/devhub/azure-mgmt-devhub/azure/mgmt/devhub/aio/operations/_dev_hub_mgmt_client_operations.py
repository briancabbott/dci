# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
from io import IOBase
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
from azure.core.pipeline.transport import AsyncHttpResponse
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator_async import distributed_trace_async
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat

from ... import models as _models
from ..._vendor import _convert_request
from ...operations._dev_hub_mgmt_client_operations import (
    build_generate_preview_artifacts_request,
    build_git_hub_o_auth_callback_request,
    build_git_hub_o_auth_request,
    build_list_git_hub_o_auth_request,
)
from .._vendor import DevHubMgmtClientMixinABC

T = TypeVar("T")
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, AsyncHttpResponse], T, Dict[str, Any]], Any]]


class DevHubMgmtClientOperationsMixin(DevHubMgmtClientMixinABC):
    @overload
    async def git_hub_o_auth(
        self,
        location: str,
        parameters: Optional[_models.GitHubOAuthCallRequest] = None,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> _models.GitHubOAuthInfoResponse:
        """Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Default value is None.
        :type parameters: ~azure.mgmt.devhub.models.GitHubOAuthCallRequest
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: GitHubOAuthInfoResponse or the result of cls(response)
        :rtype: ~azure.mgmt.devhub.models.GitHubOAuthInfoResponse
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    async def git_hub_o_auth(
        self, location: str, parameters: Optional[IO] = None, *, content_type: str = "application/json", **kwargs: Any
    ) -> _models.GitHubOAuthInfoResponse:
        """Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Default value is None.
        :type parameters: IO
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: GitHubOAuthInfoResponse or the result of cls(response)
        :rtype: ~azure.mgmt.devhub.models.GitHubOAuthInfoResponse
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace_async
    async def git_hub_o_auth(
        self, location: str, parameters: Optional[Union[_models.GitHubOAuthCallRequest, IO]] = None, **kwargs: Any
    ) -> _models.GitHubOAuthInfoResponse:
        """Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        Gets GitHubOAuth info used to authenticate users with the Developer Hub GitHub App.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Is either a GitHubOAuthCallRequest type or a IO type. Default value is None.
        :type parameters: ~azure.mgmt.devhub.models.GitHubOAuthCallRequest or IO
        :keyword content_type: Body Parameter content-type. Known values are: 'application/json'.
         Default value is None.
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: GitHubOAuthInfoResponse or the result of cls(response)
        :rtype: ~azure.mgmt.devhub.models.GitHubOAuthInfoResponse
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

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._config.api_version))
        content_type: Optional[str] = kwargs.pop("content_type", _headers.pop("Content-Type", None))
        cls: ClsType[_models.GitHubOAuthInfoResponse] = kwargs.pop("cls", None)

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(parameters, (IOBase, bytes)):
            _content = parameters
        else:
            if parameters is not None:
                _json = self._serialize.body(parameters, "GitHubOAuthCallRequest")
            else:
                _json = None

        request = build_git_hub_o_auth_request(
            location=location,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            content_type=content_type,
            json=_json,
            content=_content,
            template_url=self.git_hub_o_auth.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("GitHubOAuthInfoResponse", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    git_hub_o_auth.metadata = {
        "url": "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth/default/getGitHubOAuthInfo"
    }

    @distributed_trace_async
    async def git_hub_o_auth_callback(
        self, location: str, code: str, state: str, **kwargs: Any
    ) -> _models.GitHubOAuthResponse:
        """Callback URL to hit once authenticated with GitHub App to have the service store the OAuth
        token.

        Callback URL to hit once authenticated with GitHub App to have the service store the OAuth
        token.

        :param location: The name of Azure region. Required.
        :type location: str
        :param code: The code response from authenticating the GitHub App. Required.
        :type code: str
        :param state: The state response from authenticating the GitHub App. Required.
        :type state: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: GitHubOAuthResponse or the result of cls(response)
        :rtype: ~azure.mgmt.devhub.models.GitHubOAuthResponse
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

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._config.api_version))
        cls: ClsType[_models.GitHubOAuthResponse] = kwargs.pop("cls", None)

        request = build_git_hub_o_auth_callback_request(
            location=location,
            subscription_id=self._config.subscription_id,
            code=code,
            state=state,
            api_version=api_version,
            template_url=self.git_hub_o_auth_callback.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("GitHubOAuthResponse", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    git_hub_o_auth_callback.metadata = {
        "url": "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth/default"
    }

    @distributed_trace_async
    async def list_git_hub_o_auth(self, location: str, **kwargs: Any) -> _models.GitHubOAuthListResponse:
        """Callback URL to hit once authenticated with GitHub App to have the service store the OAuth
        token.

        Callback URL to hit once authenticated with GitHub App to have the service store the OAuth
        token.

        :param location: The name of Azure region. Required.
        :type location: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: GitHubOAuthListResponse or the result of cls(response)
        :rtype: ~azure.mgmt.devhub.models.GitHubOAuthListResponse
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

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._config.api_version))
        cls: ClsType[_models.GitHubOAuthListResponse] = kwargs.pop("cls", None)

        request = build_list_git_hub_o_auth_request(
            location=location,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            template_url=self.list_git_hub_o_auth.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("GitHubOAuthListResponse", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    list_git_hub_o_auth.metadata = {
        "url": "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/githuboauth"
    }

    @overload
    async def generate_preview_artifacts(
        self,
        location: str,
        parameters: _models.ArtifactGenerationProperties,
        *,
        content_type: str = "application/json",
        **kwargs: Any
    ) -> Dict[str, str]:
        """Generate preview dockerfile and manifests.

        Generate preview dockerfile and manifests.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Required.
        :type parameters: ~azure.mgmt.devhub.models.ArtifactGenerationProperties
        :keyword content_type: Body Parameter content-type. Content type parameter for JSON body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: dict mapping str to str or the result of cls(response)
        :rtype: dict[str, str]
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @overload
    async def generate_preview_artifacts(
        self, location: str, parameters: IO, *, content_type: str = "application/json", **kwargs: Any
    ) -> Dict[str, str]:
        """Generate preview dockerfile and manifests.

        Generate preview dockerfile and manifests.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Required.
        :type parameters: IO
        :keyword content_type: Body Parameter content-type. Content type parameter for binary body.
         Default value is "application/json".
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: dict mapping str to str or the result of cls(response)
        :rtype: dict[str, str]
        :raises ~azure.core.exceptions.HttpResponseError:
        """

    @distributed_trace_async
    async def generate_preview_artifacts(
        self, location: str, parameters: Union[_models.ArtifactGenerationProperties, IO], **kwargs: Any
    ) -> Dict[str, str]:
        """Generate preview dockerfile and manifests.

        Generate preview dockerfile and manifests.

        :param location: The name of Azure region. Required.
        :type location: str
        :param parameters: Is either a ArtifactGenerationProperties type or a IO type. Required.
        :type parameters: ~azure.mgmt.devhub.models.ArtifactGenerationProperties or IO
        :keyword content_type: Body Parameter content-type. Known values are: 'application/json'.
         Default value is None.
        :paramtype content_type: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: dict mapping str to str or the result of cls(response)
        :rtype: dict[str, str]
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

        api_version: str = kwargs.pop("api_version", _params.pop("api-version", self._config.api_version))
        content_type: Optional[str] = kwargs.pop("content_type", _headers.pop("Content-Type", None))
        cls: ClsType[Dict[str, str]] = kwargs.pop("cls", None)

        content_type = content_type or "application/json"
        _json = None
        _content = None
        if isinstance(parameters, (IOBase, bytes)):
            _content = parameters
        else:
            _json = self._serialize.body(parameters, "ArtifactGenerationProperties")

        request = build_generate_preview_artifacts_request(
            location=location,
            subscription_id=self._config.subscription_id,
            api_version=api_version,
            content_type=content_type,
            json=_json,
            content=_content,
            template_url=self.generate_preview_artifacts.metadata["url"],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)

        _stream = False
        pipeline_response: PipelineResponse = await self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request, stream=_stream, **kwargs
        )

        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize("{str}", pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    generate_preview_artifacts.metadata = {
        "url": "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/locations/{location}/generatePreviewArtifacts"
    }
