// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure;

namespace Azure.Communication.JobRouter
{
    public partial class WorkerSelectorAttachment
    {
        internal static WorkerSelectorAttachment DeserializeWorkerSelectorAttachment(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            if (element.TryGetProperty("kind", out JsonElement discriminator))
            {
                switch (discriminator.GetString())
                {
                    case "conditional": return ConditionalWorkerSelectorAttachment.DeserializeConditionalWorkerSelectorAttachment(element);
                    case "passThrough": return PassThroughWorkerSelectorAttachment.DeserializePassThroughWorkerSelectorAttachment(element);
                    case "ruleEngine": return RuleEngineWorkerSelectorAttachment.DeserializeRuleEngineWorkerSelectorAttachment(element);
                    case "static": return StaticWorkerSelectorAttachment.DeserializeStaticWorkerSelectorAttachment(element);
                    case "weightedAllocation": return WeightedAllocationWorkerSelectorAttachment.DeserializeWeightedAllocationWorkerSelectorAttachment(element);
                }
            }
            return UnknownWorkerSelectorAttachment.DeserializeUnknownWorkerSelectorAttachment(element);
        }

        /// <summary> Deserializes the model from a raw response. </summary>
        /// <param name="response"> The response to deserialize the model from. </param>
        internal static WorkerSelectorAttachment FromResponse(Response response)
        {
            using var document = JsonDocument.Parse(response.Content);
            return DeserializeWorkerSelectorAttachment(document.RootElement);
        }
    }
}
