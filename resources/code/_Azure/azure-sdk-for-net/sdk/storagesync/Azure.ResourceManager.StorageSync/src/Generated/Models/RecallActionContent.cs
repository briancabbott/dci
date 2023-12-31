// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.StorageSync.Models
{
    /// <summary> The parameters used when calling recall action on server endpoint. </summary>
    public partial class RecallActionContent
    {
        /// <summary> Initializes a new instance of <see cref="RecallActionContent"/>. </summary>
        public RecallActionContent()
        {
        }

        /// <summary> Initializes a new instance of <see cref="RecallActionContent"/>. </summary>
        /// <param name="pattern"> Pattern of the files. </param>
        /// <param name="recallPath"> Recall path. </param>
        internal RecallActionContent(string pattern, string recallPath)
        {
            Pattern = pattern;
            RecallPath = recallPath;
        }

        /// <summary> Pattern of the files. </summary>
        public string Pattern { get; set; }
        /// <summary> Recall path. </summary>
        public string RecallPath { get; set; }
    }
}
