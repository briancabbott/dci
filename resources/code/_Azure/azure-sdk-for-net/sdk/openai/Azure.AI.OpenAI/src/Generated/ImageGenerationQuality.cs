// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.ComponentModel;

namespace Azure.AI.OpenAI
{
    /// <summary>
    /// An image generation configuration that specifies how the model should prioritize quality, cost, and speed.
    /// Only configurable with dall-e-3 models.
    /// </summary>
    public readonly partial struct ImageGenerationQuality : IEquatable<ImageGenerationQuality>
    {
        private readonly string _value;

        /// <summary> Initializes a new instance of <see cref="ImageGenerationQuality"/>. </summary>
        /// <exception cref="ArgumentNullException"> <paramref name="value"/> is null. </exception>
        public ImageGenerationQuality(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        private const string StandardValue = "standard";
        private const string HdValue = "hd";

        /// <summary> Requests image generation with standard, balanced characteristics of quality, cost, and speed. </summary>
        public static ImageGenerationQuality Standard { get; } = new ImageGenerationQuality(StandardValue);
        /// <summary> Requests image generation with higher quality, higher cost and lower speed relative to standard. </summary>
        public static ImageGenerationQuality Hd { get; } = new ImageGenerationQuality(HdValue);
        /// <summary> Determines if two <see cref="ImageGenerationQuality"/> values are the same. </summary>
        public static bool operator ==(ImageGenerationQuality left, ImageGenerationQuality right) => left.Equals(right);
        /// <summary> Determines if two <see cref="ImageGenerationQuality"/> values are not the same. </summary>
        public static bool operator !=(ImageGenerationQuality left, ImageGenerationQuality right) => !left.Equals(right);
        /// <summary> Converts a string to a <see cref="ImageGenerationQuality"/>. </summary>
        public static implicit operator ImageGenerationQuality(string value) => new ImageGenerationQuality(value);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object obj) => obj is ImageGenerationQuality other && Equals(other);
        /// <inheritdoc />
        public bool Equals(ImageGenerationQuality other) => string.Equals(_value, other._value, StringComparison.InvariantCultureIgnoreCase);

        /// <inheritdoc />
        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;
        /// <inheritdoc />
        public override string ToString() => _value;
    }
}
