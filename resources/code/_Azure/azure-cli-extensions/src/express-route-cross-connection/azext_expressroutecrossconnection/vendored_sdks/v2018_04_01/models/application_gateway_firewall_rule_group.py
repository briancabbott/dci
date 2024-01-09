# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------

from msrest.serialization import Model


class ApplicationGatewayFirewallRuleGroup(Model):
    """A web application firewall rule group.

    All required parameters must be populated in order to send to Azure.

    :param rule_group_name: Required. The name of the web application firewall
     rule group.
    :type rule_group_name: str
    :param description: The description of the web application firewall rule
     group.
    :type description: str
    :param rules: Required. The rules of the web application firewall rule
     group.
    :type rules:
     list[~azure.mgmt.network.v2018_04_01.models.ApplicationGatewayFirewallRule]
    """

    _validation = {
        'rule_group_name': {'required': True},
        'rules': {'required': True},
    }

    _attribute_map = {
        'rule_group_name': {'key': 'ruleGroupName', 'type': 'str'},
        'description': {'key': 'description', 'type': 'str'},
        'rules': {'key': 'rules', 'type': '[ApplicationGatewayFirewallRule]'},
    }

    def __init__(self, **kwargs):
        super(ApplicationGatewayFirewallRuleGroup, self).__init__(**kwargs)
        self.rule_group_name = kwargs.get('rule_group_name', None)
        self.description = kwargs.get('description', None)
        self.rules = kwargs.get('rules', None)
