# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
#
# Code generated by aaz-dev-tools
# --------------------------------------------------------------------------------------------

# pylint: skip-file
# flake8: noqa

from azure.cli.core.aaz import *


class Create(AAZCommand):
    """Create a datastore in a private cloud cluster
    """

    _aaz_info = {
        "version": "2023-03-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.avs/privateclouds/{}/clusters/{}/datastores/{}", "2023-03-01"],
        ]
    }

    AZ_SUPPORT_NO_WAIT = True

    def _handler(self, command_args):
        super()._handler(command_args)
        return self.build_lro_poller(self._execute_operations, self._output)

    _args_schema = None

    @classmethod
    def _build_arguments_schema(cls, *args, **kwargs):
        if cls._args_schema is not None:
            return cls._args_schema
        cls._args_schema = super()._build_arguments_schema(*args, **kwargs)

        # define Arg Group ""

        _args_schema = cls._args_schema
        _args_schema.cluster_name = AAZStrArg(
            options=["--cluster", "--cluster-name"],
            help="Name of the cluster in the private cloud",
            required=True,
            fmt=AAZStrArgFormat(
                pattern="^[-\w\._]+$",
            ),
        )
        _args_schema.datastore_name = AAZStrArg(
            options=["-n", "--name", "--datastore-name"],
            help="Name of the datastore in the private cloud cluster",
            required=True,
            fmt=AAZStrArgFormat(
                pattern="^[-\w\._]+$",
            ),
        )
        _args_schema.private_cloud = AAZStrArg(
            options=["-c", "--private-cloud"],
            help="Name of the private cloud",
            required=True,
            fmt=AAZStrArgFormat(
                pattern="^[-\w\._]+$",
            ),
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )

        # define Arg Group "DiskPoolVolume"

        _args_schema = cls._args_schema
        _args_schema.lun_name = AAZStrArg(
            options=["--lun-name"],
            arg_group="DiskPoolVolume",
            help="Name of the LUN to be used for datastore",
        )
        _args_schema.mount_option = AAZStrArg(
            options=["--mount-option"],
            arg_group="DiskPoolVolume",
            help="Mode that describes whether the LUN has to be mounted as a datastore or attached as a LUN",
            default="MOUNT",
            enum={"ATTACH": "ATTACH", "MOUNT": "MOUNT"},
        )
        _args_schema.target_id = AAZStrArg(
            options=["--target-id"],
            arg_group="DiskPoolVolume",
            help="Azure resource ID of the iSCSI target",
        )

        # define Arg Group "NetAppVolume"

        _args_schema = cls._args_schema
        _args_schema.net_app_volumn = AAZStrArg(
            options=["--volume-id", "--net-app-volumn"],
            arg_group="NetAppVolume",
            help="Azure resource ID of the NetApp volume",
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        yield self.DatastoresCreateOrUpdate(ctx=self.ctx)()
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

    class DatastoresCreateOrUpdate(AAZHttpOperation):
        CLIENT_TYPE = "MgmtClient"

        def __call__(self, *args, **kwargs):
            request = self.make_request()
            session = self.client.send_request(request=request, stream=False, **kwargs)
            if session.http_response.status_code in [202]:
                return self.client.build_lro_polling(
                    self.ctx.args.no_wait,
                    session,
                    self.on_200_201,
                    self.on_error,
                    lro_options={"final-state-via": "azure-async-operation"},
                    path_format_arguments=self.url_parameters,
                )
            if session.http_response.status_code in [200, 201]:
                return self.client.build_lro_polling(
                    self.ctx.args.no_wait,
                    session,
                    self.on_200_201,
                    self.on_error,
                    lro_options={"final-state-via": "azure-async-operation"},
                    path_format_arguments=self.url_parameters,
                )

            return self.on_error(session.http_response)

        @property
        def url(self):
            return self.client.format_url(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/clusters/{clusterName}/datastores/{datastoreName}",
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
                    "clusterName", self.ctx.args.cluster_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "datastoreName", self.ctx.args.datastore_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "privateCloudName", self.ctx.args.private_cloud,
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
                    "api-version", "2023-03-01",
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
                typ=AAZObjectType,
                typ_kwargs={"flags": {"required": True, "client_flatten": True}}
            )
            _builder.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties")
            if properties is not None:
                properties.set_prop("diskPoolVolume", AAZObjectType)
                properties.set_prop("netAppVolume", AAZObjectType)

            disk_pool_volume = _builder.get(".properties.diskPoolVolume")
            if disk_pool_volume is not None:
                disk_pool_volume.set_prop("lunName", AAZStrType, ".lun_name", typ_kwargs={"flags": {"required": True}})
                disk_pool_volume.set_prop("mountOption", AAZStrType, ".mount_option")
                disk_pool_volume.set_prop("targetId", AAZStrType, ".target_id", typ_kwargs={"flags": {"required": True}})

            net_app_volume = _builder.get(".properties.netAppVolume")
            if net_app_volume is not None:
                net_app_volume.set_prop("id", AAZStrType, ".net_app_volumn", typ_kwargs={"flags": {"required": True}})

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

            _schema_on_200_201 = cls._schema_on_200_201
            _schema_on_200_201.id = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200_201.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200_201.properties = AAZObjectType(
                flags={"client_flatten": True},
            )
            _schema_on_200_201.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200_201.properties
            properties.disk_pool_volume = AAZObjectType(
                serialized_name="diskPoolVolume",
            )
            properties.net_app_volume = AAZObjectType(
                serialized_name="netAppVolume",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.status = AAZStrType(
                flags={"read_only": True},
            )

            disk_pool_volume = cls._schema_on_200_201.properties.disk_pool_volume
            disk_pool_volume.lun_name = AAZStrType(
                serialized_name="lunName",
                flags={"required": True},
            )
            disk_pool_volume.mount_option = AAZStrType(
                serialized_name="mountOption",
            )
            disk_pool_volume.path = AAZStrType(
                flags={"read_only": True},
            )
            disk_pool_volume.target_id = AAZStrType(
                serialized_name="targetId",
                flags={"required": True},
            )

            net_app_volume = cls._schema_on_200_201.properties.net_app_volume
            net_app_volume.id = AAZStrType(
                flags={"required": True},
            )

            return cls._schema_on_200_201


class _CreateHelper:
    """Helper class for Create"""


__all__ = ["Create"]
