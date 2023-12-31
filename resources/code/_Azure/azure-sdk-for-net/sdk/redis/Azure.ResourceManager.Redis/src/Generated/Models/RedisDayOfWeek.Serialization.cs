// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Redis.Models
{
    internal static partial class RedisDayOfWeekExtensions
    {
        public static string ToSerialString(this RedisDayOfWeek value) => value switch
        {
            RedisDayOfWeek.Monday => "Monday",
            RedisDayOfWeek.Tuesday => "Tuesday",
            RedisDayOfWeek.Wednesday => "Wednesday",
            RedisDayOfWeek.Thursday => "Thursday",
            RedisDayOfWeek.Friday => "Friday",
            RedisDayOfWeek.Saturday => "Saturday",
            RedisDayOfWeek.Sunday => "Sunday",
            RedisDayOfWeek.Everyday => "Everyday",
            RedisDayOfWeek.Weekend => "Weekend",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown RedisDayOfWeek value.")
        };

        public static RedisDayOfWeek ToRedisDayOfWeek(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Monday")) return RedisDayOfWeek.Monday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Tuesday")) return RedisDayOfWeek.Tuesday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Wednesday")) return RedisDayOfWeek.Wednesday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Thursday")) return RedisDayOfWeek.Thursday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Friday")) return RedisDayOfWeek.Friday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Saturday")) return RedisDayOfWeek.Saturday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Sunday")) return RedisDayOfWeek.Sunday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Everyday")) return RedisDayOfWeek.Everyday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Weekend")) return RedisDayOfWeek.Weekend;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown RedisDayOfWeek value.");
        }
    }
}
