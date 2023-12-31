// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> Updates network mappings input. </summary>
    public partial class A2AUpdateNetworkMappingContent : FabricSpecificUpdateNetworkMappingContent
    {
        /// <summary> Initializes a new instance of <see cref="A2AUpdateNetworkMappingContent"/>. </summary>
        public A2AUpdateNetworkMappingContent()
        {
            InstanceType = "AzureToAzure";
        }

        /// <summary> Initializes a new instance of <see cref="A2AUpdateNetworkMappingContent"/>. </summary>
        /// <param name="instanceType"> The instance type. </param>
        /// <param name="primaryNetworkId"> The primary azure vnet Id. </param>
        internal A2AUpdateNetworkMappingContent(string instanceType, ResourceIdentifier primaryNetworkId) : base(instanceType)
        {
            PrimaryNetworkId = primaryNetworkId;
            InstanceType = instanceType ?? "AzureToAzure";
        }

        /// <summary> The primary azure vnet Id. </summary>
        public ResourceIdentifier PrimaryNetworkId { get; set; }
    }
}
