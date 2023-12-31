// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.DigitalTwins.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The properties of a private endpoint connection.
    /// </summary>
    public partial class ConnectionProperties
    {
        /// <summary>
        /// Initializes a new instance of the ConnectionProperties class.
        /// </summary>
        public ConnectionProperties()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the ConnectionProperties class.
        /// </summary>
        /// <param name="provisioningState">The provisioning state. Possible
        /// values include: 'Pending', 'Approved', 'Rejected',
        /// 'Disconnected'</param>
        /// <param name="privateEndpoint">The private endpoint.</param>
        /// <param name="groupIds">The list of group ids for the private
        /// endpoint connection.</param>
        /// <param name="privateLinkServiceConnectionState">The connection
        /// state.</param>
        public ConnectionProperties(string provisioningState = default(string), PrivateEndpoint privateEndpoint = default(PrivateEndpoint), IList<string> groupIds = default(IList<string>), ConnectionPropertiesPrivateLinkServiceConnectionState privateLinkServiceConnectionState = default(ConnectionPropertiesPrivateLinkServiceConnectionState))
        {
            ProvisioningState = provisioningState;
            PrivateEndpoint = privateEndpoint;
            GroupIds = groupIds;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the provisioning state. Possible values include: 'Pending',
        /// 'Approved', 'Rejected', 'Disconnected'
        /// </summary>
        [JsonProperty(PropertyName = "provisioningState")]
        public string ProvisioningState { get; private set; }

        /// <summary>
        /// Gets or sets the private endpoint.
        /// </summary>
        [JsonProperty(PropertyName = "privateEndpoint")]
        public PrivateEndpoint PrivateEndpoint { get; set; }

        /// <summary>
        /// Gets or sets the list of group ids for the private endpoint
        /// connection.
        /// </summary>
        [JsonProperty(PropertyName = "groupIds")]
        public IList<string> GroupIds { get; set; }

        /// <summary>
        /// Gets or sets the connection state.
        /// </summary>
        [JsonProperty(PropertyName = "privateLinkServiceConnectionState")]
        public ConnectionPropertiesPrivateLinkServiceConnectionState PrivateLinkServiceConnectionState { get; set; }

    }
}
