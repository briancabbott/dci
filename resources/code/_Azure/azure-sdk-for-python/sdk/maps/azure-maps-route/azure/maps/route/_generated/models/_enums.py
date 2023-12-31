# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class AlternativeRouteType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """AlternativeRouteType."""

    #: Allow any alternative route to be returned irrespective of how it compares to the reference
    #: route in terms of optimality.
    ANY_ROUTE = "anyRoute"
    #: Return an alternative route only if it is better than the reference route according to the
    #: given planning criteria.
    BETTER_ROUTE = "betterRoute"


class ComputeTravelTime(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """ComputeTravelTime."""

    #: Does not compute additional travel times.
    NONE = "none"
    #: Computes travel times for all types of traffic information and specifies all results in the
    #: fields noTrafficTravelTimeInSeconds, historicTrafficTravelTimeInSeconds and
    #: liveTrafficIncidentsTravelTimeInSeconds being included in the summaries in the route response.
    ALL = "all"


class DelayMagnitude(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The magnitude of delay caused by the incident. These values correspond to the values of the
    response field ty of the `Get Traffic Incident Detail API
    <https://docs.microsoft.com/rest/api/maps/traffic/gettrafficincidentdetail>`_.
    """

    #: Unknown.
    UNKNOWN = "0"
    #: Minor.
    MINOR = "1"
    #: Moderate.
    MODERATE = "2"
    #: Major.
    MAJOR = "3"
    #: Undefined, used for road closures and other indefinite delays.
    UNDEFINED = "4"


class DrivingSide(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Indicates left-hand vs. right-hand side driving at the point of the maneuver."""

    #: Left side.
    LEFT = "LEFT"
    #: Right side.
    RIGHT = "RIGHT"


class GeoJsonObjectType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Specifies the ``GeoJSON`` type. Must be one of the nine valid GeoJSON object types - Point,
    MultiPoint, LineString, MultiLineString, Polygon, MultiPolygon, GeometryCollection, Feature and
    FeatureCollection.
    """

    #: ``GeoJSON Point`` geometry.
    GEO_JSON_POINT = "Point"
    #: ``GeoJSON MultiPoint`` geometry.
    GEO_JSON_MULTI_POINT = "MultiPoint"
    #: ``GeoJSON LineString`` geometry.
    GEO_JSON_LINE_STRING = "LineString"
    #: ``GeoJSON MultiLineString`` geometry.
    GEO_JSON_MULTI_LINE_STRING = "MultiLineString"
    #: ``GeoJSON Polygon`` geometry.
    GEO_JSON_POLYGON = "Polygon"
    #: ``GeoJSON MultiPolygon`` geometry.
    GEO_JSON_MULTI_POLYGON = "MultiPolygon"
    #: ``GeoJSON GeometryCollection`` geometry.
    GEO_JSON_GEOMETRY_COLLECTION = "GeometryCollection"
    #: ``GeoJSON Feature`` object.
    GEO_JSON_FEATURE = "Feature"
    #: ``GeoJSON FeatureCollection`` object.
    GEO_JSON_FEATURE_COLLECTION = "FeatureCollection"


class GuidanceInstructionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Type of the instruction, e.g., turn or change of road form."""

    #: Turn.
    TURN = "TURN"
    #: Road Change.
    ROAD_CHANGE = "ROAD_CHANGE"
    #: Departure location.
    LOCATION_DEPARTURE = "LOCATION_DEPARTURE"
    #: Arrival location.
    LOCATION_ARRIVAL = "LOCATION_ARRIVAL"
    #: Direction information.
    DIRECTION_INFO = "DIRECTION_INFO"
    #: Way point location.
    LOCATION_WAYPOINT = "LOCATION_WAYPOINT"


class GuidanceManeuver(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """A code identifying the maneuver."""

    #: You have arrived.
    ARRIVE = "ARRIVE"
    #: You have arrived. Your destination is on the left.
    ARRIVE_LEFT = "ARRIVE_LEFT"
    #: You have arrived. Your destination is on the right.
    ARRIVE_RIGHT = "ARRIVE_RIGHT"
    #: Leave.
    DEPART = "DEPART"
    #: Keep straight on.
    STRAIGHT = "STRAIGHT"
    #: Keep right.
    KEEP_RIGHT = "KEEP_RIGHT"
    #: Bear right.
    BEAR_RIGHT = "BEAR_RIGHT"
    #: Turn right.
    TURN_RIGHT = "TURN_RIGHT"
    #: Turn sharp right.
    SHARP_RIGHT = "SHARP_RIGHT"
    #: Keep left.
    KEEP_LEFT = "KEEP_LEFT"
    #: Bear left.
    BEAR_LEFT = "BEAR_LEFT"
    #: Turn left.
    TURN_LEFT = "TURN_LEFT"
    #: Turn sharp left.
    SHARP_LEFT = "SHARP_LEFT"
    #: Make a U-turn.
    MAKE_U_TURN = "MAKE_UTURN"
    #: Take the motorway.
    ENTER_MOTORWAY = "ENTER_MOTORWAY"
    #: Take the freeway.
    ENTER_FREEWAY = "ENTER_FREEWAY"
    #: Take the highway.
    ENTER_HIGHWAY = "ENTER_HIGHWAY"
    #: Take the exit.
    TAKE_EXIT = "TAKE_EXIT"
    #: Take the left exit.
    MOTORWAY_EXIT_LEFT = "MOTORWAY_EXIT_LEFT"
    #: Take the right exit.
    MOTORWAY_EXIT_RIGHT = "MOTORWAY_EXIT_RIGHT"
    #: Take the ferry.
    TAKE_FERRY = "TAKE_FERRY"
    #: Cross the roundabout.
    ROUNDABOUT_CROSS = "ROUNDABOUT_CROSS"
    #: At the roundabout take the exit on the right.
    ROUNDABOUT_RIGHT = "ROUNDABOUT_RIGHT"
    #: At the roundabout take the exit on the left.
    ROUNDABOUT_LEFT = "ROUNDABOUT_LEFT"
    #: Go around the roundabout.
    ROUNDABOUT_BACK = "ROUNDABOUT_BACK"
    #: Try to make a U-turn.
    TRY_MAKE_U_TURN = "TRY_MAKE_UTURN"
    #: Follow.
    FOLLOW = "FOLLOW"
    #: Switch to the parallel road.
    SWITCH_PARALLEL_ROAD = "SWITCH_PARALLEL_ROAD"
    #: Switch to the main road.
    SWITCH_MAIN_ROAD = "SWITCH_MAIN_ROAD"
    #: Take the ramp.
    ENTRANCE_RAMP = "ENTRANCE_RAMP"
    #: You have reached the waypoint. It is on the left.
    WAYPOINT_LEFT = "WAYPOINT_LEFT"
    #: You have reached the waypoint. It is on the right.
    WAYPOINT_RIGHT = "WAYPOINT_RIGHT"
    #: You have reached the waypoint.
    WAYPOINT_REACHED = "WAYPOINT_REACHED"


class InclineLevel(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """InclineLevel."""

    #: low
    LOW = "low"
    #: normal
    NORMAL = "normal"
    #: high
    HIGH = "high"


class JsonFormat(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """JsonFormat."""

    #: `The JavaScript Object Notation Data Interchange Format <https://tools.ietf.org/html/rfc8259>`_
    JSON = "json"


class JunctionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """The type of the junction where the maneuver takes place. For larger roundabouts, two separate
    instructions are generated for entering and leaving the roundabout.
    """

    #: regular
    REGULAR = "REGULAR"
    #: roundabout
    ROUNDABOUT = "ROUNDABOUT"
    #: bifurcation
    BIFURCATION = "BIFURCATION"


class Report(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Report."""

    #: Reports the effective parameters or data used when calling the API.
    EFFECTIVE_SETTINGS = "effectiveSettings"


class ResponseFormat(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """ResponseFormat."""

    #: `The JavaScript Object Notation Data Interchange Format <https://tools.ietf.org/html/rfc8259>`_
    JSON = "json"
    #: `The Extensible Markup Language <https://www.w3.org/TR/xml/>`_
    XML = "xml"


class ResponseSectionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Section types of the reported route response."""

    #: Sections of the route that are cars or trains.
    CAR_OR_TRAIN = "CAR_TRAIN"
    #: Sections indicating which countries the route is in.
    COUNTRY = "COUNTRY"
    #: Sections of the route that are ferries.
    FERRY = "FERRY"
    #: Sections of the route that are motorways.
    MOTORWAY = "MOTORWAY"
    #: Sections of the route that are only suited for pedestrians.
    PEDESTRIAN = "PEDESTRIAN"
    #: Sections of the route that require a toll to be payed.
    TOLL_ROAD = "TOLL_ROAD"
    #: Sections of the route that require a toll vignette to be present.
    TOLL_VIGNETTE = "TOLL_VIGNETTE"
    #: Sections of the route that contain traffic information.
    TRAFFIC = "TRAFFIC"
    #: Sections in relation to the request parameter ``travelMode``.
    TRAVEL_MODE = "TRAVEL_MODE"
    #: Sections of the route that are tunnels.
    TUNNEL = "TUNNEL"
    #: Sections of the route that require use of carpool (HOV/High Occupancy Vehicle) lanes.
    CARPOOL = "CARPOOL"
    #: Sections of the route that are located within urban areas.
    URBAN = "URBAN"


class ResponseTravelMode(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Travel mode for the calculated route. The value will be set to ``other`` if the requested mode
    of transport is not possible in this section.
    """

    #: The returned routes are optimized for cars.
    CAR = "car"
    #: The returned routes are optimized for commercial vehicles, like for trucks.
    TRUCK = "truck"
    #: The returned routes are optimized for taxis. BETA functionality.
    TAXI = "taxi"
    #: The returned routes are optimized for buses, including the use of bus only lanes. BETA
    #: functionality.
    BUS = "bus"
    #: The returned routes are optimized for vans. BETA functionality.
    VAN = "van"
    #: The returned routes are optimized for motorcycles. BETA functionality.
    MOTORCYCLE = "motorcycle"
    #: The returned routes are optimized for bicycles, including use of bicycle lanes.
    BICYCLE = "bicycle"
    #: The returned routes are optimized for pedestrians, including the use of sidewalks.
    PEDESTRIAN = "pedestrian"
    #: The given mode of transport is not possible in this section
    OTHER = "other"


class RouteAvoidType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """RouteAvoidType."""

    #: Avoids toll roads.
    TOLL_ROADS = "tollRoads"
    #: Avoids motorways
    MOTORWAYS = "motorways"
    #: Avoids ferries
    FERRIES = "ferries"
    #: Avoids unpaved roads
    UNPAVED_ROADS = "unpavedRoads"
    #: Avoids routes that require the use of carpool (HOV/High Occupancy Vehicle) lanes.
    CARPOOLS = "carpools"
    #: Avoids using the same road multiple times. Most useful in conjunction with ``routeType``\
    #: =thrilling.
    ALREADY_USED_ROADS = "alreadyUsedRoads"
    #: Avoids border crossings in route calculation.
    BORDER_CROSSINGS = "borderCrossings"


class RouteInstructionsType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """RouteInstructionsType."""

    #: Returns raw instruction data without human-readable messages.
    CODED = "coded"
    #: Returns raw instructions data with human-readable messages in plain text.
    TEXT = "text"
    #: Returns raw instruction data with tagged human-readable messages to permit formatting. A
    #: human-readable message is built up from repeatable identified elements. These are tagged to
    #: allow client applications to format them correctly. The following message components are tagged
    #: when instructionsType=tagged: street, roadNumber, signpostText, exitNumber,
    #: roundaboutExitNumber.
    #:
    #: Example of tagged 'Turn left' message:​
    #:
    #: .. code-block::
    #:
    #:    Turn left onto <roadNumber>A4</roadNumber>/<roadNumber>E19</roadNumber>
    #:    towards <signpostText>Den Haag</signpostText>
    TAGGED = "tagged"


class RouteRepresentationForBestOrder(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """RouteRepresentationForBestOrder."""

    #: Includes route geometry in the response.
    POLYLINE = "polyline"
    #: Summary as per polyline but excluding the point geometry elements for the routes in the
    #: response.
    SUMMARY_ONLY = "summaryOnly"
    #: Includes only the optimized waypoint indices but does not include the route geometry in the
    #: response.
    NONE = "none"


class RouteType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """RouteType."""

    #: The fastest route.
    FASTEST = "fastest"
    #: The shortest route by distance.
    SHORTEST = "shortest"
    #: A route balanced by economy and speed.
    ECONOMY = "eco"
    #: Includes interesting or challenging roads and uses as few motorways as possible. You can choose
    #: the level of turns included and also the degree of hilliness. See the hilliness and windingness
    #: parameters for how to set this. There is a limit of 900 km on routes planned with
    #: ``routeType``\ =thrilling
    THRILLING = "thrilling"


class SectionType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """SectionType."""

    #: Sections of the route that are cars or trains.
    CAR_OR_TRAIN = "carTrain"
    #: Sections indicating which countries the route is in.
    COUNTRY = "country"
    #: Sections of the route that are ferries.
    FERRY = "ferry"
    #: Sections of the route that are motorways.
    MOTORWAY = "motorway"
    #: Sections of the route that are only suited for pedestrians.
    PEDESTRIAN = "pedestrian"
    #: Sections of the route that require a toll to be payed.
    TOLL_ROAD = "tollRoad"
    #: Sections of the route that require a toll vignette to be present.
    TOLL_VIGNETTE = "tollVignette"
    #: Sections of the route that contain traffic information.
    TRAFFIC = "traffic"
    #: Sections in relation to the request parameter ``travelMode``.
    TRAVEL_MODE = "travelMode"
    #: Sections of the route that are tunnels.
    TUNNEL = "tunnel"
    #: Sections of the route that require use of carpool (HOV/High Occupancy Vehicle) lanes.
    CARPOOL = "carpool"
    #: Sections of the route that are located within urban areas.
    URBAN = "urban"


class SimpleCategory(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """Type of the incident. Can currently be JAM, ROAD_WORK, ROAD_CLOSURE, or OTHER. See "tec" for
    detailed information.
    """

    #: Traffic jam.
    JAM = "JAM"
    #: Road work.
    ROAD_WORK = "ROAD_WORK"
    #: Road closure.
    ROAD_CLOSURE = "ROAD_CLOSURE"
    #: Other.
    OTHER = "OTHER"


class TravelMode(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """TravelMode."""

    #: The returned routes are optimized for cars.
    CAR = "car"
    #: The returned routes are optimized for commercial vehicles, like for trucks.
    TRUCK = "truck"
    #: The returned routes are optimized for taxis. BETA functionality.
    TAXI = "taxi"
    #: The returned routes are optimized for buses, including the use of bus only lanes. BETA
    #: functionality.
    BUS = "bus"
    #: The returned routes are optimized for vans. BETA functionality.
    VAN = "van"
    #: The returned routes are optimized for motorcycles. BETA functionality.
    MOTORCYCLE = "motorcycle"
    #: The returned routes are optimized for bicycles, including use of bicycle lanes.
    BICYCLE = "bicycle"
    #: The returned routes are optimized for pedestrians, including the use of sidewalks.
    PEDESTRIAN = "pedestrian"


class VehicleEngineType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """VehicleEngineType."""

    #: Internal combustion engine.
    COMBUSTION = "combustion"
    #: Electric engine.
    ELECTRIC = "electric"


class VehicleLoadType(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """VehicleLoadType."""

    #: Explosives
    US_HAZMAT_CLASS1 = "USHazmatClass1"
    #: Compressed gas
    US_HAZMAT_CLASS2 = "USHazmatClass2"
    #: Flammable liquids
    US_HAZMAT_CLASS3 = "USHazmatClass3"
    #: Flammable solids
    US_HAZMAT_CLASS4 = "USHazmatClass4"
    #: Oxidizers
    US_HAZMAT_CLASS5 = "USHazmatClass5"
    #: Poisons
    US_HAZMAT_CLASS6 = "USHazmatClass6"
    #: Radioactive
    US_HAZMAT_CLASS7 = "USHazmatClass7"
    #: Corrosives
    US_HAZMAT_CLASS8 = "USHazmatClass8"
    #: Miscellaneous
    US_HAZMAT_CLASS9 = "USHazmatClass9"
    #: Explosives
    OTHER_HAZMAT_EXPLOSIVE = "otherHazmatExplosive"
    #: Miscellaneous
    OTHER_HAZMAT_GENERAL = "otherHazmatGeneral"
    #: Harmful to water
    OTHER_HAZMAT_HARMFUL_TO_WATER = "otherHazmatHarmfulToWater"


class WindingnessLevel(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """WindingnessLevel."""

    #: low
    LOW = "low"
    #: normal
    NORMAL = "normal"
    #: high
    HIGH = "high"
