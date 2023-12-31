// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using Azure.Core;

namespace Azure.ResourceManager.Automation.Models
{
    /// <summary> The parameters supplied to the create or update variable operation. </summary>
    public partial class AutomationVariableCreateOrUpdateContent
    {
        /// <summary> Initializes a new instance of <see cref="AutomationVariableCreateOrUpdateContent"/>. </summary>
        /// <param name="name"> Gets or sets the name of the variable. </param>
        /// <exception cref="ArgumentNullException"> <paramref name="name"/> is null. </exception>
        public AutomationVariableCreateOrUpdateContent(string name)
        {
            Argument.AssertNotNull(name, nameof(name));

            Name = name;
        }

        /// <summary> Initializes a new instance of <see cref="AutomationVariableCreateOrUpdateContent"/>. </summary>
        /// <param name="name"> Gets or sets the name of the variable. </param>
        /// <param name="value"> Gets or sets the value of the variable. </param>
        /// <param name="description"> Gets or sets the description of the variable. </param>
        /// <param name="isEncrypted"> Gets or sets the encrypted flag of the variable. </param>
        internal AutomationVariableCreateOrUpdateContent(string name, string value, string description, bool? isEncrypted)
        {
            Name = name;
            Value = value;
            Description = description;
            IsEncrypted = isEncrypted;
        }

        /// <summary> Gets or sets the name of the variable. </summary>
        public string Name { get; }
        /// <summary> Gets or sets the value of the variable. </summary>
        public string Value { get; set; }
        /// <summary> Gets or sets the description of the variable. </summary>
        public string Description { get; set; }
        /// <summary> Gets or sets the encrypted flag of the variable. </summary>
        public bool? IsEncrypted { get; set; }
    }
}
