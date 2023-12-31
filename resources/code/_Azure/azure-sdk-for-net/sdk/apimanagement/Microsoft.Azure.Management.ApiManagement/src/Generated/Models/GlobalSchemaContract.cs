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
    /// Global Schema Contract details.
    /// </summary>
    [Rest.Serialization.JsonTransformation]
    public partial class GlobalSchemaContract : ProxyResource
    {
        /// <summary>
        /// Initializes a new instance of the GlobalSchemaContract class.
        /// </summary>
        public GlobalSchemaContract()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the GlobalSchemaContract class.
        /// </summary>
        /// <param name="schemaType">Schema Type. Immutable. Possible values
        /// include: 'Xml', 'Json'</param>
        /// <param name="id">Fully qualified resource ID for the resource. Ex -
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</param>
        /// <param name="name">The name of the resource</param>
        /// <param name="type">The type of the resource. E.g.
        /// "Microsoft.Compute/virtualMachines" or
        /// "Microsoft.Storage/storageAccounts"</param>
        /// <param name="description">Free-form schema entity
        /// description.</param>
        /// <param name="value">Json-encoded string for non json-based
        /// schema.</param>
        /// <param name="document">Global Schema document object for json-based
        /// schema formats(e.g. json schema).</param>
        public GlobalSchemaContract(string schemaType, string id = default(string), string name = default(string), string type = default(string), string description = default(string), object value = default(object), object document = default(object))
            : base(id, name, type)
        {
            SchemaType = schemaType;
            Description = description;
            Value = value;
            Document = document;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets schema Type. Immutable. Possible values include:
        /// 'Xml', 'Json'
        /// </summary>
        [JsonProperty(PropertyName = "properties.schemaType")]
        public string SchemaType { get; set; }

        /// <summary>
        /// Gets or sets free-form schema entity description.
        /// </summary>
        [JsonProperty(PropertyName = "properties.description")]
        public string Description { get; set; }

        /// <summary>
        /// Gets or sets json-encoded string for non json-based schema.
        /// </summary>
        [JsonProperty(PropertyName = "properties.value")]
        public object Value { get; set; }

        /// <summary>
        /// Gets or sets global Schema document object for json-based schema
        /// formats(e.g. json schema).
        /// </summary>
        [JsonProperty(PropertyName = "properties.document")]
        public object Document { get; set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (SchemaType == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "SchemaType");
            }
        }
    }
}
