// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Hci.Models
{
    public partial class VirtualMachineInstancePropertiesOSProfileLinuxConfiguration : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(DisablePasswordAuthentication))
            {
                writer.WritePropertyName("disablePasswordAuthentication"u8);
                writer.WriteBooleanValue(DisablePasswordAuthentication.Value);
            }
            if (Optional.IsDefined(Ssh))
            {
                writer.WritePropertyName("ssh"u8);
                writer.WriteObjectValue(Ssh);
            }
            if (Optional.IsDefined(ProvisionVmAgent))
            {
                writer.WritePropertyName("provisionVMAgent"u8);
                writer.WriteBooleanValue(ProvisionVmAgent.Value);
            }
            if (Optional.IsDefined(ProvisionVmConfigAgent))
            {
                writer.WritePropertyName("provisionVMConfigAgent"u8);
                writer.WriteBooleanValue(ProvisionVmConfigAgent.Value);
            }
            writer.WriteEndObject();
        }

        internal static VirtualMachineInstancePropertiesOSProfileLinuxConfiguration DeserializeVirtualMachineInstancePropertiesOSProfileLinuxConfiguration(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<bool> disablePasswordAuthentication = default;
            Optional<SshConfiguration> ssh = default;
            Optional<bool> provisionVmAgent = default;
            Optional<bool> provisionVmConfigAgent = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("disablePasswordAuthentication"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    disablePasswordAuthentication = property.Value.GetBoolean();
                    continue;
                }
                if (property.NameEquals("ssh"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    ssh = SshConfiguration.DeserializeSshConfiguration(property.Value);
                    continue;
                }
                if (property.NameEquals("provisionVMAgent"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    provisionVmAgent = property.Value.GetBoolean();
                    continue;
                }
                if (property.NameEquals("provisionVMConfigAgent"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    provisionVmConfigAgent = property.Value.GetBoolean();
                    continue;
                }
            }
            return new VirtualMachineInstancePropertiesOSProfileLinuxConfiguration(Optional.ToNullable(disablePasswordAuthentication), ssh.Value, Optional.ToNullable(provisionVmAgent), Optional.ToNullable(provisionVmConfigAgent));
        }
    }
}
