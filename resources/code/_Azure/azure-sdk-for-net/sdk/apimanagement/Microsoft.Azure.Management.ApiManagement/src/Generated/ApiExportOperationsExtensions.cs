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
    /// Extension methods for ApiExportOperations.
    /// </summary>
    public static partial class ApiExportOperationsExtensions
    {
            /// <summary>
            /// Gets the details of the API specified by its identifier in the format
            /// specified to the Storage Blob with SAS Key valid for 5 minutes.
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
            /// <param name='format'>
            /// Format in which to export the Api Details to the Storage Blob with Sas Key
            /// valid for 5 minutes. Possible values include: 'Swagger', 'Wsdl', 'Wadl',
            /// 'Openapi', 'OpenapiJson'
            /// </param>
            public static ApiExportResult Get(this IApiExportOperations operations, string resourceGroupName, string serviceName, string apiId, string format)
            {
                return operations.GetAsync(resourceGroupName, serviceName, apiId, format).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the details of the API specified by its identifier in the format
            /// specified to the Storage Blob with SAS Key valid for 5 minutes.
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
            /// <param name='format'>
            /// Format in which to export the Api Details to the Storage Blob with Sas Key
            /// valid for 5 minutes. Possible values include: 'Swagger', 'Wsdl', 'Wadl',
            /// 'Openapi', 'OpenapiJson'
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ApiExportResult> GetAsync(this IApiExportOperations operations, string resourceGroupName, string serviceName, string apiId, string format, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, serviceName, apiId, format, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
