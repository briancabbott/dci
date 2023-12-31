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
    /// Extension methods for ApiSchemaOperations.
    /// </summary>
    public static partial class ApiSchemaOperationsExtensions
    {
            /// <summary>
            /// Get the schema configuration at the API level.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='filter'>
            /// |     Field     |     Usage     |     Supported operators     |
            /// Supported functions
            /// |&lt;/br&gt;|-------------|-------------|-------------|-------------|&lt;/br&gt;|
            /// contentType | filter | ge, le, eq, ne, gt, lt | substringof, contains,
            /// startswith, endswith |&lt;/br&gt;
            /// </param>
            /// <param name='top'>
            /// Number of records to return.
            /// </param>
            /// <param name='skip'>
            /// Number of records to skip.
            /// </param>
            public static IPage<SchemaContract> ListByApi(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string filter = default(string), int? top = default(int?), int? skip = default(int?))
            {
                return operations.ListByApiAsync(resourceGroupName, serviceName, apiId, filter, top, skip).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the schema configuration at the API level.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='filter'>
            /// |     Field     |     Usage     |     Supported operators     |
            /// Supported functions
            /// |&lt;/br&gt;|-------------|-------------|-------------|-------------|&lt;/br&gt;|
            /// contentType | filter | ge, le, eq, ne, gt, lt | substringof, contains,
            /// startswith, endswith |&lt;/br&gt;
            /// </param>
            /// <param name='top'>
            /// Number of records to return.
            /// </param>
            /// <param name='skip'>
            /// Number of records to skip.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IPage<SchemaContract>> ListByApiAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string filter = default(string), int? top = default(int?), int? skip = default(int?), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByApiWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, filter, top, skip, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Gets the entity state (Etag) version of the schema specified by its
            /// identifier.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            public static ApiSchemaGetEntityTagHeaders GetEntityTag(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId)
            {
                return operations.GetEntityTagAsync(resourceGroupName, serviceName, apiId, schemaId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the entity state (Etag) version of the schema specified by its
            /// identifier.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ApiSchemaGetEntityTagHeaders> GetEntityTagAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetEntityTagWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, schemaId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Headers;
                }
            }

            /// <summary>
            /// Get the schema configuration at the API level.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            public static SchemaContract Get(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId)
            {
                return operations.GetAsync(resourceGroupName, serviceName, apiId, schemaId).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the schema configuration at the API level.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<SchemaContract> GetAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, schemaId, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Creates or updates schema configuration for the API.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The schema contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            public static SchemaContract CreateOrUpdate(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, SchemaContract parameters, string ifMatch = default(string))
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, serviceName, apiId, schemaId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates or updates schema configuration for the API.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The schema contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<SchemaContract> CreateOrUpdateAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, SchemaContract parameters, string ifMatch = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, schemaId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Deletes the schema configuration at the Api.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            /// <param name='force'>
            /// If true removes all references to the schema before deleting it.
            /// </param>
            public static void Delete(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, string ifMatch, bool? force = default(bool?))
            {
                operations.DeleteAsync(resourceGroupName, serviceName, apiId, schemaId, ifMatch, force).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Deletes the schema configuration at the Api.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. ETag should match the current entity state from the
            /// header response of the GET request or it should be * for unconditional
            /// update.
            /// </param>
            /// <param name='force'>
            /// If true removes all references to the schema before deleting it.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, string ifMatch, bool? force = default(bool?), CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, schemaId, ifMatch, force, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

            /// <summary>
            /// Creates or updates schema configuration for the API.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The schema contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            public static SchemaContract BeginCreateOrUpdate(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, SchemaContract parameters, string ifMatch = default(string))
            {
                return operations.BeginCreateOrUpdateAsync(resourceGroupName, serviceName, apiId, schemaId, parameters, ifMatch).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Creates or updates schema configuration for the API.
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
            /// API revision identifier. Must be unique in the current API Management
            /// service instance. Non-current revision has ;rev=n as a suffix where n is
            /// the revision number.
            /// </param>
            /// <param name='schemaId'>
            /// Schema id identifier. Must be unique in the current API Management service
            /// instance.
            /// </param>
            /// <param name='parameters'>
            /// The schema contents to apply.
            /// </param>
            /// <param name='ifMatch'>
            /// ETag of the Entity. Not required when creating an entity, but required when
            /// updating an entity.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<SchemaContract> BeginCreateOrUpdateAsync(this IApiSchemaOperations operations, string resourceGroupName, string serviceName, string apiId, string schemaId, SchemaContract parameters, string ifMatch = default(string), CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.BeginCreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, schemaId, parameters, ifMatch, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Get the schema configuration at the API level.
            /// </summary>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='nextPageLink'>
            /// The NextLink from the previous successful call to List operation.
            /// </param>
            public static IPage<SchemaContract> ListByApiNext(this IApiSchemaOperations operations, string nextPageLink)
            {
                return operations.ListByApiNextAsync(nextPageLink).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get the schema configuration at the API level.
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
            public static async Task<IPage<SchemaContract>> ListByApiNextAsync(this IApiSchemaOperations operations, string nextPageLink, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByApiNextWithHttpMessagesAsync(nextPageLink, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
