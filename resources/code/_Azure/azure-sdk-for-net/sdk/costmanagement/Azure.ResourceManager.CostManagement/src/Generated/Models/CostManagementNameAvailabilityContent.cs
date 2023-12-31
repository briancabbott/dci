// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.CostManagement.Models
{
    /// <summary> The check availability request body. </summary>
    public partial class CostManagementNameAvailabilityContent
    {
        /// <summary> Initializes a new instance of <see cref="CostManagementNameAvailabilityContent"/>. </summary>
        public CostManagementNameAvailabilityContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="CostManagementNameAvailabilityContent"/>. </summary>
        /// <param name="name"> The name of the resource for which availability needs to be checked. </param>
        /// <param name="resourceType"> The resource type. </param>
        internal CostManagementNameAvailabilityContent(string name, string resourceType)
        {
            Name = name;
            ResourceType = resourceType;
        }

        /// <summary> The name of the resource for which availability needs to be checked. </summary>
        public string Name { get; set; }
        /// <summary> The resource type. </summary>
        public string ResourceType { get; set; }
    }
}
