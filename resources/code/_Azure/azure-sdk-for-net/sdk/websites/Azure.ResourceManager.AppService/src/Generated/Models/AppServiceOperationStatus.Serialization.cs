// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.AppService.Models
{
    internal static partial class AppServiceOperationStatusExtensions
    {
        public static string ToSerialString(this AppServiceOperationStatus value) => value switch
        {
            AppServiceOperationStatus.InProgress => "InProgress",
            AppServiceOperationStatus.Failed => "Failed",
            AppServiceOperationStatus.Succeeded => "Succeeded",
            AppServiceOperationStatus.TimedOut => "TimedOut",
            AppServiceOperationStatus.Created => "Created",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AppServiceOperationStatus value.")
        };

        public static AppServiceOperationStatus ToAppServiceOperationStatus(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "InProgress")) return AppServiceOperationStatus.InProgress;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Failed")) return AppServiceOperationStatus.Failed;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Succeeded")) return AppServiceOperationStatus.Succeeded;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "TimedOut")) return AppServiceOperationStatus.TimedOut;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Created")) return AppServiceOperationStatus.Created;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown AppServiceOperationStatus value.");
        }
    }
}
