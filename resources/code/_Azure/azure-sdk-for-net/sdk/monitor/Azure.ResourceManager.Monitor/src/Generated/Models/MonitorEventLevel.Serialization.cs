// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Monitor.Models
{
    internal static partial class MonitorEventLevelExtensions
    {
        public static string ToSerialString(this MonitorEventLevel value) => value switch
        {
            MonitorEventLevel.Critical => "Critical",
            MonitorEventLevel.Error => "Error",
            MonitorEventLevel.Warning => "Warning",
            MonitorEventLevel.Informational => "Informational",
            MonitorEventLevel.Verbose => "Verbose",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown MonitorEventLevel value.")
        };

        public static MonitorEventLevel ToMonitorEventLevel(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Critical")) return MonitorEventLevel.Critical;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Error")) return MonitorEventLevel.Error;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Warning")) return MonitorEventLevel.Warning;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Informational")) return MonitorEventLevel.Informational;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Verbose")) return MonitorEventLevel.Verbose;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown MonitorEventLevel value.");
        }
    }
}
