# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------


def cf_stream_analytics_cl(cli_ctx, *_):
    from azure.cli.core.commands.client_factory import get_mgmt_service_client
    from azext_stream_analytics.vendored_sdks.streamanalytics import StreamAnalyticsManagementClient
    return get_mgmt_service_client(cli_ctx,
                                   StreamAnalyticsManagementClient)


def cf_streaming_job(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).streaming_jobs


def cf_input(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).inputs


def cf_output(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).outputs


def cf_transformation(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).transformations


def cf_function(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).functions


def cf_subscription(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).subscriptions


def cf_cluster(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).clusters


def cf_private_endpoint(cli_ctx, *_):
    return cf_stream_analytics_cl(cli_ctx).private_endpoints
