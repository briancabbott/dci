// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;

namespace Azure.Messaging.EventGrid.SystemEvents
{
    /// <summary> Schema of the Data property of an EventGridEvent for a Microsoft.MachineLearningServices.DatasetDriftDetected event. </summary>
    public partial class MachineLearningServicesDatasetDriftDetectedEventData
    {
        /// <summary> Initializes a new instance of <see cref="MachineLearningServicesDatasetDriftDetectedEventData"/>. </summary>
        internal MachineLearningServicesDatasetDriftDetectedEventData()
        {
        }

        /// <summary> Initializes a new instance of <see cref="MachineLearningServicesDatasetDriftDetectedEventData"/>. </summary>
        /// <param name="dataDriftId"> The ID of the data drift monitor that triggered the event. </param>
        /// <param name="dataDriftName"> The name of the data drift monitor that triggered the event. </param>
        /// <param name="runId"> The ID of the Run that detected data drift. </param>
        /// <param name="baseDatasetId"> The ID of the base Dataset used to detect drift. </param>
        /// <param name="targetDatasetId"> The ID of the target Dataset used to detect drift. </param>
        /// <param name="driftCoefficient"> The coefficient result that triggered the event. </param>
        /// <param name="startTime"> The start time of the target dataset time series that resulted in drift detection. </param>
        /// <param name="endTime"> The end time of the target dataset time series that resulted in drift detection. </param>
        internal MachineLearningServicesDatasetDriftDetectedEventData(string dataDriftId, string dataDriftName, string runId, string baseDatasetId, string targetDatasetId, double? driftCoefficient, DateTimeOffset? startTime, DateTimeOffset? endTime)
        {
            DataDriftId = dataDriftId;
            DataDriftName = dataDriftName;
            RunId = runId;
            BaseDatasetId = baseDatasetId;
            TargetDatasetId = targetDatasetId;
            DriftCoefficient = driftCoefficient;
            StartTime = startTime;
            EndTime = endTime;
        }

        /// <summary> The ID of the data drift monitor that triggered the event. </summary>
        public string DataDriftId { get; }
        /// <summary> The name of the data drift monitor that triggered the event. </summary>
        public string DataDriftName { get; }
        /// <summary> The ID of the Run that detected data drift. </summary>
        public string RunId { get; }
        /// <summary> The ID of the base Dataset used to detect drift. </summary>
        public string BaseDatasetId { get; }
        /// <summary> The ID of the target Dataset used to detect drift. </summary>
        public string TargetDatasetId { get; }
        /// <summary> The coefficient result that triggered the event. </summary>
        public double? DriftCoefficient { get; }
        /// <summary> The start time of the target dataset time series that resulted in drift detection. </summary>
        public DateTimeOffset? StartTime { get; }
        /// <summary> The end time of the target dataset time series that resulted in drift detection. </summary>
        public DateTimeOffset? EndTime { get; }
    }
}
