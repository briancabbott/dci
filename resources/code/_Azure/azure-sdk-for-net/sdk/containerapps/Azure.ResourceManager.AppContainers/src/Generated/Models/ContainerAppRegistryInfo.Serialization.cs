// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.AppContainers.Models
{
    public partial class ContainerAppRegistryInfo : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(RegistryServer))
            {
                writer.WritePropertyName("registryUrl"u8);
                writer.WriteStringValue(RegistryServer);
            }
            if (Optional.IsDefined(RegistryUserName))
            {
                writer.WritePropertyName("registryUserName"u8);
                writer.WriteStringValue(RegistryUserName);
            }
            if (Optional.IsDefined(RegistryPassword))
            {
                writer.WritePropertyName("registryPassword"u8);
                writer.WriteStringValue(RegistryPassword);
            }
            writer.WriteEndObject();
        }

        internal static ContainerAppRegistryInfo DeserializeContainerAppRegistryInfo(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> registryUrl = default;
            Optional<string> registryUserName = default;
            Optional<string> registryPassword = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("registryUrl"u8))
                {
                    registryUrl = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("registryUserName"u8))
                {
                    registryUserName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("registryPassword"u8))
                {
                    registryPassword = property.Value.GetString();
                    continue;
                }
            }
            return new ContainerAppRegistryInfo(registryUrl.Value, registryUserName.Value, registryPassword.Value);
        }
    }
}
