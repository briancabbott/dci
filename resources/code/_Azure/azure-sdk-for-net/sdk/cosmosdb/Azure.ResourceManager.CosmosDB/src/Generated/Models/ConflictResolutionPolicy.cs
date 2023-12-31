// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.CosmosDB.Models
{
    /// <summary> The conflict resolution policy for the container. </summary>
    public partial class ConflictResolutionPolicy
    {
        /// <summary> Initializes a new instance of <see cref="ConflictResolutionPolicy"/>. </summary>
        public ConflictResolutionPolicy()
        {
        }

        /// <summary> Initializes a new instance of <see cref="ConflictResolutionPolicy"/>. </summary>
        /// <param name="mode"> Indicates the conflict resolution mode. </param>
        /// <param name="conflictResolutionPath"> The conflict resolution path in the case of LastWriterWins mode. </param>
        /// <param name="conflictResolutionProcedure"> The procedure to resolve conflicts in the case of custom mode. </param>
        internal ConflictResolutionPolicy(ConflictResolutionMode? mode, string conflictResolutionPath, string conflictResolutionProcedure)
        {
            Mode = mode;
            ConflictResolutionPath = conflictResolutionPath;
            ConflictResolutionProcedure = conflictResolutionProcedure;
        }

        /// <summary> Indicates the conflict resolution mode. </summary>
        public ConflictResolutionMode? Mode { get; set; }
        /// <summary> The conflict resolution path in the case of LastWriterWins mode. </summary>
        public string ConflictResolutionPath { get; set; }
        /// <summary> The procedure to resolve conflicts in the case of custom mode. </summary>
        public string ConflictResolutionProcedure { get; set; }
    }
}
