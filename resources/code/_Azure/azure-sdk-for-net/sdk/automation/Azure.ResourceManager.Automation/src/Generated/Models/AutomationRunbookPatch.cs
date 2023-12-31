// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.Automation.Models
{
    /// <summary> The parameters supplied to the update runbook operation. </summary>
    public partial class AutomationRunbookPatch
    {
        /// <summary> Initializes a new instance of <see cref="AutomationRunbookPatch"/>. </summary>
        public AutomationRunbookPatch()
        {
            Tags = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="AutomationRunbookPatch"/>. </summary>
        /// <param name="name"> Gets or sets the name of the resource. </param>
        /// <param name="location"> Gets or sets the location of the resource. </param>
        /// <param name="tags"> Gets or sets the tags attached to the resource. </param>
        /// <param name="description"> Gets or sets the description of the runbook. </param>
        /// <param name="isLogVerboseEnabled"> Gets or sets verbose log option. </param>
        /// <param name="isLogProgressEnabled"> Gets or sets progress log option. </param>
        /// <param name="logActivityTrace"> Gets or sets the activity-level tracing options of the runbook. </param>
        internal AutomationRunbookPatch(string name, AzureLocation? location, IDictionary<string, string> tags, string description, bool? isLogVerboseEnabled, bool? isLogProgressEnabled, int? logActivityTrace)
        {
            Name = name;
            Location = location;
            Tags = tags;
            Description = description;
            IsLogVerboseEnabled = isLogVerboseEnabled;
            IsLogProgressEnabled = isLogProgressEnabled;
            LogActivityTrace = logActivityTrace;
        }

        /// <summary> Gets or sets the name of the resource. </summary>
        public string Name { get; set; }
        /// <summary> Gets or sets the location of the resource. </summary>
        public AzureLocation? Location { get; set; }
        /// <summary> Gets or sets the tags attached to the resource. </summary>
        public IDictionary<string, string> Tags { get; }
        /// <summary> Gets or sets the description of the runbook. </summary>
        public string Description { get; set; }
        /// <summary> Gets or sets verbose log option. </summary>
        public bool? IsLogVerboseEnabled { get; set; }
        /// <summary> Gets or sets progress log option. </summary>
        public bool? IsLogProgressEnabled { get; set; }
        /// <summary> Gets or sets the activity-level tracing options of the runbook. </summary>
        public int? LogActivityTrace { get; set; }
    }
}
