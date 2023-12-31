// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Reservations.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    public partial class Catalog
    {
        /// <summary>
        /// Initializes a new instance of the Catalog class.
        /// </summary>
        public Catalog()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the Catalog class.
        /// </summary>
        /// <param name="resourceType">The type of resource the SKU applies
        /// to.</param>
        /// <param name="name">The name of SKU</param>
        /// <param name="billingPlans">The billing plan options available for
        /// this SKU.</param>
        /// <param name="terms">Available reservation terms for this
        /// resource</param>
        /// <param name="msrp">Pricing information about the SKU</param>
        /// <param name="tier">The tier of this SKU</param>
        /// <param name="size">The size of this SKU</param>
        public Catalog(string resourceType = default(string), string name = default(string), IDictionary<string, IList<string>> billingPlans = default(IDictionary<string, IList<string>>), IList<string> terms = default(IList<string>), IList<string> locations = default(IList<string>), IList<SkuProperty> skuProperties = default(IList<SkuProperty>), CatalogMsrp msrp = default(CatalogMsrp), IList<SkuRestriction> restrictions = default(IList<SkuRestriction>), string tier = default(string), string size = default(string), IList<SkuCapability> capabilities = default(IList<SkuCapability>))
        {
            ResourceType = resourceType;
            Name = name;
            BillingPlans = billingPlans;
            Terms = terms;
            Locations = locations;
            SkuProperties = skuProperties;
            Msrp = msrp;
            Restrictions = restrictions;
            Tier = tier;
            Size = size;
            Capabilities = capabilities;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the type of resource the SKU applies to.
        /// </summary>
        [JsonProperty(PropertyName = "resourceType")]
        public string ResourceType { get; private set; }

        /// <summary>
        /// Gets the name of SKU
        /// </summary>
        [JsonProperty(PropertyName = "name")]
        public string Name { get; private set; }

        /// <summary>
        /// Gets or sets the billing plan options available for this SKU.
        /// </summary>
        [JsonProperty(PropertyName = "billingPlans")]
        public IDictionary<string, IList<string>> BillingPlans { get; set; }

        /// <summary>
        /// Gets available reservation terms for this resource
        /// </summary>
        [JsonProperty(PropertyName = "terms")]
        public IList<string> Terms { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "locations")]
        public IList<string> Locations { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "skuProperties")]
        public IList<SkuProperty> SkuProperties { get; private set; }

        /// <summary>
        /// Gets pricing information about the SKU
        /// </summary>
        [JsonProperty(PropertyName = "msrp")]
        public CatalogMsrp Msrp { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "restrictions")]
        public IList<SkuRestriction> Restrictions { get; private set; }

        /// <summary>
        /// Gets the tier of this SKU
        /// </summary>
        [JsonProperty(PropertyName = "tier")]
        public string Tier { get; private set; }

        /// <summary>
        /// Gets the size of this SKU
        /// </summary>
        [JsonProperty(PropertyName = "size")]
        public string Size { get; private set; }

        /// <summary>
        /// </summary>
        [JsonProperty(PropertyName = "capabilities")]
        public IList<SkuCapability> Capabilities { get; private set; }

    }
}
