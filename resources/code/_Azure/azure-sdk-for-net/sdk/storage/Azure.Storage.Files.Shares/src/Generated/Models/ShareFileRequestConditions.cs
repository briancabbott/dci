// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Storage.Files.Shares.Models
{
    /// <summary> Parameter group. </summary>
    public partial class ShareFileRequestConditions
    {
        /// <summary> Initializes a new instance of <see cref="ShareFileRequestConditions"/>. </summary>
        public ShareFileRequestConditions()
        {
        }

        /// <summary> Initializes a new instance of <see cref="ShareFileRequestConditions"/>. </summary>
        /// <param name="leaseId"> If specified, the operation only succeeds if the resource's lease is active and matches this ID. </param>
        internal ShareFileRequestConditions(string leaseId)
        {
            LeaseId = leaseId;
        }
    }
}
