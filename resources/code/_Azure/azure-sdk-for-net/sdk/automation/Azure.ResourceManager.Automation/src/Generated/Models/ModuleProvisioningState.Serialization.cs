// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Automation.Models
{
    internal static partial class ModuleProvisioningStateExtensions
    {
        public static string ToSerialString(this ModuleProvisioningState value) => value switch
        {
            ModuleProvisioningState.Created => "Created",
            ModuleProvisioningState.Creating => "Creating",
            ModuleProvisioningState.StartingImportModuleRunbook => "StartingImportModuleRunbook",
            ModuleProvisioningState.RunningImportModuleRunbook => "RunningImportModuleRunbook",
            ModuleProvisioningState.ContentRetrieved => "ContentRetrieved",
            ModuleProvisioningState.ContentDownloaded => "ContentDownloaded",
            ModuleProvisioningState.ContentValidated => "ContentValidated",
            ModuleProvisioningState.ConnectionTypeImported => "ConnectionTypeImported",
            ModuleProvisioningState.ContentStored => "ContentStored",
            ModuleProvisioningState.ModuleDataStored => "ModuleDataStored",
            ModuleProvisioningState.ActivitiesStored => "ActivitiesStored",
            ModuleProvisioningState.ModuleImportRunbookComplete => "ModuleImportRunbookComplete",
            ModuleProvisioningState.Succeeded => "Succeeded",
            ModuleProvisioningState.Failed => "Failed",
            ModuleProvisioningState.Cancelled => "Cancelled",
            ModuleProvisioningState.Updating => "Updating",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown ModuleProvisioningState value.")
        };

        public static ModuleProvisioningState ToModuleProvisioningState(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Created")) return ModuleProvisioningState.Created;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Creating")) return ModuleProvisioningState.Creating;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "StartingImportModuleRunbook")) return ModuleProvisioningState.StartingImportModuleRunbook;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "RunningImportModuleRunbook")) return ModuleProvisioningState.RunningImportModuleRunbook;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ContentRetrieved")) return ModuleProvisioningState.ContentRetrieved;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ContentDownloaded")) return ModuleProvisioningState.ContentDownloaded;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ContentValidated")) return ModuleProvisioningState.ContentValidated;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ConnectionTypeImported")) return ModuleProvisioningState.ConnectionTypeImported;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ContentStored")) return ModuleProvisioningState.ContentStored;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ModuleDataStored")) return ModuleProvisioningState.ModuleDataStored;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ActivitiesStored")) return ModuleProvisioningState.ActivitiesStored;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "ModuleImportRunbookComplete")) return ModuleProvisioningState.ModuleImportRunbookComplete;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Succeeded")) return ModuleProvisioningState.Succeeded;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Failed")) return ModuleProvisioningState.Failed;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Cancelled")) return ModuleProvisioningState.Cancelled;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Updating")) return ModuleProvisioningState.Updating;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown ModuleProvisioningState value.");
        }
    }
}
