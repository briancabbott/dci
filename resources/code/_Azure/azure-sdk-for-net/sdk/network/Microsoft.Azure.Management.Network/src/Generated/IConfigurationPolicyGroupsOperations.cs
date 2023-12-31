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
    using System.Collections;
    using System.Collections.Generic;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// ConfigurationPolicyGroupsOperations operations.
    /// </summary>
    public partial interface IConfigurationPolicyGroupsOperations
    {
        /// <summary>
        /// Creates a ConfigurationPolicyGroup if it doesn't exist else updates
        /// the existing one.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
        /// </param>
        /// <param name='configurationPolicyGroupName'>
        /// The name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationPolicyGroupParameters'>
        /// Parameters supplied to create or update a VpnServerConfiguration
        /// PolicyGroup.
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
        Task<AzureOperationResponse<VpnServerConfigurationPolicyGroup>> CreateOrUpdateWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, string configurationPolicyGroupName, VpnServerConfigurationPolicyGroup vpnServerConfigurationPolicyGroupParameters, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Deletes a ConfigurationPolicyGroup.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
        /// </param>
        /// <param name='configurationPolicyGroupName'>
        /// The name of the ConfigurationPolicyGroup.
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
        Task<AzureOperationResponse> DeleteWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, string configurationPolicyGroupName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Retrieves the details of a ConfigurationPolicyGroup.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the VpnServerConfiguration.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
        /// </param>
        /// <param name='configurationPolicyGroupName'>
        /// The name of the ConfigurationPolicyGroup being retrieved.
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
        Task<AzureOperationResponse<VpnServerConfigurationPolicyGroup>> GetWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, string configurationPolicyGroupName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Lists all the configurationPolicyGroups in a resource group for a
        /// vpnServerConfiguration.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the VpnServerConfiguration.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
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
        Task<AzureOperationResponse<IPage<VpnServerConfigurationPolicyGroup>>> ListByVpnServerConfigurationWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Creates a ConfigurationPolicyGroup if it doesn't exist else updates
        /// the existing one.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
        /// </param>
        /// <param name='configurationPolicyGroupName'>
        /// The name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationPolicyGroupParameters'>
        /// Parameters supplied to create or update a VpnServerConfiguration
        /// PolicyGroup.
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
        Task<AzureOperationResponse<VpnServerConfigurationPolicyGroup>> BeginCreateOrUpdateWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, string configurationPolicyGroupName, VpnServerConfigurationPolicyGroup vpnServerConfigurationPolicyGroupParameters, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Deletes a ConfigurationPolicyGroup.
        /// </summary>
        /// <param name='resourceGroupName'>
        /// The resource group name of the ConfigurationPolicyGroup.
        /// </param>
        /// <param name='vpnServerConfigurationName'>
        /// The name of the VpnServerConfiguration.
        /// </param>
        /// <param name='configurationPolicyGroupName'>
        /// The name of the ConfigurationPolicyGroup.
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
        Task<AzureOperationResponse> BeginDeleteWithHttpMessagesAsync(string resourceGroupName, string vpnServerConfigurationName, string configurationPolicyGroupName, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Lists all the configurationPolicyGroups in a resource group for a
        /// vpnServerConfiguration.
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
        Task<AzureOperationResponse<IPage<VpnServerConfigurationPolicyGroup>>> ListByVpnServerConfigurationNextWithHttpMessagesAsync(string nextPageLink, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
    }
}
