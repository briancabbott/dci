// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ManagedNetworkFabric.Models
{
    public partial class StatementConditionProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(RoutePolicyConditionType))
            {
                writer.WritePropertyName("type"u8);
                writer.WriteStringValue(RoutePolicyConditionType.Value.ToString());
            }
            if (Optional.IsDefined(IPPrefixId))
            {
                writer.WritePropertyName("ipPrefixId"u8);
                writer.WriteStringValue(IPPrefixId);
            }
            if (Optional.IsCollectionDefined(IPExtendedCommunityIds))
            {
                writer.WritePropertyName("ipExtendedCommunityIds"u8);
                writer.WriteStartArray();
                foreach (var item in IPExtendedCommunityIds)
                {
                    if (item == null)
                    {
                        writer.WriteNullValue();
                        continue;
                    }
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(IPCommunityIds))
            {
                writer.WritePropertyName("ipCommunityIds"u8);
                writer.WriteStartArray();
                foreach (var item in IPCommunityIds)
                {
                    if (item == null)
                    {
                        writer.WriteNullValue();
                        continue;
                    }
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WriteEndObject();
        }

        internal static StatementConditionProperties DeserializeStatementConditionProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<RoutePolicyConditionType> type = default;
            Optional<ResourceIdentifier> ipPrefixId = default;
            Optional<IList<ResourceIdentifier>> ipExtendedCommunityIds = default;
            Optional<IList<ResourceIdentifier>> ipCommunityIds = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("type"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    type = new RoutePolicyConditionType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("ipPrefixId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    ipPrefixId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("ipExtendedCommunityIds"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<ResourceIdentifier> array = new List<ResourceIdentifier>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        if (item.ValueKind == JsonValueKind.Null)
                        {
                            array.Add(null);
                        }
                        else
                        {
                            array.Add(new ResourceIdentifier(item.GetString()));
                        }
                    }
                    ipExtendedCommunityIds = array;
                    continue;
                }
                if (property.NameEquals("ipCommunityIds"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<ResourceIdentifier> array = new List<ResourceIdentifier>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        if (item.ValueKind == JsonValueKind.Null)
                        {
                            array.Add(null);
                        }
                        else
                        {
                            array.Add(new ResourceIdentifier(item.GetString()));
                        }
                    }
                    ipCommunityIds = array;
                    continue;
                }
            }
            return new StatementConditionProperties(Optional.ToList(ipCommunityIds), Optional.ToNullable(type), ipPrefixId.Value, Optional.ToList(ipExtendedCommunityIds));
        }
    }
}
