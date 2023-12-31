// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Marketplace.Models
{
    /// <summary> The SetRulesContent. </summary>
    public partial class SetRulesContent
    {
        /// <summary> Initializes a new instance of <see cref="SetRulesContent"/>. </summary>
        public SetRulesContent()
        {
            Value = new ChangeTrackingList<MarketplaceRule>();
        }

        /// <summary> Initializes a new instance of <see cref="SetRulesContent"/>. </summary>
        /// <param name="value"></param>
        /// <param name="nextLink"> URL to get the next set of rules list results if there are any. </param>
        internal SetRulesContent(IList<MarketplaceRule> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> Gets the value. </summary>
        public IList<MarketplaceRule> Value { get; }
        /// <summary> URL to get the next set of rules list results if there are any. </summary>
        public string NextLink { get; set; }
    }
}
