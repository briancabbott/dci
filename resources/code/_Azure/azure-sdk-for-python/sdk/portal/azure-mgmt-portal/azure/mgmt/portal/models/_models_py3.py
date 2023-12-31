# coding=utf-8
# pylint: disable=too-many-lines
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

import sys
from typing import Any, Dict, List, Optional, TYPE_CHECKING

from .. import _serialization

if sys.version_info >= (3, 9):
    from collections.abc import MutableMapping
else:
    from typing import MutableMapping  # type: ignore  # pylint: disable=ungrouped-imports

if TYPE_CHECKING:
    # pylint: disable=unused-import,ungrouped-imports
    from .. import models as _models
JSON = MutableMapping[str, Any]  # pylint: disable=unsubscriptable-object


class Resource(_serialization.Model):
    """Common fields that are returned in the response for all Azure Resource Manager resources.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
    }

    def __init__(self, **kwargs):
        """ """
        super().__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None


class ProxyResource(Resource):
    """The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a location.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
    }

    def __init__(self, **kwargs):
        """ """
        super().__init__(**kwargs)


class Configuration(ProxyResource):
    """Tenant configuration.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Fully qualified resource ID for the resource. Ex -
     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}.
    :vartype id: str
    :ivar name: The name of the resource.
    :vartype name: str
    :ivar type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
     "Microsoft.Storage/storageAccounts".
    :vartype type: str
    :ivar enforce_private_markdown_storage: When flag is set to true Markdown tile will require
     external storage configuration (URI). The inline content configuration will be prohibited.
    :vartype enforce_private_markdown_storage: bool
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "enforce_private_markdown_storage": {"key": "properties.enforcePrivateMarkdownStorage", "type": "bool"},
    }

    def __init__(self, *, enforce_private_markdown_storage: Optional[bool] = None, **kwargs):
        """
        :keyword enforce_private_markdown_storage: When flag is set to true Markdown tile will require
         external storage configuration (URI). The inline content configuration will be prohibited.
        :paramtype enforce_private_markdown_storage: bool
        """
        super().__init__(**kwargs)
        self.enforce_private_markdown_storage = enforce_private_markdown_storage


class ConfigurationList(_serialization.Model):
    """List of tenant configurations.

    :ivar value: The array of tenant configurations.
    :vartype value: list[~azure.mgmt.portal.models.Configuration]
    :ivar next_link: The URL to use for getting the next set of results.
    :vartype next_link: str
    """

    _attribute_map = {
        "value": {"key": "value", "type": "[Configuration]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(
        self, *, value: Optional[List["_models.Configuration"]] = None, next_link: Optional[str] = None, **kwargs
    ):
        """
        :keyword value: The array of tenant configurations.
        :paramtype value: list[~azure.mgmt.portal.models.Configuration]
        :keyword next_link: The URL to use for getting the next set of results.
        :paramtype next_link: str
        """
        super().__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class Dashboard(_serialization.Model):
    """The shared dashboard resource definition.

    Variables are only populated by the server, and will be ignored when sending a request.

    All required parameters must be populated in order to send to Azure.

    :ivar id: Resource Id.
    :vartype id: str
    :ivar name: Resource name.
    :vartype name: str
    :ivar type: Resource type.
    :vartype type: str
    :ivar location: Resource location. Required.
    :vartype location: str
    :ivar tags: Resource tags.
    :vartype tags: dict[str, str]
    :ivar lenses: The dashboard lenses.
    :vartype lenses: list[~azure.mgmt.portal.models.DashboardLens]
    :ivar metadata: The dashboard metadata.
    :vartype metadata: dict[str, JSON]
    """

    _validation = {
        "id": {"readonly": True},
        "name": {"readonly": True},
        "type": {"readonly": True},
        "location": {"required": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "name": {"key": "name", "type": "str"},
        "type": {"key": "type", "type": "str"},
        "location": {"key": "location", "type": "str"},
        "tags": {"key": "tags", "type": "{str}"},
        "lenses": {"key": "properties.lenses", "type": "[DashboardLens]"},
        "metadata": {"key": "properties.metadata", "type": "{object}"},
    }

    def __init__(
        self,
        *,
        location: str,
        tags: Optional[Dict[str, str]] = None,
        lenses: Optional[List["_models.DashboardLens"]] = None,
        metadata: Optional[Dict[str, JSON]] = None,
        **kwargs
    ):
        """
        :keyword location: Resource location. Required.
        :paramtype location: str
        :keyword tags: Resource tags.
        :paramtype tags: dict[str, str]
        :keyword lenses: The dashboard lenses.
        :paramtype lenses: list[~azure.mgmt.portal.models.DashboardLens]
        :keyword metadata: The dashboard metadata.
        :paramtype metadata: dict[str, JSON]
        """
        super().__init__(**kwargs)
        self.id = None
        self.name = None
        self.type = None
        self.location = location
        self.tags = tags
        self.lenses = lenses
        self.metadata = metadata


class DashboardLens(_serialization.Model):
    """A dashboard lens.

    All required parameters must be populated in order to send to Azure.

    :ivar order: The lens order. Required.
    :vartype order: int
    :ivar parts: The dashboard parts. Required.
    :vartype parts: list[~azure.mgmt.portal.models.DashboardParts]
    :ivar metadata: The dashboard len's metadata.
    :vartype metadata: dict[str, JSON]
    """

    _validation = {
        "order": {"required": True},
        "parts": {"required": True},
    }

    _attribute_map = {
        "order": {"key": "order", "type": "int"},
        "parts": {"key": "parts", "type": "[DashboardParts]"},
        "metadata": {"key": "metadata", "type": "{object}"},
    }

    def __init__(
        self, *, order: int, parts: List["_models.DashboardParts"], metadata: Optional[Dict[str, JSON]] = None, **kwargs
    ):
        """
        :keyword order: The lens order. Required.
        :paramtype order: int
        :keyword parts: The dashboard parts. Required.
        :paramtype parts: list[~azure.mgmt.portal.models.DashboardParts]
        :keyword metadata: The dashboard len's metadata.
        :paramtype metadata: dict[str, JSON]
        """
        super().__init__(**kwargs)
        self.order = order
        self.parts = parts
        self.metadata = metadata


class DashboardListResult(_serialization.Model):
    """List of dashboards.

    :ivar value: The array of custom resource provider manifests.
    :vartype value: list[~azure.mgmt.portal.models.Dashboard]
    :ivar next_link: The URL to use for getting the next set of results.
    :vartype next_link: str
    """

    _attribute_map = {
        "value": {"key": "value", "type": "[Dashboard]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(self, *, value: Optional[List["_models.Dashboard"]] = None, next_link: Optional[str] = None, **kwargs):
        """
        :keyword value: The array of custom resource provider manifests.
        :paramtype value: list[~azure.mgmt.portal.models.Dashboard]
        :keyword next_link: The URL to use for getting the next set of results.
        :paramtype next_link: str
        """
        super().__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class DashboardPartMetadata(_serialization.Model):
    """A dashboard part metadata.

    You probably want to use the sub-classes and not this class directly. Known sub-classes are:
    MarkdownPartMetadata

    All required parameters must be populated in order to send to Azure.

    :ivar additional_properties: Unmatched properties from the message are deserialized to this
     collection.
    :vartype additional_properties: dict[str, any]
    :ivar type: The type of dashboard part. Required.
    :vartype type: str
    """

    _validation = {
        "type": {"required": True},
    }

    _attribute_map = {
        "additional_properties": {"key": "", "type": "{object}"},
        "type": {"key": "type", "type": "str"},
    }

    _subtype_map = {"type": {"Extension/HubsExtension/PartType/MarkdownPart": "MarkdownPartMetadata"}}

    def __init__(self, *, additional_properties: Optional[Dict[str, Any]] = None, **kwargs):
        """
        :keyword additional_properties: Unmatched properties from the message are deserialized to this
         collection.
        :paramtype additional_properties: dict[str, any]
        """
        super().__init__(**kwargs)
        self.additional_properties = additional_properties
        self.type = None  # type: Optional[str]


class DashboardParts(_serialization.Model):
    """A dashboard part.

    All required parameters must be populated in order to send to Azure.

    :ivar position: The dashboard's part position. Required.
    :vartype position: ~azure.mgmt.portal.models.DashboardPartsPosition
    :ivar metadata: The dashboard part's metadata.
    :vartype metadata: ~azure.mgmt.portal.models.DashboardPartMetadata
    """

    _validation = {
        "position": {"required": True},
    }

    _attribute_map = {
        "position": {"key": "position", "type": "DashboardPartsPosition"},
        "metadata": {"key": "metadata", "type": "DashboardPartMetadata"},
    }

    def __init__(
        self,
        *,
        position: "_models.DashboardPartsPosition",
        metadata: Optional["_models.DashboardPartMetadata"] = None,
        **kwargs
    ):
        """
        :keyword position: The dashboard's part position. Required.
        :paramtype position: ~azure.mgmt.portal.models.DashboardPartsPosition
        :keyword metadata: The dashboard part's metadata.
        :paramtype metadata: ~azure.mgmt.portal.models.DashboardPartMetadata
        """
        super().__init__(**kwargs)
        self.position = position
        self.metadata = metadata


class DashboardPartsPosition(_serialization.Model):
    """The dashboard's part position.

    All required parameters must be populated in order to send to Azure.

    :ivar x: The dashboard's part x coordinate. Required.
    :vartype x: int
    :ivar y: The dashboard's part y coordinate. Required.
    :vartype y: int
    :ivar row_span: The dashboard's part row span. Required.
    :vartype row_span: int
    :ivar col_span: The dashboard's part column span. Required.
    :vartype col_span: int
    :ivar metadata: The dashboard part's metadata.
    :vartype metadata: dict[str, JSON]
    """

    _validation = {
        "x": {"required": True},
        "y": {"required": True},
        "row_span": {"required": True},
        "col_span": {"required": True},
    }

    _attribute_map = {
        "x": {"key": "x", "type": "int"},
        "y": {"key": "y", "type": "int"},
        "row_span": {"key": "rowSpan", "type": "int"},
        "col_span": {"key": "colSpan", "type": "int"},
        "metadata": {"key": "metadata", "type": "{object}"},
    }

    def __init__(
        self, *, x: int, y: int, row_span: int, col_span: int, metadata: Optional[Dict[str, JSON]] = None, **kwargs
    ):
        """
        :keyword x: The dashboard's part x coordinate. Required.
        :paramtype x: int
        :keyword y: The dashboard's part y coordinate. Required.
        :paramtype y: int
        :keyword row_span: The dashboard's part row span. Required.
        :paramtype row_span: int
        :keyword col_span: The dashboard's part column span. Required.
        :paramtype col_span: int
        :keyword metadata: The dashboard part's metadata.
        :paramtype metadata: dict[str, JSON]
        """
        super().__init__(**kwargs)
        self.x = x
        self.y = y
        self.row_span = row_span
        self.col_span = col_span
        self.metadata = metadata


class ErrorDefinition(_serialization.Model):
    """Error definition.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar code: Service specific error code which serves as the substatus for the HTTP error code.
    :vartype code: int
    :ivar message: Description of the error.
    :vartype message: str
    :ivar details: Internal error details.
    :vartype details: list[~azure.mgmt.portal.models.ErrorDefinition]
    """

    _validation = {
        "code": {"readonly": True},
        "message": {"readonly": True},
        "details": {"readonly": True},
    }

    _attribute_map = {
        "code": {"key": "code", "type": "int"},
        "message": {"key": "message", "type": "str"},
        "details": {"key": "details", "type": "[ErrorDefinition]"},
    }

    def __init__(self, **kwargs):
        """ """
        super().__init__(**kwargs)
        self.code = None
        self.message = None
        self.details = None


class ErrorResponse(_serialization.Model):
    """Error response.

    :ivar error: The error details.
    :vartype error: ~azure.mgmt.portal.models.ErrorDefinition
    """

    _attribute_map = {
        "error": {"key": "error", "type": "ErrorDefinition"},
    }

    def __init__(self, *, error: Optional["_models.ErrorDefinition"] = None, **kwargs):
        """
        :keyword error: The error details.
        :paramtype error: ~azure.mgmt.portal.models.ErrorDefinition
        """
        super().__init__(**kwargs)
        self.error = error


class MarkdownPartMetadata(DashboardPartMetadata):
    """Markdown part metadata.

    All required parameters must be populated in order to send to Azure.

    :ivar additional_properties: Unmatched properties from the message are deserialized to this
     collection.
    :vartype additional_properties: dict[str, any]
    :ivar type: The type of dashboard part. Required.
    :vartype type: str
    :ivar inputs: Input to dashboard part.
    :vartype inputs: list[JSON]
    :ivar settings: Markdown part settings.
    :vartype settings: ~azure.mgmt.portal.models.MarkdownPartMetadataSettings
    """

    _validation = {
        "type": {"required": True},
    }

    _attribute_map = {
        "additional_properties": {"key": "", "type": "{object}"},
        "type": {"key": "type", "type": "str"},
        "inputs": {"key": "inputs", "type": "[object]"},
        "settings": {"key": "settings", "type": "MarkdownPartMetadataSettings"},
    }

    def __init__(
        self,
        *,
        additional_properties: Optional[Dict[str, Any]] = None,
        inputs: Optional[List[JSON]] = None,
        settings: Optional["_models.MarkdownPartMetadataSettings"] = None,
        **kwargs
    ):
        """
        :keyword additional_properties: Unmatched properties from the message are deserialized to this
         collection.
        :paramtype additional_properties: dict[str, any]
        :keyword inputs: Input to dashboard part.
        :paramtype inputs: list[JSON]
        :keyword settings: Markdown part settings.
        :paramtype settings: ~azure.mgmt.portal.models.MarkdownPartMetadataSettings
        """
        super().__init__(additional_properties=additional_properties, **kwargs)
        self.type = "Extension/HubsExtension/PartType/MarkdownPart"  # type: str
        self.inputs = inputs
        self.settings = settings


class MarkdownPartMetadataSettings(_serialization.Model):
    """Markdown part settings.

    :ivar content: The content of markdown part.
    :vartype content: ~azure.mgmt.portal.models.MarkdownPartMetadataSettingsContent
    """

    _attribute_map = {
        "content": {"key": "content", "type": "MarkdownPartMetadataSettingsContent"},
    }

    def __init__(self, *, content: Optional["_models.MarkdownPartMetadataSettingsContent"] = None, **kwargs):
        """
        :keyword content: The content of markdown part.
        :paramtype content: ~azure.mgmt.portal.models.MarkdownPartMetadataSettingsContent
        """
        super().__init__(**kwargs)
        self.content = content


class MarkdownPartMetadataSettingsContent(_serialization.Model):
    """The content of markdown part.

    :ivar settings: The setting of the content of markdown part.
    :vartype settings: ~azure.mgmt.portal.models.MarkdownPartMetadataSettingsContentSettings
    """

    _attribute_map = {
        "settings": {"key": "settings", "type": "MarkdownPartMetadataSettingsContentSettings"},
    }

    def __init__(self, *, settings: Optional["_models.MarkdownPartMetadataSettingsContentSettings"] = None, **kwargs):
        """
        :keyword settings: The setting of the content of markdown part.
        :paramtype settings: ~azure.mgmt.portal.models.MarkdownPartMetadataSettingsContentSettings
        """
        super().__init__(**kwargs)
        self.settings = settings


class MarkdownPartMetadataSettingsContentSettings(_serialization.Model):
    """The setting of the content of markdown part.

    :ivar content: The content of the markdown part.
    :vartype content: str
    :ivar title: The title of the markdown part.
    :vartype title: str
    :ivar subtitle: The subtitle of the markdown part.
    :vartype subtitle: str
    :ivar markdown_source: The source of the content of the markdown part.
    :vartype markdown_source: int
    :ivar markdown_uri: The uri of markdown content.
    :vartype markdown_uri: str
    """

    _attribute_map = {
        "content": {"key": "content", "type": "str"},
        "title": {"key": "title", "type": "str"},
        "subtitle": {"key": "subtitle", "type": "str"},
        "markdown_source": {"key": "markdownSource", "type": "int"},
        "markdown_uri": {"key": "markdownUri", "type": "str"},
    }

    def __init__(
        self,
        *,
        content: Optional[str] = None,
        title: Optional[str] = None,
        subtitle: Optional[str] = None,
        markdown_source: Optional[int] = None,
        markdown_uri: Optional[str] = None,
        **kwargs
    ):
        """
        :keyword content: The content of the markdown part.
        :paramtype content: str
        :keyword title: The title of the markdown part.
        :paramtype title: str
        :keyword subtitle: The subtitle of the markdown part.
        :paramtype subtitle: str
        :keyword markdown_source: The source of the content of the markdown part.
        :paramtype markdown_source: int
        :keyword markdown_uri: The uri of markdown content.
        :paramtype markdown_uri: str
        """
        super().__init__(**kwargs)
        self.content = content
        self.title = title
        self.subtitle = subtitle
        self.markdown_source = markdown_source
        self.markdown_uri = markdown_uri


class PatchableDashboard(_serialization.Model):
    """The shared dashboard resource definition.

    :ivar tags: Resource tags.
    :vartype tags: dict[str, str]
    :ivar lenses: The dashboard lenses.
    :vartype lenses: list[~azure.mgmt.portal.models.DashboardLens]
    :ivar metadata: The dashboard metadata.
    :vartype metadata: dict[str, JSON]
    """

    _attribute_map = {
        "tags": {"key": "tags", "type": "{str}"},
        "lenses": {"key": "properties.lenses", "type": "[DashboardLens]"},
        "metadata": {"key": "properties.metadata", "type": "{object}"},
    }

    def __init__(
        self,
        *,
        tags: Optional[Dict[str, str]] = None,
        lenses: Optional[List["_models.DashboardLens"]] = None,
        metadata: Optional[Dict[str, JSON]] = None,
        **kwargs
    ):
        """
        :keyword tags: Resource tags.
        :paramtype tags: dict[str, str]
        :keyword lenses: The dashboard lenses.
        :paramtype lenses: list[~azure.mgmt.portal.models.DashboardLens]
        :keyword metadata: The dashboard metadata.
        :paramtype metadata: dict[str, JSON]
        """
        super().__init__(**kwargs)
        self.tags = tags
        self.lenses = lenses
        self.metadata = metadata


class ResourceProviderOperation(_serialization.Model):
    """Supported operations of this resource provider.

    :ivar name: Operation name, in format of {provider}/{resource}/{operation}.
    :vartype name: str
    :ivar is_data_action: Indicates whether the operation applies to data-plane.
    :vartype is_data_action: str
    :ivar display: Display metadata associated with the operation.
    :vartype display: ~azure.mgmt.portal.models.ResourceProviderOperationDisplay
    """

    _attribute_map = {
        "name": {"key": "name", "type": "str"},
        "is_data_action": {"key": "isDataAction", "type": "str"},
        "display": {"key": "display", "type": "ResourceProviderOperationDisplay"},
    }

    def __init__(
        self,
        *,
        name: Optional[str] = None,
        is_data_action: Optional[str] = None,
        display: Optional["_models.ResourceProviderOperationDisplay"] = None,
        **kwargs
    ):
        """
        :keyword name: Operation name, in format of {provider}/{resource}/{operation}.
        :paramtype name: str
        :keyword is_data_action: Indicates whether the operation applies to data-plane.
        :paramtype is_data_action: str
        :keyword display: Display metadata associated with the operation.
        :paramtype display: ~azure.mgmt.portal.models.ResourceProviderOperationDisplay
        """
        super().__init__(**kwargs)
        self.name = name
        self.is_data_action = is_data_action
        self.display = display


class ResourceProviderOperationDisplay(_serialization.Model):
    """Display metadata associated with the operation.

    :ivar provider: Resource provider: Microsoft Custom Providers.
    :vartype provider: str
    :ivar resource: Resource on which the operation is performed.
    :vartype resource: str
    :ivar operation: Type of operation: get, read, delete, etc.
    :vartype operation: str
    :ivar description: Description of this operation.
    :vartype description: str
    """

    _attribute_map = {
        "provider": {"key": "provider", "type": "str"},
        "resource": {"key": "resource", "type": "str"},
        "operation": {"key": "operation", "type": "str"},
        "description": {"key": "description", "type": "str"},
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
        """
        :keyword provider: Resource provider: Microsoft Custom Providers.
        :paramtype provider: str
        :keyword resource: Resource on which the operation is performed.
        :paramtype resource: str
        :keyword operation: Type of operation: get, read, delete, etc.
        :paramtype operation: str
        :keyword description: Description of this operation.
        :paramtype description: str
        """
        super().__init__(**kwargs)
        self.provider = provider
        self.resource = resource
        self.operation = operation
        self.description = description


class ResourceProviderOperationList(_serialization.Model):
    """Results of the request to list operations.

    :ivar value: List of operations supported by this resource provider.
    :vartype value: list[~azure.mgmt.portal.models.ResourceProviderOperation]
    :ivar next_link: The URL to use for getting the next set of results.
    :vartype next_link: str
    """

    _attribute_map = {
        "value": {"key": "value", "type": "[ResourceProviderOperation]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(
        self,
        *,
        value: Optional[List["_models.ResourceProviderOperation"]] = None,
        next_link: Optional[str] = None,
        **kwargs
    ):
        """
        :keyword value: List of operations supported by this resource provider.
        :paramtype value: list[~azure.mgmt.portal.models.ResourceProviderOperation]
        :keyword next_link: The URL to use for getting the next set of results.
        :paramtype next_link: str
        """
        super().__init__(**kwargs)
        self.value = value
        self.next_link = next_link


class Violation(_serialization.Model):
    """Violation information.

    Variables are only populated by the server, and will be ignored when sending a request.

    :ivar id: Id of the item that violates tenant configuration.
    :vartype id: str
    :ivar user_id: Id of the user who owns violated item.
    :vartype user_id: str
    :ivar error_message: Error message.
    :vartype error_message: str
    """

    _validation = {
        "id": {"readonly": True},
        "user_id": {"readonly": True},
        "error_message": {"readonly": True},
    }

    _attribute_map = {
        "id": {"key": "id", "type": "str"},
        "user_id": {"key": "userId", "type": "str"},
        "error_message": {"key": "errorMessage", "type": "str"},
    }

    def __init__(self, **kwargs):
        """ """
        super().__init__(**kwargs)
        self.id = None
        self.user_id = None
        self.error_message = None


class ViolationsList(_serialization.Model):
    """List of list of items that violate tenant's configuration.

    :ivar value: The array of violations.
    :vartype value: list[~azure.mgmt.portal.models.Violation]
    :ivar next_link: The URL to use for getting the next set of results.
    :vartype next_link: str
    """

    _attribute_map = {
        "value": {"key": "value", "type": "[Violation]"},
        "next_link": {"key": "nextLink", "type": "str"},
    }

    def __init__(self, *, value: Optional[List["_models.Violation"]] = None, next_link: Optional[str] = None, **kwargs):
        """
        :keyword value: The array of violations.
        :paramtype value: list[~azure.mgmt.portal.models.Violation]
        :keyword next_link: The URL to use for getting the next set of results.
        :paramtype next_link: str
        """
        super().__init__(**kwargs)
        self.value = value
        self.next_link = next_link
