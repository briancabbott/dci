// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Security
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Collections;
    using System.Collections.Generic;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// GovernanceAssignmentsOperations operations.
    /// </summary>
    public partial interface IGovernanceAssignmentsOperations
    {
        /// <summary>
        /// Get security governanceAssignments on all your resources inside a
        /// scope
        /// </summary>
        /// <param name='scope'>
        /// Scope of the query, can be subscription
        /// (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management
        /// group (/providers/Microsoft.Management/managementGroups/mgName).
        /// </param>
        /// <param name='assessmentName'>
        /// The Assessment Key - Unique key for the assessment type
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
        Task<AzureOperationResponse<IPage<GovernanceAssignment>>> ListWithHttpMessagesAsync(string scope, string assessmentName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Get a specific governanceAssignment for the requested scope by
        /// AssignmentKey
        /// </summary>
        /// <param name='scope'>
        /// Scope of the query, can be subscription
        /// (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management
        /// group (/providers/Microsoft.Management/managementGroups/mgName).
        /// </param>
        /// <param name='assessmentName'>
        /// The Assessment Key - Unique key for the assessment type
        /// </param>
        /// <param name='assignmentKey'>
        /// The security governance assignment key - the assessment key of the
        /// required governance assignment
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
        Task<AzureOperationResponse<GovernanceAssignment>> GetWithHttpMessagesAsync(string scope, string assessmentName, string assignmentKey, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Creates or update a security GovernanceAssignment on the given
        /// subscription.
        /// </summary>
        /// <param name='scope'>
        /// Scope of the query, can be subscription
        /// (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management
        /// group (/providers/Microsoft.Management/managementGroups/mgName).
        /// </param>
        /// <param name='assessmentName'>
        /// The Assessment Key - Unique key for the assessment type
        /// </param>
        /// <param name='assignmentKey'>
        /// The security governance assignment key - the assessment key of the
        /// required governance assignment
        /// </param>
        /// <param name='governanceAssignment'>
        /// GovernanceAssignment over a subscription scope
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
        Task<AzureOperationResponse<GovernanceAssignment>> CreateOrUpdateWithHttpMessagesAsync(string scope, string assessmentName, string assignmentKey, GovernanceAssignment governanceAssignment, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Delete a GovernanceAssignment over a given scope
        /// </summary>
        /// <param name='scope'>
        /// Scope of the query, can be subscription
        /// (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management
        /// group (/providers/Microsoft.Management/managementGroups/mgName).
        /// </param>
        /// <param name='assessmentName'>
        /// The Assessment Key - Unique key for the assessment type
        /// </param>
        /// <param name='assignmentKey'>
        /// The security governance assignment key - the assessment key of the
        /// required governance assignment
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
        Task<AzureOperationResponse> DeleteWithHttpMessagesAsync(string scope, string assessmentName, string assignmentKey, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Get security governanceAssignments on all your resources inside a
        /// scope
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
        /// <exception cref="Microsoft.Rest.Azure.CloudException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<AzureOperationResponse<IPage<GovernanceAssignment>>> ListNextWithHttpMessagesAsync(string nextPageLink, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
    }
}
