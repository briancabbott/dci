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
    "dataprotection backup-instance adhoc-backup",
    is_experimental=True,
)
class AdhocBackup(AAZCommand):
    """Trigger adhoc backup.

    :example: Fetch Backup Policy details for rule name and Trigger Adhoc Backup
        az dataprotection backup-policy show -g "000pikumar" --vault-name "PratikPrivatePreviewVault1" -n "backupPolicy"
        az dataprotection backup-instance adhoc-backup --name "testInstance1" --rule-name "BackupWeekly" --retention-tag-override "yearly" --resource-group "000pikumar" --vault-name "PratikPrivatePreviewVault1"
    """

    _aaz_info = {
        "version": "2023-05-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.dataprotection/backupvaults/{}/backupinstances/{}/backup", "2023-05-01"],
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
        _args_schema.backup_instance_name = AAZStrArg(
            options=["-n", "--name", "--backup-instance-name"],
            help="The name of the backup instance.",
            required=True,
            id_part="child_name_1",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        _args_schema.vault_name = AAZStrArg(
            options=["-v", "--vault-name"],
            help="The name of the backup vault.",
            required=True,
            id_part="name",
        )

        # define Arg Group "Backup Rule Options"

        _args_schema = cls._args_schema
        _args_schema.rule_name = AAZStrArg(
            options=["--rule-name"],
            arg_group="Backup Rule Options",
            help="Specify backup policy rule name.",
            required=True,
        )

        # define Arg Group "Backup Rule Options Trigger Option"

        _args_schema = cls._args_schema
        _args_schema.retention_tag_override = AAZStrArg(
            options=["--retention-tag-override"],
            arg_group="Backup Rule Options Trigger Option",
            help="Specify retention override tag.",
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        yield self.BackupInstancesAdhocBackup(ctx=self.ctx)()
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

    class BackupInstancesAdhocBackup(AAZHttpOperation):
        CLIENT_TYPE = "MgmtClient"

        def __call__(self, *args, **kwargs):
            request = self.make_request()
            session = self.client.send_request(request=request, stream=False, **kwargs)
            if session.http_response.status_code in [202]:
                return self.client.build_lro_polling(
                    self.ctx.args.no_wait,
                    session,
                    self.on_200,
                    self.on_error,
                    lro_options={"final-state-via": "location"},
                    path_format_arguments=self.url_parameters,
                )
            if session.http_response.status_code in [200]:
                return self.client.build_lro_polling(
                    self.ctx.args.no_wait,
                    session,
                    self.on_200,
                    self.on_error,
                    lro_options={"final-state-via": "location"},
                    path_format_arguments=self.url_parameters,
                )

            return self.on_error(session.http_response)

        @property
        def url(self):
            return self.client.format_url(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/backup",
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
                    "backupInstanceName", self.ctx.args.backup_instance_name,
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
                **self.serialize_url_param(
                    "vaultName", self.ctx.args.vault_name,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2023-05-01",
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
            _builder.set_prop("backupRuleOptions", AAZObjectType, ".", typ_kwargs={"flags": {"required": True}})

            backup_rule_options = _builder.get(".backupRuleOptions")
            if backup_rule_options is not None:
                backup_rule_options.set_prop("ruleName", AAZStrType, ".rule_name", typ_kwargs={"flags": {"required": True}})
                backup_rule_options.set_prop("triggerOption", AAZObjectType, ".", typ_kwargs={"flags": {"required": True}})

            trigger_option = _builder.get(".backupRuleOptions.triggerOption")
            if trigger_option is not None:
                trigger_option.set_prop("retentionTagOverride", AAZStrType, ".retention_tag_override")

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
            _schema_on_200.job_id = AAZStrType(
                serialized_name="jobId",
            )
            _schema_on_200.object_type = AAZStrType(
                serialized_name="objectType",
                flags={"required": True},
            )

            return cls._schema_on_200


class _AdhocBackupHelper:
    """Helper class for AdhocBackup"""


__all__ = ["AdhocBackup"]
