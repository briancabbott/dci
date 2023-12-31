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
    "network vhub route-map rule show",
)
class Show(AAZCommand):
    """Show route map rule

    :example: Show route map rule
        az network vhub route-map rule show -g rg --route-map-name routemap-name --vhub-name vhub-name --rule-index 0
    """

    _aaz_info = {
        "version": "2022-05-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.network/virtualhubs/{}/routemaps/{}", "2022-05-01", "properties.rules[]"],
        ]
    }

    def _handler(self, command_args):
        super()._handler(command_args)
        self.SubresourceSelector(ctx=self.ctx, name="subresource")
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
        _args_schema.route_map_name = AAZStrArg(
            options=["--route-map-name"],
            help="The name of the RouteMap.",
            required=True,
        )
        _args_schema.vhub_name = AAZStrArg(
            options=["--vhub-name"],
            help="The name of the VirtualHub containing the RouteMap.",
            required=True,
        )
        _args_schema.rule_index = AAZIntArg(
            options=["--rule-index"],
            help="The index of the route map rule",
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.RouteMapsGet(ctx=self.ctx)()
        self.post_operations()

    @register_callback
    def pre_operations(self):
        pass

    @register_callback
    def post_operations(self):
        pass

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.selectors.subresource.required(), client_flatten=True)
        return result

    class SubresourceSelector(AAZJsonSelector):

        def _get(self):
            result = self.ctx.vars.instance
            result = result.properties.rules
            filters = enumerate(result)
            filters = filter(
                lambda e: e[0] == self.ctx.args.rule_index,
                filters
            )
            idx = next(filters)[0]
            return result[idx]

        def _set(self, value):
            result = self.ctx.vars.instance
            result = result.properties.rules
            filters = enumerate(result)
            filters = filter(
                lambda e: e[0] == self.ctx.args.rule_index,
                filters
            )
            idx = next(filters, [len(result)])[0]
            self.ctx.args.rule_index = idx
            result[idx] = value
            return

    class RouteMapsGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}",
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
                    "routeMapName", self.ctx.args.route_map_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "subscriptionId", self.ctx.subscription_id,
                    required=True,
                ),
                **self.serialize_url_param(
                    "virtualHubName", self.ctx.args.vhub_name,
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
            _ShowHelper._build_schema_route_map_read(cls._schema_on_200)

            return cls._schema_on_200


class _ShowHelper:
    """Helper class for Show"""

    _schema_route_map_read = None

    @classmethod
    def _build_schema_route_map_read(cls, _schema):
        if cls._schema_route_map_read is not None:
            _schema.etag = cls._schema_route_map_read.etag
            _schema.id = cls._schema_route_map_read.id
            _schema.name = cls._schema_route_map_read.name
            _schema.properties = cls._schema_route_map_read.properties
            _schema.type = cls._schema_route_map_read.type
            return

        cls._schema_route_map_read = _schema_route_map_read = AAZObjectType()

        route_map_read = _schema_route_map_read
        route_map_read.etag = AAZStrType(
            flags={"read_only": True},
        )
        route_map_read.id = AAZStrType(
            flags={"read_only": True},
        )
        route_map_read.name = AAZStrType(
            flags={"read_only": True},
        )
        route_map_read.properties = AAZObjectType(
            flags={"client_flatten": True},
        )
        route_map_read.type = AAZStrType(
            flags={"read_only": True},
        )

        properties = _schema_route_map_read.properties
        properties.associated_inbound_connections = AAZListType(
            serialized_name="associatedInboundConnections",
        )
        properties.associated_outbound_connections = AAZListType(
            serialized_name="associatedOutboundConnections",
        )
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
            flags={"read_only": True},
        )
        properties.rules = AAZListType()

        associated_inbound_connections = _schema_route_map_read.properties.associated_inbound_connections
        associated_inbound_connections.Element = AAZStrType()

        associated_outbound_connections = _schema_route_map_read.properties.associated_outbound_connections
        associated_outbound_connections.Element = AAZStrType()

        rules = _schema_route_map_read.properties.rules
        rules.Element = AAZObjectType()

        _element = _schema_route_map_read.properties.rules.Element
        _element.actions = AAZListType()
        _element.match_criteria = AAZListType(
            serialized_name="matchCriteria",
        )
        _element.name = AAZStrType()
        _element.next_step_if_matched = AAZStrType(
            serialized_name="nextStepIfMatched",
        )

        actions = _schema_route_map_read.properties.rules.Element.actions
        actions.Element = AAZObjectType()

        _element = _schema_route_map_read.properties.rules.Element.actions.Element
        _element.parameters = AAZListType()
        _element.type = AAZStrType()

        parameters = _schema_route_map_read.properties.rules.Element.actions.Element.parameters
        parameters.Element = AAZObjectType()

        _element = _schema_route_map_read.properties.rules.Element.actions.Element.parameters.Element
        _element.as_path = AAZListType(
            serialized_name="asPath",
        )
        _element.community = AAZListType()
        _element.route_prefix = AAZListType(
            serialized_name="routePrefix",
        )

        as_path = _schema_route_map_read.properties.rules.Element.actions.Element.parameters.Element.as_path
        as_path.Element = AAZStrType()

        community = _schema_route_map_read.properties.rules.Element.actions.Element.parameters.Element.community
        community.Element = AAZStrType()

        route_prefix = _schema_route_map_read.properties.rules.Element.actions.Element.parameters.Element.route_prefix
        route_prefix.Element = AAZStrType()

        match_criteria = _schema_route_map_read.properties.rules.Element.match_criteria
        match_criteria.Element = AAZObjectType()

        _element = _schema_route_map_read.properties.rules.Element.match_criteria.Element
        _element.as_path = AAZListType(
            serialized_name="asPath",
        )
        _element.community = AAZListType()
        _element.match_condition = AAZStrType(
            serialized_name="matchCondition",
        )
        _element.route_prefix = AAZListType(
            serialized_name="routePrefix",
        )

        as_path = _schema_route_map_read.properties.rules.Element.match_criteria.Element.as_path
        as_path.Element = AAZStrType()

        community = _schema_route_map_read.properties.rules.Element.match_criteria.Element.community
        community.Element = AAZStrType()

        route_prefix = _schema_route_map_read.properties.rules.Element.match_criteria.Element.route_prefix
        route_prefix.Element = AAZStrType()

        _schema.etag = cls._schema_route_map_read.etag
        _schema.id = cls._schema_route_map_read.id
        _schema.name = cls._schema_route_map_read.name
        _schema.properties = cls._schema_route_map_read.properties
        _schema.type = cls._schema_route_map_read.type


__all__ = ["Show"]
