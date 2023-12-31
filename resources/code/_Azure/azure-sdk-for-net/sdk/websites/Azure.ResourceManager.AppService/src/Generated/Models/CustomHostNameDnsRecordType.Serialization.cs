// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.AppService.Models
{
    internal static partial class CustomHostNameDnsRecordTypeExtensions
    {
        public static string ToSerialString(this CustomHostNameDnsRecordType value) => value switch
        {
            CustomHostNameDnsRecordType.CName => "CName",
            CustomHostNameDnsRecordType.A => "A",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown CustomHostNameDnsRecordType value.")
        };

        public static CustomHostNameDnsRecordType ToCustomHostNameDnsRecordType(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "CName")) return CustomHostNameDnsRecordType.CName;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "A")) return CustomHostNameDnsRecordType.A;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown CustomHostNameDnsRecordType value.");
        }
    }
}
