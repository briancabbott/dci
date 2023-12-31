// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.AppService.Models
{
    internal static partial class BuiltInAuthenticationProviderExtensions
    {
        public static string ToSerialString(this BuiltInAuthenticationProvider value) => value switch
        {
            BuiltInAuthenticationProvider.AzureActiveDirectory => "AzureActiveDirectory",
            BuiltInAuthenticationProvider.Facebook => "Facebook",
            BuiltInAuthenticationProvider.Google => "Google",
            BuiltInAuthenticationProvider.MicrosoftAccount => "MicrosoftAccount",
            BuiltInAuthenticationProvider.Twitter => "Twitter",
            BuiltInAuthenticationProvider.Github => "Github",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BuiltInAuthenticationProvider value.")
        };

        public static BuiltInAuthenticationProvider ToBuiltInAuthenticationProvider(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "AzureActiveDirectory")) return BuiltInAuthenticationProvider.AzureActiveDirectory;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Facebook")) return BuiltInAuthenticationProvider.Facebook;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Google")) return BuiltInAuthenticationProvider.Google;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "MicrosoftAccount")) return BuiltInAuthenticationProvider.MicrosoftAccount;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Twitter")) return BuiltInAuthenticationProvider.Twitter;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Github")) return BuiltInAuthenticationProvider.Github;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BuiltInAuthenticationProvider value.");
        }
    }
}
