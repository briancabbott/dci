// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;
using Azure.ResourceManager.Resources.Models;

namespace Azure.ResourceManager.Compute.Models
{
    /// <summary> The parameters of a capacity reservation Profile. </summary>
    internal partial class CapacityReservationProfile
    {
        /// <summary> Initializes a new instance of <see cref="CapacityReservationProfile"/>. </summary>
        public CapacityReservationProfile()
        {
        }

        /// <summary> Initializes a new instance of <see cref="CapacityReservationProfile"/>. </summary>
        /// <param name="capacityReservationGroup"> Specifies the capacity reservation group resource id that should be used for allocating the virtual machine or scaleset vm instances provided enough capacity has been reserved. Please refer to https://aka.ms/CapacityReservation for more details. </param>
        internal CapacityReservationProfile(WritableSubResource capacityReservationGroup)
        {
            CapacityReservationGroup = capacityReservationGroup;
        }

        /// <summary> Specifies the capacity reservation group resource id that should be used for allocating the virtual machine or scaleset vm instances provided enough capacity has been reserved. Please refer to https://aka.ms/CapacityReservation for more details. </summary>
        internal WritableSubResource CapacityReservationGroup { get; set; }
        /// <summary> Gets or sets Id. </summary>
        public ResourceIdentifier CapacityReservationGroupId
        {
            get => CapacityReservationGroup is null ? default : CapacityReservationGroup.Id;
            set
            {
                if (CapacityReservationGroup is null)
                    CapacityReservationGroup = new WritableSubResource();
                CapacityReservationGroup.Id = value;
            }
        }
    }
}
