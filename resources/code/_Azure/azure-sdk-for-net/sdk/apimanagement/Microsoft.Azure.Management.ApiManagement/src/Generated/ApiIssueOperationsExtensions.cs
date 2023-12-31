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
    /// Extension methods for ApiIssueOperations.
    /// </summary>
    public static partial class ApiIssueOperationsExtensions
    {
            /// <summary>
            /// Lists all issues associated with the specified API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='odataQuery'>
            /// OData parameters to apply to the operation.
            /// </param>
            /// <param name='expandCommentsAttachments'>
            /// Expand the comment attachments.
            /// </param>
            public static IPage<IssueContract> ListByService(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, ODataQuery<IssueContract> odataQuery = default(ODataQuery<IssueContract>), bool? expandCommentsAttachments = default(bool?))
            {
                return operations.ListByServiceAsync(resourceGroupName, serviceName, apiId, odataQuery, expandCommentsAttachments).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Lists all issues associated with the specified API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='odataQuery'>
            /// OData parameters to apply to the operation.
            /// </param>
            /// <param name='expandCommentsAttachments'>
            /// Expand the comment attachments.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<IssueContract>> ListByServiceAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, ODataQuery<IssueContract> odataQuery = default(ODataQuery<IssueContract>), bool? expandCommentsAttachments = default(bool?), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByServiceWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, odataQuery, expandCommentsAttachments, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Gets the entity state (Etag) version of the Issue for an API specified by
            /// its identifier.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            public static ApiIssueGetEntityTagHeaders GetEntityTag(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId)
            {
                return operations.GetEntityTagAsync(resourceGroupName, serviceName, apiId, issueId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the entity state (Etag) version of the Issue for an API specified by
            /// its identifier.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ApiIssueGetEntityTagHeaders> GetEntityTagAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetEntityTagWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, issueId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Headers;
                }
            }

            /// <summary>
            /// Gets the details of the Issue for an API specified by its identifier.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='expandCommentsAttachments'>
            /// Expand the comment attachments.
            /// </param>
            public static IssueContract Get(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, bool? expandCommentsAttachments = default(bool?))
            {
                return operations.GetAsync(resourceGroupName, serviceName, apiId, issueId, expandCommentsAttachments).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the details of the Issue for an API specified by its identifier.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='expandCommentsAttachments'>
            /// Expand the comment attachments.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IssueContract> GetAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, bool? expandCommentsAttachments = default(bool?), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, issueId, expandCommentsAttachments, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates a new Issue for an API or updates an existing one.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// Create parameters.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            public static IssueContract CreateOrUpdate(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, IssueContract parameters, string ifMatch = default(string))
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, serviceName, apiId, issueId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates a new Issue for an API or updates an existing one.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
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
            public static async Task<IssueContract> CreateOrUpdateAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, IssueContract parameters, string ifMatch = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, issueId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Updates an existing issue for an API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// Update parameters.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            public static IssueContract Update(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, IssueUpdateContract parameters, string ifMatch)
            {
                return operations.UpdateAsync(resourceGroupName, serviceName, apiId, issueId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Updates an existing issue for an API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// Update parameters.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IssueContract> UpdateAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, IssueUpdateContract parameters, string ifMatch, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.UpdateWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, issueId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes the specified Issue from an API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            public static void Delete(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, string ifMatch)
            {
                operations.DeleteAsync(resourceGroupName, serviceName, apiId, issueId, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes the specified Issue from an API.
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
            /// <param name='apiId'>
            /// API identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='issueId'>
            /// Issue identifier. Must be unique in the current API Management service
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
            public static async Task DeleteAsync(this IApiIssueOperations operations, string resourceGroupName, string serviceName, string apiId, string issueId, string ifMatch, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, issueId, ifMatch, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Lists all issues associated with the specified API.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<IssueContract> ListByServiceNext(this IApiIssueOperations operations, string nextPageLink)
            {
                return operations.ListByServiceNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Lists all issues associated with the specified API.
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
            public static async Task<IPage<IssueContract>> ListByServiceNextAsync(this IApiIssueOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByServiceNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
