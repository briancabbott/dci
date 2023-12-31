// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.ApiManagement
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for TenantAccessGitOperations.
    /// </summary>
    public static partial class TenantAccessGitOperationsExtensions
    {
            /// <summary>
            /// Regenerate primary access key for GIT.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='serviceName'>
            /// The name of the API Management service.
            /// </param>
            /// <param name='accessName'>
            /// The identifier of the Access configuration. Possible values include:
            /// 'access', 'gitAccess'
            /// </param>
            public static void RegeneratePrimaryKey(this ITenantAccessGitOperations operations, string resourceGroupName, string serviceName, string accessName)
            {
                operations.RegeneratePrimaryKeyAsync(resourceGroupName, serviceName, accessName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Regenerate primary access key for GIT.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='serviceName'>
            /// The name of the API Management service.
            /// </param>
            /// <param name='accessName'>
            /// The identifier of the Access configuration. Possible values include:
            /// 'access', 'gitAccess'
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task RegeneratePrimaryKeyAsync(this ITenantAccessGitOperations operations, string resourceGroupName, string serviceName, string accessName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.RegeneratePrimaryKeyWithHttpMessagesAsync(resourceGroupName, serviceName, accessName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Regenerate secondary access key for GIT.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='serviceName'>
            /// The name of the API Management service.
            /// </param>
            /// <param name='accessName'>
            /// The identifier of the Access configuration. Possible values include:
            /// 'access', 'gitAccess'
            /// </param>
            public static void RegenerateSecondaryKey(this ITenantAccessGitOperations operations, string resourceGroupName, string serviceName, string accessName)
            {
                operations.RegenerateSecondaryKeyAsync(resourceGroupName, serviceName, accessName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Regenerate secondary access key for GIT.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group. The name is case insensitive.
            /// </param>
            /// <param name='serviceName'>
            /// The name of the API Management service.
            /// </param>
            /// <param name='accessName'>
            /// The identifier of the Access configuration. Possible values include:
            /// 'access', 'gitAccess'
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task RegenerateSecondaryKeyAsync(this ITenantAccessGitOperations operations, string resourceGroupName, string serviceName, string accessName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.RegenerateSecondaryKeyWithHttpMessagesAsync(resourceGroupName, serviceName, accessName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

    }
}
