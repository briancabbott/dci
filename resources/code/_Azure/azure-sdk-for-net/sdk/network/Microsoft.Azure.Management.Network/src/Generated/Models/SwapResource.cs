// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Network.Models
{
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// SwapResource to represent slot type on the specified cloud service.
    /// </summary>
    public partial class SwapResource
    {
        /// <summary>
        /// Initializes a new instance of the SwapResource class.
        /// </summary>
        public SwapResource()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the SwapResource class.
        /// </summary>
        /// <param name="id">Resource Id.</param>
        /// <param name="name">Resource name.</param>
        /// <param name="type">Resource type.</param>
        public SwapResource(string id = default(string), string name = default(string), string type = default(string), SwapResourceProperties properties = default(SwapResourceProperties))
        {
            Id = id;
            Name = name;
            Type = type;
            Properties = properties;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets resource Id.
        /// </summary>
        [JsonProperty(PropertyName = "id")]
        public string Id { get; private set; }

        /// <summary>
        /// Gets resource name.
        /// </summary>
        [JsonProperty(PropertyName = "name")]
        public string Name { get; private set; }

        /// <summary>
        /// Gets resource type.
        /// </summary>
        [JsonProperty(PropertyName = "type")]
        public string Type { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "properties")]
        public SwapResourceProperties Properties { get; set; }

    }
}
