// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.RecoveryServices
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for VaultCertificatesOperations.
    /// </summary>
    public static partial class VaultCertificatesOperationsExtensions
    {
            /// <summary>
            /// Uploads a certificate for a resource.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='vaultName'>
            /// The name of the recovery services vault.
            /// </param>
            /// <param name='certificateName'>
            /// Certificate friendly name.
            /// </param>
            /// <param name='properties'>
            /// </param>
            public static VaultCertificateResponse Create(this IVaultCertificatesOperations operations, string resourceGroupName, string vaultName, string certificateName, RawCertificateData properties = default(RawCertificateData))
            {
                return operations.CreateAsync(resourceGroupName, vaultName, certificateName, properties).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Uploads a certificate for a resource.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='vaultName'>
            /// The name of the recovery services vault.
            /// </param>
            /// <param name='certificateName'>
            /// Certificate friendly name.
            /// </param>
            /// <param name='properties'>
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<VaultCertificateResponse> CreateAsync(this IVaultCertificatesOperations operations, string resourceGroupName, string vaultName, string certificateName, RawCertificateData properties = default(RawCertificateData), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateWithHttpMessagesAsync(resourceGroupName, vaultName, certificateName, properties, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
