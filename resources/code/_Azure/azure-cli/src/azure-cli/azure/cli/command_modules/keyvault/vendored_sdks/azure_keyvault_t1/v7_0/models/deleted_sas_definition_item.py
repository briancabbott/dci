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
from .sas_definition_item import SasDefinitionItem


class DeletedSasDefinitionItem(SasDefinitionItem):
    """The deleted SAS definition item containing metadata about the deleted SAS
    definition.

    Variables are only populated by the server, and will be ignored when
    sending a request.

    :ivar id: The storage SAS identifier.
    :vartype id: str
    :ivar secret_id: The storage account SAS definition secret id.
    :vartype secret_id: str
    :ivar attributes: The SAS definition management attributes.
    :vartype attributes: ~azure.keyvault.v7_0.models.SasDefinitionAttributes
    :ivar tags: Application specific metadata in the form of key-value pairs.
    :vartype tags: dict[str, str]
    :param recovery_id: The url of the recovery object, used to identify and
     recover the deleted SAS definition.
    :type recovery_id: str
    :ivar scheduled_purge_date: The time when the SAS definition is scheduled
     to be purged, in UTC
    :vartype scheduled_purge_date: datetime
    :ivar deleted_date: The time when the SAS definition was deleted, in UTC
    :vartype deleted_date: datetime
    """

    _validation = {
        'id': {'readonly': True},
        'secret_id': {'readonly': True},
        'attributes': {'readonly': True},
        'tags': {'readonly': True},
        'scheduled_purge_date': {'readonly': True},
        'deleted_date': {'readonly': True},
    }

    _attribute_map = {
        'id': {'key': 'id', 'type': 'str'},
        'secret_id': {'key': 'sid', 'type': 'str'},
        'attributes': {'key': 'attributes', 'type': 'SasDefinitionAttributes'},
        'tags': {'key': 'tags', 'type': '{str}'},
        'recovery_id': {'key': 'recoveryId', 'type': 'str'},
        'scheduled_purge_date': {'key': 'scheduledPurgeDate', 'type': 'unix-time'},
        'deleted_date': {'key': 'deletedDate', 'type': 'unix-time'},
    }

    def __init__(self, **kwargs):
        super(DeletedSasDefinitionItem, self).__init__(**kwargs)
        self.recovery_id = kwargs.get('recovery_id', None)
        self.scheduled_purge_date = None
        self.deleted_date = None
