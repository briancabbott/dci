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
    "networkfabric externalnetwork wait",
)
class Wait(AAZWaitCommand):
    """Place the CLI in a waiting state until a condition is met.
    """

    _aaz_info = {
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.managednetworkfabric/l3isolationdomains/{}/externalnetworks/{}", "2023-06-15"],
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
        _args_schema.resource_name = AAZStrArg(
            options=["--resource-name"],
            help="Name of the External Network.",
            required=True,
            id_part="child_name_1",
        )
        _args_schema.l3_isolation_domain_name = AAZStrArg(
            options=["--l3domain", "--l3-isolation-domain-name"],
            help="Name of the L3 Isolation Domain.",
            required=True,
            id_part="name",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            help="Name of the resource group",
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.ExternalNetworksGet(ctx=self.ctx)()
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

    class ExternalNetworksGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/{l3IsolationDomainName}/externalNetworks/{externalNetworkName}",
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
                    "externalNetworkName", self.ctx.args.resource_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "l3IsolationDomainName", self.ctx.args.l3_isolation_domain_name,
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
                    "api-version", "2023-06-15",
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
            _schema_on_200.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.properties = AAZObjectType(
                flags={"required": True, "client_flatten": True},
            )
            _schema_on_200.system_data = AAZObjectType(
                serialized_name="systemData",
                flags={"read_only": True},
            )
            _schema_on_200.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.properties
            properties.administrative_state = AAZStrType(
                serialized_name="administrativeState",
                flags={"read_only": True},
            )
            properties.annotation = AAZStrType()
            properties.configuration_state = AAZStrType(
                serialized_name="configurationState",
                flags={"read_only": True},
            )
            properties.export_route_policy = AAZObjectType(
                serialized_name="exportRoutePolicy",
            )
            properties.export_route_policy_id = AAZStrType(
                serialized_name="exportRoutePolicyId",
            )
            properties.import_route_policy = AAZObjectType(
                serialized_name="importRoutePolicy",
            )
            properties.import_route_policy_id = AAZStrType(
                serialized_name="importRoutePolicyId",
            )
            properties.network_to_network_interconnect_id = AAZStrType(
                serialized_name="networkToNetworkInterconnectId",
                flags={"read_only": True},
            )
            properties.option_a_properties = AAZObjectType(
                serialized_name="optionAProperties",
            )
            properties.option_b_properties = AAZObjectType(
                serialized_name="optionBProperties",
            )
            properties.peering_option = AAZStrType(
                serialized_name="peeringOption",
                flags={"required": True},
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )

            export_route_policy = cls._schema_on_200.properties.export_route_policy
            export_route_policy.export_ipv4_route_policy_id = AAZStrType(
                serialized_name="exportIpv4RoutePolicyId",
            )
            export_route_policy.export_ipv6_route_policy_id = AAZStrType(
                serialized_name="exportIpv6RoutePolicyId",
            )

            import_route_policy = cls._schema_on_200.properties.import_route_policy
            import_route_policy.import_ipv4_route_policy_id = AAZStrType(
                serialized_name="importIpv4RoutePolicyId",
            )
            import_route_policy.import_ipv6_route_policy_id = AAZStrType(
                serialized_name="importIpv6RoutePolicyId",
            )

            option_a_properties = cls._schema_on_200.properties.option_a_properties
            option_a_properties.bfd_configuration = AAZObjectType(
                serialized_name="bfdConfiguration",
            )
            option_a_properties.egress_acl_id = AAZStrType(
                serialized_name="egressAclId",
            )
            option_a_properties.fabric_asn = AAZIntType(
                serialized_name="fabricASN",
                flags={"read_only": True},
            )
            option_a_properties.ingress_acl_id = AAZStrType(
                serialized_name="ingressAclId",
            )
            option_a_properties.mtu = AAZIntType()
            option_a_properties.peer_asn = AAZIntType(
                serialized_name="peerASN",
                flags={"required": True},
            )
            option_a_properties.primary_ipv4_prefix = AAZStrType(
                serialized_name="primaryIpv4Prefix",
            )
            option_a_properties.primary_ipv6_prefix = AAZStrType(
                serialized_name="primaryIpv6Prefix",
                nullable=True,
            )
            option_a_properties.secondary_ipv4_prefix = AAZStrType(
                serialized_name="secondaryIpv4Prefix",
            )
            option_a_properties.secondary_ipv6_prefix = AAZStrType(
                serialized_name="secondaryIpv6Prefix",
                nullable=True,
            )
            option_a_properties.vlan_id = AAZIntType(
                serialized_name="vlanId",
                flags={"required": True},
            )

            bfd_configuration = cls._schema_on_200.properties.option_a_properties.bfd_configuration
            bfd_configuration.administrative_state = AAZStrType(
                serialized_name="administrativeState",
                flags={"read_only": True},
            )
            bfd_configuration.interval_in_milli_seconds = AAZIntType(
                serialized_name="intervalInMilliSeconds",
            )
            bfd_configuration.multiplier = AAZIntType()

            option_b_properties = cls._schema_on_200.properties.option_b_properties
            option_b_properties.export_route_targets = AAZListType(
                serialized_name="exportRouteTargets",
            )
            option_b_properties.import_route_targets = AAZListType(
                serialized_name="importRouteTargets",
            )
            option_b_properties.route_targets = AAZObjectType(
                serialized_name="routeTargets",
            )

            export_route_targets = cls._schema_on_200.properties.option_b_properties.export_route_targets
            export_route_targets.Element = AAZStrType()

            import_route_targets = cls._schema_on_200.properties.option_b_properties.import_route_targets
            import_route_targets.Element = AAZStrType()

            route_targets = cls._schema_on_200.properties.option_b_properties.route_targets
            route_targets.export_ipv4_route_targets = AAZListType(
                serialized_name="exportIpv4RouteTargets",
            )
            route_targets.export_ipv6_route_targets = AAZListType(
                serialized_name="exportIpv6RouteTargets",
            )
            route_targets.import_ipv4_route_targets = AAZListType(
                serialized_name="importIpv4RouteTargets",
            )
            route_targets.import_ipv6_route_targets = AAZListType(
                serialized_name="importIpv6RouteTargets",
            )

            export_ipv4_route_targets = cls._schema_on_200.properties.option_b_properties.route_targets.export_ipv4_route_targets
            export_ipv4_route_targets.Element = AAZStrType()

            export_ipv6_route_targets = cls._schema_on_200.properties.option_b_properties.route_targets.export_ipv6_route_targets
            export_ipv6_route_targets.Element = AAZStrType()

            import_ipv4_route_targets = cls._schema_on_200.properties.option_b_properties.route_targets.import_ipv4_route_targets
            import_ipv4_route_targets.Element = AAZStrType()

            import_ipv6_route_targets = cls._schema_on_200.properties.option_b_properties.route_targets.import_ipv6_route_targets
            import_ipv6_route_targets.Element = AAZStrType()

            system_data = cls._schema_on_200.system_data
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

            return cls._schema_on_200


class _WaitHelper:
    """Helper class for Wait"""


__all__ = ["Wait"]
