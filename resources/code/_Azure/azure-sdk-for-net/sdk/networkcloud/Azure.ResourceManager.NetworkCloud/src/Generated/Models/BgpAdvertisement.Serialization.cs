// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.NetworkCloud.Models
{
    public partial class BgpAdvertisement : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(AdvertiseToFabric))
            {
                writer.WritePropertyName("advertiseToFabric"u8);
                writer.WriteStringValue(AdvertiseToFabric.Value.ToString());
            }
            if (Optional.IsCollectionDefined(Communities))
            {
                writer.WritePropertyName("communities"u8);
                writer.WriteStartArray();
                foreach (var item in Communities)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WritePropertyName("ipAddressPools"u8);
            writer.WriteStartArray();
            foreach (var item in IPAddressPools)
            {
                writer.WriteStringValue(item);
            }
            writer.WriteEndArray();
            if (Optional.IsCollectionDefined(Peers))
            {
                writer.WritePropertyName("peers"u8);
                writer.WriteStartArray();
                foreach (var item in Peers)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WriteEndObject();
        }

        internal static BgpAdvertisement DeserializeBgpAdvertisement(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<AdvertiseToFabric> advertiseToFabric = default;
            Optional<IList<string>> communities = default;
            IList<string> ipAddressPools = default;
            Optional<IList<string>> peers = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("advertiseToFabric"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    advertiseToFabric = new AdvertiseToFabric(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("communities"u8))
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
                    communities = array;
                    continue;
                }
                if (property.NameEquals("ipAddressPools"u8))
                {
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    ipAddressPools = array;
                    continue;
                }
                if (property.NameEquals("peers"u8))
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
                    peers = array;
                    continue;
                }
            }
            return new BgpAdvertisement(Optional.ToNullable(advertiseToFabric), Optional.ToList(communities), ipAddressPools, Optional.ToList(peers));
        }
    }
}
