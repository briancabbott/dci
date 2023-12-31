// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Storage.Files.DataLake.Models
{
    /// <summary> The AclFailedEntry. </summary>
    internal partial class AclFailedEntry
    {
        /// <summary> Initializes a new instance of <see cref="AclFailedEntry"/>. </summary>
        internal AclFailedEntry()
        {
        }

        /// <summary> Initializes a new instance of <see cref="AclFailedEntry"/>. </summary>
        /// <param name="name"></param>
        /// <param name="type"></param>
        /// <param name="errorMessage"></param>
        internal AclFailedEntry(string name, string type, string errorMessage)
        {
            Name = name;
            Type = type;
            ErrorMessage = errorMessage;
        }
    }
}
