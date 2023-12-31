﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.ComponentModel;

namespace Azure.ResourceManager.Compute.Models
{
    public partial class LinuxPatchSettings
    {
        /// <summary> Specifies the reboot setting for all AutomaticByPlatform patch installation operations. </summary>
        [EditorBrowsable(EditorBrowsableState.Never)]
        public LinuxVmGuestPatchAutomaticByPlatformRebootSetting? AutomaticByPlatformRebootSetting
        {
            get => AutomaticByPlatformSettings is null ? default : AutomaticByPlatformSettings.RebootSetting;
            set
            {
                if (AutomaticByPlatformSettings is null)
                    AutomaticByPlatformSettings = new LinuxVmGuestPatchAutomaticByPlatformSettings();
                AutomaticByPlatformSettings.RebootSetting = value;
            }
        }
    }
}
