// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;

namespace Azure.ResourceManager.Communication.Models
{
    /// <summary> The check availability request body. </summary>
    public partial class CommunicationNameAvailabilityContent
    {
        /// <summary> Initializes a new instance of <see cref="CommunicationNameAvailabilityContent"/>. </summary>
        public CommunicationNameAvailabilityContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="CommunicationNameAvailabilityContent"/>. </summary>
        /// <param name="name"> The name of the resource for which availability needs to be checked. </param>
        /// <param name="resourceType"> The resource type. </param>
        internal CommunicationNameAvailabilityContent(string name, ResourceType? resourceType)
        {
            Name = name;
            ResourceType = resourceType;
        }

        /// <summary> The name of the resource for which availability needs to be checked. </summary>
        public string Name { get; set; }
        /// <summary> The resource type. </summary>
        public ResourceType? ResourceType { get; set; }
    }
}
