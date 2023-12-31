// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Sql.Models
{
    internal static partial class LedgerDigestUploadsStateExtensions
    {
        public static string ToSerialString(this LedgerDigestUploadsState value) => value switch
        {
            LedgerDigestUploadsState.Enabled => "Enabled",
            LedgerDigestUploadsState.Disabled => "Disabled",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown LedgerDigestUploadsState value.")
        };

        public static LedgerDigestUploadsState ToLedgerDigestUploadsState(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Enabled")) return LedgerDigestUploadsState.Enabled;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Disabled")) return LedgerDigestUploadsState.Disabled;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown LedgerDigestUploadsState value.");
        }
    }
}
