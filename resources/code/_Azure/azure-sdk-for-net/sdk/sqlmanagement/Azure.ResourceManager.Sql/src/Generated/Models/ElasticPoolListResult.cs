// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Sql;

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> The result of an elastic pool list request. </summary>
    internal partial class ElasticPoolListResult
    {
        /// <summary> Initializes a new instance of <see cref="ElasticPoolListResult"/>. </summary>
        internal ElasticPoolListResult()
        {
            Value = new ChangeTrackingList<ElasticPoolData>();
        }

        /// <summary> Initializes a new instance of <see cref="ElasticPoolListResult"/>. </summary>
        /// <param name="value"> Array of results. </param>
        /// <param name="nextLink"> Link to retrieve next page of results. </param>
        internal ElasticPoolListResult(IReadOnlyList<ElasticPoolData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> Array of results. </summary>
        public IReadOnlyList<ElasticPoolData> Value { get; }
        /// <summary> Link to retrieve next page of results. </summary>
        public string NextLink { get; }
    }
}
