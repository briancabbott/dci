// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> DiskExclusionInput when doing enable protection of virtual machine in InMage provider. </summary>
    public partial class InMageDiskExclusionContent
    {
        /// <summary> Initializes a new instance of <see cref="InMageDiskExclusionContent"/>. </summary>
        public InMageDiskExclusionContent()
        {
            VolumeOptions = new ChangeTrackingList<InMageVolumeExclusionOptions>();
            DiskSignatureOptions = new ChangeTrackingList<InMageDiskSignatureExclusionOptions>();
        }

        /// <summary> Initializes a new instance of <see cref="InMageDiskExclusionContent"/>. </summary>
        /// <param name="volumeOptions"> The volume label based option for disk exclusion. </param>
        /// <param name="diskSignatureOptions"> The guest disk signature based option for disk exclusion. </param>
        internal InMageDiskExclusionContent(IList<InMageVolumeExclusionOptions> volumeOptions, IList<InMageDiskSignatureExclusionOptions> diskSignatureOptions)
        {
            VolumeOptions = volumeOptions;
            DiskSignatureOptions = diskSignatureOptions;
        }

        /// <summary> The volume label based option for disk exclusion. </summary>
        public IList<InMageVolumeExclusionOptions> VolumeOptions { get; }
        /// <summary> The guest disk signature based option for disk exclusion. </summary>
        public IList<InMageDiskSignatureExclusionOptions> DiskSignatureOptions { get; }
    }
}
