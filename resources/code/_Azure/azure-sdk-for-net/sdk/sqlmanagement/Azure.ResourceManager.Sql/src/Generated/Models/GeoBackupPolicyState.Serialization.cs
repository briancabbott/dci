// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Sql.Models
{
    internal static partial class GeoBackupPolicyStateExtensions
    {
        public static string ToSerialString(this GeoBackupPolicyState value) => value switch
        {
            GeoBackupPolicyState.Disabled => "Disabled",
            GeoBackupPolicyState.Enabled => "Enabled",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown GeoBackupPolicyState value.")
        };

        public static GeoBackupPolicyState ToGeoBackupPolicyState(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Disabled")) return GeoBackupPolicyState.Disabled;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Enabled")) return GeoBackupPolicyState.Enabled;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown GeoBackupPolicyState value.");
        }
    }
}
