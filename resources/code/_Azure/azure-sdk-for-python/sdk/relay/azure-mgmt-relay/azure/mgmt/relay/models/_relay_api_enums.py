# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class AccessRights(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """AccessRights."""

    MANAGE = "Manage"
    SEND = "Send"
    LISTEN = "Listen"


class CreatedByType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of identity that created the resource."""

    USER = "User"
    APPLICATION = "Application"
    MANAGED_IDENTITY = "ManagedIdentity"
    KEY = "Key"


class DefaultAction(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Default Action for Network Rule Set."""

    ALLOW = "Allow"
    DENY = "Deny"


class EndPointProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Provisioning state of the Private Endpoint Connection."""

    CREATING = "Creating"
    UPDATING = "Updating"
    DELETING = "Deleting"
    SUCCEEDED = "Succeeded"
    CANCELED = "Canceled"
    FAILED = "Failed"


class KeyType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The access key to regenerate."""

    PRIMARY_KEY = "PrimaryKey"
    SECONDARY_KEY = "SecondaryKey"


class NetworkRuleIPAction(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The IP Filter Action."""

    ALLOW = "Allow"


class PrivateLinkConnectionStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Status of the connection."""

    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"
    DISCONNECTED = "Disconnected"


class PublicNetworkAccess(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """This determines if traffic is allowed over public network. By default it is enabled."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"
    SECURED_BY_PERIMETER = "SecuredByPerimeter"


class Relaytype(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """WCF relay type."""

    NET_TCP = "NetTcp"
    HTTP = "Http"


class SkuName(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Name of this SKU."""

    STANDARD = "Standard"


class SkuTier(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The tier of this SKU."""

    STANDARD = "Standard"


class UnavailableReason(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Specifies the reason for the unavailability of the service."""

    NONE = "None"
    INVALID_NAME = "InvalidName"
    SUBSCRIPTION_IS_DISABLED = "SubscriptionIsDisabled"
    NAME_IN_USE = "NameInUse"
    NAME_IN_LOCKDOWN = "NameInLockdown"
    TOO_MANY_NAMESPACE_IN_CURRENT_SUBSCRIPTION = "TooManyNamespaceInCurrentSubscription"
