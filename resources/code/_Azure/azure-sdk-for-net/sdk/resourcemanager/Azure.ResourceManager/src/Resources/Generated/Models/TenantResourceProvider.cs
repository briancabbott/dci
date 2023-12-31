// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Resources.Models
{
    /// <summary> Resource provider information. </summary>
    public partial class TenantResourceProvider
    {
        /// <summary> Initializes a new instance of <see cref="TenantResourceProvider"/>. </summary>
        internal TenantResourceProvider()
        {
            ResourceTypes = new ChangeTrackingList<ProviderResourceType>();
        }

        /// <summary> Initializes a new instance of <see cref="TenantResourceProvider"/>. </summary>
        /// <param name="namespace"> The namespace of the resource provider. </param>
        /// <param name="resourceTypes"> The collection of provider resource types. </param>
        internal TenantResourceProvider(string @namespace, IReadOnlyList<ProviderResourceType> resourceTypes)
        {
            Namespace = @namespace;
            ResourceTypes = resourceTypes;
        }

        /// <summary> The namespace of the resource provider. </summary>
        public string Namespace { get; }
        /// <summary> The collection of provider resource types. </summary>
        public IReadOnlyList<ProviderResourceType> ResourceTypes { get; }
    }
}
