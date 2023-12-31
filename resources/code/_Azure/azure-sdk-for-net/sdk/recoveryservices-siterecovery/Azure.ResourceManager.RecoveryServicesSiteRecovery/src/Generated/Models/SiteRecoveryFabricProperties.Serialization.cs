// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class SiteRecoveryFabricProperties
    {
        internal static SiteRecoveryFabricProperties DeserializeSiteRecoveryFabricProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> friendlyName = default;
            Optional<SiteRecoveryEncryptionDetails> encryptionDetails = default;
            Optional<SiteRecoveryEncryptionDetails> rolloverEncryptionDetails = default;
            Optional<string> internalIdentifier = default;
            Optional<string> bcdrState = default;
            Optional<FabricSpecificDetails> customDetails = default;
            Optional<IReadOnlyList<SiteRecoveryHealthError>> healthErrorDetails = default;
            Optional<string> health = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("friendlyName"u8))
                {
                    friendlyName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("encryptionDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    encryptionDetails = SiteRecoveryEncryptionDetails.DeserializeSiteRecoveryEncryptionDetails(property.Value);
                    continue;
                }
                if (property.NameEquals("rolloverEncryptionDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    rolloverEncryptionDetails = SiteRecoveryEncryptionDetails.DeserializeSiteRecoveryEncryptionDetails(property.Value);
                    continue;
                }
                if (property.NameEquals("internalIdentifier"u8))
                {
                    internalIdentifier = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("bcdrState"u8))
                {
                    bcdrState = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("customDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    customDetails = FabricSpecificDetails.DeserializeFabricSpecificDetails(property.Value);
                    continue;
                }
                if (property.NameEquals("healthErrorDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<SiteRecoveryHealthError> array = new List<SiteRecoveryHealthError>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(SiteRecoveryHealthError.DeserializeSiteRecoveryHealthError(item));
                    }
                    healthErrorDetails = array;
                    continue;
                }
                if (property.NameEquals("health"u8))
                {
                    health = property.Value.GetString();
                    continue;
                }
            }
            return new SiteRecoveryFabricProperties(friendlyName.Value, encryptionDetails.Value, rolloverEncryptionDetails.Value, internalIdentifier.Value, bcdrState.Value, customDetails.Value, Optional.ToList(healthErrorDetails), health.Value);
        }
    }
}
