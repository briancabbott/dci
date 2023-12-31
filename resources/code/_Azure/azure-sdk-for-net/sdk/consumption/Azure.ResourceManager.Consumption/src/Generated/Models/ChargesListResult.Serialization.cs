// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Consumption.Models
{
    internal partial class ChargesListResult
    {
        internal static ChargesListResult DeserializeChargesListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<ConsumptionChargeSummary>> value = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<ConsumptionChargeSummary> array = new List<ConsumptionChargeSummary>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(ConsumptionChargeSummary.DeserializeConsumptionChargeSummary(item));
                    }
                    value = array;
                    continue;
                }
            }
            return new ChargesListResult(Optional.ToList(value));
        }
    }
}
