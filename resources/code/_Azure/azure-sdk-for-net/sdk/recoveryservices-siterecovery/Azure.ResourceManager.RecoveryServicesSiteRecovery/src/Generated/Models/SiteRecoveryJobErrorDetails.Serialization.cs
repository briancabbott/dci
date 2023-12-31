// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class SiteRecoveryJobErrorDetails
    {
        internal static SiteRecoveryJobErrorDetails DeserializeSiteRecoveryJobErrorDetails(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<SiteRecoveryServiceError> serviceErrorDetails = default;
            Optional<SiteRecoveryJobProviderError> providerErrorDetails = default;
            Optional<string> errorLevel = default;
            Optional<DateTimeOffset> creationTime = default;
            Optional<string> taskId = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("serviceErrorDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    serviceErrorDetails = SiteRecoveryServiceError.DeserializeSiteRecoveryServiceError(property.Value);
                    continue;
                }
                if (property.NameEquals("providerErrorDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    providerErrorDetails = SiteRecoveryJobProviderError.DeserializeSiteRecoveryJobProviderError(property.Value);
                    continue;
                }
                if (property.NameEquals("errorLevel"u8))
                {
                    errorLevel = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("creationTime"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    creationTime = property.Value.GetDateTimeOffset("O");
                    continue;
                }
                if (property.NameEquals("taskId"u8))
                {
                    taskId = property.Value.GetString();
                    continue;
                }
            }
            return new SiteRecoveryJobErrorDetails(serviceErrorDetails.Value, providerErrorDetails.Value, errorLevel.Value, Optional.ToNullable(creationTime), taskId.Value);
        }
    }
}
