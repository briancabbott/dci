// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.RecoveryServices.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Capabilities information
    /// </summary>
    public partial class CapabilitiesProperties
    {
        /// <summary>
        /// Initializes a new instance of the CapabilitiesProperties class.
        /// </summary>
        public CapabilitiesProperties()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the CapabilitiesProperties class.
        /// </summary>
        public CapabilitiesProperties(IList<DNSZone> dnsZones = default(IList<DNSZone>))
        {
            DnsZones = dnsZones;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "dnsZones")]
        public IList<DNSZone> DnsZones { get; set; }

    }
}
