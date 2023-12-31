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
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// A list of AzureVMSS instances which can be included or excluded to run
    /// packet capture. If both included and excluded are empty, then the
    /// packet capture will run on all instances of AzureVMSS.
    /// </summary>
    public partial class PacketCaptureMachineScope
    {
        /// <summary>
        /// Initializes a new instance of the PacketCaptureMachineScope class.
        /// </summary>
        public PacketCaptureMachineScope()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the PacketCaptureMachineScope class.
        /// </summary>
        /// <param name="include">List of AzureVMSS instances to run packet
        /// capture on.</param>
        /// <param name="exclude">List of AzureVMSS instances which has to be
        /// excluded from the AzureVMSS from running packet capture.</param>
        public PacketCaptureMachineScope(IList<string> include = default(IList<string>), IList<string> exclude = default(IList<string>))
        {
            Include = include;
            Exclude = exclude;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets list of AzureVMSS instances to run packet capture on.
        /// </summary>
        [JsonProperty(PropertyName = "include")]
        public IList<string> Include { get; set; }

        /// <summary>
        /// Gets or sets list of AzureVMSS instances which has to be excluded
        /// from the AzureVMSS from running packet capture.
        /// </summary>
        [JsonProperty(PropertyName = "exclude")]
        public IList<string> Exclude { get; set; }

    }
}
