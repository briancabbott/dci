// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;
using Azure.ResourceManager.Network;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Response for ListExpressRouteCrossConnection API service call. </summary>
    internal partial class ExpressRouteCrossConnectionListResult
    {
        /// <summary> Initializes a new instance of <see cref="ExpressRouteCrossConnectionListResult"/>. </summary>
        internal ExpressRouteCrossConnectionListResult()
        {
            Value = new ChangeTrackingList<ExpressRouteCrossConnectionData>();
        }

        /// <summary> Initializes a new instance of <see cref="ExpressRouteCrossConnectionListResult"/>. </summary>
        /// <param name="value"> A list of ExpressRouteCrossConnection resources. </param>
        /// <param name="nextLink"> The URL to get the next set of results. </param>
        internal ExpressRouteCrossConnectionListResult(IReadOnlyList<ExpressRouteCrossConnectionData> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> A list of ExpressRouteCrossConnection resources. </summary>
        public IReadOnlyList<ExpressRouteCrossConnectionData> Value { get; }
        /// <summary> The URL to get the next set of results. </summary>
        public string NextLink { get; }
    }
}
