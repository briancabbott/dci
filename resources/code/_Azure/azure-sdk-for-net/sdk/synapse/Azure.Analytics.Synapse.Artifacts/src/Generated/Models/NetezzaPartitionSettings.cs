// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Analytics.Synapse.Artifacts.Models
{
    /// <summary> The settings that will be leveraged for Netezza source partitioning. </summary>
    public partial class NetezzaPartitionSettings
    {
        /// <summary> Initializes a new instance of <see cref="NetezzaPartitionSettings"/>. </summary>
        public NetezzaPartitionSettings()
        {
        }

        /// <summary> Initializes a new instance of <see cref="NetezzaPartitionSettings"/>. </summary>
        /// <param name="partitionColumnName"> The name of the column in integer type that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </param>
        /// <param name="partitionUpperBound"> The maximum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </param>
        /// <param name="partitionLowerBound"> The minimum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </param>
        internal NetezzaPartitionSettings(object partitionColumnName, object partitionUpperBound, object partitionLowerBound)
        {
            PartitionColumnName = partitionColumnName;
            PartitionUpperBound = partitionUpperBound;
            PartitionLowerBound = partitionLowerBound;
        }

        /// <summary> The name of the column in integer type that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </summary>
        public object PartitionColumnName { get; set; }
        /// <summary> The maximum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </summary>
        public object PartitionUpperBound { get; set; }
        /// <summary> The minimum value of column specified in partitionColumnName that will be used for proceeding range partitioning. Type: string (or Expression with resultType string). </summary>
        public object PartitionLowerBound { get; set; }
    }
}
