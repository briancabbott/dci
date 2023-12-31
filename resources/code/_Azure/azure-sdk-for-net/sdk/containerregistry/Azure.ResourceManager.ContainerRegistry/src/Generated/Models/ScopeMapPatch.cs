// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.ContainerRegistry.Models
{
    /// <summary> The properties for updating the scope map. </summary>
    public partial class ScopeMapPatch
    {
        /// <summary> Initializes a new instance of <see cref="ScopeMapPatch"/>. </summary>
        public ScopeMapPatch()
        {
            Actions = new ChangeTrackingList<string>();
        }

        /// <summary> Initializes a new instance of <see cref="ScopeMapPatch"/>. </summary>
        /// <param name="description"> The user friendly description of the scope map. </param>
        /// <param name="actions">
        /// The list of scope permissions for registry artifacts.
        /// E.g. repositories/repository-name/pull,
        /// repositories/repository-name/delete
        /// </param>
        internal ScopeMapPatch(string description, IList<string> actions)
        {
            Description = description;
            Actions = actions;
        }

        /// <summary> The user friendly description of the scope map. </summary>
        public string Description { get; set; }
        /// <summary>
        /// The list of scope permissions for registry artifacts.
        /// E.g. repositories/repository-name/pull,
        /// repositories/repository-name/delete
        /// </summary>
        public IList<string> Actions { get; }
    }
}
