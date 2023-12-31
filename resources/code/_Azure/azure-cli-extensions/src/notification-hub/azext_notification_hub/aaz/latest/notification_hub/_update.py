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
    "notification-hub update",
    is_experimental=True,
)
class Update(AAZCommand):
    """Update a notification hub in a namespace.

    :example: Update the notification hub
        az notification-hub update --resource-group MyResourceGroup --namespace-name my-namespace --name "sdk-notificationHubs-8708"
    """

    _aaz_info = {
        "version": "2017-04-01",
        "resources": [
            ["mgmt-plane", "/subscriptions/{}/resourcegroups/{}/providers/microsoft.notificationhubs/namespaces/{}/notificationhubs/{}", "2017-04-01"],
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
        _args_schema.namespace_name = AAZStrArg(
            options=["--namespace-name"],
            help="The namespace name.",
            required=True,
            id_part="name",
        )
        _args_schema.notification_hub_name = AAZStrArg(
            options=["-n", "--name", "--notification-hub-name"],
            help="The notification hub name.",
            required=True,
            id_part="child_name_1",
        )
        _args_schema.resource_group = AAZResourceGroupNameArg(
            required=True,
        )

        # define Arg Group "Parameters"

        _args_schema = cls._args_schema
        _args_schema.location = AAZResourceLocationArg(
            arg_group="Parameters",
            help="Resource location",
            nullable=True,
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

        _args_schema = cls._args_schema
        _args_schema.adm_credential = AAZObjectArg(
            options=["--adm-credential"],
            arg_group="Properties",
            help="The AdmCredential of the created NotificationHub",
            nullable=True,
        )
        _args_schema.apns_credential = AAZObjectArg(
            options=["--apns-credential"],
            arg_group="Properties",
            help="The ApnsCredential of the created NotificationHub",
            nullable=True,
        )
        _args_schema.baidu_credential = AAZObjectArg(
            options=["--baidu-credential"],
            arg_group="Properties",
            help="The BaiduCredential of the created NotificationHub",
            nullable=True,
        )
        _args_schema.gcm_credential = AAZObjectArg(
            options=["--gcm-credential"],
            arg_group="Properties",
            help="The GcmCredential of the created NotificationHub",
            nullable=True,
        )
        _args_schema.mpns_credential = AAZObjectArg(
            options=["--mpns-credential"],
            arg_group="Properties",
            help="The MpnsCredential of the created NotificationHub",
            nullable=True,
        )
        _args_schema.wns_credential = AAZObjectArg(
            options=["--wns-credential"],
            arg_group="Properties",
            help="The WnsCredential of the created NotificationHub",
            nullable=True,
        )

        adm_credential = cls._args_schema.adm_credential
        adm_credential.auth_token_url = AAZStrArg(
            options=["auth-token-url"],
            help="The URL of the authorization token.",
            nullable=True,
        )
        adm_credential.client_id = AAZStrArg(
            options=["client-id"],
            help="The client identifier.",
            nullable=True,
        )
        adm_credential.client_secret = AAZStrArg(
            options=["client-secret"],
            help="The credential secret access key.",
            nullable=True,
        )

        apns_credential = cls._args_schema.apns_credential
        apns_credential.apns_certificate = AAZStrArg(
            options=["apns-certificate"],
            help="The APNS certificate. Specify if using Certificate Authentication Mode.",
            nullable=True,
        )
        apns_credential.app_id = AAZStrArg(
            options=["app-id"],
            help="The issuer (iss) registered claim key. The value is a 10-character TeamId, obtained from your developer account. Specify if using Token Authentication Mode.",
            nullable=True,
        )
        apns_credential.app_name = AAZStrArg(
            options=["app-name"],
            help="The name of the application or BundleId. Specify if using Token Authentication Mode.",
            nullable=True,
        )
        apns_credential.certificate_key = AAZStrArg(
            options=["certificate-key"],
            help="The APNS certificate password if it exists.",
            nullable=True,
        )
        apns_credential.endpoint = AAZStrArg(
            options=["endpoint"],
            help="The APNS endpoint of this credential. If using Certificate Authentication Mode and Sandbox specify 'gateway.sandbox.push.apple.com'. If using Certificate Authentication Mode and Production specify 'gateway.push.apple.com'. If using Token Authentication Mode and Sandbox specify 'https://api.development.push.apple.com:443/3/device'. If using Token Authentication Mode and Production specify 'https://api.push.apple.com:443/3/device'.",
            nullable=True,
        )
        apns_credential.key_id = AAZStrArg(
            options=["key-id"],
            help="A 10-character key identifier (kid) key, obtained from your developer account. Specify if using Token Authentication Mode.",
            nullable=True,
        )
        apns_credential.thumbprint = AAZStrArg(
            options=["thumbprint"],
            help="The APNS certificate thumbprint. Specify if using Certificate Authentication Mode.",
            nullable=True,
        )
        apns_credential.token = AAZStrArg(
            options=["token"],
            help="Provider Authentication Token, obtained through your developer account. Specify if using Token Authentication Mode.",
            nullable=True,
        )

        baidu_credential = cls._args_schema.baidu_credential
        baidu_credential.baidu_api_key = AAZStrArg(
            options=["baidu-api-key"],
            help="Baidu Api Key.",
            nullable=True,
        )
        baidu_credential.baidu_end_point = AAZStrArg(
            options=["baidu-end-point"],
            help="Baidu Endpoint.",
            nullable=True,
        )
        baidu_credential.baidu_secret_key = AAZStrArg(
            options=["baidu-secret-key"],
            help="Baidu Secret Key",
            nullable=True,
        )

        gcm_credential = cls._args_schema.gcm_credential
        gcm_credential.gcm_endpoint = AAZStrArg(
            options=["gcm-endpoint"],
            help="The FCM legacy endpoint. Default value is 'https://fcm.googleapis.com/fcm/send'",
            nullable=True,
        )
        gcm_credential.google_api_key = AAZStrArg(
            options=["google-api-key"],
            help="The Google API key.",
            nullable=True,
        )

        mpns_credential = cls._args_schema.mpns_credential
        mpns_credential.certificate_key = AAZStrArg(
            options=["certificate-key"],
            help="The certificate key for this credential.",
            nullable=True,
        )
        mpns_credential.mpns_certificate = AAZStrArg(
            options=["mpns-certificate"],
            help="The MPNS certificate.",
            nullable=True,
        )
        mpns_credential.thumbprint = AAZStrArg(
            options=["thumbprint"],
            help="The MPNS certificate Thumbprint",
            nullable=True,
        )

        wns_credential = cls._args_schema.wns_credential
        wns_credential.package_sid = AAZStrArg(
            options=["package-sid"],
            help="The package ID for this credential.",
            nullable=True,
        )
        wns_credential.secret_key = AAZStrArg(
            options=["secret-key"],
            help="The secret key.",
            nullable=True,
        )
        wns_credential.windows_live_endpoint = AAZStrArg(
            options=["windows-live-endpoint"],
            help="The Windows Live endpoint.",
            nullable=True,
        )
        return cls._args_schema

    def _execute_operations(self):
        self.pre_operations()
        self.NotificationHubsGet(ctx=self.ctx)()
        self.pre_instance_update(self.ctx.vars.instance)
        self.InstanceUpdateByJson(ctx=self.ctx)()
        self.InstanceUpdateByGeneric(ctx=self.ctx)()
        self.post_instance_update(self.ctx.vars.instance)
        self.NotificationHubsCreateOrUpdate(ctx=self.ctx)()
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

    class NotificationHubsGet(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}",
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
                    "namespaceName", self.ctx.args.namespace_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "notificationHubName", self.ctx.args.notification_hub_name,
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
            _UpdateHelper._build_schema_notification_hub_resource_read(cls._schema_on_200)

            return cls._schema_on_200

    class NotificationHubsCreateOrUpdate(AAZHttpOperation):
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
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NotificationHubs/namespaces/{namespaceName}/notificationHubs/{notificationHubName}",
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
                    "namespaceName", self.ctx.args.namespace_name,
                    required=True,
                ),
                **self.serialize_url_param(
                    "notificationHubName", self.ctx.args.notification_hub_name,
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
            _UpdateHelper._build_schema_notification_hub_resource_read(cls._schema_on_200_201)

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
            _builder.set_prop("location", AAZStrType, ".location")
            _builder.set_prop("properties", AAZObjectType, ".", typ_kwargs={"flags": {"required": True, "client_flatten": True}})
            _builder.set_prop("tags", AAZDictType, ".tags")

            properties = _builder.get(".properties")
            if properties is not None:
                properties.set_prop("admCredential", AAZObjectType, ".adm_credential")
                properties.set_prop("apnsCredential", AAZObjectType, ".apns_credential")
                properties.set_prop("baiduCredential", AAZObjectType, ".baidu_credential")
                properties.set_prop("gcmCredential", AAZObjectType, ".gcm_credential")
                properties.set_prop("mpnsCredential", AAZObjectType, ".mpns_credential")
                properties.set_prop("name", AAZStrType, ".notification_hub_name")
                properties.set_prop("wnsCredential", AAZObjectType, ".wns_credential")

            adm_credential = _builder.get(".properties.admCredential")
            if adm_credential is not None:
                adm_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.admCredential.properties")
            if properties is not None:
                properties.set_prop("authTokenUrl", AAZStrType, ".auth_token_url")
                properties.set_prop("clientId", AAZStrType, ".client_id")
                properties.set_prop("clientSecret", AAZStrType, ".client_secret")

            apns_credential = _builder.get(".properties.apnsCredential")
            if apns_credential is not None:
                apns_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.apnsCredential.properties")
            if properties is not None:
                properties.set_prop("apnsCertificate", AAZStrType, ".apns_certificate")
                properties.set_prop("appId", AAZStrType, ".app_id")
                properties.set_prop("appName", AAZStrType, ".app_name")
                properties.set_prop("certificateKey", AAZStrType, ".certificate_key")
                properties.set_prop("endpoint", AAZStrType, ".endpoint")
                properties.set_prop("keyId", AAZStrType, ".key_id")
                properties.set_prop("thumbprint", AAZStrType, ".thumbprint")
                properties.set_prop("token", AAZStrType, ".token")

            baidu_credential = _builder.get(".properties.baiduCredential")
            if baidu_credential is not None:
                baidu_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.baiduCredential.properties")
            if properties is not None:
                properties.set_prop("baiduApiKey", AAZStrType, ".baidu_api_key")
                properties.set_prop("baiduEndPoint", AAZStrType, ".baidu_end_point")
                properties.set_prop("baiduSecretKey", AAZStrType, ".baidu_secret_key")

            gcm_credential = _builder.get(".properties.gcmCredential")
            if gcm_credential is not None:
                gcm_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.gcmCredential.properties")
            if properties is not None:
                properties.set_prop("gcmEndpoint", AAZStrType, ".gcm_endpoint")
                properties.set_prop("googleApiKey", AAZStrType, ".google_api_key")

            mpns_credential = _builder.get(".properties.mpnsCredential")
            if mpns_credential is not None:
                mpns_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.mpnsCredential.properties")
            if properties is not None:
                properties.set_prop("certificateKey", AAZStrType, ".certificate_key")
                properties.set_prop("mpnsCertificate", AAZStrType, ".mpns_certificate")
                properties.set_prop("thumbprint", AAZStrType, ".thumbprint")

            wns_credential = _builder.get(".properties.wnsCredential")
            if wns_credential is not None:
                wns_credential.set_prop("properties", AAZObjectType, typ_kwargs={"flags": {"client_flatten": True}})

            properties = _builder.get(".properties.wnsCredential.properties")
            if properties is not None:
                properties.set_prop("packageSid", AAZStrType, ".package_sid")
                properties.set_prop("secretKey", AAZStrType, ".secret_key")
                properties.set_prop("windowsLiveEndpoint", AAZStrType, ".windows_live_endpoint")

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

    _schema_notification_hub_resource_read = None

    @classmethod
    def _build_schema_notification_hub_resource_read(cls, _schema):
        if cls._schema_notification_hub_resource_read is not None:
            _schema.id = cls._schema_notification_hub_resource_read.id
            _schema.location = cls._schema_notification_hub_resource_read.location
            _schema.name = cls._schema_notification_hub_resource_read.name
            _schema.properties = cls._schema_notification_hub_resource_read.properties
            _schema.sku = cls._schema_notification_hub_resource_read.sku
            _schema.tags = cls._schema_notification_hub_resource_read.tags
            _schema.type = cls._schema_notification_hub_resource_read.type
            return

        cls._schema_notification_hub_resource_read = _schema_notification_hub_resource_read = AAZObjectType()

        notification_hub_resource_read = _schema_notification_hub_resource_read
        notification_hub_resource_read.id = AAZStrType(
            flags={"read_only": True},
        )
        notification_hub_resource_read.location = AAZStrType()
        notification_hub_resource_read.name = AAZStrType(
            flags={"read_only": True},
        )
        notification_hub_resource_read.properties = AAZObjectType(
            flags={"client_flatten": True},
        )
        notification_hub_resource_read.sku = AAZObjectType()
        notification_hub_resource_read.tags = AAZDictType()
        notification_hub_resource_read.type = AAZStrType(
            flags={"read_only": True},
        )

        properties = _schema_notification_hub_resource_read.properties
        properties.adm_credential = AAZObjectType(
            serialized_name="admCredential",
        )
        properties.apns_credential = AAZObjectType(
            serialized_name="apnsCredential",
        )
        properties.authorization_rules = AAZListType(
            serialized_name="authorizationRules",
        )
        properties.baidu_credential = AAZObjectType(
            serialized_name="baiduCredential",
        )
        properties.gcm_credential = AAZObjectType(
            serialized_name="gcmCredential",
        )
        properties.mpns_credential = AAZObjectType(
            serialized_name="mpnsCredential",
        )
        properties.name = AAZStrType()
        properties.registration_ttl = AAZStrType(
            serialized_name="registrationTtl",
        )
        properties.wns_credential = AAZObjectType(
            serialized_name="wnsCredential",
        )

        adm_credential = _schema_notification_hub_resource_read.properties.adm_credential
        adm_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.adm_credential.properties
        properties.auth_token_url = AAZStrType(
            serialized_name="authTokenUrl",
        )
        properties.client_id = AAZStrType(
            serialized_name="clientId",
        )
        properties.client_secret = AAZStrType(
            serialized_name="clientSecret",
        )

        apns_credential = _schema_notification_hub_resource_read.properties.apns_credential
        apns_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.apns_credential.properties
        properties.apns_certificate = AAZStrType(
            serialized_name="apnsCertificate",
        )
        properties.app_id = AAZStrType(
            serialized_name="appId",
        )
        properties.app_name = AAZStrType(
            serialized_name="appName",
        )
        properties.certificate_key = AAZStrType(
            serialized_name="certificateKey",
        )
        properties.endpoint = AAZStrType()
        properties.key_id = AAZStrType(
            serialized_name="keyId",
        )
        properties.thumbprint = AAZStrType()
        properties.token = AAZStrType()

        authorization_rules = _schema_notification_hub_resource_read.properties.authorization_rules
        authorization_rules.Element = AAZObjectType()

        _element = _schema_notification_hub_resource_read.properties.authorization_rules.Element
        _element.claim_type = AAZStrType(
            serialized_name="claimType",
            flags={"read_only": True},
        )
        _element.claim_value = AAZStrType(
            serialized_name="claimValue",
            flags={"read_only": True},
        )
        _element.created_time = AAZStrType(
            serialized_name="createdTime",
            flags={"read_only": True},
        )
        _element.key_name = AAZStrType(
            serialized_name="keyName",
            flags={"read_only": True},
        )
        _element.modified_time = AAZStrType(
            serialized_name="modifiedTime",
            flags={"read_only": True},
        )
        _element.primary_key = AAZStrType(
            serialized_name="primaryKey",
            flags={"read_only": True},
        )
        _element.revision = AAZIntType(
            flags={"read_only": True},
        )
        _element.rights = AAZListType()
        _element.secondary_key = AAZStrType(
            serialized_name="secondaryKey",
            flags={"read_only": True},
        )

        rights = _schema_notification_hub_resource_read.properties.authorization_rules.Element.rights
        rights.Element = AAZStrType()

        baidu_credential = _schema_notification_hub_resource_read.properties.baidu_credential
        baidu_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.baidu_credential.properties
        properties.baidu_api_key = AAZStrType(
            serialized_name="baiduApiKey",
        )
        properties.baidu_end_point = AAZStrType(
            serialized_name="baiduEndPoint",
        )
        properties.baidu_secret_key = AAZStrType(
            serialized_name="baiduSecretKey",
        )

        gcm_credential = _schema_notification_hub_resource_read.properties.gcm_credential
        gcm_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.gcm_credential.properties
        properties.gcm_endpoint = AAZStrType(
            serialized_name="gcmEndpoint",
        )
        properties.google_api_key = AAZStrType(
            serialized_name="googleApiKey",
        )

        mpns_credential = _schema_notification_hub_resource_read.properties.mpns_credential
        mpns_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.mpns_credential.properties
        properties.certificate_key = AAZStrType(
            serialized_name="certificateKey",
        )
        properties.mpns_certificate = AAZStrType(
            serialized_name="mpnsCertificate",
        )
        properties.thumbprint = AAZStrType()

        wns_credential = _schema_notification_hub_resource_read.properties.wns_credential
        wns_credential.properties = AAZObjectType(
            flags={"client_flatten": True},
        )

        properties = _schema_notification_hub_resource_read.properties.wns_credential.properties
        properties.package_sid = AAZStrType(
            serialized_name="packageSid",
        )
        properties.secret_key = AAZStrType(
            serialized_name="secretKey",
        )
        properties.windows_live_endpoint = AAZStrType(
            serialized_name="windowsLiveEndpoint",
        )

        sku = _schema_notification_hub_resource_read.sku
        sku.capacity = AAZIntType()
        sku.family = AAZStrType()
        sku.name = AAZStrType(
            flags={"required": True},
        )
        sku.size = AAZStrType()
        sku.tier = AAZStrType()

        tags = _schema_notification_hub_resource_read.tags
        tags.Element = AAZStrType()

        _schema.id = cls._schema_notification_hub_resource_read.id
        _schema.location = cls._schema_notification_hub_resource_read.location
        _schema.name = cls._schema_notification_hub_resource_read.name
        _schema.properties = cls._schema_notification_hub_resource_read.properties
        _schema.sku = cls._schema_notification_hub_resource_read.sku
        _schema.tags = cls._schema_notification_hub_resource_read.tags
        _schema.type = cls._schema_notification_hub_resource_read.type


__all__ = ["Update"]
