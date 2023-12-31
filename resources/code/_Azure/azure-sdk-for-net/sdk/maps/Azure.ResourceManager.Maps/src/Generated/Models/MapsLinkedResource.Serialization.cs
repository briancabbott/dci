// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Maps.Models
{
    public partial class MapsLinkedResource : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("uniqueName"u8);
            writer.WriteStringValue(UniqueName);
            writer.WritePropertyName("id"u8);
            writer.WriteStringValue(Id);
            writer.WriteEndObject();
        }

        internal static MapsLinkedResource DeserializeMapsLinkedResource(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            string uniqueName = default;
            string id = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("uniqueName"u8))
                {
                    uniqueName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("id"u8))
                {
                    id = property.Value.GetString();
                    continue;
                }
            }
            return new MapsLinkedResource(uniqueName, id);
        }
    }
}
