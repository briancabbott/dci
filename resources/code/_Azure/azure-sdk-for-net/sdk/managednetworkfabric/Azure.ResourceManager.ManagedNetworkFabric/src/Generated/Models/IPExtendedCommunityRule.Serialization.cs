// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ManagedNetworkFabric.Models
{
    public partial class IPExtendedCommunityRule : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("action"u8);
            writer.WriteStringValue(Action.ToString());
            writer.WritePropertyName("sequenceNumber"u8);
            writer.WriteNumberValue(SequenceNumber);
            writer.WritePropertyName("routeTargets"u8);
            writer.WriteStartArray();
            foreach (var item in RouteTargets)
            {
                writer.WriteStringValue(item);
            }
            writer.WriteEndArray();
            writer.WriteEndObject();
        }

        internal static IPExtendedCommunityRule DeserializeIPExtendedCommunityRule(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            CommunityActionType action = default;
            long sequenceNumber = default;
            IList<string> routeTargets = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("action"u8))
                {
                    action = new CommunityActionType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("sequenceNumber"u8))
                {
                    sequenceNumber = property.Value.GetInt64();
                    continue;
                }
                if (property.NameEquals("routeTargets"u8))
                {
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    routeTargets = array;
                    continue;
                }
            }
            return new IPExtendedCommunityRule(action, sequenceNumber, routeTargets);
        }
    }
}
