﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;

namespace Azure.Storage.Blobs.Models
{
    /// <summary>
    /// PageInfo.
    /// </summary>
    public class PageInfo
    {
        /// <summary>
        /// The ETag contains a value that you can use to perform operations conditionally.
        /// If the request version is 2011-08-18 or newer, the ETag value will be in quotes.
        /// </summary>
        public ETag ETag { get; internal set; }

        /// <summary>
        /// Returns the date and time the container was last modified.
        /// Any operation that modifies the blob, including an update of the blob's metadata or properties,
        /// changes the last-modified time of the blob.
        /// </summary>
        public DateTimeOffset LastModified { get; internal set; }

        /// <summary>
        /// If the blob has an MD5 hash and this operation is to read the full blob,
        /// this response header is returned so that the client can check for message content integrity.
        /// </summary>
#pragma warning disable CA1819 // Properties should not return arrays
        public byte[] ContentHash { get; internal set; }
#pragma warning restore CA1819 // Properties should not return arrays

        /// <summary>
        /// This header is returned so that the client can check for message content integrity.
        /// The value of this header is computed by the Blob service; it is not necessarily the same value specified in the request headers.
        /// </summary>
#pragma warning disable CA1819 // Properties should not return arrays
        public byte[] ContentCrc64 { get; internal set; }
#pragma warning restore CA1819 // Properties should not return arrays

        /// <summary>
        /// The current sequence number for the page blob.  This is only returned for page blobs.
        /// </summary>
        public long BlobSequenceNumber { get; internal set; }

        /// <summary>
        /// The SHA-256 hash of the encryption key used to encrypt the pages.
        /// This header is only returned when the pages were encrypted with a customer-provided key.
        /// </summary>
        public string EncryptionKeySha256 { get; internal set; }

        /// <summary>
        /// Returns the name of the encryption scope used to encrypt the blob contents and application metadata.
        /// Note that the absence of this header implies use of the default account encryption scope.
        /// </summary>
        public string EncryptionScope { get; internal set; }

        /// <summary>
        /// Prevent direct instantiation of PageInfo instances.
        /// You can use BlobsModelFactory.PageInfo instead.
        /// </summary>
        internal PageInfo() { }
    }
}
