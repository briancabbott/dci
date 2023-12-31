// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class DisableProtectionProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(DisableProtectionReason))
            {
                writer.WritePropertyName("disableProtectionReason"u8);
                writer.WriteStringValue(DisableProtectionReason.Value.ToString());
            }
            if (Optional.IsDefined(ReplicationProviderContent))
            {
                writer.WritePropertyName("replicationProviderInput"u8);
                writer.WriteObjectValue(ReplicationProviderContent);
            }
            writer.WriteEndObject();
        }
    }
}
