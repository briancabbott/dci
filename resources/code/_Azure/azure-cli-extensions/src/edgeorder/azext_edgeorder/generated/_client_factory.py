# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------


def cf_edgeorder_cl(cli_ctx, *_):
    from azure.cli.core.commands.client_factory import get_mgmt_service_client
    from azext_edgeorder.vendored_sdks.edgeorder import EdgeOrderManagementClient
    return get_mgmt_service_client(cli_ctx,
                                   EdgeOrderManagementClient)


def cf_address(cli_ctx, *_):
    return cf_edgeorder_cl(cli_ctx).addresses


def cf_order(cli_ctx, *_):
    return cf_edgeorder_cl(cli_ctx).order


def cf_order_item(cli_ctx, *_):
    return cf_edgeorder_cl(cli_ctx).order_items
