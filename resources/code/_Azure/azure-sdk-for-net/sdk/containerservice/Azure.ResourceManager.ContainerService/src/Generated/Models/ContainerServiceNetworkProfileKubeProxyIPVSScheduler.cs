// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.ResourceManager.ContainerService.Models
{
    /// <summary> IPVS scheduler, for more information please see http://www.linuxvirtualserver.org/docs/scheduling.html. </summary>
    public readonly partial struct ContainerServiceNetworkProfileKubeProxyIPVSScheduler : IEquatable<ContainerServiceNetworkProfileKubeProxyIPVSScheduler>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="ContainerServiceNetworkProfileKubeProxyIPVSScheduler"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public ContainerServiceNetworkProfileKubeProxyIPVSScheduler(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string RoundRobinValue = "RoundRobin";
        private const string LeastConnectionValue = "LeastConnection";

        /// <summary> Round Robin. </summary>
        public static ContainerServiceNetworkProfileKubeProxyIPVSScheduler RoundRobin { get; } = new ContainerServiceNetworkProfileKubeProxyIPVSScheduler(RoundRobinValue);
        /// <summary> Least Connection. </summary>
        public static ContainerServiceNetworkProfileKubeProxyIPVSScheduler LeastConnection { get; } = new ContainerServiceNetworkProfileKubeProxyIPVSScheduler(LeastConnectionValue);
        /// <summary> Determines if two <see cref="ContainerServiceNetworkProfileKubeProxyIPVSScheduler"/> values are the same. </summary>
        public static bool operator ==(ContainerServiceNetworkProfileKubeProxyIPVSScheduler left, ContainerServiceNetworkProfileKubeProxyIPVSScheduler right) => left.Equals(right);
        /// <summary> Determines if two <see cref="ContainerServiceNetworkProfileKubeProxyIPVSScheduler"/> values are not the same. </summary>
        public static bool operator !=(ContainerServiceNetworkProfileKubeProxyIPVSScheduler left, ContainerServiceNetworkProfileKubeProxyIPVSScheduler right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="ContainerServiceNetworkProfileKubeProxyIPVSScheduler"/>. </summary>
        public static implicit operator ContainerServiceNetworkProfileKubeProxyIPVSScheduler(string value) => new ContainerServiceNetworkProfileKubeProxyIPVSScheduler(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is ContainerServiceNetworkProfileKubeProxyIPVSScheduler other && Equals(other);
        /// <inheritdoc />
        public bool Equals(ContainerServiceNetworkProfileKubeProxyIPVSScheduler other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
