# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class AddonProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the addon provisioning."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELLED = "Cancelled"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class AddonType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of private cloud addon."""

    SRM = "SRM"
    VR = "VR"
    HCX = "HCX"
    ARC = "Arc"


class AffinityStrength(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """VM-Host placement policy affinity strength (should/must)."""

    SHOULD = "Should"
    MUST = "Must"


class AffinityType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Placement policy affinity type."""

    AFFINITY = "Affinity"
    ANTI_AFFINITY = "AntiAffinity"


class AvailabilityStrategy(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The availability strategy for the private cloud."""

    SINGLE_ZONE = "SingleZone"
    DUAL_ZONE = "DualZone"


class AzureHybridBenefitType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Placement policy hosts opt-in Azure Hybrid Benefit type."""

    SQL_HOST = "SqlHost"
    NONE = "None"


class CloudLinkStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the cloud link."""

    ACTIVE = "Active"
    BUILDING = "Building"
    DELETING = "Deleting"
    FAILED = "Failed"
    DISCONNECTED = "Disconnected"


class ClusterProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the cluster provisioning."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELLED = "Cancelled"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class DatastoreProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the datastore provisioning."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELLED = "Cancelled"
    PENDING = "Pending"
    CREATING = "Creating"
    UPDATING = "Updating"
    DELETING = "Deleting"
    CANCELED = "Canceled"


class DatastoreStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The operational status of the datastore."""

    UNKNOWN = "Unknown"
    ACCESSIBLE = "Accessible"
    INACCESSIBLE = "Inaccessible"
    ATTACHED = "Attached"
    DETACHED = "Detached"
    LOST_COMMUNICATION = "LostCommunication"
    DEAD_OR_ERROR = "DeadOrError"


class DhcpTypeEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Type of DHCP: SERVER or RELAY."""

    SERVER = "SERVER"
    RELAY = "RELAY"


class DnsServiceLogLevelEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """DNS Service log level."""

    DEBUG = "DEBUG"
    INFO = "INFO"
    WARNING = "WARNING"
    ERROR = "ERROR"
    FATAL = "FATAL"


class DnsServiceStatusEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """DNS Service status."""

    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"


class EncryptionKeyStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of key provided."""

    CONNECTED = "Connected"
    ACCESS_DENIED = "AccessDenied"


class EncryptionState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Status of customer managed encryption key."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class EncryptionVersionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Property of the key if user provided or auto detected."""

    FIXED = "Fixed"
    AUTO_DETECTED = "AutoDetected"


class ExpressRouteAuthorizationProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the  ExpressRoute Circuit Authorization provisioning."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class GlobalReachConnectionProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the  ExpressRoute Circuit Authorization provisioning."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class GlobalReachConnectionStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The connection status of the global reach connection."""

    CONNECTED = "Connected"
    CONNECTING = "Connecting"
    DISCONNECTED = "Disconnected"


class HcxEnterpriseSiteStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The status of the HCX Enterprise Site."""

    AVAILABLE = "Available"
    CONSUMED = "Consumed"
    DEACTIVATED = "Deactivated"
    DELETED = "Deleted"


class InternetEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Connectivity to internet is enabled or disabled."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class MountOptionEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Mode that describes whether the LUN has to be mounted as a datastore or attached as a LUN."""

    MOUNT = "MOUNT"
    ATTACH = "ATTACH"


class NsxPublicIpQuotaRaisedEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Flag to indicate whether the private cloud has the quota for provisioned NSX Public IP count
    raised from 64 to 1024.
    """

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class OptionalParamEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Is this parameter required or optional."""

    OPTIONAL = "Optional"
    REQUIRED = "Required"


class PlacementPolicyProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class PlacementPolicyState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Whether the placement policy is enabled or disabled."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class PlacementPolicyType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """placement policy type."""

    VM_VM = "VmVm"
    VM_HOST = "VmHost"


class PortMirroringDirectionEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Direction of port mirroring profile."""

    INGRESS = "INGRESS"
    EGRESS = "EGRESS"
    BIDIRECTIONAL = "BIDIRECTIONAL"


class PortMirroringStatusEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Port Mirroring Status."""

    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"


class PrivateCloudProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELLED = "Cancelled"
    PENDING = "Pending"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class QuotaEnabled(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Host quota is active for current subscription."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class ResourceIdentityType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of identity used for the private cloud. The type 'SystemAssigned' refers to an
    implicitly created identity. The type 'None' will remove any identities from the Private Cloud.
    """

    SYSTEM_ASSIGNED = "SystemAssigned"
    NONE = "None"


class ScriptExecutionParameterType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of execution parameter."""

    VALUE = "Value"
    SECURE_VALUE = "SecureValue"
    CREDENTIAL = "Credential"


class ScriptExecutionProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The state of the script execution resource."""

    PENDING = "Pending"
    RUNNING = "Running"
    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    CANCELLING = "Cancelling"
    CANCELLED = "Cancelled"
    DELETING = "Deleting"
    CANCELED = "Canceled"


class ScriptOutputStreamType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """ScriptOutputStreamType."""

    INFORMATION = "Information"
    WARNING = "Warning"
    OUTPUT = "Output"
    ERROR = "Error"


class ScriptParameterTypes(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of parameter the script is expecting. psCredential is a PSCredentialObject."""

    STRING = "String"
    SECURE_STRING = "SecureString"
    CREDENTIAL = "Credential"
    INT = "Int"
    BOOL = "Bool"
    FLOAT = "Float"


class SegmentStatusEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Segment status."""

    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"


class SslEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Protect LDAP communication using SSL certificate (LDAPS)."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class TrialStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Trial status."""

    TRIAL_AVAILABLE = "TrialAvailable"
    TRIAL_USED = "TrialUsed"
    TRIAL_DISABLED = "TrialDisabled"


class VirtualMachineRestrictMovementState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Whether VM DRS-driven movement is restricted (enabled) or not (disabled)."""

    ENABLED = "Enabled"
    DISABLED = "Disabled"


class VisibilityParameterEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Should this parameter be visible to arm and passed in the parameters argument when executing."""

    VISIBLE = "Visible"
    HIDDEN = "Hidden"


class VMGroupStatusEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """VM Group status."""

    SUCCESS = "SUCCESS"
    FAILURE = "FAILURE"


class VMTypeEnum(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Virtual machine type."""

    REGULAR = "REGULAR"
    EDGE = "EDGE"
    SERVICE = "SERVICE"


class WorkloadNetworkDhcpProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkDnsServiceProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkDnsZoneProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkName(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """WorkloadNetworkName."""

    DEFAULT = "default"


class WorkloadNetworkPortMirroringProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkPublicIPProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkSegmentProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"


class WorkloadNetworkVMGroupProvisioningState(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The provisioning state."""

    SUCCEEDED = "Succeeded"
    FAILED = "Failed"
    BUILDING = "Building"
    DELETING = "Deleting"
    UPDATING = "Updating"
    CANCELED = "Canceled"
