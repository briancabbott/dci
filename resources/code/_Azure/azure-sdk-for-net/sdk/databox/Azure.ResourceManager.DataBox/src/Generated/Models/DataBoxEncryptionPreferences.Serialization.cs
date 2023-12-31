// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.DataBox.Models
{
    public partial class DataBoxEncryptionPreferences : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(DoubleEncryption))
            {
                writer.WritePropertyName("doubleEncryption"u8);
                writer.WriteStringValue(DoubleEncryption.Value.ToSerialString());
            }
            if (Optional.IsDefined(HardwareEncryption))
            {
                writer.WritePropertyName("hardwareEncryption"u8);
                writer.WriteStringValue(HardwareEncryption.Value.ToSerialString());
            }
            writer.WriteEndObject();
        }

        internal static DataBoxEncryptionPreferences DeserializeDataBoxEncryptionPreferences(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<DataBoxDoubleEncryption> doubleEncryption = default;
            Optional<HardwareEncryption> hardwareEncryption = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("doubleEncryption"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    doubleEncryption = property.Value.GetString().ToDataBoxDoubleEncryption();
                    continue;
                }
                if (property.NameEquals("hardwareEncryption"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    hardwareEncryption = property.Value.GetString().ToHardwareEncryption();
                    continue;
                }
            }
            return new DataBoxEncryptionPreferences(Optional.ToNullable(doubleEncryption), Optional.ToNullable(hardwareEncryption));
        }
    }
}
