# pylint: disable=too-many-lines
# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------
from typing import Any, Callable, Dict, IO, Optional, TypeVar, Union, cast

from msrest import Serializer

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError, ResourceExistsError, ResourceNotFoundError, map_error
from azure.core.pipeline import PipelineResponse
from azure.core.pipeline.transport import HttpResponse
from azure.core.polling import LROPoller, NoPolling, PollingMethod
from azure.core.rest import HttpRequest
from azure.core.tracing.decorator import distributed_trace
from azure.core.utils import case_insensitive_dict
from azure.mgmt.core.exceptions import ARMErrorFormat
from azure.mgmt.core.polling.arm_polling import ARMPolling

from .. import models as _models
from .._vendor import MixinABC, _convert_request, _format_url_section
T = TypeVar('T')
ClsType = Optional[Callable[[PipelineResponse[HttpRequest, HttpResponse], T, Dict[str, Any]], Any]]

_SERIALIZER = Serializer()
_SERIALIZER.client_side_validation = False

def build_get_content_request(
    subscription_id: str,
    resource_group_name: str,
    automation_account_name: str,
    runbook_name: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
    accept = _headers.pop('Accept', "text/powershell")

    # Construct URL
    _url = kwargs.pop("template_url", "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content")  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, 'str'),
        "resourceGroupName": _SERIALIZER.url("resource_group_name", resource_group_name, 'str', max_length=90, min_length=1, pattern=r'^[-\w\._]+$'),
        "automationAccountName": _SERIALIZER.url("automation_account_name", automation_account_name, 'str'),
        "runbookName": _SERIALIZER.url("runbook_name", runbook_name, 'str'),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params['api-version'] = _SERIALIZER.query("api_version", api_version, 'str')

    # Construct headers
    _headers['Accept'] = _SERIALIZER.header("accept", accept, 'str')

    return HttpRequest(
        method="GET",
        url=_url,
        params=_params,
        headers=_headers,
        **kwargs
    )


def build_replace_content_request_initial(
    subscription_id: str,
    resource_group_name: str,
    automation_account_name: str,
    runbook_name: str,
    *,
    content: Any,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
    content_type = kwargs.pop('content_type', _headers.pop('Content-Type', None))  # type: Optional[str]
    accept = _headers.pop('Accept', "application/json")

    # Construct URL
    _url = kwargs.pop("template_url", "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content")  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, 'str'),
        "resourceGroupName": _SERIALIZER.url("resource_group_name", resource_group_name, 'str', max_length=90, min_length=1, pattern=r'^[-\w\._]+$'),
        "automationAccountName": _SERIALIZER.url("automation_account_name", automation_account_name, 'str'),
        "runbookName": _SERIALIZER.url("runbook_name", runbook_name, 'str'),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params['api-version'] = _SERIALIZER.query("api_version", api_version, 'str')

    # Construct headers
    if content_type is not None:
        _headers['Content-Type'] = _SERIALIZER.header("content_type", content_type, 'str')
    _headers['Accept'] = _SERIALIZER.header("accept", accept, 'str')

    return HttpRequest(
        method="PUT",
        url=_url,
        params=_params,
        headers=_headers,
        content=content,
        **kwargs
    )


def build_get_request(
    subscription_id: str,
    resource_group_name: str,
    automation_account_name: str,
    runbook_name: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
    accept = _headers.pop('Accept', "application/json")

    # Construct URL
    _url = kwargs.pop("template_url", "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft")  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, 'str'),
        "resourceGroupName": _SERIALIZER.url("resource_group_name", resource_group_name, 'str', max_length=90, min_length=1, pattern=r'^[-\w\._]+$'),
        "automationAccountName": _SERIALIZER.url("automation_account_name", automation_account_name, 'str'),
        "runbookName": _SERIALIZER.url("runbook_name", runbook_name, 'str'),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params['api-version'] = _SERIALIZER.query("api_version", api_version, 'str')

    # Construct headers
    _headers['Accept'] = _SERIALIZER.header("accept", accept, 'str')

    return HttpRequest(
        method="GET",
        url=_url,
        params=_params,
        headers=_headers,
        **kwargs
    )


def build_undo_edit_request(
    subscription_id: str,
    resource_group_name: str,
    automation_account_name: str,
    runbook_name: str,
    **kwargs: Any
) -> HttpRequest:
    _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
    _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

    api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
    accept = _headers.pop('Accept', "application/json")

    # Construct URL
    _url = kwargs.pop("template_url", "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/undoEdit")  # pylint: disable=line-too-long
    path_format_arguments = {
        "subscriptionId": _SERIALIZER.url("subscription_id", subscription_id, 'str'),
        "resourceGroupName": _SERIALIZER.url("resource_group_name", resource_group_name, 'str', max_length=90, min_length=1, pattern=r'^[-\w\._]+$'),
        "automationAccountName": _SERIALIZER.url("automation_account_name", automation_account_name, 'str'),
        "runbookName": _SERIALIZER.url("runbook_name", runbook_name, 'str'),
    }

    _url = _format_url_section(_url, **path_format_arguments)

    # Construct parameters
    _params['api-version'] = _SERIALIZER.query("api_version", api_version, 'str')

    # Construct headers
    _headers['Accept'] = _SERIALIZER.header("accept", accept, 'str')

    return HttpRequest(
        method="POST",
        url=_url,
        params=_params,
        headers=_headers,
        **kwargs
    )

class RunbookDraftOperations:
    """
    .. warning::
        **DO NOT** instantiate this class directly.

        Instead, you should access the following operations through
        :class:`~azure.mgmt.automation.AutomationClient`'s
        :attr:`runbook_draft` attribute.
    """

    models = _models

    def __init__(self, *args, **kwargs):
        input_args = list(args)
        self._client = input_args.pop(0) if input_args else kwargs.pop("client")
        self._config = input_args.pop(0) if input_args else kwargs.pop("config")
        self._serialize = input_args.pop(0) if input_args else kwargs.pop("serializer")
        self._deserialize = input_args.pop(0) if input_args else kwargs.pop("deserializer")


    @distributed_trace
    def get_content(
        self,
        resource_group_name: str,
        automation_account_name: str,
        runbook_name: str,
        **kwargs: Any
    ) -> IO:
        """Retrieve the content of runbook draft identified by runbook name.

        :param resource_group_name: Name of an Azure Resource group.
        :type resource_group_name: str
        :param automation_account_name: The name of the automation account.
        :type automation_account_name: str
        :param runbook_name: The runbook name.
        :type runbook_name: str
        :keyword api_version: Api Version. Default value is "2018-06-30". Note that overriding this
         default value may result in unsupported behavior.
        :paramtype api_version: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: IO, or the result of cls(response)
        :rtype: IO
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
        cls = kwargs.pop('cls', None)  # type: ClsType[IO]

        
        request = build_get_content_request(
            subscription_id=self._config.subscription_id,
            resource_group_name=resource_group_name,
            automation_account_name=automation_account_name,
            runbook_name=runbook_name,
            api_version=api_version,
            template_url=self.get_content.metadata['url'],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = self._deserialize('IO', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get_content.metadata = {'url': "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content"}  # type: ignore


    def _replace_content_initial(
        self,
        resource_group_name: str,
        automation_account_name: str,
        runbook_name: str,
        runbook_content: str,
        **kwargs: Any
    ) -> Optional[IO]:
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}) or {})

        _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
        content_type = kwargs.pop('content_type', _headers.pop('Content-Type', "text/powershell"))  # type: Optional[str]
        cls = kwargs.pop('cls', None)  # type: ClsType[Optional[IO]]

        _content = runbook_content

        request = build_replace_content_request_initial(
            subscription_id=self._config.subscription_id,
            resource_group_name=resource_group_name,
            automation_account_name=automation_account_name,
            runbook_name=runbook_name,
            api_version=api_version,
            content_type=content_type,
            content=_content,
            template_url=self._replace_content_initial.metadata['url'],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request,
            stream=True,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200, 202]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            raise HttpResponseError(response=response, error_format=ARMErrorFormat)

        deserialized = None
        response_headers = {}
        if response.status_code == 200:
            deserialized = response.stream_download(self._client._pipeline)

        if response.status_code == 202:
            response_headers['location']=self._deserialize('str', response.headers.get('location'))
            

        if cls:
            return cls(pipeline_response, deserialized, response_headers)

        return deserialized

    _replace_content_initial.metadata = {'url': "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content"}  # type: ignore


    @distributed_trace
    def begin_replace_content(
        self,
        resource_group_name: str,
        automation_account_name: str,
        runbook_name: str,
        runbook_content: str,
        **kwargs: Any
    ) -> LROPoller[IO]:
        """Replaces the runbook draft content.

        :param resource_group_name: Name of an Azure Resource group.
        :type resource_group_name: str
        :param automation_account_name: The name of the automation account.
        :type automation_account_name: str
        :param runbook_name: The runbook name.
        :type runbook_name: str
        :param runbook_content: The runbook draft content.
        :type runbook_content: str
        :keyword api_version: Api Version. Default value is "2018-06-30". Note that overriding this
         default value may result in unsupported behavior.
        :paramtype api_version: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :keyword str continuation_token: A continuation token to restart a poller from a saved state.
        :keyword polling: By default, your polling method will be ARMPolling. Pass in False for this
         operation to not poll, or pass in your own initialized polling object for a personal polling
         strategy.
        :paramtype polling: bool or ~azure.core.polling.PollingMethod
        :keyword int polling_interval: Default waiting time between two polls for LRO operations if no
         Retry-After header is present.
        :return: An instance of LROPoller that returns either IO or the result of cls(response)
        :rtype: ~azure.core.polling.LROPoller[IO]
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        _headers = case_insensitive_dict(kwargs.pop("headers", {}) or {})
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
        content_type = kwargs.pop('content_type', _headers.pop('Content-Type', "text/powershell"))  # type: Optional[str]
        cls = kwargs.pop('cls', None)  # type: ClsType[IO]
        polling = kwargs.pop('polling', True)  # type: Union[bool, PollingMethod]
        lro_delay = kwargs.pop(
            'polling_interval',
            self._config.polling_interval
        )
        cont_token = kwargs.pop('continuation_token', None)  # type: Optional[str]
        if cont_token is None:
            raw_result = self._replace_content_initial(  # type: ignore
                resource_group_name=resource_group_name,
                automation_account_name=automation_account_name,
                runbook_name=runbook_name,
                runbook_content=runbook_content,
                api_version=api_version,
                content_type=content_type,
                cls=lambda x,y,z: x,
                headers=_headers,
                params=_params,
                **kwargs
            )
        kwargs.pop('error_map', None)

        def get_long_running_output(pipeline_response):
            deserialized = response.stream_download(self._client._pipeline)
            if cls:
                return cls(pipeline_response, deserialized, {})
            return deserialized


        if polling is True:
            polling_method = cast(PollingMethod, ARMPolling(
                lro_delay,
                
                
                **kwargs
        ))  # type: PollingMethod
        elif polling is False: polling_method = cast(PollingMethod, NoPolling())
        else: polling_method = polling
        if cont_token:
            return LROPoller.from_continuation_token(
                polling_method=polling_method,
                continuation_token=cont_token,
                client=self._client,
                deserialization_callback=get_long_running_output
            )
        return LROPoller(self._client, raw_result, get_long_running_output, polling_method)

    begin_replace_content.metadata = {'url': "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/content"}  # type: ignore

    @distributed_trace
    def get(
        self,
        resource_group_name: str,
        automation_account_name: str,
        runbook_name: str,
        **kwargs: Any
    ) -> _models.RunbookDraft:
        """Retrieve the runbook draft identified by runbook name.

        :param resource_group_name: Name of an Azure Resource group.
        :type resource_group_name: str
        :param automation_account_name: The name of the automation account.
        :type automation_account_name: str
        :param runbook_name: The runbook name.
        :type runbook_name: str
        :keyword api_version: Api Version. Default value is "2018-06-30". Note that overriding this
         default value may result in unsupported behavior.
        :paramtype api_version: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: RunbookDraft, or the result of cls(response)
        :rtype: ~azure.mgmt.automation.models.RunbookDraft
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
        cls = kwargs.pop('cls', None)  # type: ClsType[_models.RunbookDraft]

        
        request = build_get_request(
            subscription_id=self._config.subscription_id,
            resource_group_name=resource_group_name,
            automation_account_name=automation_account_name,
            runbook_name=runbook_name,
            api_version=api_version,
            template_url=self.get.metadata['url'],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize('RunbookDraft', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    get.metadata = {'url': "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft"}  # type: ignore


    @distributed_trace
    def undo_edit(
        self,
        resource_group_name: str,
        automation_account_name: str,
        runbook_name: str,
        **kwargs: Any
    ) -> _models.RunbookDraftUndoEditResult:
        """Undo draft edit to last known published state identified by runbook name.

        :param resource_group_name: Name of an Azure Resource group.
        :type resource_group_name: str
        :param automation_account_name: The name of the automation account.
        :type automation_account_name: str
        :param runbook_name: The runbook name.
        :type runbook_name: str
        :keyword api_version: Api Version. Default value is "2018-06-30". Note that overriding this
         default value may result in unsupported behavior.
        :paramtype api_version: str
        :keyword callable cls: A custom type or function that will be passed the direct response
        :return: RunbookDraftUndoEditResult, or the result of cls(response)
        :rtype: ~azure.mgmt.automation.models.RunbookDraftUndoEditResult
        :raises: ~azure.core.exceptions.HttpResponseError
        """
        error_map = {
            401: ClientAuthenticationError, 404: ResourceNotFoundError, 409: ResourceExistsError
        }
        error_map.update(kwargs.pop('error_map', {}) or {})

        _headers = kwargs.pop("headers", {}) or {}
        _params = case_insensitive_dict(kwargs.pop("params", {}) or {})

        api_version = kwargs.pop('api_version', _params.pop('api-version', "2018-06-30"))  # type: str
        cls = kwargs.pop('cls', None)  # type: ClsType[_models.RunbookDraftUndoEditResult]

        
        request = build_undo_edit_request(
            subscription_id=self._config.subscription_id,
            resource_group_name=resource_group_name,
            automation_account_name=automation_account_name,
            runbook_name=runbook_name,
            api_version=api_version,
            template_url=self.undo_edit.metadata['url'],
            headers=_headers,
            params=_params,
        )
        request = _convert_request(request)
        request.url = self._client.format_url(request.url)  # type: ignore

        pipeline_response = self._client._pipeline.run(  # type: ignore # pylint: disable=protected-access
            request,
            stream=False,
            **kwargs
        )
        response = pipeline_response.http_response

        if response.status_code not in [200]:
            map_error(status_code=response.status_code, response=response, error_map=error_map)
            error = self._deserialize.failsafe_deserialize(_models.ErrorResponse, pipeline_response)
            raise HttpResponseError(response=response, model=error, error_format=ARMErrorFormat)

        deserialized = self._deserialize('RunbookDraftUndoEditResult', pipeline_response)

        if cls:
            return cls(pipeline_response, deserialized, {})

        return deserialized

    undo_edit.metadata = {'url': "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/undoEdit"}  # type: ignore

