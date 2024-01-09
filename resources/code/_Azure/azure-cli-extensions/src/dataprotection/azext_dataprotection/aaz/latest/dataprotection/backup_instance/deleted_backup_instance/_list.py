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
    "dataprotection backup-instance deleted-backup-instance list",
    is_experimental=True,
)
class List(AAZCommand):
    """List deleted backup instances belonging to a backup vault
    """

    _aaz_info = {
        "version": "2023-05-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.dataprotection/backupvaults/{}/deletedbackupinstances", "2023-05-01"],
        ]
    }

    AZ_SUPPORT_PAGINATION = True

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
        _args_schema.vault_name = AAZStrArg(
            options=["-v", "--vault-name"],
            help="The name of the backup vault.",
            required=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.DeletedBackupInstancesList(ctx=self.ctx)()
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

    class DeletedBackupInstancesList(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/deletedBackupInstances",
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
            )
            _schema_on_200.value = AAZListType()

            value = cls._schema_on_200.value
            value.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element
            _element.id = AAZStrType(
                flags={"read_only": True},
            )
            _element.name = AAZStrType(
                flags={"read_only": True},
            )
            _element.properties = AAZObjectType()
            _element.system_data = AAZObjectType(
                serialized_name="systemData",
                flags={"read_only": True},
            )
            _element.type = AAZStrType(
                flags={"read_only": True},
            )

            properties = cls._schema_on_200.value.Element.properties
            properties.current_protection_state = AAZStrType(
                serialized_name="currentProtectionState",
                flags={"read_only": True},
            )
            properties.data_source_info = AAZObjectType(
                serialized_name="dataSourceInfo",
                flags={"required": True},
            )
            properties.data_source_set_info = AAZObjectType(
                serialized_name="dataSourceSetInfo",
            )
            properties.datasource_auth_credentials = AAZObjectType(
                serialized_name="datasourceAuthCredentials",
            )
            properties.deletion_info = AAZObjectType(
                serialized_name="deletionInfo",
            )
            properties.friendly_name = AAZStrType(
                serialized_name="friendlyName",
            )
            properties.identity_details = AAZObjectType(
                serialized_name="identityDetails",
            )
            properties.object_type = AAZStrType(
                serialized_name="objectType",
                flags={"required": True},
            )
            properties.policy_info = AAZObjectType(
                serialized_name="policyInfo",
                flags={"required": True},
            )
            properties.protection_error_details = AAZObjectType(
                serialized_name="protectionErrorDetails",
            )
            _ListHelper._build_schema_user_facing_error_read(properties.protection_error_details)
            properties.protection_status = AAZObjectType(
                serialized_name="protectionStatus",
            )
            properties.provisioning_state = AAZStrType(
                serialized_name="provisioningState",
                flags={"read_only": True},
            )
            properties.validation_type = AAZStrType(
                serialized_name="validationType",
            )

            data_source_info = cls._schema_on_200.value.Element.properties.data_source_info
            data_source_info.datasource_type = AAZStrType(
                serialized_name="datasourceType",
            )
            data_source_info.object_type = AAZStrType(
                serialized_name="objectType",
            )
            data_source_info.resource_id = AAZStrType(
                serialized_name="resourceID",
                flags={"required": True},
            )
            data_source_info.resource_location = AAZStrType(
                serialized_name="resourceLocation",
            )
            data_source_info.resource_name = AAZStrType(
                serialized_name="resourceName",
            )
            data_source_info.resource_properties = AAZObjectType(
                serialized_name="resourceProperties",
            )
            _ListHelper._build_schema_base_resource_properties_read(data_source_info.resource_properties)
            data_source_info.resource_type = AAZStrType(
                serialized_name="resourceType",
            )
            data_source_info.resource_uri = AAZStrType(
                serialized_name="resourceUri",
            )

            data_source_set_info = cls._schema_on_200.value.Element.properties.data_source_set_info
            data_source_set_info.datasource_type = AAZStrType(
                serialized_name="datasourceType",
            )
            data_source_set_info.object_type = AAZStrType(
                serialized_name="objectType",
            )
            data_source_set_info.resource_id = AAZStrType(
                serialized_name="resourceID",
                flags={"required": True},
            )
            data_source_set_info.resource_location = AAZStrType(
                serialized_name="resourceLocation",
            )
            data_source_set_info.resource_name = AAZStrType(
                serialized_name="resourceName",
            )
            data_source_set_info.resource_properties = AAZObjectType(
                serialized_name="resourceProperties",
            )
            _ListHelper._build_schema_base_resource_properties_read(data_source_set_info.resource_properties)
            data_source_set_info.resource_type = AAZStrType(
                serialized_name="resourceType",
            )
            data_source_set_info.resource_uri = AAZStrType(
                serialized_name="resourceUri",
            )

            datasource_auth_credentials = cls._schema_on_200.value.Element.properties.datasource_auth_credentials
            datasource_auth_credentials.object_type = AAZStrType(
                serialized_name="objectType",
                flags={"required": True},
            )

            disc_secret_store_based_auth_credentials = cls._schema_on_200.value.Element.properties.datasource_auth_credentials.discriminate_by("object_type", "SecretStoreBasedAuthCredentials")
            disc_secret_store_based_auth_credentials.secret_store_resource = AAZObjectType(
                serialized_name="secretStoreResource",
            )

            secret_store_resource = cls._schema_on_200.value.Element.properties.datasource_auth_credentials.discriminate_by("object_type", "SecretStoreBasedAuthCredentials").secret_store_resource
            secret_store_resource.secret_store_type = AAZStrType(
                serialized_name="secretStoreType",
                flags={"required": True},
            )
            secret_store_resource.uri = AAZStrType()
            secret_store_resource.value = AAZStrType()

            deletion_info = cls._schema_on_200.value.Element.properties.deletion_info
            deletion_info.billing_end_date = AAZStrType(
                serialized_name="billingEndDate",
                flags={"read_only": True},
            )
            deletion_info.delete_activity_id = AAZStrType(
                serialized_name="deleteActivityID",
                flags={"read_only": True},
            )
            deletion_info.deletion_time = AAZStrType(
                serialized_name="deletionTime",
                flags={"read_only": True},
            )
            deletion_info.scheduled_purge_time = AAZStrType(
                serialized_name="scheduledPurgeTime",
                flags={"read_only": True},
            )

            identity_details = cls._schema_on_200.value.Element.properties.identity_details
            identity_details.use_system_assigned_identity = AAZBoolType(
                serialized_name="useSystemAssignedIdentity",
            )
            identity_details.user_assigned_identity_arm_url = AAZStrType(
                serialized_name="userAssignedIdentityArmUrl",
            )

            policy_info = cls._schema_on_200.value.Element.properties.policy_info
            policy_info.policy_id = AAZStrType(
                serialized_name="policyId",
                flags={"required": True},
            )
            policy_info.policy_parameters = AAZObjectType(
                serialized_name="policyParameters",
            )
            policy_info.policy_version = AAZStrType(
                serialized_name="policyVersion",
                flags={"read_only": True},
            )

            policy_parameters = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters
            policy_parameters.backup_datasource_parameters_list = AAZListType(
                serialized_name="backupDatasourceParametersList",
            )
            policy_parameters.data_store_parameters_list = AAZListType(
                serialized_name="dataStoreParametersList",
            )

            backup_datasource_parameters_list = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list
            backup_datasource_parameters_list.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element
            _element.object_type = AAZStrType(
                serialized_name="objectType",
                flags={"required": True},
            )

            disc_blob_backup_datasource_parameters = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "BlobBackupDatasourceParameters")
            disc_blob_backup_datasource_parameters.containers_list = AAZListType(
                serialized_name="containersList",
                flags={"required": True},
            )

            containers_list = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "BlobBackupDatasourceParameters").containers_list
            containers_list.Element = AAZStrType()

            disc_kubernetes_cluster_backup_datasource_parameters = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters")
            disc_kubernetes_cluster_backup_datasource_parameters.backup_hook_references = AAZListType(
                serialized_name="backupHookReferences",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.excluded_namespaces = AAZListType(
                serialized_name="excludedNamespaces",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.excluded_resource_types = AAZListType(
                serialized_name="excludedResourceTypes",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.include_cluster_scope_resources = AAZBoolType(
                serialized_name="includeClusterScopeResources",
                flags={"required": True},
            )
            disc_kubernetes_cluster_backup_datasource_parameters.included_namespaces = AAZListType(
                serialized_name="includedNamespaces",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.included_resource_types = AAZListType(
                serialized_name="includedResourceTypes",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.label_selectors = AAZListType(
                serialized_name="labelSelectors",
            )
            disc_kubernetes_cluster_backup_datasource_parameters.snapshot_volumes = AAZBoolType(
                serialized_name="snapshotVolumes",
                flags={"required": True},
            )

            backup_hook_references = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").backup_hook_references
            backup_hook_references.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").backup_hook_references.Element
            _element.name = AAZStrType()
            _element.namespace = AAZStrType()

            excluded_namespaces = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").excluded_namespaces
            excluded_namespaces.Element = AAZStrType()

            excluded_resource_types = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").excluded_resource_types
            excluded_resource_types.Element = AAZStrType()

            included_namespaces = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").included_namespaces
            included_namespaces.Element = AAZStrType()

            included_resource_types = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").included_resource_types
            included_resource_types.Element = AAZStrType()

            label_selectors = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.backup_datasource_parameters_list.Element.discriminate_by("object_type", "KubernetesClusterBackupDatasourceParameters").label_selectors
            label_selectors.Element = AAZStrType()

            data_store_parameters_list = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.data_store_parameters_list
            data_store_parameters_list.Element = AAZObjectType()

            _element = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.data_store_parameters_list.Element
            _element.data_store_type = AAZStrType(
                serialized_name="dataStoreType",
                flags={"required": True},
            )
            _element.object_type = AAZStrType(
                serialized_name="objectType",
                flags={"required": True},
            )

            disc_azure_operational_store_parameters = cls._schema_on_200.value.Element.properties.policy_info.policy_parameters.data_store_parameters_list.Element.discriminate_by("object_type", "AzureOperationalStoreParameters")
            disc_azure_operational_store_parameters.resource_group_id = AAZStrType(
                serialized_name="resourceGroupId",
            )

            protection_status = cls._schema_on_200.value.Element.properties.protection_status
            protection_status.error_details = AAZObjectType(
                serialized_name="errorDetails",
            )
            _ListHelper._build_schema_user_facing_error_read(protection_status.error_details)
            protection_status.status = AAZStrType()

            system_data = cls._schema_on_200.value.Element.system_data
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


class _ListHelper:
    """Helper class for List"""

    _schema_base_resource_properties_read = None

    @classmethod
    def _build_schema_base_resource_properties_read(cls, _schema):
        if cls._schema_base_resource_properties_read is not None:
            _schema.object_type = cls._schema_base_resource_properties_read.object_type
            return

        cls._schema_base_resource_properties_read = _schema_base_resource_properties_read = AAZObjectType()

        base_resource_properties_read = _schema_base_resource_properties_read
        base_resource_properties_read.object_type = AAZStrType(
            serialized_name="objectType",
            flags={"required": True},
        )

        _schema.object_type = cls._schema_base_resource_properties_read.object_type

    _schema_inner_error_read = None

    @classmethod
    def _build_schema_inner_error_read(cls, _schema):
        if cls._schema_inner_error_read is not None:
            _schema.additional_info = cls._schema_inner_error_read.additional_info
            _schema.code = cls._schema_inner_error_read.code
            _schema.embedded_inner_error = cls._schema_inner_error_read.embedded_inner_error
            return

        cls._schema_inner_error_read = _schema_inner_error_read = AAZObjectType()

        inner_error_read = _schema_inner_error_read
        inner_error_read.additional_info = AAZDictType(
            serialized_name="additionalInfo",
        )
        inner_error_read.code = AAZStrType()
        inner_error_read.embedded_inner_error = AAZObjectType(
            serialized_name="embeddedInnerError",
        )
        cls._build_schema_inner_error_read(inner_error_read.embedded_inner_error)

        additional_info = _schema_inner_error_read.additional_info
        additional_info.Element = AAZStrType()

        _schema.additional_info = cls._schema_inner_error_read.additional_info
        _schema.code = cls._schema_inner_error_read.code
        _schema.embedded_inner_error = cls._schema_inner_error_read.embedded_inner_error

    _schema_user_facing_error_read = None

    @classmethod
    def _build_schema_user_facing_error_read(cls, _schema):
        if cls._schema_user_facing_error_read is not None:
            _schema.code = cls._schema_user_facing_error_read.code
            _schema.details = cls._schema_user_facing_error_read.details
            _schema.inner_error = cls._schema_user_facing_error_read.inner_error
            _schema.is_retryable = cls._schema_user_facing_error_read.is_retryable
            _schema.is_user_error = cls._schema_user_facing_error_read.is_user_error
            _schema.message = cls._schema_user_facing_error_read.message
            _schema.properties = cls._schema_user_facing_error_read.properties
            _schema.recommended_action = cls._schema_user_facing_error_read.recommended_action
            _schema.target = cls._schema_user_facing_error_read.target
            return

        cls._schema_user_facing_error_read = _schema_user_facing_error_read = AAZObjectType()

        user_facing_error_read = _schema_user_facing_error_read
        user_facing_error_read.code = AAZStrType()
        user_facing_error_read.details = AAZListType()
        user_facing_error_read.inner_error = AAZObjectType(
            serialized_name="innerError",
        )
        cls._build_schema_inner_error_read(user_facing_error_read.inner_error)
        user_facing_error_read.is_retryable = AAZBoolType(
            serialized_name="isRetryable",
        )
        user_facing_error_read.is_user_error = AAZBoolType(
            serialized_name="isUserError",
        )
        user_facing_error_read.message = AAZStrType()
        user_facing_error_read.properties = AAZDictType()
        user_facing_error_read.recommended_action = AAZListType(
            serialized_name="recommendedAction",
        )
        user_facing_error_read.target = AAZStrType()

        details = _schema_user_facing_error_read.details
        details.Element = AAZObjectType()
        cls._build_schema_user_facing_error_read(details.Element)

        properties = _schema_user_facing_error_read.properties
        properties.Element = AAZStrType()

        recommended_action = _schema_user_facing_error_read.recommended_action
        recommended_action.Element = AAZStrType()

        _schema.code = cls._schema_user_facing_error_read.code
        _schema.details = cls._schema_user_facing_error_read.details
        _schema.inner_error = cls._schema_user_facing_error_read.inner_error
        _schema.is_retryable = cls._schema_user_facing_error_read.is_retryable
        _schema.is_user_error = cls._schema_user_facing_error_read.is_user_error
        _schema.message = cls._schema_user_facing_error_read.message
        _schema.properties = cls._schema_user_facing_error_read.properties
        _schema.recommended_action = cls._schema_user_facing_error_read.recommended_action
        _schema.target = cls._schema_user_facing_error_read.target


__all__ = ["List"]
