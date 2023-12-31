// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class SiteRecoveryAddVCenterProperties : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(FriendlyName))
            {
                writer.WritePropertyName("friendlyName"u8);
                writer.WriteStringValue(FriendlyName);
            }
            if (Optional.IsDefined(IPAddress))
            {
                writer.WritePropertyName("ipAddress"u8);
                writer.WriteStringValue(IPAddress.ToString());
            }
            if (Optional.IsDefined(ProcessServerId))
            {
                writer.WritePropertyName("processServerId"u8);
                writer.WriteStringValue(ProcessServerId.Value);
            }
            if (Optional.IsDefined(Port))
            {
                writer.WritePropertyName("port"u8);
                writer.WriteStringValue(Port);
            }
            if (Optional.IsDefined(RunAsAccountId))
            {
                writer.WritePropertyName("runAsAccountId"u8);
                writer.WriteStringValue(RunAsAccountId);
            }
            writer.WriteEndObject();
        }
    }
}
