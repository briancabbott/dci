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
    using System.Collections;
    using System.Collections.Generic;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// ChannelsOperations operations.
    /// </summary>
    public partial interface IChannelsOperations
    {
        /// <summary>
        /// Get a channel.
        /// </summary>
        /// <remarks>
        /// Get properties of a channel.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the channel.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<Channel>> GetWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Create or update a channel.
        /// </summary>
        /// <remarks>
        /// Synchronously creates or updates a new channel with the specified
        /// parameters.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the channel.
        /// </param>
        /// <param name='channelInfo'>
        /// Channel information.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<Channel>> CreateOrUpdateWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, Channel channelInfo, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Delete a channel.
        /// </summary>
        /// <remarks>
        /// Delete an existing channel.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the channel.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse> DeleteWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Update a Channel.
        /// </summary>
        /// <remarks>
        /// Synchronously updates a channel with the specified parameters.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the channel.
        /// </param>
        /// <param name='channelUpdateParameters'>
        /// Channel update information.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse> UpdateWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, ChannelUpdateParameters channelUpdateParameters, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// List channels.
        /// </summary>
        /// <remarks>
        /// List all the channels in a partner namespace.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='filter'>
        /// The query used to filter the search results using OData syntax.
        /// Filtering is permitted on the 'name' property only and with limited
        /// number of OData operations. These operations are: the 'contains'
        /// function as well as the following logical operations: not, and, or,
        /// eq (for equal), and ne (for not equal). No arithmetic operations
        /// are supported. The following is a valid filter example:
        /// $filter=contains(namE, 'PATTERN') and name ne 'PATTERN-1'. The
        /// following is not a valid filter example: $filter=location eq
        /// 'westus'.
        /// </param>
        /// <param name='top'>
        /// The number of results to return per page for the list operation.
        /// Valid range for top parameter is 1 to 100. If not specified, the
        /// default number of results to be returned is 20 items per page.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<IPage<Channel>>> ListByPartnerNamespaceWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string filter = default(string), int? top = default(int?), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Get full URL of partner destination channel.
        /// </summary>
        /// <remarks>
        /// Get the full endpoint URL of a partner destination channel.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the Channel.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<EventSubscriptionFullUrl>> GetFullUrlWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Delete a channel.
        /// </summary>
        /// <remarks>
        /// Delete an existing channel.
        /// </remarks>
        /// <param name='resourceGroupName'>
        /// The name of the resource group within the partners subscription.
        /// </param>
        /// <param name='partnerNamespaceName'>
        /// Name of the partner namespace.
        /// </param>
        /// <param name='channelName'>
        /// Name of the channel.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse> BeginDeleteWithHttpMessagesAsync(string resourceGroupName, string partnerNamespaceName, string channelName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// List channels.
        /// </summary>
        /// <remarks>
        /// List all the channels in a partner namespace.
        /// </remarks>
        /// <param name='nextPageLink'>
        /// The NextLink from the previous successful call to List operation.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<IPage<Channel>>> ListByPartnerNamespaceNextWithHttpMessagesAsync(string nextPageLink, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
    }
}
