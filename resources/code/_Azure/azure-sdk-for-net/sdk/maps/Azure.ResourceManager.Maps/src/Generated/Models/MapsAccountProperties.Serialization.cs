// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Maps.Models
{
    public partial class MapsAccountProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(DisableLocalAuth))
            {
                writer.WritePropertyName("disableLocalAuth"u8);
                writer.WriteBooleanValue(DisableLocalAuth.Value);
            }
            if (Optional.IsCollectionDefined(LinkedResources))
            {
                writer.WritePropertyName("linkedResources"u8);
                writer.WriteStartArray();
                foreach (var item in LinkedResources)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsDefined(Cors))
            {
                writer.WritePropertyName("cors"u8);
                writer.WriteObjectValue(Cors);
            }
            writer.WriteEndObject();
        }

        internal static MapsAccountProperties DeserializeMapsAccountProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<Guid> uniqueId = default;
            Optional<bool> disableLocalAuth = default;
            Optional<string> provisioningState = default;
            Optional<IList<MapsLinkedResource>> linkedResources = default;
            Optional<CorsRules> cors = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("uniqueId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    uniqueId = property.Value.GetGuid();
                    continue;
                }
                if (property.NameEquals("disableLocalAuth"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    disableLocalAuth = property.Value.GetBoolean();
                    continue;
                }
                if (property.NameEquals("provisioningState"u8))
                {
                    provisioningState = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("linkedResources"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<MapsLinkedResource> array = new List<MapsLinkedResource>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(MapsLinkedResource.DeserializeMapsLinkedResource(item));
                    }
                    linkedResources = array;
                    continue;
                }
                if (property.NameEquals("cors"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    cors = CorsRules.DeserializeCorsRules(property.Value);
                    continue;
                }
            }
            return new MapsAccountProperties(Optional.ToNullable(uniqueId), Optional.ToNullable(disableLocalAuth), provisioningState.Value, Optional.ToList(linkedResources), cors.Value);
        }
    }
}
