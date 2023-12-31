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
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// OAuth2 settings details
    /// </summary>
    public partial class AuthorizationProviderOAuth2Settings
    {
        /// <summary>
        /// Initializes a new instance of the
        /// AuthorizationProviderOAuth2Settings class.
        /// </summary>
        public AuthorizationProviderOAuth2Settings()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the
        /// AuthorizationProviderOAuth2Settings class.
        /// </summary>
        /// <param name="redirectUrl">Redirect URL to be set in the OAuth
        /// application.</param>
        /// <param name="grantTypes">OAuth2 settings</param>
        public AuthorizationProviderOAuth2Settings(string redirectUrl = default(string), AuthorizationProviderOAuth2GrantTypes grantTypes = default(AuthorizationProviderOAuth2GrantTypes))
        {
            RedirectUrl = redirectUrl;
            GrantTypes = grantTypes;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets redirect URL to be set in the OAuth application.
        /// </summary>
        [JsonProperty(PropertyName = "redirectUrl")]
        public string RedirectUrl { get; set; }

        /// <summary>
        /// Gets or sets oAuth2 settings
        /// </summary>
        [JsonProperty(PropertyName = "grantTypes")]
        public AuthorizationProviderOAuth2GrantTypes GrantTypes { get; set; }

    }
}
