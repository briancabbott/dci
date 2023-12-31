// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> Contains SKU in an ExpressRouteCircuit. </summary>
    public partial class ExpressRouteCircuitSku
    {
        /// <summary> Initializes a new instance of <see cref="ExpressRouteCircuitSku"/>. </summary>
        public ExpressRouteCircuitSku()
        {
        }

        /// <summary> Initializes a new instance of <see cref="ExpressRouteCircuitSku"/>. </summary>
        /// <param name="name"> The name of the SKU. </param>
        /// <param name="tier"> The tier of the SKU. </param>
        /// <param name="family"> The family of the SKU. </param>
        internal ExpressRouteCircuitSku(string name, ExpressRouteCircuitSkuTier? tier, ExpressRouteCircuitSkuFamily? family)
        {
            Name = name;
            Tier = tier;
            Family = family;
        }

        /// <summary> The name of the SKU. </summary>
        public string Name { get; set; }
        /// <summary> The tier of the SKU. </summary>
        public ExpressRouteCircuitSkuTier? Tier { get; set; }
        /// <summary> The family of the SKU. </summary>
        public ExpressRouteCircuitSkuFamily? Family { get; set; }
    }
}
