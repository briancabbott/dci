// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesDataReplication.Models
{
    public partial class DataReplicationIdentity : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("tenantId"u8);
            writer.WriteStringValue(TenantId);
            writer.WritePropertyName("applicationId"u8);
            writer.WriteStringValue(ApplicationId);
            writer.WritePropertyName("objectId"u8);
            writer.WriteStringValue(ObjectId);
            writer.WritePropertyName("audience"u8);
            writer.WriteStringValue(Audience);
            writer.WritePropertyName("aadAuthority"u8);
            writer.WriteStringValue(AadAuthority);
            writer.WriteEndObject();
        }

        internal static DataReplicationIdentity DeserializeDataReplicationIdentity(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Guid tenantId = default;
            string applicationId = default;
            string objectId = default;
            string audience = default;
            string aadAuthority = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("tenantId"u8))
                {
                    tenantId = property.Value.GetGuid();
                    continue;
                }
                if (property.NameEquals("applicationId"u8))
                {
                    applicationId = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("objectId"u8))
                {
                    objectId = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("audience"u8))
                {
                    audience = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("aadAuthority"u8))
                {
                    aadAuthority = property.Value.GetString();
                    continue;
                }
            }
            return new DataReplicationIdentity(tenantId, applicationId, objectId, audience, aadAuthority);
        }
    }
}
