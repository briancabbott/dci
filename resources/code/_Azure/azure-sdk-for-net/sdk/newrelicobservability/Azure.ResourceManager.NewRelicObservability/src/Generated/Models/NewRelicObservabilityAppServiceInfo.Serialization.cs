// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.NewRelicObservability.Models
{
    public partial class NewRelicObservabilityAppServiceInfo
    {
        internal static NewRelicObservabilityAppServiceInfo DeserializeNewRelicObservabilityAppServiceInfo(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<ResourceIdentifier> azureResourceId = default;
            Optional<string> agentVersion = default;
            Optional<string> agentStatus = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("azureResourceId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    azureResourceId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("agentVersion"u8))
                {
                    agentVersion = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("agentStatus"u8))
                {
                    agentStatus = property.Value.GetString();
                    continue;
                }
            }
            return new NewRelicObservabilityAppServiceInfo(azureResourceId.Value, agentVersion.Value, agentStatus.Value);
        }
    }
}
