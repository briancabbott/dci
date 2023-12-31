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
    using Microsoft.Rest.Azure.OData;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for AuthorizationOperations.
    /// </summary>
    public static partial class AuthorizationOperationsExtensions
    {
            /// <summary>
            /// Lists a collection of authorization providers defined within a
            /// authorization provider.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='odataQuery'>
            /// OData parameters to apply to the operation.
            /// </param>
            public static IPage<AuthorizationContract> ListByAuthorizationProvider(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, ODataQuery<AuthorizationContract> odataQuery = default(ODataQuery<AuthorizationContract>))
            {
                return operations.ListByAuthorizationProviderAsync(resourceGroupName, serviceName, authorizationProviderId, odataQuery).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Lists a collection of authorization providers defined within a
            /// authorization provider.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='odataQuery'>
            /// OData parameters to apply to the operation.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<AuthorizationContract>> ListByAuthorizationProviderAsync(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, ODataQuery<AuthorizationContract> odataQuery = default(ODataQuery<AuthorizationContract>), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByAuthorizationProviderWithHttpMessagesAsync(resourceGroupName, serviceName, authorizationProviderId, odataQuery, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Gets the details of the authorization specified by its identifier.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            public static AuthorizationContract Get(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId)
            {
                return operations.GetAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the details of the authorization specified by its identifier.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<AuthorizationContract> GetAsync(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates or updates authorization.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='parameters'>
            /// Create parameters.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            public static AuthorizationContract CreateOrUpdate(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, AuthorizationContract parameters, string ifMatch = default(string))
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates or updates authorization.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='parameters'>
            /// Create parameters.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<AuthorizationContract> CreateOrUpdateAsync(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, AuthorizationContract parameters, string ifMatch = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes specific Authorization from the Authorization provider.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            public static void Delete(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, string ifMatch)
            {
                operations.DeleteAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes specific Authorization from the Authorization provider.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, string ifMatch, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, ifMatch, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Confirm valid consent code to suppress Authorizations anti-phishing page.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='parameters'>
            /// Create parameters.
            /// </param>
            public static AuthorizationConfirmConsentCodeHeaders ConfirmConsentCode(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, AuthorizationConfirmConsentCodeRequestContract parameters)
            {
                return operations.ConfirmConsentCodeAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, parameters).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Confirm valid consent code to suppress Authorizations anti-phishing page.
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
            /// <param name='authorizationProviderId'>
            /// Identifier of the authorization provider.
            /// </param>
            /// <param name='authorizationId'>
            /// Identifier of the authorization.
            /// </param>
            /// <param name='parameters'>
            /// Create parameters.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<AuthorizationConfirmConsentCodeHeaders> ConfirmConsentCodeAsync(this IAuthorizationOperations operations, string resourceGroupName, string serviceName, string authorizationProviderId, string authorizationId, AuthorizationConfirmConsentCodeRequestContract parameters, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ConfirmConsentCodeWithHttpMessagesAsync(resourceGroupName, serviceName, authorizationProviderId, authorizationId, parameters, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Headers;
                }
            }

            /// <summary>
            /// Lists a collection of authorization providers defined within a
            /// authorization provider.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<AuthorizationContract> ListByAuthorizationProviderNext(this IAuthorizationOperations operations, string nextPageLink)
            {
                return operations.ListByAuthorizationProviderNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Lists a collection of authorization providers defined within a
            /// authorization provider.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<AuthorizationContract>> ListByAuthorizationProviderNextAsync(this IAuthorizationOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByAuthorizationProviderNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
