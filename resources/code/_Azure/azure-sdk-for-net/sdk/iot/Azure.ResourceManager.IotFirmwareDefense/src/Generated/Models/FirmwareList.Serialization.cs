// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.IotFirmwareDefense;

namespace Azure.ResourceManager.IotFirmwareDefense.Models
{
    internal partial class FirmwareList
    {
        internal static FirmwareList DeserializeFirmwareList(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<FirmwareData>> value = default;
            Optional<string> nextLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<FirmwareData> array = new List<FirmwareData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(FirmwareData.DeserializeFirmwareData(item));
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
            return new FirmwareList(Optional.ToList(value), nextLink.Value);
        }
    }
}
