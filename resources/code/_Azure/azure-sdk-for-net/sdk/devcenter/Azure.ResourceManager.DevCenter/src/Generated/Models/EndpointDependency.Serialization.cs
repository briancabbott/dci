// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.DevCenter.Models
{
    public partial class EndpointDependency
    {
        internal static EndpointDependency DeserializeEndpointDependency(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> domainName = default;
            Optional<string> description = default;
            Optional<IReadOnlyList<DevCenterEndpointDetail>> endpointDetails = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("domainName"u8))
                {
                    domainName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("description"u8))
                {
                    description = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("endpointDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<DevCenterEndpointDetail> array = new List<DevCenterEndpointDetail>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(DevCenterEndpointDetail.DeserializeDevCenterEndpointDetail(item));
                    }
                    endpointDetails = array;
                    continue;
                }
            }
            return new EndpointDependency(domainName.Value, description.Value, Optional.ToList(endpointDetails));
        }
    }
}
