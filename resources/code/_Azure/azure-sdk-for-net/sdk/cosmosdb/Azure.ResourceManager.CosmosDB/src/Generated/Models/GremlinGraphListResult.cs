// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.CosmosDB;

namespace Azure.ResourceManager.CosmosDB.Models
{
    /// <summary> The List operation response, that contains the graphs and their properties. </summary>
    internal partial class GremlinGraphListResult
    {
        /// <summary> Initializes a new instance of <see cref="GremlinGraphListResult"/>. </summary>
        internal GremlinGraphListResult()
        {
            Value = new ChangeTrackingList<GremlinGraphData>();
        }

        /// <summary> Initializes a new instance of <see cref="GremlinGraphListResult"/>. </summary>
        /// <param name="value"> List of graphs and their properties. </param>
        internal GremlinGraphListResult(IReadOnlyList<GremlinGraphData> value)
        {
            Value = value;
        }

        /// <summary> List of graphs and their properties. </summary>
        public IReadOnlyList<GremlinGraphData> Value { get; }
    }
}
