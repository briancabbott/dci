// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Compute.Models
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Microsoft.Rest.Serialization;
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// Describes a network interface reference.
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class NetworkInterfaceReference : IResource
    {
        /// <summary>
        /// Initializes a new instance of the NetworkInterfaceReference class.
        /// </summary>
        public NetworkInterfaceReference()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the NetworkInterfaceReference class.
        /// </summary>
        /// <param name="primary">Specifies the primary network interface in
        /// case the virtual machine has more than 1 network interface.</param>
        /// <param name="deleteOption">Specify what happens to the network
        /// interface when the VM is deleted. Possible values include:
        /// 'Delete', 'Detach'</param>
        public NetworkInterfaceReference(bool? primary = default(bool?), string deleteOption = default(string))
        {
            Primary = primary;
            DeleteOption = deleteOption;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets specifies the primary network interface in case the
        /// virtual machine has more than 1 network interface.
        /// </summary>
        [JsonProperty(PropertyName = "properties.primary")]
        public bool? Primary { get; set; }

        /// <summary>
        /// Gets or sets specify what happens to the network interface when the
        /// VM is deleted. Possible values include: 'Delete', 'Detach'
        /// </summary>
        [JsonProperty(PropertyName = "properties.deleteOption")]
        public string DeleteOption { get; set; }

    }
}
