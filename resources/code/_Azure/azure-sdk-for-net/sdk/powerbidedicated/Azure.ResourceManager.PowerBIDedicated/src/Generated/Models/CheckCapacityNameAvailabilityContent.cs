// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.PowerBIDedicated.Models
{
    /// <summary> Details of capacity name request body. </summary>
    public partial class CheckCapacityNameAvailabilityContent
    {
        /// <summary> Initializes a new instance of <see cref="CheckCapacityNameAvailabilityContent"/>. </summary>
        public CheckCapacityNameAvailabilityContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="CheckCapacityNameAvailabilityContent"/>. </summary>
        /// <param name="name"> Name for checking availability. </param>
        /// <param name="resourceType"> The resource type of PowerBI dedicated. </param>
        internal CheckCapacityNameAvailabilityContent(string name, string resourceType)
        {
            Name = name;
            ResourceType = resourceType;
        }

        /// <summary> Name for checking availability. </summary>
        public string Name { get; set; }
        /// <summary> The resource type of PowerBI dedicated. </summary>
        public string ResourceType { get; set; }
    }
}
