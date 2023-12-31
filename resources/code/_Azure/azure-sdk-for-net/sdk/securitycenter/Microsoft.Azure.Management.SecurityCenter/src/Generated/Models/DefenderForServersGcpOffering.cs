// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Security.Models
{
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// The Defender for Servers GCP offering configurations
    /// </summary>
    [Newtonsoft.Json.JsonObject("DefenderForServersGcp")]
    public partial class DefenderForServersGcpOffering : CloudOffering
    {
        /// <summary>
        /// Initializes a new instance of the DefenderForServersGcpOffering
        /// class.
        /// </summary>
        public DefenderForServersGcpOffering()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the DefenderForServersGcpOffering
        /// class.
        /// </summary>
        /// <param name="description">The offering description.</param>
        /// <param name="defenderForServers">The Defender for servers
        /// connection configuration</param>
        /// <param name="arcAutoProvisioning">The ARC autoprovisioning
        /// configuration</param>
        /// <param name="vaAutoProvisioning">The Vulnerability Assessment
        /// autoprovisioning configuration</param>
        /// <param name="mdeAutoProvisioning">The Microsoft Defender for
        /// Endpoint autoprovisioning configuration</param>
        /// <param name="subPlan">configuration for the servers offering
        /// subPlan</param>
        public DefenderForServersGcpOffering(string description = default(string), DefenderForServersGcpOfferingDefenderForServers defenderForServers = default(DefenderForServersGcpOfferingDefenderForServers), DefenderForServersGcpOfferingArcAutoProvisioning arcAutoProvisioning = default(DefenderForServersGcpOfferingArcAutoProvisioning), DefenderForServersGcpOfferingVaAutoProvisioning vaAutoProvisioning = default(DefenderForServersGcpOfferingVaAutoProvisioning), DefenderForServersGcpOfferingMdeAutoProvisioning mdeAutoProvisioning = default(DefenderForServersGcpOfferingMdeAutoProvisioning), DefenderForServersGcpOfferingSubPlan subPlan = default(DefenderForServersGcpOfferingSubPlan))
            : base(description)
        {
            DefenderForServers = defenderForServers;
            ArcAutoProvisioning = arcAutoProvisioning;
            VaAutoProvisioning = vaAutoProvisioning;
            MdeAutoProvisioning = mdeAutoProvisioning;
            SubPlan = subPlan;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the Defender for servers connection configuration
        /// </summary>
        [JsonProperty(PropertyName = "defenderForServers")]
        public DefenderForServersGcpOfferingDefenderForServers DefenderForServers { get; set; }

        /// <summary>
        /// Gets or sets the ARC autoprovisioning configuration
        /// </summary>
        [JsonProperty(PropertyName = "arcAutoProvisioning")]
        public DefenderForServersGcpOfferingArcAutoProvisioning ArcAutoProvisioning { get; set; }

        /// <summary>
        /// Gets or sets the Vulnerability Assessment autoprovisioning
        /// configuration
        /// </summary>
        [JsonProperty(PropertyName = "vaAutoProvisioning")]
        public DefenderForServersGcpOfferingVaAutoProvisioning VaAutoProvisioning { get; set; }

        /// <summary>
        /// Gets or sets the Microsoft Defender for Endpoint autoprovisioning
        /// configuration
        /// </summary>
        [JsonProperty(PropertyName = "mdeAutoProvisioning")]
        public DefenderForServersGcpOfferingMdeAutoProvisioning MdeAutoProvisioning { get; set; }

        /// <summary>
        /// Gets or sets configuration for the servers offering subPlan
        /// </summary>
        [JsonProperty(PropertyName = "subPlan")]
        public DefenderForServersGcpOfferingSubPlan SubPlan { get; set; }

    }
}
