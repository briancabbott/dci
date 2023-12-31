// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.DefenderEasm;

namespace Azure.ResourceManager.DefenderEasm.Models
{
    internal partial class EasmLabelListResult
    {
        internal static EasmLabelListResult DeserializeEasmLabelListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<EasmLabelData>> value = default;
            Optional<string> nextLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<EasmLabelData> array = new List<EasmLabelData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(EasmLabelData.DeserializeEasmLabelData(item));
                    }
                    value = array;
                    continue;
                }
                if (property.NameEquals("nextLink"u8))
                {
                    nextLink = property.Value.GetString();
                    continue;
                }
            }
            return new EasmLabelListResult(Optional.ToList(value), nextLink.Value);
        }
    }
}
