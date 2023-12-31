// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Compute.Models
{
    /// <summary> A UEFI key signature. </summary>
    public partial class UefiKey
    {
        /// <summary> Initializes a new instance of <see cref="UefiKey"/>. </summary>
        public UefiKey()
        {
            Value = new ChangeTrackingList<string>();
        }

        /// <summary> Initializes a new instance of <see cref="UefiKey"/>. </summary>
        /// <param name="keyType"> The type of key signature. </param>
        /// <param name="value"> The value of the key signature. </param>
        internal UefiKey(UefiKeyType? keyType, IList<string> value)
        {
            KeyType = keyType;
            Value = value;
        }

        /// <summary> The type of key signature. </summary>
        public UefiKeyType? KeyType { get; set; }
        /// <summary> The value of the key signature. </summary>
        public IList<string> Value { get; }
    }
}
