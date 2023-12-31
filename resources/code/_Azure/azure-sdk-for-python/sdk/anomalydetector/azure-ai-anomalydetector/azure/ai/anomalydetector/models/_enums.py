# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# Code generated by Microsoft (R) Python Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from azure.core import CaseInsensitiveEnumMeta


class AlignMode(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """AlignMode."""

    INNER = "Inner"
    OUTER = "Outer"


class AnomalyDetectorErrorCodes(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """AnomalyDetectorErrorCodes."""

    INVALID_CUSTOM_INTERVAL = "InvalidCustomInterval"
    BAD_ARGUMENT = "BadArgument"
    INVALID_GRANULARITY = "InvalidGranularity"
    INVALID_PERIOD = "InvalidPeriod"
    INVALID_MODEL_ARGUMENT = "InvalidModelArgument"
    INVALID_SERIES = "InvalidSeries"
    INVALID_JSON_FORMAT = "InvalidJsonFormat"
    REQUIRED_GRANULARITY = "RequiredGranularity"
    REQUIRED_SERIES = "RequiredSeries"
    INVALID_IMPUTE_MODE = "InvalidImputeMode"
    INVALID_IMPUTE_FIXED_VALUE = "InvalidImputeFixedValue"


class APIVersion(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """APIVersion."""

    V1_1 = "v1.1"


class DataSchema(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """DataSchema."""

    #: OneTable means that your input data are all in one CSV file, which contains one 'timestamp'
    #: column and several variable columns. The default DataSchema is OneTable.
    ONE_TABLE = "OneTable"
    #: MultiTable means that your input data are separated in multiple CSV files, in each file
    #: containing one 'timestamp' column and one 'variable' column, and the CSV file name should
    #: indicate the name of the variable. The default DataSchema is OneTable.
    MULTI_TABLE = "MultiTable"


class FillNAMethod(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """FillNAMethod."""

    PREVIOUS = "Previous"
    SUBSEQUENT = "Subsequent"
    LINEAR = "Linear"
    ZERO = "Zero"
    FIXED = "Fixed"


class ImputeMode(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """ImputeMode."""

    AUTO = "auto"
    PREVIOUS = "previous"
    LINEAR = "linear"
    FIXED = "fixed"
    ZERO = "zero"
    NOT_FILL = "notFill"


class ModelStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """ModelStatus."""

    CREATED = "CREATED"
    RUNNING = "RUNNING"
    READY = "READY"
    FAILED = "FAILED"


class MultivariateBatchDetectionStatus(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """MultivariateBatchDetectionStatus."""

    CREATED = "CREATED"
    RUNNING = "RUNNING"
    READY = "READY"
    FAILED = "FAILED"


class TimeGranularity(str, Enum, metaclass=CaseInsensitiveEnumMeta):
    """TimeGranularity."""

    YEARLY = "yearly"
    MONTHLY = "monthly"
    WEEKLY = "weekly"
    DAILY = "daily"
    HOURLY = "hourly"
    PER_MINUTE = "minutely"
    PER_SECOND = "secondly"
    MICROSECOND = "microsecond"
    NONE = "none"
