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
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The common properties of a DigitalTwinsInstance.
    /// </summary>
    public partial class DigitalTwinsResource : IResource
    {
        /// <summary>
        /// Initializes a new instance of the DigitalTwinsResource class.
        /// </summary>
        public DigitalTwinsResource()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the DigitalTwinsResource class.
        /// </summary>
        /// <param name="location">The resource location.</param>
        /// <param name="id">The resource identifier.</param>
        /// <param name="name">The resource name.</param>
        /// <param name="type">The resource type.</param>
        /// <param name="tags">The resource tags.</param>
        /// <param name="identity">The managed identity for the
        /// DigitalTwinsInstance.</param>
        /// <param name="systemData">Metadata pertaining to creation and last
        /// modification of the DigitalTwinsInstance.</param>
        public DigitalTwinsResource(string location, string id = default(string), string name = default(string), string type = default(string), IDictionary<string, string> tags = default(IDictionary<string, string>), DigitalTwinsIdentity identity = default(DigitalTwinsIdentity), SystemData systemData = default(SystemData))
        {
            Id = id;
            Name = name;
            Type = type;
            Location = location;
            Tags = tags;
            Identity = identity;
            SystemData = systemData;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the resource identifier.
        /// </summary>
        [JsonProperty(PropertyName = "id")]
        public string Id { get; private set; }

        /// <summary>
        /// Gets the resource name.
        /// </summary>
        [JsonProperty(PropertyName = "name")]
        public string Name { get; private set; }

        /// <summary>
        /// Gets the resource type.
        /// </summary>
        [JsonProperty(PropertyName = "type")]
        public string Type { get; private set; }

        /// <summary>
        /// Gets or sets the resource location.
        /// </summary>
        [JsonProperty(PropertyName = "location")]
        public string Location { get; set; }

        /// <summary>
        /// Gets or sets the resource tags.
        /// </summary>
        [JsonProperty(PropertyName = "tags")]
        public IDictionary<string, string> Tags { get; set; }

        /// <summary>
        /// Gets or sets the managed identity for the DigitalTwinsInstance.
        /// </summary>
        [JsonProperty(PropertyName = "identity")]
        public DigitalTwinsIdentity Identity { get; set; }

        /// <summary>
        /// Gets metadata pertaining to creation and last modification of the
        /// DigitalTwinsInstance.
        /// </summary>
        [JsonProperty(PropertyName = "systemData")]
        public SystemData SystemData { get; private set; }

    }
}
