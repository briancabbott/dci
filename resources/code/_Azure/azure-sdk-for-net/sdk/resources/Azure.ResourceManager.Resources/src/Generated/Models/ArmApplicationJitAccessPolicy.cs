// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Resources.Models
{
    /// <summary> Managed application Jit access policy. </summary>
    public partial class ArmApplicationJitAccessPolicy
    {
        /// <summary> Initializes a new instance of <see cref="ArmApplicationJitAccessPolicy"/>. </summary>
        /// <param name="jitAccessEnabled"> Whether the JIT access is enabled. </param>
        public ArmApplicationJitAccessPolicy(bool jitAccessEnabled)
        {
            JitAccessEnabled = jitAccessEnabled;
            JitApprovers = new ChangeTrackingList<JitApprover>();
        }

        /// <summary> Initializes a new instance of <see cref="ArmApplicationJitAccessPolicy"/>. </summary>
        /// <param name="jitAccessEnabled"> Whether the JIT access is enabled. </param>
        /// <param name="jitApprovalMode"> JIT approval mode. </param>
        /// <param name="jitApprovers"> The JIT approvers. </param>
        /// <param name="maximumJitAccessDuration"> The maximum duration JIT access is granted. This is an ISO8601 time period value. </param>
        internal ArmApplicationJitAccessPolicy(bool jitAccessEnabled, JitApprovalMode? jitApprovalMode, IList<JitApprover> jitApprovers, TimeSpan? maximumJitAccessDuration)
        {
            JitAccessEnabled = jitAccessEnabled;
            JitApprovalMode = jitApprovalMode;
            JitApprovers = jitApprovers;
            MaximumJitAccessDuration = maximumJitAccessDuration;
        }

        /// <summary> Whether the JIT access is enabled. </summary>
        public bool JitAccessEnabled { get; set; }
        /// <summary> JIT approval mode. </summary>
        public JitApprovalMode? JitApprovalMode { get; set; }
        /// <summary> The JIT approvers. </summary>
        public IList<JitApprover> JitApprovers { get; }
        /// <summary> The maximum duration JIT access is granted. This is an ISO8601 time period value. </summary>
        public TimeSpan? MaximumJitAccessDuration { get; set; }
    }
}
