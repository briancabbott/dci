// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Network;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Response for ListBackendAddressPool API service call. </summary>
    internal partial class LoadBalancerBackendAddressPoolListResult
    {
        /// <summary> Initializes a new instance of <see cref="LoadBalancerBackendAddressPoolListResult"/>. </summary>
        internal LoadBalancerBackendAddressPoolListResult()
        {
            Value = new ChangeTrackingList<BackendAddressPoolData>();
        }

        /// <summary> Initializes a new instance of <see cref="LoadBalancerBackendAddressPoolListResult"/>. </summary>
        /// <param name="value"> A list of backend address pools in a load balancer. </param>
        /// <param name="nextLink"> The URL to get the next set of results. </param>
        internal LoadBalancerBackendAddressPoolListResult(IReadOnlyList<BackendAddressPoolData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> A list of backend address pools in a load balancer. </summary>
        public IReadOnlyList<BackendAddressPoolData> Value { get; }
        /// <summary> The URL to get the next set of results. </summary>
        public string NextLink { get; }
    }
}
