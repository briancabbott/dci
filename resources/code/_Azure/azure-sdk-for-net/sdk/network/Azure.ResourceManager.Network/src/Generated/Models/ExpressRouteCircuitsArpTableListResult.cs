// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Response for ListArpTable associated with the Express Route Circuits API. </summary>
    public partial class ExpressRouteCircuitsArpTableListResult
    {
        /// <summary> Initializes a new instance of <see cref="ExpressRouteCircuitsArpTableListResult"/>. </summary>
        internal ExpressRouteCircuitsArpTableListResult()
        {
            Value = new ChangeTrackingList<ExpressRouteCircuitArpTable>();
        }

        /// <summary> Initializes a new instance of <see cref="ExpressRouteCircuitsArpTableListResult"/>. </summary>
        /// <param name="value"> A list of the ARP tables. </param>
        /// <param name="nextLink"> The URL to get the next set of results. </param>
        internal ExpressRouteCircuitsArpTableListResult(IReadOnlyList<ExpressRouteCircuitArpTable> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> A list of the ARP tables. </summary>
        public IReadOnlyList<ExpressRouteCircuitArpTable> Value { get; }
        /// <summary> The URL to get the next set of results. </summary>
        public string NextLink { get; }
    }
}
