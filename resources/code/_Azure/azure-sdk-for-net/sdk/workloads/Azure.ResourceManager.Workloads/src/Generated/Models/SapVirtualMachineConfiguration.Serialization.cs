// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Workloads.Models
{
    public partial class SapVirtualMachineConfiguration : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("vmSize"u8);
            writer.WriteStringValue(VmSize);
            writer.WritePropertyName("imageReference"u8);
            writer.WriteObjectValue(ImageReference);
            writer.WritePropertyName("osProfile"u8);
            writer.WriteObjectValue(OSProfile);
            writer.WriteEndObject();
        }

        internal static SapVirtualMachineConfiguration DeserializeSapVirtualMachineConfiguration(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            string vmSize = default;
            SapImageReference imageReference = default;
            SapOSProfile osProfile = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("vmSize"u8))
                {
                    vmSize = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("imageReference"u8))
                {
                    imageReference = SapImageReference.DeserializeSapImageReference(property.Value);
                    continue;
                }
                if (property.NameEquals("osProfile"u8))
                {
                    osProfile = SapOSProfile.DeserializeSapOSProfile(property.Value);
                    continue;
                }
            }
            return new SapVirtualMachineConfiguration(vmSize, imageReference, osProfile);
        }
    }
}
