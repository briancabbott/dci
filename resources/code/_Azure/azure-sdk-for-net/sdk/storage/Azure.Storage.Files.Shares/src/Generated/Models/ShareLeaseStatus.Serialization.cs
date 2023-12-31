// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Storage.Files.Shares.Models
{
    internal static partial class ShareLeaseStatusExtensions
    {
        public static string ToSerialString(this ShareLeaseStatus value) => value switch
        {
            ShareLeaseStatus.Locked => "locked",
            ShareLeaseStatus.Unlocked => "unlocked",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown ShareLeaseStatus value.")
        };

        public static ShareLeaseStatus ToShareLeaseStatus(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "locked")) return ShareLeaseStatus.Locked;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "unlocked")) return ShareLeaseStatus.Unlocked;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown ShareLeaseStatus value.");
        }
    }
}
