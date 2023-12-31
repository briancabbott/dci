// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;

namespace Azure.ResourceManager.DevTestLabs.Models
{
    /// <summary> Properties of the disk to attach. </summary>
    public partial class DevTestLabDiskAttachContent
    {
        /// <summary> Initializes a new instance of <see cref="DevTestLabDiskAttachContent"/>. </summary>
        public DevTestLabDiskAttachContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="DevTestLabDiskAttachContent"/>. </summary>
        /// <param name="leasedByLabVmId"> The resource ID of the Lab virtual machine to which the disk is attached. </param>
        internal DevTestLabDiskAttachContent(ResourceIdentifier leasedByLabVmId)
        {
            LeasedByLabVmId = leasedByLabVmId;
        }

        /// <summary> The resource ID of the Lab virtual machine to which the disk is attached. </summary>
        public ResourceIdentifier LeasedByLabVmId { get; set; }
    }
}
