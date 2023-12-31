// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.ResourceManager.Network.Models
{
    /// <summary> The protocol for which the DDoS protection policy is being customized. </summary>
    [Obsolete("This struct is obsolete and will be removed in a future release", false)]
    [EditorBrowsable(EditorBrowsableState.Never)]
    public readonly struct DdosCustomPolicyProtocol : IEquatable<DdosCustomPolicyProtocol>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="DdosCustomPolicyProtocol"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public DdosCustomPolicyProtocol(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string TcpValue = "Tcp";
        private const string UdpValue = "Udp";
        private const string SynValue = "Syn";

        /// <summary> Tcp. </summary>
        public static DdosCustomPolicyProtocol Tcp { get; } = new DdosCustomPolicyProtocol(TcpValue);
        /// <summary> Udp. </summary>
        public static DdosCustomPolicyProtocol Udp { get; } = new DdosCustomPolicyProtocol(UdpValue);
        /// <summary> Syn. </summary>
        public static DdosCustomPolicyProtocol Syn { get; } = new DdosCustomPolicyProtocol(SynValue);
        /// <summary> Determines if two <see cref="DdosCustomPolicyProtocol"/> values are the same. </summary>
        public static bool operator ==(DdosCustomPolicyProtocol left, DdosCustomPolicyProtocol right) => left.Equals(right);
        /// <summary> Determines if two <see cref="DdosCustomPolicyProtocol"/> values are not the same. </summary>
        public static bool operator !=(DdosCustomPolicyProtocol left, DdosCustomPolicyProtocol right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="DdosCustomPolicyProtocol"/>. </summary>
        public static implicit operator DdosCustomPolicyProtocol(string value) => new DdosCustomPolicyProtocol(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is DdosCustomPolicyProtocol other && Equals(other);
        /// <inheritdoc />
        public bool Equals(DdosCustomPolicyProtocol other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
