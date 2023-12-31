// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Monitor.Models
{
    internal static partial class MonitorAggregationTypeExtensions
    {
        public static string ToSerialString(this MonitorAggregationType value) => value switch
        {
            MonitorAggregationType.None => "None",
            MonitorAggregationType.Average => "Average",
            MonitorAggregationType.Count => "Count",
            MonitorAggregationType.Minimum => "Minimum",
            MonitorAggregationType.Maximum => "Maximum",
            MonitorAggregationType.Total => "Total",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown MonitorAggregationType value.")
        };

        public static MonitorAggregationType ToMonitorAggregationType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "None")) return MonitorAggregationType.None;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Average")) return MonitorAggregationType.Average;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Count")) return MonitorAggregationType.Count;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Minimum")) return MonitorAggregationType.Minimum;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Maximum")) return MonitorAggregationType.Maximum;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Total")) return MonitorAggregationType.Total;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown MonitorAggregationType value.");
        }
    }
}
