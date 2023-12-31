# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------
import json
from azext_aks_preview.azuremonitormetrics.constants import MapToClosestMACRegion
from azext_aks_preview.azuremonitormetrics.dc.defaults import get_default_region, sanitize_name
from azext_aks_preview.azuremonitormetrics.constants import (
    DC_TYPE,
    DC_API
)
from knack.util import CLIError


def get_default_dcr_name(cmd, mac_region, cluster_name):
    region = MapToClosestMACRegion.get(mac_region, get_default_region(cmd))
    default_dcr_name = "MSProm-" + region + "-" + cluster_name
    return sanitize_name(default_dcr_name, DC_TYPE.DCR, 64)


# pylint: disable=too-many-locals,too-many-branches,too-many-statements
def create_dcr(
    cmd,
    mac_region,
    azure_monitor_workspace_resource_id,
    cluster_subscription,
    cluster_resource_group_name,
    cluster_name,
    dce_resource_id
):
    from azure.cli.core.util import send_raw_request
    dcr_name = get_default_dcr_name(cmd, mac_region, cluster_name)
    dcr_resource_id = (
        f"/subscriptions/{cluster_subscription}/resourceGroups/{cluster_resource_group_name}/providers/"
        f"Microsoft.Insights/dataCollectionRules/{dcr_name}"
    )
    dcr_creation_body = json.dumps(
        {
            "location": mac_region,
            "kind": "Linux",
            "properties": {
                "dataCollectionEndpointId": dce_resource_id,
                "dataSources": {
                    "prometheusForwarder": [
                        {
                            "name": "PrometheusDataSource",
                            "streams": ["Microsoft-PrometheusMetrics"],
                            "labelIncludeFilter": {},
                        }
                    ]
                },
                "dataFlows": [
                    {
                        "destinations": ["MonitoringAccount1"],
                        "streams": ["Microsoft-PrometheusMetrics"],
                    }
                ],
                "description": "DCR description",
                "destinations": {
                    "monitoringAccounts": [
                        {
                            "accountResourceId": azure_monitor_workspace_resource_id,
                            "name": "MonitoringAccount1",
                        }
                    ]
                },
            },
        }
    )
    armendpoint = cmd.cli_ctx.cloud.endpoints.resource_manager
    dcr_url = f"{armendpoint}{dcr_resource_id}?api-version={DC_API}"
    try:
        headers = ['User-Agent=azuremonitormetrics.create_dcr']
        send_raw_request(cmd.cli_ctx, "PUT",
                         dcr_url, body=dcr_creation_body, headers=headers)
        return dcr_resource_id
    except CLIError as error:
        raise error
