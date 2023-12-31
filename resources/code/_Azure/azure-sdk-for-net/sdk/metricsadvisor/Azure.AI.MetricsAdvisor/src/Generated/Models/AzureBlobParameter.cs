// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.AI.MetricsAdvisor.Models
{
    /// <summary> The AzureBlobParameter. </summary>
    internal partial class AzureBlobParameter
    {
        /// <summary> Initializes a new instance of <see cref="AzureBlobParameter"/>. </summary>
        /// <param name="container"> The container name in this Azure Blob. </param>
        /// <param name="blobTemplate"> The path template in this container. </param>
        public AzureBlobParameter(string container, string blobTemplate)
        {
            Container = container;
            BlobTemplate = blobTemplate;
        }

        /// <summary> Initializes a new instance of <see cref="AzureBlobParameter"/>. </summary>
        /// <param name="connectionString"> The connection string of this Azure Blob. </param>
        /// <param name="container"> The container name in this Azure Blob. </param>
        /// <param name="blobTemplate"> The path template in this container. </param>
        internal AzureBlobParameter(string connectionString, string container, string blobTemplate)
        {
            ConnectionString = connectionString;
            Container = container;
            BlobTemplate = blobTemplate;
        }

        /// <summary> The connection string of this Azure Blob. </summary>
        public string ConnectionString { get; set; }
        /// <summary> The container name in this Azure Blob. </summary>
        public string Container { get; set; }
        /// <summary> The path template in this container. </summary>
        public string BlobTemplate { get; set; }
    }
}
