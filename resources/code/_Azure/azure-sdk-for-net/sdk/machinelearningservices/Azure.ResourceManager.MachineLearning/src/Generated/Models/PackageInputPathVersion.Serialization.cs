// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    public partial class PackageInputPathVersion : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(ResourceName))
            {
                if (ResourceName != null)
                {
                    writer.WritePropertyName("resourceName"u8);
                    writer.WriteStringValue(ResourceName);
                }
                else
                {
                    writer.WriteNull("resourceName");
                }
            }
            if (Optional.IsDefined(ResourceVersion))
            {
                if (ResourceVersion != null)
                {
                    writer.WritePropertyName("resourceVersion"u8);
                    writer.WriteStringValue(ResourceVersion);
                }
                else
                {
                    writer.WriteNull("resourceVersion");
                }
            }
            writer.WritePropertyName("inputPathType"u8);
            writer.WriteStringValue(InputPathType.ToString());
            writer.WriteEndObject();
        }

        internal static PackageInputPathVersion DeserializePackageInputPathVersion(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> resourceName = default;
            Optional<string> resourceVersion = default;
            InputPathType inputPathType = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("resourceName"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        resourceName = null;
                        continue;
                    }
                    resourceName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("resourceVersion"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        resourceVersion = null;
                        continue;
                    }
                    resourceVersion = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("inputPathType"u8))
                {
                    inputPathType = new InputPathType(property.Value.GetString());
                    continue;
                }
            }
            return new PackageInputPathVersion(inputPathType, resourceName.Value, resourceVersion.Value);
        }
    }
}
