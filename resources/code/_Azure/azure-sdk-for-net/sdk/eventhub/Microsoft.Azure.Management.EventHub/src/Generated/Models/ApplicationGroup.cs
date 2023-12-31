// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.EventHub.Models
{
    using Microsoft.Rest;
    using Microsoft.Rest.Serialization;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// The Application Group object
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class ApplicationGroup : ProxyResource
    {
        /// <summary>
        /// Initializes a new instance of the ApplicationGroup class.
        /// </summary>
        public ApplicationGroup()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the ApplicationGroup class.
        /// </summary>
        /// <param name="clientAppGroupIdentifier">The Unique identifier for
        /// application group.Supports SAS(SASKeyName=KeyName) or
        /// AAD(AADAppID=Guid)</param>
        /// <param name="id">Fully qualified resource ID for the resource. Ex -
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</param>
        /// <param name="name">The name of the resource</param>
        /// <param name="type">The type of the resource. E.g.
        /// "Microsoft.EventHub/Namespaces" or
        /// "Microsoft.EventHub/Namespaces/EventHubs"</param>
        /// <param name="location">The geo-location where the resource
        /// lives</param>
        /// <param name="isEnabled">Determines if Application Group is allowed
        /// to create connection with namespace or not. Once the isEnabled is
        /// set to false, all the existing connections of application group
        /// gets dropped and no new connections will be allowed</param>
        /// <param name="policies">List of group policies that define the
        /// behavior of application group. The policies can support resource
        /// governance scenarios such as limiting ingress or egress
        /// traffic.</param>
        /// <param name="systemData">The system meta data relating to this
        /// resource.</param>
        public ApplicationGroup(string clientAppGroupIdentifier, string id = default(string), string name = default(string), string type = default(string), string location = default(string), bool? isEnabled = default(bool?), IList<ApplicationGroupPolicy> policies = default(IList<ApplicationGroupPolicy>), SystemData systemData = default(SystemData))
            : base(id, name, type, location)
        {
            IsEnabled = isEnabled;
            ClientAppGroupIdentifier = clientAppGroupIdentifier;
            Policies = policies;
            SystemData = systemData;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets determines if Application Group is allowed to create
        /// connection with namespace or not. Once the isEnabled is set to
        /// false, all the existing connections of application group gets
        /// dropped and no new connections will be allowed
        /// </summary>
        [JsonProperty(PropertyName = "properties.isEnabled")]
        public bool? IsEnabled { get; set; }

        /// <summary>
        /// Gets or sets the Unique identifier for application group.Supports
        /// SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
        /// </summary>
        [JsonProperty(PropertyName = "properties.clientAppGroupIdentifier")]
        public string ClientAppGroupIdentifier { get; set; }

        /// <summary>
        /// Gets or sets list of group policies that define the behavior of
        /// application group. The policies can support resource governance
        /// scenarios such as limiting ingress or egress traffic.
        /// </summary>
        [JsonProperty(PropertyName = "properties.policies")]
        public IList<ApplicationGroupPolicy> Policies { get; set; }

        /// <summary>
        /// Gets the system meta data relating to this resource.
        /// </summary>
        [JsonProperty(PropertyName = "systemData")]
        public SystemData SystemData { get; private set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (ClientAppGroupIdentifier == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "ClientAppGroupIdentifier");
            }
            if (Policies != null)
            {
                foreach (var element in Policies)
                {
                    if (element != null)
                    {
                        element.Validate();
                    }
                }
            }
        }
    }
}
