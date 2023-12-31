// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.MachineLearning;

namespace Azure.ResourceManager.MachineLearning.Models
{
    internal partial class WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult
    {
        internal static WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult DeserializeWorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> nextLink = default;
            Optional<IReadOnlyList<MachineLearningWorkspaceConnectionData>> value = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("nextLink"u8))
                {
                    nextLink = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<MachineLearningWorkspaceConnectionData> array = new List<MachineLearningWorkspaceConnectionData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(MachineLearningWorkspaceConnectionData.DeserializeMachineLearningWorkspaceConnectionData(item));
                    }
                    value = array;
                    continue;
                }
            }
            return new WorkspaceConnectionPropertiesV2BasicResourceArmPaginatedResult(nextLink.Value, Optional.ToList(value));
        }
    }
}
