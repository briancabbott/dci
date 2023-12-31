// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    internal partial class UnknownMonitoringFeatureFilterBase : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("filterType"u8);
            writer.WriteStringValue(FilterType.ToString());
            writer.WriteEndObject();
        }

        internal static UnknownMonitoringFeatureFilterBase DeserializeUnknownMonitoringFeatureFilterBase(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            MonitoringFeatureFilterType filterType = "Unknown";
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("filterType"u8))
                {
                    filterType = new MonitoringFeatureFilterType(property.Value.GetString());
                    continue;
                }
            }
            return new UnknownMonitoringFeatureFilterBase(filterType);
        }
    }
}
