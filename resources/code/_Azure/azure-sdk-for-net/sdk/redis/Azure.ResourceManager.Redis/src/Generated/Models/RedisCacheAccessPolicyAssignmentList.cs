// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Redis;

namespace Azure.ResourceManager.Redis.Models
{
    /// <summary> List of access policies assignments (with properties) of a Redis cache. </summary>
    internal partial class RedisCacheAccessPolicyAssignmentList
    {
        /// <summary> Initializes a new instance of <see cref="RedisCacheAccessPolicyAssignmentList"/>. </summary>
        internal RedisCacheAccessPolicyAssignmentList()
        {
            Value = new ChangeTrackingList<RedisCacheAccessPolicyAssignmentData>();
        }

        /// <summary> Initializes a new instance of <see cref="RedisCacheAccessPolicyAssignmentList"/>. </summary>
        /// <param name="value"> List of access policies assignments (with properties) of a Redis cache. </param>
        /// <param name="nextLink"> Link for next set. </param>
        internal RedisCacheAccessPolicyAssignmentList(IReadOnlyList<RedisCacheAccessPolicyAssignmentData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> List of access policies assignments (with properties) of a Redis cache. </summary>
        public IReadOnlyList<RedisCacheAccessPolicyAssignmentData> Value { get; }
        /// <summary> Link for next set. </summary>
        public string NextLink { get; }
    }
}
