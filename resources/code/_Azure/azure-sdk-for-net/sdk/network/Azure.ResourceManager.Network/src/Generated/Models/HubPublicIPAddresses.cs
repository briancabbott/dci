// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Public IP addresses associated with azure firewall. </summary>
    public partial class HubPublicIPAddresses
    {
        /// <summary> Initializes a new instance of <see cref="HubPublicIPAddresses"/>. </summary>
        public HubPublicIPAddresses()
        {
            Addresses = new ChangeTrackingList<AzureFirewallPublicIPAddress>();
        }

        /// <summary> Initializes a new instance of <see cref="HubPublicIPAddresses"/>. </summary>
        /// <param name="addresses"> The list of Public IP addresses associated with azure firewall or IP addresses to be retained. </param>
        /// <param name="count"> The number of Public IP addresses associated with azure firewall. </param>
        internal HubPublicIPAddresses(IList<AzureFirewallPublicIPAddress> addresses, int? count)
        {
            Addresses = addresses;
            Count = count;
        }

        /// <summary> The list of Public IP addresses associated with azure firewall or IP addresses to be retained. </summary>
        public IList<AzureFirewallPublicIPAddress> Addresses { get; }
        /// <summary> The number of Public IP addresses associated with azure firewall. </summary>
        public int? Count { get; set; }
    }
}
