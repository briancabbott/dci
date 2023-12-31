// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Storage.Files.Shares.Models
{
    /// <summary> Parameter group. </summary>
    internal partial class SourceModifiedAccessConditions
    {
        /// <summary> Initializes a new instance of <see cref="SourceModifiedAccessConditions"/>. </summary>
        public SourceModifiedAccessConditions()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SourceModifiedAccessConditions"/>. </summary>
        /// <param name="sourceIfMatchCrc64"> Specify the crc64 value to operate only on range with a matching crc64 checksum. </param>
        /// <param name="sourceIfNoneMatchCrc64"> Specify the crc64 value to operate only on range without a matching crc64 checksum. </param>
        internal SourceModifiedAccessConditions(byte[] sourceIfMatchCrc64, byte[] sourceIfNoneMatchCrc64)
        {
            SourceIfMatchCrc64 = sourceIfMatchCrc64;
            SourceIfNoneMatchCrc64 = sourceIfNoneMatchCrc64;
        }

        /// <summary> Specify the crc64 value to operate only on range with a matching crc64 checksum. </summary>
        public byte[] SourceIfMatchCrc64 { get; set; }
        /// <summary> Specify the crc64 value to operate only on range without a matching crc64 checksum. </summary>
        public byte[] SourceIfNoneMatchCrc64 { get; set; }
    }
}
