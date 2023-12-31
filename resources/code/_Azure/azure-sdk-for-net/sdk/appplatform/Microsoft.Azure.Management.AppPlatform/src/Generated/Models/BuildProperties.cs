// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.AppPlatform.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// Build resource properties payload
    /// </summary>
    public partial class BuildProperties
    {
        /// <summary>
        /// Initializes a new instance of the BuildProperties class.
        /// </summary>
        public BuildProperties()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the BuildProperties class.
        /// </summary>
        /// <param name="relativePath">The relative path of source code</param>
        /// <param name="builder">The resource id of builder to build the
        /// source code</param>
        /// <param name="agentPool">The resource id of agent pool</param>
        /// <param name="provisioningState">Provisioning state of the KPack
        /// build result. Possible values include: 'Creating', 'Updating',
        /// 'Succeeded', 'Failed', 'Deleting'</param>
        /// <param name="env">The environment variables for this build</param>
        /// <param name="triggeredBuildResult"> The build result triggered by
        /// this build</param>
        /// <param name="resourceRequests">The customized build resource for
        /// this build</param>
        public BuildProperties(string relativePath = default(string), string builder = default(string), string agentPool = default(string), string provisioningState = default(string), IDictionary<string, string> env = default(IDictionary<string, string>), TriggeredBuildResult triggeredBuildResult = default(TriggeredBuildResult), BuildResourceRequests resourceRequests = default(BuildResourceRequests))
        {
            RelativePath = relativePath;
            Builder = builder;
            AgentPool = agentPool;
            ProvisioningState = provisioningState;
            Env = env;
            TriggeredBuildResult = triggeredBuildResult;
            ResourceRequests = resourceRequests;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the relative path of source code
        /// </summary>
        [JsonProperty(PropertyName = "relativePath")]
        public string RelativePath { get; set; }

        /// <summary>
        /// Gets or sets the resource id of builder to build the source code
        /// </summary>
        [JsonProperty(PropertyName = "builder")]
        public string Builder { get; set; }

        /// <summary>
        /// Gets or sets the resource id of agent pool
        /// </summary>
        [JsonProperty(PropertyName = "agentPool")]
        public string AgentPool { get; set; }

        /// <summary>
        /// Gets provisioning state of the KPack build result. Possible values
        /// include: 'Creating', 'Updating', 'Succeeded', 'Failed', 'Deleting'
        /// </summary>
        [JsonProperty(PropertyName = "provisioningState")]
        public string ProvisioningState { get; private set; }

        /// <summary>
        /// Gets or sets the environment variables for this build
        /// </summary>
        [JsonProperty(PropertyName = "env")]
        public IDictionary<string, string> Env { get; set; }

        /// <summary>
        /// Gets or sets  The build result triggered by this build
        /// </summary>
        [JsonProperty(PropertyName = "triggeredBuildResult")]
        public TriggeredBuildResult TriggeredBuildResult { get; set; }

        /// <summary>
        /// Gets or sets the customized build resource for this build
        /// </summary>
        [JsonProperty(PropertyName = "resourceRequests")]
        public BuildResourceRequests ResourceRequests { get; set; }

    }
}
