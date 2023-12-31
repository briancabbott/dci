# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class AlertNotifications(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Whether to send security alerts notifications to the security contact."""

    ON = "On"
    """Get notifications on new alerts"""
    OFF = "Off"
    """Don't get notifications on new alerts"""


class AlertsToAdmins(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Whether to send security alerts notifications to subscription admins."""

    ON = "On"
    """Send notification on new alerts to the subscription's admins"""
    OFF = "Off"
    """Don't send notification on new alerts to the subscription's admins"""


class AutoProvision(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Describes what kind of security agent provisioning action to take."""

    ON = "On"
    """Install missing security agent on VMs automatically"""
    OFF = "Off"
    """Do not install security agent on the VMs automatically"""


class InformationProtectionPolicyName(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """InformationProtectionPolicyName."""

    EFFECTIVE = "effective"
    CUSTOM = "custom"


class Rank(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The rank of the sensitivity label."""

    NONE = "None"
    LOW = "Low"
    MEDIUM = "Medium"
    HIGH = "High"
    CRITICAL = "Critical"
