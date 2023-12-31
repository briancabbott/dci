﻿// < auto - generated >
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Compute.Models
{
    using Microsoft.Rest;
    using Microsoft.Rest.Serialization;
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    public partial class VirtualMachineNetworkInterfaceConfiguration
    {
        public VirtualMachineNetworkInterfaceConfiguration(string name, IList<VirtualMachineNetworkInterfaceIPConfiguration> ipConfigurations, bool? primary, string deleteOption, bool? enableAcceleratedNetworking, bool? enableFpga, bool? enableIPForwarding = default(bool?), SubResource networkSecurityGroup = default(SubResource), VirtualMachineNetworkInterfaceDnsSettingsConfiguration dnsSettings = default(VirtualMachineNetworkInterfaceDnsSettingsConfiguration), SubResource dscpConfiguration = default(SubResource))
        {
            Name = name;
            Primary = primary;
            DeleteOption = deleteOption;
            EnableAcceleratedNetworking = enableAcceleratedNetworking;
            EnableFpga = enableFpga;
            EnableIPForwarding = enableIPForwarding;
            NetworkSecurityGroup = networkSecurityGroup;
            DnsSettings = dnsSettings;
            IpConfigurations = ipConfigurations;
            DscpConfiguration = dscpConfiguration;
            CustomInit();
        }
    }
}