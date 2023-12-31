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
    "billing-benefits reservation-order-aliases create",
)
class Create(AAZCommand):
    """Create a reservation order alias.

    :example: Create a Reservation
        az billing-benefits reservation-order-aliases create --order-alias-name TestRO --location westus --applied-scope-type Single --applied-scope-prop "{subscription-id:/subscriptions/30000000-aaaa-bbbb-cccc-200000000004}" --billing-plan P1M --billing-scope-id /subscriptions/30000000-aaaa-bbbb-cccc-200000000004 --display-name TestRO --quantity 1 --renew false --reserved-resource-type VirtualMachines --sku Standard_B1ls  --term P1Y --instance-flexibility On
    """

    _aaz_info = {
        "version": "2022-11-01",
        "resources": [
            ["mgmt-plane", "/providers/microsoft.billingbenefits/reservationorderaliases/{}", "2022-11-01"],
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
        _args_schema.order_alias_name = AAZStrArg(
            options=["--order-alias-name"],
            help="Name of the reservation order alias",
            required=True,
            fmt=AAZStrArgFormat(
                pattern="^[a-zA-Z0-9_\-\.]+$",
            ),
        )
        _args_schema.sku = AAZStrArg(
            options=["--sku"],
            help="Name of the SKU to be applied",
        )

        # define Arg Group "Body"

        _args_schema = cls._args_schema
        _args_schema.location = AAZResourceLocationArg(
            arg_group="Body",
            help="The Azure Region where the reservation benefits are applied to.",
        )

        # define Arg Group "Properties"

        _args_schema = cls._args_schema
        _args_schema.applied_scope_prop = AAZObjectArg(
            options=["--applied-scope-prop"],
            arg_group="Properties",
            help="Properties specific to applied scope type. Not required if not applicable.",
        )
        _args_schema.applied_scope_type = AAZStrArg(
            options=["--applied-scope-type"],
            arg_group="Properties",
            help="Type of the Applied Scope.",
            enum={"ManagementGroup": "ManagementGroup", "Shared": "Shared", "Single": "Single"},
        )
        _args_schema.billing_plan = AAZStrArg(
            options=["--billing-plan"],
            arg_group="Properties",
            help="Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.",
            enum={"P1M": "P1M"},
        )
        _args_schema.billing_scope_id = AAZStrArg(
            options=["--billing-scope-id"],
            arg_group="Properties",
            help="Subscription that will be charged for purchasing the benefit",
        )
        _args_schema.display_name = AAZStrArg(
            options=["--display-name"],
            arg_group="Properties",
            help="Display name",
        )
        _args_schema.quantity = AAZIntArg(
            options=["--quantity"],
            arg_group="Properties",
            help="Total Quantity of the SKUs purchased in the Reservation.",
            fmt=AAZIntArgFormat(
                minimum=1,
            ),
        )
        _args_schema.renew = AAZBoolArg(
            options=["--renew"],
            arg_group="Properties",
            help="Setting this to true will automatically purchase a new benefit on the expiration date time.",
            default=False,
        )
        _args_schema.reserved_resource_type = AAZStrArg(
            options=["--reserved-resource-type"],
            arg_group="Properties",
            help="The type of the resource that is being reserved.",
            enum={"AVS": "AVS", "AppService": "AppService", "AzureDataExplorer": "AzureDataExplorer", "AzureFiles": "AzureFiles", "BlockBlob": "BlockBlob", "CosmosDb": "CosmosDb", "DataFactory": "DataFactory", "Databricks": "Databricks", "DedicatedHost": "DedicatedHost", "ManagedDisk": "ManagedDisk", "MariaDb": "MariaDb", "MySql": "MySql", "NetAppStorage": "NetAppStorage", "PostgreSql": "PostgreSql", "RedHat": "RedHat", "RedHatOsa": "RedHatOsa", "RedisCache": "RedisCache", "SapHana": "SapHana", "SqlAzureHybridBenefit": "SqlAzureHybridBenefit", "SqlDataWarehouse": "SqlDataWarehouse", "SqlDatabases": "SqlDatabases", "SqlEdge": "SqlEdge", "SuseLinux": "SuseLinux", "VMwareCloudSimple": "VMwareCloudSimple", "VirtualMachineSoftware": "VirtualMachineSoftware", "VirtualMachines": "VirtualMachines"},
        )
        _args_schema.review_date_time = AAZDateTimeArg(
            options=["--review-date-time"],
            arg_group="Properties",
            help="This is the date-time when the Azure Hybrid Benefit needs to be reviewed.",
        )
        _args_schema.term = AAZStrArg(
            options=["--term"],
            arg_group="Properties",
            help="Represent benefit term in ISO 8601 format.",
            enum={"P1Y": "P1Y", "P3Y": "P3Y", "P5Y": "P5Y"},
        )

        applied_scope_prop = cls._args_schema.applied_scope_prop
        applied_scope_prop.display_name = AAZStrArg(
            options=["display-name"],
            help="Display name",
        )
        applied_scope_prop.management_group_id = AAZStrArg(
            options=["management-group-id"],
            help="Fully-qualified identifier of the management group where the benefit must be applied.",
        )
        applied_scope_prop.resource_group_id = AAZStrArg(
            options=["resource-group-id"],
            help="Fully-qualified identifier of the resource group.",
        )
        applied_scope_prop.subscription_id = AAZStrArg(
            options=["subscription-id"],
            help="Fully-qualified identifier of the subscription.",
        )
        applied_scope_prop.tenant_id = AAZStrArg(
            options=["tenant-id"],
            help="Tenant ID where the benefit is applied.",
        )

        # define Arg Group "ReservedResourceProperties"

        _args_schema = cls._args_schema
        _args_schema.instance_flexibility = AAZStrArg(
            options=["--instance-flexibility"],
            arg_group="ReservedResourceProperties",
            help="Turning this on will apply the reservation discount to other VMs in the same VM size group.",
            enum={"Off": "Off", "On": "On"},
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        yield self.ReservationOrderAliasCreate(ctx=self.ctx)()
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

    class ReservationOrderAliasCreate(AAZHttpOperation):
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
                "/providers/Microsoft.BillingBenefits/reservationOrderAliases/{reservationOrderAliasName}",
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
                    "reservationOrderAliasName", self.ctx.args.order_alias_name,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2022-11-01",
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
            _builder.set_prop("location", AAZStrType, ".location")
            _builder.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})
            _builder.set_prop("sku", AAZObjectType, ".", typ_kwargs={"flags": {"required": True}})

            properties = _builder.get(".properties")
            if properties is not None:
                properties.set_prop("appliedScopeProperties", AAZObjectType, ".applied_scope_prop")
                properties.set_prop("appliedScopeType", AAZStrType, ".applied_scope_type")
                properties.set_prop("billingPlan", AAZStrType, ".billing_plan")
                properties.set_prop("billingScopeId", AAZStrType, ".billing_scope_id")
                properties.set_prop("displayName", AAZStrType, ".display_name")
                properties.set_prop("quantity", AAZIntType, ".quantity")
                properties.set_prop("renew", AAZBoolType, ".renew")
                properties.set_prop("reservedResourceProperties", AAZObjectType)
                properties.set_prop("reservedResourceType", AAZStrType, ".reserved_resource_type")
                properties.set_prop("reviewDateTime", AAZStrType, ".review_date_time")
                properties.set_prop("term", AAZStrType, ".term")

            applied_scope_properties = _builder.get(".properties.appliedScopeProperties")
            if applied_scope_properties is not None:
                applied_scope_properties.set_prop("displayName", AAZStrType, ".display_name")
                applied_scope_properties.set_prop("managementGroupId", AAZStrType, ".management_group_id")
                applied_scope_properties.set_prop("resourceGroupId", AAZStrType, ".resource_group_id")
                applied_scope_properties.set_prop("subscriptionId", AAZStrType, ".subscription_id")
                applied_scope_properties.set_prop("tenantId", AAZStrType, ".tenant_id")

            reserved_resource_properties = _builder.get(".properties.reservedResourceProperties")
            if reserved_resource_properties is not None:
                reserved_resource_properties.set_prop("instanceFlexibility", AAZStrType, ".instance_flexibility")

            sku = _builder.get(".sku")
            if sku is not None:
                sku.set_prop("name", AAZStrType, ".sku")

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
            _schema_on_200_201.location = AAZStrType()
            _schema_on_200_201.name = AAZStrType(
                flags={"read_only": True},
            )
            _schema_on_200_201.properties = AAZObjectType(
                flags={"client_flatten": True},
            )
            _schema_on_200_201.sku = AAZObjectType(
                flags={"required": True},
            )
            _schema_on_200_201.system_data = AAZObjectType(
                serialized_name="systemData",
                flags={"read_only": True},
            )
            _schema_on_200_201.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200_201.properties
            properties.applied_scope_properties = AAZObjectType(
                serialized_name="appliedScopeProperties",
            )
            properties.applied_scope_type = AAZStrType(
                serialized_name="appliedScopeType",
            )
            properties.billing_plan = AAZStrType(
                serialized_name="billingPlan",
            )
            properties.billing_scope_id = AAZStrType(
                serialized_name="billingScopeId",
            )
            properties.display_name = AAZStrType(
                serialized_name="displayName",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.quantity = AAZIntType()
            properties.renew = AAZBoolType()
            properties.reservation_order_id = AAZStrType(
                serialized_name="reservationOrderId",
                flags={"read_only": True},
            )
            properties.reserved_resource_properties = AAZObjectType(
                serialized_name="reservedResourceProperties",
            )
            properties.reserved_resource_type = AAZStrType(
                serialized_name="reservedResourceType",
            )
            properties.review_date_time = AAZStrType(
                serialized_name="reviewDateTime",
            )
            properties.term = AAZStrType()

            applied_scope_properties = cls._schema_on_200_201.properties.applied_scope_properties
            applied_scope_properties.display_name = AAZStrType(
                serialized_name="displayName",
            )
            applied_scope_properties.management_group_id = AAZStrType(
                serialized_name="managementGroupId",
            )
            applied_scope_properties.resource_group_id = AAZStrType(
                serialized_name="resourceGroupId",
            )
            applied_scope_properties.subscription_id = AAZStrType(
                serialized_name="subscriptionId",
            )
            applied_scope_properties.tenant_id = AAZStrType(
                serialized_name="tenantId",
            )

            reserved_resource_properties = cls._schema_on_200_201.properties.reserved_resource_properties
            reserved_resource_properties.instance_flexibility = AAZStrType(
                serialized_name="instanceFlexibility",
            )

            sku = cls._schema_on_200_201.sku
            sku.name = AAZStrType()

            system_data = cls._schema_on_200_201.system_data
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

            return cls._schema_on_200_201


class _CreateHelper:
    """Helper class for Create"""


__all__ = ["Create"]
