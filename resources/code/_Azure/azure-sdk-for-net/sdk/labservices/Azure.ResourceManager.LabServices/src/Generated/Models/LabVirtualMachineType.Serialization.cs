// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.LabServices.Models
{
    internal static partial class LabVirtualMachineTypeExtensions
    {
        public static string ToSerialString(this LabVirtualMachineType value) => value switch
        {
            LabVirtualMachineType.User => "User",
            LabVirtualMachineType.Template => "Template",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown LabVirtualMachineType value.")
        };

        public static LabVirtualMachineType ToLabVirtualMachineType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "User")) return LabVirtualMachineType.User;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Template")) return LabVirtualMachineType.Template;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown LabVirtualMachineType value.");
        }
    }
}
