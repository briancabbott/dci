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
    using Microsoft.Rest.Serialization;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Authorization contract.
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class AuthorizationContract : ProxyResource
    {
        /// <summary>
        /// Initializes a new instance of the AuthorizationContract class.
        /// </summary>
        public AuthorizationContract()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the AuthorizationContract class.
        /// </summary>
        /// <param name="id">Fully qualified resource ID for the resource. Ex -
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</param>
        /// <param name="name">The name of the resource</param>
        /// <param name="type">The type of the resource. E.g.
        /// "Microsoft.Compute/virtualMachines" or
        /// "Microsoft.Storage/storageAccounts"</param>
        /// <param name="authorizationType">Authorization type options.
        /// Possible values include: 'OAuth2'</param>
        /// <param name="oAuth2GrantType">OAuth2 grant type options. Possible
        /// values include: 'AuthorizationCode', 'ClientCredentials'</param>
        /// <param name="parameters">Authorization parameters</param>
        /// <param name="status">Status of the Authorization</param>
        public AuthorizationContract(string id = default(string), string name = default(string), string type = default(string), string authorizationType = default(string), string oAuth2GrantType = default(string), IDictionary<string, string> parameters = default(IDictionary<string, string>), AuthorizationError error = default(AuthorizationError), string status = default(string))
            : base(id, name, type)
        {
            AuthorizationType = authorizationType;
            OAuth2GrantType = oAuth2GrantType;
            Parameters = parameters;
            Error = error;
            Status = status;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets authorization type options. Possible values include:
        /// 'OAuth2'
        /// </summary>
        [JsonProperty(PropertyName = "properties.authorizationType")]
        public string AuthorizationType { get; set; }

        /// <summary>
        /// Gets or sets oAuth2 grant type options. Possible values include:
        /// 'AuthorizationCode', 'ClientCredentials'
        /// </summary>
        [JsonProperty(PropertyName = "properties.oauth2grantType")]
        public string OAuth2GrantType { get; set; }

        /// <summary>
        /// Gets or sets authorization parameters
        /// </summary>
        [JsonProperty(PropertyName = "properties.parameters")]
        public IDictionary<string, string> Parameters { get; set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "properties.error")]
        public AuthorizationError Error { get; set; }

        /// <summary>
        /// Gets or sets status of the Authorization
        /// </summary>
        [JsonProperty(PropertyName = "properties.status")]
        public string Status { get; set; }

    }
}
