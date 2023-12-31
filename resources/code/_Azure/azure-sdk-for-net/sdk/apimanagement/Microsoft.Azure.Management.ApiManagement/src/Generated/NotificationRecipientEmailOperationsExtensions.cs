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
    /// Extension methods for NotificationRecipientEmailOperations.
    /// </summary>
    public static partial class NotificationRecipientEmailOperationsExtensions
    {
            /// <summary>
            /// Gets the list of the Notification Recipient Emails subscribed to a
            /// notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            public static RecipientEmailCollection ListByNotification(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName)
            {
                return operations.ListByNotificationAsync(resourceGroupName, serviceName, notificationName).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Gets the list of the Notification Recipient Emails subscribed to a
            /// notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<RecipientEmailCollection> ListByNotificationAsync(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.ListByNotificationWithHttpMessagesAsync(resourceGroupName, serviceName, notificationName, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Determine if Notification Recipient Email subscribed to the notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            public static bool CheckEntityExists(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email)
            {
                return operations.CheckEntityExistsAsync(resourceGroupName, serviceName, notificationName, email).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Determine if Notification Recipient Email subscribed to the notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<bool> CheckEntityExistsAsync(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CheckEntityExistsWithHttpMessagesAsync(resourceGroupName, serviceName, notificationName, email, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Adds the Email address to the list of Recipients for the Notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            public static RecipientEmailContract CreateOrUpdate(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email)
            {
                return operations.CreateOrUpdateAsync(resourceGroupName, serviceName, notificationName, email).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Adds the Email address to the list of Recipients for the Notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task<RecipientEmailContract> CreateOrUpdateAsync(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email, CancellationToken cancellationToken = default(CancellationToken))
            {
                using (var _result = await operations.CreateOrUpdateWithHttpMessagesAsync(resourceGroupName, serviceName, notificationName, email, null, cancellationToken).ConfigureAwait(false))
                {
                    return _result.Body;
                }
            }

            /// <summary>
            /// Removes the email from the list of Notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            public static void Delete(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email)
            {
                operations.DeleteAsync(resourceGroupName, serviceName, notificationName, email).GetAwaiter().GetResult();
            }

            /// <summary>
            /// Removes the email from the list of Notification.
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
            /// <param name='notificationName'>
            /// Notification Name Identifier. Possible values include:
            /// 'RequestPublisherNotificationMessage',
            /// 'PurchasePublisherNotificationMessage',
            /// 'NewApplicationNotificationMessage', 'BCC',
            /// 'NewIssuePublisherNotificationMessage', 'AccountClosedPublisher',
            /// 'QuotaLimitApproachingPublisherNotificationMessage'
            /// </param>
            /// <param name='email'>
            /// Email identifier.
            /// </param>
            /// <param name='cancellationToken'>
            /// The cancellation token.
            /// </param>
            public static async Task DeleteAsync(this INotificationRecipientEmailOperations operations, string resourceGroupName, string serviceName, string notificationName, string email, CancellationToken cancellationToken = default(CancellationToken))
            {
                (await operations.DeleteWithHttpMessagesAsync(resourceGroupName, serviceName, notificationName, email, null, cancellationToken).ConfigureAwait(false)).Dispose();
            }

    }
}
