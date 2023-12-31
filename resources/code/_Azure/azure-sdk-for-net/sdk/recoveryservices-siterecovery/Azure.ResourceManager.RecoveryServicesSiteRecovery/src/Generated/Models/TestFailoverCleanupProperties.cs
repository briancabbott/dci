// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    /// <summary> Input definition for test failover cleanup input properties. </summary>
    public partial class TestFailoverCleanupProperties
    {
        /// <summary> Initializes a new instance of <see cref="TestFailoverCleanupProperties"/>. </summary>
        public TestFailoverCleanupProperties()
        {
        }

        /// <summary> Initializes a new instance of <see cref="TestFailoverCleanupProperties"/>. </summary>
        /// <param name="comments"> Test failover cleanup comments. </param>
        internal TestFailoverCleanupProperties(string comments)
        {
            Comments = comments;
        }

        /// <summary> Test failover cleanup comments. </summary>
        public string Comments { get; set; }
    }
}
