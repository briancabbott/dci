// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ManagedNetworkFabric.Models
{
    public partial class VlanMatchCondition : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsCollectionDefined(Vlans))
            {
                writer.WritePropertyName("vlans"u8);
                writer.WriteStartArray();
                foreach (var item in Vlans)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(InnerVlans))
            {
                writer.WritePropertyName("innerVlans"u8);
                writer.WriteStartArray();
                foreach (var item in InnerVlans)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(VlanGroupNames))
            {
                writer.WritePropertyName("vlanGroupNames"u8);
                writer.WriteStartArray();
                foreach (var item in VlanGroupNames)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WriteEndObject();
        }

        internal static VlanMatchCondition DeserializeVlanMatchCondition(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IList<string>> vlans = default;
            Optional<IList<string>> innerVlans = default;
            Optional<IList<string>> vlanGroupNames = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("vlans"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    vlans = array;
                    continue;
                }
                if (property.NameEquals("innerVlans"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    innerVlans = array;
                    continue;
                }
                if (property.NameEquals("vlanGroupNames"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    vlanGroupNames = array;
                    continue;
                }
            }
            return new VlanMatchCondition(Optional.ToList(vlans), Optional.ToList(innerVlans), Optional.ToList(vlanGroupNames));
        }
    }
}
