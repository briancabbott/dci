// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.AppContainers.Models
{
    public partial class SecretVolumeItem : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(SecretRef))
            {
                writer.WritePropertyName("secretRef"u8);
                writer.WriteStringValue(SecretRef);
            }
            if (Optional.IsDefined(Path))
            {
                writer.WritePropertyName("path"u8);
                writer.WriteStringValue(Path);
            }
            writer.WriteEndObject();
        }

        internal static SecretVolumeItem DeserializeSecretVolumeItem(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> secretRef = default;
            Optional<string> path = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("secretRef"u8))
                {
                    secretRef = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("path"u8))
                {
                    path = property.Value.GetString();
                    continue;
                }
            }
            return new SecretVolumeItem(secretRef.Value, path.Value);
        }
    }
}
