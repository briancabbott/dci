// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.HealthcareApis.Models
{
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// The properties indicating the operation result of an operation on a
    /// service.
    /// </summary>
    public partial class OperationResultsDescription
    {
        /// <summary>
        /// Initializes a new instance of the OperationResultsDescription
        /// class.
        /// </summary>
        public OperationResultsDescription()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the OperationResultsDescription
        /// class.
        /// </summary>
        /// <param name="id">The ID of the operation returned.</param>
        /// <param name="name">The name of the operation result.</param>
        /// <param name="status">The status of the operation being performed.
        /// Possible values include: 'Canceled', 'Succeeded', 'Failed',
        /// 'Requested', 'Running'</param>
        /// <param name="startTime">The time that the operation was
        /// started.</param>
        /// <param name="endTime">The time that the operation finished.</param>
        /// <param name="properties">Additional properties of the operation
        /// result.</param>
        public OperationResultsDescription(string id = default(string), string name = default(string), string status = default(string), string startTime = default(string), string endTime = default(string), object properties = default(object))
        {
            Id = id;
            Name = name;
            Status = status;
            StartTime = startTime;
            EndTime = endTime;
            Properties = properties;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets the ID of the operation returned.
        /// </summary>
        [JsonProperty(PropertyName = "id")]
        public string Id { get; private set; }

        /// <summary>
        /// Gets the name of the operation result.
        /// </summary>
        [JsonProperty(PropertyName = "name")]
        public string Name { get; private set; }

        /// <summary>
        /// Gets the status of the operation being performed. Possible values
        /// include: 'Canceled', 'Succeeded', 'Failed', 'Requested', 'Running'
        /// </summary>
        [JsonProperty(PropertyName = "status")]
        public string Status { get; private set; }

        /// <summary>
        /// Gets the time that the operation was started.
        /// </summary>
        [JsonProperty(PropertyName = "startTime")]
        public string StartTime { get; private set; }

        /// <summary>
        /// Gets the time that the operation finished.
        /// </summary>
        [JsonProperty(PropertyName = "endTime")]
        public string EndTime { get; private set; }

        /// <summary>
        /// Gets or sets additional properties of the operation result.
        /// </summary>
        [JsonProperty(PropertyName = "properties")]
        public object Properties { get; set; }

    }
}
