# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from ._patch import AnomalyAlert
from ._patch import AnomalyAlertConfiguration
from ._patch import AnomalyDetectionConfiguration
from ._patch import AnomalyFeedback
from ._patch import AnomalyIncident
from ._patch import AzureApplicationInsightsDataFeedSource
from ._patch import AzureBlobDataFeedSource
from ._patch import AzureCosmosDbDataFeedSource
from ._patch import AzureDataExplorerDataFeedSource
from ._patch import AzureDataLakeStorageGen2DataFeedSource
from ._patch import AzureEventHubsDataFeedSource
from ._patch import AzureLogAnalyticsDataFeedSource
from ._patch import AzureTableDataFeedSource
from ._patch import ChangePointFeedback
from ._patch import ChangeThresholdCondition
from ._patch import CommentFeedback
from ._patch import DataFeed
from ._patch import DataFeedDimension
from ._patch import DataFeedIngestionProgress
from ._models import DataFeedIngestionStatus
from ._patch import DataFeedMetric
from ._patch import DataPointAnomaly
from ._patch import DatasourceCredential
from ._patch import DatasourceDataLakeGen2SharedKey
from ._patch import DatasourceServicePrincipal
from ._patch import DatasourceServicePrincipalInKeyVault
from ._patch import DatasourceSqlConnectionString
from ._patch import DetectionAnomalyFilterCondition
from ._models import DetectionIncidentFilterCondition
from ._models import DimensionGroupIdentity
from ._patch import EmailNotificationHook
from ._models import EnrichmentStatus
from ._patch import HardThresholdCondition
from ._patch import IncidentRootCause
from ._patch import InfluxDbDataFeedSource
from ._patch import MetricAlertConfiguration
from ._patch import MetricAnomalyAlertSnoozeCondition
from ._patch import MetricBoundaryCondition
from ._patch import MetricDetectionCondition
from ._patch import MetricEnrichedSeriesData
from ._patch import MetricFeedback
from ._patch import MetricSeriesData
from ._models import MetricSeriesDefinition
from ._patch import MetricSeriesGroupDetectionCondition
from ._patch import MetricSingleSeriesDetectionCondition
from ._patch import MongoDbDataFeedSource
from ._patch import NotificationHook
from ._patch import PeriodFeedback
from ._patch import SeriesIdentity
from ._patch import SeverityCondition
from ._models import SeverityFilterCondition
from ._patch import SmartDetectionCondition
from ._patch import SuppressCondition
from ._patch import TopNGroupScope
from ._patch import WebNotificationHook


from ._enums import AnomalyDetectorDirection
from ._enums import AnomalyIncidentStatus
from ._enums import AnomalySeverity
from ._enums import AnomalyValue
from ._enums import ChangePointValue
from ._enums import DataFeedAccessMode
from ._enums import DataFeedAutoRollupMethod
from ._enums import DataFeedGranularityType
from ._enums import DataFeedRollupType
from ._enums import DataFeedStatus
from ._enums import DatasourceAuthenticationType
from ._enums import DatasourceCredentialType
from ._enums import DatasourceMissingDataPointFillType
from ._enums import DatasourceType
from ._enums import DetectionConditionOperator
from ._enums import FeedbackType
from ._enums import MetricAnomalyAlertConfigurationsOperator
from ._enums import MetricAnomalyAlertScopeType
from ._enums import PeriodType
from ._enums import SnoozeScope

from ._patch import DataFeedGranularity
from ._patch import DataFeedIngestionSettings
from ._patch import DataFeedMissingDataPointFillSettings
from ._patch import DataFeedRollupSettings
from ._patch import DataFeedSchema
from ._patch import MetricAnomalyAlertScope
from ._patch import SqlServerDataFeedSource
from ._patch import MySqlDataFeedSource
from ._patch import PostgreSqlDataFeedSource
from ._patch import MetricAnomalyAlertConditions
from ._patch import DataFeedSource
from ._patch import AlertQueryTimeMode
from ._patch import FeedbackQueryTimeMode
from ._patch import patch_sdk as _patch_sdk

__all__ = [
    "DataFeedGranularity",
    "DataFeedIngestionSettings",
    "DataFeedMissingDataPointFillSettings",
    "DataFeedRollupSettings",
    "DataFeedSchema",
    "MetricAnomalyAlertScope",
    "SqlServerDataFeedSource",
    "MySqlDataFeedSource",
    "PostgreSqlDataFeedSource",
    "MetricAnomalyAlertConditions",
    "DataFeedSource",
    "AlertQueryTimeMode",
    "FeedbackQueryTimeMode",
    "AnomalyAlert",
    "AnomalyAlertConfiguration",
    "AnomalyDetectionConfiguration",
    "AnomalyFeedback",
    "AnomalyIncident",
    "AzureApplicationInsightsDataFeedSource",
    "AzureBlobDataFeedSource",
    "AzureCosmosDbDataFeedSource",
    "AzureDataExplorerDataFeedSource",
    "AzureDataLakeStorageGen2DataFeedSource",
    "AzureEventHubsDataFeedSource",
    "AzureLogAnalyticsDataFeedSource",
    "AzureTableDataFeedSource",
    "ChangePointFeedback",
    "ChangeThresholdCondition",
    "CommentFeedback",
    "DataFeed",
    "DataFeedDimension",
    "DataFeedIngestionProgress",
    "DataFeedIngestionStatus",
    "DataFeedMetric",
    "DataPointAnomaly",
    "DatasourceCredential",
    "DatasourceDataLakeGen2SharedKey",
    "DatasourceServicePrincipal",
    "DatasourceServicePrincipalInKeyVault",
    "DatasourceSqlConnectionString",
    "DetectionAnomalyFilterCondition",
    "DetectionIncidentFilterCondition",
    "DimensionGroupIdentity",
    "EmailNotificationHook",
    "EnrichmentStatus",
    "HardThresholdCondition",
    "IncidentRootCause",
    "InfluxDbDataFeedSource",
    "MetricAlertConfiguration",
    "MetricAnomalyAlertSnoozeCondition",
    "MetricBoundaryCondition",
    "MetricDetectionCondition",
    "MetricEnrichedSeriesData",
    "MetricFeedback",
    "MetricSeriesData",
    "MetricSeriesDefinition",
    "MetricSeriesGroupDetectionCondition",
    "MetricSingleSeriesDetectionCondition",
    "MongoDbDataFeedSource",
    "NotificationHook",
    "PeriodFeedback",
    "SeriesIdentity",
    "SeverityCondition",
    "SeverityFilterCondition",
    "SmartDetectionCondition",
    "SuppressCondition",
    "TopNGroupScope",
    "WebNotificationHook",
    "AnomalyDetectorDirection",
    "AnomalyIncidentStatus",
    "AnomalySeverity",
    "AnomalyValue",
    "ChangePointValue",
    "DataFeedAccessMode",
    "DataFeedAutoRollupMethod",
    "DataFeedGranularityType",
    "DataFeedRollupType",
    "DataFeedStatus",
    "DatasourceAuthenticationType",
    "DatasourceCredentialType",
    "DatasourceMissingDataPointFillType",
    "DatasourceType",
    "DetectionConditionOperator",
    "FeedbackType",
    "MetricAnomalyAlertConfigurationsOperator",
    "MetricAnomalyAlertScopeType",
    "PeriodType",
    "SnoozeScope",
]

_patch_sdk()
