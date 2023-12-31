// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.ResourceManager.Chaos.Models
{
    /// <summary> Enum of the Target reference type. </summary>
    public readonly partial struct TargetReferenceType : IEquatable<TargetReferenceType>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="TargetReferenceType"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public TargetReferenceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string ChaosTargetValue = "ChaosTarget";

        /// <summary> ChaosTarget. </summary>
        public static TargetReferenceType ChaosTarget { get; } = new TargetReferenceType(ChaosTargetValue);
        /// <summary> Determines if two <see cref="TargetReferenceType"/> values are the same. </summary>
        public static bool operator ==(TargetReferenceType left, TargetReferenceType right) => left.Equals(right);
        /// <summary> Determines if two <see cref="TargetReferenceType"/> values are not the same. </summary>
        public static bool operator !=(TargetReferenceType left, TargetReferenceType right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="TargetReferenceType"/>. </summary>
        public static implicit operator TargetReferenceType(string value) => new TargetReferenceType(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is TargetReferenceType other && Equals(other);
        /// <inheritdoc />
        public bool Equals(TargetReferenceType other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
