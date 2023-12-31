// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using Azure.Core;

namespace Azure.ResourceManager.DataBoxEdge.Models
{
    /// <summary> The Data Box Edge/Gateway device extended info patch. </summary>
    public partial class DataBoxEdgeDeviceExtendedInfoPatch
    {
        /// <summary> Initializes a new instance of <see cref="DataBoxEdgeDeviceExtendedInfoPatch"/>. </summary>
        public DataBoxEdgeDeviceExtendedInfoPatch()
        {
        }

        /// <summary> Initializes a new instance of <see cref="DataBoxEdgeDeviceExtendedInfoPatch"/>. </summary>
        /// <param name="clientSecretStoreId"> The Key Vault ARM Id for client secrets. </param>
        /// <param name="clientSecretStoreUri"> The url to access the Client Key Vault. </param>
        /// <param name="channelIntegrityKeyName"> The name for Channel Integrity Key stored in the Client Key Vault. </param>
        /// <param name="channelIntegrityKeyVersion"> The version of Channel Integrity Key stored in the Client Key Vault. </param>
        /// <param name="syncStatus"> For changing or to initiate the resync to key-vault set the status to KeyVaultSyncPending, rest of the status will not be applicable. </param>
        internal DataBoxEdgeDeviceExtendedInfoPatch(ResourceIdentifier clientSecretStoreId, Uri clientSecretStoreUri, string channelIntegrityKeyName, string channelIntegrityKeyVersion, EdgeKeyVaultSyncStatus? syncStatus)
        {
            ClientSecretStoreId = clientSecretStoreId;
            ClientSecretStoreUri = clientSecretStoreUri;
            ChannelIntegrityKeyName = channelIntegrityKeyName;
            ChannelIntegrityKeyVersion = channelIntegrityKeyVersion;
            SyncStatus = syncStatus;
        }

        /// <summary> The Key Vault ARM Id for client secrets. </summary>
        public ResourceIdentifier ClientSecretStoreId { get; set; }
        /// <summary> The url to access the Client Key Vault. </summary>
        public Uri ClientSecretStoreUri { get; set; }
        /// <summary> The name for Channel Integrity Key stored in the Client Key Vault. </summary>
        public string ChannelIntegrityKeyName { get; set; }
        /// <summary> The version of Channel Integrity Key stored in the Client Key Vault. </summary>
        public string ChannelIntegrityKeyVersion { get; set; }
        /// <summary> For changing or to initiate the resync to key-vault set the status to KeyVaultSyncPending, rest of the status will not be applicable. </summary>
        public EdgeKeyVaultSyncStatus? SyncStatus { get; set; }
    }
}
