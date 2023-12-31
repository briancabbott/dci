// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class ReplicationEligibilityResultProperties
    {
        internal static ReplicationEligibilityResultProperties DeserializeReplicationEligibilityResultProperties(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> clientRequestId = default;
            Optional<IReadOnlyList<ReplicationEligibilityResultErrorInfo>> errors = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("clientRequestId"u8))
                {
                    clientRequestId = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("errors"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<ReplicationEligibilityResultErrorInfo> array = new List<ReplicationEligibilityResultErrorInfo>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(ReplicationEligibilityResultErrorInfo.DeserializeReplicationEligibilityResultErrorInfo(item));
                    }
                    errors = array;
                    continue;
                }
            }
            return new ReplicationEligibilityResultProperties(clientRequestId.Value, Optional.ToList(errors));
        }
    }
}
