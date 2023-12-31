// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.NetworkAnalytics.Models
{
    public partial class DataProductNetworkAcls : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("virtualNetworkRule"u8);
            writer.WriteStartArray();
            foreach (var item in VirtualNetworkRule)
            {
                writer.WriteObjectValue(item);
            }
            writer.WriteEndArray();
            writer.WritePropertyName("ipRules"u8);
            writer.WriteStartArray();
            foreach (var item in IPRules)
            {
                writer.WriteObjectValue(item);
            }
            writer.WriteEndArray();
            writer.WritePropertyName("allowedQueryIpRangeList"u8);
            writer.WriteStartArray();
            foreach (var item in AllowedQueryIPRangeList)
            {
                writer.WriteStringValue(item);
            }
            writer.WriteEndArray();
            writer.WritePropertyName("defaultAction"u8);
            writer.WriteStringValue(DefaultAction.ToString());
            writer.WriteEndObject();
        }

        internal static DataProductNetworkAcls DeserializeDataProductNetworkAcls(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            IList<NetworkAnalyticsVirtualNetworkRule> virtualNetworkRule = default;
            IList<NetworkAnalyticsIPRules> ipRules = default;
            IList<string> allowedQueryIPRangeList = default;
            NetworkAclDefaultAction defaultAction = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("virtualNetworkRule"u8))
                {
                    List<NetworkAnalyticsVirtualNetworkRule> array = new List<NetworkAnalyticsVirtualNetworkRule>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(NetworkAnalyticsVirtualNetworkRule.DeserializeNetworkAnalyticsVirtualNetworkRule(item));
                    }
                    virtualNetworkRule = array;
                    continue;
                }
                if (property.NameEquals("ipRules"u8))
                {
                    List<NetworkAnalyticsIPRules> array = new List<NetworkAnalyticsIPRules>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(NetworkAnalyticsIPRules.DeserializeNetworkAnalyticsIPRules(item));
                    }
                    ipRules = array;
                    continue;
                }
                if (property.NameEquals("allowedQueryIpRangeList"u8))
                {
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    allowedQueryIPRangeList = array;
                    continue;
                }
                if (property.NameEquals("defaultAction"u8))
                {
                    defaultAction = new NetworkAclDefaultAction(property.Value.GetString());
                    continue;
                }
            }
            return new DataProductNetworkAcls(virtualNetworkRule, ipRules, allowedQueryIPRangeList, defaultAction);
        }
    }
}
