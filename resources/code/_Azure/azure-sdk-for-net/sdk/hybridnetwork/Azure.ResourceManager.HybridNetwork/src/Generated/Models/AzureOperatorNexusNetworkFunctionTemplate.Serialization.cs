// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.HybridNetwork.Models
{
    public partial class AzureOperatorNexusNetworkFunctionTemplate : IUtf8JsonSerializable
    {
        void IUtf8JsonSerializable.Write(Utf8JsonWriter writer)
        {
            writer.WriteStartObject();
            if (Optional.IsCollectionDefined(NetworkFunctionApplications))
            {
                writer.WritePropertyName("networkFunctionApplications"u8);
                writer.WriteStartArray();
                foreach (var item in NetworkFunctionApplications)
                {
                    writer.WriteObjectValue(item);
                }
                writer.WriteEndArray();
            }
            writer.WritePropertyName("nfviType"u8);
            writer.WriteStringValue(NfviType.ToString());
            writer.WriteEndObject();
        }

        internal static AzureOperatorNexusNetworkFunctionTemplate DeserializeAzureOperatorNexusNetworkFunctionTemplate(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<IList<AzureOperatorNexusNetworkFunctionApplication>> networkFunctionApplications = default;
            VirtualNetworkFunctionNfviType nfviType = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("networkFunctionApplications"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    List<AzureOperatorNexusNetworkFunctionApplication> array = new List<AzureOperatorNexusNetworkFunctionApplication>();
                    foreach (var item in property.Value.EnumerateArray())
                    {
                        array.Add(AzureOperatorNexusNetworkFunctionApplication.DeserializeAzureOperatorNexusNetworkFunctionApplication(item));
                    }
                    networkFunctionApplications = array;
                    continue;
                }
                if (property.NameEquals("nfviType"u8))
                {
                    nfviType = new VirtualNetworkFunctionNfviType(property.Value.GetString());
                    continue;
                }
            }
            return new AzureOperatorNexusNetworkFunctionTemplate(nfviType, Optional.ToList(networkFunctionApplications));
        }
    }
}
