# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
#
# Code generated by aaz-dev-tools
# --------------------------------------------------------------------------------------------

# pylint: skip-file
# flake8: noqa

from azure.cli.core.aaz import *


class List(AAZCommand):
    """List all VPN connections.

    :example: List all VPN connections in a resource group.
        az network vpn-connection list -g MyResourceGroup
    """

    _aaz_info = {
        "version": "2017-10-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.network/connections", "2017-10-01"],
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
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.VirtualNetworkGatewayConnectionsList(ctx=self.ctx)()
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

    class VirtualNetworkGatewayConnectionsList(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections",
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
                    "subscriptionId", self.ctx.subscription_id,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2017-10-01",
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
                flags={"read_only": True},
            )
            _schema_on_200.value = AAZListType()

            value = cls._schema_on_200.value
            value.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element
            _element.etag = AAZStrType()
            _element.id = AAZStrType()
            _element.location = AAZStrType()
            _element.name = AAZStrType(
                flags={"read_only": True},
            )
            _element.properties = AAZObjectType(
                flags={"required": True, "client_flatten": True},
            )
            _element.tags = AAZDictType()
            _element.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.value.Element.properties
            properties.authorization_key = AAZStrType(
                serialized_name="authorizationKey",
            )
            properties.connection_status = AAZStrType(
                serialized_name="connectionStatus",
                flags={"read_only": True},
            )
            properties.connection_type = AAZStrType(
                serialized_name="connectionType",
                flags={"required": True},
            )
            properties.egress_bytes_transferred = AAZIntType(
                serialized_name="egressBytesTransferred",
                flags={"read_only": True},
            )
            properties.enable_bgp = AAZBoolType(
                serialized_name="enableBgp",
            )
            properties.ingress_bytes_transferred = AAZIntType(
                serialized_name="ingressBytesTransferred",
                flags={"read_only": True},
            )
            properties.ipsec_policies = AAZListType(
                serialized_name="ipsecPolicies",
            )
            properties.local_network_gateway2 = AAZObjectType(
                serialized_name="localNetworkGateway2",
            )
            properties.peer = AAZObjectType()
            _ListHelper._build_schema_sub_resource_read(properties.peer)
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.resource_guid = AAZStrType(
                serialized_name="resourceGuid",
            )
            properties.routing_weight = AAZIntType(
                serialized_name="routingWeight",
            )
            properties.shared_key = AAZStrType(
                serialized_name="sharedKey",
            )
            properties.tunnel_connection_status = AAZListType(
                serialized_name="tunnelConnectionStatus",
                flags={"read_only": True},
            )
            properties.use_policy_based_traffic_selectors = AAZBoolType(
                serialized_name="usePolicyBasedTrafficSelectors",
            )
            properties.virtual_network_gateway1 = AAZObjectType(
                serialized_name="virtualNetworkGateway1",
                flags={"required": True},
            )
            _ListHelper._build_schema_virtual_network_gateway_read(properties.virtual_network_gateway1)
            properties.virtual_network_gateway2 = AAZObjectType(
                serialized_name="virtualNetworkGateway2",
            )
            _ListHelper._build_schema_virtual_network_gateway_read(properties.virtual_network_gateway2)

            ipsec_policies = cls._schema_on_200.value.Element.properties.ipsec_policies
            ipsec_policies.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.ipsec_policies.Element
            _element.dh_group = AAZStrType(
                serialized_name="dhGroup",
                flags={"required": True},
            )
            _element.ike_encryption = AAZStrType(
                serialized_name="ikeEncryption",
                flags={"required": True},
            )
            _element.ike_integrity = AAZStrType(
                serialized_name="ikeIntegrity",
                flags={"required": True},
            )
            _element.ipsec_encryption = AAZStrType(
                serialized_name="ipsecEncryption",
                flags={"required": True},
            )
            _element.ipsec_integrity = AAZStrType(
                serialized_name="ipsecIntegrity",
                flags={"required": True},
            )
            _element.pfs_group = AAZStrType(
                serialized_name="pfsGroup",
                flags={"required": True},
            )
            _element.sa_data_size_kilobytes = AAZIntType(
                serialized_name="saDataSizeKilobytes",
                flags={"required": True},
            )
            _element.sa_life_time_seconds = AAZIntType(
                serialized_name="saLifeTimeSeconds",
                flags={"required": True},
            )

            local_network_gateway2 = cls._schema_on_200.value.Element.properties.local_network_gateway2
            local_network_gateway2.etag = AAZStrType()
            local_network_gateway2.id = AAZStrType()
            local_network_gateway2.location = AAZStrType()
            local_network_gateway2.name = AAZStrType(
                flags={"read_only": True},
            )
            local_network_gateway2.properties = AAZObjectType(
                flags={"required": True, "client_flatten": True},
            )
            local_network_gateway2.tags = AAZDictType()
            local_network_gateway2.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.value.Element.properties.local_network_gateway2.properties
            properties.bgp_settings = AAZObjectType(
                serialized_name="bgpSettings",
            )
            _ListHelper._build_schema_bgp_settings_read(properties.bgp_settings)
            properties.gateway_ip_address = AAZStrType(
                serialized_name="gatewayIpAddress",
            )
            properties.local_network_address_space = AAZObjectType(
                serialized_name="localNetworkAddressSpace",
            )
            _ListHelper._build_schema_address_space_read(properties.local_network_address_space)
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.resource_guid = AAZStrType(
                serialized_name="resourceGuid",
            )

            tags = cls._schema_on_200.value.Element.properties.local_network_gateway2.tags
            tags.Element = AAZStrType()

            tunnel_connection_status = cls._schema_on_200.value.Element.properties.tunnel_connection_status
            tunnel_connection_status.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.tunnel_connection_status.Element
            _element.connection_status = AAZStrType(
                serialized_name="connectionStatus",
                flags={"read_only": True},
            )
            _element.egress_bytes_transferred = AAZIntType(
                serialized_name="egressBytesTransferred",
                flags={"read_only": True},
            )
            _element.ingress_bytes_transferred = AAZIntType(
                serialized_name="ingressBytesTransferred",
                flags={"read_only": True},
            )
            _element.last_connection_established_utc_time = AAZStrType(
                serialized_name="lastConnectionEstablishedUtcTime",
                flags={"read_only": True},
            )
            _element.tunnel = AAZStrType(
                flags={"read_only": True},
            )

            tags = cls._schema_on_200.value.Element.tags
            tags.Element = AAZStrType()

            return cls._schema_on_200


class _ListHelper:
    """Helper class for List"""

    _schema_address_space_read = None

    @classmethod
    def _build_schema_address_space_read(cls, _schema):
        if cls._schema_address_space_read is not None:
            _schema.address_prefixes = cls._schema_address_space_read.address_prefixes
            return

        cls._schema_address_space_read = _schema_address_space_read = AAZObjectType()

        address_space_read = _schema_address_space_read
        address_space_read.address_prefixes = AAZListType(
            serialized_name="addressPrefixes",
        )

        address_prefixes = _schema_address_space_read.address_prefixes
        address_prefixes.Element = AAZStrType()

        _schema.address_prefixes = cls._schema_address_space_read.address_prefixes

    _schema_bgp_settings_read = None

    @classmethod
    def _build_schema_bgp_settings_read(cls, _schema):
        if cls._schema_bgp_settings_read is not None:
            _schema.asn = cls._schema_bgp_settings_read.asn
            _schema.bgp_peering_address = cls._schema_bgp_settings_read.bgp_peering_address
            _schema.peer_weight = cls._schema_bgp_settings_read.peer_weight
            return

        cls._schema_bgp_settings_read = _schema_bgp_settings_read = AAZObjectType()

        bgp_settings_read = _schema_bgp_settings_read
        bgp_settings_read.asn = AAZIntType()
        bgp_settings_read.bgp_peering_address = AAZStrType(
            serialized_name="bgpPeeringAddress",
        )
        bgp_settings_read.peer_weight = AAZIntType(
            serialized_name="peerWeight",
        )

        _schema.asn = cls._schema_bgp_settings_read.asn
        _schema.bgp_peering_address = cls._schema_bgp_settings_read.bgp_peering_address
        _schema.peer_weight = cls._schema_bgp_settings_read.peer_weight

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

    _schema_virtual_network_gateway_read = None

    @classmethod
    def _build_schema_virtual_network_gateway_read(cls, _schema):
        if cls._schema_virtual_network_gateway_read is not None:
            _schema.etag = cls._schema_virtual_network_gateway_read.etag
            _schema.id = cls._schema_virtual_network_gateway_read.id
            _schema.location = cls._schema_virtual_network_gateway_read.location
            _schema.name = cls._schema_virtual_network_gateway_read.name
            _schema.properties = cls._schema_virtual_network_gateway_read.properties
            _schema.tags = cls._schema_virtual_network_gateway_read.tags
            _schema.type = cls._schema_virtual_network_gateway_read.type
            return

        cls._schema_virtual_network_gateway_read = _schema_virtual_network_gateway_read = AAZObjectType()

        virtual_network_gateway_read = _schema_virtual_network_gateway_read
        virtual_network_gateway_read.etag = AAZStrType()
        virtual_network_gateway_read.id = AAZStrType()
        virtual_network_gateway_read.location = AAZStrType()
        virtual_network_gateway_read.name = AAZStrType(
            flags={"read_only": True},
        )
        virtual_network_gateway_read.properties = AAZObjectType(
            flags={"required": True, "client_flatten": True},
        )
        virtual_network_gateway_read.tags = AAZDictType()
        virtual_network_gateway_read.type = AAZStrType(
            flags={"read_only": True},
        )

        properties = _schema_virtual_network_gateway_read.properties
        properties.active_active = AAZBoolType(
            serialized_name="activeActive",
        )
        properties.bgp_settings = AAZObjectType(
            serialized_name="bgpSettings",
        )
        cls._build_schema_bgp_settings_read(properties.bgp_settings)
        properties.enable_bgp = AAZBoolType(
            serialized_name="enableBgp",
        )
        properties.gateway_default_site = AAZObjectType(
            serialized_name="gatewayDefaultSite",
        )
        cls._build_schema_sub_resource_read(properties.gateway_default_site)
        properties.gateway_type = AAZStrType(
            serialized_name="gatewayType",
        )
        properties.ip_configurations = AAZListType(
            serialized_name="ipConfigurations",
        )
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
            flags={"read_only": True},
        )
        properties.resource_guid = AAZStrType(
            serialized_name="resourceGuid",
        )
        properties.sku = AAZObjectType()
        properties.vpn_client_configuration = AAZObjectType(
            serialized_name="vpnClientConfiguration",
        )
        properties.vpn_type = AAZStrType(
            serialized_name="vpnType",
        )

        ip_configurations = _schema_virtual_network_gateway_read.properties.ip_configurations
        ip_configurations.Element = AAZObjectType()

        _element = _schema_virtual_network_gateway_read.properties.ip_configurations.Element
        _element.etag = AAZStrType()
        _element.id = AAZStrType()
        _element.name = AAZStrType()
        _element.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_virtual_network_gateway_read.properties.ip_configurations.Element.properties
        properties.private_ip_allocation_method = AAZStrType(
            serialized_name="privateIPAllocationMethod",
        )
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
            flags={"read_only": True},
        )
        properties.public_ip_address = AAZObjectType(
            serialized_name="publicIPAddress",
        )
        cls._build_schema_sub_resource_read(properties.public_ip_address)
        properties.subnet = AAZObjectType()
        cls._build_schema_sub_resource_read(properties.subnet)

        sku = _schema_virtual_network_gateway_read.properties.sku
        sku.capacity = AAZIntType()
        sku.name = AAZStrType()
        sku.tier = AAZStrType()

        vpn_client_configuration = _schema_virtual_network_gateway_read.properties.vpn_client_configuration
        vpn_client_configuration.radius_server_address = AAZStrType(
            serialized_name="radiusServerAddress",
        )
        vpn_client_configuration.radius_server_secret = AAZStrType(
            serialized_name="radiusServerSecret",
        )
        vpn_client_configuration.vpn_client_address_pool = AAZObjectType(
            serialized_name="vpnClientAddressPool",
        )
        cls._build_schema_address_space_read(vpn_client_configuration.vpn_client_address_pool)
        vpn_client_configuration.vpn_client_protocols = AAZListType(
            serialized_name="vpnClientProtocols",
        )
        vpn_client_configuration.vpn_client_revoked_certificates = AAZListType(
            serialized_name="vpnClientRevokedCertificates",
        )
        vpn_client_configuration.vpn_client_root_certificates = AAZListType(
            serialized_name="vpnClientRootCertificates",
        )

        vpn_client_protocols = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_protocols
        vpn_client_protocols.Element = AAZStrType()

        vpn_client_revoked_certificates = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_revoked_certificates
        vpn_client_revoked_certificates.Element = AAZObjectType()

        _element = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_revoked_certificates.Element
        _element.etag = AAZStrType()
        _element.id = AAZStrType()
        _element.name = AAZStrType()
        _element.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_revoked_certificates.Element.properties
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
            flags={"read_only": True},
        )
        properties.thumbprint = AAZStrType()

        vpn_client_root_certificates = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_root_certificates
        vpn_client_root_certificates.Element = AAZObjectType()

        _element = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_root_certificates.Element
        _element.etag = AAZStrType()
        _element.id = AAZStrType()
        _element.name = AAZStrType()
        _element.properties = AAZObjectType(
            flags={"required": True, "client_flatten": True},
        )

        properties = _schema_virtual_network_gateway_read.properties.vpn_client_configuration.vpn_client_root_certificates.Element.properties
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
            flags={"read_only": True},
        )
        properties.public_cert_data = AAZStrType(
            serialized_name="publicCertData",
            flags={"required": True},
        )

        tags = _schema_virtual_network_gateway_read.tags
        tags.Element = AAZStrType()

        _schema.etag = cls._schema_virtual_network_gateway_read.etag
        _schema.id = cls._schema_virtual_network_gateway_read.id
        _schema.location = cls._schema_virtual_network_gateway_read.location
        _schema.name = cls._schema_virtual_network_gateway_read.name
        _schema.properties = cls._schema_virtual_network_gateway_read.properties
        _schema.tags = cls._schema_virtual_network_gateway_read.tags
        _schema.type = cls._schema_virtual_network_gateway_read.type


__all__ = ["List"]
