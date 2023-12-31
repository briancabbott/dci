// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.ServiceFabricManagedClusters.Models
{
    /// <summary> Managed cluster update request. </summary>
    public partial class ServiceFabricManagedClusterPatch
    {
        /// <summary> Initializes a new instance of <see cref="ServiceFabricManagedClusterPatch"/>. </summary>
        public ServiceFabricManagedClusterPatch()
        {
            Tags = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="ServiceFabricManagedClusterPatch"/>. </summary>
        /// <param name="tags"> Managed cluster update parameters. </param>
        internal ServiceFabricManagedClusterPatch(IDictionary<string, string> tags)
        {
            Tags = tags;
        }

        /// <summary> Managed cluster update parameters. </summary>
        public IDictionary<string, string> Tags { get; }
    }
}
