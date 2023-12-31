// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Text.Json;
using Azure;
using Azure.Core;

namespace Azure.AI.DocumentIntelligence
{
    public partial class AzureBlobFileListContentSource : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("containerUrl"u8);
            writer.WriteStringValue(ContainerUrl.AbsoluteUri);
            writer.WritePropertyName("fileList"u8);
            writer.WriteStringValue(FileList);
            writer.WriteEndObject();
        }

        internal static AzureBlobFileListContentSource DeserializeAzureBlobFileListContentSource(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Uri containerUrl = default;
            string fileList = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("containerUrl"u8))
                {
                    containerUrl = new Uri(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("fileList"u8))
                {
                    fileList = property.Value.GetString();
                    continue;
                }
            }
            return new AzureBlobFileListContentSource(containerUrl, fileList);
        }

        /// <summary> Deserializes the model from a raw response. </summary>
        /// <param name="response"> The response to deserialize the model from. </param>
        internal static AzureBlobFileListContentSource FromResponse(Response response)
        {
            using var document = JsonDocument.Parse(response.Content);
            return DeserializeAzureBlobFileListContentSource(document.RootElement);
        }

        /// <summary> Convert into a Utf8JsonRequestContent. </summary>
        internal virtual RequestContent ToRequestContent()
        {
            var content = new Utf8JsonRequestContent();
            content.JsonWriter.WriteObjectValue(this);
            return content;
        }
    }
}
