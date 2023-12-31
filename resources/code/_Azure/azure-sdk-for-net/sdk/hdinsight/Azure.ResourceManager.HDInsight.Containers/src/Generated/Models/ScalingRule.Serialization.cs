// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.HDInsight.Containers.Models
{
    public partial class ScalingRule : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            writer.WritePropertyName("actionType"u8);
            writer.WriteStringValue(ActionType.ToString());
            writer.WritePropertyName("evaluationCount"u8);
            writer.WriteNumberValue(EvaluationCount);
            writer.WritePropertyName("scalingMetric"u8);
            writer.WriteStringValue(ScalingMetric);
            writer.WritePropertyName("comparisonRule"u8);
            writer.WriteObjectValue(ComparisonRule);
            writer.WriteEndObject();
        }

        internal static ScalingRule DeserializeScalingRule(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            ScaleActionType actionType = default;
            int evaluationCount = default;
            string scalingMetric = default;
            HDInsightComparisonRule comparisonRule = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("actionType"u8))
                {
                    actionType = new ScaleActionType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("evaluationCount"u8))
                {
                    evaluationCount = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("scalingMetric"u8))
                {
                    scalingMetric = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("comparisonRule"u8))
                {
                    comparisonRule = HDInsightComparisonRule.DeserializeHDInsightComparisonRule(property.Value);
                    continue;
                }
            }
            return new ScalingRule(actionType, evaluationCount, scalingMetric, comparisonRule);
        }
    }
}
