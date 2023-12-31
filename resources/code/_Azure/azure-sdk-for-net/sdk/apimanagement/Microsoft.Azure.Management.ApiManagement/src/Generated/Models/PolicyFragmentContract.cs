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
    using System.Linq;

    /// <summary>
    /// Policy fragment contract details.
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class PolicyFragmentContract : ProxyResource
    {
        /// <summary>
        /// Initializes a new instance of the PolicyFragmentContract class.
        /// </summary>
        public PolicyFragmentContract()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the PolicyFragmentContract class.
        /// </summary>
        /// <param name="value">Contents of the policy fragment.</param>
        /// <param name="id">Fully qualified resource ID for the resource. Ex -
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</param>
        /// <param name="name">The name of the resource</param>
        /// <param name="type">The type of the resource. E.g.
        /// "Microsoft.Compute/virtualMachines" or
        /// "Microsoft.Storage/storageAccounts"</param>
        /// <param name="description">Policy fragment description.</param>
        /// <param name="format">Format of the policy fragment content.
        /// Possible values include: 'xml', 'rawxml'</param>
        public PolicyFragmentContract(string value, string id = default(string), string name = default(string), string type = default(string), string description = default(string), string format = default(string))
            : base(id, name, type)
        {
            Value = value;
            Description = description;
            Format = format;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets contents of the policy fragment.
        /// </summary>
        [JsonProperty(PropertyName = "properties.value")]
        public string Value { get; set; }

        /// <summary>
        /// Gets or sets policy fragment description.
        /// </summary>
        [JsonProperty(PropertyName = "properties.description")]
        public string Description { get; set; }

        /// <summary>
        /// Gets or sets format of the policy fragment content. Possible values
        /// include: 'xml', 'rawxml'
        /// </summary>
        [JsonProperty(PropertyName = "properties.format")]
        public string Format { get; set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (Value == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "Value");
            }
            if (Description != null)
            {
                if (Description.Length > 1000)
                {
                    throw new ValidationException(ValidationRules.MaxLength, "Description", 1000);
                }
                if (Description.Length < 0)
                {
                    throw new ValidationException(ValidationRules.MinLength, "Description", 0);
                }
            }
        }
    }
}
