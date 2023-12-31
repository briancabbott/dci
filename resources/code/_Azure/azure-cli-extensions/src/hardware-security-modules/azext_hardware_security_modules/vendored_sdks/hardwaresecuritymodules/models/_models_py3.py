# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import datetime
from typing import Dict, List, Optional, Union

from azure.core.exceptions import HttpResponseError
import msrest.serialization

from ._azure_dedicated_hsm_resource_provider_enums import *


class ApiEntityReference(msrest.serialization.Model):
    """The API entity reference.

    :param id: The ARM resource id in the form of
     /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
    :type id: str
    """

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        id: Optional[str] = None,
        **kwargs
    ):
        super(ApiEntityReference, self).__init__(**kwargs)
        self.id = id


class Resource(msrest.serialization.Model):
    """Dedicated HSM resource.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: The Azure Resource Manager resource ID for the dedicated HSM.
    :vartype id: str
    :ivar name: The name of the dedicated HSM.
    :vartype name: str
    :ivar type: The resource type of the dedicated HSM.
    :vartype type: str
    :param location: Required. The supported Azure location where the dedicated HSM should be
     created.
    :type location: str
    :param sku: SKU details.
    :type sku: ~azure_dedicated_hsm_resource_provider.models.Sku
    :param zones: The Dedicated Hsm zones.
    :type zones: list[str]
    :param tags: A set of tags. Resource tags.
    :type tags: dict[str, str]
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'location': {'key': 'location', 'type': 'str'},
        'sku': {'key': 'sku', 'type': 'Sku'},
        'zones': {'key': 'zones', 'type': '[str]'},
        'tags': {'key': 'tags', 'type': '{str}'},
    }

    def __init__(
        self,
        *,
        location: str,
        sku: Optional["Sku"] = None,
        zones: Optional[List[str]] = None,
        tags: Optional[Dict[str, str]] = None,
        **kwargs
    ):
        super(Resource, self).__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None
        self.location = location
        self.sku = sku
        self.zones = zones
        self.tags = tags


class DedicatedHsm(Resource):
    """Resource information with extended details.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: The Azure Resource Manager resource ID for the dedicated HSM.
    :vartype id: str
    :ivar name: The name of the dedicated HSM.
    :vartype name: str
    :ivar type: The resource type of the dedicated HSM.
    :vartype type: str
    :param location: Required. The supported Azure location where the dedicated HSM should be
     created.
    :type location: str
    :param sku: SKU details.
    :type sku: ~azure_dedicated_hsm_resource_provider.models.Sku
    :param zones: The Dedicated Hsm zones.
    :type zones: list[str]
    :param tags: A set of tags. Resource tags.
    :type tags: dict[str, str]
    :ivar system_data: Metadata pertaining to creation and last modification of the resource.
    :vartype system_data: ~azure_dedicated_hsm_resource_provider.models.SystemData
    :param network_profile: Specifies the network interfaces of the dedicated hsm.
    :type network_profile: ~azure_dedicated_hsm_resource_provider.models.NetworkProfile
    :param management_network_profile: Specifies the management network interfaces of the dedicated
     hsm.
    :type management_network_profile: ~azure_dedicated_hsm_resource_provider.models.NetworkProfile
    :param stamp_id: This field will be used when RP does not support Availability zones.
    :type stamp_id: str
    :ivar status_message: Resource Status Message.
    :vartype status_message: str
    :ivar provisioning_state: Provisioning state. Possible values include: "Succeeded",
     "Provisioning", "Allocating", "Connecting", "Failed", "CheckingQuota", "Deleting".
    :vartype provisioning_state: str or
     ~azure_dedicated_hsm_resource_provider.models.JsonWebKeyType
    """

    _validation = {
        'id': {'readonly': True},
        'name': {'readonly': True},
        'type': {'readonly': True},
        'location': {'required': True},
        'system_data': {'readonly': True},
        'status_message': {'readonly': True},
        'provisioning_state': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'name': {'key': 'name', 'type': 'str'},
        'type': {'key': 'type', 'type': 'str'},
        'location': {'key': 'location', 'type': 'str'},
        'sku': {'key': 'sku', 'type': 'Sku'},
        'zones': {'key': 'zones', 'type': '[str]'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'system_data': {'key': 'systemData', 'type': 'SystemData'},
        'network_profile': {'key': 'properties.networkProfile', 'type': 'NetworkProfile'},
        'management_network_profile': {'key': 'properties.managementNetworkProfile', 'type': 'NetworkProfile'},
        'stamp_id': {'key': 'properties.stampId', 'type': 'str'},
        'status_message': {'key': 'properties.statusMessage', 'type': 'str'},
        'provisioning_state': {'key': 'properties.provisioningState', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        location: str,
        sku: Optional["Sku"] = None,
        zones: Optional[List[str]] = None,
        tags: Optional[Dict[str, str]] = None,
        network_profile: Optional["NetworkProfile"] = None,
        management_network_profile: Optional["NetworkProfile"] = None,
        stamp_id: Optional[str] = None,
        **kwargs
    ):
        super(DedicatedHsm, self).__init__(location=location, sku=sku, zones=zones, tags=tags, **kwargs)
        self.system_data = None
        self.network_profile = network_profile
        self.management_network_profile = management_network_profile
        self.stamp_id = stamp_id
        self.status_message = None
        self.provisioning_state = None


class DedicatedHsmError(msrest.serialization.Model):
    """The error exception.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar error: The error detail of the operation if any.
    :vartype error: ~azure_dedicated_hsm_resource_provider.models.Error
    """

    _validation = {
        'error': {'readonly': True},
    }

    _attribute_map = {
        'error': {'key': 'error', 'type': 'Error'},
    }

    def __init__(
        self,
        **kwargs
    ):
        super(DedicatedHsmError, self).__init__(**kwargs)
        self.error = None


class DedicatedHsmListResult(msrest.serialization.Model):
    """List of dedicated HSMs.

    :param value: The list of dedicated HSMs.
    :type value: list[~azure_dedicated_hsm_resource_provider.models.DedicatedHsm]
    :param next_link: The URL to get the next set of dedicated hsms.
    :type next_link: str
    """

    _attribute_map = {
        'value': {'key': 'value', 'type': '[DedicatedHsm]'},
        'next_link': {'key': 'nextLink', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        value: Optional[List["DedicatedHsm"]] = None,
        next_link: Optional[str] = None,
        **kwargs
    ):
        super(DedicatedHsmListResult, self).__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class DedicatedHsmOperation(msrest.serialization.Model):
    """REST API operation.

    Variables are only populated by the server, and will be ignored when sending a request.

    :param name: The name of the Dedicated HSM Resource Provider Operation.
    :type name: str
    :ivar is_data_action: Gets or sets a value indicating whether it is a data plane action.
    :vartype is_data_action: str
    :param display: The display string.
    :type display: ~azure_dedicated_hsm_resource_provider.models.DedicatedHsmOperationDisplay
    """

    _validation = {
        'is_data_action': {'readonly': True},
    }

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
        'is_data_action': {'key': 'isDataAction', 'type': 'str'},
        'display': {'key': 'display', 'type': 'DedicatedHsmOperationDisplay'},
    }

    def __init__(
        self,
        *,
        name: Optional[str] = None,
        display: Optional["DedicatedHsmOperationDisplay"] = None,
        **kwargs
    ):
        super(DedicatedHsmOperation, self).__init__(**kwargs)
        self.name = name
        self.is_data_action = None
        self.display = display


class DedicatedHsmOperationDisplay(msrest.serialization.Model):
    """The display string.

    :param provider: The Resource Provider of the operation.
    :type provider: str
    :param resource: Resource on which the operation is performed.
    :type resource: str
    :param operation: Operation type: Read, write, delete, etc.
    :type operation: str
    :param description: The object that represents the operation.
    :type description: str
    """

    _attribute_map = {
        'provider': {'key': 'provider', 'type': 'str'},
        'resource': {'key': 'resource', 'type': 'str'},
        'operation': {'key': 'operation', 'type': 'str'},
        'description': {'key': 'description', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        provider: Optional[str] = None,
        resource: Optional[str] = None,
        operation: Optional[str] = None,
        description: Optional[str] = None,
        **kwargs
    ):
        super(DedicatedHsmOperationDisplay, self).__init__(**kwargs)
        self.provider = provider
        self.resource = resource
        self.operation = operation
        self.description = description


class DedicatedHsmOperationListResult(msrest.serialization.Model):
    """Result of the request to list Dedicated HSM Provider operations. It contains a list of operations.

    :param value: List of Dedicated HSM Resource Provider operations.
    :type value: list[~azure_dedicated_hsm_resource_provider.models.DedicatedHsmOperation]
    """

    _attribute_map = {
        'value': {'key': 'value', 'type': '[DedicatedHsmOperation]'},
    }

    def __init__(
        self,
        *,
        value: Optional[List["DedicatedHsmOperation"]] = None,
        **kwargs
    ):
        super(DedicatedHsmOperationListResult, self).__init__(**kwargs)
        self.value = value


class DedicatedHsmPatchParameters(msrest.serialization.Model):
    """Patchable properties of the dedicated HSM.

    :param tags: A set of tags. Resource tags.
    :type tags: dict[str, str]
    """

    _attribute_map = {
        'tags': {'key': 'tags', 'type': '{str}'},
    }

    def __init__(
        self,
        *,
        tags: Optional[Dict[str, str]] = None,
        **kwargs
    ):
        super(DedicatedHsmPatchParameters, self).__init__(**kwargs)
        self.tags = tags


class EndpointDependency(msrest.serialization.Model):
    """A domain name that dedicated hsm services are reaching at.

    :param domain_name: The domain name of the dependency.
    :type domain_name: str
    :param endpoint_details: The Ports and Protocols used when connecting to domainName.
    :type endpoint_details: list[~azure_dedicated_hsm_resource_provider.models.EndpointDetail]
    """

    _attribute_map = {
        'domain_name': {'key': 'domainName', 'type': 'str'},
        'endpoint_details': {'key': 'endpointDetails', 'type': '[EndpointDetail]'},
    }

    def __init__(
        self,
        *,
        domain_name: Optional[str] = None,
        endpoint_details: Optional[List["EndpointDetail"]] = None,
        **kwargs
    ):
        super(EndpointDependency, self).__init__(**kwargs)
        self.domain_name = domain_name
        self.endpoint_details = endpoint_details


class EndpointDetail(msrest.serialization.Model):
    """Connect information from the dedicated hsm service to a single endpoint.

    :param ip_address: An IP Address that Domain Name currently resolves to.
    :type ip_address: str
    :param port: The port an endpoint is connected to.
    :type port: int
    :param protocol: The protocol used for connection.
    :type protocol: str
    :param description: Description of the detail.
    :type description: str
    """

    _attribute_map = {
        'ip_address': {'key': 'ipAddress', 'type': 'str'},
        'port': {'key': 'port', 'type': 'int'},
        'protocol': {'key': 'protocol', 'type': 'str'},
        'description': {'key': 'description', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        ip_address: Optional[str] = None,
        port: Optional[int] = None,
        protocol: Optional[str] = None,
        description: Optional[str] = None,
        **kwargs
    ):
        super(EndpointDetail, self).__init__(**kwargs)
        self.ip_address = ip_address
        self.port = port
        self.protocol = protocol
        self.description = description


class Error(msrest.serialization.Model):
    """The key vault server error.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar code: The error code.
    :vartype code: str
    :ivar message: The error message.
    :vartype message: str
    :ivar inner_error: Contains more specific error that narrows down the cause. May be null.
    :vartype inner_error: ~azure_dedicated_hsm_resource_provider.models.Error
    """

    _validation = {
        'code': {'readonly': True},
        'message': {'readonly': True},
        'inner_error': {'readonly': True},
    }

    _attribute_map = {
        'code': {'key': 'code', 'type': 'str'},
        'message': {'key': 'message', 'type': 'str'},
        'inner_error': {'key': 'innererror', 'type': 'Error'},
    }

    def __init__(
        self,
        **kwargs
    ):
        super(Error, self).__init__(**kwargs)
        self.code = None
        self.message = None
        self.inner_error = None


class NetworkInterface(msrest.serialization.Model):
    """The network interface definition.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: The ARM resource id in the form of
     /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
    :vartype id: str
    :param private_ip_address: Private Ip address of the interface.
    :type private_ip_address: str
    """

    _validation = {
        'id': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'private_ip_address': {'key': 'privateIpAddress', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        private_ip_address: Optional[str] = None,
        **kwargs
    ):
        super(NetworkInterface, self).__init__(**kwargs)
        self.id = None
        self.private_ip_address = private_ip_address


class NetworkProfile(msrest.serialization.Model):
    """The network profile definition.

    :param subnet: Specifies the identifier of the subnet.
    :type subnet: ~azure_dedicated_hsm_resource_provider.models.ApiEntityReference
    :param network_interfaces: Specifies the list of resource Ids for the network interfaces
     associated with the dedicated HSM.
    :type network_interfaces: list[~azure_dedicated_hsm_resource_provider.models.NetworkInterface]
    """

    _attribute_map = {
        'subnet': {'key': 'subnet', 'type': 'ApiEntityReference'},
        'network_interfaces': {'key': 'networkInterfaces', 'type': '[NetworkInterface]'},
    }

    def __init__(
        self,
        *,
        subnet: Optional["ApiEntityReference"] = None,
        network_interfaces: Optional[List["NetworkInterface"]] = None,
        **kwargs
    ):
        super(NetworkProfile, self).__init__(**kwargs)
        self.subnet = subnet
        self.network_interfaces = network_interfaces


class OutboundEnvironmentEndpoint(msrest.serialization.Model):
    """Egress endpoints which dedicated hsm service connects to for common purpose.

    :param category: The category of endpoints accessed by the dedicated hsm service, e.g. azure-
     resource-management, apiserver, etc.
    :type category: str
    :param endpoints: The endpoints that dedicated hsm service connects to.
    :type endpoints: list[~azure_dedicated_hsm_resource_provider.models.EndpointDependency]
    """

    _attribute_map = {
        'category': {'key': 'category', 'type': 'str'},
        'endpoints': {'key': 'endpoints', 'type': '[EndpointDependency]'},
    }

    def __init__(
        self,
        *,
        category: Optional[str] = None,
        endpoints: Optional[List["EndpointDependency"]] = None,
        **kwargs
    ):
        super(OutboundEnvironmentEndpoint, self).__init__(**kwargs)
        self.category = category
        self.endpoints = endpoints


class OutboundEnvironmentEndpointCollection(msrest.serialization.Model):
    """Collection of OutboundEnvironmentEndpoint.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :param value: Required. Collection of resources.
    :type value: list[~azure_dedicated_hsm_resource_provider.models.OutboundEnvironmentEndpoint]
    :ivar next_link: Link to next page of resources.
    :vartype next_link: str
    """

    _validation = {
        'value': {'required': True},
        'next_link': {'readonly': True},
    }

    _attribute_map = {
        'value': {'key': 'value', 'type': '[OutboundEnvironmentEndpoint]'},
        'next_link': {'key': 'nextLink', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        value: List["OutboundEnvironmentEndpoint"],
        **kwargs
    ):
        super(OutboundEnvironmentEndpointCollection, self).__init__(**kwargs)
        self.value = value
        self.next_link = None


class ResourceListResult(msrest.serialization.Model):
    """List of dedicated HSM resources.

    :param value: The list of dedicated HSM resources.
    :type value: list[~azure_dedicated_hsm_resource_provider.models.Resource]
    :param next_link: The URL to get the next set of dedicated HSM resources.
    :type next_link: str
    """

    _attribute_map = {
        'value': {'key': 'value', 'type': '[Resource]'},
        'next_link': {'key': 'nextLink', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        value: Optional[List["Resource"]] = None,
        next_link: Optional[str] = None,
        **kwargs
    ):
        super(ResourceListResult, self).__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class Sku(msrest.serialization.Model):
    """SKU of the dedicated HSM.

    :param name: SKU of the dedicated HSM. Possible values include: "SafeNet Luna Network HSM
     A790", "payShield10K_LMK1_CPS60", "payShield10K_LMK1_CPS250", "payShield10K_LMK1_CPS2500",
     "payShield10K_LMK2_CPS60", "payShield10K_LMK2_CPS250", "payShield10K_LMK2_CPS2500".
    :type name: str or ~azure_dedicated_hsm_resource_provider.models.SkuName
    """

    _attribute_map = {
        'name': {'key': 'name', 'type': 'str'},
    }

    def __init__(
        self,
        *,
        name: Optional[Union[str, "SkuName"]] = None,
        **kwargs
    ):
        super(Sku, self).__init__(**kwargs)
        self.name = name


class SystemData(msrest.serialization.Model):
    """Metadata pertaining to creation and last modification of dedicated hsm resource.

    :param created_by: The identity that created dedicated hsm resource.
    :type created_by: str
    :param created_by_type: The type of identity that created dedicated hsm resource. Possible
     values include: "User", "Application", "ManagedIdentity", "Key".
    :type created_by_type: str or ~azure_dedicated_hsm_resource_provider.models.IdentityType
    :param created_at: The timestamp of dedicated hsm resource creation (UTC).
    :type created_at: ~datetime.datetime
    :param last_modified_by: The identity that last modified dedicated hsm resource.
    :type last_modified_by: str
    :param last_modified_by_type: The type of identity that last modified dedicated hsm resource.
     Possible values include: "User", "Application", "ManagedIdentity", "Key".
    :type last_modified_by_type: str or ~azure_dedicated_hsm_resource_provider.models.IdentityType
    :param last_modified_at: The timestamp of dedicated hsm resource last modification (UTC).
    :type last_modified_at: ~datetime.datetime
    """

    _attribute_map = {
        'created_by': {'key': 'createdBy', 'type': 'str'},
        'created_by_type': {'key': 'createdByType', 'type': 'str'},
        'created_at': {'key': 'createdAt', 'type': 'iso-8601'},
        'last_modified_by': {'key': 'lastModifiedBy', 'type': 'str'},
        'last_modified_by_type': {'key': 'lastModifiedByType', 'type': 'str'},
        'last_modified_at': {'key': 'lastModifiedAt', 'type': 'iso-8601'},
    }

    def __init__(
        self,
        *,
        created_by: Optional[str] = None,
        created_by_type: Optional[Union[str, "IdentityType"]] = None,
        created_at: Optional[datetime.datetime] = None,
        last_modified_by: Optional[str] = None,
        last_modified_by_type: Optional[Union[str, "IdentityType"]] = None,
        last_modified_at: Optional[datetime.datetime] = None,
        **kwargs
    ):
        super(SystemData, self).__init__(**kwargs)
        self.created_by = created_by
        self.created_by_type = created_by_type
        self.created_at = created_at
        self.last_modified_by = last_modified_by
        self.last_modified_by_type = last_modified_by_type
        self.last_modified_at = last_modified_at
