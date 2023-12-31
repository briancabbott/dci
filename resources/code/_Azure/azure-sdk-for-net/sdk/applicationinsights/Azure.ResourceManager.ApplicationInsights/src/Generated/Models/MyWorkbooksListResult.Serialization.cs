// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;
using Azure.ResourceManager.ApplicationInsights;

namespace Azure.ResourceManager.ApplicationInsights.Models
{
    internal partial class MyWorkbooksListResult
    {
        internal static MyWorkbooksListResult DeserializeMyWorkbooksListResult(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IReadOnlyList<MyWorkbookData>> value = default;
            Optional<string> nextLink = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("value"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<MyWorkbookData> array = new List<MyWorkbookData>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(MyWorkbookData.DeserializeMyWorkbookData(item));
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
            return new MyWorkbooksListResult(Optional.ToList(value), nextLink.Value);
        }
    }
}
