// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.DesktopVirtualization.Models
{
    /// <summary> Application properties that can be patched. </summary>
    public partial class VirtualApplicationPatch
    {
        /// <summary> Initializes a new instance of <see cref="VirtualApplicationPatch"/>. </summary>
        public VirtualApplicationPatch()
        {
            Tags = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="VirtualApplicationPatch"/>. </summary>
        /// <param name="tags"> tags to be updated. </param>
        /// <param name="description"> Description of Application. </param>
        /// <param name="friendlyName"> Friendly name of Application. </param>
        /// <param name="filePath"> Specifies a path for the executable file for the application. </param>
        /// <param name="commandLineSetting"> Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all. </param>
        /// <param name="commandLineArguments"> Command Line Arguments for Application. </param>
        /// <param name="showInPortal"> Specifies whether to show the RemoteApp program in the RD Web Access server. </param>
        /// <param name="iconPath"> Path to icon. </param>
        /// <param name="iconIndex"> Index of the icon. </param>
        /// <param name="msixPackageFamilyName"> Specifies the package family name for MSIX applications. </param>
        /// <param name="msixPackageApplicationId"> Specifies the package application Id for MSIX applications. </param>
        /// <param name="applicationType"> Resource Type of Application. </param>
        internal VirtualApplicationPatch(IDictionary<string, string> tags, string description, string friendlyName, string filePath, VirtualApplicationCommandLineSetting? commandLineSetting, string commandLineArguments, bool? showInPortal, string iconPath, int? iconIndex, string msixPackageFamilyName, string msixPackageApplicationId, RemoteApplicationType? applicationType)
        {
            Tags = tags;
            Description = description;
            FriendlyName = friendlyName;
            FilePath = filePath;
            CommandLineSetting = commandLineSetting;
            CommandLineArguments = commandLineArguments;
            ShowInPortal = showInPortal;
            IconPath = iconPath;
            IconIndex = iconIndex;
            MsixPackageFamilyName = msixPackageFamilyName;
            MsixPackageApplicationId = msixPackageApplicationId;
            ApplicationType = applicationType;
        }

        /// <summary> tags to be updated. </summary>
        public IDictionary<string, string> Tags { get; }
        /// <summary> Description of Application. </summary>
        public string Description { get; set; }
        /// <summary> Friendly name of Application. </summary>
        public string FriendlyName { get; set; }
        /// <summary> Specifies a path for the executable file for the application. </summary>
        public string FilePath { get; set; }
        /// <summary> Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all. </summary>
        public VirtualApplicationCommandLineSetting? CommandLineSetting { get; set; }
        /// <summary> Command Line Arguments for Application. </summary>
        public string CommandLineArguments { get; set; }
        /// <summary> Specifies whether to show the RemoteApp program in the RD Web Access server. </summary>
        public bool? ShowInPortal { get; set; }
        /// <summary> Path to icon. </summary>
        public string IconPath { get; set; }
        /// <summary> Index of the icon. </summary>
        public int? IconIndex { get; set; }
        /// <summary> Specifies the package family name for MSIX applications. </summary>
        public string MsixPackageFamilyName { get; set; }
        /// <summary> Specifies the package application Id for MSIX applications. </summary>
        public string MsixPackageApplicationId { get; set; }
        /// <summary> Resource Type of Application. </summary>
        public RemoteApplicationType? ApplicationType { get; set; }
    }
}
