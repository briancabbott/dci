// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.CosmosDBForPostgreSql;

namespace Azure.ResourceManager.CosmosDBForPostgreSql.Models
{
    internal partial class CosmosDBForPostgreSqlClusterServerListResult
    {
        internal static CosmosDBForPostgreSqlClusterServerListResult DeserializeCosmosDBForPostgreSqlClusterServerListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<CosmosDBForPostgreSqlClusterServerData>> value = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<CosmosDBForPostgreSqlClusterServerData> array = new List<CosmosDBForPostgreSqlClusterServerData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(CosmosDBForPostgreSqlClusterServerData.DeserializeCosmosDBForPostgreSqlClusterServerData(item));
                    }
                    value = array;
                    continue;
                }
            }
            return new CosmosDBForPostgreSqlClusterServerListResult(Optional.ToList(value));
        }
    }
}
