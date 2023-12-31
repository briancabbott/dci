// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Network
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for RouteMapsOperations.
    /// </summary>
    public static partial class RouteMapsOperationsExtensions
    {
            /// <summary>
            /// Retrieves the details of a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            public static RouteMap Get(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName)
            {
                return operations.GetAsync(resourceGroupName, virtualHubName, routeMapName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Retrieves the details of a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<RouteMap> GetAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, virtualHubName, routeMapName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates a RouteMap if it doesn't exist else updates the existing one.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='routeMapParameters'>
            /// Parameters supplied to create or update a RouteMap.
            /// </param>
            public static RouteMap CreateOrUpdate(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, RouteMap routeMapParameters)
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, virtualHubName, routeMapName, routeMapParameters).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates a RouteMap if it doesn't exist else updates the existing one.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='routeMapParameters'>
            /// Parameters supplied to create or update a RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<RouteMap> CreateOrUpdateAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, RouteMap routeMapParameters, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, virtualHubName, routeMapName, routeMapParameters, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            public static void Delete(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName)
            {
                operations.DeleteAsync(resourceGroupName, virtualHubName, routeMapName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, virtualHubName, routeMapName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Retrieves the details of all RouteMaps.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group'.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            public static IPage<RouteMap> List(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName)
            {
                return operations.ListAsync(resourceGroupName, virtualHubName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Retrieves the details of all RouteMaps.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group'.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<RouteMap>> ListAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListWithHttpMessagesAsync(resourceGroupName, virtualHubName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates a RouteMap if it doesn't exist else updates the existing one.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='routeMapParameters'>
            /// Parameters supplied to create or update a RouteMap.
            /// </param>
            public static RouteMap BeginCreateOrUpdate(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, RouteMap routeMapParameters)
            {
                return operations.BeginCreateOrUpdateAsync(resourceGroupName, virtualHubName, routeMapName, routeMapParameters).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates a RouteMap if it doesn't exist else updates the existing one.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='routeMapParameters'>
            /// Parameters supplied to create or update a RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<RouteMap> BeginCreateOrUpdateAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, RouteMap routeMapParameters, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.BeginCreateOrUpdateWithHttpMessagesAsync(resourceGroupName, virtualHubName, routeMapName, routeMapParameters, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            public static void BeginDelete(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName)
            {
                operations.BeginDeleteAsync(resourceGroupName, virtualHubName, routeMapName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes a RouteMap.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// The resource group name of the RouteMap's resource group.
            /// </param>
            /// <param name='virtualHubName'>
            /// The name of the VirtualHub containing the RouteMap.
            /// </param>
            /// <param name='routeMapName'>
            /// The name of the RouteMap.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task BeginDeleteAsync(this IRouteMapsOperations operations, string resourceGroupName, string virtualHubName, string routeMapName, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.BeginDeleteWithHttpMessagesAsync(resourceGroupName, virtualHubName, routeMapName, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Retrieves the details of all RouteMaps.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<RouteMap> ListNext(this IRouteMapsOperations operations, string nextPageLink)
            {
                return operations.ListNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Retrieves the details of all RouteMaps.
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
            public static async Task<IPage<RouteMap>> ListNextAsync(this IRouteMapsOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
