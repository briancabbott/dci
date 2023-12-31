// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Synapse.Models
{
    internal static partial class SynapseDataMaskingFunctionExtensions
    {
        public static string ToSerialString(this SynapseDataMaskingFunction value) => value switch
        {
            SynapseDataMaskingFunction.Default => "Default",
            SynapseDataMaskingFunction.Ccn => "CCN",
            SynapseDataMaskingFunction.Email => "Email",
            SynapseDataMaskingFunction.Number => "Number",
            SynapseDataMaskingFunction.Ssn => "SSN",
            SynapseDataMaskingFunction.Text => "Text",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SynapseDataMaskingFunction value.")
        };

        public static SynapseDataMaskingFunction ToSynapseDataMaskingFunction(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Default")) return SynapseDataMaskingFunction.Default;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "CCN")) return SynapseDataMaskingFunction.Ccn;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Email")) return SynapseDataMaskingFunction.Email;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Number")) return SynapseDataMaskingFunction.Number;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "SSN")) return SynapseDataMaskingFunction.Ssn;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Text")) return SynapseDataMaskingFunction.Text;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown SynapseDataMaskingFunction value.");
        }
    }
}
