// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.AppService.Models
{
    internal static partial class AppServiceResourceTypeExtensions
    {
        public static string ToSerialString(this AppServiceResourceType value) => value switch
        {
            AppServiceResourceType.Website => "Website",
            AppServiceResourceType.TrafficManager => "TrafficManager",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AppServiceResourceType value.")
        };

        public static AppServiceResourceType ToAppServiceResourceType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Website")) return AppServiceResourceType.Website;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "TrafficManager")) return AppServiceResourceType.TrafficManager;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AppServiceResourceType value.");
        }
    }
}
