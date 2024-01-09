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
    "monitor autoscale show",
)
class Show(AAZCommand):
    """Get an autoscale setting

    :example: Show autoscale setting details.
        az monitor autoscale show --name MyAutoscaleSettings --resource-group MyResourceGroup
    """

    _aaz_info = {
        "version": "2022-10-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.insights/autoscalesettings/{}", "2022-10-01"],
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
        _args_schema.autoscale_name = AAZStrArg(
            options=["-n", "--name", "--autoscale-name"],
            help="The autoscale setting name.",
            required=True,
            id_part="name",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.AutoscaleSettingsGet(ctx=self.ctx)()
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

    class AutoscaleSettingsGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Insights/autoscalesettings/{autoscaleSettingName}",
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
                    "autoscaleSettingName", self.ctx.args.autoscale_name,
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
                    "api-version", "2022-10-01",
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
            _schema_on_200.location = AAZStrType(
                flags={"required": True},
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
            _ShowHelper._build_schema_system_data_read(_schema_on_200.system_data)
            _schema_on_200.tags = AAZDictType()
            _schema_on_200.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.properties
            properties.enabled = AAZBoolType()
            properties.name = AAZStrType()
            properties.notifications = AAZListType()
            properties.predictive_autoscale_policy = AAZObjectType(
                serialized_name="predictiveAutoscalePolicy",
            )
            properties.profiles = AAZListType(
                flags={"required": True},
            )
            properties.target_resource_location = AAZStrType(
                serialized_name="targetResourceLocation",
            )
            properties.target_resource_uri = AAZStrType(
                serialized_name="targetResourceUri",
            )

            notifications = cls._schema_on_200.properties.notifications
            notifications.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.notifications.Element
            _element.email = AAZObjectType()
            _element.operation = AAZStrType(
                flags={"required": True},
            )
            _element.webhooks = AAZListType()

            email = cls._schema_on_200.properties.notifications.Element.email
            email.custom_emails = AAZListType(
                serialized_name="customEmails",
            )
            email.send_to_subscription_administrator = AAZBoolType(
                serialized_name="sendToSubscriptionAdministrator",
            )
            email.send_to_subscription_co_administrators = AAZBoolType(
                serialized_name="sendToSubscriptionCoAdministrators",
            )

            custom_emails = cls._schema_on_200.properties.notifications.Element.email.custom_emails
            custom_emails.Element = AAZStrType()

            webhooks = cls._schema_on_200.properties.notifications.Element.webhooks
            webhooks.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.notifications.Element.webhooks.Element
            _element.properties = AAZDictType()
            _element.service_uri = AAZStrType(
                serialized_name="serviceUri",
            )

            properties = cls._schema_on_200.properties.notifications.Element.webhooks.Element.properties
            properties.Element = AAZStrType()

            predictive_autoscale_policy = cls._schema_on_200.properties.predictive_autoscale_policy
            predictive_autoscale_policy.scale_look_ahead_time = AAZStrType(
                serialized_name="scaleLookAheadTime",
            )
            predictive_autoscale_policy.scale_mode = AAZStrType(
                serialized_name="scaleMode",
                flags={"required": True},
            )

            profiles = cls._schema_on_200.properties.profiles
            profiles.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.profiles.Element
            _element.capacity = AAZObjectType(
                flags={"required": True},
            )
            _element.fixed_date = AAZObjectType(
                serialized_name="fixedDate",
            )
            _element.name = AAZStrType(
                flags={"required": True},
            )
            _element.recurrence = AAZObjectType()
            _element.rules = AAZListType(
                flags={"required": True},
            )

            capacity = cls._schema_on_200.properties.profiles.Element.capacity
            capacity.default = AAZStrType(
                flags={"required": True},
            )
            capacity.maximum = AAZStrType(
                flags={"required": True},
            )
            capacity.minimum = AAZStrType(
                flags={"required": True},
            )

            fixed_date = cls._schema_on_200.properties.profiles.Element.fixed_date
            fixed_date.end = AAZStrType(
                flags={"required": True},
            )
            fixed_date.start = AAZStrType(
                flags={"required": True},
            )
            fixed_date.time_zone = AAZStrType(
                serialized_name="timeZone",
            )

            recurrence = cls._schema_on_200.properties.profiles.Element.recurrence
            recurrence.frequency = AAZStrType(
                flags={"required": True},
            )
            recurrence.schedule = AAZObjectType(
                flags={"required": True},
            )

            schedule = cls._schema_on_200.properties.profiles.Element.recurrence.schedule
            schedule.days = AAZListType(
                flags={"required": True},
            )
            schedule.hours = AAZListType(
                flags={"required": True},
            )
            schedule.minutes = AAZListType(
                flags={"required": True},
            )
            schedule.time_zone = AAZStrType(
                serialized_name="timeZone",
                flags={"required": True},
            )

            days = cls._schema_on_200.properties.profiles.Element.recurrence.schedule.days
            days.Element = AAZStrType()

            hours = cls._schema_on_200.properties.profiles.Element.recurrence.schedule.hours
            hours.Element = AAZIntType()

            minutes = cls._schema_on_200.properties.profiles.Element.recurrence.schedule.minutes
            minutes.Element = AAZIntType()

            rules = cls._schema_on_200.properties.profiles.Element.rules
            rules.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.profiles.Element.rules.Element
            _element.metric_trigger = AAZObjectType(
                serialized_name="metricTrigger",
                flags={"required": True},
            )
            _element.scale_action = AAZObjectType(
                serialized_name="scaleAction",
                flags={"required": True},
            )

            metric_trigger = cls._schema_on_200.properties.profiles.Element.rules.Element.metric_trigger
            metric_trigger.dimensions = AAZListType()
            metric_trigger.divide_per_instance = AAZBoolType(
                serialized_name="dividePerInstance",
            )
            metric_trigger.metric_name = AAZStrType(
                serialized_name="metricName",
                flags={"required": True},
            )
            metric_trigger.metric_namespace = AAZStrType(
                serialized_name="metricNamespace",
            )
            metric_trigger.metric_resource_location = AAZStrType(
                serialized_name="metricResourceLocation",
            )
            metric_trigger.metric_resource_uri = AAZStrType(
                serialized_name="metricResourceUri",
                flags={"required": True},
            )
            metric_trigger.operator = AAZStrType(
                flags={"required": True},
            )
            metric_trigger.statistic = AAZStrType(
                flags={"required": True},
            )
            metric_trigger.threshold = AAZFloatType(
                flags={"required": True},
            )
            metric_trigger.time_aggregation = AAZStrType(
                serialized_name="timeAggregation",
                flags={"required": True},
            )
            metric_trigger.time_grain = AAZStrType(
                serialized_name="timeGrain",
                flags={"required": True},
            )
            metric_trigger.time_window = AAZStrType(
                serialized_name="timeWindow",
                flags={"required": True},
            )

            dimensions = cls._schema_on_200.properties.profiles.Element.rules.Element.metric_trigger.dimensions
            dimensions.Element = AAZObjectType()

            _element = cls._schema_on_200.properties.profiles.Element.rules.Element.metric_trigger.dimensions.Element
            _element.dimension_name = AAZStrType(
                serialized_name="DimensionName",
                flags={"required": True},
            )
            _element.operator = AAZStrType(
                serialized_name="Operator",
                flags={"required": True},
            )
            _element.values = AAZListType(
                serialized_name="Values",
                flags={"required": True},
            )

            values = cls._schema_on_200.properties.profiles.Element.rules.Element.metric_trigger.dimensions.Element.values
            values.Element = AAZStrType()

            scale_action = cls._schema_on_200.properties.profiles.Element.rules.Element.scale_action
            scale_action.cooldown = AAZStrType(
                flags={"required": True},
            )
            scale_action.direction = AAZStrType(
                flags={"required": True},
            )
            scale_action.type = AAZStrType(
                flags={"required": True},
            )
            scale_action.value = AAZStrType()

            tags = cls._schema_on_200.tags
            tags.Element = AAZStrType()

            return cls._schema_on_200


class _ShowHelper:
    """Helper class for Show"""

    _schema_system_data_read = None

    @classmethod
    def _build_schema_system_data_read(cls, _schema):
        if cls._schema_system_data_read is not None:
            _schema.created_at = cls._schema_system_data_read.created_at
            _schema.created_by = cls._schema_system_data_read.created_by
            _schema.created_by_type = cls._schema_system_data_read.created_by_type
            _schema.last_modified_at = cls._schema_system_data_read.last_modified_at
            _schema.last_modified_by = cls._schema_system_data_read.last_modified_by
            _schema.last_modified_by_type = cls._schema_system_data_read.last_modified_by_type
            return

        cls._schema_system_data_read = _schema_system_data_read = AAZObjectType(
            flags={"read_only": True}
        )

        system_data_read = _schema_system_data_read
        system_data_read.created_at = AAZStrType(
            serialized_name="createdAt",
        )
        system_data_read.created_by = AAZStrType(
            serialized_name="createdBy",
        )
        system_data_read.created_by_type = AAZStrType(
            serialized_name="createdByType",
        )
        system_data_read.last_modified_at = AAZStrType(
            serialized_name="lastModifiedAt",
        )
        system_data_read.last_modified_by = AAZStrType(
            serialized_name="lastModifiedBy",
        )
        system_data_read.last_modified_by_type = AAZStrType(
            serialized_name="lastModifiedByType",
        )

        _schema.created_at = cls._schema_system_data_read.created_at
        _schema.created_by = cls._schema_system_data_read.created_by
        _schema.created_by_type = cls._schema_system_data_read.created_by_type
        _schema.last_modified_at = cls._schema_system_data_read.last_modified_at
        _schema.last_modified_by = cls._schema_system_data_read.last_modified_by
        _schema.last_modified_by_type = cls._schema_system_data_read.last_modified_by_type


__all__ = ["Show"]
