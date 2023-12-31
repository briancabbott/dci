// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Text.Json;
using Azure;
using Azure.Core;

namespace Azure.AI.DocumentIntelligence
{
    public partial class QuotaDetails
    {
        internal static QuotaDetails DeserializeQuotaDetails(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            int used = default;
            int quota = default;
            DateTimeOffset quotaResetDateTime = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("used"u8))
                {
                    used = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("quota"u8))
                {
                    quota = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("quotaResetDateTime"u8))
                {
                    quotaResetDateTime = property.Value.GetDateTimeOffset("O");
                    continue;
                }
            }
            return new QuotaDetails(used, quota, quotaResetDateTime);
        }

        /// <summary> Deserializes the model from a raw response. </summary>
        /// <param name="response"> The response to deserialize the model from. </param>
        internal static QuotaDetails FromResponse(Response response)
        {
            using var document = JsonDocument.Parse(response.Content);
            return DeserializeQuotaDetails(document.RootElement);
        }
    }
}
