// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.CosmosDB.Models
{
    public partial class ListConnectionStringsResult
    {
        internal static ListConnectionStringsResult DeserializeListConnectionStringsResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<CosmosDBConnectionString>> connectionStrings = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("connectionStrings"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<CosmosDBConnectionString> array = new List<CosmosDBConnectionString>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(CosmosDBConnectionString.DeserializeCosmosDBConnectionString(item));
                    }
                    connectionStrings = array;
                    continue;
                }
            }
            return new ListConnectionStringsResult(Optional.ToList(connectionStrings));
        }
    }
}
