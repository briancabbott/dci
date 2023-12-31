// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Compute.Models
{
    /// <summary> Information about a specific patch that was encountered during an installation action. </summary>
    public partial class PatchInstallationDetail
    {
        /// <summary> Initializes a new instance of <see cref="PatchInstallationDetail"/>. </summary>
        internal PatchInstallationDetail()
        {
            Classifications = new ChangeTrackingList<string>();
        }

        /// <summary> Initializes a new instance of <see cref="PatchInstallationDetail"/>. </summary>
        /// <param name="patchId"> A unique identifier for the patch. </param>
        /// <param name="name"> The friendly name of the patch. </param>
        /// <param name="version"> The version string of the package. It may conform to Semantic Versioning. Only applies to Linux. </param>
        /// <param name="kbId"> The KBID of the patch. Only applies to Windows patches. </param>
        /// <param name="classifications"> The classification(s) of the patch as provided by the patch publisher. </param>
        /// <param name="installationState"> The state of the patch after the installation operation completed. </param>
        internal PatchInstallationDetail(string patchId, string name, string version, string kbId, IReadOnlyList<string> classifications, PatchInstallationState? installationState)
        {
            PatchId = patchId;
            Name = name;
            Version = version;
            KbId = kbId;
            Classifications = classifications;
            InstallationState = installationState;
        }

        /// <summary> A unique identifier for the patch. </summary>
        public string PatchId { get; }
        /// <summary> The friendly name of the patch. </summary>
        public string Name { get; }
        /// <summary> The version string of the package. It may conform to Semantic Versioning. Only applies to Linux. </summary>
        public string Version { get; }
        /// <summary> The KBID of the patch. Only applies to Windows patches. </summary>
        public string KbId { get; }
        /// <summary> The classification(s) of the patch as provided by the patch publisher. </summary>
        public IReadOnlyList<string> Classifications { get; }
        /// <summary> The state of the patch after the installation operation completed. </summary>
        public PatchInstallationState? InstallationState { get; }
    }
}
