// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> Read-write endpoint of the failover group instance. </summary>
    public partial class FailoverGroupReadWriteEndpoint
    {
        /// <summary> Initializes a new instance of <see cref="FailoverGroupReadWriteEndpoint"/>. </summary>
        /// <param name="failoverPolicy"> Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required. </param>
        public FailoverGroupReadWriteEndpoint(ReadWriteEndpointFailoverPolicy failoverPolicy)
        {
            FailoverPolicy = failoverPolicy;
        }

        /// <summary> Initializes a new instance of <see cref="FailoverGroupReadWriteEndpoint"/>. </summary>
        /// <param name="failoverPolicy"> Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required. </param>
        /// <param name="failoverWithDataLossGracePeriodMinutes"> Grace period before failover with data loss is attempted for the read-write endpoint. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required. </param>
        internal FailoverGroupReadWriteEndpoint(ReadWriteEndpointFailoverPolicy failoverPolicy, int? failoverWithDataLossGracePeriodMinutes)
        {
            FailoverPolicy = failoverPolicy;
            FailoverWithDataLossGracePeriodMinutes = failoverWithDataLossGracePeriodMinutes;
        }

        /// <summary> Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required. </summary>
        public ReadWriteEndpointFailoverPolicy FailoverPolicy { get; set; }
        /// <summary> Grace period before failover with data loss is attempted for the read-write endpoint. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required. </summary>
        public int? FailoverWithDataLossGracePeriodMinutes { get; set; }
    }
}
