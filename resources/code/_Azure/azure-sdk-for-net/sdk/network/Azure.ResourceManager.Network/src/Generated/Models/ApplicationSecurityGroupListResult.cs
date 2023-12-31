// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Network;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> A list of application security groups. </summary>
    internal partial class ApplicationSecurityGroupListResult
    {
        /// <summary> Initializes a new instance of <see cref="ApplicationSecurityGroupListResult"/>. </summary>
        internal ApplicationSecurityGroupListResult()
        {
            Value = new ChangeTrackingList<ApplicationSecurityGroupData>();
        }

        /// <summary> Initializes a new instance of <see cref="ApplicationSecurityGroupListResult"/>. </summary>
        /// <param name="value"> A list of application security groups. </param>
        /// <param name="nextLink"> The URL to get the next set of results. </param>
        internal ApplicationSecurityGroupListResult(IReadOnlyList<ApplicationSecurityGroupData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> A list of application security groups. </summary>
        public IReadOnlyList<ApplicationSecurityGroupData> Value { get; }
        /// <summary> The URL to get the next set of results. </summary>
        public string NextLink { get; }
    }
}
