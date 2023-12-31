# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from typing import TYPE_CHECKING

from azure.mgmt.core import ARMPipelineClient
from msrest import Deserializer, Serializer

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from typing import Any, Optional

    from azure.core.credentials import TokenCredential

from ._configuration import MaintenanceManagementClientConfiguration
from .operations import PublicMaintenanceConfigurationsOperations
from .operations import ApplyUpdatesOperations
from .operations import ConfigurationAssignmentsOperations
from .operations import MaintenanceConfigurationsOperations
from .operations import MaintenanceConfigurationsForResourceGroupOperations
from .operations import ApplyUpdateForResourceGroupOperations
from .operations import ConfigurationAssignmentsWithinSubscriptionOperations
from .operations import Operations
from .operations import UpdatesOperations
from . import models


class MaintenanceManagementClient(object):
    """Azure Maintenance Management Client.

    :ivar public_maintenance_configurations: PublicMaintenanceConfigurationsOperations operations
    :vartype public_maintenance_configurations: maintenance_management_client.operations.PublicMaintenanceConfigurationsOperations
    :ivar apply_updates: ApplyUpdatesOperations operations
    :vartype apply_updates: maintenance_management_client.operations.ApplyUpdatesOperations
    :ivar configuration_assignments: ConfigurationAssignmentsOperations operations
    :vartype configuration_assignments: maintenance_management_client.operations.ConfigurationAssignmentsOperations
    :ivar maintenance_configurations: MaintenanceConfigurationsOperations operations
    :vartype maintenance_configurations: maintenance_management_client.operations.MaintenanceConfigurationsOperations
    :ivar maintenance_configurations_for_resource_group: MaintenanceConfigurationsForResourceGroupOperations operations
    :vartype maintenance_configurations_for_resource_group: maintenance_management_client.operations.MaintenanceConfigurationsForResourceGroupOperations
    :ivar apply_update_for_resource_group: ApplyUpdateForResourceGroupOperations operations
    :vartype apply_update_for_resource_group: maintenance_management_client.operations.ApplyUpdateForResourceGroupOperations
    :ivar configuration_assignments_within_subscription: ConfigurationAssignmentsWithinSubscriptionOperations operations
    :vartype configuration_assignments_within_subscription: maintenance_management_client.operations.ConfigurationAssignmentsWithinSubscriptionOperations
    :ivar operations: Operations operations
    :vartype operations: maintenance_management_client.operations.Operations
    :ivar updates: UpdatesOperations operations
    :vartype updates: maintenance_management_client.operations.UpdatesOperations
    :param credential: Credential needed for the client to connect to Azure.
    :type credential: ~azure.core.credentials.TokenCredential
    :param subscription_id: Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.
    :type subscription_id: str
    :param str base_url: Service URL
    """

    def __init__(
        self,
        credential,  # type: "TokenCredential"
        subscription_id,  # type: str
        base_url=None,  # type: Optional[str]
        **kwargs  # type: Any
    ):
        # type: (...) -> None
        if not base_url:
            base_url = 'https://management.azure.com'
        self._config = MaintenanceManagementClientConfiguration(credential, subscription_id, **kwargs)
        self._client = ARMPipelineClient(base_url=base_url, config=self._config, **kwargs)

        client_models = {k: v for k, v in models.__dict__.items() if isinstance(v, type)}
        self._serialize = Serializer(client_models)
        self._deserialize = Deserializer(client_models)

        self.public_maintenance_configurations = PublicMaintenanceConfigurationsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.apply_updates = ApplyUpdatesOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.configuration_assignments = ConfigurationAssignmentsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.maintenance_configurations = MaintenanceConfigurationsOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.maintenance_configurations_for_resource_group = MaintenanceConfigurationsForResourceGroupOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.apply_update_for_resource_group = ApplyUpdateForResourceGroupOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.configuration_assignments_within_subscription = ConfigurationAssignmentsWithinSubscriptionOperations(
            self._client, self._config, self._serialize, self._deserialize)
        self.operations = Operations(
            self._client, self._config, self._serialize, self._deserialize)
        self.updates = UpdatesOperations(
            self._client, self._config, self._serialize, self._deserialize)

    def close(self):
        # type: () -> None
        self._client.close()

    def __enter__(self):
        # type: () -> MaintenanceManagementClient
        self._client.__enter__()
        return self

    def __exit__(self, *exc_details):
        # type: (Any) -> None
        self._client.__exit__(*exc_details)
