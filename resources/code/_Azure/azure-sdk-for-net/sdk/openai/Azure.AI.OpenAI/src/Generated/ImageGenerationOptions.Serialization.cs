// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure;
using Azure.Core;

namespace Azure.AI.OpenAI
{
    public partial class ImageGenerationOptions : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(DeploymentName))
            {
                writer.WritePropertyName("model"u8);
                writer.WriteStringValue(DeploymentName);
            }
            writer.WritePropertyName("prompt"u8);
            writer.WriteStringValue(Prompt);
            if (Optional.IsDefined(ImageCount))
            {
                writer.WritePropertyName("n"u8);
                writer.WriteNumberValue(ImageCount.Value);
            }
            if (Optional.IsDefined(Size))
            {
                writer.WritePropertyName("size"u8);
                writer.WriteStringValue(Size.Value.ToString());
            }
            if (Optional.IsDefined(ResponseFormat))
            {
                writer.WritePropertyName("response_format"u8);
                writer.WriteStringValue(ResponseFormat.Value.ToString());
            }
            if (Optional.IsDefined(Quality))
            {
                writer.WritePropertyName("quality"u8);
                writer.WriteStringValue(Quality.Value.ToString());
            }
            if (Optional.IsDefined(Style))
            {
                writer.WritePropertyName("style"u8);
                writer.WriteStringValue(Style.Value.ToString());
            }
            if (Optional.IsDefined(User))
            {
                writer.WritePropertyName("user"u8);
                writer.WriteStringValue(User);
            }
            writer.WriteEndObject();
        }

        internal static ImageGenerationOptions DeserializeImageGenerationOptions(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> model = default;
            string prompt = default;
            Optional<int> n = default;
            Optional<ImageSize> size = default;
            Optional<ImageGenerationResponseFormat> responseFormat = default;
            Optional<ImageGenerationQuality> quality = default;
            Optional<ImageGenerationStyle> style = default;
            Optional<string> user = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("model"u8))
                {
                    model = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("prompt"u8))
                {
                    prompt = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("n"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    n = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("size"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    size = new ImageSize(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("response_format"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    responseFormat = new ImageGenerationResponseFormat(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("quality"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    quality = new ImageGenerationQuality(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("style"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    style = new ImageGenerationStyle(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("user"u8))
                {
                    user = property.Value.GetString();
                    continue;
                }
            }
            return new ImageGenerationOptions(model.Value, prompt, Optional.ToNullable(n), Optional.ToNullable(size), Optional.ToNullable(responseFormat), Optional.ToNullable(quality), Optional.ToNullable(style), user.Value);
        }

        /// <summary> Deserializes the model from a raw response. </summary>
        /// <param name="response"> The response to deserialize the model from. </param>
        internal static ImageGenerationOptions FromResponse(Response response)
        {
            using var document = JsonDocument.Parse(response.Content);
            return DeserializeImageGenerationOptions(document.RootElement);
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
