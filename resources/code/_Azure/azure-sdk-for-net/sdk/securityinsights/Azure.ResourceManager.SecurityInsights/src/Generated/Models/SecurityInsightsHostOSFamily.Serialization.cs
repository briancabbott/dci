// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.SecurityInsights.Models
{
    internal static partial class SecurityInsightsHostOSFamilyExtensions
    {
        public static string ToSerialString(this SecurityInsightsHostOSFamily value) => value switch
        {
            SecurityInsightsHostOSFamily.Unknown => "Unknown",
            SecurityInsightsHostOSFamily.Linux => "Linux",
            SecurityInsightsHostOSFamily.Windows => "Windows",
            SecurityInsightsHostOSFamily.Android => "Android",
            SecurityInsightsHostOSFamily.Ios => "IOS",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SecurityInsightsHostOSFamily value.")
        };

        public static SecurityInsightsHostOSFamily ToSecurityInsightsHostOSFamily(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Unknown")) return SecurityInsightsHostOSFamily.Unknown;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Linux")) return SecurityInsightsHostOSFamily.Linux;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Windows")) return SecurityInsightsHostOSFamily.Windows;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Android")) return SecurityInsightsHostOSFamily.Android;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "IOS")) return SecurityInsightsHostOSFamily.Ios;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SecurityInsightsHostOSFamily value.");
        }
    }
}
