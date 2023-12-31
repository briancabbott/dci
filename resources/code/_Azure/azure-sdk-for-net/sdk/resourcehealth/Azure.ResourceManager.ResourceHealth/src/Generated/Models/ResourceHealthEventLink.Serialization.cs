// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.ResourceHealth.Models
{
    public partial class ResourceHealthEventLink
    {
        internal static ResourceHealthEventLink DeserializeResourceHealthEventLink(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<ResourceHealthEventLinkTypeValue> type = default;
            Optional<ResourceHealthEventLinkDisplayText> displayText = default;
            Optional<string> extensionName = default;
            Optional<string> bladeName = default;
            Optional<BinaryData> parameters = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("type"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    type = new ResourceHealthEventLinkTypeValue(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("displayText"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    displayText = ResourceHealthEventLinkDisplayText.DeserializeResourceHealthEventLinkDisplayText(property.Value);
                    continue;
                }
                if (property.NameEquals("extensionName"u8))
                {
                    extensionName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("bladeName"u8))
                {
                    bladeName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("parameters"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    parameters = BinaryData.FromString(property.Value.GetRawText());
                    continue;
                }
            }
            return new ResourceHealthEventLink(Optional.ToNullable(type), displayText.Value, extensionName.Value, bladeName.Value, parameters.Value);
        }
    }
}
