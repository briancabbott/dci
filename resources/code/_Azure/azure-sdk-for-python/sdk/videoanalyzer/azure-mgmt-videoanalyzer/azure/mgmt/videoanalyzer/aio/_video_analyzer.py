# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from typing import Any, Optional, TYPE_CHECKING

from azure.core.pipeline.transport import AsyncHttpResponse, HttpRequest
from azure.mgmt.core import AsyncARMPipelineClient
from msrest import Deserializer, Serializer

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from azure.core.credentials_async import AsyncTokenCredential

from ._configuration import VideoAnalyzerConfiguration
from .operations import EdgeModulesOperations
from .operations import PipelineTopologiesOperations
from .operations import LivePipelinesOperations
from .operations import PipelineJobsOperations
from .operations import LivePipelineOperationStatusesOperations
from .operations import PipelineJobOperationStatusesOperations
from .operations import Operations
from .operations import VideoAnalyzersOperations
from .operations import PrivateLinkResourcesOperations
from .operations import PrivateEndpointConnectionsOperations
from .operations import OperationStatusesOperations
from .operations import OperationResultsOperations
from .operations import VideoAnalyzerOperationStatusesOperations
from .operations import VideoAnalyzerOperationResultsOperations
from .operations import LocationsOperations
from .operations import VideosOperations
from .operations import AccessPoliciesOperations
from .. import models


class VideoAnalyzer(object):
    """Azure Video Analyzer provides a platform for you to build intelligent video applications that span the edge and the cloud.

    :ivar edge_modules: EdgeModulesOperations operations
    :vartype edge_modules: video_analyzer.aio.operations.EdgeModulesOperations
    :ivar pipeline_topologies: PipelineTopologiesOperations operations
    :vartype pipeline_topologies: video_analyzer.aio.operations.PipelineTopologiesOperations
    :ivar live_pipelines: LivePipelinesOperations operations
    :vartype live_pipelines: video_analyzer.aio.operations.LivePipelinesOperations
    :ivar pipeline_jobs: PipelineJobsOperations operations
    :vartype pipeline_jobs: video_analyzer.aio.operations.PipelineJobsOperations
    :ivar live_pipeline_operation_statuses: LivePipelineOperationStatusesOperations operations
    :vartype live_pipeline_operation_statuses: video_analyzer.aio.operations.LivePipelineOperationStatusesOperations
    :ivar pipeline_job_operation_statuses: PipelineJobOperationStatusesOperations operations
    :vartype pipeline_job_operation_statuses: video_analyzer.aio.operations.PipelineJobOperationStatusesOperations
    :ivar operations: Operations operations
    :vartype operations: video_analyzer.aio.operations.Operations
    :ivar video_analyzers: VideoAnalyzersOperations operations
    :vartype video_analyzers: video_analyzer.aio.operations.VideoAnalyzersOperations
    :ivar private_link_resources: PrivateLinkResourcesOperations operations
    :vartype private_link_resources: video_analyzer.aio.operations.PrivateLinkResourcesOperations
    :ivar private_endpoint_connections: PrivateEndpointConnectionsOperations operations
    :vartype private_endpoint_connections: video_analyzer.aio.operations.PrivateEndpointConnectionsOperations
    :ivar operation_statuses: OperationStatusesOperations operations
    :vartype operation_statuses: video_analyzer.aio.operations.OperationStatusesOperations
    :ivar operation_results: OperationResultsOperations operations
    :vartype operation_results: video_analyzer.aio.operations.OperationResultsOperations
    :ivar video_analyzer_operation_statuses: VideoAnalyzerOperationStatusesOperations operations
    :vartype video_analyzer_operation_statuses: video_analyzer.aio.operations.VideoAnalyzerOperationStatusesOperations
    :ivar video_analyzer_operation_results: VideoAnalyzerOperationResultsOperations operations
    :vartype video_analyzer_operation_results: video_analyzer.aio.operations.VideoAnalyzerOperationResultsOperations
    :ivar locations: LocationsOperations operations
    :vartype locations: video_analyzer.aio.operations.LocationsOperations
    :ivar videos: VideosOperations operations
    :vartype videos: video_analyzer.aio.operations.VideosOperations
    :ivar access_policies: AccessPoliciesOperations operations
    :vartype access_policies: video_analyzer.aio.operations.AccessPoliciesOperations
    :param credential: Credential needed for the client to connect to Azure.
    :type credential: ~azure.core.credentials_async.AsyncTokenCredential
    :param subscription_id: The ID of the target subscription.
    :type subscription_id: str
    :param str base_url: Service URL
    :keyword int polling_interval: Default waiting time between two polls for LRO operations if no Retry-After header is present.
    """

    def __init__(
        self,
        credential: "AsyncTokenCredential",
        subscription_id: str,
        base_url: Optional[str] = None,
        **kwargs: Any
    ) -> None:
        if not base_url:
            base_url = 'https://management.azure.com'
        self._config = VideoAnalyzerConfiguration(credential, subscription_id, **kwargs)
        self._client = AsyncARMPipelineClient(base_url=base_url, config=self._config, **kwargs)

        client_models = {k: v for k, v in models.__dict__.items() if isinstance(v, type)}
        self._serialize = Serializer(client_models)
        self._serialize.client_side_validation = False
        self._deserialize = Deserializer(client_models)

        self.edge_modules = EdgeModulesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.pipeline_topologies = PipelineTopologiesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.live_pipelines = LivePipelinesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.pipeline_jobs = PipelineJobsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.live_pipeline_operation_statuses = LivePipelineOperationStatusesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.pipeline_job_operation_statuses = PipelineJobOperationStatusesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.operations = Operations(
            self._client, self._config, self._serialize, self._deserialize)
        self.video_analyzers = VideoAnalyzersOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.private_link_resources = PrivateLinkResourcesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.private_endpoint_connections = PrivateEndpointConnectionsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.operation_statuses = OperationStatusesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.operation_results = OperationResultsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.video_analyzer_operation_statuses = VideoAnalyzerOperationStatusesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.video_analyzer_operation_results = VideoAnalyzerOperationResultsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.locations = LocationsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.videos = VideosOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.access_policies = AccessPoliciesOperations(
            self._client, self._config, self._serialize, self._deserialize)

    async def _send_request(self, http_request: HttpRequest, **kwargs: Any) -> AsyncHttpResponse:
        """Runs the network request through the client's chained policies.

        :param http_request: The network request you want to make. Required.
        :type http_request: ~azure.core.pipeline.transport.HttpRequest
        :keyword bool stream: Whether the response payload will be streamed. Defaults to True.
        :return: The response of your network call. Does not do error handling on your response.
        :rtype: ~azure.core.pipeline.transport.AsyncHttpResponse
        """
        path_format_arguments = {
            'subscriptionId': self._serialize.url("self._config.subscription_id", self._config.subscription_id, 'str', min_length=1),
        }
        http_request.url = self._client.format_url(http_request.url, **path_format_arguments)
        stream = kwargs.pop("stream", True)
        pipeline_response = await self._client._pipeline.run(http_request, stream=stream, **kwargs)
        return pipeline_response.http_response

    async def close(self) -> None:
        await self._client.close()

    async def __aenter__(self) -> "VideoAnalyzer":
        await self._client.__aenter__()
        return self

    async def __aexit__(self, *exc_details) -> None:
        await self._client.__aexit__(*exc_details)
