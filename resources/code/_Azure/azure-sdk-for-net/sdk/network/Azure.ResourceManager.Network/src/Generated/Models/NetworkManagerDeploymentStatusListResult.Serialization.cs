// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    internal partial class NetworkManagerDeploymentStatusListResult
    {
        internal static NetworkManagerDeploymentStatusListResult DeserializeNetworkManagerDeploymentStatusListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<NetworkManagerDeploymentStatus>> value = default;
            Optional<string> skipToken = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<NetworkManagerDeploymentStatus> array = new List<NetworkManagerDeploymentStatus>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(NetworkManagerDeploymentStatus.DeserializeNetworkManagerDeploymentStatus(item));
                    }
                    value = array;
                    continue;
                }
                if (property.NameEquals("skipToken"u8))
                {
                    skipToken = property.Value.GetString();
                    continue;
                }
            }
            return new NetworkManagerDeploymentStatusListResult(Optional.ToList(value), skipToken.Value);
        }
    }
}
