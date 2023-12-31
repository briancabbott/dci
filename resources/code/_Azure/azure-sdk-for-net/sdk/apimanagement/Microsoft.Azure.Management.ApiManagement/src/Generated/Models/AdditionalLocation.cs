// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.ApiManagement.Models
{
    using Microsoft.Rest;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Description of an additional API Management resource location.
    /// </summary>
    public partial class AdditionalLocation
    {
        /// <summary>
        /// Initializes a new instance of the AdditionalLocation class.
        /// </summary>
        public AdditionalLocation()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the AdditionalLocation class.
        /// </summary>
        /// <param name="location">The location name of the additional region
        /// among Azure Data center regions.</param>
        /// <param name="sku">SKU properties of the API Management
        /// service.</param>
        /// <param name="zones">A list of availability zones denoting where the
        /// resource needs to come from.</param>
        /// <param name="publicIPAddresses">Public Static Load Balanced IP
        /// addresses of the API Management service in the additional location.
        /// Available only for Basic, Standard, Premium and Isolated
        /// SKU.</param>
        /// <param name="privateIPAddresses">Private Static Load Balanced IP
        /// addresses of the API Management service which is deployed in an
        /// Internal Virtual Network in a particular additional location.
        /// Available only for Basic, Standard, Premium and Isolated
        /// SKU.</param>
        /// <param name="publicIpAddressId">Public Standard SKU IP V4 based IP
        /// address to be associated with Virtual Network deployed service in
        /// the location. Supported only for Premium SKU being deployed in
        /// Virtual Network.</param>
        /// <param name="virtualNetworkConfiguration">Virtual network
        /// configuration for the location.</param>
        /// <param name="gatewayRegionalUrl">Gateway URL of the API Management
        /// service in the Region.</param>
        /// <param name="natGatewayState">Property can be used to enable NAT
        /// Gateway for this API Management service. Possible values include:
        /// 'Enabled', 'Disabled'</param>
        /// <param name="outboundPublicIPAddresses">Outbound public IPV4
        /// address prefixes associated with NAT Gateway deployed service.
        /// Available only for Premium SKU on stv2 platform.</param>
        /// <param name="disableGateway">Property only valid for an Api
        /// Management service deployed in multiple locations. This can be used
        /// to disable the gateway in this additional location.</param>
        /// <param name="platformVersion">Compute Platform Version running the
        /// service. Possible values include: 'undetermined', 'stv1', 'stv2',
        /// 'mtv1'</param>
        public AdditionalLocation(string location, ApiManagementServiceSkuProperties sku, IList<string> zones = default(IList<string>), IList<string> publicIPAddresses = default(IList<string>), IList<string> privateIPAddresses = default(IList<string>), string publicIpAddressId = default(string), VirtualNetworkConfiguration virtualNetworkConfiguration = default(VirtualNetworkConfiguration), string gatewayRegionalUrl = default(string), string natGatewayState = default(string), IList<string> outboundPublicIPAddresses = default(IList<string>), bool? disableGateway = default(bool?), string platformVersion = default(string))
        {
            Location = location;
            Sku = sku;
            Zones = zones;
            PublicIPAddresses = publicIPAddresses;
            PrivateIPAddresses = privateIPAddresses;
            PublicIpAddressId = publicIpAddressId;
            VirtualNetworkConfiguration = virtualNetworkConfiguration;
            GatewayRegionalUrl = gatewayRegionalUrl;
            NatGatewayState = natGatewayState;
            OutboundPublicIPAddresses = outboundPublicIPAddresses;
            DisableGateway = disableGateway;
            PlatformVersion = platformVersion;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the location name of the additional region among Azure
        /// Data center regions.
        /// </summary>
        [JsonProperty(PropertyName = "location")]
        public string Location { get; set; }

        /// <summary>
        /// Gets or sets SKU properties of the API Management service.
        /// </summary>
        [JsonProperty(PropertyName = "sku")]
        public ApiManagementServiceSkuProperties Sku { get; set; }

        /// <summary>
        /// Gets or sets a list of availability zones denoting where the
        /// resource needs to come from.
        /// </summary>
        [JsonProperty(PropertyName = "zones")]
        public IList<string> Zones { get; set; }

        /// <summary>
        /// Gets public Static Load Balanced IP addresses of the API Management
        /// service in the additional location. Available only for Basic,
        /// Standard, Premium and Isolated SKU.
        /// </summary>
        [JsonProperty(PropertyName = "publicIPAddresses")]
        public IList<string> PublicIPAddresses { get; private set; }

        /// <summary>
        /// Gets private Static Load Balanced IP addresses of the API
        /// Management service which is deployed in an Internal Virtual Network
        /// in a particular additional location. Available only for Basic,
        /// Standard, Premium and Isolated SKU.
        /// </summary>
        [JsonProperty(PropertyName = "privateIPAddresses")]
        public IList<string> PrivateIPAddresses { get; private set; }

        /// <summary>
        /// Gets or sets public Standard SKU IP V4 based IP address to be
        /// associated with Virtual Network deployed service in the location.
        /// Supported only for Premium SKU being deployed in Virtual Network.
        /// </summary>
        [JsonProperty(PropertyName = "publicIpAddressId")]
        public string PublicIpAddressId { get; set; }

        /// <summary>
        /// Gets or sets virtual network configuration for the location.
        /// </summary>
        [JsonProperty(PropertyName = "virtualNetworkConfiguration")]
        public VirtualNetworkConfiguration VirtualNetworkConfiguration { get; set; }

        /// <summary>
        /// Gets gateway URL of the API Management service in the Region.
        /// </summary>
        [JsonProperty(PropertyName = "gatewayRegionalUrl")]
        public string GatewayRegionalUrl { get; private set; }

        /// <summary>
        /// Gets or sets property can be used to enable NAT Gateway for this
        /// API Management service. Possible values include: 'Enabled',
        /// 'Disabled'
        /// </summary>
        [JsonProperty(PropertyName = "natGatewayState")]
        public string NatGatewayState { get; set; }

        /// <summary>
        /// Gets outbound public IPV4 address prefixes associated with NAT
        /// Gateway deployed service. Available only for Premium SKU on stv2
        /// platform.
        /// </summary>
        [JsonProperty(PropertyName = "outboundPublicIPAddresses")]
        public IList<string> OutboundPublicIPAddresses { get; private set; }

        /// <summary>
        /// Gets or sets property only valid for an Api Management service
        /// deployed in multiple locations. This can be used to disable the
        /// gateway in this additional location.
        /// </summary>
        [JsonProperty(PropertyName = "disableGateway")]
        public bool? DisableGateway { get; set; }

        /// <summary>
        /// Gets compute Platform Version running the service. Possible values
        /// include: 'undetermined', 'stv1', 'stv2', 'mtv1'
        /// </summary>
        [JsonProperty(PropertyName = "platformVersion")]
        public string PlatformVersion { get; private set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (Location == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "Location");
            }
            if (Sku == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "Sku");
            }
            if (Sku != null)
            {
                Sku.Validate();
            }
            if (VirtualNetworkConfiguration != null)
            {
                VirtualNetworkConfiguration.Validate();
            }
        }
    }
}
