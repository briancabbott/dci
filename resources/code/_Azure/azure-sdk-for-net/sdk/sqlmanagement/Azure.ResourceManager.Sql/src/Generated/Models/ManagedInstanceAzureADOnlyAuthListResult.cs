// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Sql;

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> A list of active directory only authentications. </summary>
    internal partial class ManagedInstanceAzureADOnlyAuthListResult
    {
        /// <summary> Initializes a new instance of <see cref="ManagedInstanceAzureADOnlyAuthListResult"/>. </summary>
        internal ManagedInstanceAzureADOnlyAuthListResult()
        {
            Value = new ChangeTrackingList<ManagedInstanceAzureADOnlyAuthenticationData>();
        }

        /// <summary> Initializes a new instance of <see cref="ManagedInstanceAzureADOnlyAuthListResult"/>. </summary>
        /// <param name="value"> Array of results. </param>
        /// <param name="nextLink"> Link to retrieve next page of results. </param>
        internal ManagedInstanceAzureADOnlyAuthListResult(IReadOnlyList<ManagedInstanceAzureADOnlyAuthenticationData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> Array of results. </summary>
        public IReadOnlyList<ManagedInstanceAzureADOnlyAuthenticationData> Value { get; }
        /// <summary> Link to retrieve next page of results. </summary>
        public string NextLink { get; }
    }
}
