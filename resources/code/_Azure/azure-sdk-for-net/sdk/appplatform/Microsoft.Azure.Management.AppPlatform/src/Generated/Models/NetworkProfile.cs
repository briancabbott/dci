// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.AppPlatform.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Service network profile payload
    /// </summary>
    public partial class NetworkProfile
    {
        /// <summary>
        /// Initializes a new instance of the NetworkProfile class.
        /// </summary>
        public NetworkProfile()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the NetworkProfile class.
        /// </summary>
        /// <param name="serviceRuntimeSubnetId">Fully qualified resource Id of
        /// the subnet to host Azure Spring Apps Service Runtime</param>
        /// <param name="appSubnetId">Fully qualified resource Id of the subnet
        /// to host customer apps in Azure Spring Apps</param>
        /// <param name="serviceCidr">Azure Spring Apps service reserved
        /// CIDR</param>
        /// <param name="serviceRuntimeNetworkResourceGroup">Name of the
        /// resource group containing network resources of Azure Spring Apps
        /// Service Runtime</param>
        /// <param name="appNetworkResourceGroup">Name of the resource group
        /// containing network resources for customer apps in Azure Spring
        /// Apps</param>
        /// <param name="outboundIPs">Desired outbound IP resources for Azure
        /// Spring Apps resource.</param>
        /// <param name="requiredTraffics">Required inbound or outbound
        /// traffics for Azure Spring Apps resource.</param>
        /// <param name="ingressConfig">Ingress configuration payload for Azure
        /// Spring Apps resource.</param>
        public NetworkProfile(string serviceRuntimeSubnetId = default(string), string appSubnetId = default(string), string serviceCidr = default(string), string serviceRuntimeNetworkResourceGroup = default(string), string appNetworkResourceGroup = default(string), NetworkProfileOutboundIPs outboundIPs = default(NetworkProfileOutboundIPs), IList<RequiredTraffic> requiredTraffics = default(IList<RequiredTraffic>), IngressConfig ingressConfig = default(IngressConfig))
        {
            ServiceRuntimeSubnetId = serviceRuntimeSubnetId;
            AppSubnetId = appSubnetId;
            ServiceCidr = serviceCidr;
            ServiceRuntimeNetworkResourceGroup = serviceRuntimeNetworkResourceGroup;
            AppNetworkResourceGroup = appNetworkResourceGroup;
            OutboundIPs = outboundIPs;
            RequiredTraffics = requiredTraffics;
            IngressConfig = ingressConfig;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets fully qualified resource Id of the subnet to host
        /// Azure Spring Apps Service Runtime
        /// </summary>
        [JsonProperty(PropertyName = "serviceRuntimeSubnetId")]
        public string ServiceRuntimeSubnetId { get; set; }

        /// <summary>
        /// Gets or sets fully qualified resource Id of the subnet to host
        /// customer apps in Azure Spring Apps
        /// </summary>
        [JsonProperty(PropertyName = "appSubnetId")]
        public string AppSubnetId { get; set; }

        /// <summary>
        /// Gets or sets azure Spring Apps service reserved CIDR
        /// </summary>
        [JsonProperty(PropertyName = "serviceCidr")]
        public string ServiceCidr { get; set; }

        /// <summary>
        /// Gets or sets name of the resource group containing network
        /// resources of Azure Spring Apps Service Runtime
        /// </summary>
        [JsonProperty(PropertyName = "serviceRuntimeNetworkResourceGroup")]
        public string ServiceRuntimeNetworkResourceGroup { get; set; }

        /// <summary>
        /// Gets or sets name of the resource group containing network
        /// resources for customer apps in Azure Spring Apps
        /// </summary>
        [JsonProperty(PropertyName = "appNetworkResourceGroup")]
        public string AppNetworkResourceGroup { get; set; }

        /// <summary>
        /// Gets desired outbound IP resources for Azure Spring Apps resource.
        /// </summary>
        [JsonProperty(PropertyName = "outboundIPs")]
        public NetworkProfileOutboundIPs OutboundIPs { get; private set; }

        /// <summary>
        /// Gets required inbound or outbound traffics for Azure Spring Apps
        /// resource.
        /// </summary>
        [JsonProperty(PropertyName = "requiredTraffics")]
        public IList<RequiredTraffic> RequiredTraffics { get; private set; }

        /// <summary>
        /// Gets or sets ingress configuration payload for Azure Spring Apps
        /// resource.
        /// </summary>
        [JsonProperty(PropertyName = "ingressConfig")]
        public IngressConfig IngressConfig { get; set; }

    }
}
