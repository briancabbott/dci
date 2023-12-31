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
    "adp workspace update",
)
class Update(AAZCommand):
    """Update a Workspace

    :example: update workspace to 3 storage account
        az adp workspace update --subscription sample-rg --resource-group sample-rg --workspace-name sample-ws --set properties.StorageAccountCount=3
    """

    _aaz_info = {
        "version": "2022-09-01-preview",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.autonomousdevelopmentplatform/workspaces/{}", "2022-09-01-preview"],
        ]
    }

    AZ_SUPPORT_NO_WAIT = True

    AZ_SUPPORT_GENERIC_UPDATE = True

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
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )
        _args_schema.workspace_name = AAZStrArg(
            options=["-n", "--name", "--workspace-name"],
            help="Workspace Name",
            required=True,
            id_part="name",
            fmt=AAZStrArgFormat(
                pattern="^[a-z0-9][-a-z0-9]{0,45}$",
                max_length=46,
            ),
        )

        # define Arg Group "Resource"

        _args_schema = cls._args_schema
        _args_schema.identity = AAZObjectArg(
            options=["--identity"],
            arg_group="Resource",
            help="The managed service identities assigned to this resource.",
            nullable=True,
        )
        _args_schema.tags = AAZDictArg(
            options=["--tags"],
            arg_group="Resource",
            help="Resource tags.",
            nullable=True,
        )

        identity = cls._args_schema.identity
        identity.identity = AAZObjectArg(
            options=["identity"],
            help="The managed service identities assigned to this resource.",
            nullable=True,
        )

        identity = cls._args_schema.identity.identity
        identity.type = AAZStrArg(
            options=["type"],
            help="The type of managed identity assigned to this resource.",
            enum={"None": "None", "SystemAssigned": "SystemAssigned", "SystemAssigned, UserAssigned": "SystemAssigned, UserAssigned", "UserAssigned": "UserAssigned"},
        )
        identity.user_assigned_identities = AAZObjectArg(
            options=["user-assigned-identities"],
            help="The identities assigned to this resource by the user.",
            nullable=True,
        )

        tags = cls._args_schema.tags
        tags.Element = AAZStrArg(
            nullable=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.WorkspacesGet(ctx=self.ctx)()
        self.InstanceUpdateByJson(ctx=self.ctx)()
        self.InstanceUpdateByGeneric(ctx=self.ctx)()
        yield self.WorkspacesCreateOrUpdate(ctx=self.ctx)()

    def _output(self, *args, **kwargs):
        result = self.deserialize_output(self.ctx.vars.instance, client_flatten=True)
        return result

    class WorkspacesGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/workspaces/{workspaceName}",
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
                    "workspaceName", self.ctx.args.workspace_name,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2022-09-01-preview",
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
            _build_schema_workspace_read(cls._schema_on_200)

            return cls._schema_on_200

    class WorkspacesCreateOrUpdate(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AutonomousDevelopmentPlatform/workspaces/{workspaceName}",
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
                    "resourceGroupName", self.ctx.args.resource_group,
                    required=True,
                ),
                **self.serialize_url_param(
                    "subscriptionId", self.ctx.subscription_id,
                    required=True,
                ),
                **self.serialize_url_param(
                    "workspaceName", self.ctx.args.workspace_name,
                    required=True,
                ),
            }
            return parameters

        @property
        def query_parameters(self):
            parameters = {
                **self.serialize_query_param(
                    "api-version", "2022-09-01-preview",
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
            _build_schema_workspace_read(cls._schema_on_200_201)

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
            _builder.set_prop("identity", AAZObjectType, ".identity")
            _builder.set_prop("tags", AAZDictType, ".tags")

            identity = _builder.get(".identity")
            if identity is not None:
                identity.set_prop("identity", AAZObjectType, ".identity")

            identity = _builder.get(".identity.identity")
            if identity is not None:
                identity.set_prop("type", AAZStrType, ".type", typ_kwargs={"flags": {"required": True}})
                identity.set_prop("userAssignedIdentities", AAZObjectType, ".user_assigned_identities")

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


_schema_workspace_read = None


def _build_schema_workspace_read(_schema):
    global _schema_workspace_read
    if _schema_workspace_read is not None:
        _schema.id = _schema_workspace_read.id
        _schema.identity = _schema_workspace_read.identity
        _schema.location = _schema_workspace_read.location
        _schema.name = _schema_workspace_read.name
        _schema.properties = _schema_workspace_read.properties
        _schema.system_data = _schema_workspace_read.system_data
        _schema.tags = _schema_workspace_read.tags
        _schema.type = _schema_workspace_read.type
        return

    _schema_workspace_read = AAZObjectType()

    workspace_read = _schema_workspace_read
    workspace_read.id = AAZStrType(
        flags={"read_only": True},
    )
    workspace_read.identity = AAZObjectType()
    workspace_read.location = AAZStrType(
        flags={"required": True},
    )
    workspace_read.name = AAZStrType(
        flags={"read_only": True},
    )
    workspace_read.properties = AAZObjectType(
        flags={"client_flatten": True},
    )
    workspace_read.system_data = AAZObjectType(
        serialized_name="systemData",
        flags={"read_only": True},
    )
    workspace_read.tags = AAZDictType()
    workspace_read.type = AAZStrType(
        flags={"read_only": True},
    )

    identity = _schema_workspace_read.identity
    identity.identity = AAZObjectType()

    identity = _schema_workspace_read.identity.identity
    identity.principal_id = AAZStrType(
        serialized_name="principalId",
        flags={"read_only": True},
    )
    identity.tenant_id = AAZStrType(
        serialized_name="tenantId",
        flags={"read_only": True},
    )
    identity.type = AAZStrType(
        flags={"required": True},
    )
    identity.user_assigned_identities = AAZDictType(
        serialized_name="userAssignedIdentities",
    )

    user_assigned_identities = _schema_workspace_read.identity.identity.user_assigned_identities
    user_assigned_identities.Element = AAZObjectType()

    _element = _schema_workspace_read.identity.identity.user_assigned_identities.Element
    _element.client_id = AAZStrType(
        serialized_name="clientId",
        flags={"read_only": True},
    )
    _element.principal_id = AAZStrType(
        serialized_name="principalId",
        flags={"read_only": True},
    )

    properties = _schema_workspace_read.properties
    properties.auto_generated_domain_name_label_scope = AAZStrType(
        serialized_name="autoGeneratedDomainNameLabelScope",
    )
    properties.batch_accounts = AAZListType(
        serialized_name="batchAccounts",
    )
    properties.data_catalog = AAZObjectType(
        serialized_name="dataCatalog",
    )
    properties.data_location = AAZStrType(
        serialized_name="dataLocation",
    )
    properties.direct_read_access = AAZStrType(
        serialized_name="directReadAccess",
    )
    properties.encryption = AAZObjectType()
    properties.endpoint = AAZStrType(
        flags={"read_only": True},
    )
    properties.provisioning_state = AAZStrType(
        serialized_name="provisioningState",
        flags={"read_only": True},
    )
    properties.resim = AAZObjectType()
    properties.source_resource_id = AAZStrType(
        serialized_name="sourceResourceId",
    )
    properties.storage_account_count = AAZIntType(
        serialized_name="storageAccountCount",
    )
    properties.storage_sku = AAZObjectType(
        serialized_name="storageSku",
    )

    batch_accounts = _schema_workspace_read.properties.batch_accounts
    batch_accounts.Element = AAZObjectType()

    _element = _schema_workspace_read.properties.batch_accounts.Element
    _element.batch_account_resource_id = AAZStrType(
        serialized_name="batchAccountResourceId",
        flags={"required": True},
    )
    _element.user_assigned_identity_resource_id = AAZStrType(
        serialized_name="userAssignedIdentityResourceId",
    )

    data_catalog = _schema_workspace_read.properties.data_catalog
    data_catalog.data_explorer = AAZObjectType(
        serialized_name="dataExplorer",
    )
    data_catalog.external_workspace_ids = AAZListType(
        serialized_name="externalWorkspaceIds",
    )
    data_catalog.state = AAZStrType(
        flags={"required": True},
    )

    data_explorer = _schema_workspace_read.properties.data_catalog.data_explorer
    data_explorer.azure_sku = AAZObjectType(
        serialized_name="azureSku",
        flags={"required": True, "client_flatten": True},
    )

    azure_sku = _schema_workspace_read.properties.data_catalog.data_explorer.azure_sku
    azure_sku.capacity = AAZIntType()
    azure_sku.name = AAZStrType()
    azure_sku.tier = AAZStrType()

    external_workspace_ids = _schema_workspace_read.properties.data_catalog.external_workspace_ids
    external_workspace_ids.Element = AAZStrType()

    encryption = _schema_workspace_read.properties.encryption
    encryption.customer_managed_key_encryption = AAZObjectType(
        serialized_name="customerManagedKeyEncryption",
        flags={"client_flatten": True},
    )

    customer_managed_key_encryption = _schema_workspace_read.properties.encryption.customer_managed_key_encryption
    customer_managed_key_encryption.key_encryption_key_identity = AAZObjectType(
        serialized_name="keyEncryptionKeyIdentity",
        flags={"required": True, "client_flatten": True},
    )
    customer_managed_key_encryption.key_encryption_key_url = AAZStrType(
        serialized_name="keyEncryptionKeyUrl",
        flags={"required": True},
    )

    key_encryption_key_identity = _schema_workspace_read.properties.encryption.customer_managed_key_encryption.key_encryption_key_identity
    key_encryption_key_identity.user_assigned_identity_resource_id = AAZStrType(
        serialized_name="userAssignedIdentityResourceId",
        flags={"required": True},
    )

    resim = _schema_workspace_read.properties.resim
    resim.state = AAZStrType(
        flags={"required": True},
    )

    storage_sku = _schema_workspace_read.properties.storage_sku
    storage_sku.name = AAZStrType(
        flags={"required": True},
    )

    system_data = _schema_workspace_read.system_data
    system_data.created_at = AAZStrType(
        serialized_name="createdAt",
        flags={"read_only": True},
    )
    system_data.created_by = AAZStrType(
        serialized_name="createdBy",
        flags={"read_only": True},
    )
    system_data.created_by_type = AAZStrType(
        serialized_name="createdByType",
        flags={"read_only": True},
    )
    system_data.last_modified_at = AAZStrType(
        serialized_name="lastModifiedAt",
        flags={"read_only": True},
    )
    system_data.last_modified_by = AAZStrType(
        serialized_name="lastModifiedBy",
        flags={"read_only": True},
    )
    system_data.last_modified_by_type = AAZStrType(
        serialized_name="lastModifiedByType",
        flags={"read_only": True},
    )

    tags = _schema_workspace_read.tags
    tags.Element = AAZStrType()

    _schema.id = _schema_workspace_read.id
    _schema.identity = _schema_workspace_read.identity
    _schema.location = _schema_workspace_read.location
    _schema.name = _schema_workspace_read.name
    _schema.properties = _schema_workspace_read.properties
    _schema.system_data = _schema_workspace_read.system_data
    _schema.tags = _schema_workspace_read.tags
    _schema.type = _schema_workspace_read.type


__all__ = ["Update"]
