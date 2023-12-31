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
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The List VmImages in EdgeZone operation response.
    /// </summary>
    public partial class VmImagesInEdgeZoneListResult
    {
        /// <summary>
        /// Initializes a new instance of the VmImagesInEdgeZoneListResult
        /// class.
        /// </summary>
        public VmImagesInEdgeZoneListResult()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the VmImagesInEdgeZoneListResult
        /// class.
        /// </summary>
        /// <param name="value">The list of VMImages in EdgeZone</param>
        /// <param name="nextLink">The URI to fetch the next page of VMImages
        /// in EdgeZone. Call ListNext() with this URI to fetch the next page
        /// of VmImages.</param>
        public VmImagesInEdgeZoneListResult(IList<VirtualMachineImageResource> value = default(IList<VirtualMachineImageResource>), string nextLink = default(string))
        {
            Value = value;
            NextLink = nextLink;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the list of VMImages in EdgeZone
        /// </summary>
        [JsonProperty(PropertyName = "value")]
        public IList<VirtualMachineImageResource> Value { get; set; }

        /// <summary>
        /// Gets or sets the URI to fetch the next page of VMImages in
        /// EdgeZone. Call ListNext() with this URI to fetch the next page of
        /// VmImages.
        /// </summary>
        [JsonProperty(PropertyName = "nextLink")]
        public string NextLink { get; set; }

    }
}
