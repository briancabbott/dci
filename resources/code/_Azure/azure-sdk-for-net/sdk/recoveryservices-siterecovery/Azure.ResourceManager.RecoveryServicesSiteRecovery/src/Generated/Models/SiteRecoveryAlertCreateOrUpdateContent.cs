// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> Request to configure alerts for the system. </summary>
    public partial class SiteRecoveryAlertCreateOrUpdateContent
    {
        /// <summary> Initializes a new instance of <see cref="SiteRecoveryAlertCreateOrUpdateContent"/>. </summary>
        public SiteRecoveryAlertCreateOrUpdateContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SiteRecoveryAlertCreateOrUpdateContent"/>. </summary>
        /// <param name="properties"> The properties of a configure alert request. </param>
        internal SiteRecoveryAlertCreateOrUpdateContent(SiteRecoveryConfigureAlertProperties properties)
        {
            Properties = properties;
        }

        /// <summary> The properties of a configure alert request. </summary>
        public SiteRecoveryConfigureAlertProperties Properties { get; set; }
    }
}
