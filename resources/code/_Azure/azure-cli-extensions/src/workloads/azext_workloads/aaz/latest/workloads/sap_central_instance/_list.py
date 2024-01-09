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
    "workloads sap-central-instance list",
    is_preview=True,
)
class List(AAZCommand):
    """List the SAP Central Services Instance resource for the given Virtual Instance for SAP solutions resource.

    :example: Get an overview of the Central service Instance in a Virtual instance for SAP solutions (VIS)
        az workloads sap-central-instance list -g <Resource-group-name> --sap-virtual-instance-name <VIS name>
    """

    _aaz_info = {
        "version": "2023-04-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.workloads/sapvirtualinstances/{}/centralinstances", "2023-04-01"],
        ]
    }

    def _handler(self, command_args):
        super()._handler(command_args)
        return self.build_paging(self._execute_operations, self._output)

    _args_schema = None

    @classmethod
    def _build_arguments_schema(cls, *args, **kwargs):
        if cls._args_schema is not None:
            return cls._args_schema
        cls._args_schema = super()._build_arguments_schema(*args, **kwargs)

        # define Arg Group ""

        _args_schema = cls._args_schema
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        _args_schema.sap_virtual_instance_name = AAZStrArg(
            options=["--vis-name", "--sap-virtual-instance-name"],
            help="The name of the Virtual Instances for SAP solutions resource",
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.SAPCentralInstancesList(ctx=self.ctx)()
        self.post_operations()

    @register_callback
    def pre_operations(self):
        pass

    @register_callback
    def post_operations(self):
        pass

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.vars.instance.value, client_flatten=True)
        next_link = self.deserialize_output(self.ctx.vars.instance.next_link)
        return result, next_link

    class SAPCentralInstancesList(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/centralInstances",
                **self.url_parameters
            )

        @property
        def method(self):
            return "GET"

        @property
        def error_format(self):
            return "MgmtErrorFormat"

        @property
        def url_parameters(self):
            parameters = {
                **self.serialize_url_param(
                    "resourceGroupName", self.ctx.args.resource_group,
                    required=True,
                ),
                **self.serialize_url_param(
                    "sapVirtualInstanceName", self.ctx.args.sap_virtual_instance_name,
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
                    "api-version", "2023-04-01",
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
            _schema_on_200.next_link = AAZStrType(
                serialized_name="nextLink",
            )
            _schema_on_200.value = AAZListType()

            value = cls._schema_on_200.value
            value.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element
            _element.id = AAZStrType(
                flags={"read_only": True},
            )
            _element.location = AAZStrType(
                flags={"required": True},
            )
            _element.name = AAZStrType(
                flags={"read_only": True},
            )
            _element.properties = AAZObjectType(
                flags={"client_flatten": True},
            )
            _element.system_data = AAZObjectType(
                serialized_name="systemData",
                flags={"read_only": True},
            )
            _element.tags = AAZDictType()
            _element.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.value.Element.properties
            properties.enqueue_replication_server_properties = AAZObjectType(
                serialized_name="enqueueReplicationServerProperties",
            )
            properties.enqueue_server_properties = AAZObjectType(
                serialized_name="enqueueServerProperties",
            )
            properties.errors = AAZObjectType()
            properties.gateway_server_properties = AAZObjectType(
                serialized_name="gatewayServerProperties",
            )
            properties.health = AAZStrType()
            properties.instance_no = AAZStrType(
                serialized_name="instanceNo",
                flags={"read_only": True},
            )
            properties.kernel_patch = AAZStrType(
                serialized_name="kernelPatch",
                nullable=True,
                flags={"read_only": True},
            )
            properties.kernel_version = AAZStrType(
                serialized_name="kernelVersion",
                nullable=True,
                flags={"read_only": True},
            )
            properties.load_balancer_details = AAZObjectType(
                serialized_name="loadBalancerDetails",
            )
            properties.message_server_properties = AAZObjectType(
                serialized_name="messageServerProperties",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.status = AAZStrType()
            properties.subnet = AAZStrType(
                flags={"read_only": True},
            )
            properties.vm_details = AAZListType(
                serialized_name="vmDetails",
                flags={"read_only": True},
            )

            enqueue_replication_server_properties = cls._schema_on_200.value.Element.properties.enqueue_replication_server_properties
            enqueue_replication_server_properties.ers_version = AAZStrType(
                serialized_name="ersVersion",
                flags={"read_only": True},
            )
            enqueue_replication_server_properties.health = AAZStrType()
            enqueue_replication_server_properties.hostname = AAZStrType(
                flags={"read_only": True},
            )
            enqueue_replication_server_properties.instance_no = AAZStrType(
                serialized_name="instanceNo",
                flags={"read_only": True},
            )
            enqueue_replication_server_properties.ip_address = AAZStrType(
                serialized_name="ipAddress",
                flags={"read_only": True},
            )
            enqueue_replication_server_properties.kernel_patch = AAZStrType(
                serialized_name="kernelPatch",
                flags={"read_only": True},
            )
            enqueue_replication_server_properties.kernel_version = AAZStrType(
                serialized_name="kernelVersion",
                flags={"read_only": True},
            )

            enqueue_server_properties = cls._schema_on_200.value.Element.properties.enqueue_server_properties
            enqueue_server_properties.health = AAZStrType()
            enqueue_server_properties.hostname = AAZStrType(
                flags={"read_only": True},
            )
            enqueue_server_properties.ip_address = AAZStrType(
                serialized_name="ipAddress",
                flags={"read_only": True},
            )
            enqueue_server_properties.port = AAZIntType(
                nullable=True,
                flags={"read_only": True},
            )

            errors = cls._schema_on_200.value.Element.properties.errors
            errors.properties = AAZObjectType()
            _ListHelper._build_schema_error_definition_read(errors.properties)

            gateway_server_properties = cls._schema_on_200.value.Element.properties.gateway_server_properties
            gateway_server_properties.health = AAZStrType()
            gateway_server_properties.port = AAZIntType(
                nullable=True,
                flags={"read_only": True},
            )

            load_balancer_details = cls._schema_on_200.value.Element.properties.load_balancer_details
            load_balancer_details.id = AAZStrType(
                flags={"read_only": True},
            )

            message_server_properties = cls._schema_on_200.value.Element.properties.message_server_properties
            message_server_properties.health = AAZStrType()
            message_server_properties.hostname = AAZStrType(
                flags={"read_only": True},
            )
            message_server_properties.http_port = AAZIntType(
                serialized_name="httpPort",
                nullable=True,
                flags={"read_only": True},
            )
            message_server_properties.https_port = AAZIntType(
                serialized_name="httpsPort",
                nullable=True,
                flags={"read_only": True},
            )
            message_server_properties.internal_ms_port = AAZIntType(
                serialized_name="internalMsPort",
                nullable=True,
                flags={"read_only": True},
            )
            message_server_properties.ip_address = AAZStrType(
                serialized_name="ipAddress",
                flags={"read_only": True},
            )
            message_server_properties.ms_port = AAZIntType(
                serialized_name="msPort",
                nullable=True,
                flags={"read_only": True},
            )

            vm_details = cls._schema_on_200.value.Element.properties.vm_details
            vm_details.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.vm_details.Element
            _element.storage_details = AAZListType(
                serialized_name="storageDetails",
                flags={"read_only": True},
            )
            _element.type = AAZStrType(
                flags={"read_only": True},
            )
            _element.virtual_machine_id = AAZStrType(
                serialized_name="virtualMachineId",
                flags={"read_only": True},
            )

            storage_details = cls._schema_on_200.value.Element.properties.vm_details.Element.storage_details
            storage_details.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.vm_details.Element.storage_details.Element
            _element.id = AAZStrType(
                flags={"read_only": True},
            )

            system_data = cls._schema_on_200.value.Element.system_data
            system_data.created_at = AAZStrType(
                serialized_name="createdAt",
            )
            system_data.created_by = AAZStrType(
                serialized_name="createdBy",
            )
            system_data.created_by_type = AAZStrType(
                serialized_name="createdByType",
            )
            system_data.last_modified_at = AAZStrType(
                serialized_name="lastModifiedAt",
            )
            system_data.last_modified_by = AAZStrType(
                serialized_name="lastModifiedBy",
            )
            system_data.last_modified_by_type = AAZStrType(
                serialized_name="lastModifiedByType",
            )

            tags = cls._schema_on_200.value.Element.tags
            tags.Element = AAZStrType()

            return cls._schema_on_200


class _ListHelper:
    """Helper class for List"""

    _schema_error_definition_read = None

    @classmethod
    def _build_schema_error_definition_read(cls, _schema):
        if cls._schema_error_definition_read is not None:
            _schema.code = cls._schema_error_definition_read.code
            _schema.details = cls._schema_error_definition_read.details
            _schema.message = cls._schema_error_definition_read.message
            return

        cls._schema_error_definition_read = _schema_error_definition_read = AAZObjectType()

        error_definition_read = _schema_error_definition_read
        error_definition_read.code = AAZStrType(
            flags={"read_only": True},
        )
        error_definition_read.details = AAZListType(
            flags={"read_only": True},
        )
        error_definition_read.message = AAZStrType(
            flags={"read_only": True},
        )

        details = _schema_error_definition_read.details
        details.Element = AAZObjectType()
        cls._build_schema_error_definition_read(details.Element)

        _schema.code = cls._schema_error_definition_read.code
        _schema.details = cls._schema_error_definition_read.details
        _schema.message = cls._schema_error_definition_read.message


__all__ = ["List"]
