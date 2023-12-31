// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Analytics.Synapse.Artifacts.Models
{
    /// <summary> The MetastoreRequestSuccessResponse. </summary>
    public partial class MetastoreRequestSuccessResponse
    {
        /// <summary> Initializes a new instance of <see cref="MetastoreRequestSuccessResponse"/>. </summary>
        internal MetastoreRequestSuccessResponse()
        {
        }

        /// <summary> Initializes a new instance of <see cref="MetastoreRequestSuccessResponse"/>. </summary>
        /// <param name="status"> Enumerates possible Status of the resource. </param>
        internal MetastoreRequestSuccessResponse(ResourceStatus? status)
        {
            Status = status;
        }

        /// <summary> Enumerates possible Status of the resource. </summary>
        public ResourceStatus? Status { get; }
    }
}
