// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.CostManagement.Models
{
    internal partial class BenefitUtilizationSummariesListResult
    {
        internal static BenefitUtilizationSummariesListResult DeserializeBenefitUtilizationSummariesListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<BenefitUtilizationSummary>> value = default;
            Optional<Uri> nextLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<BenefitUtilizationSummary> array = new List<BenefitUtilizationSummary>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(BenefitUtilizationSummary.DeserializeBenefitUtilizationSummary(item));
                    }
                    value = array;
                    continue;
                }
                if (property.NameEquals("nextLink"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    nextLink = new Uri(property.Value.GetString());
                    continue;
                }
            }
            return new BenefitUtilizationSummariesListResult(Optional.ToList(value), nextLink.Value);
        }
    }
}
