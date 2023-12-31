// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.SqlVirtualMachine.Models
{
    internal static partial class SqlVmAutoPatchingDayOfWeekExtensions
    {
        public static string ToSerialString(this SqlVmAutoPatchingDayOfWeek value) => value switch
        {
            SqlVmAutoPatchingDayOfWeek.Everyday => "Everyday",
            SqlVmAutoPatchingDayOfWeek.Monday => "Monday",
            SqlVmAutoPatchingDayOfWeek.Tuesday => "Tuesday",
            SqlVmAutoPatchingDayOfWeek.Wednesday => "Wednesday",
            SqlVmAutoPatchingDayOfWeek.Thursday => "Thursday",
            SqlVmAutoPatchingDayOfWeek.Friday => "Friday",
            SqlVmAutoPatchingDayOfWeek.Saturday => "Saturday",
            SqlVmAutoPatchingDayOfWeek.Sunday => "Sunday",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SqlVmAutoPatchingDayOfWeek value.")
        };

        public static SqlVmAutoPatchingDayOfWeek ToSqlVmAutoPatchingDayOfWeek(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Everyday")) return SqlVmAutoPatchingDayOfWeek.Everyday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Monday")) return SqlVmAutoPatchingDayOfWeek.Monday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Tuesday")) return SqlVmAutoPatchingDayOfWeek.Tuesday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Wednesday")) return SqlVmAutoPatchingDayOfWeek.Wednesday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Thursday")) return SqlVmAutoPatchingDayOfWeek.Thursday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Friday")) return SqlVmAutoPatchingDayOfWeek.Friday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Saturday")) return SqlVmAutoPatchingDayOfWeek.Saturday;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Sunday")) return SqlVmAutoPatchingDayOfWeek.Sunday;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SqlVmAutoPatchingDayOfWeek value.");
        }
    }
}
