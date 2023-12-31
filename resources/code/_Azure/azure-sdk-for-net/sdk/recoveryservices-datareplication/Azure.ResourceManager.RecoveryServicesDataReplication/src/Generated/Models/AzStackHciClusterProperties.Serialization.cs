// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesDataReplication.Models
{
    public partial class AzStackHciClusterProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("clusterName"u8);
            writer.WriteStringValue(ClusterName);
            writer.WritePropertyName("resourceName"u8);
            writer.WriteStringValue(ResourceName);
            writer.WritePropertyName("storageAccountName"u8);
            writer.WriteStringValue(StorageAccountName);
            writer.WritePropertyName("storageContainers"u8);
            writer.WriteStartArray();
            foreach (var item in StorageContainers)
            {
                writer.WriteObjectValue(item);
            }
            writer.WriteEndArray();
            writer.WriteEndObject();
        }

        internal static AzStackHciClusterProperties DeserializeAzStackHciClusterProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            string clusterName = default;
            string resourceName = default;
            string storageAccountName = default;
            IList<StorageContainerProperties> storageContainers = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("clusterName"u8))
                {
                    clusterName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("resourceName"u8))
                {
                    resourceName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("storageAccountName"u8))
                {
                    storageAccountName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("storageContainers"u8))
                {
                    List<StorageContainerProperties> array = new List<StorageContainerProperties>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(StorageContainerProperties.DeserializeStorageContainerProperties(item));
                    }
                    storageContainers = array;
                    continue;
                }
            }
            return new AzStackHciClusterProperties(clusterName, resourceName, storageAccountName, storageContainers);
        }
    }
}
