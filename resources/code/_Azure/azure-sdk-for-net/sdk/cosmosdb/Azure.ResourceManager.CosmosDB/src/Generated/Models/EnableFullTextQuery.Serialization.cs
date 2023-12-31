// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.CosmosDB.Models
{
    internal static partial class EnableFullTextQueryExtensions
    {
        public static string ToSerialString(this EnableFullTextQuery value) => value switch
        {
            EnableFullTextQuery.None => "None",
            EnableFullTextQuery.True => "True",
            EnableFullTextQuery.False => "False",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown EnableFullTextQuery value.")
        };

        public static EnableFullTextQuery ToEnableFullTextQuery(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "None")) return EnableFullTextQuery.None;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "True")) return EnableFullTextQuery.True;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "False")) return EnableFullTextQuery.False;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown EnableFullTextQuery value.");
        }
    }
}
