// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.Communication.MediaComposition
{
    /// <summary> The scaling mode for the view of a video stream in a cell. </summary>
    public readonly partial struct ScalingMode : IEquatable<ScalingMode>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="ScalingMode"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public ScalingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string StretchValue = "stretch";
        private const string CropValue = "crop";
        private const string FitValue = "fit";

        /// <summary> stretch. </summary>
        public static ScalingMode Stretch { get; } = new ScalingMode(StretchValue);
        /// <summary> crop. </summary>
        public static ScalingMode Crop { get; } = new ScalingMode(CropValue);
        /// <summary> fit. </summary>
        public static ScalingMode Fit { get; } = new ScalingMode(FitValue);
        /// <summary> Determines if two <see cref="ScalingMode"/> values are the same. </summary>
        public static bool operator ==(ScalingMode left, ScalingMode right) => left.Equals(right);
        /// <summary> Determines if two <see cref="ScalingMode"/> values are not the same. </summary>
        public static bool operator !=(ScalingMode left, ScalingMode right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="ScalingMode"/>. </summary>
        public static implicit operator ScalingMode(string value) => new ScalingMode(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is ScalingMode other && Equals(other);
        /// <inheritdoc />
        public bool Equals(ScalingMode other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
