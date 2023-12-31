// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.CosmosDB.Models
{
    internal static partial class CosmosDBSqlRoleDefinitionTypeExtensions
    {
        public static string ToSerialString(this CosmosDBSqlRoleDefinitionType value) => value switch
        {
            CosmosDBSqlRoleDefinitionType.BuiltInRole => "BuiltInRole",
            CosmosDBSqlRoleDefinitionType.CustomRole => "CustomRole",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown CosmosDBSqlRoleDefinitionType value.")
        };

        public static CosmosDBSqlRoleDefinitionType ToCosmosDBSqlRoleDefinitionType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "BuiltInRole")) return CosmosDBSqlRoleDefinitionType.BuiltInRole;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "CustomRole")) return CosmosDBSqlRoleDefinitionType.CustomRole;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown CosmosDBSqlRoleDefinitionType value.");
        }
    }
}
