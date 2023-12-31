// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.OperationalInsights.Models
{
    internal static partial class OperationalInsightsClusterCapacityExtensions
    {
        public static long ToSerialInt64(this OperationalInsightsClusterCapacity value) => value switch
        {
            OperationalInsightsClusterCapacity.FiveHundred => 500L,
            OperationalInsightsClusterCapacity.TenHundred => 1000L,
            OperationalInsightsClusterCapacity.TwoThousand => 2000L,
            OperationalInsightsClusterCapacity.FiveThousand => 5000L,
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown OperationalInsightsClusterCapacity value.")
        };

        public static OperationalInsightsClusterCapacity ToOperationalInsightsClusterCapacity(this long value)
        {
            if (value == 500L) return OperationalInsightsClusterCapacity.FiveHundred;
            if (value == 1000L) return OperationalInsightsClusterCapacity.TenHundred;
            if (value == 2000L) return OperationalInsightsClusterCapacity.TwoThousand;
            if (value == 5000L) return OperationalInsightsClusterCapacity.FiveThousand;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown OperationalInsightsClusterCapacity value.");
        }
    }
}
