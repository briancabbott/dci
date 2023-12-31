// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.HealthcareApis.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Authentication configuration information
    /// </summary>
    public partial class DicomServiceAuthenticationConfiguration
    {
        /// <summary>
        /// Initializes a new instance of the
        /// DicomServiceAuthenticationConfiguration class.
        /// </summary>
        public DicomServiceAuthenticationConfiguration()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the
        /// DicomServiceAuthenticationConfiguration class.
        /// </summary>
        /// <param name="authority">The authority url for the service</param>
        /// <param name="audiences">The audiences for the service</param>
        public DicomServiceAuthenticationConfiguration(string authority = default(string), IList<string> audiences = default(IList<string>))
        {
            Authority = authority;
            Audiences = audiences;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the authority url for the service
        /// </summary>
        [JsonProperty(PropertyName = "authority")]
        public string Authority { get; private set; }

        /// <summary>
        /// Gets the audiences for the service
        /// </summary>
        [JsonProperty(PropertyName = "audiences")]
        public IList<string> Audiences { get; private set; }

    }
}
