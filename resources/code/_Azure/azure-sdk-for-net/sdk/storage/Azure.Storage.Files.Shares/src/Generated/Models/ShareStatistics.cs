// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Storage.Files.Shares.Models
{
    /// <summary> Stats for the share. </summary>
    public partial class ShareStatistics
    {
        /// <summary> Initializes a new instance of <see cref="ShareStatistics"/>. </summary>
        /// <param name="shareUsageBytes"> The approximate size of the data stored in bytes. Note that this value may not include all recently created or recently resized files. </param>
        internal ShareStatistics(int shareUsageBytes)
        {
            ShareUsageBytes = shareUsageBytes;
        }
    }
}
