// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    public partial class FeatureSetVersionBackfillContent : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Description))
            {
                writer.WritePropertyName("description"u8);
                writer.WriteStringValue(Description);
            }
            if (Optional.IsDefined(DisplayName))
            {
                writer.WritePropertyName("displayName"u8);
                writer.WriteStringValue(DisplayName);
            }
            if (Optional.IsDefined(FeatureWindow))
            {
                if (FeatureWindow != null)
                {
                    writer.WritePropertyName("featureWindow"u8);
                    writer.WriteObjectValue(FeatureWindow);
                }
                else
                {
                    writer.WriteNull("featureWindow");
                }
            }
            if (Optional.IsDefined(Resource))
            {
                writer.WritePropertyName("resource"u8);
                writer.WriteObjectValue(Resource);
            }
            if (Optional.IsCollectionDefined(SparkConfiguration))
            {
                writer.WritePropertyName("sparkConfiguration"u8);
                writer.WriteStartObject();
                foreach (var item in SparkConfiguration)
                {
                    writer.WritePropertyName(item.Key);
                    writer.WriteStringValue(item.Value);
                }
                writer.WriteEndObject();
            }
            if (Optional.IsCollectionDefined(Tags))
            {
                writer.WritePropertyName("tags"u8);
                writer.WriteStartObject();
                foreach (var item in Tags)
                {
                    writer.WritePropertyName(item.Key);
                    writer.WriteStringValue(item.Value);
                }
                writer.WriteEndObject();
            }
            writer.WriteEndObject();
        }
    }
}
