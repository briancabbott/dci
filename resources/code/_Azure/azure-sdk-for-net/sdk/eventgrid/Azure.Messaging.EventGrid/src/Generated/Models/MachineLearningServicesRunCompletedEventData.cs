// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Messaging.EventGrid.SystemEvents
{
    /// <summary> Schema of the Data property of an EventGridEvent for a Microsoft.MachineLearningServices.RunCompleted event. </summary>
    public partial class MachineLearningServicesRunCompletedEventData
    {
        /// <summary> Initializes a new instance of <see cref="MachineLearningServicesRunCompletedEventData"/>. </summary>
        internal MachineLearningServicesRunCompletedEventData()
        {
        }

        /// <summary> Initializes a new instance of <see cref="MachineLearningServicesRunCompletedEventData"/>. </summary>
        /// <param name="experimentId"> The ID of the experiment that the run belongs to. </param>
        /// <param name="experimentName"> The name of the experiment that the run belongs to. </param>
        /// <param name="runId"> The ID of the Run that was completed. </param>
        /// <param name="runType"> The Run Type of the completed Run. </param>
        /// <param name="runTags"> The tags of the completed Run. </param>
        /// <param name="runProperties"> The properties of the completed Run. </param>
        internal MachineLearningServicesRunCompletedEventData(string experimentId, string experimentName, string runId, string runType, object runTags, object runProperties)
        {
            ExperimentId = experimentId;
            ExperimentName = experimentName;
            RunId = runId;
            RunType = runType;
            RunTags = runTags;
            RunProperties = runProperties;
        }

        /// <summary> The ID of the experiment that the run belongs to. </summary>
        public string ExperimentId { get; }
        /// <summary> The name of the experiment that the run belongs to. </summary>
        public string ExperimentName { get; }
        /// <summary> The ID of the Run that was completed. </summary>
        public string RunId { get; }
        /// <summary> The Run Type of the completed Run. </summary>
        public string RunType { get; }
        /// <summary> The tags of the completed Run. </summary>
        public object RunTags { get; }
        /// <summary> The properties of the completed Run. </summary>
        public object RunProperties { get; }
    }
}
