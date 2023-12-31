// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.Messaging.EventGrid.SystemEvents
{
    /// <summary> The geofence geometry. </summary>
    public partial class MapsGeofenceGeometry
    {
        /// <summary> Initializes a new instance of <see cref="MapsGeofenceGeometry"/>. </summary>
        internal MapsGeofenceGeometry()
        {
        }

        /// <summary> Initializes a new instance of <see cref="MapsGeofenceGeometry"/>. </summary>
        /// <param name="deviceId"> ID of the device. </param>
        /// <param name="distance"> Distance from the coordinate to the closest border of the geofence. Positive means the coordinate is outside of the geofence. If the coordinate is outside of the geofence, but more than the value of searchBuffer away from the closest geofence border, then the value is 999. Negative means the coordinate is inside of the geofence. If the coordinate is inside the polygon, but more than the value of searchBuffer away from the closest geofencing border,then the value is -999. A value of 999 means that there is great confidence the coordinate is well outside the geofence. A value of -999 means that there is great confidence the coordinate is well within the geofence. </param>
        /// <param name="geometryId"> The unique ID for the geofence geometry. </param>
        /// <param name="nearestLat"> Latitude of the nearest point of the geometry. </param>
        /// <param name="nearestLon"> Longitude of the nearest point of the geometry. </param>
        /// <param name="udId"> The unique id returned from user upload service when uploading a geofence. Will not be included in geofencing post API. </param>
        internal MapsGeofenceGeometry(string deviceId, float? distance, string geometryId, float? nearestLat, float? nearestLon, string udId)
        {
            DeviceId = deviceId;
            Distance = distance;
            GeometryId = geometryId;
            NearestLat = nearestLat;
            NearestLon = nearestLon;
            UdId = udId;
        }

        /// <summary> ID of the device. </summary>
        public string DeviceId { get; }
        /// <summary> Distance from the coordinate to the closest border of the geofence. Positive means the coordinate is outside of the geofence. If the coordinate is outside of the geofence, but more than the value of searchBuffer away from the closest geofence border, then the value is 999. Negative means the coordinate is inside of the geofence. If the coordinate is inside the polygon, but more than the value of searchBuffer away from the closest geofencing border,then the value is -999. A value of 999 means that there is great confidence the coordinate is well outside the geofence. A value of -999 means that there is great confidence the coordinate is well within the geofence. </summary>
        public float? Distance { get; }
        /// <summary> The unique ID for the geofence geometry. </summary>
        public string GeometryId { get; }
        /// <summary> Latitude of the nearest point of the geometry. </summary>
        public float? NearestLat { get; }
        /// <summary> Longitude of the nearest point of the geometry. </summary>
        public float? NearestLon { get; }
        /// <summary> The unique id returned from user upload service when uploading a geofence. Will not be included in geofencing post API. </summary>
        public string UdId { get; }
    }
}
