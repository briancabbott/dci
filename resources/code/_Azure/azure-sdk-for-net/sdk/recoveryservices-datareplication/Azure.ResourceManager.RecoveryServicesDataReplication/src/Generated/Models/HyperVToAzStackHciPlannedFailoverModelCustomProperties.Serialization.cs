// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesDataReplication.Models
{
    public partial class HyperVToAzStackHciPlannedFailoverModelCustomProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("shutdownSourceVM"u8);
            writer.WriteBooleanValue(ShutdownSourceVm);
            writer.WritePropertyName("instanceType"u8);
            writer.WriteStringValue(InstanceType);
            writer.WriteEndObject();
        }

        internal static HyperVToAzStackHciPlannedFailoverModelCustomProperties DeserializeHyperVToAzStackHciPlannedFailoverModelCustomProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            bool shutdownSourceVm = default;
            string instanceType = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("shutdownSourceVM"u8))
                {
                    shutdownSourceVm = property.Value.GetBoolean();
                    continue;
                }
                if (property.NameEquals("instanceType"u8))
                {
                    instanceType = property.Value.GetString();
                    continue;
                }
            }
            return new HyperVToAzStackHciPlannedFailoverModelCustomProperties(instanceType, shutdownSourceVm);
        }
    }
}
