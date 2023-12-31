# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------


# pylint: disable=protected-access

# pylint: disable=no-self-use


import argparse
from collections import defaultdict
from knack.util import CLIError


class AddAccessPolicies(argparse._AppendAction):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        super(AddAccessPolicies, self).__call__(parser, namespace, action, option_string)

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'object-id':
                d['object_id'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter access-policies. All possible keys are: object-id'
                    .format(k)
                )

        return d


class AddCosmosDbConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.cosmos_db_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'offer-throughput':
                d['offer_throughput'] = v[0]
            elif kl == 'key-vault-key-uri':
                d['key_vault_key_uri'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter cosmos-db-configuration. All possible keys are:'
                    ' offer-throughput, key-vault-key-uri'.format(k)
                )

        return d


class AddAuthenticationConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.authentication_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'authority':
                d['authority'] = v[0]
            elif kl == 'audience':
                d['audience'] = v[0]
            elif kl == 'smart-proxy-enabled':
                d['smart_proxy_enabled'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter authentication-configuration. All possible keys are:'
                    ' authority, audience, smart-proxy-enabled'.format(k)
                )

        return d


class AddCorsConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.cors_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'origins':
                d['origins'] = v
            elif kl == 'headers':
                d['headers'] = v
            elif kl == 'methods':
                d['methods'] = v
            elif kl == 'max-age':
                d['max_age'] = v[0]
            elif kl == 'allow-credentials':
                d['allow_credentials'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter cors-configuration. All possible keys are: origins,'
                    ' headers, methods, max-age, allow-credentials'.format(k)
                )
        return d


class AddServicesOciArtifacts(argparse._AppendAction):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        super(AddServicesOciArtifacts, self).__call__(parser, namespace, action, option_string)

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'login-server':
                d['login_server'] = v[0]
            elif kl == 'image-name':
                d['image_name'] = v[0]
            elif kl == 'digest':
                d['digest'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter oci-artifacts. All possible keys are: login-server,'
                    ' image-name, digest'.format(k)
                )
        return d


class AddPrivateEndpointConnections(argparse._AppendAction):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        super(AddPrivateEndpointConnections, self).__call__(parser, namespace, action, option_string)

    def get_action(self, values, option_string):  # pylint: disable=no-self-use
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'status':
                d['status'] = v[0]
            elif kl == 'description':
                d['description'] = v[0]
            elif kl == 'actions-required':
                d['actions_required'] = v[0]
        return d


class AddPrivateLinkServiceConnectionState(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.private_link_service_connection_state = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'status':
                d['status'] = v[0]
            elif kl == 'description':
                d['description'] = v[0]
            elif kl == 'actions-required':
                d['actions_required'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter private-link-service-connection-state. All possible'
                    ' keys are: status, description, actions-required'.format(k)
                )
        return d


class AddIngestionEndpointConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.ingestion_endpoint_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'event-hub-name':
                d['event_hub_name'] = v[0]
            elif kl == 'consumer-group':
                d['consumer_group'] = v[0]
            elif kl == 'fully-qualified-event-hub-namespace':
                d['fully_qualified_event_hub_namespace'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter ingestion-endpoint-configuration. All possible keys'
                    ' are: event-hub-name, consumer-group, fully-qualified-event-hub-namespace'.format(k)
                )
        return d


class AddFhirservicesAccessPolicies(argparse._AppendAction):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        super(AddFhirservicesAccessPolicies, self).__call__(parser, namespace, action, option_string)

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'object-id':
                d['object_id'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter access-policies. All possible keys are: object-id'
                    .format(k)
                )
        return d


class AddFhirservicesAuthenticationConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.authentication_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'authority':
                d['authority'] = v[0]
            elif kl == 'audience':
                d['audience'] = v[0]
            elif kl == 'smart-proxy-enabled':
                d['smart_proxy_enabled'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter authentication-configuration. All possible keys are:'
                    ' authority, audience, smart-proxy-enabled'.format(k)
                )
        return d


class AddFhirservicesCorsConfiguration(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.cors_configuration = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'origins':
                d['origins'] = v
            elif kl == 'headers':
                d['headers'] = v
            elif kl == 'methods':
                d['methods'] = v
            elif kl == 'max-age':
                d['max_age'] = v[0]
            elif kl == 'allow-credentials':
                d['allow_credentials'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter cors-configuration. All possible keys are: origins,'
                    ' headers, methods, max-age, allow-credentials'.format(k)
                )
        return d


class AddResourceTypeOverrides(argparse.Action):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        namespace.resource_type_overrides = action

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            v = properties[k]
            d[k] = v[0]
        return d


class AddFhirservicesOciArtifacts(argparse._AppendAction):
    def __call__(self, parser, namespace, values, option_string=None):
        action = self.get_action(values, option_string)
        super(AddFhirservicesOciArtifacts, self).__call__(parser, namespace, action, option_string)

    def get_action(self, values, option_string):
        try:
            properties = defaultdict(list)
            for (k, v) in (x.split('=', 1) for x in values):
                properties[k].append(v)
            properties = dict(properties)
        except ValueError:
            raise CLIError('usage error: {} [KEY=VALUE ...]'.format(option_string))
        d = {}
        for k in properties:
            kl = k.lower()
            v = properties[k]
            if kl == 'login-server':
                d['login_server'] = v[0]
            elif kl == 'image-name':
                d['image_name'] = v[0]
            elif kl == 'digest':
                d['digest'] = v[0]
            else:
                raise CLIError(
                    'Unsupported Key {} is provided for parameter oci-artifacts. All possible keys are: login-server,'
                    ' image-name, digest'.format(k)
                )
        return d
