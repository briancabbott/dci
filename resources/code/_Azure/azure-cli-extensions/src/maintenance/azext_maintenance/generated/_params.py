# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------
# pylint: disable=too-many-lines
# pylint: disable=too-many-statements

from azure.cli.core.commands.parameters import (
    tags_type,
    get_enum_type,
    resource_group_name_type,
    get_location_type
)
from azure.cli.core.commands.validators import (
    get_default_location_from_resource_group,
    validate_file_or_dict
)
from azext_maintenance.action import (
    AddExtensionProperties,
    AddWindowsParameters,
    AddLinuxParameters
)


def load_arguments(self, _):

    with self.argument_context('maintenance public-configuration show') as c:
        c.argument('resource_name', type=str, help='Maintenance Configuration Name', id_part='name')

    with self.argument_context('maintenance applyupdate show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('apply_update_name', options_list=['--name', '-n', '--apply-update-name'], type=str,
                   help='applyUpdate Id')

    with self.argument_context('maintenance applyupdate create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')

    with self.argument_context('maintenance applyupdate create-or-update-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')

    with self.argument_context('maintenance applyupdate show-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('apply_update_name', options_list=['--name', '-n', '--apply-update-name'], type=str,
                   help='applyUpdate Id')

    with self.argument_context('maintenance assignment list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')

    with self.argument_context('maintenance assignment show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Configuration assignment name')

    with self.argument_context('maintenance assignment create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Configuration assignment name')
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('maintenance_configuration_id', type=str, help='The maintenance configuration Id')
        c.argument('resource_id', type=str, help='The unique resourceId')

    with self.argument_context('maintenance assignment update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Configuration assignment name')
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('maintenance_configuration_id', type=str, help='The maintenance configuration Id')
        c.argument('resource_id', type=str, help='The unique resourceId')
        c.ignore('configuration_assignment')

    with self.argument_context('maintenance assignment delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Unique configuration assignment name')

    with self.argument_context('maintenance assignment create-or-update-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Configuration assignment name')
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('maintenance_configuration_id', type=str, help='The maintenance configuration Id')
        c.argument('resource_id', type=str, help='The unique resourceId')

    with self.argument_context('maintenance assignment delete-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Unique configuration assignment name')

    with self.argument_context('maintenance assignment list-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')

    with self.argument_context('maintenance assignment show-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
        c.argument('configuration_assignment_name', options_list=['--name', '-n', '--configuration-assignment-name'],
                   type=str, help='Configuration assignment name')

    with self.argument_context('maintenance configuration show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('resource_name', type=str, help='Maintenance Configuration Name')

    with self.argument_context('maintenance configuration create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('resource_name', type=str, help='Maintenance Configuration Name')
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('tags', tags_type)
        c.argument('namespace', type=str, help='Gets or sets namespace of the resource')
        c.argument('extension_properties', action=AddExtensionProperties, nargs='+', help='Gets or sets '
                   'extensionProperties of the maintenanceConfiguration Expect value: KEY1=VALUE1 KEY2=VALUE2 ...')
        c.argument('maintenance_scope', arg_type=get_enum_type(['Host', 'OSImage', 'Extension', 'InGuestPatch',
                                                                'SQLDB', 'SQLManagedInstance']), help='Gets or sets '
                   'maintenanceScope of the configuration')
        c.argument('visibility', arg_type=get_enum_type(['Custom', 'Public']), help='Gets or sets the visibility of '
                   'the configuration. The default value is \'Custom\'')
        c.argument('start_date_time', options_list=['--maintenance-window-start-date-time', '--start-date-time'],
                   type=str, help='Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The '
                   'start date can be set to either the current date or future date. The window will be created in the '
                   'time zone provided and adjusted to daylight savings according to that time zone.',
                   arg_group='Maintenance Window')
        c.argument('expiration_date_time', options_list=['--maintenance-window-expiration-date-time',
                                                         '--expiration-date-time'], type=str, help='Effective '
                   'expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created '
                   'in the time zone provided and adjusted to daylight savings according to that time zone. Expiration '
                   'date must be set to a future date. If not provided, it will be set to the maximum datetime '
                   '9999-12-31 23:59:59.', arg_group='Maintenance Window')
        c.argument('duration', options_list=['--maintenance-window-duration', '--duration'], type=str, help='Duration '
                   'of the maintenance window in HH:mm format. If not provided, default value will be used based on '
                   'maintenance scope provided. Example: 05:00.', arg_group='Maintenance Window')
        c.argument('time_zone', options_list=['--maintenance-window-time-zone', '--time-zone'], type=str, help='Name '
                   'of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTim'
                   'eZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea '
                   'Standard Time, Cen. Australia Standard Time.', arg_group='Maintenance Window')
        c.argument('recur_every', options_list=['--maintenance-window-recur-every', '--recur-every'], type=str,
                   help='Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, '
                   'weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as '
                   'integer][\'Day(s)\']. If no frequency is provided, the default frequency is 1. Daily schedule '
                   'examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: '
                   '[Frequency as integer][\'Week(s)\'] [Optional comma separated list of weekdays Monday-Sunday]. '
                   'Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly '
                   'schedules are formatted as [Frequency as integer][\'Month(s)\'] [Comma separated list of month '
                   'days] or [Frequency as integer][\'Month(s)\'] [Week of Month (First, Second, Third, Fourth, Last)] '
                   '[Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 '
                   'inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month '
                   'day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month '
                   'Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.', arg_group='Maintenance Window')
        c.argument('reboot_setting', arg_type=get_enum_type(['IfRequired', 'Never', 'Always']), help='Possible reboot '
                   'preference as defined by the user based on which it would be decided to reboot the machine or not '
                   'after the patch operation is completed.', arg_group='Install Patches')
        c.argument('windows_parameters', options_list=['--install-patches-windows-parameters', '--windows-parameters'],
                   action=AddWindowsParameters, nargs='+', help='Input parameters specific to patching a Windows '
                   'machine. For Linux machines, do not pass this property.', arg_group='Install Patches')
        c.argument('linux_parameters', options_list=['--install-patches-linux-parameters', '--linux-parameters'],
                   action=AddLinuxParameters, nargs='+', help='Input parameters specific to patching Linux machine. '
                   'For Windows machines, do not pass this property.', arg_group='Install Patches')
        c.argument('pre_tasks', options_list=['--install-patches-pre-tasks', '--pre-tasks'],
                   type=validate_file_or_dict, help='List of pre tasks. e.g. [{\'source\' :\'runbook\', \'taskScope\': '
                   '\'Global\', \'parameters\': { \'arg1\': \'value1\'}}] Expected value: '
                   'json-string/json-file/@json-file.', arg_group='Install Patches Tasks')
        c.argument('post_tasks', options_list=['--install-patches-post-tasks', '--post-tasks'],
                   type=validate_file_or_dict, help='List of post tasks. e.g. [{\'source\' :\'runbook\', '
                   '\'taskScope\': \'Resource\', \'parameters\': { \'arg1\': \'value1\'}}] Expected value: '
                   'json-string/json-file/@json-file.', arg_group='Install Patches Tasks')

    with self.argument_context('maintenance configuration update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('resource_name', type=str, help='Maintenance Configuration Name')
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('tags', tags_type)
        c.argument('namespace', type=str, help='Gets or sets namespace of the resource')
        c.argument('extension_properties', action=AddExtensionProperties, nargs='+', help='Gets or sets '
                   'extensionProperties of the maintenanceConfiguration Expect value: KEY1=VALUE1 KEY2=VALUE2 ...')
        c.argument('maintenance_scope', arg_type=get_enum_type(['Host', 'OSImage', 'Extension', 'InGuestPatch',
                                                                'SQLDB', 'SQLManagedInstance']), help='Gets or sets '
                   'maintenanceScope of the configuration')
        c.argument('visibility', arg_type=get_enum_type(['Custom', 'Public']), help='Gets or sets the visibility of '
                   'the configuration. The default value is \'Custom\'')
        c.argument('start_date_time', options_list=['--maintenance-window-start-date-time', '--start-date-time'],
                   type=str, help='Effective start date of the maintenance window in YYYY-MM-DD hh:mm format. The '
                   'start date can be set to either the current date or future date. The window will be created in the '
                   'time zone provided and adjusted to daylight savings according to that time zone.',
                   arg_group='Maintenance Window')
        c.argument('expiration_date_time', options_list=['--maintenance-window-expiration-date-time',
                                                         '--expiration-date-time'], type=str, help='Effective '
                   'expiration date of the maintenance window in YYYY-MM-DD hh:mm format. The window will be created '
                   'in the time zone provided and adjusted to daylight savings according to that time zone. Expiration '
                   'date must be set to a future date. If not provided, it will be set to the maximum datetime '
                   '9999-12-31 23:59:59.', arg_group='Maintenance Window')
        c.argument('duration', options_list=['--maintenance-window-duration', '--duration'], type=str, help='Duration '
                   'of the maintenance window in HH:mm format. If not provided, default value will be used based on '
                   'maintenance scope provided. Example: 05:00.', arg_group='Maintenance Window')
        c.argument('time_zone', options_list=['--maintenance-window-time-zone', '--time-zone'], type=str, help='Name '
                   'of the timezone. List of timezones can be obtained by executing [System.TimeZoneInfo]::GetSystemTim'
                   'eZones() in PowerShell. Example: Pacific Standard Time, UTC, W. Europe Standard Time, Korea '
                   'Standard Time, Cen. Australia Standard Time.', arg_group='Maintenance Window')
        c.argument('recur_every', options_list=['--maintenance-window-recur-every', '--recur-every'], type=str,
                   help='Rate at which a Maintenance window is expected to recur. The rate can be expressed as daily, '
                   'weekly, or monthly schedules. Daily schedule are formatted as recurEvery: [Frequency as '
                   'integer][\'Day(s)\']. If no frequency is provided, the default frequency is 1. Daily schedule '
                   'examples are recurEvery: Day, recurEvery: 3Days.  Weekly schedule are formatted as recurEvery: '
                   '[Frequency as integer][\'Week(s)\'] [Optional comma separated list of weekdays Monday-Sunday]. '
                   'Weekly schedule examples are recurEvery: 3Weeks, recurEvery: Week Saturday,Sunday. Monthly '
                   'schedules are formatted as [Frequency as integer][\'Month(s)\'] [Comma separated list of month '
                   'days] or [Frequency as integer][\'Month(s)\'] [Week of Month (First, Second, Third, Fourth, Last)] '
                   '[Weekday Monday-Sunday] [Optional Offset(No. of days)]. Offset value must be between -6 to 6 '
                   'inclusive. Monthly schedule examples are recurEvery: Month, recurEvery: 2Months, recurEvery: Month '
                   'day23,day24, recurEvery: Month Last Sunday, recurEvery: Month Fourth Monday, recurEvery: Month '
                   'Last Sunday Offset-3, recurEvery: Month Third Sunday Offset6.', arg_group='Maintenance Window')
        c.argument('reboot_setting', arg_type=get_enum_type(['IfRequired', 'Never', 'Always']), help='Possible reboot '
                   'preference as defined by the user based on which it would be decided to reboot the machine or not '
                   'after the patch operation is completed.', arg_group='Install Patches')
        c.argument('windows_parameters', options_list=['--install-patches-windows-parameters', '--windows-parameters'],
                   action=AddWindowsParameters, nargs='+', help='Input parameters specific to patching a Windows '
                   'machine. For Linux machines, do not pass this property.', arg_group='Install Patches')
        c.argument('linux_parameters', options_list=['--install-patches-linux-parameters', '--linux-parameters'],
                   action=AddLinuxParameters, nargs='+', help='Input parameters specific to patching Linux machine. '
                   'For Windows machines, do not pass this property.', arg_group='Install Patches')
        c.argument('pre_tasks', options_list=['--install-patches-pre-tasks', '--pre-tasks'],
                   type=validate_file_or_dict, help='List of pre tasks. e.g. [{\'source\' :\'runbook\', \'taskScope\': '
                   '\'Global\', \'parameters\': { \'arg1\': \'value1\'}}] Expected value: '
                   'json-string/json-file/@json-file.', arg_group='Install Patches Tasks')
        c.argument('post_tasks', options_list=['--install-patches-post-tasks', '--post-tasks'],
                   type=validate_file_or_dict, help='List of post tasks. e.g. [{\'source\' :\'runbook\', '
                   '\'taskScope\': \'Resource\', \'parameters\': { \'arg1\': \'value1\'}}] Expected value: '
                   'json-string/json-file/@json-file.', arg_group='Install Patches Tasks')

    with self.argument_context('maintenance configuration delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('resource_name', type=str, help='Maintenance Configuration Name')

    with self.argument_context('maintenance update list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')

    with self.argument_context('maintenance update list-parent') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('provider_name', type=str, help='Resource provider name')
        c.argument('resource_parent_type', type=str, help='Resource parent type')
        c.argument('resource_parent_name', type=str, help='Resource parent identifier')
        c.argument('resource_type', type=str, help='Resource type')
        c.argument('resource_name', type=str, help='Resource identifier')
