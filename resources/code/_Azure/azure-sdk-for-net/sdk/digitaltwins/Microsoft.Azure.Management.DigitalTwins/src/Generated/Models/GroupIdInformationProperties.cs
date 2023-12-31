// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.DigitalTwins.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The properties for a group information object.
    /// </summary>
    public partial class GroupIdInformationProperties
    {
        /// <summary>
        /// Initializes a new instance of the GroupIdInformationProperties
        /// class.
        /// </summary>
        public GroupIdInformationProperties()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the GroupIdInformationProperties
        /// class.
        /// </summary>
        /// <param name="groupId">The group id.</param>
        /// <param name="requiredMembers">The required members for a specific
        /// group id.</param>
        /// <param name="requiredZoneNames">The required DNS zones for a
        /// specific group id.</param>
        public GroupIdInformationProperties(string groupId = default(string), IList<string> requiredMembers = default(IList<string>), IList<string> requiredZoneNames = default(IList<string>))
        {
            GroupId = groupId;
            RequiredMembers = requiredMembers;
            RequiredZoneNames = requiredZoneNames;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the group id.
        /// </summary>
        [JsonProperty(PropertyName = "groupId")]
        public string GroupId { get; set; }

        /// <summary>
        /// Gets or sets the required members for a specific group id.
        /// </summary>
        [JsonProperty(PropertyName = "requiredMembers")]
        public IList<string> RequiredMembers { get; set; }

        /// <summary>
        /// Gets or sets the required DNS zones for a specific group id.
        /// </summary>
        [JsonProperty(PropertyName = "requiredZoneNames")]
        public IList<string> RequiredZoneNames { get; set; }

    }
}
