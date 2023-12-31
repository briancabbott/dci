// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

namespace Azure.ResourceManager.Network
{
    /// <summary>
    /// A Class representing a VirtualHub along with the instance operations that can be performed on it.
    /// If you have a <see cref="ResourceIdentifier" /> you can construct a <see cref="VirtualHubResource" />
    /// from an instance of <see cref="ArmClient" /> using the GetVirtualHubResource method.
    /// Otherwise you can get one from its parent resource <see cref="ResourceGroupResource" /> using the GetVirtualHub method.
    /// </summary>
    public partial class VirtualHubResource : ArmResource
    {
        /// <summary>
        /// Gets the effective routes configured for the Virtual Hub resource or the specified resource .
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/effectiveRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetEffectiveVirtualHubRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the effective routes for a specific resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual async Task<ArmOperation> GetEffectiveVirtualHubRoutesAsync(WaitUntil waitUntil, EffectiveRoutesContent content = null, CancellationToken cancellationToken = default)
        {
            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetEffectiveVirtualHubRoutes");
            scope.Start();
            try
            {
                var response = await _virtualHubRestClient.GetEffectiveVirtualHubRoutesAsync(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken).ConfigureAwait(false);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetEffectiveVirtualHubRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    await operation.WaitForCompletionResponseAsync(cancellationToken).ConfigureAwait(false);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }

        /// <summary>
        /// Gets the effective routes configured for the Virtual Hub resource or the specified resource .
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/effectiveRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetEffectiveVirtualHubRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the effective routes for a specific resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual ArmOperation GetEffectiveVirtualHubRoutes(WaitUntil waitUntil, EffectiveRoutesContent content = null, CancellationToken cancellationToken = default)
        {
            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetEffectiveVirtualHubRoutes");
            scope.Start();
            try
            {
                var response = _virtualHubRestClient.GetEffectiveVirtualHubRoutes(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetEffectiveVirtualHubRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    operation.WaitForCompletionResponse(cancellationToken);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }

        /// <summary>
        /// Gets the inbound routes configured for the Virtual Hub on a particular connection.
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/inboundRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetInboundRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the inbound routes for a connection resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="content"/> is null. </exception>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual async Task<ArmOperation> GetInboundRoutesAsync(WaitUntil waitUntil, VirtualHubInboundRoutesContent content, CancellationToken cancellationToken = default)
        {
            Argument.AssertNotNull(content, nameof(content));

            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetInboundRoutes");
            scope.Start();
            try
            {
                var response = await _virtualHubRestClient.GetInboundRoutesAsync(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken).ConfigureAwait(false);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetInboundRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    await operation.WaitForCompletionResponseAsync(cancellationToken).ConfigureAwait(false);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }

        /// <summary>
        /// Gets the inbound routes configured for the Virtual Hub on a particular connection.
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/inboundRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetInboundRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the inbound routes for a connection resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="content"/> is null. </exception>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual ArmOperation GetInboundRoutes(WaitUntil waitUntil, VirtualHubInboundRoutesContent content, CancellationToken cancellationToken = default)
        {
            Argument.AssertNotNull(content, nameof(content));

            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetInboundRoutes");
            scope.Start();
            try
            {
                var response = _virtualHubRestClient.GetInboundRoutes(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetInboundRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    operation.WaitForCompletionResponse(cancellationToken);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }

        /// <summary>
        /// Gets the outbound routes configured for the Virtual Hub on a particular connection.
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/outboundRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetOutboundRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the outbound routes for a connection resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="content"/> is null. </exception>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual async Task<ArmOperation> GetOutboundRoutesAsync(WaitUntil waitUntil, VirtualHubOutboundRoutesContent content, CancellationToken cancellationToken = default)
        {
            Argument.AssertNotNull(content, nameof(content));

            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetOutboundRoutes");
            scope.Start();
            try
            {
                var response = await _virtualHubRestClient.GetOutboundRoutesAsync(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken).ConfigureAwait(false);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetOutboundRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    await operation.WaitForCompletionResponseAsync(cancellationToken).ConfigureAwait(false);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }

        /// <summary>
        /// Gets the outbound routes configured for the Virtual Hub on a particular connection.
        /// <list type="bullet">
        /// <item>
        /// <term>Request Path</term>
        /// <description>/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/outboundRoutes</description>
        /// </item>
        /// <item>
        /// <term>Operation Id</term>
        /// <description>VirtualHubs_GetOutboundRoutes</description>
        /// </item>
        /// </list>
        /// </summary>
        /// <param name="waitUntil"> <see cref="WaitUntil.Completed"/> if the method should wait to return until the long-running operation has completed on the service; <see cref="WaitUntil.Started"/> if it should return after starting the operation. For more information on long-running operations, please see <see href="https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/samples/LongRunningOperations.md"> Azure.Core Long-Running Operation samples</see>. </param>
        /// <param name="content"> Parameters supplied to get the outbound routes for a connection resource. </param>
        /// <param name="cancellationToken"> The cancellation token to use. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="content"/> is null. </exception>
        [Obsolete("This method is obsolete and will be removed in a future release", false)]
        [EditorBrowsable(EditorBrowsableState.Never)]
        public virtual ArmOperation GetOutboundRoutes(WaitUntil waitUntil, VirtualHubOutboundRoutesContent content, CancellationToken cancellationToken = default)
        {
            Argument.AssertNotNull(content, nameof(content));

            using var scope = _virtualHubClientDiagnostics.CreateScope("VirtualHubResource.GetOutboundRoutes");
            scope.Start();
            try
            {
                var response = _virtualHubRestClient.GetOutboundRoutes(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content, cancellationToken);
                var operation = new NetworkArmOperation(_virtualHubClientDiagnostics, Pipeline, _virtualHubRestClient.CreateGetOutboundRoutesRequest(Id.SubscriptionId, Id.ResourceGroupName, Id.Name, content).Request, response, OperationFinalStateVia.Location);
                if (waitUntil == WaitUntil.Completed)
                    operation.WaitForCompletionResponse(cancellationToken);
                return operation;
            }
            catch (Exception e)
            {
                scope.Failed(e);
                throw;
            }
        }
    }
}
