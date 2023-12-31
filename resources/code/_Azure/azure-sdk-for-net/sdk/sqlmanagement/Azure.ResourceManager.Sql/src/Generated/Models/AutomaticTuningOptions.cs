// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> Automatic tuning properties for individual advisors. </summary>
    public partial class AutomaticTuningOptions
    {
        /// <summary> Initializes a new instance of <see cref="AutomaticTuningOptions"/>. </summary>
        public AutomaticTuningOptions()
        {
        }

        /// <summary> Initializes a new instance of <see cref="AutomaticTuningOptions"/>. </summary>
        /// <param name="desiredState"> Automatic tuning option desired state. </param>
        /// <param name="actualState"> Automatic tuning option actual state. </param>
        /// <param name="reasonCode"> Reason code if desired and actual state are different. </param>
        /// <param name="reasonDesc"> Reason description if desired and actual state are different. </param>
        internal AutomaticTuningOptions(AutomaticTuningOptionModeDesired? desiredState, AutomaticTuningOptionModeActual? actualState, int? reasonCode, AutomaticTuningDisabledReason? reasonDesc)
        {
            DesiredState = desiredState;
            ActualState = actualState;
            ReasonCode = reasonCode;
            ReasonDesc = reasonDesc;
        }

        /// <summary> Automatic tuning option desired state. </summary>
        public AutomaticTuningOptionModeDesired? DesiredState { get; set; }
        /// <summary> Automatic tuning option actual state. </summary>
        public AutomaticTuningOptionModeActual? ActualState { get; }
        /// <summary> Reason code if desired and actual state are different. </summary>
        public int? ReasonCode { get; }
        /// <summary> Reason description if desired and actual state are different. </summary>
        public AutomaticTuningDisabledReason? ReasonDesc { get; }
    }
}
