// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.RecoveryServicesDataReplication;

namespace Azure.ResourceManager.RecoveryServicesDataReplication.Models
{
    public partial class DataReplicationTask
    {
        internal static DataReplicationTask DeserializeDataReplicationTask(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> taskName = default;
            Optional<DataReplicationTaskState> state = default;
            Optional<DateTimeOffset> startTime = default;
            Optional<DateTimeOffset> endTime = default;
            Optional<TaskModelCustomProperties> customProperties = default;
            Optional<IReadOnlyList<DataReplicationWorkflowData>> childrenWorkflows = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("taskName"u8))
                {
                    taskName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("state"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    state = new DataReplicationTaskState(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("startTime"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    startTime = property.Value.GetDateTimeOffset("O");
                    continue;
                }
                if (property.NameEquals("endTime"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    endTime = property.Value.GetDateTimeOffset("O");
                    continue;
                }
                if (property.NameEquals("customProperties"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    customProperties = TaskModelCustomProperties.DeserializeTaskModelCustomProperties(property.Value);
                    continue;
                }
                if (property.NameEquals("childrenWorkflows"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<DataReplicationWorkflowData> array = new List<DataReplicationWorkflowData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(DataReplicationWorkflowData.DeserializeDataReplicationWorkflowData(item));
                    }
                    childrenWorkflows = array;
                    continue;
                }
            }
            return new DataReplicationTask(taskName.Value, Optional.ToNullable(state), Optional.ToNullable(startTime), Optional.ToNullable(endTime), customProperties.Value, Optional.ToList(childrenWorkflows));
        }
    }
}
