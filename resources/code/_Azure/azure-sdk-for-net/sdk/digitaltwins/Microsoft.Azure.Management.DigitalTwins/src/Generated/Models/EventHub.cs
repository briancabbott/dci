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
    using System.Linq;

    /// <summary>
    /// Properties related to EventHub.
    /// </summary>
    public partial class EventHub : DigitalTwinsEndpointResourceProperties
    {
        /// <summary>
        /// Initializes a new instance of the EventHub class.
        /// </summary>
        public EventHub()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the EventHub class.
        /// </summary>
        /// <param name="provisioningState">The provisioning state. Possible
        /// values include: 'Provisioning', 'Deleting', 'Updating',
        /// 'Succeeded', 'Failed', 'Canceled', 'Deleted', 'Warning',
        /// 'Suspending', 'Restoring', 'Moving', 'Disabled'</param>
        /// <param name="createdTime">Time when the Endpoint was added to
        /// DigitalTwinsInstance.</param>
        /// <param name="authenticationType">Specifies the authentication type
        /// being used for connecting to the endpoint. Defaults to 'KeyBased'.
        /// If 'KeyBased' is selected, a connection string must be specified
        /// (at least the primary connection string). If 'IdentityBased' is
        /// select, the endpointUri and entityPath properties must be
        /// specified. Possible values include: 'KeyBased',
        /// 'IdentityBased'</param>
        /// <param name="deadLetterSecret">Dead letter storage secret for
        /// key-based authentication. Will be obfuscated during read.</param>
        /// <param name="deadLetterUri">Dead letter storage URL for
        /// identity-based authentication.</param>
        /// <param name="connectionStringPrimaryKey">PrimaryConnectionString of
        /// the endpoint for key-based authentication. Will be obfuscated
        /// during read.</param>
        /// <param
        /// name="connectionStringSecondaryKey">SecondaryConnectionString of
        /// the endpoint for key-based authentication. Will be obfuscated
        /// during read.</param>
        /// <param name="endpointUri">The URL of the EventHub namespace for
        /// identity-based authentication. It must include the protocol
        /// 'sb://'.</param>
        /// <param name="entityPath">The EventHub name in the EventHub
        /// namespace for identity-based authentication.</param>
        public EventHub(string provisioningState = default(string), System.DateTime? createdTime = default(System.DateTime?), string authenticationType = default(string), string deadLetterSecret = default(string), string deadLetterUri = default(string), string connectionStringPrimaryKey = default(string), string connectionStringSecondaryKey = default(string), string endpointUri = default(string), string entityPath = default(string))
            : base(provisioningState, createdTime, authenticationType, deadLetterSecret, deadLetterUri)
        {
            ConnectionStringPrimaryKey = connectionStringPrimaryKey;
            ConnectionStringSecondaryKey = connectionStringSecondaryKey;
            EndpointUri = endpointUri;
            EntityPath = entityPath;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets primaryConnectionString of the endpoint for key-based
        /// authentication. Will be obfuscated during read.
        /// </summary>
        [JsonProperty(PropertyName = "connectionStringPrimaryKey")]
        public string ConnectionStringPrimaryKey { get; set; }

        /// <summary>
        /// Gets or sets secondaryConnectionString of the endpoint for
        /// key-based authentication. Will be obfuscated during read.
        /// </summary>
        [JsonProperty(PropertyName = "connectionStringSecondaryKey")]
        public string ConnectionStringSecondaryKey { get; set; }

        /// <summary>
        /// Gets or sets the URL of the EventHub namespace for identity-based
        /// authentication. It must include the protocol 'sb://'.
        /// </summary>
        [JsonProperty(PropertyName = "endpointUri")]
        public string EndpointUri { get; set; }

        /// <summary>
        /// Gets or sets the EventHub name in the EventHub namespace for
        /// identity-based authentication.
        /// </summary>
        [JsonProperty(PropertyName = "entityPath")]
        public string EntityPath { get; set; }

    }
}
