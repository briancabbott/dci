# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------
# pylint: disable=too-many-statements
# pylint: disable=too-many-locals
# pylint: disable=line-too-long

from azure.cli.core.commands import CliCommandType


def load_command_table(self, _):

    from azext_amcs._client_factory import cf_data_collection_rule

    monitor_control_service_data_collection_rule = CliCommandType(
        operations_tmpl='azext_amcs.vendored_sdks.amcs.operations._data_collection_rules_operations#DataCollectionRules'
                        'Operations.{}',
        client_factory=cf_data_collection_rule)

    with self.command_group('monitor data-collection rule association') as g:
        g.custom_command('list', 'monitor_data_collection_rule_association_list')

    from .custom import EndpointCreate
    self.command_table['monitor data-collection endpoint create'] = EndpointCreate(loader=self)

    with self.command_group('monitor data-collection rule') as g:
        # g.custom_command('list', 'data_collection_rules_list')
        # g.custom_command('create', 'data_collection_rules_create')
        # g.custom_command('update', 'data_collection_rules_update')
        from .custom import RuleCreate, RuleUpdate
        self.command_table['monitor data-collection rule create'] = RuleCreate(loader=self)
        self.command_table['monitor data-collection rule update'] = RuleUpdate(loader=self)

    with self.command_group('monitor data-collection rule data-flow', monitor_control_service_data_collection_rule,
                            client_factory=cf_data_collection_rule) as g:
        g.custom_show_command('list', 'data_collection_rules_data_flows_list')
        g.custom_command('add', 'data_collection_rules_data_flows_add')

    with self.command_group('monitor data-collection rule log-analytics', monitor_control_service_data_collection_rule,
                            client_factory=cf_data_collection_rule) as g:
        g.custom_show_command('list', 'data_collection_rules_log_analytics_list')
        g.custom_show_command('show', 'data_collection_rules_log_analytics_show')
        g.custom_command('add', 'data_collection_rules_log_analytics_add')
        g.custom_command('delete', 'data_collection_rules_log_analytics_delete')
        g.custom_command('update', 'data_collection_rules_log_analytics_update')

    with self.command_group('monitor data-collection rule performance-counter',
                            monitor_control_service_data_collection_rule,
                            client_factory=cf_data_collection_rule) as g:
        g.custom_show_command('list', 'data_collection_rules_performance_counters_list')
        g.custom_show_command('show', 'data_collection_rules_performance_counters_show')
        g.custom_command('add', 'data_collection_rules_performance_counters_add')
        g.custom_command('delete', 'data_collection_rules_performance_counters_delete')
        g.custom_command('update', 'data_collection_rules_performance_counters_update')

    with self.command_group('monitor data-collection rule windows-event-log',
                            monitor_control_service_data_collection_rule,
                            client_factory=cf_data_collection_rule) as g:
        g.custom_show_command('list', 'data_collection_rules_windows_event_logs_list')
        g.custom_show_command('show', 'data_collection_rules_windows_event_logs_show')
        g.custom_command('add', 'data_collection_rules_windows_event_logs_add')
        g.custom_command('delete', 'data_collection_rules_windows_event_logs_delete')
        g.custom_command('update', 'data_collection_rules_windows_event_logs_update')

    with self.command_group('monitor data-collection rule syslog', monitor_control_service_data_collection_rule,
                            client_factory=cf_data_collection_rule) as g:
        g.custom_show_command('list', 'data_collection_rules_syslog_list')
        g.custom_show_command('show', 'data_collection_rules_syslog_show')
        g.custom_command('add', 'data_collection_rules_syslog_add')
        g.custom_command('delete', 'data_collection_rules_syslog_delete')
        g.custom_command('update', 'data_collection_rules_syslog_update')
