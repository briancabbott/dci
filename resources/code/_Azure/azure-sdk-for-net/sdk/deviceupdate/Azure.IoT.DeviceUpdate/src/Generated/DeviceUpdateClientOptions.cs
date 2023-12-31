// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using Azure.Core;

namespace Azure.IoT.DeviceUpdate
{
    /// <summary> Client options for DeviceUpdate library clients. </summary>
    public partial class DeviceUpdateClientOptions : ClientOptions
    {
        private const ServiceVersion LatestVersion = ServiceVersion.V2022_10_01;

        /// <summary> The version of the service to use. </summary>
        public enum ServiceVersion
        {
            /// <summary> Service version "2022-10-01". </summary>
            V2022_10_01 = 1,
        }

        internal string Version { get; }

        /// <summary> Initializes new instance of DeviceUpdateClientOptions. </summary>
        public DeviceUpdateClientOptions(ServiceVersion version = LatestVersion)
        {
            Version = version switch
            {
                ServiceVersion.V2022_10_01 => "2022-10-01",
                _ => throw new NotSupportedException()
            };
        }
    }
}
