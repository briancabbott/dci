// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Sql;

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> A list of short term retention policies. </summary>
    internal partial class ManagedBackupShortTermRetentionPolicyListResult
    {
        /// <summary> Initializes a new instance of <see cref="ManagedBackupShortTermRetentionPolicyListResult"/>. </summary>
        internal ManagedBackupShortTermRetentionPolicyListResult()
        {
            Value = new ChangeTrackingList<ManagedBackupShortTermRetentionPolicyData>();
        }

        /// <summary> Initializes a new instance of <see cref="ManagedBackupShortTermRetentionPolicyListResult"/>. </summary>
        /// <param name="value"> Array of results. </param>
        /// <param name="nextLink"> Link to retrieve next page of results. </param>
        internal ManagedBackupShortTermRetentionPolicyListResult(IReadOnlyList<ManagedBackupShortTermRetentionPolicyData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> Array of results. </summary>
        public IReadOnlyList<ManagedBackupShortTermRetentionPolicyData> Value { get; }
        /// <summary> Link to retrieve next page of results. </summary>
        public string NextLink { get; }
    }
}
