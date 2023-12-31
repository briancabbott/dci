// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class SiteRecoveryVaultSettingProperties
    {
        internal static SiteRecoveryVaultSettingProperties DeserializeSiteRecoveryVaultSettingProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<ResourceIdentifier> migrationSolutionId = default;
            Optional<string> vmwareToAzureProviderType = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("migrationSolutionId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    migrationSolutionId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("vmwareToAzureProviderType"u8))
                {
                    vmwareToAzureProviderType = property.Value.GetString();
                    continue;
                }
            }
            return new SiteRecoveryVaultSettingProperties(migrationSolutionId.Value, vmwareToAzureProviderType.Value);
        }
    }
}
