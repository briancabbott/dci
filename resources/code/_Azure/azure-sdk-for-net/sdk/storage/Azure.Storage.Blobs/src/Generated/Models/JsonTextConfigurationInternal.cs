// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Storage.Blobs.Models
{
    /// <summary> json text configuration. </summary>
    internal partial class JsonTextConfigurationInternal
    {
        /// <summary> Initializes a new instance of <see cref="JsonTextConfigurationInternal"/>. </summary>
        public JsonTextConfigurationInternal()
        {
        }

        /// <summary> Initializes a new instance of <see cref="JsonTextConfigurationInternal"/>. </summary>
        /// <param name="recordSeparator"> The string used to separate records. </param>
        internal JsonTextConfigurationInternal(string recordSeparator)
        {
            RecordSeparator = recordSeparator;
        }

        /// <summary> The string used to separate records. </summary>
        public string RecordSeparator { get; set; }
    }
}
