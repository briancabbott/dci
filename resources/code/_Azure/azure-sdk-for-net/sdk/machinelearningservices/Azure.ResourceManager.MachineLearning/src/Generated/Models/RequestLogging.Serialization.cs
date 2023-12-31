// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    internal partial class RequestLogging : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsCollectionDefined(CaptureHeaders))
            {
                if (CaptureHeaders != null)
                {
                    writer.WritePropertyName("captureHeaders"u8);
                    writer.WriteStartArray();
                    foreach (var item in CaptureHeaders)
                    {
                        writer.WriteStringValue(item);
                    }
                    writer.WriteEndArray();
                }
                else
                {
                    writer.WriteNull("captureHeaders");
                }
            }
            writer.WriteEndObject();
        }

        internal static RequestLogging DeserializeRequestLogging(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IList<string>> captureHeaders = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("captureHeaders"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        captureHeaders = null;
                        continue;
                    }
                    List<string> array = new List<string>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(item.GetString());
                    }
                    captureHeaders = array;
                    continue;
                }
            }
            return new RequestLogging(Optional.ToList(captureHeaders));
        }
    }
}
