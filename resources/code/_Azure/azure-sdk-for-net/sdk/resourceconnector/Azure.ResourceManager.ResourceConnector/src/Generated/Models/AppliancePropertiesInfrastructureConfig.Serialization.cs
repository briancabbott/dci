// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ResourceConnector.Models
{
    internal partial class AppliancePropertiesInfrastructureConfig : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Provider))
            {
                writer.WritePropertyName("provider"u8);
                writer.WriteStringValue(Provider.Value.ToString());
            }
            writer.WriteEndObject();
        }

        internal static AppliancePropertiesInfrastructureConfig DeserializeAppliancePropertiesInfrastructureConfig(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<ApplianceProvider> provider = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("provider"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    provider = new ApplianceProvider(property.Value.GetString());
                    continue;
                }
            }
            return new AppliancePropertiesInfrastructureConfig(Optional.ToNullable(provider));
        }
    }
}
