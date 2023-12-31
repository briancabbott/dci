// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.RecoveryServicesBackup.Models
{
    internal static partial class BackupTargetDiskNetworkAccessOptionExtensions
    {
        public static string ToSerialString(this BackupTargetDiskNetworkAccessOption value) => value switch
        {
            BackupTargetDiskNetworkAccessOption.SameAsOnSourceDisks => "SameAsOnSourceDisks",
            BackupTargetDiskNetworkAccessOption.EnablePrivateAccessForAllDisks => "EnablePrivateAccessForAllDisks",
            BackupTargetDiskNetworkAccessOption.EnablePublicAccessForAllDisks => "EnablePublicAccessForAllDisks",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BackupTargetDiskNetworkAccessOption value.")
        };

        public static BackupTargetDiskNetworkAccessOption ToBackupTargetDiskNetworkAccessOption(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "SameAsOnSourceDisks")) return BackupTargetDiskNetworkAccessOption.SameAsOnSourceDisks;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "EnablePrivateAccessForAllDisks")) return BackupTargetDiskNetworkAccessOption.EnablePrivateAccessForAllDisks;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "EnablePublicAccessForAllDisks")) return BackupTargetDiskNetworkAccessOption.EnablePublicAccessForAllDisks;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BackupTargetDiskNetworkAccessOption value.");
        }
    }
}
