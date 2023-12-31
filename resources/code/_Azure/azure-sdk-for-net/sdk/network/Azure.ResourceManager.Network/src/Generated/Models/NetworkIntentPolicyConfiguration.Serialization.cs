// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    public partial class NetworkIntentPolicyConfiguration : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(NetworkIntentPolicyName))
            {
                writer.WritePropertyName("networkIntentPolicyName"u8);
                writer.WriteStringValue(NetworkIntentPolicyName);
            }
            if (Optional.IsDefined(SourceNetworkIntentPolicy))
            {
                writer.WritePropertyName("sourceNetworkIntentPolicy"u8);
                writer.WriteObjectValue(SourceNetworkIntentPolicy);
            }
            writer.WriteEndObject();
        }
    }
}
