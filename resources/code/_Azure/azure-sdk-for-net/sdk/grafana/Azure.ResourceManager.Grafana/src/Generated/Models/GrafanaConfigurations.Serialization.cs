// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Grafana.Models
{
    internal partial class GrafanaConfigurations : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Smtp))
            {
                writer.WritePropertyName("smtp"u8);
                writer.WriteObjectValue(Smtp);
            }
            writer.WriteEndObject();
        }

        internal static GrafanaConfigurations DeserializeGrafanaConfigurations(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<Smtp> smtp = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("smtp"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    smtp = Smtp.DeserializeSmtp(property.Value);
                    continue;
                }
            }
            return new GrafanaConfigurations(smtp.Value);
        }
    }
}
