// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.Consumption.Models
{
    internal partial class ReservationRecommendationsListResult
    {
        internal static ReservationRecommendationsListResult DeserializeReservationRecommendationsListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<ConsumptionReservationRecommendation>> value = default;
            Optional<string> nextLink = default;
            Optional<string> previousLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<ConsumptionReservationRecommendation> array = new List<ConsumptionReservationRecommendation>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(ConsumptionReservationRecommendation.DeserializeConsumptionReservationRecommendation(item));
                    }
                    value = array;
                    continue;
                }
                if (property.NameEquals("nextLink"u8))
                {
                    nextLink = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("previousLink"u8))
                {
                    previousLink = property.Value.GetString();
                    continue;
                }
            }
            return new ReservationRecommendationsListResult(Optional.ToList(value), nextLink.Value, previousLink.Value);
        }
    }
}
