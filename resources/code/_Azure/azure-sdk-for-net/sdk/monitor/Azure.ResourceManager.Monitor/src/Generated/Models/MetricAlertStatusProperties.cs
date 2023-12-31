// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Monitor.Models
{
    /// <summary> An alert status properties. </summary>
    public partial class MetricAlertStatusProperties
    {
        /// <summary> Initializes a new instance of <see cref="MetricAlertStatusProperties"/>. </summary>
        internal MetricAlertStatusProperties()
        {
            Dimensions = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="MetricAlertStatusProperties"/>. </summary>
        /// <param name="dimensions"> An object describing the type of the dimensions. </param>
        /// <param name="status"> status value. </param>
        /// <param name="timestamp"> UTC time when the status was checked. </param>
        internal MetricAlertStatusProperties(IReadOnlyDictionary<string, string> dimensions, string status, DateTimeOffset? timestamp)
        {
            Dimensions = dimensions;
            Status = status;
            Timestamp = timestamp;
        }

        /// <summary> An object describing the type of the dimensions. </summary>
        public IReadOnlyDictionary<string, string> Dimensions { get; }
        /// <summary> status value. </summary>
        public string Status { get; }
        /// <summary> UTC time when the status was checked. </summary>
        public DateTimeOffset? Timestamp { get; }
    }
}
