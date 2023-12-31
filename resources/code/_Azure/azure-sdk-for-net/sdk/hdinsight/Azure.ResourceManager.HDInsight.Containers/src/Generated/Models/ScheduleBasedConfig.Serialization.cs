// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.HDInsight.Containers.Models
{
    public partial class ScheduleBasedConfig : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("timeZone"u8);
            writer.WriteStringValue(TimeZone);
            writer.WritePropertyName("defaultCount"u8);
            writer.WriteNumberValue(DefaultCount);
            writer.WritePropertyName("schedules"u8);
            writer.WriteStartArray();
            foreach (var item in Schedules)
            {
                writer.WriteObjectValue(item);
            }
            writer.WriteEndArray();
            writer.WriteEndObject();
        }

        internal static ScheduleBasedConfig DeserializeScheduleBasedConfig(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            string timeZone = default;
            int defaultCount = default;
            IList<AutoscaleSchedule> schedules = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("timeZone"u8))
                {
                    timeZone = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("defaultCount"u8))
                {
                    defaultCount = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("schedules"u8))
                {
                    List<AutoscaleSchedule> array = new List<AutoscaleSchedule>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(AutoscaleSchedule.DeserializeAutoscaleSchedule(item));
                    }
                    schedules = array;
                    continue;
                }
            }
            return new ScheduleBasedConfig(timeZone, defaultCount, schedules);
        }
    }
}
