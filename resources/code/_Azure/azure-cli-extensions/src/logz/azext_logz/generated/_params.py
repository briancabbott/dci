# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------
# pylint: disable=line-too-long
# pylint: disable=too-many-lines
# pylint: disable=too-many-statements

from azure.cli.core.commands.parameters import (
    tags_type,
    get_three_state_flag,
    get_enum_type,
    resource_group_name_type,
    get_location_type
)
from azure.cli.core.commands.validators import get_default_location_from_resource_group
from azext_logz.action import (
    AddLogzOrganizationProperties,
    AddUserInfo,
    AddPlanData,
    AddVmResourceIds,
    AddFilteringTags,
    AddProperties
)


def load_arguments(self, _):

    with self.argument_context('logz monitor list') as c:
        c.argument('resource_group_name', resource_group_name_type)

    with self.argument_context('logz monitor show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name', id_part='name')

    with self.argument_context('logz monitor create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name')
        c.argument('tags', tags_type)
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('type_', options_list=['--type'], arg_type=get_enum_type(['SystemAssigned', 'UserAssigned']),
                   help='Identity type.', arg_group='Identity')
        c.argument('monitoring_status', arg_type=get_enum_type(['Enabled', 'Disabled']), help='Flag specifying if the '
                   'resource monitoring is enabled or disabled.')
        c.argument('marketplace_subscription_status', options_list=['--subscription-status'],
                   arg_type=get_enum_type(['Active', 'Suspended']),
                   help='Flag specifying the Marketplace Subscription Status of the resource. If payment is not made '
                   'in time, the resource will go in Suspended state.')
        c.argument('logz_organization_properties', options_list=['--org-properties'],
                   action=AddLogzOrganizationProperties, nargs='+', help='')
        c.argument('user_info', action=AddUserInfo, nargs='+', help='')
        c.argument('plan_data', action=AddPlanData, nargs='+', help='')

    with self.argument_context('logz monitor update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name', id_part='name')
        c.argument('tags', tags_type)
        c.argument('monitoring_status', arg_type=get_enum_type(['Enabled', 'Disabled']), help='Flag specifying if the '
                   'resource monitoring is enabled or disabled.')

    with self.argument_context('logz monitor delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name', id_part='name')

    with self.argument_context('logz monitor list-payload') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name')

    with self.argument_context('logz monitor list-resource') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name')

    with self.argument_context('logz monitor list-role') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name')
        c.argument('email_address', type=str, help='Email of the user used by Logz for contacting them if needed')

    with self.argument_context('logz monitor list-vm') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name')

    with self.argument_context('logz monitor update-vm') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name', id_part='name')
        c.argument('vm_resource_ids', action=AddVmResourceIds, nargs='+', help='Request of a list vm host update '
                   'operation.')
        c.argument('state', arg_type=get_enum_type(['Install', 'Delete']), help='Specifies the state of the operation '
                   '- install/ delete.')

    with self.argument_context('logz monitor wait') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', options_list=['--name', '-n', '--monitor-name'], type=str, help='Monitor resource '
                   'name', id_part='name')

    with self.argument_context('logz rule list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')

    with self.argument_context('logz rule show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.', id_part='child_name_1')

    with self.argument_context('logz rule create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.')
        c.argument('send_aad_logs', arg_type=get_three_state_flag(), help='Flag specifying if AAD logs should be sent '
                   'for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_subscription_logs', arg_type=get_three_state_flag(), help='Flag specifying if subscription '
                   'logs should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_activity_logs', arg_type=get_three_state_flag(), help='Flag specifying if activity logs from '
                   'Azure resources should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('filtering_tags', action=AddFilteringTags, nargs='+', help='List of filtering tags to be used for '
                   'capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all '
                   'resources will be captured. If only Exclude action is specified, the rules will apply to the list '
                   'of all available resources. If Include actions are specified, the rules will only include '
                   'resources with the associated tags.', arg_group='Log Rules')

    with self.argument_context('logz rule update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.', id_part='child_name_1')
        c.argument('send_aad_logs', arg_type=get_three_state_flag(), help='Flag specifying if AAD logs should be sent '
                   'for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_subscription_logs', arg_type=get_three_state_flag(), help='Flag specifying if subscription '
                   'logs should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_activity_logs', arg_type=get_three_state_flag(), help='Flag specifying if activity logs from '
                   'Azure resources should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('filtering_tags', action=AddFilteringTags, nargs='+', help='List of filtering tags to be used for '
                   'capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all '
                   'resources will be captured. If only Exclude action is specified, the rules will apply to the list '
                   'of all available resources. If Include actions are specified, the rules will only include '
                   'resources with the associated tags.', arg_group='Log Rules')
        c.ignore('body')

    with self.argument_context('logz rule delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.', id_part='child_name_1')

    with self.argument_context('logz sso list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')

    with self.argument_context('logz sso show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('configuration_name', type=str, help='Single sign-on configuration name.', id_part='child_name_1')

    with self.argument_context('logz sso create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('configuration_name', type=str, help='Single sign-on configuration name.')
        c.argument('properties', action=AddProperties, nargs='+', help='')

    with self.argument_context('logz sso update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('configuration_name', type=str, help='Single sign-on configuration name.', id_part='child_name_1')
        c.argument('properties', action=AddProperties, nargs='+', help='')
        c.ignore('body')

    with self.argument_context('logz sso wait') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('configuration_name', type=str, help='Single sign-on configuration name.', id_part='child_name_1')

    with self.argument_context('logz sub-account list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')

    with self.argument_context('logz sub-account show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name', id_part='child_name_1')

    with self.argument_context('logz sub-account create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name')
        c.argument('tags', tags_type)
        c.argument('location', arg_type=get_location_type(self.cli_ctx), required=False,
                   validator=get_default_location_from_resource_group)
        c.argument('type_', options_list=['--type'], arg_type=get_enum_type(['SystemAssigned', 'UserAssigned']),
                   help='Identity type.', arg_group='Identity')
        c.argument('monitoring_status', arg_type=get_enum_type(['Enabled', 'Disabled']), help='Flag specifying if the '
                   'resource monitoring is enabled or disabled.')
        c.argument('marketplace_subscription_status', options_list=['--subscription-status'],
                   arg_type=get_enum_type(['Active', 'Suspended']),
                   help='Flag specifying the Marketplace Subscription Status of the resource. If payment is not made '
                   'in time, the resource will go in Suspended state.')
        c.argument('logz_organization_properties', options_list=['--org-properties'],
                   action=AddLogzOrganizationProperties, nargs='+', help='')
        c.argument('user_info', action=AddUserInfo, nargs='+', help='')
        c.argument('plan_data', action=AddPlanData, nargs='+', help='')

    with self.argument_context('logz sub-account update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name', id_part='child_name_1')
        c.argument('tags', tags_type)
        c.argument('monitoring_status', arg_type=get_enum_type(['Enabled', 'Disabled']), help='Flag specifying if the '
                   'resource monitoring is enabled or disabled.')

    with self.argument_context('logz sub-account delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name', id_part='child_name_1')

    with self.argument_context('logz sub-account list-payload') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name')

    with self.argument_context('logz sub-account list-resource') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name')

    with self.argument_context('logz sub-account list-vm') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name')

    with self.argument_context('logz sub-account update-vm') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name', id_part='child_name_1')
        c.argument('vm_resource_ids', action=AddVmResourceIds, nargs='+', help='Request of a list vm host update '
                   'operation.')
        c.argument('state', arg_type=get_enum_type(['Install', 'Delete']), help='Specifies the state of the operation '
                   '- install/ delete.')

    with self.argument_context('logz sub-account wait') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', options_list=['--name', '-n', '--sub-account-name'], type=str, help='Sub '
                   'Account resource name', id_part='child_name_1')

    with self.argument_context('logz sub-rule list') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', type=str, help='Sub Account resource name')

    with self.argument_context('logz sub-rule show') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', type=str, help='Sub Account resource name', id_part='child_name_1')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.', id_part='child_name_2')

    with self.argument_context('logz sub-rule create') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name')
        c.argument('sub_account_name', type=str, help='Sub Account resource name')
        c.argument('rule_set_name', type=str, help='Rule set name of sub account.')
        c.argument('send_aad_logs', arg_type=get_three_state_flag(), help='Flag specifying if AAD logs should be sent '
                   'for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_subscription_logs', arg_type=get_three_state_flag(), help='Flag specifying if subscription '
                   'logs should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_activity_logs', arg_type=get_three_state_flag(), help='Flag specifying if activity logs from '
                   'Azure resources should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('filtering_tags', action=AddFilteringTags, nargs='+', help='List of filtering tags to be used for '
                   'capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all '
                   'resources will be captured. If only Exclude action is specified, the rules will apply to the list '
                   'of all available resources. If Include actions are specified, the rules will only include '
                   'resources with the associated tags.', arg_group='Log Rules')

    with self.argument_context('logz sub-rule update') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', type=str, help='Sub Account resource name', id_part='child_name_1')
        c.argument('rule_set_name', type=str, help='Rule set name of sub account.', id_part='child_name_2')
        c.argument('send_aad_logs', arg_type=get_three_state_flag(), help='Flag specifying if AAD logs should be sent '
                   'for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_subscription_logs', arg_type=get_three_state_flag(), help='Flag specifying if subscription '
                   'logs should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('send_activity_logs', arg_type=get_three_state_flag(), help='Flag specifying if activity logs from '
                   'Azure resources should be sent for the Monitor resource.', arg_group='Log Rules')
        c.argument('filtering_tags', action=AddFilteringTags, nargs='+', help='List of filtering tags to be used for '
                   'capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty, all '
                   'resources will be captured. If only Exclude action is specified, the rules will apply to the list '
                   'of all available resources. If Include actions are specified, the rules will only include '
                   'resources with the associated tags.', arg_group='Log Rules')
        c.ignore('body')

    with self.argument_context('logz sub-rule delete') as c:
        c.argument('resource_group_name', resource_group_name_type)
        c.argument('monitor_name', type=str, help='Monitor resource name', id_part='name')
        c.argument('sub_account_name', type=str, help='Sub Account resource name', id_part='child_name_1')
        c.argument('rule_set_name', type=str, help='Rule set name of monitor.', id_part='child_name_2')
