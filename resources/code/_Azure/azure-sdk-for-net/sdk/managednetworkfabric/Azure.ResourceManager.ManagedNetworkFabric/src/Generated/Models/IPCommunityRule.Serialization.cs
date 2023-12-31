// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ManagedNetworkFabric.Models
{
    public partial class IPCommunityRule : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("action"u8);
            writer.WriteStringValue(Action.ToString());
            writer.WritePropertyName("sequenceNumber"u8);
            writer.WriteNumberValue(SequenceNumber);
            if (Optional.IsCollectionDefined(WellKnownCommunities))
            {
                writer.WritePropertyName("wellKnownCommunities"u8);
                writer.WriteStartArray();
                foreach (var item in WellKnownCommunities)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            writer.WritePropertyName("communityMembers"u8);
            writer.WriteStartArray();
            foreach (var item in CommunityMembers)
            {
                writer.WriteStringValue(item);
            }
            writer.WriteEndArray();
            writer.WriteEndObject();
        }

        internal static IPCommunityRule DeserializeIPCommunityRule(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            CommunityActionType action = default;
            long sequenceNumber = default;
            Optional<IList<WellKnownCommunity>> wellKnownCommunities = default;
            IList<string> communityMembers = default;
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
                if (property.NameEquals("wellKnownCommunities"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<WellKnownCommunity> array = new List<WellKnownCommunity>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(new WellKnownCommunity(item.GetString()));
                    }
                    wellKnownCommunities = array;
                    continue;
                }
                if (property.NameEquals("communityMembers"u8))
                {
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    communityMembers = array;
                    continue;
                }
            }
            return new IPCommunityRule(action, sequenceNumber, Optional.ToList(wellKnownCommunities), communityMembers);
        }
    }
}
