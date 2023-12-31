# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class Category(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Identity system of the device."""

    AZURE_AD = "AzureAD"
    ADFS = "ADFS"


class CompatibilityIssue(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Compatibility issue."""

    HIGHER_DEVICE_VERSION_REQUIRED = "HigherDeviceVersionRequired"
    LOWER_DEVICE_VERSION_REQUIRED = "LowerDeviceVersionRequired"
    CAPACITY_BILLING_MODEL_REQUIRED = "CapacityBillingModelRequired"
    PAY_AS_YOU_GO_BILLING_MODEL_REQUIRED = "PayAsYouGoBillingModelRequired"
    DEVELOPMENT_BILLING_MODEL_REQUIRED = "DevelopmentBillingModelRequired"
    AZURE_AD_IDENTITY_SYSTEM_REQUIRED = "AzureADIdentitySystemRequired"
    ADFS_IDENTITY_SYSTEM_REQUIRED = "ADFSIdentitySystemRequired"
    CONNECTION_TO_INTERNET_REQUIRED = "ConnectionToInternetRequired"
    CONNECTION_TO_AZURE_REQUIRED = "ConnectionToAzureRequired"
    DISCONNECTED_ENVIRONMENT_REQUIRED = "DisconnectedEnvironmentRequired"


class ComputeRole(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Compute role type (IaaS or PaaS)."""

    NONE = "None"
    IAA_S = "IaaS"
    PAA_S = "PaaS"


class Location(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Location of the resource."""

    GLOBAL = "global"


class OperatingSystem(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Operating system type (Windows or Linux)."""

    NONE = "None"
    WINDOWS = "Windows"
    LINUX = "Linux"


class ProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state of the resource."""

    CREATING = "Creating"
    FAILED = "Failed"
    SUCCEEDED = "Succeeded"
    CANCELED = "Canceled"
