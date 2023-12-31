// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.HybridNetwork.Models
{
    public partial class HelmUpgradeConfig : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Atomic))
            {
                writer.WritePropertyName("atomic"u8);
                writer.WriteStringValue(Atomic);
            }
            if (Optional.IsDefined(Wait))
            {
                writer.WritePropertyName("wait"u8);
                writer.WriteStringValue(Wait);
            }
            if (Optional.IsDefined(Timeout))
            {
                writer.WritePropertyName("timeout"u8);
                writer.WriteStringValue(Timeout);
            }
            writer.WriteEndObject();
        }

        internal static HelmUpgradeConfig DeserializeHelmUpgradeConfig(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> atomic = default;
            Optional<string> wait = default;
            Optional<string> timeout = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("atomic"u8))
                {
                    atomic = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("wait"u8))
                {
                    wait = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("timeout"u8))
                {
                    timeout = property.Value.GetString();
                    continue;
                }
            }
            return new HelmUpgradeConfig(atomic.Value, wait.Value, timeout.Value);
        }
    }
}
