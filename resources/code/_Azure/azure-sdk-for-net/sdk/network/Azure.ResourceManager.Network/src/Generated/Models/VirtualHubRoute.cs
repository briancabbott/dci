// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> VirtualHub route. </summary>
    public partial class VirtualHubRoute
    {
        /// <summary> Initializes a new instance of <see cref="VirtualHubRoute"/>. </summary>
        public VirtualHubRoute()
        {
            AddressPrefixes = new ChangeTrackingList<string>();
        }

        /// <summary> Initializes a new instance of <see cref="VirtualHubRoute"/>. </summary>
        /// <param name="addressPrefixes"> List of all addressPrefixes. </param>
        /// <param name="nextHopIPAddress"> NextHop ip address. </param>
        internal VirtualHubRoute(IList<string> addressPrefixes, string nextHopIPAddress)
        {
            AddressPrefixes = addressPrefixes;
            NextHopIPAddress = nextHopIPAddress;
        }

        /// <summary> List of all addressPrefixes. </summary>
        public IList<string> AddressPrefixes { get; }
        /// <summary> NextHop ip address. </summary>
        public string NextHopIPAddress { get; set; }
    }
}
