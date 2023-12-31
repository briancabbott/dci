// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Security.KeyVault.Administration.Models
{
    /// <summary> Role Assignments filter. </summary>
    internal partial class RoleAssignmentFilter
    {
        /// <summary> Initializes a new instance of <see cref="RoleAssignmentFilter"/>. </summary>
        internal RoleAssignmentFilter()
        {
        }

        /// <summary> Initializes a new instance of <see cref="RoleAssignmentFilter"/>. </summary>
        /// <param name="principalId"> Returns role assignment of the specific principal. </param>
        internal RoleAssignmentFilter(string principalId)
        {
            PrincipalId = principalId;
        }

        /// <summary> Returns role assignment of the specific principal. </summary>
        public string PrincipalId { get; }
    }
}
