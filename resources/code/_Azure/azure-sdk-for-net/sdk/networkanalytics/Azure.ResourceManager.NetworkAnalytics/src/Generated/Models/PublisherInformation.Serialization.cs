// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.NetworkAnalytics.Models
{
    public partial class PublisherInformation : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("publisherName"u8);
            writer.WriteStringValue(PublisherName);
            writer.WritePropertyName("dataProducts"u8);
            writer.WriteStartArray();
            foreach (var item in DataProducts)
            {
                writer.WriteObjectValue(item);
            }
            writer.WriteEndArray();
            writer.WriteEndObject();
        }

        internal static PublisherInformation DeserializePublisherInformation(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            string publisherName = default;
            IList<DataProductInformation> dataProducts = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("publisherName"u8))
                {
                    publisherName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("dataProducts"u8))
                {
                    List<DataProductInformation> array = new List<DataProductInformation>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(DataProductInformation.DeserializeDataProductInformation(item));
                    }
                    dataProducts = array;
                    continue;
                }
            }
            return new PublisherInformation(publisherName, dataProducts);
        }
    }
}
