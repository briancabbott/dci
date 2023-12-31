// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.NetworkCloud.Models
{
    public partial class ControlPlaneNodeConfiguration : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(AdministratorConfiguration))
            {
                writer.WritePropertyName("administratorConfiguration"u8);
                writer.WriteObjectValue(AdministratorConfiguration);
            }
            if (Optional.IsCollectionDefined(AvailabilityZones))
            {
                writer.WritePropertyName("availabilityZones"u8);
                writer.WriteStartArray();
                foreach (var item in AvailabilityZones)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WritePropertyName("count"u8);
            writer.WriteNumberValue(Count);
            writer.WritePropertyName("vmSkuName"u8);
            writer.WriteStringValue(VmSkuName);
            writer.WriteEndObject();
        }

        internal static ControlPlaneNodeConfiguration DeserializeControlPlaneNodeConfiguration(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<AdministratorConfiguration> administratorConfiguration = default;
            Optional<IList<string>> availabilityZones = default;
            long count = default;
            string vmSkuName = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("administratorConfiguration"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    administratorConfiguration = AdministratorConfiguration.DeserializeAdministratorConfiguration(property.Value);
                    continue;
                }
                if (property.NameEquals("availabilityZones"u8))
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
                    availabilityZones = array;
                    continue;
                }
                if (property.NameEquals("count"u8))
                {
                    count = property.Value.GetInt64();
                    continue;
                }
                if (property.NameEquals("vmSkuName"u8))
                {
                    vmSkuName = property.Value.GetString();
                    continue;
                }
            }
            return new ControlPlaneNodeConfiguration(administratorConfiguration.Value, Optional.ToList(availabilityZones), count, vmSkuName);
        }
    }
}
