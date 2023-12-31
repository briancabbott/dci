// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Security
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for SecurityConnectorApplicationOperations.
    /// </summary>
    public static partial class SecurityConnectorApplicationOperationsExtensions
    {
            /// <summary>
            /// Get a specific application for the requested scope by applicationId
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            public static Application Get(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId)
            {
                return operations.GetAsync(resourceGroupName, securityConnectorName, applicationId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get a specific application for the requested scope by applicationId
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<Application> GetAsync(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, securityConnectorName, applicationId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates or update a security Application on the given security connector.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            /// <param name='application'>
            /// Application over a subscription scope
            /// </param>
            public static Application CreateOrUpdate(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId, Application application)
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, securityConnectorName, applicationId, application).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates or update a security Application on the given security connector.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            /// <param name='application'>
            /// Application over a subscription scope
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<Application> CreateOrUpdateAsync(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId, Application application, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, securityConnectorName, applicationId, application, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Delete an Application over a given scope
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            public static void Delete(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId)
            {
                operations.DeleteAsync(resourceGroupName, securityConnectorName, applicationId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Delete an Application over a given scope
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription. The name is
            /// case insensitive.
            /// </param>
            /// <param name='securityConnectorName'>
            /// The security connector name.
            /// </param>
            /// <param name='applicationId'>
            /// The security Application key - unique key for the standard application
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this ISecurityConnectorApplicationOperations operations, string resourceGroupName, string securityConnectorName, string applicationId, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, securityConnectorName, applicationId, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

    }
}
