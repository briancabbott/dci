// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Resources.Models
{
    /// <summary> List of resource types of a resource provider. </summary>
    internal partial class ProviderResourceTypeListResult
    {
        /// <summary> Initializes a new instance of <see cref="ProviderResourceTypeListResult"/>. </summary>
        internal ProviderResourceTypeListResult()
        {
            Value = new ChangeTrackingList<ProviderResourceType>();
        }

        /// <summary> Initializes a new instance of <see cref="ProviderResourceTypeListResult"/>. </summary>
        /// <param name="value"> An array of resource types. </param>
        /// <param name="nextLink"> The URL to use for getting the next set of results. </param>
        internal ProviderResourceTypeListResult(IReadOnlyList<ProviderResourceType> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> An array of resource types. </summary>
        public IReadOnlyList<ProviderResourceType> Value { get; }
        /// <summary> The URL to use for getting the next set of results. </summary>
        public string NextLink { get; }
    }
}
