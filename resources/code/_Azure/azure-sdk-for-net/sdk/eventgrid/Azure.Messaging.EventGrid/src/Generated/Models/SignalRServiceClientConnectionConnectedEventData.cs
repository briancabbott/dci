// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Messaging.EventGrid.SystemEvents
{
    /// <summary> Schema of the Data property of an EventGridEvent for a Microsoft.SignalRService.ClientConnectionConnected event. </summary>
    public partial class SignalRServiceClientConnectionConnectedEventData
    {
        /// <summary> Initializes a new instance of <see cref="SignalRServiceClientConnectionConnectedEventData"/>. </summary>
        internal SignalRServiceClientConnectionConnectedEventData()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SignalRServiceClientConnectionConnectedEventData"/>. </summary>
        /// <param name="timestamp"> The time at which the event occurred. </param>
        /// <param name="hubName"> The hub of connected client connection. </param>
        /// <param name="connectionId"> The connection Id of connected client connection. </param>
        /// <param name="userId"> The user Id of connected client connection. </param>
        internal SignalRServiceClientConnectionConnectedEventData(DateTimeOffset? timestamp, string hubName, string connectionId, string userId)
        {
            Timestamp = timestamp;
            HubName = hubName;
            ConnectionId = connectionId;
            UserId = userId;
        }

        /// <summary> The time at which the event occurred. </summary>
        public DateTimeOffset? Timestamp { get; }
        /// <summary> The hub of connected client connection. </summary>
        public string HubName { get; }
        /// <summary> The connection Id of connected client connection. </summary>
        public string ConnectionId { get; }
        /// <summary> The user Id of connected client connection. </summary>
        public string UserId { get; }
    }
}
