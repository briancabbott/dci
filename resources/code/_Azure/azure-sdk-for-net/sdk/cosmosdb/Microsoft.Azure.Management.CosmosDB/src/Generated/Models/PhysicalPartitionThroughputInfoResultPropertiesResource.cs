// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.CosmosDB.Models
{
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// properties of physical partition throughput info
    /// </summary>
    public partial class PhysicalPartitionThroughputInfoResultPropertiesResource : PhysicalPartitionThroughputInfoProperties
    {
        /// <summary>
        /// Initializes a new instance of the
        /// PhysicalPartitionThroughputInfoResultPropertiesResource class.
        /// </summary>
        public PhysicalPartitionThroughputInfoResultPropertiesResource()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the
        /// PhysicalPartitionThroughputInfoResultPropertiesResource class.
        /// </summary>
        /// <param name="physicalPartitionThroughputInfo">Array of physical
        /// partition throughput info objects</param>
        public PhysicalPartitionThroughputInfoResultPropertiesResource(IList<PhysicalPartitionThroughputInfoResource> physicalPartitionThroughputInfo = default(IList<PhysicalPartitionThroughputInfoResource>))
            : base(physicalPartitionThroughputInfo)
        {
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

    }
}
