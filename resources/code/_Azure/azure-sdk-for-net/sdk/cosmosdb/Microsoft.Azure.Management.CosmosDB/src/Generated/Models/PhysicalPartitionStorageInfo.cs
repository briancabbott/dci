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
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// The storage of a physical partition
    /// </summary>
    public partial class PhysicalPartitionStorageInfo
    {
        /// <summary>
        /// Initializes a new instance of the PhysicalPartitionStorageInfo
        /// class.
        /// </summary>
        public PhysicalPartitionStorageInfo()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the PhysicalPartitionStorageInfo
        /// class.
        /// </summary>
        /// <param name="id">The unique identifier of the partition.</param>
        /// <param name="storageInKB">The storage in KB for the physical
        /// partition.</param>
        public PhysicalPartitionStorageInfo(string id = default(string), double? storageInKB = default(double?))
        {
            Id = id;
            StorageInKB = storageInKB;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the unique identifier of the partition.
        /// </summary>
        [JsonProperty(PropertyName = "id")]
        public string Id { get; private set; }

        /// <summary>
        /// Gets the storage in KB for the physical partition.
        /// </summary>
        [JsonProperty(PropertyName = "storageInKB")]
        public double? StorageInKB { get; private set; }

    }
}
