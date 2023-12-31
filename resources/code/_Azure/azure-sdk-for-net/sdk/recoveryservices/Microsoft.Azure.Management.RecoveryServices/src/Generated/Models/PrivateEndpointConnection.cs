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
    /// Private Endpoint Connection Response Properties.
    /// </summary>
    public partial class PrivateEndpointConnection
    {
        /// <summary>
        /// Initializes a new instance of the PrivateEndpointConnection class.
        /// </summary>
        public PrivateEndpointConnection()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the PrivateEndpointConnection class.
        /// </summary>
        /// <param name="provisioningState">Gets or sets provisioning state of
        /// the private endpoint connection. Possible values include:
        /// 'Succeeded', 'Deleting', 'Failed', 'Pending'</param>
        /// <param name="groupIds">Group Ids for the Private Endpoint</param>
        public PrivateEndpointConnection(string provisioningState = default(string), PrivateEndpoint privateEndpoint = default(PrivateEndpoint), PrivateLinkServiceConnectionState privateLinkServiceConnectionState = default(PrivateLinkServiceConnectionState), IList<string> groupIds = default(IList<string>))
        {
            ProvisioningState = provisioningState;
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            GroupIds = groupIds;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets provisioning state of the private endpoint connection.
        /// Possible values include: 'Succeeded', 'Deleting', 'Failed',
        /// 'Pending'
        /// </summary>
        [JsonProperty(PropertyName = "provisioningState")]
        public string ProvisioningState { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "privateEndpoint")]
        public PrivateEndpoint PrivateEndpoint { get; set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "privateLinkServiceConnectionState")]
        public PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Gets or sets group Ids for the Private Endpoint
        /// </summary>
        [JsonProperty(PropertyName = "groupIds")]
        public IList<string> GroupIds { get; set; }

    }
}
