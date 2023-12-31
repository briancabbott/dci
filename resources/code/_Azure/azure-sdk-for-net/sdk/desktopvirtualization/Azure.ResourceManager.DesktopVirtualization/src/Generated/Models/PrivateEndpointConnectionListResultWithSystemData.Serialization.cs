// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.DesktopVirtualization;

namespace Azure.ResourceManager.DesktopVirtualization.Models
{
    internal partial class PrivateEndpointConnectionListResultWithSystemData
    {
        internal static PrivateEndpointConnectionListResultWithSystemData DeserializePrivateEndpointConnectionListResultWithSystemData(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<DesktopVirtualizationPrivateEndpointConnectionDataData>> value = default;
            Optional<string> nextLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<DesktopVirtualizationPrivateEndpointConnectionDataData> array = new List<DesktopVirtualizationPrivateEndpointConnectionDataData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(DesktopVirtualizationPrivateEndpointConnectionDataData.DeserializeDesktopVirtualizationPrivateEndpointConnectionDataData(item));
                    }
                    value = array;
                    continue;
                }
                if (property.NameEquals("nextLink"u8))
                {
                    nextLink = property.Value.GetString();
                    continue;
                }
            }
            return new PrivateEndpointConnectionListResultWithSystemData(Optional.ToList(value), nextLink.Value);
        }
    }
}
