// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources.Models;

namespace Azure.ResourceManager.ConnectedVMwarevSphere
{
    public partial class VMwareVmInstanceData : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(ExtendedLocation))
            {
                writer.WritePropertyName("extendedLocation"u8);
                JsonSerializer.Serialize(writer, ExtendedLocation);
            }
            writer.WritePropertyName("properties"u8);
            writer.WriteStartObject();
            if (Optional.IsDefined(PlacementProfile))
            {
                writer.WritePropertyName("placementProfile"u8);
                writer.WriteObjectValue(PlacementProfile);
            }
            if (Optional.IsDefined(OSProfile))
            {
                writer.WritePropertyName("osProfile"u8);
                writer.WriteObjectValue(OSProfile);
            }
            if (Optional.IsDefined(HardwareProfile))
            {
                writer.WritePropertyName("hardwareProfile"u8);
                writer.WriteObjectValue(HardwareProfile);
            }
            if (Optional.IsDefined(NetworkProfile))
            {
                writer.WritePropertyName("networkProfile"u8);
                writer.WriteObjectValue(NetworkProfile);
            }
            if (Optional.IsDefined(StorageProfile))
            {
                writer.WritePropertyName("storageProfile"u8);
                writer.WriteObjectValue(StorageProfile);
            }
            if (Optional.IsDefined(SecurityProfile))
            {
                writer.WritePropertyName("securityProfile"u8);
                writer.WriteObjectValue(SecurityProfile);
            }
            if (Optional.IsDefined(InfrastructureProfile))
            {
                writer.WritePropertyName("infrastructureProfile"u8);
                writer.WriteObjectValue(InfrastructureProfile);
            }
            writer.WriteEndObject();
            writer.WriteEndObject();
        }

        internal static VMwareVmInstanceData DeserializeVMwareVmInstanceData(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<ExtendedLocation> extendedLocation = default;
            ResourceIdentifier id = default;
            string name = default;
            ResourceType type = default;
            Optional<SystemData> systemData = default;
            Optional<PlacementProfile> placementProfile = default;
            Optional<OSProfileForVmInstance> osProfile = default;
            Optional<VmInstanceHardwareProfile> hardwareProfile = default;
            Optional<VMwareNetworkProfile> networkProfile = default;
            Optional<VMwareStorageProfile> storageProfile = default;
            Optional<SecurityProfile> securityProfile = default;
            Optional<VCenterInfrastructureProfile> infrastructureProfile = default;
            Optional<string> powerState = default;
            Optional<IReadOnlyList<VMwareResourceStatus>> statuses = default;
            Optional<VMwareResourceProvisioningState> provisioningState = default;
            Optional<string> resourceUid = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("extendedLocation"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    extendedLocation = JsonSerializer.Deserialize<ExtendedLocation>(property.Value.GetRawText());
                    continue;
                }
                if (property.NameEquals("id"u8))
                {
                    id = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("name"u8))
                {
                    name = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("type"u8))
                {
                    type = new ResourceType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("systemData"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    systemData = JsonSerializer.Deserialize<SystemData>(property.Value.GetRawText());
                    continue;
                }
                if (property.NameEquals("properties"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        property.ThrowNonNullablePropertyIsNull();
                        continue;
                    }
                    foreach (var property0 in property.Value.EnumerateObject())
                    {
                        if (property0.NameEquals("placementProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            placementProfile = PlacementProfile.DeserializePlacementProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("osProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            osProfile = OSProfileForVmInstance.DeserializeOSProfileForVmInstance(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("hardwareProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            hardwareProfile = VmInstanceHardwareProfile.DeserializeVmInstanceHardwareProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("networkProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            networkProfile = VMwareNetworkProfile.DeserializeVMwareNetworkProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("storageProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            storageProfile = VMwareStorageProfile.DeserializeVMwareStorageProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("securityProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            securityProfile = SecurityProfile.DeserializeSecurityProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("infrastructureProfile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            infrastructureProfile = VCenterInfrastructureProfile.DeserializeVCenterInfrastructureProfile(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("powerState"u8))
                        {
                            powerState = property0.Value.GetString();
                            continue;
                        }
                        if (property0.NameEquals("statuses"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            List<VMwareResourceStatus> array = new List<VMwareResourceStatus>();
                            foreach (var item in property0.Value.EnumerateArray())
                            {
                                array.Add(VMwareResourceStatus.DeserializeVMwareResourceStatus(item));
                            }
                            statuses = array;
                            continue;
                        }
                        if (property0.NameEquals("provisioningState"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            provisioningState = new VMwareResourceProvisioningState(property0.Value.GetString());
                            continue;
                        }
                        if (property0.NameEquals("resourceUid"u8))
                        {
                            resourceUid = property0.Value.GetString();
                            continue;
                        }
                    }
                    continue;
                }
            }
            return new VMwareVmInstanceData(id, name, type, systemData.Value, extendedLocation, placementProfile.Value, osProfile.Value, hardwareProfile.Value, networkProfile.Value, storageProfile.Value, securityProfile.Value, infrastructureProfile.Value, powerState.Value, Optional.ToList(statuses), Optional.ToNullable(provisioningState), resourceUid.Value);
        }
    }
}
