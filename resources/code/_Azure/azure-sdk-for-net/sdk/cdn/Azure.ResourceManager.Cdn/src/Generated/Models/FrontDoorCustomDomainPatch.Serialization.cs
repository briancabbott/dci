// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Cdn.Models
{
    public partial class FrontDoorCustomDomainPatch : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("properties"u8);
            writer.WriteStartObject();
            if (Optional.IsDefined(TlsSettings))
            {
                writer.WritePropertyName("tlsSettings"u8);
                writer.WriteObjectValue(TlsSettings);
            }
            if (Optional.IsDefined(DnsZone))
            {
                writer.WritePropertyName("azureDnsZone"u8);
                JsonSerializer.Serialize(writer, DnsZone);
            }
            if (Optional.IsDefined(PreValidatedCustomDomainResource))
            {
                if (PreValidatedCustomDomainResource != null)
                {
                    writer.WritePropertyName("preValidatedCustomDomainResourceId"u8);
                    writer.WriteObjectValue(PreValidatedCustomDomainResource);
                }
                else
                {
                    writer.WriteNull("preValidatedCustomDomainResourceId");
                }
            }
            writer.WriteEndObject();
            writer.WriteEndObject();
        }
    }
}
