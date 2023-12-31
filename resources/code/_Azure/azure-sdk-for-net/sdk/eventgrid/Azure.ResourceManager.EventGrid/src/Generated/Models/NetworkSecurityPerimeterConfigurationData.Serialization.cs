// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.Models;

namespace Azure.ResourceManager.EventGrid
{
    public partial class NetworkSecurityPerimeterConfigurationData : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("properties"u8);
            writer.WriteStartObject();
            if (Optional.IsDefined(ProvisioningState))
            {
                writer.WritePropertyName("provisioningState"u8);
                writer.WriteStringValue(ProvisioningState.Value.ToString());
            }
            if (Optional.IsCollectionDefined(ProvisioningIssues))
            {
                writer.WritePropertyName("provisioningIssues"u8);
                writer.WriteStartArray();
                foreach (var item in ProvisioningIssues)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsDefined(NetworkSecurityPerimeter))
            {
                writer.WritePropertyName("networkSecurityPerimeter"u8);
                writer.WriteObjectValue(NetworkSecurityPerimeter);
            }
            if (Optional.IsDefined(ResourceAssociation))
            {
                writer.WritePropertyName("resourceAssociation"u8);
                writer.WriteObjectValue(ResourceAssociation);
            }
            if (Optional.IsDefined(Profile))
            {
                writer.WritePropertyName("profile"u8);
                writer.WriteObjectValue(Profile);
            }
            writer.WriteEndObject();
            writer.WriteEndObject();
        }

        internal static NetworkSecurityPerimeterConfigurationData DeserializeNetworkSecurityPerimeterConfigurationData(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            ResourceIdentifier id = default;
            string name = default;
            ResourceType type = default;
            Optional<SystemData> systemData = default;
            Optional<NetworkSecurityPerimeterConfigProvisioningState> provisioningState = default;
            Optional<IList<NetworkSecurityPerimeterConfigurationIssues>> provisioningIssues = default;
            Optional<NetworkSecurityPerimeterInfo> networkSecurityPerimeter = default;
            Optional<ResourceAssociation> resourceAssociation = default;
            Optional<NetworkSecurityPerimeterConfigurationProfile> profile = default;
            foreach (var property in element.EnumerateObject())
            {
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
                        if (property0.NameEquals("provisioningState"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            provisioningState = new NetworkSecurityPerimeterConfigProvisioningState(property0.Value.GetString());
                            continue;
                        }
                        if (property0.NameEquals("provisioningIssues"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            List<NetworkSecurityPerimeterConfigurationIssues> array = new List<NetworkSecurityPerimeterConfigurationIssues>();
                            foreach (var item in property0.Value.EnumerateArray())
                            {
                                array.Add(NetworkSecurityPerimeterConfigurationIssues.DeserializeNetworkSecurityPerimeterConfigurationIssues(item));
                            }
                            provisioningIssues = array;
                            continue;
                        }
                        if (property0.NameEquals("networkSecurityPerimeter"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            networkSecurityPerimeter = NetworkSecurityPerimeterInfo.DeserializeNetworkSecurityPerimeterInfo(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("resourceAssociation"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            resourceAssociation = ResourceAssociation.DeserializeResourceAssociation(property0.Value);
                            continue;
                        }
                        if (property0.NameEquals("profile"u8))
                        {
                            if (property0.Value.ValueKind == JsonValueKind.Null)
                            {
                                continue;
                            }
                            profile = NetworkSecurityPerimeterConfigurationProfile.DeserializeNetworkSecurityPerimeterConfigurationProfile(property0.Value);
                            continue;
                        }
                    }
                    continue;
                }
            }
            return new NetworkSecurityPerimeterConfigurationData(id, name, type, systemData.Value, Optional.ToNullable(provisioningState), Optional.ToList(provisioningIssues), networkSecurityPerimeter.Value, resourceAssociation.Value, profile.Value);
        }
    }
}
