// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.HybridNetwork.Models
{
    public partial class DeploymentResourceIdReference : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("idType"u8);
            writer.WriteStringValue(IdType.ToString());
            writer.WriteEndObject();
        }

        internal static DeploymentResourceIdReference DeserializeDeploymentResourceIdReference(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            if (element.TryGetProperty("idType", out JsonElement discriminator))
            {
                switch (discriminator.GetString())
                {
                    case "Open": return OpenDeploymentResourceReference.DeserializeOpenDeploymentResourceReference(element);
                    case "Secret": return SecretDeploymentResourceReference.DeserializeSecretDeploymentResourceReference(element);
                }
            }
            return UnknownDeploymentResourceIdReference.DeserializeUnknownDeploymentResourceIdReference(element);
        }
    }
}
