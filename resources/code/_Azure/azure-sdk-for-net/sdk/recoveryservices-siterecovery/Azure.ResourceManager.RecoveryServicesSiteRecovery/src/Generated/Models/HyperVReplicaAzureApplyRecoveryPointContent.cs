// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> ApplyRecoveryPoint input specific to HyperVReplicaAzure provider. </summary>
    public partial class HyperVReplicaAzureApplyRecoveryPointContent : SiteRecoveryApplyRecoveryPointProviderSpecificContent
    {
        /// <summary> Initializes a new instance of <see cref="HyperVReplicaAzureApplyRecoveryPointContent"/>. </summary>
        public HyperVReplicaAzureApplyRecoveryPointContent()
        {
            InstanceType = "HyperVReplicaAzure";
        }

        /// <summary> Initializes a new instance of <see cref="HyperVReplicaAzureApplyRecoveryPointContent"/>. </summary>
        /// <param name="instanceType"> The class type. </param>
        /// <param name="primaryKekCertificatePfx"> The primary kek certificate pfx. </param>
        /// <param name="secondaryKekCertificatePfx"> The secondary kek certificate pfx. </param>
        internal HyperVReplicaAzureApplyRecoveryPointContent(string instanceType, string primaryKekCertificatePfx, string secondaryKekCertificatePfx) : base(instanceType)
        {
            PrimaryKekCertificatePfx = primaryKekCertificatePfx;
            SecondaryKekCertificatePfx = secondaryKekCertificatePfx;
            InstanceType = instanceType ?? "HyperVReplicaAzure";
        }

        /// <summary> The primary kek certificate pfx. </summary>
        public string PrimaryKekCertificatePfx { get; set; }
        /// <summary> The secondary kek certificate pfx. </summary>
        public string SecondaryKekCertificatePfx { get; set; }
    }
}
