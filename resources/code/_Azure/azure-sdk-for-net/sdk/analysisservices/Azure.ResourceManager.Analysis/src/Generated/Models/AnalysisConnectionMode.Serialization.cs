// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Analysis.Models
{
    internal static partial class AnalysisConnectionModeExtensions
    {
        public static string ToSerialString(this AnalysisConnectionMode value) => value switch
        {
            AnalysisConnectionMode.All => "All",
            AnalysisConnectionMode.ReadOnly => "ReadOnly",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AnalysisConnectionMode value.")
        };

        public static AnalysisConnectionMode ToAnalysisConnectionMode(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "All")) return AnalysisConnectionMode.All;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ReadOnly")) return AnalysisConnectionMode.ReadOnly;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AnalysisConnectionMode value.");
        }
    }
}
