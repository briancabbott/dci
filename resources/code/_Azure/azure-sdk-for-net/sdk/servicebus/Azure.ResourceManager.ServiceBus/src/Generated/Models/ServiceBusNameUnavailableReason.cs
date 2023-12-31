// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.ServiceBus.Models
{
    /// <summary> Specifies the reason for the unavailability of the service. </summary>
    public enum ServiceBusNameUnavailableReason
    {
        /// <summary> None. </summary>
        None,
        /// <summary> InvalidName. </summary>
        InvalidName,
        /// <summary> SubscriptionIsDisabled. </summary>
        SubscriptionIsDisabled,
        /// <summary> NameInUse. </summary>
        NameInUse,
        /// <summary> NameInLockdown. </summary>
        NameInLockdown,
        /// <summary> TooManyNamespaceInCurrentSubscription. </summary>
        TooManyNamespaceInCurrentSubscription
    }
}
