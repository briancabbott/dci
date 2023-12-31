// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Reservations
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Collections;
    using System.Collections.Generic;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// QuotaRequestStatusOperations operations.
    /// </summary>
    public partial interface IQuotaRequestStatusOperations
    {
        /// <summary>
        /// For the specified Azure region (location), get the details and
        /// status of the quota request by the quota request ID for the
        /// resources of the resource provider. The PUT request for the quota
        /// (service limit) returns a response with the requestId parameter.
        /// </summary>
        /// <param name='subscriptionId'>
        /// Azure subscription ID.
        /// </param>
        /// <param name='providerId'>
        /// Azure resource provider ID.
        /// </param>
        /// <param name='location'>
        /// Azure region.
        /// </param>
        /// <param name='id'>
        /// Quota Request ID.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="ExceptionResponseException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<QuotaRequestDetails>> GetWithHttpMessagesAsync(string subscriptionId, string providerId, string location, string id, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// For the specified Azure region (location), subscription, and
        /// resource provider, get the history of the quota requests for the
        /// past year. To select specific quota requests, use the oData filter.
        /// </summary>
        /// <param name='subscriptionId'>
        /// Azure subscription ID.
        /// </param>
        /// <param name='providerId'>
        /// Azure resource provider ID.
        /// </param>
        /// <param name='location'>
        /// Azure region.
        /// </param>
        /// <param name='filter'>
        /// | Field | Supported operators |
        /// |---------------------|------------------------|
        /// |requestSubmitTime | ge, le, eq, gt, lt |
        /// </param>
        /// <param name='top'>
        /// Number of records to return.
        /// </param>
        /// <param name='skiptoken'>
        /// Skiptoken is only used if a previous operation returned a partial
        /// result. If a previous response contains a nextLink element, the
        /// value of the nextLink element includes a skiptoken parameter that
        /// specifies a starting point to use for subsequent calls.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="ExceptionResponseException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<IPage<QuotaRequestDetails>>> ListWithHttpMessagesAsync(string subscriptionId, string providerId, string location, string filter = default(string), int? top = default(int?), string skiptoken = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// For the specified Azure region (location), subscription, and
        /// resource provider, get the history of the quota requests for the
        /// past year. To select specific quota requests, use the oData filter.
        /// </summary>
        /// <param name='nextPageLink'>
        /// The NextLink from the previous successful call to List operation.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="ExceptionResponseException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<IPage<QuotaRequestDetails>>> ListNextWithHttpMessagesAsync(string nextPageLink, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
    }
}
