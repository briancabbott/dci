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


class MpnsCredential(Model):
    """Description of a NotificationHub MpnsCredential.

    :param mpns_certificate: The MPNS certificate.
    :type mpns_certificate: str
    :param certificate_key: The certificate key for this credential.
    :type certificate_key: str
    :param thumbprint: The MPNS certificate Thumbprint
    :type thumbprint: str
    """

    _attribute_map = {
        'mpns_certificate': {'key': 'properties.mpnsCertificate', 'type': 'str'},
        'certificate_key': {'key': 'properties.certificateKey', 'type': 'str'},
        'thumbprint': {'key': 'properties.thumbprint', 'type': 'str'},
    }

    def __init__(self, *, mpns_certificate: str=None, certificate_key: str=None, thumbprint: str=None, **kwargs) -> None:
        super(MpnsCredential, self).__init__(**kwargs)
        self.mpns_certificate = mpns_certificate
        self.certificate_key = certificate_key
        self.thumbprint = thumbprint
