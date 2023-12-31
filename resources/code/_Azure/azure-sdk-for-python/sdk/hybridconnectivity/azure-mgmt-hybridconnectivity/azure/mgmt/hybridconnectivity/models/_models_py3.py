# coding=utf-8
# pylint: disable=too-many-lines
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import datetime
from typing import Any, List, Optional, TYPE_CHECKING, Union

from .. import _serialization

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models


class EndpointAccessResource(_serialization.Model):
    """The endpoint access for the target resource.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar namespace_name: The namespace name.
    :vartype namespace_name: str
    :ivar namespace_name_suffix: The suffix domain name of relay namespace.
    :vartype namespace_name_suffix: str
    :ivar hybrid_connection_name: Azure Relay hybrid connection name for the resource.
    :vartype hybrid_connection_name: str
    :ivar access_key: Access key for hybrid connection.
    :vartype access_key: str
    :ivar expires_on: The expiration of access key in unix time.
    :vartype expires_on: int
    :ivar service_configuration_token: The token to access the enabled service.
    :vartype service_configuration_token: str
    """

    _validation = {
        "namespace_name": {"max_length": 200, "min_length": 1},
        "namespace_name_suffix": {"max_length": 100, "min_length": 1},
        "access_key": {"readonly": True},
    }

    _attribute_map = {
        "namespace_name": {"key": "relay.namespaceName", "type": "str"},
        "namespace_name_suffix": {"key": "relay.namespaceNameSuffix", "type": "str"},
        "hybrid_connection_name": {"key": "relay.hybridConnectionName", "type": "str"},
        "access_key": {"key": "relay.accessKey", "type": "str"},
        "expires_on": {"key": "relay.expiresOn", "type": "int"},
        "service_configuration_token": {"key": "relay.serviceConfigurationToken", "type": "str"},
    }

    def __init__(
        self,
        *,
        namespace_name: Optional[str] = None,
        namespace_name_suffix: Optional[str] = None,
        hybrid_connection_name: Optional[str] = None,
        expires_on: Optional[int] = None,
        service_configuration_token: Optional[str] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword namespace_name: The namespace name.
        :paramtype namespace_name: str
        :keyword namespace_name_suffix: The suffix domain name of relay namespace.
        :paramtype namespace_name_suffix: str
        :keyword hybrid_connection_name: Azure Relay hybrid connection name for the resource.
        :paramtype hybrid_connection_name: str
        :keyword expires_on: The expiration of access key in unix time.
        :paramtype expires_on: int
        :keyword service_configuration_token: The token to access the enabled service.
        :paramtype service_configuration_token: str
        """
        super().__init__(**kwargs)
        self.namespace_name = namespace_name
        self.namespace_name_suffix = namespace_name_suffix
        self.hybrid_connection_name = hybrid_connection_name
        self.access_key = None
        self.expires_on = expires_on
        self.service_configuration_token = service_configuration_token


class EndpointProperties(_serialization.Model):
    """Endpoint details.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar type: The type of endpoint. Required. Known values are: "default" and "custom".
    :vartype type: str or ~azure.mgmt.hybridconnectivity.models.Type
    :ivar resource_id: The resource Id of the connectivity endpoint (optional).
    :vartype resource_id: str
    :ivar provisioning_state: The resource provisioning state.
    :vartype provisioning_state: str
    """

    _validation = {
        "type": {"required": True},
        "provisioning_state": {"readonly": True},
    }

    _attribute_map = {
        "type": {"key": "type", "type": "str"},
        "resource_id": {"key": "resourceId", "type": "str"},
        "provisioning_state": {"key": "provisioningState", "type": "str"},
    }

    def __init__(self, *, type: Union[str, "_models.Type"], resource_id: Optional[str] = None, **kwargs: Any) -> None:
        """
        :keyword type: The type of endpoint. Required. Known values are: "default" and "custom".
        :paramtype type: str or ~azure.mgmt.hybridconnectivity.models.Type
        :keyword resource_id: The resource Id of the connectivity endpoint (optional).
        :paramtype resource_id: str
        """
        super().__init__(**kwargs)
        self.type = type
        self.resource_id = resource_id
        self.provisioning_state = None


class Resource(_serialization.Model):
    """Common fields that are returned in the response for all Azure Resource Manager resources.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. E.g.
     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}".
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar system_data: Azure Resource Manager metadata containing createdBy and modifiedBy
     information.
    :vartype system_data: ~azure.mgmt.hybridconnectivity.models.SystemData
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
        "system_data": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "system_data": {"key": "systemData", "type": "SystemData"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None
        self.system_data = None


class ProxyResource(Resource):
    """The resource model definition for a Azure Resource Manager proxy resource. It will not have
    tags and a location.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. E.g.
     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}".
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar system_data: Azure Resource Manager metadata containing createdBy and modifiedBy
     information.
    :vartype system_data: ~azure.mgmt.hybridconnectivity.models.SystemData
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
        "system_data": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "system_data": {"key": "systemData", "type": "SystemData"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)


class EndpointResource(ProxyResource):
    """The endpoint for the target resource.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. E.g.
     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}".
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar system_data: Azure Resource Manager metadata containing createdBy and modifiedBy
     information.
    :vartype system_data: ~azure.mgmt.hybridconnectivity.models.SystemData
    :ivar properties: The endpoint properties.
    :vartype properties: ~azure.mgmt.hybridconnectivity.models.EndpointProperties
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
        "system_data": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "system_data": {"key": "systemData", "type": "SystemData"},
        "properties": {"key": "properties", "type": "EndpointProperties"},
    }

    def __init__(self, *, properties: Optional["_models.EndpointProperties"] = None, **kwargs: Any) -> None:
        """
        :keyword properties: The endpoint properties.
        :paramtype properties: ~azure.mgmt.hybridconnectivity.models.EndpointProperties
        """
        super().__init__(**kwargs)
        self.properties = properties


class EndpointsList(_serialization.Model):
    """The list of endpoints.

    :ivar next_link: The link used to get the next page of endpoints list.
    :vartype next_link: str
    :ivar value: The list of endpoint.
    :vartype value: list[~azure.mgmt.hybridconnectivity.models.EndpointResource]
    """

    _attribute_map = {
        "next_link": {"key": "nextLink", "type": "str"},
        "value": {"key": "value", "type": "[EndpointResource]"},
    }

    def __init__(
        self,
        *,
        next_link: Optional[str] = None,
        value: Optional[List["_models.EndpointResource"]] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword next_link: The link used to get the next page of endpoints list.
        :paramtype next_link: str
        :keyword value: The list of endpoint.
        :paramtype value: list[~azure.mgmt.hybridconnectivity.models.EndpointResource]
        """
        super().__init__(**kwargs)
        self.next_link = next_link
        self.value = value


class ErrorAdditionalInfo(_serialization.Model):
    """The resource management error additional info.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar type: The additional info type.
    :vartype type: str
    :ivar info: The additional info.
    :vartype info: JSON
    """

    _validation = {
        "type": {"readonly": True},
        "info": {"readonly": True},
    }

    _attribute_map = {
        "type": {"key": "type", "type": "str"},
        "info": {"key": "info", "type": "object"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)
        self.type = None
        self.info = None


class ErrorDetail(_serialization.Model):
    """The error detail.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar code: The error code.
    :vartype code: str
    :ivar message: The error message.
    :vartype message: str
    :ivar target: The error target.
    :vartype target: str
    :ivar details: The error details.
    :vartype details: list[~azure.mgmt.hybridconnectivity.models.ErrorDetail]
    :ivar additional_info: The error additional info.
    :vartype additional_info: list[~azure.mgmt.hybridconnectivity.models.ErrorAdditionalInfo]
    """

    _validation = {
        "code": {"readonly": True},
        "message": {"readonly": True},
        "target": {"readonly": True},
        "details": {"readonly": True},
        "additional_info": {"readonly": True},
    }

    _attribute_map = {
        "code": {"key": "code", "type": "str"},
        "message": {"key": "message", "type": "str"},
        "target": {"key": "target", "type": "str"},
        "details": {"key": "details", "type": "[ErrorDetail]"},
        "additional_info": {"key": "additionalInfo", "type": "[ErrorAdditionalInfo]"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)
        self.code = None
        self.message = None
        self.target = None
        self.details = None
        self.additional_info = None


class ErrorResponse(_serialization.Model):
    """Common error response for all Azure Resource Manager APIs to return error details for failed
    operations. (This also follows the OData error response format.).

    :ivar error: The error object.
    :vartype error: ~azure.mgmt.hybridconnectivity.models.ErrorDetail
    """

    _attribute_map = {
        "error": {"key": "error", "type": "ErrorDetail"},
    }

    def __init__(self, *, error: Optional["_models.ErrorDetail"] = None, **kwargs: Any) -> None:
        """
        :keyword error: The error object.
        :paramtype error: ~azure.mgmt.hybridconnectivity.models.ErrorDetail
        """
        super().__init__(**kwargs)
        self.error = error


class IngressGatewayResource(_serialization.Model):
    """The ingress gateway access credentials.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar hostname: The ingress hostname.
    :vartype hostname: str
    :ivar server_id: The arc ingress gateway server app id.
    :vartype server_id: str
    :ivar tenant_id: The target resource home tenant id.
    :vartype tenant_id: str
    :ivar namespace_name: The namespace name.
    :vartype namespace_name: str
    :ivar namespace_name_suffix: The suffix domain name of relay namespace.
    :vartype namespace_name_suffix: str
    :ivar hybrid_connection_name: Azure Relay hybrid connection name for the resource.
    :vartype hybrid_connection_name: str
    :ivar access_key: Access key for hybrid connection.
    :vartype access_key: str
    :ivar expires_on: The expiration of access key in unix time.
    :vartype expires_on: int
    :ivar service_configuration_token: The token to access the enabled service.
    :vartype service_configuration_token: str
    """

    _validation = {
        "namespace_name": {"max_length": 200, "min_length": 1},
        "namespace_name_suffix": {"max_length": 100, "min_length": 1},
        "access_key": {"readonly": True},
    }

    _attribute_map = {
        "hostname": {"key": "ingress.hostname", "type": "str"},
        "server_id": {"key": "ingress.aadProfile.serverId", "type": "str"},
        "tenant_id": {"key": "ingress.aadProfile.tenantId", "type": "str"},
        "namespace_name": {"key": "relay.namespaceName", "type": "str"},
        "namespace_name_suffix": {"key": "relay.namespaceNameSuffix", "type": "str"},
        "hybrid_connection_name": {"key": "relay.hybridConnectionName", "type": "str"},
        "access_key": {"key": "relay.accessKey", "type": "str"},
        "expires_on": {"key": "relay.expiresOn", "type": "int"},
        "service_configuration_token": {"key": "relay.serviceConfigurationToken", "type": "str"},
    }

    def __init__(
        self,
        *,
        hostname: Optional[str] = None,
        server_id: Optional[str] = None,
        tenant_id: Optional[str] = None,
        namespace_name: Optional[str] = None,
        namespace_name_suffix: Optional[str] = None,
        hybrid_connection_name: Optional[str] = None,
        expires_on: Optional[int] = None,
        service_configuration_token: Optional[str] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword hostname: The ingress hostname.
        :paramtype hostname: str
        :keyword server_id: The arc ingress gateway server app id.
        :paramtype server_id: str
        :keyword tenant_id: The target resource home tenant id.
        :paramtype tenant_id: str
        :keyword namespace_name: The namespace name.
        :paramtype namespace_name: str
        :keyword namespace_name_suffix: The suffix domain name of relay namespace.
        :paramtype namespace_name_suffix: str
        :keyword hybrid_connection_name: Azure Relay hybrid connection name for the resource.
        :paramtype hybrid_connection_name: str
        :keyword expires_on: The expiration of access key in unix time.
        :paramtype expires_on: int
        :keyword service_configuration_token: The token to access the enabled service.
        :paramtype service_configuration_token: str
        """
        super().__init__(**kwargs)
        self.hostname = hostname
        self.server_id = server_id
        self.tenant_id = tenant_id
        self.namespace_name = namespace_name
        self.namespace_name_suffix = namespace_name_suffix
        self.hybrid_connection_name = hybrid_connection_name
        self.access_key = None
        self.expires_on = expires_on
        self.service_configuration_token = service_configuration_token


class ListCredentialsRequest(_serialization.Model):
    """The details of the service for which credentials needs to be returned.

    :ivar service_name: The name of the service. If not provided, the request will by pass the
     generation of service configuration token. Known values are: "SSH" and "WAC".
    :vartype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
    """

    _attribute_map = {
        "service_name": {"key": "serviceName", "type": "str"},
    }

    def __init__(self, *, service_name: Optional[Union[str, "_models.ServiceName"]] = None, **kwargs: Any) -> None:
        """
        :keyword service_name: The name of the service. If not provided, the request will by pass the
         generation of service configuration token. Known values are: "SSH" and "WAC".
        :paramtype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
        """
        super().__init__(**kwargs)
        self.service_name = service_name


class ListIngressGatewayCredentialsRequest(_serialization.Model):
    """Represent ListIngressGatewayCredentials Request object.

    :ivar service_name: The name of the service. Known values are: "SSH" and "WAC".
    :vartype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
    """

    _attribute_map = {
        "service_name": {"key": "serviceName", "type": "str"},
    }

    def __init__(self, *, service_name: Optional[Union[str, "_models.ServiceName"]] = None, **kwargs: Any) -> None:
        """
        :keyword service_name: The name of the service. Known values are: "SSH" and "WAC".
        :paramtype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
        """
        super().__init__(**kwargs)
        self.service_name = service_name


class ManagedProxyRequest(_serialization.Model):
    """Represent ManageProxy Request object.

    All required parameters must be populated in order to send to Azure.

    :ivar service: The name of the service. Required.
    :vartype service: str
    :ivar hostname: The target host name.
    :vartype hostname: str
    :ivar service_name: The name of the service. It is an optional property, if not provided,
     service configuration tokens issue code would be by passed. Known values are: "SSH" and "WAC".
    :vartype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
    """

    _validation = {
        "service": {"required": True},
    }

    _attribute_map = {
        "service": {"key": "service", "type": "str"},
        "hostname": {"key": "hostname", "type": "str"},
        "service_name": {"key": "serviceName", "type": "str"},
    }

    def __init__(
        self,
        *,
        service: str,
        hostname: Optional[str] = None,
        service_name: Optional[Union[str, "_models.ServiceName"]] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword service: The name of the service. Required.
        :paramtype service: str
        :keyword hostname: The target host name.
        :paramtype hostname: str
        :keyword service_name: The name of the service. It is an optional property, if not provided,
         service configuration tokens issue code would be by passed. Known values are: "SSH" and "WAC".
        :paramtype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
        """
        super().__init__(**kwargs)
        self.service = service
        self.hostname = hostname
        self.service_name = service_name


class ManagedProxyResource(_serialization.Model):
    """Managed Proxy.

    All required parameters must be populated in order to send to Azure.

    :ivar proxy: The short lived proxy name. Required.
    :vartype proxy: str
    :ivar expires_on: The expiration time of short lived proxy name in unix epoch. Required.
    :vartype expires_on: int
    """

    _validation = {
        "proxy": {"required": True},
        "expires_on": {"required": True},
    }

    _attribute_map = {
        "proxy": {"key": "proxy", "type": "str"},
        "expires_on": {"key": "expiresOn", "type": "int"},
    }

    def __init__(self, *, proxy: str, expires_on: int, **kwargs: Any) -> None:
        """
        :keyword proxy: The short lived proxy name. Required.
        :paramtype proxy: str
        :keyword expires_on: The expiration time of short lived proxy name in unix epoch. Required.
        :paramtype expires_on: int
        """
        super().__init__(**kwargs)
        self.proxy = proxy
        self.expires_on = expires_on


class Operation(_serialization.Model):
    """Details of a REST API operation, returned from the Resource Provider Operations API.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar name: The name of the operation, as per Resource-Based Access Control (RBAC). Examples:
     "Microsoft.Compute/virtualMachines/write", "Microsoft.Compute/virtualMachines/capture/action".
    :vartype name: str
    :ivar is_data_action: Whether the operation applies to data-plane. This is "true" for
     data-plane operations and "false" for ARM/control-plane operations.
    :vartype is_data_action: bool
    :ivar display: Localized display information for this particular operation.
    :vartype display: ~azure.mgmt.hybridconnectivity.models.OperationDisplay
    :ivar origin: The intended executor of the operation; as in Resource Based Access Control
     (RBAC) and audit logs UX. Default value is "user,system". Known values are: "user", "system",
     and "user,system".
    :vartype origin: str or ~azure.mgmt.hybridconnectivity.models.Origin
    :ivar action_type: Enum. Indicates the action type. "Internal" refers to actions that are for
     internal only APIs. "Internal"
    :vartype action_type: str or ~azure.mgmt.hybridconnectivity.models.ActionType
    """

    _validation = {
        "name": {"readonly": True},
        "is_data_action": {"readonly": True},
        "origin": {"readonly": True},
        "action_type": {"readonly": True},
    }

    _attribute_map = {
        "name": {"key": "name", "type": "str"},
        "is_data_action": {"key": "isDataAction", "type": "bool"},
        "display": {"key": "display", "type": "OperationDisplay"},
        "origin": {"key": "origin", "type": "str"},
        "action_type": {"key": "actionType", "type": "str"},
    }

    def __init__(self, *, display: Optional["_models.OperationDisplay"] = None, **kwargs: Any) -> None:
        """
        :keyword display: Localized display information for this particular operation.
        :paramtype display: ~azure.mgmt.hybridconnectivity.models.OperationDisplay
        """
        super().__init__(**kwargs)
        self.name = None
        self.is_data_action = None
        self.display = display
        self.origin = None
        self.action_type = None


class OperationDisplay(_serialization.Model):
    """Localized display information for this particular operation.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar provider: The localized friendly form of the resource provider name, e.g. "Microsoft
     Monitoring Insights" or "Microsoft Compute".
    :vartype provider: str
    :ivar resource: The localized friendly name of the resource type related to this operation.
     E.g. "Virtual Machines" or "Job Schedule Collections".
    :vartype resource: str
    :ivar operation: The concise, localized friendly name for the operation; suitable for
     dropdowns. E.g. "Create or Update Virtual Machine", "Restart Virtual Machine".
    :vartype operation: str
    :ivar description: The short, localized friendly description of the operation; suitable for
     tool tips and detailed views.
    :vartype description: str
    """

    _validation = {
        "provider": {"readonly": True},
        "resource": {"readonly": True},
        "operation": {"readonly": True},
        "description": {"readonly": True},
    }

    _attribute_map = {
        "provider": {"key": "provider", "type": "str"},
        "resource": {"key": "resource", "type": "str"},
        "operation": {"key": "operation", "type": "str"},
        "description": {"key": "description", "type": "str"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)
        self.provider = None
        self.resource = None
        self.operation = None
        self.description = None


class OperationListResult(_serialization.Model):
    """A list of REST API operations supported by an Azure Resource Provider. It contains an URL link
    to get the next set of results.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar value: List of operations supported by the resource provider.
    :vartype value: list[~azure.mgmt.hybridconnectivity.models.Operation]
    :ivar next_link: URL to get the next set of operation list results (if there are any).
    :vartype next_link: str
    """

    _validation = {
        "value": {"readonly": True},
        "next_link": {"readonly": True},
    }

    _attribute_map = {
        "value": {"key": "value", "type": "[Operation]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(self, **kwargs: Any) -> None:
        """ """
        super().__init__(**kwargs)
        self.value = None
        self.next_link = None


class ServiceConfigurationList(_serialization.Model):
    """The paginated list of serviceConfigurations.

    :ivar value: The list of service configuration.
    :vartype value: list[~azure.mgmt.hybridconnectivity.models.ServiceConfigurationResource]
    :ivar next_link: The link to fetch the next page of connected cluster.
    :vartype next_link: str
    """

    _attribute_map = {
        "value": {"key": "value", "type": "[ServiceConfigurationResource]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(
        self,
        *,
        value: Optional[List["_models.ServiceConfigurationResource"]] = None,
        next_link: Optional[str] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword value: The list of service configuration.
        :paramtype value: list[~azure.mgmt.hybridconnectivity.models.ServiceConfigurationResource]
        :keyword next_link: The link to fetch the next page of connected cluster.
        :paramtype next_link: str
        """
        super().__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class ServiceConfigurationResource(ProxyResource):
    """The service configuration details associated with the target resource.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. E.g.
     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}".
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar system_data: Azure Resource Manager metadata containing createdBy and modifiedBy
     information.
    :vartype system_data: ~azure.mgmt.hybridconnectivity.models.SystemData
    :ivar service_name: Name of the service. Known values are: "SSH" and "WAC".
    :vartype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
    :ivar resource_id: The resource Id of the connectivity endpoint (optional).
    :vartype resource_id: str
    :ivar port: The port on which service is enabled.
    :vartype port: int
    :ivar provisioning_state: The resource provisioning state. Known values are: "Succeeded",
     "Creating", "Updating", "Failed", and "Canceled".
    :vartype provisioning_state: str or ~azure.mgmt.hybridconnectivity.models.ProvisioningState
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
        "system_data": {"readonly": True},
        "provisioning_state": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "system_data": {"key": "systemData", "type": "SystemData"},
        "service_name": {"key": "properties.serviceName", "type": "str"},
        "resource_id": {"key": "properties.resourceId", "type": "str"},
        "port": {"key": "properties.port", "type": "int"},
        "provisioning_state": {"key": "properties.provisioningState", "type": "str"},
    }

    def __init__(
        self,
        *,
        service_name: Optional[Union[str, "_models.ServiceName"]] = None,
        resource_id: Optional[str] = None,
        port: Optional[int] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword service_name: Name of the service. Known values are: "SSH" and "WAC".
        :paramtype service_name: str or ~azure.mgmt.hybridconnectivity.models.ServiceName
        :keyword resource_id: The resource Id of the connectivity endpoint (optional).
        :paramtype resource_id: str
        :keyword port: The port on which service is enabled.
        :paramtype port: int
        """
        super().__init__(**kwargs)
        self.service_name = service_name
        self.resource_id = resource_id
        self.port = port
        self.provisioning_state = None


class ServiceConfigurationResourcePatch(_serialization.Model):
    """The service details under service configuration for the target endpoint resource.

    :ivar port: The port on which service is enabled.
    :vartype port: int
    """

    _attribute_map = {
        "port": {"key": "properties.port", "type": "int"},
    }

    def __init__(self, *, port: Optional[int] = None, **kwargs: Any) -> None:
        """
        :keyword port: The port on which service is enabled.
        :paramtype port: int
        """
        super().__init__(**kwargs)
        self.port = port


class SystemData(_serialization.Model):
    """Metadata pertaining to creation and last modification of the resource.

    :ivar created_by: The identity that created the resource.
    :vartype created_by: str
    :ivar created_by_type: The type of identity that created the resource. Known values are:
     "User", "Application", "ManagedIdentity", and "Key".
    :vartype created_by_type: str or ~azure.mgmt.hybridconnectivity.models.CreatedByType
    :ivar created_at: The timestamp of resource creation (UTC).
    :vartype created_at: ~datetime.datetime
    :ivar last_modified_by: The identity that last modified the resource.
    :vartype last_modified_by: str
    :ivar last_modified_by_type: The type of identity that last modified the resource. Known values
     are: "User", "Application", "ManagedIdentity", and "Key".
    :vartype last_modified_by_type: str or ~azure.mgmt.hybridconnectivity.models.CreatedByType
    :ivar last_modified_at: The timestamp of resource last modification (UTC).
    :vartype last_modified_at: ~datetime.datetime
    """

    _attribute_map = {
        "created_by": {"key": "createdBy", "type": "str"},
        "created_by_type": {"key": "createdByType", "type": "str"},
        "created_at": {"key": "createdAt", "type": "iso-8601"},
        "last_modified_by": {"key": "lastModifiedBy", "type": "str"},
        "last_modified_by_type": {"key": "lastModifiedByType", "type": "str"},
        "last_modified_at": {"key": "lastModifiedAt", "type": "iso-8601"},
    }

    def __init__(
        self,
        *,
        created_by: Optional[str] = None,
        created_by_type: Optional[Union[str, "_models.CreatedByType"]] = None,
        created_at: Optional[datetime.datetime] = None,
        last_modified_by: Optional[str] = None,
        last_modified_by_type: Optional[Union[str, "_models.CreatedByType"]] = None,
        last_modified_at: Optional[datetime.datetime] = None,
        **kwargs: Any
    ) -> None:
        """
        :keyword created_by: The identity that created the resource.
        :paramtype created_by: str
        :keyword created_by_type: The type of identity that created the resource. Known values are:
         "User", "Application", "ManagedIdentity", and "Key".
        :paramtype created_by_type: str or ~azure.mgmt.hybridconnectivity.models.CreatedByType
        :keyword created_at: The timestamp of resource creation (UTC).
        :paramtype created_at: ~datetime.datetime
        :keyword last_modified_by: The identity that last modified the resource.
        :paramtype last_modified_by: str
        :keyword last_modified_by_type: The type of identity that last modified the resource. Known
         values are: "User", "Application", "ManagedIdentity", and "Key".
        :paramtype last_modified_by_type: str or ~azure.mgmt.hybridconnectivity.models.CreatedByType
        :keyword last_modified_at: The timestamp of resource last modification (UTC).
        :paramtype last_modified_at: ~datetime.datetime
        """
        super().__init__(**kwargs)
        self.created_by = created_by
        self.created_by_type = created_by_type
        self.created_at = created_at
        self.last_modified_by = last_modified_by
        self.last_modified_by_type = last_modified_by_type
        self.last_modified_at = last_modified_at
