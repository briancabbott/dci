// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;
using Azure.ResourceManager.Models;

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> Database query. </summary>
    public partial class ManagedInstanceQuery : ResourceData
    {
        /// <summary> Initializes a new instance of <see cref="ManagedInstanceQuery"/>. </summary>
        public ManagedInstanceQuery()
        {
        }

        /// <summary> Initializes a new instance of <see cref="ManagedInstanceQuery"/>. </summary>
        /// <param name="id"> The id. </param>
        /// <param name="name"> The name. </param>
        /// <param name="resourceType"> The resourceType. </param>
        /// <param name="systemData"> The systemData. </param>
        /// <param name="queryText"> Query text. </param>
        internal ManagedInstanceQuery(ResourceIdentifier id, string name, ResourceType resourceType, SystemData systemData, string queryText) : base(id, name, resourceType, systemData)
        {
            QueryText = queryText;
        }

        /// <summary> Query text. </summary>
        public string QueryText { get; set; }
    }
}
