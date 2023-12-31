// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure;

namespace Azure.ResourceManager.Resources.Models
{
    /// <summary> Operation status message object. </summary>
    public partial class StatusMessage
    {
        /// <summary> Initializes a new instance of <see cref="StatusMessage"/>. </summary>
        internal StatusMessage()
        {
        }

        /// <summary> Initializes a new instance of <see cref="StatusMessage"/>. </summary>
        /// <param name="status"> Status of the deployment operation. </param>
        /// <param name="error"> The error reported by the operation. </param>
        internal StatusMessage(string status, ResponseError error)
        {
            Status = status;
            Error = error;
        }

        /// <summary> Status of the deployment operation. </summary>
        public string Status { get; }
        /// <summary> The error reported by the operation. </summary>
        public ResponseError Error { get; }
    }
}
