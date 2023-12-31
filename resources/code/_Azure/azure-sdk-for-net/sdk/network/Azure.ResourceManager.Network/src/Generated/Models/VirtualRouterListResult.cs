// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Network;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Response for ListVirtualRouters API service call. </summary>
    internal partial class VirtualRouterListResult
    {
        /// <summary> Initializes a new instance of <see cref="VirtualRouterListResult"/>. </summary>
        internal VirtualRouterListResult()
        {
            Value = new ChangeTrackingList<VirtualRouterData>();
        }

        /// <summary> Initializes a new instance of <see cref="VirtualRouterListResult"/>. </summary>
        /// <param name="value"> List of Virtual Routers. </param>
        /// <param name="nextLink"> URL to get the next set of results. </param>
        internal VirtualRouterListResult(IReadOnlyList<VirtualRouterData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> List of Virtual Routers. </summary>
        public IReadOnlyList<VirtualRouterData> Value { get; }
        /// <summary> URL to get the next set of results. </summary>
        public string NextLink { get; }
    }
}
