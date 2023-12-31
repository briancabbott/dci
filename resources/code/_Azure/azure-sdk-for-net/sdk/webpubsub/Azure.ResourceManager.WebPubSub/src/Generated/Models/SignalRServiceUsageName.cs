// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.WebPubSub.Models
{
    /// <summary> Localizable String object containing the name and a localized value. </summary>
    public partial class SignalRServiceUsageName
    {
        /// <summary> Initializes a new instance of <see cref="SignalRServiceUsageName"/>. </summary>
        internal SignalRServiceUsageName()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SignalRServiceUsageName"/>. </summary>
        /// <param name="value"> The identifier of the usage. </param>
        /// <param name="localizedValue"> Localized name of the usage. </param>
        internal SignalRServiceUsageName(string value, string localizedValue)
        {
            Value = value;
            LocalizedValue = localizedValue;
        }

        /// <summary> The identifier of the usage. </summary>
        public string Value { get; }
        /// <summary> Localized name of the usage. </summary>
        public string LocalizedValue { get; }
    }
}
