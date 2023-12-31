# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
#
# Code generated by aaz-dev-tools
# --------------------------------------------------------------------------------------------

# pylint: skip-file
# flake8: noqa

from azure.cli.core.aaz import *


@register_command(
    "site-recovery event show",
)
class Show(AAZCommand):
    """Get operation to get the details of an Azure Site recovery event.

    :example: event show
        az site-recovery event show -g rg --vault-name vault_name -n event_name
    """

    _aaz_info = {
        "version": "2022-08-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.recoveryservices/vaults/{}/replicationevents/{}", "2022-08-01"],
        ]
    }

    def _handler(self, command_args):
        super()._handler(command_args)
        self._execute_operations()
        return self._output()

    _args_schema = None

    @classmethod
    def _build_arguments_schema(cls, *args, **kwargs):
        if cls._args_schema is not None:
            return cls._args_schema
        cls._args_schema = super()._build_arguments_schema(*args, **kwargs)

        # define Arg Group ""

        _args_schema = cls._args_schema
        _args_schema.event_name = AAZStrArg(
            options=["-n", "--name", "--event-name"],
            help="The name of the Azure Site Recovery event.",
            required=True,
            id_part="child_name_1",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        _args_schema.vault_name = AAZStrArg(
            options=["--vault-name"],
            help="The name of the recovery services vault.",
            required=True,
            id_part="name",
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.ReplicationEventsGet(ctx=self.ctx)()
        self.post_operations()

    @register_callback
    def pre_operations(self):
        pass

    @register_callback
    def post_operations(self):
        pass

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.vars.instance, client_flatten=True)
        return result

    class ReplicationEventsGet(AAZHttpOperation):
        CLIENT_TYPE = "MgmtClient"

        def __call__(self, *args, **kwargs):
            request = self.make_request()
            session = self.client.send_request(request=request, stream=False, **kwargs)
            if session.http_response.status_code in [200]:
                return self.on_200(session)

            return self.on_error(session.http_response)

        @property
        def url(self):
            return self.client.format_url(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationEvents/{eventName}",
                **self.url_parameters
            )

        @property
        def method(self):
            return "GET"

        @property
        def error_format(self):
            return "ODataV4Format"

        @property
        def url_parameters(self):
            parameters = {
                **self.serialize_url_param(
                    "eventName", self.ctx.args.event_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "resourceGroupName", self.ctx.args.resource_group,
                    required=True,
                ),
                **self.serialize_url_param(
                    "resourceName", self.ctx.args.vault_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "subscriptionId", self.ctx.subscription_id,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2022-08-01",
                    required=True,
                ),
            }
            return parameters

        @property
        def header_parameters(self):
            parameters = {
                **self.serialize_header_param(
                    "Accept", "application/json",
                ),
            }
            return parameters

        def on_200(self, session):
            data = self.deserialize_http_content(session)
            self.ctx.set_var(
                "instance",
                data,
                schema_builder=self._build_schema_on_200
            )

        _schema_on_200 = None

        @classmethod
        def _build_schema_on_200(cls):
            if cls._schema_on_200 is not None:
                return cls._schema_on_200

            cls._schema_on_200 = AAZObjectType()

            _schema_on_200 = cls._schema_on_200
            _schema_on_200.id = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.location = AAZStrType()
            _schema_on_200.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.properties = AAZObjectType()
            _schema_on_200.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.properties
            properties.affected_object_correlation_id = AAZStrType(
                serialized_name="affectedObjectCorrelationId",
            )
            properties.affected_object_friendly_name = AAZStrType(
                serialized_name="affectedObjectFriendlyName",
            )
            properties.description = AAZStrType()
            properties.event_code = AAZStrType(
                serialized_name="eventCode",
            )
            properties.event_specific_details = AAZObjectType(
                serialized_name="eventSpecificDetails",
            )
            properties.event_type = AAZStrType(
                serialized_name="eventType",
            )
            properties.fabric_id = AAZStrType(
                serialized_name="fabricId",
            )
            properties.health_errors = AAZListType(
                serialized_name="healthErrors",
            )
            properties.provider_specific_details = AAZObjectType(
                serialized_name="providerSpecificDetails",
            )
            properties.severity = AAZStrType()
            properties.time_of_occurrence = AAZStrType(
                serialized_name="timeOfOccurrence",
            )

            event_specific_details = cls._schema_on_200.properties.event_specific_details
            event_specific_details.instance_type = AAZStrType(
                serialized_name="instanceType",
                flags={"required": True},
            )

            disc_job_status = cls._schema_on_200.properties.event_specific_details.discriminate_by("instance_type", "JobStatus")
            disc_job_status.affected_object_type = AAZStrType(
                serialized_name="affectedObjectType",
            )
            disc_job_status.job_friendly_name = AAZStrType(
                serialized_name="jobFriendlyName",
            )
            disc_job_status.job_id = AAZStrType(
                serialized_name="jobId",
            )
            disc_job_status.job_status = AAZStrType(
                serialized_name="jobStatus",
            )

            health_errors = cls._schema_on_200.properties.health_errors
            health_errors.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.health_errors.Element
            _element.creation_time_utc = AAZStrType(
                serialized_name="creationTimeUtc",
            )
            _element.customer_resolvability = AAZStrType(
                serialized_name="customerResolvability",
            )
            _element.entity_id = AAZStrType(
                serialized_name="entityId",
            )
            _element.error_category = AAZStrType(
                serialized_name="errorCategory",
            )
            _element.error_code = AAZStrType(
                serialized_name="errorCode",
            )
            _element.error_id = AAZStrType(
                serialized_name="errorId",
            )
            _element.error_level = AAZStrType(
                serialized_name="errorLevel",
            )
            _element.error_message = AAZStrType(
                serialized_name="errorMessage",
            )
            _element.error_source = AAZStrType(
                serialized_name="errorSource",
            )
            _element.error_type = AAZStrType(
                serialized_name="errorType",
            )
            _element.inner_health_errors = AAZListType(
                serialized_name="innerHealthErrors",
            )
            _element.possible_causes = AAZStrType(
                serialized_name="possibleCauses",
            )
            _element.recommended_action = AAZStrType(
                serialized_name="recommendedAction",
            )
            _element.recovery_provider_error_message = AAZStrType(
                serialized_name="recoveryProviderErrorMessage",
            )
            _element.summary_message = AAZStrType(
                serialized_name="summaryMessage",
            )

            inner_health_errors = cls._schema_on_200.properties.health_errors.Element.inner_health_errors
            inner_health_errors.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.health_errors.Element.inner_health_errors.Element
            _element.creation_time_utc = AAZStrType(
                serialized_name="creationTimeUtc",
            )
            _element.customer_resolvability = AAZStrType(
                serialized_name="customerResolvability",
            )
            _element.entity_id = AAZStrType(
                serialized_name="entityId",
            )
            _element.error_category = AAZStrType(
                serialized_name="errorCategory",
            )
            _element.error_code = AAZStrType(
                serialized_name="errorCode",
            )
            _element.error_id = AAZStrType(
                serialized_name="errorId",
            )
            _element.error_level = AAZStrType(
                serialized_name="errorLevel",
            )
            _element.error_message = AAZStrType(
                serialized_name="errorMessage",
            )
            _element.error_source = AAZStrType(
                serialized_name="errorSource",
            )
            _element.error_type = AAZStrType(
                serialized_name="errorType",
            )
            _element.possible_causes = AAZStrType(
                serialized_name="possibleCauses",
            )
            _element.recommended_action = AAZStrType(
                serialized_name="recommendedAction",
            )
            _element.recovery_provider_error_message = AAZStrType(
                serialized_name="recoveryProviderErrorMessage",
            )
            _element.summary_message = AAZStrType(
                serialized_name="summaryMessage",
            )

            provider_specific_details = cls._schema_on_200.properties.provider_specific_details
            provider_specific_details.instance_type = AAZStrType(
                serialized_name="instanceType",
                flags={"required": True},
            )

            disc_a2_a = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "A2A")
            disc_a2_a.fabric_location = AAZStrType(
                serialized_name="fabricLocation",
            )
            disc_a2_a.fabric_name = AAZStrType(
                serialized_name="fabricName",
            )
            disc_a2_a.fabric_object_id = AAZStrType(
                serialized_name="fabricObjectId",
            )
            disc_a2_a.protected_item_name = AAZStrType(
                serialized_name="protectedItemName",
            )
            disc_a2_a.remote_fabric_location = AAZStrType(
                serialized_name="remoteFabricLocation",
            )
            disc_a2_a.remote_fabric_name = AAZStrType(
                serialized_name="remoteFabricName",
            )

            disc_hyper_v_replica2012 = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "HyperVReplica2012")
            disc_hyper_v_replica2012.container_name = AAZStrType(
                serialized_name="containerName",
            )
            disc_hyper_v_replica2012.fabric_name = AAZStrType(
                serialized_name="fabricName",
            )
            disc_hyper_v_replica2012.remote_container_name = AAZStrType(
                serialized_name="remoteContainerName",
            )
            disc_hyper_v_replica2012.remote_fabric_name = AAZStrType(
                serialized_name="remoteFabricName",
            )

            disc_hyper_v_replica2012_r2 = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "HyperVReplica2012R2")
            disc_hyper_v_replica2012_r2.container_name = AAZStrType(
                serialized_name="containerName",
            )
            disc_hyper_v_replica2012_r2.fabric_name = AAZStrType(
                serialized_name="fabricName",
            )
            disc_hyper_v_replica2012_r2.remote_container_name = AAZStrType(
                serialized_name="remoteContainerName",
            )
            disc_hyper_v_replica2012_r2.remote_fabric_name = AAZStrType(
                serialized_name="remoteFabricName",
            )

            disc_hyper_v_replica_azure = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "HyperVReplicaAzure")
            disc_hyper_v_replica_azure.container_name = AAZStrType(
                serialized_name="containerName",
            )
            disc_hyper_v_replica_azure.fabric_name = AAZStrType(
                serialized_name="fabricName",
            )
            disc_hyper_v_replica_azure.remote_container_name = AAZStrType(
                serialized_name="remoteContainerName",
            )

            disc_hyper_v_replica_base_event_details = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "HyperVReplicaBaseEventDetails")
            disc_hyper_v_replica_base_event_details.container_name = AAZStrType(
                serialized_name="containerName",
            )
            disc_hyper_v_replica_base_event_details.fabric_name = AAZStrType(
                serialized_name="fabricName",
            )
            disc_hyper_v_replica_base_event_details.remote_container_name = AAZStrType(
                serialized_name="remoteContainerName",
            )
            disc_hyper_v_replica_base_event_details.remote_fabric_name = AAZStrType(
                serialized_name="remoteFabricName",
            )

            disc_in_mage_azure_v2 = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "InMageAzureV2")
            disc_in_mage_azure_v2.category = AAZStrType()
            disc_in_mage_azure_v2.component = AAZStrType()
            disc_in_mage_azure_v2.corrective_action = AAZStrType(
                serialized_name="correctiveAction",
            )
            disc_in_mage_azure_v2.details = AAZStrType()
            disc_in_mage_azure_v2.event_type = AAZStrType(
                serialized_name="eventType",
            )
            disc_in_mage_azure_v2.site_name = AAZStrType(
                serialized_name="siteName",
            )
            disc_in_mage_azure_v2.summary = AAZStrType()

            disc_in_mage_rcm = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "InMageRcm")
            disc_in_mage_rcm.appliance_name = AAZStrType(
                serialized_name="applianceName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.component_display_name = AAZStrType(
                serialized_name="componentDisplayName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.fabric_name = AAZStrType(
                serialized_name="fabricName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.job_id = AAZStrType(
                serialized_name="jobId",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.latest_agent_version = AAZStrType(
                serialized_name="latestAgentVersion",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.protected_item_name = AAZStrType(
                serialized_name="protectedItemName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.server_type = AAZStrType(
                serialized_name="serverType",
                flags={"read_only": True},
            )
            disc_in_mage_rcm.vm_name = AAZStrType(
                serialized_name="vmName",
                flags={"read_only": True},
            )

            disc_in_mage_rcm_failback = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "InMageRcmFailback")
            disc_in_mage_rcm_failback.appliance_name = AAZStrType(
                serialized_name="applianceName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm_failback.component_display_name = AAZStrType(
                serialized_name="componentDisplayName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm_failback.protected_item_name = AAZStrType(
                serialized_name="protectedItemName",
                flags={"read_only": True},
            )
            disc_in_mage_rcm_failback.server_type = AAZStrType(
                serialized_name="serverType",
                flags={"read_only": True},
            )
            disc_in_mage_rcm_failback.vm_name = AAZStrType(
                serialized_name="vmName",
                flags={"read_only": True},
            )

            disc_v_mware_cbt = cls._schema_on_200.properties.provider_specific_details.discriminate_by("instance_type", "VMwareCbt")
            disc_v_mware_cbt.migration_item_name = AAZStrType(
                serialized_name="migrationItemName",
                flags={"read_only": True},
            )

            return cls._schema_on_200


class _ShowHelper:
    """Helper class for Show"""


__all__ = ["Show"]
