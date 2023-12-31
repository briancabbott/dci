// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesDataReplication.Models
{
    public partial class TestFailoverCleanupWorkflowModelCustomProperties
    {
        internal static TestFailoverCleanupWorkflowModelCustomProperties DeserializeTestFailoverCleanupWorkflowModelCustomProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> comments = default;
            string instanceType = default;
            Optional<IReadOnlyDictionary<string, string>> affectedObjectDetails = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("comments"u8))
                {
                    comments = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("instanceType"u8))
                {
                    instanceType = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("affectedObjectDetails"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    Dictionary<string, string> dictionary = new Dictionary<string, string>();
                    foreach (var property0 in property.Value.EnumerateObject())
                    {
                        dictionary.Add(property0.Name, property0.Value.GetString());
                    }
                    affectedObjectDetails = dictionary;
                    continue;
                }
            }
            return new TestFailoverCleanupWorkflowModelCustomProperties(instanceType, Optional.ToDictionary(affectedObjectDetails), comments.Value);
        }
    }
}
