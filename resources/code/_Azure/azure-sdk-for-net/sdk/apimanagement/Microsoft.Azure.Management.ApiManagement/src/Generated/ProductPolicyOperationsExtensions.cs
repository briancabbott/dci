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
    /// Extension methods for ProductPolicyOperations.
    /// </summary>
    public static partial class ProductPolicyOperationsExtensions
    {
            /// <summary>
            /// Get the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            public static PolicyCollection ListByProduct(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId)
            {
                return operations.ListByProductAsync(resourceGroupName, serviceName, productId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PolicyCollection> ListByProductAsync(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByProductWithHttpMessagesAsync(resourceGroupName, serviceName, productId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Get the ETag of the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            public static ProductPolicyGetEntityTagHeaders GetEntityTag(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId)
            {
                return operations.GetEntityTagAsync(resourceGroupName, serviceName, productId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the ETag of the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ProductPolicyGetEntityTagHeaders> GetEntityTagAsync(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetEntityTagWithHttpMessagesAsync(resourceGroupName, serviceName, productId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Headers;
                }
            }

            /// <summary>
            /// Get the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='format'>
            /// Policy Export Format. Possible values include: 'xml', 'rawxml'
            /// </param>
            public static PolicyContract Get(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, string format = default(string))
            {
                return operations.GetAsync(resourceGroupName, serviceName, productId, format).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the policy configuration at the Product level.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='format'>
            /// Policy Export Format. Possible values include: 'xml', 'rawxml'
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PolicyContract> GetAsync(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, string format = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, serviceName, productId, format, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates or updates policy configuration for the Product.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The policy contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            public static PolicyContract CreateOrUpdate(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, PolicyContract parameters, string ifMatch = default(string))
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, serviceName, productId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates or updates policy configuration for the Product.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The policy contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<PolicyContract> CreateOrUpdateAsync(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, PolicyContract parameters, string ifMatch = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, productId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes the policy configuration at the Product.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            public static void Delete(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, string ifMatch)
            {
                operations.DeleteAsync(resourceGroupName, serviceName, productId, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes the policy configuration at the Product.
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
            /// <param name='productId'>
            /// Product identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this IProductPolicyOperations operations, string resourceGroupName, string serviceName, string productId, string ifMatch, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, serviceName, productId, ifMatch, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

    }
}
