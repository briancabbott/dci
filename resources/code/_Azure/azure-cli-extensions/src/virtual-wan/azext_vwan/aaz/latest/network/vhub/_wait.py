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
    "network vhub wait",
)
class Wait(AAZWaitCommand):
    """Place the CLI in a waiting state until a condition is met.
    """

    _aaz_info = {
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.network/virtualhubs/{}", "2022-05-01"],
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
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        _args_schema.name = AAZStrArg(
            options=["-n", "--name"],
            help="Name of the virtual hub.",
            required=True,
            id_part="name",
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.VirtualHubsGet(ctx=self.ctx)()
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

    class VirtualHubsGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}",
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
                    "resourceGroupName", self.ctx.args.resource_group,
                    required=True,
                ),
                **self.serialize_url_param(
                    "subscriptionId", self.ctx.subscription_id,
                    required=True,
                ),
                **self.serialize_url_param(
                    "virtualHubName", self.ctx.args.name,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2022-05-01",
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
            _schema_on_200.id = AAZStrType()
            _schema_on_200.kind = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.location = AAZStrType(
                flags={"required": True},
            )
            _schema_on_200.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200.properties = AAZObjectType(
                flags={"client_flatten": True},
            )
            _schema_on_200.tags = AAZDictType()
            _schema_on_200.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.properties
            properties.address_prefix = AAZStrType(
                serialized_name="addressPrefix",
            )
            properties.allow_branch_to_branch_traffic = AAZBoolType(
                serialized_name="allowBranchToBranchTraffic",
            )
            properties.azure_firewall = AAZObjectType(
                serialized_name="azureFirewall",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.azure_firewall)
            properties.bgp_connections = AAZListType(
                serialized_name="bgpConnections",
                flags={"read_only": True},
            )
            properties.express_route_gateway = AAZObjectType(
                serialized_name="expressRouteGateway",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.express_route_gateway)
            properties.hub_routing_preference = AAZStrType(
                serialized_name="hubRoutingPreference",
            )
            properties.ip_configurations = AAZListType(
                serialized_name="ipConfigurations",
                flags={"read_only": True},
            )
            properties.p2_s_vpn_gateway = AAZObjectType(
                serialized_name="p2SVpnGateway",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.p2_s_vpn_gateway)
            properties.preferred_routing_gateway = AAZStrType(
                serialized_name="preferredRoutingGateway",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.route_maps = AAZListType(
                serialized_name="routeMaps",
                flags={"read_only": True},
            )
            properties.route_table = AAZObjectType(
                serialized_name="routeTable",
            )
            properties.routing_state = AAZStrType(
                serialized_name="routingState",
                flags={"read_only": True},
            )
            properties.security_partner_provider = AAZObjectType(
                serialized_name="securityPartnerProvider",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.security_partner_provider)
            properties.security_provider_name = AAZStrType(
                serialized_name="securityProviderName",
            )
            properties.sku = AAZStrType()
            properties.virtual_hub_route_table_v2s = AAZListType(
                serialized_name="virtualHubRouteTableV2s",
            )
            properties.virtual_router_asn = AAZIntType(
                serialized_name="virtualRouterAsn",
            )
            properties.virtual_router_auto_scale_configuration = AAZObjectType(
                serialized_name="virtualRouterAutoScaleConfiguration",
            )
            properties.virtual_router_ips = AAZListType(
                serialized_name="virtualRouterIps",
            )
            properties.virtual_wan = AAZObjectType(
                serialized_name="virtualWan",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.virtual_wan)
            properties.vpn_gateway = AAZObjectType(
                serialized_name="vpnGateway",
            )
            _WaitHelper._build_schema_sub_resource_read(properties.vpn_gateway)

            bgp_connections = cls._schema_on_200.properties.bgp_connections
            bgp_connections.Element = AAZObjectType()
            _WaitHelper._build_schema_sub_resource_read(bgp_connections.Element)

            ip_configurations = cls._schema_on_200.properties.ip_configurations
            ip_configurations.Element = AAZObjectType()
            _WaitHelper._build_schema_sub_resource_read(ip_configurations.Element)

            route_maps = cls._schema_on_200.properties.route_maps
            route_maps.Element = AAZObjectType()
            _WaitHelper._build_schema_sub_resource_read(route_maps.Element)

            route_table = cls._schema_on_200.properties.route_table
            route_table.routes = AAZListType()

            routes = cls._schema_on_200.properties.route_table.routes
            routes.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.route_table.routes.Element
            _element.address_prefixes = AAZListType(
                serialized_name="addressPrefixes",
            )
            _element.next_hop_ip_address = AAZStrType(
                serialized_name="nextHopIpAddress",
            )

            address_prefixes = cls._schema_on_200.properties.route_table.routes.Element.address_prefixes
            address_prefixes.Element = AAZStrType()

            virtual_hub_route_table_v2s = cls._schema_on_200.properties.virtual_hub_route_table_v2s
            virtual_hub_route_table_v2s.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element
            _element.etag = AAZStrType(
                flags={"read_only": True},
            )
            _element.id = AAZStrType()
            _element.name = AAZStrType()
            _element.properties = AAZObjectType(
                flags={"client_flatten": True},
            )

            properties = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties
            properties.attached_connections = AAZListType(
                serialized_name="attachedConnections",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.routes = AAZListType()

            attached_connections = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties.attached_connections
            attached_connections.Element = AAZStrType()

            routes = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties.routes
            routes.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties.routes.Element
            _element.destination_type = AAZStrType(
                serialized_name="destinationType",
            )
            _element.destinations = AAZListType()
            _element.next_hop_type = AAZStrType(
                serialized_name="nextHopType",
            )
            _element.next_hops = AAZListType(
                serialized_name="nextHops",
            )

            destinations = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties.routes.Element.destinations
            destinations.Element = AAZStrType()

            next_hops = cls._schema_on_200.properties.virtual_hub_route_table_v2s.Element.properties.routes.Element.next_hops
            next_hops.Element = AAZStrType()

            virtual_router_auto_scale_configuration = cls._schema_on_200.properties.virtual_router_auto_scale_configuration
            virtual_router_auto_scale_configuration.min_capacity = AAZIntType(
                serialized_name="minCapacity",
            )

            virtual_router_ips = cls._schema_on_200.properties.virtual_router_ips
            virtual_router_ips.Element = AAZStrType()

            tags = cls._schema_on_200.tags
            tags.Element = AAZStrType()

            return cls._schema_on_200


class _WaitHelper:
    """Helper class for Wait"""

    _schema_sub_resource_read = None

    @classmethod
    def _build_schema_sub_resource_read(cls, _schema):
        if cls._schema_sub_resource_read is not None:
            _schema.id = cls._schema_sub_resource_read.id
            return

        cls._schema_sub_resource_read = _schema_sub_resource_read = AAZObjectType()

        sub_resource_read = _schema_sub_resource_read
        sub_resource_read.id = AAZStrType()

        _schema.id = cls._schema_sub_resource_read.id


__all__ = ["Wait"]
