// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Net;
using System.Text.Json;
using Azure.Core;

namespace Azure.ResourceManager.RecoveryServicesSiteRecovery.Models
{
    public partial class VMwareCbtNicDetails
    {
        internal static VMwareCbtNicDetails DeserializeVMwareCbtNicDetails(JsonElement element)
        {
            if (element.ValueKind == JsonValueKind.Null)
            {
                return null;
            }
            Optional<string> nicId = default;
            Optional<string> isPrimaryNic = default;
            Optional<IPAddress> sourceIPAddress = default;
            Optional<SiteRecoveryEthernetAddressType> sourceIPAddressType = default;
            Optional<ResourceIdentifier> sourceNetworkId = default;
            Optional<IPAddress> targetIPAddress = default;
            Optional<SiteRecoveryEthernetAddressType> targetIPAddressType = default;
            Optional<string> targetSubnetName = default;
            Optional<ResourceIdentifier> testNetworkId = default;
            Optional<string> testSubnetName = default;
            Optional<IPAddress> testIPAddress = default;
            Optional<SiteRecoveryEthernetAddressType> testIPAddressType = default;
            Optional<string> targetNicName = default;
            Optional<string> isSelectedForMigration = default;
            foreach (var property in element.EnumerateObject())
            {
                if (property.NameEquals("nicId"u8))
                {
                    nicId = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("isPrimaryNic"u8))
                {
                    isPrimaryNic = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("sourceIPAddress"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    sourceIPAddress = IPAddress.Parse(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("sourceIPAddressType"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    sourceIPAddressType = new SiteRecoveryEthernetAddressType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("sourceNetworkId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    sourceNetworkId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("targetIPAddress"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    targetIPAddress = IPAddress.Parse(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("targetIPAddressType"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    targetIPAddressType = new SiteRecoveryEthernetAddressType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("targetSubnetName"u8))
                {
                    targetSubnetName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("testNetworkId"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    testNetworkId = new ResourceIdentifier(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("testSubnetName"u8))
                {
                    testSubnetName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("testIPAddress"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    testIPAddress = IPAddress.Parse(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("testIPAddressType"u8))
                {
                    if (property.Value.ValueKind == JsonValueKind.Null)
                    {
                        continue;
                    }
                    testIPAddressType = new SiteRecoveryEthernetAddressType(property.Value.GetString());
                    continue;
                }
                if (property.NameEquals("targetNicName"u8))
                {
                    targetNicName = property.Value.GetString();
                    continue;
                }
                if (property.NameEquals("isSelectedForMigration"u8))
                {
                    isSelectedForMigration = property.Value.GetString();
                    continue;
                }
            }
            return new VMwareCbtNicDetails(nicId.Value, isPrimaryNic.Value, sourceIPAddress.Value, Optional.ToNullable(sourceIPAddressType), sourceNetworkId.Value, targetIPAddress.Value, Optional.ToNullable(targetIPAddressType), targetSubnetName.Value, testNetworkId.Value, testSubnetName.Value, testIPAddress.Value, Optional.ToNullable(testIPAddressType), targetNicName.Value, isSelectedForMigration.Value);
        }
    }
}
