// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.ApiManagement.Models
{
    using Microsoft.Rest;
    using Microsoft.Rest.Serialization;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Logger details.
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class LoggerContract : ProxyResource
    {
        /// <summary>
        /// Initializes a new instance of the LoggerContract class.
        /// </summary>
        public LoggerContract()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the LoggerContract class.
        /// </summary>
        /// <param name="loggerType">Logger type. Possible values include:
        /// 'azureEventHub', 'applicationInsights', 'azureMonitor'</param>
        /// <param name="id">Fully qualified resource ID for the resource. Ex -
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</param>
        /// <param name="name">The name of the resource</param>
        /// <param name="type">The type of the resource. E.g.
        /// "Microsoft.Compute/virtualMachines" or
        /// "Microsoft.Storage/storageAccounts"</param>
        /// <param name="description">Logger description.</param>
        /// <param name="credentials">The name and SendRule connection string
        /// of the event hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.</param>
        /// <param name="isBuffered">Whether records are buffered in the logger
        /// before publishing. Default is assumed to be true.</param>
        /// <param name="resourceId">Azure Resource Id of a log target (either
        /// Azure Event Hub resource or Azure Application Insights
        /// resource).</param>
        public LoggerContract(string loggerType, string id = default(string), string name = default(string), string type = default(string), string description = default(string), IDictionary<string, string> credentials = default(IDictionary<string, string>), bool? isBuffered = default(bool?), string resourceId = default(string))
            : base(id, name, type)
        {
            LoggerType = loggerType;
            Description = description;
            Credentials = credentials;
            IsBuffered = isBuffered;
            ResourceId = resourceId;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets logger type. Possible values include: 'azureEventHub',
        /// 'applicationInsights', 'azureMonitor'
        /// </summary>
        [JsonProperty(PropertyName = "properties.loggerType")]
        public string LoggerType { get; set; }

        /// <summary>
        /// Gets or sets logger description.
        /// </summary>
        [JsonProperty(PropertyName = "properties.description")]
        public string Description { get; set; }

        /// <summary>
        /// Gets or sets the name and SendRule connection string of the event
        /// hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.
        /// </summary>
        [JsonProperty(PropertyName = "properties.credentials")]
        public IDictionary<string, string> Credentials { get; set; }

        /// <summary>
        /// Gets or sets whether records are buffered in the logger before
        /// publishing. Default is assumed to be true.
        /// </summary>
        [JsonProperty(PropertyName = "properties.isBuffered")]
        public bool? IsBuffered { get; set; }

        /// <summary>
        /// Gets or sets azure Resource Id of a log target (either Azure Event
        /// Hub resource or Azure Application Insights resource).
        /// </summary>
        [JsonProperty(PropertyName = "properties.resourceId")]
        public string ResourceId { get; set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (LoggerType == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "LoggerType");
            }
            if (Description != null)
            {
                if (Description.Length > 256)
                {
                    throw new ValidationException(ValidationRules.MaxLength, "Description", 256);
                }
            }
        }
    }
}
