// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ApiManagement.Models
{
    public partial class ApiManagementProductPatch : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("properties"u8);
            writer.WriteStartObject();
            if (Optional.IsDefined(Description))
            {
                writer.WritePropertyName("description"u8);
                writer.WriteStringValue(Description);
            }
            if (Optional.IsDefined(Terms))
            {
                writer.WritePropertyName("terms"u8);
                writer.WriteStringValue(Terms);
            }
            if (Optional.IsDefined(IsSubscriptionRequired))
            {
                writer.WritePropertyName("subscriptionRequired"u8);
                writer.WriteBooleanValue(IsSubscriptionRequired.Value);
            }
            if (Optional.IsDefined(IsApprovalRequired))
            {
                writer.WritePropertyName("approvalRequired"u8);
                writer.WriteBooleanValue(IsApprovalRequired.Value);
            }
            if (Optional.IsDefined(SubscriptionsLimit))
            {
                writer.WritePropertyName("subscriptionsLimit"u8);
                writer.WriteNumberValue(SubscriptionsLimit.Value);
            }
            if (Optional.IsDefined(State))
            {
                writer.WritePropertyName("state"u8);
                writer.WriteStringValue(State.Value.ToSerialString());
            }
            if (Optional.IsDefined(DisplayName))
            {
                writer.WritePropertyName("displayName"u8);
                writer.WriteStringValue(DisplayName);
            }
            writer.WriteEndObject();
            writer.WriteEndObject();
        }
    }
}
