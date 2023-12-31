// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.EventGrid
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for PartnerNamespacesOperations.
    /// </summary>
    public static partial class PartnerNamespacesOperationsExtensions
    {
            /// <summary>
            /// Get a partner namespace.
            /// </summary>
            /// <remarks>
            /// Get properties of a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            public static PartnerNamespace Get(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName)
            {
                return operations.GetAsync(resourceGroupName, partnerNamespaceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get a partner namespace.
            /// </summary>
            /// <remarks>
            /// Get properties of a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespace> GetAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Create a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously creates a new partner namespace with the specified
            /// parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceInfo'>
            /// PartnerNamespace information.
            /// </param>
            public static PartnerNamespace CreateOrUpdate(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespace partnerNamespaceInfo)
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceInfo).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Create a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously creates a new partner namespace with the specified
            /// parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceInfo'>
            /// PartnerNamespace information.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespace> CreateOrUpdateAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespace partnerNamespaceInfo, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceInfo, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Delete a partner namespace.
            /// </summary>
            /// <remarks>
            /// Delete existing partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            public static void Delete(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName)
            {
                operations.DeleteAsync(resourceGroupName, partnerNamespaceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Delete a partner namespace.
            /// </summary>
            /// <remarks>
            /// Delete existing partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Update a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously updates a partner namespace with the specified parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceUpdateParameters'>
            /// Partner namespace update information.
            /// </param>
            public static PartnerNamespace Update(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespaceUpdateParameters partnerNamespaceUpdateParameters)
            {
                return operations.UpdateAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Update a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously updates a partner namespace with the specified parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceUpdateParameters'>
            /// Partner namespace update information.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespace> UpdateAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespaceUpdateParameters partnerNamespaceUpdateParameters, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.UpdateWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List partner namespaces under an Azure subscription.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under an Azure subscription.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='filter'>
            /// The query used to filter the search results using OData syntax. Filtering
            /// is permitted on the 'name' property only and with limited number of OData
            /// operations. These operations are: the 'contains' function as well as the
            /// following logical operations: not, and, or, eq (for equal), and ne (for not
            /// equal). No arithmetic operations are supported. The following is a valid
            /// filter example: $filter=contains(namE, 'PATTERN') and name ne 'PATTERN-1'.
            /// The following is not a valid filter example: $filter=location eq 'westus'.
            /// </param>
            /// <param name='top'>
            /// The number of results to return per page for the list operation. Valid
            /// range for top parameter is 1 to 100. If not specified, the default number
            /// of results to be returned is 20 items per page.
            /// </param>
            public static IPage<PartnerNamespace> ListBySubscription(this IPartnerNamespacesOperations operations, string filter = default(string), int? top = default(int?))
            {
                return operations.ListBySubscriptionAsync(filter, top).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List partner namespaces under an Azure subscription.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under an Azure subscription.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='filter'>
            /// The query used to filter the search results using OData syntax. Filtering
            /// is permitted on the 'name' property only and with limited number of OData
            /// operations. These operations are: the 'contains' function as well as the
            /// following logical operations: not, and, or, eq (for equal), and ne (for not
            /// equal). No arithmetic operations are supported. The following is a valid
            /// filter example: $filter=contains(namE, 'PATTERN') and name ne 'PATTERN-1'.
            /// The following is not a valid filter example: $filter=location eq 'westus'.
            /// </param>
            /// <param name='top'>
            /// The number of results to return per page for the list operation. Valid
            /// range for top parameter is 1 to 100. If not specified, the default number
            /// of results to be returned is 20 items per page.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<PartnerNamespace>> ListBySubscriptionAsync(this IPartnerNamespacesOperations operations, string filter = default(string), int? top = default(int?), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListBySubscriptionWithHttpMessagesAsync(filter, top, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List partner namespaces under a resource group.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under a resource group.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='filter'>
            /// The query used to filter the search results using OData syntax. Filtering
            /// is permitted on the 'name' property only and with limited number of OData
            /// operations. These operations are: the 'contains' function as well as the
            /// following logical operations: not, and, or, eq (for equal), and ne (for not
            /// equal). No arithmetic operations are supported. The following is a valid
            /// filter example: $filter=contains(namE, 'PATTERN') and name ne 'PATTERN-1'.
            /// The following is not a valid filter example: $filter=location eq 'westus'.
            /// </param>
            /// <param name='top'>
            /// The number of results to return per page for the list operation. Valid
            /// range for top parameter is 1 to 100. If not specified, the default number
            /// of results to be returned is 20 items per page.
            /// </param>
            public static IPage<PartnerNamespace> ListByResourceGroup(this IPartnerNamespacesOperations operations, string resourceGroupName, string filter = default(string), int? top = default(int?))
            {
                return operations.ListByResourceGroupAsync(resourceGroupName, filter, top).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List partner namespaces under a resource group.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under a resource group.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='filter'>
            /// The query used to filter the search results using OData syntax. Filtering
            /// is permitted on the 'name' property only and with limited number of OData
            /// operations. These operations are: the 'contains' function as well as the
            /// following logical operations: not, and, or, eq (for equal), and ne (for not
            /// equal). No arithmetic operations are supported. The following is a valid
            /// filter example: $filter=contains(namE, 'PATTERN') and name ne 'PATTERN-1'.
            /// The following is not a valid filter example: $filter=location eq 'westus'.
            /// </param>
            /// <param name='top'>
            /// The number of results to return per page for the list operation. Valid
            /// range for top parameter is 1 to 100. If not specified, the default number
            /// of results to be returned is 20 items per page.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<PartnerNamespace>> ListByResourceGroupAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string filter = default(string), int? top = default(int?), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByResourceGroupWithHttpMessagesAsync(resourceGroupName, filter, top, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List keys for a partner namespace.
            /// </summary>
            /// <remarks>
            /// List the two keys used to publish to a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            public static PartnerNamespaceSharedAccessKeys ListSharedAccessKeys(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName)
            {
                return operations.ListSharedAccessKeysAsync(resourceGroupName, partnerNamespaceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List keys for a partner namespace.
            /// </summary>
            /// <remarks>
            /// List the two keys used to publish to a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespaceSharedAccessKeys> ListSharedAccessKeysAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListSharedAccessKeysWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Regenerate key for a partner namespace.
            /// </summary>
            /// <remarks>
            /// Regenerate a shared access key for a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='keyName'>
            /// Key name to regenerate (key1 or key2).
            /// </param>
            public static PartnerNamespaceSharedAccessKeys RegenerateKey(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, string keyName)
            {
                return operations.RegenerateKeyAsync(resourceGroupName, partnerNamespaceName, keyName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Regenerate key for a partner namespace.
            /// </summary>
            /// <remarks>
            /// Regenerate a shared access key for a partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='keyName'>
            /// Key name to regenerate (key1 or key2).
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespaceSharedAccessKeys> RegenerateKeyAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, string keyName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.RegenerateKeyWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, keyName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Create a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously creates a new partner namespace with the specified
            /// parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceInfo'>
            /// PartnerNamespace information.
            /// </param>
            public static PartnerNamespace BeginCreateOrUpdate(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespace partnerNamespaceInfo)
            {
                return operations.BeginCreateOrUpdateAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceInfo).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Create a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously creates a new partner namespace with the specified
            /// parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceInfo'>
            /// PartnerNamespace information.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespace> BeginCreateOrUpdateAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespace partnerNamespaceInfo, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.BeginCreateOrUpdateWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceInfo, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Delete a partner namespace.
            /// </summary>
            /// <remarks>
            /// Delete existing partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            public static void BeginDelete(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName)
            {
                operations.BeginDeleteAsync(resourceGroupName, partnerNamespaceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Delete a partner namespace.
            /// </summary>
            /// <remarks>
            /// Delete existing partner namespace.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task BeginDeleteAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.BeginDeleteWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Update a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously updates a partner namespace with the specified parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceUpdateParameters'>
            /// Partner namespace update information.
            /// </param>
            public static PartnerNamespace BeginUpdate(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespaceUpdateParameters partnerNamespaceUpdateParameters)
            {
                return operations.BeginUpdateAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Update a partner namespace.
            /// </summary>
            /// <remarks>
            /// Asynchronously updates a partner namespace with the specified parameters.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The name of the resource group within the user's subscription.
            /// </param>
            /// <param name='partnerNamespaceName'>
            /// Name of the partner namespace.
            /// </param>
            /// <param name='partnerNamespaceUpdateParameters'>
            /// Partner namespace update information.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PartnerNamespace> BeginUpdateAsync(this IPartnerNamespacesOperations operations, string resourceGroupName, string partnerNamespaceName, PartnerNamespaceUpdateParameters partnerNamespaceUpdateParameters, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.BeginUpdateWithHttpMessagesAsync(resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List partner namespaces under an Azure subscription.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under an Azure subscription.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<PartnerNamespace> ListBySubscriptionNext(this IPartnerNamespacesOperations operations, string nextPageLink)
            {
                return operations.ListBySubscriptionNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List partner namespaces under an Azure subscription.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under an Azure subscription.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<PartnerNamespace>> ListBySubscriptionNextAsync(this IPartnerNamespacesOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListBySubscriptionNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List partner namespaces under a resource group.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under a resource group.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<PartnerNamespace> ListByResourceGroupNext(this IPartnerNamespacesOperations operations, string nextPageLink)
            {
                return operations.ListByResourceGroupNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List partner namespaces under a resource group.
            /// </summary>
            /// <remarks>
            /// List all the partner namespaces under a resource group.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<PartnerNamespace>> ListByResourceGroupNextAsync(this IPartnerNamespacesOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByResourceGroupNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
