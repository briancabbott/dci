// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.ResourceManager.IotFirmwareDefense.Models
{
    /// <summary> PIE flag. </summary>
    public readonly partial struct PieFlag : IEquatable<PieFlag>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="PieFlag"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public PieFlag(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string TrueValue = "True";
        private const string FalseValue = "False";

        /// <summary> True. </summary>
        public static PieFlag True { get; } = new PieFlag(TrueValue);
        /// <summary> False. </summary>
        public static PieFlag False { get; } = new PieFlag(FalseValue);
        /// <summary> Determines if two <see cref="PieFlag"/> values are the same. </summary>
        public static bool operator ==(PieFlag left, PieFlag right) => left.Equals(right);
        /// <summary> Determines if two <see cref="PieFlag"/> values are not the same. </summary>
        public static bool operator !=(PieFlag left, PieFlag right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="PieFlag"/>. </summary>
        public static implicit operator PieFlag(string value) => new PieFlag(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is PieFlag other && Equals(other);
        /// <inheritdoc />
        public bool Equals(PieFlag other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
