// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.Batch.Models
{
    internal static partial class BatchNodeDeallocationOptionExtensions
    {
        public static string ToSerialString(this BatchNodeDeallocationOption value) => value switch
        {
            BatchNodeDeallocationOption.Requeue => "Requeue",
            BatchNodeDeallocationOption.Terminate => "Terminate",
            BatchNodeDeallocationOption.TaskCompletion => "TaskCompletion",
            BatchNodeDeallocationOption.RetainedData => "RetainedData",
            _ => throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BatchNodeDeallocationOption value.")
        };

        public static BatchNodeDeallocationOption ToBatchNodeDeallocationOption(this string value)
        {
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Requeue")) return BatchNodeDeallocationOption.Requeue;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "Terminate")) return BatchNodeDeallocationOption.Terminate;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "TaskCompletion")) return BatchNodeDeallocationOption.TaskCompletion;
            if (StringComparer.OrdinalIgnoreCase.Equals(value, "RetainedData")) return BatchNodeDeallocationOption.RetainedData;
            throw new ArgumentOutOfRangeException(nameof(value), value, "Unknown BatchNodeDeallocationOption value.");
        }
    }
}
