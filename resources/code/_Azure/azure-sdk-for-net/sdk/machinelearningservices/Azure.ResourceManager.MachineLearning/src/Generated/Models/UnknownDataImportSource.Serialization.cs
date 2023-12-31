// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    internal partial class UnknownDataImportSource : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Connection))
            {
                if (Connection != null)
                {
                    writer.WritePropertyName("connection"u8);
                    writer.WriteStringValue(Connection);
                }
                else
                {
                    writer.WriteNull("connection");
                }
            }
            writer.WritePropertyName("sourceType"u8);
            writer.WriteStringValue(SourceType.ToString());
            writer.WriteEndObject();
        }

        internal static UnknownDataImportSource DeserializeUnknownDataImportSource(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> connection = default;
            DataImportSourceType sourceType = "Unknown";
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("connection"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        connection = null;
                        continue;
                    }
                    connection = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("sourceType"u8))
                {
                    sourceType = new DataImportSourceType(property.Value.GetString());
                    continue;
                }
            }
            return new UnknownDataImportSource(connection.Value, sourceType);
        }
    }
}
