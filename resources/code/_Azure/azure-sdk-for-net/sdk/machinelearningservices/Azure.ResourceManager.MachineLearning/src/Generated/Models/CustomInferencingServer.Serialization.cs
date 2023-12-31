// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    public partial class CustomInferencingServer : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(InferenceConfiguration))
            {
                if (InferenceConfiguration != null)
                {
                    writer.WritePropertyName("inferenceConfiguration"u8);
                    writer.WriteObjectValue(InferenceConfiguration);
                }
                else
                {
                    writer.WriteNull("inferenceConfiguration");
                }
            }
            writer.WritePropertyName("serverType"u8);
            writer.WriteStringValue(ServerType.ToString());
            writer.WriteEndObject();
        }

        internal static CustomInferencingServer DeserializeCustomInferencingServer(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<OnlineInferenceConfiguration> inferenceConfiguration = default;
            InferencingServerType serverType = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("inferenceConfiguration"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        inferenceConfiguration = null;
                        continue;
                    }
                    inferenceConfiguration = OnlineInferenceConfiguration.DeserializeOnlineInferenceConfiguration(property.Value);
                    continue;
                }
                if (property.NameEquals("serverType"u8))
                {
                    serverType = new InferencingServerType(property.Value.GetString());
                    continue;
                }
            }
            return new CustomInferencingServer(serverType, inferenceConfiguration.Value);
        }
    }
}
