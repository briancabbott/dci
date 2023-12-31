// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.Resources.Models;

namespace Azure.ResourceManager.Hci.Models
{
    public partial class IPConfigurationProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(PrivateIPAddress))
            {
                writer.WritePropertyName("privateIPAddress"u8);
                writer.WriteStringValue(PrivateIPAddress);
            }
            if (Optional.IsDefined(Subnet))
            {
                writer.WritePropertyName("subnet"u8);
                JsonSerializer.Serialize(writer, Subnet);
            }
            writer.WriteEndObject();
        }

        internal static IPConfigurationProperties DeserializeIPConfigurationProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> gateway = default;
            Optional<string> prefixLength = default;
            Optional<string> privateIPAddress = default;
            Optional<WritableSubResource> subnet = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("gateway"u8))
                {
                    gateway = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("prefixLength"u8))
                {
                    prefixLength = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("privateIPAddress"u8))
                {
                    privateIPAddress = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("subnet"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    subnet = JsonSerializer.Deserialize<WritableSubResource>(property.Value.GetRawText());
                    continue;
                }
            }
            return new IPConfigurationProperties(gateway.Value, prefixLength.Value, privateIPAddress.Value, subnet);
        }
    }
}
