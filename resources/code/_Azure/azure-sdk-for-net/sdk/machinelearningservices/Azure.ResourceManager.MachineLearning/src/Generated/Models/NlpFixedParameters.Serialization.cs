// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.MachineLearning.Models
{
    public partial class NlpFixedParameters : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsDefined(GradientAccumulationSteps))
            {
                if (GradientAccumulationSteps != null)
                {
                    writer.WritePropertyName("gradientAccumulationSteps"u8);
                    writer.WriteNumberValue(GradientAccumulationSteps.Value);
                }
                else
                {
                    writer.WriteNull("gradientAccumulationSteps");
                }
            }
            if (Optional.IsDefined(LearningRate))
            {
                if (LearningRate != null)
                {
                    writer.WritePropertyName("learningRate"u8);
                    writer.WriteNumberValue(LearningRate.Value);
                }
                else
                {
                    writer.WriteNull("learningRate");
                }
            }
            if (Optional.IsDefined(LearningRateScheduler))
            {
                writer.WritePropertyName("learningRateScheduler"u8);
                writer.WriteStringValue(LearningRateScheduler.Value.ToString());
            }
            if (Optional.IsDefined(ModelName))
            {
                if (ModelName != null)
                {
                    writer.WritePropertyName("modelName"u8);
                    writer.WriteStringValue(ModelName);
                }
                else
                {
                    writer.WriteNull("modelName");
                }
            }
            if (Optional.IsDefined(NumberOfEpochs))
            {
                if (NumberOfEpochs != null)
                {
                    writer.WritePropertyName("numberOfEpochs"u8);
                    writer.WriteNumberValue(NumberOfEpochs.Value);
                }
                else
                {
                    writer.WriteNull("numberOfEpochs");
                }
            }
            if (Optional.IsDefined(TrainingBatchSize))
            {
                if (TrainingBatchSize != null)
                {
                    writer.WritePropertyName("trainingBatchSize"u8);
                    writer.WriteNumberValue(TrainingBatchSize.Value);
                }
                else
                {
                    writer.WriteNull("trainingBatchSize");
                }
            }
            if (Optional.IsDefined(ValidationBatchSize))
            {
                if (ValidationBatchSize != null)
                {
                    writer.WritePropertyName("validationBatchSize"u8);
                    writer.WriteNumberValue(ValidationBatchSize.Value);
                }
                else
                {
                    writer.WriteNull("validationBatchSize");
                }
            }
            if (Optional.IsDefined(WarmupRatio))
            {
                if (WarmupRatio != null)
                {
                    writer.WritePropertyName("warmupRatio"u8);
                    writer.WriteNumberValue(WarmupRatio.Value);
                }
                else
                {
                    writer.WriteNull("warmupRatio");
                }
            }
            if (Optional.IsDefined(WeightDecay))
            {
                if (WeightDecay != null)
                {
                    writer.WritePropertyName("weightDecay"u8);
                    writer.WriteNumberValue(WeightDecay.Value);
                }
                else
                {
                    writer.WriteNull("weightDecay");
                }
            }
            writer.WriteEndObject();
        }

        internal static NlpFixedParameters DeserializeNlpFixedParameters(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<int?> gradientAccumulationSteps = default;
            Optional<float?> learningRate = default;
            Optional<NlpLearningRateScheduler> learningRateScheduler = default;
            Optional<string> modelName = default;
            Optional<int?> numberOfEpochs = default;
            Optional<int?> trainingBatchSize = default;
            Optional<int?> validationBatchSize = default;
            Optional<float?> warmupRatio = default;
            Optional<float?> weightDecay = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("gradientAccumulationSteps"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        gradientAccumulationSteps = null;
                        continue;
                    }
                    gradientAccumulationSteps = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("learningRate"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        learningRate = null;
                        continue;
                    }
                    learningRate = property.Value.GetSingle();
                    continue;
                }
                if (property.NameEquals("learningRateScheduler"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    learningRateScheduler = new NlpLearningRateScheduler(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("modelName"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        modelName = null;
                        continue;
                    }
                    modelName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("numberOfEpochs"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        numberOfEpochs = null;
                        continue;
                    }
                    numberOfEpochs = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("trainingBatchSize"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        trainingBatchSize = null;
                        continue;
                    }
                    trainingBatchSize = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("validationBatchSize"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        validationBatchSize = null;
                        continue;
                    }
                    validationBatchSize = property.Value.GetInt32();
                    continue;
                }
                if (property.NameEquals("warmupRatio"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        warmupRatio = null;
                        continue;
                    }
                    warmupRatio = property.Value.GetSingle();
                    continue;
                }
                if (property.NameEquals("weightDecay"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        weightDecay = null;
                        continue;
                    }
                    weightDecay = property.Value.GetSingle();
                    continue;
                }
            }
            return new NlpFixedParameters(Optional.ToNullable(gradientAccumulationSteps), Optional.ToNullable(learningRate), Optional.ToNullable(learningRateScheduler), modelName.Value, Optional.ToNullable(numberOfEpochs), Optional.ToNullable(trainingBatchSize), Optional.ToNullable(validationBatchSize), Optional.ToNullable(warmupRatio), Optional.ToNullable(weightDecay));
        }
    }
}
