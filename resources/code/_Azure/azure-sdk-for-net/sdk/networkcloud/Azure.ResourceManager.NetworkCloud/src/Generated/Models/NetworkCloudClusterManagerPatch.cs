// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.NetworkCloud.Models
{
    /// <summary> ClusterManagerPatchParameters represents the body of the request to patch the cluster properties. </summary>
    public partial class NetworkCloudClusterManagerPatch
    {
        /// <summary> Initializes a new instance of <see cref="NetworkCloudClusterManagerPatch"/>. </summary>
        public NetworkCloudClusterManagerPatch()
        {
            Tags = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="NetworkCloudClusterManagerPatch"/>. </summary>
        /// <param name="tags"> The Azure resource tags that will replace the existing ones. </param>
        internal NetworkCloudClusterManagerPatch(IDictionary<string, string> tags)
        {
            Tags = tags;
        }

        /// <summary> The Azure resource tags that will replace the existing ones. </summary>
        public IDictionary<string, string> Tags { get; }
    }
}
