// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> Resolve health input properties. </summary>
    internal partial class ResolveHealthContentProperties
    {
        /// <summary> Initializes a new instance of <see cref="ResolveHealthContentProperties"/>. </summary>
        public ResolveHealthContentProperties()
        {
            HealthErrors = new ChangeTrackingList<ResolveHealthError>();
        }

        /// <summary> Initializes a new instance of <see cref="ResolveHealthContentProperties"/>. </summary>
        /// <param name="healthErrors"> Health errors. </param>
        internal ResolveHealthContentProperties(IList<ResolveHealthError> healthErrors)
        {
            HealthErrors = healthErrors;
        }

        /// <summary> Health errors. </summary>
        public IList<ResolveHealthError> HealthErrors { get; }
    }
}
