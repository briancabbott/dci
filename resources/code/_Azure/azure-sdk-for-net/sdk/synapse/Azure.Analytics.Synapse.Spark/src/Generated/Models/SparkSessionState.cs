// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Analytics.Synapse.Spark.Models
{
    /// <summary> The SparkSessionState. </summary>
    public partial class SparkSessionState
    {
        /// <summary> Initializes a new instance of <see cref="SparkSessionState"/>. </summary>
        internal SparkSessionState()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SparkSessionState"/>. </summary>
        /// <param name="notStartedAt"></param>
        /// <param name="startingAt"></param>
        /// <param name="idleAt"></param>
        /// <param name="deadAt"></param>
        /// <param name="shuttingDownAt"></param>
        /// <param name="terminatedAt"></param>
        /// <param name="recoveringAt"></param>
        /// <param name="busyAt"></param>
        /// <param name="errorAt"></param>
        /// <param name="currentState"></param>
        /// <param name="jobCreationRequest"></param>
        internal SparkSessionState(DateTimeOffset? notStartedAt, DateTimeOffset? startingAt, DateTimeOffset? idleAt, DateTimeOffset? deadAt, DateTimeOffset? shuttingDownAt, DateTimeOffset? terminatedAt, DateTimeOffset? recoveringAt, DateTimeOffset? busyAt, DateTimeOffset? errorAt, string currentState, SparkRequest jobCreationRequest)
        {
            NotStartedAt = notStartedAt;
            StartingAt = startingAt;
            IdleAt = idleAt;
            DeadAt = deadAt;
            ShuttingDownAt = shuttingDownAt;
            TerminatedAt = terminatedAt;
            RecoveringAt = recoveringAt;
            BusyAt = busyAt;
            ErrorAt = errorAt;
            CurrentState = currentState;
            JobCreationRequest = jobCreationRequest;
        }

        /// <summary> Gets the not started at. </summary>
        public DateTimeOffset? NotStartedAt { get; }
        /// <summary> Gets the starting at. </summary>
        public DateTimeOffset? StartingAt { get; }
        /// <summary> Gets the idle at. </summary>
        public DateTimeOffset? IdleAt { get; }
        /// <summary> Gets the dead at. </summary>
        public DateTimeOffset? DeadAt { get; }
        /// <summary> Gets the shutting down at. </summary>
        public DateTimeOffset? ShuttingDownAt { get; }
        /// <summary> Gets the terminated at. </summary>
        public DateTimeOffset? TerminatedAt { get; }
        /// <summary> Gets the recovering at. </summary>
        public DateTimeOffset? RecoveringAt { get; }
        /// <summary> Gets the busy at. </summary>
        public DateTimeOffset? BusyAt { get; }
        /// <summary> Gets the error at. </summary>
        public DateTimeOffset? ErrorAt { get; }
        /// <summary> Gets the current state. </summary>
        public string CurrentState { get; }
        /// <summary> Gets the job creation request. </summary>
        public SparkRequest JobCreationRequest { get; }
    }
}
