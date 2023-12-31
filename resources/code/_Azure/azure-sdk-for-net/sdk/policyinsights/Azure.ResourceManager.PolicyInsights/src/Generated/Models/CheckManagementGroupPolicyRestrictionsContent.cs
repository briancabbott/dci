// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.PolicyInsights.Models
{
    /// <summary> The check policy restrictions parameters describing the resource that is being evaluated. </summary>
    public partial class CheckManagementGroupPolicyRestrictionsContent
    {
        /// <summary> Initializes a new instance of <see cref="CheckManagementGroupPolicyRestrictionsContent"/>. </summary>
        public CheckManagementGroupPolicyRestrictionsContent()
        {
            PendingFields = new ChangeTrackingList<PendingField>();
        }

        /// <summary> Initializes a new instance of <see cref="CheckManagementGroupPolicyRestrictionsContent"/>. </summary>
        /// <param name="resourceDetails"> The information about the resource that will be evaluated. </param>
        /// <param name="pendingFields"> The list of fields and values that should be evaluated for potential restrictions. </param>
        internal CheckManagementGroupPolicyRestrictionsContent(CheckRestrictionsResourceDetails resourceDetails, IList<PendingField> pendingFields)
        {
            ResourceDetails = resourceDetails;
            PendingFields = pendingFields;
        }

        /// <summary> The information about the resource that will be evaluated. </summary>
        public CheckRestrictionsResourceDetails ResourceDetails { get; set; }
        /// <summary> The list of fields and values that should be evaluated for potential restrictions. </summary>
        public IList<PendingField> PendingFields { get; }
    }
}
