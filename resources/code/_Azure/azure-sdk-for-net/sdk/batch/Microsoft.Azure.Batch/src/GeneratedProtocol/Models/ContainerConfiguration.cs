// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Batch.Protocol.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The configuration for container-enabled Pools.
    /// </summary>
    public partial class ContainerConfiguration
    {
        /// <summary>
        /// Initializes a new instance of the ContainerConfiguration class.
        /// </summary>
        public ContainerConfiguration()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the ContainerConfiguration class.
        /// </summary>
        /// <param name="type">The container technology to be used.</param>
        /// <param name="containerImageNames">The collection of container Image
        /// names.</param>
        /// <param name="containerRegistries">Additional private registries
        /// from which containers can be pulled.</param>
        public ContainerConfiguration(string type, IList<string> containerImageNames = default(IList<string>), IList<ContainerRegistry> containerRegistries = default(IList<ContainerRegistry>))
        {
            Type = type;
            ContainerImageNames = containerImageNames;
            ContainerRegistries = containerRegistries;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the container technology to be used.
        /// </summary>
        /// <remarks>
        /// Possible values include: 'dockerCompatible', 'criCompatible'
        /// </remarks>
        [JsonProperty(PropertyName = "type")]
        public string Type { get; set; }

        /// <summary>
        /// Gets or sets the collection of container Image names.
        /// </summary>
        /// <remarks>
        /// This is the full Image reference, as would be specified to "docker
        /// pull". An Image will be sourced from the default Docker registry
        /// unless the Image is fully qualified with an alternative registry.
        /// </remarks>
        [JsonProperty(PropertyName = "containerImageNames")]
        public IList<string> ContainerImageNames { get; set; }

        /// <summary>
        /// Gets or sets additional private registries from which containers
        /// can be pulled.
        /// </summary>
        /// <remarks>
        /// If any Images must be downloaded from a private registry which
        /// requires credentials, then those credentials must be provided here.
        /// </remarks>
        [JsonProperty(PropertyName = "containerRegistries")]
        public IList<ContainerRegistry> ContainerRegistries { get; set; }

    }
}
