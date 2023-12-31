// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Maintenance
{
    using Microsoft.Rest;
    using Microsoft.Rest.Azure;
    using Models;
    using System.Collections;
    using System.Collections.Generic;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// Extension methods for ConfigurationAssignmentsOperations.
    /// </summary>
    public static partial class ConfigurationAssignmentsOperationsExtensions
    {
            /// <summary>
            /// Get configuration assignment
            /// </summary>
            /// <remarks>
            /// Get configuration assignment for resource..
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            public static ConfigurationAssignment GetParent(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName)
            {
                return operations.GetParentAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get configuration assignment
            /// </summary>
            /// <remarks>
            /// Get configuration assignment for resource..
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> GetParentAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetParentWithHttpMessagesAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Create configuration assignment
            /// </summary>
            /// <remarks>
            /// Register configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='configurationAssignment'>
            /// The configurationAssignment
            /// </param>
            public static ConfigurationAssignment CreateOrUpdateParent(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName, ConfigurationAssignment configurationAssignment)
            {
                return operations.CreateOrUpdateParentAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, configurationAssignment).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Create configuration assignment
            /// </summary>
            /// <remarks>
            /// Register configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='configurationAssignment'>
            /// The configurationAssignment
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> CreateOrUpdateParentAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName, ConfigurationAssignment configurationAssignment, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateParentWithHttpMessagesAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Unregister configuration for resource
            /// </summary>
            /// <remarks>
            /// Unregister configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Unique configuration assignment name
            /// </param>
            public static ConfigurationAssignment DeleteParent(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName)
            {
                return operations.DeleteParentAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Unregister configuration for resource
            /// </summary>
            /// <remarks>
            /// Unregister configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Unique configuration assignment name
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> DeleteParentAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, string configurationAssignmentName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.DeleteParentWithHttpMessagesAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Get configuration assignment
            /// </summary>
            /// <remarks>
            /// Get configuration assignment for resource..
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            public static ConfigurationAssignment Get(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName)
            {
                return operations.GetAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Get configuration assignment
            /// </summary>
            /// <remarks>
            /// Get configuration assignment for resource..
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> GetAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.GetWithHttpMessagesAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Create configuration assignment
            /// </summary>
            /// <remarks>
            /// Register configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='configurationAssignment'>
            /// The configurationAssignment
            /// </param>
            public static ConfigurationAssignment CreateOrUpdate(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName, ConfigurationAssignment configurationAssignment)
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, configurationAssignment).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Create configuration assignment
            /// </summary>
            /// <remarks>
            /// Register configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Configuration assignment name
            /// </param>
            /// <param name='configurationAssignment'>
            /// The configurationAssignment
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> CreateOrUpdateAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName, ConfigurationAssignment configurationAssignment, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Unregister configuration for resource
            /// </summary>
            /// <remarks>
            /// Unregister configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Unique configuration assignment name
            /// </param>
            public static ConfigurationAssignment Delete(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName)
            {
                return operations.DeleteAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Unregister configuration for resource
            /// </summary>
            /// <remarks>
            /// Unregister configuration for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='configurationAssignmentName'>
            /// Unique configuration assignment name
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<ConfigurationAssignment> DeleteAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, string configurationAssignmentName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.DeleteWithHttpMessagesAsync(resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List configurationAssignments for resource
            /// </summary>
            /// <remarks>
            /// List configurationAssignments for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            public static IEnumerable<ConfigurationAssignment> ListParent(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName)
            {
                return operations.ListParentAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List configurationAssignments for resource
            /// </summary>
            /// <remarks>
            /// List configurationAssignments for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceParentType'>
            /// Resource parent type
            /// </param>
            /// <param name='resourceParentName'>
            /// Resource parent identifier
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IEnumerable<ConfigurationAssignment>> ListParentAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceParentType, string resourceParentName, string resourceType, string resourceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListParentWithHttpMessagesAsync(resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// List configurationAssignments for resource
            /// </summary>
            /// <remarks>
            /// List configurationAssignments for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            public static IEnumerable<ConfigurationAssignment> List(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName)
            {
                return operations.ListAsync(resourceGroupName, providerName, resourceType, resourceName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// List configurationAssignments for resource
            /// </summary>
            /// <remarks>
            /// List configurationAssignments for resource.
            /// </remarks>
            /// <param name='operations'>
            /// The operations group for this extension method.
            /// </param>
            /// <param name='resourceGroupName'>
            /// Resource group name
            /// </param>
            /// <param name='providerName'>
            /// Resource provider name
            /// </param>
            /// <param name='resourceType'>
            /// Resource type
            /// </param>
            /// <param name='resourceName'>
            /// Resource identifier
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<IEnumerable<ConfigurationAssignment>> ListAsync(this IConfigurationAssignmentsOperations operations, string resourceGroupName, string providerName, string resourceType, string resourceName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListWithHttpMessagesAsync(resourceGroupName, providerName, resourceType, resourceName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

    }
}
