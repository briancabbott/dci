// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.ResourceManager.HealthcareApis.Models
{
    /// <summary> Settings to encrypt a service. </summary>
    internal partial class Encryption
    {
        /// <summary> Initializes a new instance of <see cref="Encryption"/>. </summary>
        public Encryption()
        {
        }

        /// <summary> Initializes a new instance of <see cref="Encryption"/>. </summary>
        /// <param name="customerManagedKeyEncryption"> The encryption settings for the customer-managed key. </param>
        internal Encryption(EncryptionCustomerManagedKeyEncryption customerManagedKeyEncryption)
        {
            CustomerManagedKeyEncryption = customerManagedKeyEncryption;
        }

        /// <summary> The encryption settings for the customer-managed key. </summary>
        internal EncryptionCustomerManagedKeyEncryption CustomerManagedKeyEncryption { get; set; }
        /// <summary> The URL of the key to use for encryption. </summary>
        public Uri KeyEncryptionKeyUri
        {
            get => CustomerManagedKeyEncryption is null ? default : CustomerManagedKeyEncryption.KeyEncryptionKeyUri;
            set
            {
                if (CustomerManagedKeyEncryption is null)
                    CustomerManagedKeyEncryption = new EncryptionCustomerManagedKeyEncryption();
                CustomerManagedKeyEncryption.KeyEncryptionKeyUri = value;
            }
        }
    }
}
