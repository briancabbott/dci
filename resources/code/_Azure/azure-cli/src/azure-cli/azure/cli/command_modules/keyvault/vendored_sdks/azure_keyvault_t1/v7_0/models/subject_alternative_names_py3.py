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
# pylint: skip-file
# flake8: noqa
from msrest.serialization import Model


class SubjectAlternativeNames(Model):
    """The subject alternate names of a X509 object.

    :param emails: Email addresses.
    :type emails: list[str]
    :param dns_names: Domain names.
    :type dns_names: list[str]
    :param upns: User principal names.
    :type upns: list[str]
    """

    _attribute_map = {
        'emails': {'key': 'emails', 'type': '[str]'},
        'dns_names': {'key': 'dns_names', 'type': '[str]'},
        'upns': {'key': 'upns', 'type': '[str]'},
    }

    def __init__(self, *, emails=None, dns_names=None, upns=None, **kwargs) -> None:
        super(SubjectAlternativeNames, self).__init__(**kwargs)
        self.emails = emails
        self.dns_names = dns_names
        self.upns = upns
