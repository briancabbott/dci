// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ManagedNetworkFabric.Models
{
    public partial class ManagementNetworkConfigurationPatchableProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(InfrastructureVpnConfiguration))
            {
                writer.WritePropertyName("infrastructureVpnConfiguration"u8);
                writer.WriteObjectValue(InfrastructureVpnConfiguration);
            }
            if (Optional.IsDefined(WorkloadVpnConfiguration))
            {
                writer.WritePropertyName("workloadVpnConfiguration"u8);
                writer.WriteObjectValue(WorkloadVpnConfiguration);
            }
            writer.WriteEndObject();
        }

        internal static ManagementNetworkConfigurationPatchableProperties DeserializeManagementNetworkConfigurationPatchableProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<VpnConfigurationPatchableProperties> infrastructureVpnConfiguration = default;
            Optional<VpnConfigurationPatchableProperties> workloadVpnConfiguration = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("infrastructureVpnConfiguration"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    infrastructureVpnConfiguration = VpnConfigurationPatchableProperties.DeserializeVpnConfigurationPatchableProperties(property.Value);
                    continue;
                }
                if (property.NameEquals("workloadVpnConfiguration"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    workloadVpnConfiguration = VpnConfigurationPatchableProperties.DeserializeVpnConfigurationPatchableProperties(property.Value);
                    continue;
                }
            }
            return new ManagementNetworkConfigurationPatchableProperties(infrastructureVpnConfiguration.Value, workloadVpnConfiguration.Value);
        }
    }
}
