// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.AI.OpenAI
{
    /// <summary>
    /// The authentication options for Azure OpenAI On Your Data.
    /// Please note <see cref="OnYourDataAuthenticationOptions"/> is the base class. According to the scenario, a derived class of the base class might need to be assigned here, or this property needs to be casted to one of the possible derived classes.
    /// The available derived classes include <see cref="OnYourDataApiKeyAuthenticationOptions"/>, <see cref="OnYourDataConnectionStringAuthenticationOptions"/>, <see cref="OnYourDataKeyAndKeyIdAuthenticationOptions"/>, <see cref="OnYourDataSystemAssignedManagedIdentityAuthenticationOptions"/> and <see cref="OnYourDataUserAssignedManagedIdentityAuthenticationOptions"/>.
    /// </summary>
    public abstract partial class OnYourDataAuthenticationOptions
    {
        /// <summary> Initializes a new instance of <see cref="OnYourDataAuthenticationOptions"/>. </summary>
        protected OnYourDataAuthenticationOptions()
        {
        }

        /// <summary> Initializes a new instance of <see cref="OnYourDataAuthenticationOptions"/>. </summary>
        /// <param name="type"> The authentication type. </param>
        internal OnYourDataAuthenticationOptions(OnYourDataAuthenticationType type)
        {
            Type = type;
        }

        /// <summary> The authentication type. </summary>
        internal OnYourDataAuthenticationType Type { get; set; }
    }
}
