// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Batch.Models
{
    internal static partial class BatchUserAccountElevationLevelExtensions
    {
        public static string ToSerialString(this BatchUserAccountElevationLevel value) => value switch
        {
            BatchUserAccountElevationLevel.NonAdmin => "NonAdmin",
            BatchUserAccountElevationLevel.Admin => "Admin",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BatchUserAccountElevationLevel value.")
        };

        public static BatchUserAccountElevationLevel ToBatchUserAccountElevationLevel(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "NonAdmin")) return BatchUserAccountElevationLevel.NonAdmin;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Admin")) return BatchUserAccountElevationLevel.Admin;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BatchUserAccountElevationLevel value.");
        }
    }
}
