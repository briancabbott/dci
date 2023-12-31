// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.AppService.Models
{
    internal static partial class AutoHealActionTypeExtensions
    {
        public static string ToSerialString(this AutoHealActionType value) => value switch
        {
            AutoHealActionType.Recycle => "Recycle",
            AutoHealActionType.LogEvent => "LogEvent",
            AutoHealActionType.CustomAction => "CustomAction",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AutoHealActionType value.")
        };

        public static AutoHealActionType ToAutoHealActionType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Recycle")) return AutoHealActionType.Recycle;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "LogEvent")) return AutoHealActionType.LogEvent;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "CustomAction")) return AutoHealActionType.CustomAction;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AutoHealActionType value.");
        }
    }
}
