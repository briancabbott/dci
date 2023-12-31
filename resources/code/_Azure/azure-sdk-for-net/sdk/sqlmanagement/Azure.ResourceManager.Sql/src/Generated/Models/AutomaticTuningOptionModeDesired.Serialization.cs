// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Sql.Models
{
    internal static partial class AutomaticTuningOptionModeDesiredExtensions
    {
        public static string ToSerialString(this AutomaticTuningOptionModeDesired value) => value switch
        {
            AutomaticTuningOptionModeDesired.Default => "Default",
            AutomaticTuningOptionModeDesired.Off => "Off",
            AutomaticTuningOptionModeDesired.On => "On",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AutomaticTuningOptionModeDesired value.")
        };

        public static AutomaticTuningOptionModeDesired ToAutomaticTuningOptionModeDesired(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Default")) return AutomaticTuningOptionModeDesired.Default;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Off")) return AutomaticTuningOptionModeDesired.Off;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "On")) return AutomaticTuningOptionModeDesired.On;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AutomaticTuningOptionModeDesired value.");
        }
    }
}
