// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Compute.Models
{
    using Newtonsoft.Json;
    using Newtonsoft.Json.Converters;
    using System.Runtime;
    using System.Runtime.Serialization;

    /// <summary>
    /// Defines values for GalleryApplicationCustomActionParameterType.
    /// </summary>
    [JsonConverter(typeof(StringEnumConverter))]
    public enum GalleryApplicationCustomActionParameterType
    {
        [EnumMember(Value = "String")]
        String,
        [EnumMember(Value = "ConfigurationDataBlob")]
        ConfigurationDataBlob,
        [EnumMember(Value = "LogOutputBlob")]
        LogOutputBlob
    }
    internal static class GalleryApplicationCustomActionParameterTypeEnumExtension
    {
        internal static string ToSerializedValue(this GalleryApplicationCustomActionParameterType? value)
        {
            return value == null ? null : ((GalleryApplicationCustomActionParameterType)value).ToSerializedValue();
        }

        internal static string ToSerializedValue(this GalleryApplicationCustomActionParameterType value)
        {
            switch( value )
            {
                case GalleryApplicationCustomActionParameterType.String:
                    return "String";
                case GalleryApplicationCustomActionParameterType.ConfigurationDataBlob:
                    return "ConfigurationDataBlob";
                case GalleryApplicationCustomActionParameterType.LogOutputBlob:
                    return "LogOutputBlob";
            }
            return null;
        }

        internal static GalleryApplicationCustomActionParameterType? ParseGalleryApplicationCustomActionParameterType(this string value)
        {
            switch( value )
            {
                case "String":
                    return GalleryApplicationCustomActionParameterType.String;
                case "ConfigurationDataBlob":
                    return GalleryApplicationCustomActionParameterType.ConfigurationDataBlob;
                case "LogOutputBlob":
                    return GalleryApplicationCustomActionParameterType.LogOutputBlob;
            }
            return null;
        }
    }
}
