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
    using System.Linq;

    /// <summary>
    /// Operation detail payload
    /// </summary>
    public partial class OperationDetail
    {
        /// <summary>
        /// Initializes a new instance of the OperationDetail class.
        /// </summary>
        public OperationDetail()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the OperationDetail class.
        /// </summary>
        /// <param name="name">Name of the operation</param>
        /// <param name="isDataAction">Indicates whether the operation is a
        /// data action</param>
        /// <param name="display">Display of the operation</param>
        /// <param name="actionType">Enum. Indicates the action type.
        /// "Internal" refers to actions that are for internal only APIs.
        /// Possible values include: 'Internal'</param>
        /// <param name="origin">Origin of the operation</param>
        /// <param name="properties">Properties of the operation</param>
        public OperationDetail(string name = default(string), bool? isDataAction = default(bool?), OperationDisplay display = default(OperationDisplay), string actionType = default(string), string origin = default(string), OperationProperties properties = default(OperationProperties))
        {
            Name = name;
            IsDataAction = isDataAction;
            Display = display;
            ActionType = actionType;
            Origin = origin;
            Properties = properties;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets name of the operation
        /// </summary>
        [JsonProperty(PropertyName = "name")]
        public string Name { get; set; }

        /// <summary>
        /// Gets or sets indicates whether the operation is a data action
        /// </summary>
        [JsonProperty(PropertyName = "isDataAction")]
        public bool? IsDataAction { get; set; }

        /// <summary>
        /// Gets or sets display of the operation
        /// </summary>
        [JsonProperty(PropertyName = "display")]
        public OperationDisplay Display { get; set; }

        /// <summary>
        /// Gets enum. Indicates the action type. "Internal" refers to actions
        /// that are for internal only APIs. Possible values include:
        /// 'Internal'
        /// </summary>
        [JsonProperty(PropertyName = "actionType")]
        public string ActionType { get; private set; }

        /// <summary>
        /// Gets or sets origin of the operation
        /// </summary>
        [JsonProperty(PropertyName = "origin")]
        public string Origin { get; set; }

        /// <summary>
        /// Gets or sets properties of the operation
        /// </summary>
        [JsonProperty(PropertyName = "properties")]
        public OperationProperties Properties { get; set; }

    }
}
