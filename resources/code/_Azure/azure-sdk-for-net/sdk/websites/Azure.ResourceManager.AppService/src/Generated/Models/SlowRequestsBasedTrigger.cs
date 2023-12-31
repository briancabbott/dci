// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.AppService.Models
{
    /// <summary> Trigger based on request execution time. </summary>
    public partial class SlowRequestsBasedTrigger
    {
        /// <summary> Initializes a new instance of <see cref="SlowRequestsBasedTrigger"/>. </summary>
        public SlowRequestsBasedTrigger()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SlowRequestsBasedTrigger"/>. </summary>
        /// <param name="timeTaken"> Time taken. </param>
        /// <param name="path"> Request Path. </param>
        /// <param name="count"> Request Count. </param>
        /// <param name="timeInterval"> Time interval. </param>
        internal SlowRequestsBasedTrigger(string timeTaken, string path, int? count, string timeInterval)
        {
            TimeTaken = timeTaken;
            Path = path;
            Count = count;
            TimeInterval = timeInterval;
        }

        /// <summary> Time taken. </summary>
        public string TimeTaken { get; set; }
        /// <summary> Request Path. </summary>
        public string Path { get; set; }
        /// <summary> Request Count. </summary>
        public int? Count { get; set; }
        /// <summary> Time interval. </summary>
        public string TimeInterval { get; set; }
    }
}
