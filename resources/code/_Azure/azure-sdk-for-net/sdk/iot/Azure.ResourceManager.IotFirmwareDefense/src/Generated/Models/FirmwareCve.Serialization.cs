// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.IotFirmwareDefense.Models
{
    public partial class FirmwareCve
    {
        internal static FirmwareCve DeserializeFirmwareCve(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> cveId = default;
            Optional<BinaryData> component = default;
            Optional<string> severity = default;
            Optional<string> name = default;
            Optional<string> cvssScore = default;
            Optional<string> cvssVersion = default;
            Optional<string> cvssV2Score = default;
            Optional<string> cvssV3Score = default;
            Optional<DateTimeOffset> publishDate = default;
            Optional<DateTimeOffset> updatedDate = default;
            Optional<IReadOnlyList<CveLink>> links = default;
            Optional<string> description = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("cveId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        cveId = null;
                        continue;
                    }
                    cveId = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("component"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    component = BinaryData.FromString(property.Value.GetRawText());
                    continue;
                }
                if (property.NameEquals("severity"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        severity = null;
                        continue;
                    }
                    severity = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("name"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        name = null;
                        continue;
                    }
                    name = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("cvssScore"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        cvssScore = null;
                        continue;
                    }
                    cvssScore = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("cvssVersion"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        cvssVersion = null;
                        continue;
                    }
                    cvssVersion = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("cvssV2Score"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        cvssV2Score = null;
                        continue;
                    }
                    cvssV2Score = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("cvssV3Score"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        cvssV3Score = null;
                        continue;
                    }
                    cvssV3Score = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("publishDate"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    publishDate = property.Value.GetDateTimeOffset("O");
                    continue;
                }
                if (property.NameEquals("updatedDate"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    updatedDate = property.Value.GetDateTimeOffset("O");
                    continue;
                }
                if (property.NameEquals("links"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<CveLink> array = new List<CveLink>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(CveLink.DeserializeCveLink(item));
                    }
                    links = array;
                    continue;
                }
                if (property.NameEquals("description"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        description = null;
                        continue;
                    }
                    description = property.Value.GetString();
                    continue;
                }
            }
            return new FirmwareCve(cveId.Value, component.Value, severity.Value, name.Value, cvssScore.Value, cvssVersion.Value, cvssV2Score.Value, cvssV3Score.Value, Optional.ToNullable(publishDate), Optional.ToNullable(updatedDate), Optional.ToList(links), description.Value);
        }
    }
}
