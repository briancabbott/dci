// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.AI.MetricsAdvisor.Models
{
    /// <summary> The IngestionStatusQueryOptions. </summary>
    internal partial class IngestionStatusQueryOptions
    {
        /// <summary> Initializes a new instance of <see cref="IngestionStatusQueryOptions"/>. </summary>
        /// <param name="startTime"> the start point of time range to query data ingestion status. </param>
        /// <param name="endTime"> the end point of time range to query data ingestion status. </param>
        public IngestionStatusQueryOptions(DateTimeOffset startTime, DateTimeOffset endTime)
        {
            StartTime = startTime;
            EndTime = endTime;
        }

        /// <summary> the start point of time range to query data ingestion status. </summary>
        public DateTimeOffset StartTime { get; }
        /// <summary> the end point of time range to query data ingestion status. </summary>
        public DateTimeOffset EndTime { get; }
    }
}
