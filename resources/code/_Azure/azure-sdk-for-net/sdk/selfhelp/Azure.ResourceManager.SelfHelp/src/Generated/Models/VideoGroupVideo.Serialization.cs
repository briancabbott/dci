// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.SelfHelp.Models
{
    public partial class VideoGroupVideo : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(Src))
            {
                writer.WritePropertyName("src"u8);
                writer.WriteStringValue(Src);
            }
            if (Optional.IsDefined(Title))
            {
                writer.WritePropertyName("title"u8);
                writer.WriteStringValue(Title);
            }
            writer.WriteEndObject();
        }

        internal static VideoGroupVideo DeserializeVideoGroupVideo(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> src = default;
            Optional<string> title = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("src"u8))
                {
                    src = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("title"u8))
                {
                    title = property.Value.GetString();
                    continue;
                }
            }
            return new VideoGroupVideo(src.Value, title.Value);
        }
    }
}
