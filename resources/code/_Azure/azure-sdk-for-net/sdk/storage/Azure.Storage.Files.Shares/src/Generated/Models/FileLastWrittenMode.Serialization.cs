// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Storage.Files.Shares.Models
{
    internal static partial class FileLastWrittenModeExtensions
    {
        public static string ToSerialString(this FileLastWrittenMode value) => value switch
        {
            FileLastWrittenMode.Now => "Now",
            FileLastWrittenMode.Preserve => "Preserve",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown FileLastWrittenMode value.")
        };

        public static FileLastWrittenMode ToFileLastWrittenMode(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Now")) return FileLastWrittenMode.Now;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Preserve")) return FileLastWrittenMode.Preserve;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown FileLastWrittenMode value.");
        }
    }
}
