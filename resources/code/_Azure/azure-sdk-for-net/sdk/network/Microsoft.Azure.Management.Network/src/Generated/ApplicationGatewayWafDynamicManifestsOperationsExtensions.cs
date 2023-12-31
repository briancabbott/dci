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
    /// Extension methods for ApplicationGatewayWafDynamicManifestsOperations.
    /// </summary>
    public static partial class ApplicationGatewayWafDynamicManifestsOperationsExtensions
    {
            /// <summary>
            /// Gets the regional application gateway waf manifest.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='location'>
            /// The region where the nrp are located at.
            /// </param>
            public static IPage<ApplicationGatewayWafDynamicManifestResult> Get(this IApplicationGatewayWafDynamicManifestsOperations operations, string location)
            {
                return operations.GetAsync(location).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the regional application gateway waf manifest.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='location'>
            /// The region where the nrp are located at.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<ApplicationGatewayWafDynamicManifestResult>> GetAsync(this IApplicationGatewayWafDynamicManifestsOperations operations, string location, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(location, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Gets the regional application gateway waf manifest.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<ApplicationGatewayWafDynamicManifestResult> GetNext(this IApplicationGatewayWafDynamicManifestsOperations operations, string nextPageLink)
            {
                return operations.GetNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the regional application gateway waf manifest.
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
            public static async Task<IPage<ApplicationGatewayWafDynamicManifestResult>> GetNextAsync(this IApplicationGatewayWafDynamicManifestsOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
