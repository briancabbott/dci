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
    "network perimeter link wait",
)
class Wait(AAZWaitCommand):
    """Place the CLI in a waiting state until a condition is met.
    """

    _aaz_info = {
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.network/networksecurityperimeters/{}/links/{}", "2023-07-01-preview"],
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
        _args_schema.link_name = AAZStrArg(
            options=["-n", "--name", "--link-name"],
            help="The name of the NSP link.",
            required=True,
            id_part="child_name_1",
        )
        _args_schema.perimeter_name = AAZStrArg(
            options=["--perimeter-name"],
            help="The name of the network security perimeter.",
            required=True,
            id_part="name",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.NspLinksGet(ctx=self.ctx)()
        self.post_operations()

    @register_callback
    def pre_operations(self):
        pass

    @register_callback
    def post_operations(self):
        pass

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.vars.instance, client_flatten=False)
        return result

    class NspLinksGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityPerimeters/{networkSecurityPerimeterName}/links/{linkName}",
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
                    "linkName", self.ctx.args.link_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "networkSecurityPerimeterName", self.ctx.args.perimeter_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "resourceGroupName", self.ctx.args.resource_group,
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
                    "api-version", "2023-07-01-preview",
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
            _schema_on_200.etag = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.id = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.properties = AAZObjectType()
            _schema_on_200.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.properties
            properties.auto_approved_remote_perimeter_resource_id = AAZStrType(
                serialized_name="autoApprovedRemotePerimeterResourceId",
            )
            properties.description = AAZStrType()
            properties.local_inbound_profiles = AAZListType(
                serialized_name="localInboundProfiles",
            )
            properties.local_outbound_profiles = AAZListType(
                serialized_name="localOutboundProfiles",
                flags={"read_only": True},
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.remote_inbound_profiles = AAZListType(
                serialized_name="remoteInboundProfiles",
            )
            properties.remote_outbound_profiles = AAZListType(
                serialized_name="remoteOutboundProfiles",
                flags={"read_only": True},
            )
            properties.remote_perimeter_guid = AAZStrType(
                serialized_name="remotePerimeterGuid",
                flags={"read_only": True},
            )
            properties.remote_perimeter_location = AAZStrType(
                serialized_name="remotePerimeterLocation",
                flags={"read_only": True},
            )
            properties.status = AAZStrType(
                flags={"read_only": True},
            )

            local_inbound_profiles = cls._schema_on_200.properties.local_inbound_profiles
            local_inbound_profiles.Element = AAZStrType()

            local_outbound_profiles = cls._schema_on_200.properties.local_outbound_profiles
            local_outbound_profiles.Element = AAZStrType()

            remote_inbound_profiles = cls._schema_on_200.properties.remote_inbound_profiles
            remote_inbound_profiles.Element = AAZStrType()

            remote_outbound_profiles = cls._schema_on_200.properties.remote_outbound_profiles
            remote_outbound_profiles.Element = AAZStrType()

            return cls._schema_on_200


class _WaitHelper:
    """Helper class for Wait"""


__all__ = ["Wait"]
