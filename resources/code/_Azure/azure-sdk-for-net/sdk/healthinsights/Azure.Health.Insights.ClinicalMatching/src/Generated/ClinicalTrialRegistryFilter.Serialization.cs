// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.Health.Insights.ClinicalMatching
{
    public partial class ClinicalTrialRegistryFilter : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsCollectionDefined(Conditions))
            {
                writer.WritePropertyName("conditions"u8);
                writer.WriteStartArray();
                foreach (var item in Conditions)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(StudyTypes))
            {
                writer.WritePropertyName("studyTypes"u8);
                writer.WriteStartArray();
                foreach (var item in StudyTypes)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(RecruitmentStatuses))
            {
                writer.WritePropertyName("recruitmentStatuses"u8);
                writer.WriteStartArray();
                foreach (var item in RecruitmentStatuses)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(Sponsors))
            {
                writer.WritePropertyName("sponsors"u8);
                writer.WriteStartArray();
                foreach (var item in Sponsors)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(Phases))
            {
                writer.WritePropertyName("phases"u8);
                writer.WriteStartArray();
                foreach (var item in Phases)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(Purposes))
            {
                writer.WritePropertyName("purposes"u8);
                writer.WriteStartArray();
                foreach (var item in Purposes)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(Ids))
            {
                writer.WritePropertyName("ids"u8);
                writer.WriteStartArray();
                foreach (var item in Ids)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(Sources))
            {
                writer.WritePropertyName("sources"u8);
                writer.WriteStartArray();
                foreach (var item in Sources)
                {
                    writer.WriteStringValue(item.ToString());
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(FacilityNames))
            {
                writer.WritePropertyName("facilityNames"u8);
                writer.WriteStartArray();
                foreach (var item in FacilityNames)
                {
                    writer.WriteStringValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(FacilityLocations))
            {
                writer.WritePropertyName("facilityLocations"u8);
                writer.WriteStartArray();
                foreach (var item in FacilityLocations)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            if (Optional.IsCollectionDefined(FacilityAreas))
            {
                writer.WritePropertyName("facilityAreas"u8);
                writer.WriteStartArray();
                foreach (var item in FacilityAreas)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WriteEndObject();
        }

        /// <summary> Convert into a Utf8JsonRequestContent. </summary>
        internal virtual RequestContent ToRequestContent()
        {
            var content = new Utf8JsonRequestContent();
            content.JsonWriter.WriteObjectValue(this);
            return content;
        }
    }
}
