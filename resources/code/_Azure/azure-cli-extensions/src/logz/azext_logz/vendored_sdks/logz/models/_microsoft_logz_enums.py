# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum, EnumMeta
from six import with_metaclass

class _CaseInsensitiveEnumMeta(EnumMeta):
    def __getitem__(self, name):
        return super().__getitem__(name.upper())

    def __getattr__(cls, name):
        """Return the enum member matching `name`
        We use __getattr__ instead of descriptors or inserting into the enum
        class' __dict__ in order to support `name` and `value` being both
        properties for enum members (which live in the class' __dict__) and
        enum members themselves.
        """
        try:
            return cls._member_map_[name.upper()]
        except KeyError:
            raise AttributeError(name)


class CreatedByType(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """The type of identity that created the resource.
    """

    USER = "User"
    APPLICATION = "Application"
    MANAGED_IDENTITY = "ManagedIdentity"
    KEY = "Key"

class LiftrResourceCategories(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):

    UNKNOWN = "Unknown"
    MONITOR_LOGS = "MonitorLogs"

class ManagedIdentityTypes(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Identity type.
    """

    SYSTEM_ASSIGNED = "SystemAssigned"
    USER_ASSIGNED = "UserAssigned"

class MarketplaceSubscriptionStatus(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Flag specifying the Marketplace Subscription Status of the resource. If payment is not made in
    time, the resource will go in Suspended state.
    """

    ACTIVE = "Active"
    SUSPENDED = "Suspended"

class MonitoringStatus(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Flag specifying if the resource monitoring is enabled or disabled.
    """

    ENABLED = "Enabled"
    DISABLED = "Disabled"

class ProvisioningState(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Flag specifying if the resource provisioning state as tracked by ARM.
    """

    ACCEPTED = "Accepted"
    CREATING = "Creating"
    UPDATING = "Updating"
    DELETING = "Deleting"
    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELED = "Canceled"
    DELETED = "Deleted"
    NOT_SPECIFIED = "NotSpecified"

class SingleSignOnStates(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Various states of the SSO resource
    """

    INITIAL = "Initial"
    ENABLE = "Enable"
    DISABLE = "Disable"
    EXISTING = "Existing"

class TagAction(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Valid actions for a filtering tag. Exclusion takes priority over inclusion.
    """

    INCLUDE = "Include"
    EXCLUDE = "Exclude"

class UserRole(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """User roles on configured in Logz.io account.
    """

    NONE = "None"
    USER = "User"
    ADMIN = "Admin"

class VmHostUpdateStates(with_metaclass(_CaseInsensitiveEnumMeta, str, Enum)):
    """Various states of the updating vm extension on resource
    """

    INSTALL = "Install"
    DELETE = "Delete"
