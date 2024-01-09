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
    "workloads sap-sizing-recommendation",
    is_preview=True,
)
class SapSizingRecommendation(AAZCommand):
    """Show SAP sizing recommendations by providing input SAPS for application tier and memory required for database tier

    :example: Get sizing recommendations for a Distributed with High Availability (HA) system by providing SAPS for application tier and memory required for database tier of the SAP system
        az workloads sap-sizing-recommendation --app-location "eastus" --database-type "HANA" --db-memory 1024 --deployment-type "ThreeTier" --environment "Prod" --high-availability-type "AvailabilitySet" --sap-product "S4HANA" --saps 75000 --location "eastus2" --db-scale-method ScaleUp

    :example: Get sizing recommendations for a Distributed system by providing SAPS for application tier and memory required for database tier of the SAP system
        az workloads sap-sizing-recommendation --app-location "eastus" --database-type "HANA" --db-memory 1024 --deployment-type "ThreeTier" --environment "Prod" --sap-product "S4HANA" --saps 20000 --location "northeurope" --db-scale-method ScaleUp
    """

    _aaz_info = {
        "version": "2023-04-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/providers/microsoft.workloads/locations/{}/sapvirtualinstancemetadata/default/getsizingrecommendations", "2023-04-01"],
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
        _args_schema.location = AAZResourceLocationArg(
            required=True,
            id_part="name",
        )

        # define Arg Group "SAPSizingRecommendation"

        _args_schema = cls._args_schema
        _args_schema.app_location = AAZStrArg(
            options=["--app-location"],
            arg_group="SAPSizingRecommendation",
            help="The geo-location where the resource is to be created.",
        )
        _args_schema.database_type = AAZStrArg(
            options=["--database-type"],
            arg_group="SAPSizingRecommendation",
            help="The database type.",
            enum={"DB2": "DB2", "HANA": "HANA"},
        )
        _args_schema.db_memory = AAZIntArg(
            options=["--db-memory"],
            arg_group="SAPSizingRecommendation",
            help="The database memory configuration.",
        )
        _args_schema.db_scale_method = AAZStrArg(
            options=["--db-scale-method"],
            arg_group="SAPSizingRecommendation",
            help="The DB scale method.",
            enum={"ScaleUp": "ScaleUp"},
        )
        _args_schema.deployment_type = AAZStrArg(
            options=["--deployment-type"],
            arg_group="SAPSizingRecommendation",
            help="The deployment type. Eg: SingleServer/ThreeTier",
            enum={"SingleServer": "SingleServer", "ThreeTier": "ThreeTier"},
        )
        _args_schema.environment = AAZStrArg(
            options=["--environment"],
            arg_group="SAPSizingRecommendation",
            help="Defines the environment type - Production/Non Production.",
            enum={"NonProd": "NonProd", "Prod": "Prod"},
        )
        _args_schema.high_availability_type = AAZStrArg(
            options=["--high-availability-type"],
            arg_group="SAPSizingRecommendation",
            help="The high availability type.",
            enum={"AvailabilitySet": "AvailabilitySet", "AvailabilityZone": "AvailabilityZone"},
        )
        _args_schema.sap_product = AAZStrArg(
            options=["--sap-product"],
            arg_group="SAPSizingRecommendation",
            help="Defines the SAP Product type.",
            enum={"ECC": "ECC", "Other": "Other", "S4HANA": "S4HANA"},
        )
        _args_schema.saps = AAZIntArg(
            options=["--saps"],
            arg_group="SAPSizingRecommendation",
            help="The SAP Application Performance Standard measurement.",
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.SAPSizingRecommendations(ctx=self.ctx)()
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

    class SAPSizingRecommendations(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/providers/Microsoft.Workloads/locations/{location}/sapVirtualInstanceMetadata/default/getSizingRecommendations",
                **self.url_parameters
            )

        @property
        def method(self):
            return "POST"

        @property
        def error_format(self):
            return "MgmtErrorFormat"

        @property
        def url_parameters(self):
            parameters = {
                **self.serialize_url_param(
                    "location", self.ctx.args.location,
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
                typ_kwargs={"flags": {"client_flatten": True}}
            )
            _builder.set_prop("appLocation", AAZStrType, ".app_location", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("databaseType", AAZStrType, ".database_type", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("dbMemory", AAZIntType, ".db_memory", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("dbScaleMethod", AAZStrType, ".db_scale_method")
            _builder.set_prop("deploymentType", AAZStrType, ".deployment_type", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("environment", AAZStrType, ".environment", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("highAvailabilityType", AAZStrType, ".high_availability_type")
            _builder.set_prop("sapProduct", AAZStrType, ".sap_product", typ_kwargs={"flags": {"required": True}})
            _builder.set_prop("saps", AAZIntType, ".saps", typ_kwargs={"flags": {"required": True}})

            return self.serialize_content(_content_value)

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
            _schema_on_200.deployment_type = AAZStrType(
                serialized_name="deploymentType",
                flags={"required": True},
            )

            disc_single_server = cls._schema_on_200.discriminate_by("deployment_type", "SingleServer")
            disc_single_server.vm_sku = AAZStrType(
                serialized_name="vmSku",
            )

            disc_three_tier = cls._schema_on_200.discriminate_by("deployment_type", "ThreeTier")
            disc_three_tier.application_server_instance_count = AAZIntType(
                serialized_name="applicationServerInstanceCount",
            )
            disc_three_tier.application_server_vm_sku = AAZStrType(
                serialized_name="applicationServerVmSku",
            )
            disc_three_tier.central_server_instance_count = AAZIntType(
                serialized_name="centralServerInstanceCount",
            )
            disc_three_tier.central_server_vm_sku = AAZStrType(
                serialized_name="centralServerVmSku",
            )
            disc_three_tier.database_instance_count = AAZIntType(
                serialized_name="databaseInstanceCount",
            )
            disc_three_tier.db_vm_sku = AAZStrType(
                serialized_name="dbVmSku",
            )

            return cls._schema_on_200


class _SapSizingRecommendationHelper:
    """Helper class for SapSizingRecommendation"""


__all__ = ["SapSizingRecommendation"]
