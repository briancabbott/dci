# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------

from argparse import ArgumentError
from azure.cli.core.azclierror import InvalidArgumentValueError, ArgumentUsageError
from azure.cli.command_modules.iot.custom import SimpleAccessRights, iot_central_app_get


def validate_policy_permissions(ns):
    if ns.permissions is None or ns.permissions == []:
        raise ArgumentError(None, 'the following arguments are required: --permissions')

    allowed = [x.value.lower() for x in SimpleAccessRights]
    for p in ns.permissions:
        if p not in allowed:
            raise ValueError('Unrecognized permission "{}"'.format(p))


def validate_retention_days(ns):
    if ns.retention_day and ns.retention_day not in range(1, 8, 1):
        raise ArgumentError(None, 'Please specify the retention time for device-to-cloud messages'
                                  ' from 1 to 7 days only.')


def validate_fileupload_notification_max_delivery_count(ns):
    if (ns.fileupload_notification_max_delivery_count and
            ns.fileupload_notification_max_delivery_count not in range(1, 101, 1)):
        raise ArgumentError(None, 'Please specify the number of retry from 1 to 100 only to deliver a message')


def validate_fileupload_notification_ttl(ns):
    if (ns.fileupload_notification_ttl and
            ns.fileupload_notification_ttl not in range(1, 49, 1)):
        raise ArgumentError(None, 'Please specify the period of time from 1 to 48 hours for which a file upload'
                                  ' notification is available to consume before it is expired.')


def validate_fileupload_sas_ttl(ns):
    if (ns.fileupload_sas_ttl and
            ns.fileupload_sas_ttl not in range(1, 25, 1)):
        raise ArgumentError(None, 'Please specify the period of time from 1 to 24 hours for which the SAS URI'
                                  ' generated by IoT Hub for file upload is valid.')


def validate_feedback_ttl(ns):
    if (ns.feedback_ttl and
            ns.feedback_ttl not in range(1, 49, 1)):
        raise ArgumentError(None, 'Please specify the period of time from 1 to 48 hours for which a message is'
                                  ' available to consume before it is expired by the IoT hub.')


def validate_feedback_lock_duration(ns):
    if (ns.feedback_lock_duration and
            ns.feedback_lock_duration not in range(5, 301, 1)):
        raise InvalidArgumentValueError('Please specify the feedback lock duration from 5 to 300 seconds only.')


def validate_fileupload_notification_lock_duration(ns):
    if (ns.fileupload_notification_lock_duration and
            ns.fileupload_notification_lock_duration not in range(5, 301, 1)):
        raise InvalidArgumentValueError('Please specify the notification lock duration from 5 to 300 seconds only.')


def validate_feedback_max_delivery_count(ns):
    if (ns.feedback_max_delivery_count and
            ns.feedback_max_delivery_count not in range(1, 101, 1)):
        raise ArgumentError(None, 'Please specify the number of retry from 1 to 100 only to deliver a message'
                                  ' on the feedback queue.')


def validate_c2d_max_delivery_count(ns):
    if (ns.c2d_max_delivery_count and
            ns.c2d_max_delivery_count not in range(1, 101, 1)):
        raise ArgumentError(None, 'Please specify the maximum delivery count from 1 to 100 only for cloud-to-device'
                                  ' messages in the device queue.')


def validate_c2d_ttl(ns):
    if (ns.c2d_ttl and
            ns.c2d_ttl not in range(1, 49, 1)):
        raise ArgumentError(None, 'Please specify the default time from 1 to 48 hours to live for cloud-to-device'
                                  ' messages in the device queue.')


def validate_private_endpoint_connection_id(cmd, namespace):
    ns = namespace
    if ns.connection_id:
        id_parts = ns.connection_id.split('/')
        ns.private_endpoint_connection_name = id_parts[-1]
        ns.account_name = id_parts[-3]
        ns.resource_group_name = id_parts[-7]
        del ns.connection_id
    if ns.account_name and not ns.resource_group_name:
        ns.resource_group_name = iot_central_app_get(cmd.cli_ctx, ns.account_name).resourceGroup

    if not all([ns.account_name, ns.resource_group_name, ns.private_endpoint_connection_name]):
        raise ArgumentUsageError('incorrect usage: [--id ID | --name NAME --account-name NAME]')
