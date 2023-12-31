// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using Azure.Core;
using Azure.Core.GeoJson;

namespace Azure.Maps.Search.Models
{
    /// <summary> Describes the address range on both sides of the street for a search result. Coordinates for the start and end locations of the address range are included. </summary>
    [CodeGenModel("AddressRanges")]
    public partial class AddressRanges
    {
        /// <summary> A location represented as a latitude and longitude using short names &apos;lat&apos; &amp; &apos;lon&apos;. </summary>
        [CodeGenMember("From")]
        internal LatLongPairAbbreviated FromInternal { get; }

        /// <summary> A location represented as a latitude and longitude using short names &apos;lat&apos; &amp; &apos;lon&apos;. </summary>
        [CodeGenMember("To")]
        internal LatLongPairAbbreviated ToInternal { get; }

        /// <summary> A location represented as a latitude and longitude using short names &apos;lat&apos; &amp; &apos;lon&apos;. </summary>
        public GeoPosition From 
        {
            get { return new GeoPosition((double) FromInternal.Lon, (double) FromInternal.Lat); }
        }

        /// <summary> A location represented as a latitude and longitude using short names &apos;lat&apos; &amp; &apos;lon&apos;. </summary>
        public GeoPosition To 
        {
            get { return new GeoPosition((double) ToInternal.Lon, (double) ToInternal.Lat); }
        }
    }
}
