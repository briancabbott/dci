// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Cdn.Models
{
    internal static partial class SharedPrivateLinkResourceStatusExtensions
    {
        public static string ToSerialString(this SharedPrivateLinkResourceStatus value) => value switch
        {
            SharedPrivateLinkResourceStatus.Pending => "Pending",
            SharedPrivateLinkResourceStatus.Approved => "Approved",
            SharedPrivateLinkResourceStatus.Rejected => "Rejected",
            SharedPrivateLinkResourceStatus.Disconnected => "Disconnected",
            SharedPrivateLinkResourceStatus.Timeout => "Timeout",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SharedPrivateLinkResourceStatus value.")
        };

        public static SharedPrivateLinkResourceStatus ToSharedPrivateLinkResourceStatus(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Pending")) return SharedPrivateLinkResourceStatus.Pending;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Approved")) return SharedPrivateLinkResourceStatus.Approved;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Rejected")) return SharedPrivateLinkResourceStatus.Rejected;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Disconnected")) return SharedPrivateLinkResourceStatus.Disconnected;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Timeout")) return SharedPrivateLinkResourceStatus.Timeout;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SharedPrivateLinkResourceStatus value.");
        }
    }
}
