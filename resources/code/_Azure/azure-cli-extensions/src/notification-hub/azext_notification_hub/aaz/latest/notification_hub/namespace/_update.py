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
    "notification-hub namespace update",
    is_experimental=True,
)
class Update(AAZCommand):
    """Update a service namespace. The namespace's resource manifest is immutable and cannot be modified.

    :example: Update the namespace
        az notification-hub namespace update --resource-group MyResourceGroup --name my-namespace --sku Standard
    """

    _aaz_info = {
        "version": "2017-04-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.notificationhubs/namespaces/{}", "2017-04-01"],
        ]
    }

    AZ_SUPPORT_GENERIC_UPDATE = True

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
        _args_schema.name = AAZStrArg(
            options=["-n", "--name"],
            help="The namespace name.",
            required=True,
            id_part="name",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )

        # define Arg Group "Parameters"

        _args_schema = cls._args_schema
        _args_schema.location = AAZResourceLocationArg(
            arg_group="Parameters",
            help="Resource location",
            fmt=AAZResourceLocationArgFormat(
                resource_group_arg="resource_group",
            ),
        )
        _args_schema.tags = AAZDictArg(
            options=["--tags"],
            arg_group="Parameters",
            help="Resource tags",
            nullable=True,
        )

        tags = cls._args_schema.tags
        tags.Element = AAZStrArg(
            nullable=True,
        )

        # define Arg Group "Properties"

        # define Arg Group "Sku"

        _args_schema = cls._args_schema
        _args_schema.sku = AAZStrArg(
            options=["--sku"],
            arg_group="Sku",
            help="Name of the notification hub sku",
            enum={"Basic": "Basic", "Free": "Free", "Standard": "Standard"},
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.NamespacesGet(ctx=self.ctx)()
        self.pre_instance_update(self.ctx.vars.instance)
        self.InstanceUpdateByJson(ctx=self.ctx)()
        self.InstanceUpdateByGeneric(ctx=self.ctx)()
        self.post_instance_update(self.ctx.vars.instance)
        self.NamespacesCreateOrUpdate(ctx=self.ctx)()
        self.post_operations()

    @register_callback
    def pre_operations(self):
        pass

    @register_callback
    def post_operations(self):
        pass

    @register_callback
    def pre_instance_update(self, instance):
        pass

    @register_callback
    def post_instance_update(self, instance):
        pass

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.vars.instance, client_flatten=True)
        return result

    class NamespacesGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}",
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
                    "namespaceName", self.ctx.args.name,
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
                    "api-version", "2017-04-01",
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
            _UpdateHelper._build_schema_namespace_resource_read(cls._schema_on_200)

            return cls._schema_on_200

    class NamespacesCreateOrUpdate(AAZHttpOperation):
        CLIENT_TYPE = "MgmtClient"

        def __call__(self, *args, **kwargs):
            request = self.make_request()
            session = self.client.send_request(request=request, stream=False, **kwargs)
            if session.http_response.status_code in [200, 201]:
                return self.on_200_201(session)

            return self.on_error(session.http_response)

        @property
        def url(self):
            return self.client.format_url(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}",
                **self.url_parameters
            )

        @property
        def method(self):
            return "PUT"

        @property
        def error_format(self):
            return "MgmtErrorFormat"

        @property
        def url_parameters(self):
            parameters = {
                **self.serialize_url_param(
                    "namespaceName", self.ctx.args.name,
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
                    "api-version", "2017-04-01",
                    required=True,
                ),
            }
            return parameters

        @property
        def header_parameters(self):
            parameters = {
                **self.serialize_header_param(
                    "Content-Type", "application/json",
                ),
                **self.serialize_header_param(
                    "Accept", "application/json",
                ),
            }
            return parameters

        @property
        def content(self):
            _content_value, _builder = self.new_content_builder(
                self.ctx.args,
                value=self.ctx.vars.instance,
            )

            return self.serialize_content(_content_value)

        def on_200_201(self, session):
            data = self.deserialize_http_content(session)
            self.ctx.set_var(
                "instance",
                data,
                schema_builder=self._build_schema_on_200_201
            )

        _schema_on_200_201 = None

        @classmethod
        def _build_schema_on_200_201(cls):
            if cls._schema_on_200_201 is not None:
                return cls._schema_on_200_201

            cls._schema_on_200_201 = AAZObjectType()
            _UpdateHelper._build_schema_namespace_resource_read(cls._schema_on_200_201)

            return cls._schema_on_200_201

    class InstanceUpdateByJson(AAZJsonInstanceUpdateOperation):

        def __call__(self, *args, **kwargs):
            self._update_instance(self.ctx.vars.instance)

        def _update_instance(self, instance):
            _instance_value, _builder = self.new_content_builder(
                self.ctx.args,
                value=instance,
                typ=AAZObjectType
            )
            _builder.set_prop("location", AAZStrType, ".location", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})
            _builder.set_prop("sku", AAZObjectType)
            _builder.set_prop("tags", AAZDictType, ".tags")

            properties = _builder.get(".properties")
            if properties is not None:
                properties.set_prop("name", AAZStrType, ".name")

            sku = _builder.get(".sku")
            if sku is not None:
                sku.set_prop("name", AAZStrType, ".sku", typ_kwargs={"flags": {"required": True}})

            tags = _builder.get(".tags")
            if tags is not None:
                tags.set_elements(AAZStrType, ".")

            return _instance_value

    class InstanceUpdateByGeneric(AAZGenericInstanceUpdateOperation):

        def __call__(self, *args, **kwargs):
            self._update_instance_by_generic(
                self.ctx.vars.instance,
                self.ctx.generic_update_args
            )


class _UpdateHelper:
    """Helper class for Update"""

    _schema_namespace_resource_read = None

    @classmethod
    def _build_schema_namespace_resource_read(cls, _schema):
        if cls._schema_namespace_resource_read is not None:
            _schema.id = cls._schema_namespace_resource_read.id
            _schema.location = cls._schema_namespace_resource_read.location
            _schema.name = cls._schema_namespace_resource_read.name
            _schema.properties = cls._schema_namespace_resource_read.properties
            _schema.sku = cls._schema_namespace_resource_read.sku
            _schema.tags = cls._schema_namespace_resource_read.tags
            _schema.type = cls._schema_namespace_resource_read.type
            return

        cls._schema_namespace_resource_read = _schema_namespace_resource_read = AAZObjectType()

        namespace_resource_read = _schema_namespace_resource_read
        namespace_resource_read.id = AAZStrType(
            flags={"read_only": True},
        )
        namespace_resource_read.location = AAZStrType()
        namespace_resource_read.name = AAZStrType(
            flags={"read_only": True},
        )
        namespace_resource_read.properties = AAZObjectType(
            flags={"client_flatten": True},
        )
        namespace_resource_read.sku = AAZObjectType()
        namespace_resource_read.tags = AAZDictType()
        namespace_resource_read.type = AAZStrType(
            flags={"read_only": True},
        )

        properties = _schema_namespace_resource_read.properties
        properties.created_at = AAZStrType(
            serialized_name="createdAt",
        )
        properties.critical = AAZBoolType()
        properties.data_center = AAZStrType(
            serialized_name="dataCenter",
        )
        properties.enabled = AAZBoolType()
        properties.metric_id = AAZStrType(
            serialized_name="metricId",
            flags={"read_only": True},
        )
        properties.name = AAZStrType()
        properties.namespace_type = AAZStrType(
            serialized_name="namespaceType",
        )
        properties.provisioning_state = AAZStrType(
            serialized_name="provisioningState",
        )
        properties.region = AAZStrType()
        properties.scale_unit = AAZStrType(
            serialized_name="scaleUnit",
        )
        properties.service_bus_endpoint = AAZStrType(
            serialized_name="serviceBusEndpoint",
        )
        properties.status = AAZStrType()
        properties.subscription_id = AAZStrType(
            serialized_name="subscriptionId",
        )
        properties.updated_at = AAZStrType(
            serialized_name="updatedAt",
        )

        sku = _schema_namespace_resource_read.sku
        sku.capacity = AAZIntType()
        sku.family = AAZStrType()
        sku.name = AAZStrType(
            flags={"required": True},
        )
        sku.size = AAZStrType()
        sku.tier = AAZStrType()

        tags = _schema_namespace_resource_read.tags
        tags.Element = AAZStrType()

        _schema.id = cls._schema_namespace_resource_read.id
        _schema.location = cls._schema_namespace_resource_read.location
        _schema.name = cls._schema_namespace_resource_read.name
        _schema.properties = cls._schema_namespace_resource_read.properties
        _schema.sku = cls._schema_namespace_resource_read.sku
        _schema.tags = cls._schema_namespace_resource_read.tags
        _schema.type = cls._schema_namespace_resource_read.type


__all__ = ["Update"]
