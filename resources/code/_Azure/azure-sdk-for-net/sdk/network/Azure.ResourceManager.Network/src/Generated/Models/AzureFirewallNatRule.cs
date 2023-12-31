// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Properties of a NAT rule. </summary>
    public partial class AzureFirewallNatRule
    {
        /// <summary> Initializes a new instance of <see cref="AzureFirewallNatRule"/>. </summary>
        public AzureFirewallNatRule()
        {
            SourceAddresses = new ChangeTrackingList<string>();
            DestinationAddresses = new ChangeTrackingList<string>();
            DestinationPorts = new ChangeTrackingList<string>();
            Protocols = new ChangeTrackingList<AzureFirewallNetworkRuleProtocol>();
            SourceIPGroups = new ChangeTrackingList<string>();
        }

        /// <summary> Initializes a new instance of <see cref="AzureFirewallNatRule"/>. </summary>
        /// <param name="name"> Name of the NAT rule. </param>
        /// <param name="description"> Description of the rule. </param>
        /// <param name="sourceAddresses"> List of source IP addresses for this rule. </param>
        /// <param name="destinationAddresses"> List of destination IP addresses for this rule. Supports IP ranges, prefixes, and service tags. </param>
        /// <param name="destinationPorts"> List of destination ports. </param>
        /// <param name="protocols"> Array of AzureFirewallNetworkRuleProtocols applicable to this NAT rule. </param>
        /// <param name="translatedAddress"> The translated address for this NAT rule. </param>
        /// <param name="translatedPort"> The translated port for this NAT rule. </param>
        /// <param name="translatedFqdn"> The translated FQDN for this NAT rule. </param>
        /// <param name="sourceIPGroups"> List of source IpGroups for this rule. </param>
        internal AzureFirewallNatRule(string name, string description, IList<string> sourceAddresses, IList<string> destinationAddresses, IList<string> destinationPorts, IList<AzureFirewallNetworkRuleProtocol> protocols, string translatedAddress, string translatedPort, string translatedFqdn, IList<string> sourceIPGroups)
        {
            Name = name;
            Description = description;
            SourceAddresses = sourceAddresses;
            DestinationAddresses = destinationAddresses;
            DestinationPorts = destinationPorts;
            Protocols = protocols;
            TranslatedAddress = translatedAddress;
            TranslatedPort = translatedPort;
            TranslatedFqdn = translatedFqdn;
            SourceIPGroups = sourceIPGroups;
        }

        /// <summary> Name of the NAT rule. </summary>
        public string Name { get; set; }
        /// <summary> Description of the rule. </summary>
        public string Description { get; set; }
        /// <summary> List of source IP addresses for this rule. </summary>
        public IList<string> SourceAddresses { get; }
        /// <summary> List of destination IP addresses for this rule. Supports IP ranges, prefixes, and service tags. </summary>
        public IList<string> DestinationAddresses { get; }
        /// <summary> List of destination ports. </summary>
        public IList<string> DestinationPorts { get; }
        /// <summary> Array of AzureFirewallNetworkRuleProtocols applicable to this NAT rule. </summary>
        public IList<AzureFirewallNetworkRuleProtocol> Protocols { get; }
        /// <summary> The translated address for this NAT rule. </summary>
        public string TranslatedAddress { get; set; }
        /// <summary> The translated port for this NAT rule. </summary>
        public string TranslatedPort { get; set; }
        /// <summary> The translated FQDN for this NAT rule. </summary>
        public string TranslatedFqdn { get; set; }
        /// <summary> List of source IpGroups for this rule. </summary>
        public IList<string> SourceIPGroups { get; }
    }
}
