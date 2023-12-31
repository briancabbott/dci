// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Net;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.SpringAppDiscovery.Models
{
    public partial class SpringBootServerProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Port))
            {
                writer.WritePropertyName("port"u8);
                writer.WriteNumberValue(Port.Value);
            }
            writer.WritePropertyName("server"u8);
            writer.WriteStringValue(Server);
            if (Optional.IsCollectionDefined(FqdnAndIPAddressList))
            {
                writer.WritePropertyName("fqdnAndIpAddressList"u8);
                writer.WriteStartArray();
                foreach (var item in FqdnAndIPAddressList)
                {
                    if (item == null)
                    {
                        writer.WriteNullValue();
                        continue;
                    }
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsDefined(MachineArmId))
            {
                writer.WritePropertyName("machineArmId"u8);
                writer.WriteStringValue(MachineArmId);
            }
            if (Optional.IsDefined(TotalApps))
            {
                writer.WritePropertyName("totalApps"u8);
                writer.WriteNumberValue(TotalApps.Value);
            }
            if (Optional.IsDefined(SpringBootApps))
            {
                writer.WritePropertyName("springBootApps"u8);
                writer.WriteNumberValue(SpringBootApps.Value);
            }
            if (Optional.IsCollectionDefined(Errors))
            {
                writer.WritePropertyName("errors"u8);
                writer.WriteStartArray();
                foreach (var item in Errors)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsDefined(ProvisioningState))
            {
                writer.WritePropertyName("provisioningState"u8);
                writer.WriteStringValue(ProvisioningState.Value.ToString());
            }
            writer.WriteEndObject();
        }

        internal static SpringBootServerProperties DeserializeSpringBootServerProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<int> port = default;
            string server = default;
            Optional<IList<IPAddress>> fqdnAndIPAddressList = default;
            Optional<ResourceIdentifier> machineArmId = default;
            Optional<int> totalApps = default;
            Optional<int> springBootApps = default;
            Optional<IList<SpringBootSiteError>> errors = default;
            Optional<SpringAppDiscoveryProvisioningState> provisioningState = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("port"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    port = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("server"u8))
                {
                    server = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("fqdnAndIpAddressList"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<IPAddress> array = new List<IPAddress>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        if (item.ValueKind == JsonValueKind.Null)
                        {
                            array.Add(null);
                        }
                        else
                        {
                            array.Add(IPAddress.Parse(item.GetString()));
                        }
                    }
                    fqdnAndIPAddressList = array;
                    continue;
                }
                if (property.NameEquals("machineArmId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    machineArmId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("totalApps"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    totalApps = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("springBootApps"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    springBootApps = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("errors"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<SpringBootSiteError> array = new List<SpringBootSiteError>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(SpringBootSiteError.DeserializeSpringBootSiteError(item));
                    }
                    errors = array;
                    continue;
                }
                if (property.NameEquals("provisioningState"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    provisioningState = new SpringAppDiscoveryProvisioningState(property.Value.GetString());
                    continue;
                }
            }
            return new SpringBootServerProperties(Optional.ToNullable(port), server, Optional.ToList(fqdnAndIPAddressList), machineArmId.Value, Optional.ToNullable(totalApps), Optional.ToNullable(springBootApps), Optional.ToList(errors), Optional.ToNullable(provisioningState));
        }
    }
}
