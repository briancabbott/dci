// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;
using Azure.ResourceManager.Models;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> The information of an AvailablePrivateEndpointType. </summary>
    public partial class AvailablePrivateEndpointType : ResourceData
    {
        /// <summary> Initializes a new instance of <see cref="AvailablePrivateEndpointType"/>. </summary>
        internal AvailablePrivateEndpointType()
        {
        }

        /// <summary> Initializes a new instance of <see cref="AvailablePrivateEndpointType"/>. </summary>
        /// <param name="id"> The id. </param>
        /// <param name="name"> The name. </param>
        /// <param name="resourceType"> The resourceType. </param>
        /// <param name="systemData"> The systemData. </param>
        /// <param name="resourceName"> The name of the service and resource. </param>
        /// <param name="displayName"> Display name of the resource. </param>
        internal AvailablePrivateEndpointType(ResourceIdentifier id, string name, ResourceType resourceType, SystemData systemData, string resourceName, string displayName) : base(id, name, resourceType, systemData)
        {
            ResourceName = resourceName;
            DisplayName = displayName;
        }

        /// <summary> The name of the service and resource. </summary>
        public string ResourceName { get; }
        /// <summary> Display name of the resource. </summary>
        public string DisplayName { get; }
    }
}
