// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.MachineLearning.Models
{
    /// <summary> The MachineLearningDeploymentLogsContent. </summary>
    public partial class MachineLearningDeploymentLogsContent
    {
        /// <summary> Initializes a new instance of <see cref="MachineLearningDeploymentLogsContent"/>. </summary>
        public MachineLearningDeploymentLogsContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="MachineLearningDeploymentLogsContent"/>. </summary>
        /// <param name="containerType"> The type of container to retrieve logs from. </param>
        /// <param name="tail"> The maximum number of lines to tail. </param>
        internal MachineLearningDeploymentLogsContent(MachineLearningContainerType? containerType, int? tail)
        {
            ContainerType = containerType;
            Tail = tail;
        }

        /// <summary> The type of container to retrieve logs from. </summary>
        public MachineLearningContainerType? ContainerType { get; set; }
        /// <summary> The maximum number of lines to tail. </summary>
        public int? Tail { get; set; }
    }
}
