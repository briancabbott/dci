"""Generated message classes for apikeys version v2alpha1.

Manages the API keys associated with developer projects.
"""
# NOTE: This file is autogenerated and should not be edited by hand.

from __future__ import absolute_import

from apitools.base.protorpclite import messages as _messages
from apitools.base.py import encoding
from apitools.base.py import extra_types


package = 'apikeys'


class ApikeysGetKeyStringNameRequest(_messages.Message):
  r"""A ApikeysGetKeyStringNameRequest object.

  Fields:
    keyString: Required. Finds the project that owns the key string value.
  """

  keyString = _messages.StringField(1)


class ApikeysOperationsGetRequest(_messages.Message):
  r"""A ApikeysOperationsGetRequest object.

  Fields:
    name: The name of the operation resource.
  """

  name = _messages.StringField(1, required=True)


class ApikeysProjectsKeysCloneRequest(_messages.Message):
  r"""A ApikeysProjectsKeysCloneRequest object.

  Fields:
    name: Required. The resource name of the Api key to be cloned under same
      parent. `apikeys.keys.get permission` and `apikeys.keys.create
      permission` are required for parent resource.
    v2alpha1CloneKeyRequest: A V2alpha1CloneKeyRequest resource to be passed
      as the request body.
  """

  name = _messages.StringField(1, required=True)
  v2alpha1CloneKeyRequest = _messages.MessageField('V2alpha1CloneKeyRequest', 2)


class ApikeysProjectsKeysCreateRequest(_messages.Message):
  r"""A ApikeysProjectsKeysCreateRequest object.

  Fields:
    parent: Required. The project for which this API key will be created.
    v2alpha1ApiKey: A V2alpha1ApiKey resource to be passed as the request
      body.
  """

  parent = _messages.StringField(1, required=True)
  v2alpha1ApiKey = _messages.MessageField('V2alpha1ApiKey', 2)


class ApikeysProjectsKeysDeleteRequest(_messages.Message):
  r"""A ApikeysProjectsKeysDeleteRequest object.

  Fields:
    name: Required. The resource name of the API key to be deleted.
  """

  name = _messages.StringField(1, required=True)


class ApikeysProjectsKeysGetKeyStringRequest(_messages.Message):
  r"""A ApikeysProjectsKeysGetKeyStringRequest object.

  Fields:
    name: Required. The resource name of the API key to be retrieved.
  """

  name = _messages.StringField(1, required=True)


class ApikeysProjectsKeysGetRequest(_messages.Message):
  r"""A ApikeysProjectsKeysGetRequest object.

  Fields:
    name: Required. The resource name of the API key to be retrieved.
  """

  name = _messages.StringField(1, required=True)


class ApikeysProjectsKeysListRequest(_messages.Message):
  r"""A ApikeysProjectsKeysListRequest object.

  Fields:
    filter: Optional. Only list keys that conform to the given filter. The
      allowed filter strings are `state:ACTIVE` and `state:DELETED`. By
      default, ListKeys will return active keys.
    pageSize: Optional. Specifies the maximum number of results to be returned
      at a time.
    pageToken: Optional. Requests a specific page of results.
    parent: Required. Lists all API keys associated with this project.
  """

  filter = _messages.StringField(1)
  pageSize = _messages.IntegerField(2, variant=_messages.Variant.INT32)
  pageToken = _messages.StringField(3)
  parent = _messages.StringField(4, required=True)


class ApikeysProjectsKeysPatchRequest(_messages.Message):
  r"""A ApikeysProjectsKeysPatchRequest object.

  Fields:
    name: Required. The resource name of the API key to be modified.
    updateMask: Required. The field mask specifies which fields should be
      updated as part of this request. All other fields will be ignored.
      Allowed field mask: `display_name` and `restrictions`
    v2alpha1ApiKey: A V2alpha1ApiKey resource to be passed as the request
      body.
  """

  name = _messages.StringField(1, required=True)
  updateMask = _messages.StringField(2)
  v2alpha1ApiKey = _messages.MessageField('V2alpha1ApiKey', 3)


class ApikeysProjectsKeysUndeleteRequest(_messages.Message):
  r"""A ApikeysProjectsKeysUndeleteRequest object.

  Fields:
    name: Required. The resource name of the API key to be undeleted.
    v2alpha1UndeleteKeyRequest: A V2alpha1UndeleteKeyRequest resource to be
      passed as the request body.
  """

  name = _messages.StringField(1, required=True)
  v2alpha1UndeleteKeyRequest = _messages.MessageField('V2alpha1UndeleteKeyRequest', 2)


class Operation(_messages.Message):
  r"""This resource represents a long-running operation that is the result of
  a network API call.

  Messages:
    MetadataValue: Service-specific metadata associated with the operation. It
      typically contains progress information and common metadata such as
      create time. Some services might not provide such metadata. Any method
      that returns a long-running operation should document the metadata type,
      if any.
    ResponseValue: The normal, successful response of the operation. If the
      original method returns no data on success, such as `Delete`, the
      response is `google.protobuf.Empty`. If the original method is standard
      `Get`/`Create`/`Update`, the response should be the resource. For other
      methods, the response should have the type `XxxResponse`, where `Xxx` is
      the original method name. For example, if the original method name is
      `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.

  Fields:
    done: If the value is `false`, it means the operation is still in
      progress. If `true`, the operation is completed, and either `error` or
      `response` is available.
    error: The error result of the operation in case of failure or
      cancellation.
    metadata: Service-specific metadata associated with the operation. It
      typically contains progress information and common metadata such as
      create time. Some services might not provide such metadata. Any method
      that returns a long-running operation should document the metadata type,
      if any.
    name: The server-assigned name, which is only unique within the same
      service that originally returns it. If you use the default HTTP mapping,
      the `name` should be a resource name ending with
      `operations/{unique_id}`.
    response: The normal, successful response of the operation. If the
      original method returns no data on success, such as `Delete`, the
      response is `google.protobuf.Empty`. If the original method is standard
      `Get`/`Create`/`Update`, the response should be the resource. For other
      methods, the response should have the type `XxxResponse`, where `Xxx` is
      the original method name. For example, if the original method name is
      `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
  """

  @encoding.MapUnrecognizedFields('additionalProperties')
  class MetadataValue(_messages.Message):
    r"""Service-specific metadata associated with the operation. It typically
    contains progress information and common metadata such as create time.
    Some services might not provide such metadata. Any method that returns a
    long-running operation should document the metadata type, if any.

    Messages:
      AdditionalProperty: An additional property for a MetadataValue object.

    Fields:
      additionalProperties: Properties of the object. Contains field @type
        with type URL.
    """

    class AdditionalProperty(_messages.Message):
      r"""An additional property for a MetadataValue object.

      Fields:
        key: Name of the additional property.
        value: A extra_types.JsonValue attribute.
      """

      key = _messages.StringField(1)
      value = _messages.MessageField('extra_types.JsonValue', 2)

    additionalProperties = _messages.MessageField('AdditionalProperty', 1, repeated=True)

  @encoding.MapUnrecognizedFields('additionalProperties')
  class ResponseValue(_messages.Message):
    r"""The normal, successful response of the operation. If the original
    method returns no data on success, such as `Delete`, the response is
    `google.protobuf.Empty`. If the original method is standard
    `Get`/`Create`/`Update`, the response should be the resource. For other
    methods, the response should have the type `XxxResponse`, where `Xxx` is
    the original method name. For example, if the original method name is
    `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.

    Messages:
      AdditionalProperty: An additional property for a ResponseValue object.

    Fields:
      additionalProperties: Properties of the object. Contains field @type
        with type URL.
    """

    class AdditionalProperty(_messages.Message):
      r"""An additional property for a ResponseValue object.

      Fields:
        key: Name of the additional property.
        value: A extra_types.JsonValue attribute.
      """

      key = _messages.StringField(1)
      value = _messages.MessageField('extra_types.JsonValue', 2)

    additionalProperties = _messages.MessageField('AdditionalProperty', 1, repeated=True)

  done = _messages.BooleanField(1)
  error = _messages.MessageField('Status', 2)
  metadata = _messages.MessageField('MetadataValue', 3)
  name = _messages.StringField(4)
  response = _messages.MessageField('ResponseValue', 5)


class StandardQueryParameters(_messages.Message):
  r"""Query parameters accepted by all methods.

  Enums:
    FXgafvValueValuesEnum: V1 error format.
    AltValueValuesEnum: Data format for response.

  Fields:
    f__xgafv: V1 error format.
    access_token: OAuth access token.
    alt: Data format for response.
    callback: JSONP
    fields: Selector specifying which fields to include in a partial response.
    key: API key. Your API key identifies your project and provides you with
      API access, quota, and reports. Required unless you provide an OAuth 2.0
      token.
    oauth_token: OAuth 2.0 token for the current user.
    prettyPrint: Returns response with indentations and line breaks.
    quotaUser: Available to use for quota purposes for server-side
      applications. Can be any arbitrary string assigned to a user, but should
      not exceed 40 characters.
    trace: A tracing token of the form "token:<tokenid>" to include in api
      requests.
    uploadType: Legacy upload protocol for media (e.g. "media", "multipart").
    upload_protocol: Upload protocol for media (e.g. "raw", "multipart").
  """

  class AltValueValuesEnum(_messages.Enum):
    r"""Data format for response.

    Values:
      json: Responses with Content-Type of application/json
      media: Media download with context-dependent Content-Type
      proto: Responses with Content-Type of application/x-protobuf
    """
    json = 0
    media = 1
    proto = 2

  class FXgafvValueValuesEnum(_messages.Enum):
    r"""V1 error format.

    Values:
      _1: v1 error format
      _2: v2 error format
    """
    _1 = 0
    _2 = 1

  f__xgafv = _messages.EnumField('FXgafvValueValuesEnum', 1)
  access_token = _messages.StringField(2)
  alt = _messages.EnumField('AltValueValuesEnum', 3, default='json')
  callback = _messages.StringField(4)
  fields = _messages.StringField(5)
  key = _messages.StringField(6)
  oauth_token = _messages.StringField(7)
  prettyPrint = _messages.BooleanField(8, default=True)
  quotaUser = _messages.StringField(9)
  trace = _messages.StringField(10)
  uploadType = _messages.StringField(11)
  upload_protocol = _messages.StringField(12)


class Status(_messages.Message):
  r"""The `Status` type defines a logical error model that is suitable for
  different programming environments, including REST APIs and RPC APIs. It is
  used by [gRPC](https://github.com/grpc). Each `Status` message contains
  three pieces of data: error code, error message, and error details. You can
  find out more about this error model and how to work with it in the [API
  Design Guide](https://cloud.google.com/apis/design/errors).

  Messages:
    DetailsValueListEntry: A DetailsValueListEntry object.

  Fields:
    code: The status code, which should be an enum value of google.rpc.Code.
    details: A list of messages that carry the error details. There is a
      common set of message types for APIs to use.
    message: A developer-facing error message, which should be in English. Any
      user-facing error message should be localized and sent in the
      google.rpc.Status.details field, or localized by the client.
  """

  @encoding.MapUnrecognizedFields('additionalProperties')
  class DetailsValueListEntry(_messages.Message):
    r"""A DetailsValueListEntry object.

    Messages:
      AdditionalProperty: An additional property for a DetailsValueListEntry
        object.

    Fields:
      additionalProperties: Properties of the object. Contains field @type
        with type URL.
    """

    class AdditionalProperty(_messages.Message):
      r"""An additional property for a DetailsValueListEntry object.

      Fields:
        key: Name of the additional property.
        value: A extra_types.JsonValue attribute.
      """

      key = _messages.StringField(1)
      value = _messages.MessageField('extra_types.JsonValue', 2)

    additionalProperties = _messages.MessageField('AdditionalProperty', 1, repeated=True)

  code = _messages.IntegerField(1, variant=_messages.Variant.INT32)
  details = _messages.MessageField('DetailsValueListEntry', 2, repeated=True)
  message = _messages.StringField(3)


class V2alpha1AndroidApplication(_messages.Message):
  r"""Identifier of an Android application for API key use.

  Fields:
    packageName: The package name of the application.
    sha1Fingerprint: The SHA1 fingerprint of the application. For example,
      both sha1 formats are acceptable as input:
      DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or
      DA39A3EE5E6B4B0D3255BFEF95601890AFD80709. Output format is the latter.
  """

  packageName = _messages.StringField(1)
  sha1Fingerprint = _messages.StringField(2)


class V2alpha1AndroidKeyRestrictions(_messages.Message):
  r"""Key restrictions that are specific to android keys.

  Fields:
    allowedApplications: A list of Android applications that are allowed to
      make API calls with this key.
  """

  allowedApplications = _messages.MessageField('V2alpha1AndroidApplication', 1, repeated=True)


class V2alpha1ApiKey(_messages.Message):
  r"""The representation of an API key managed by the `ApiKeys` API. An API
  key is used for programmatic access to a project by a service account.

  Enums:
    StateValueValuesEnum: Whether Key is active or deleted.

  Fields:
    createTime: Output only. A timestamp identifying the time this API key was
      originally created.
    creator: Email address of the user who originally created this API key.
    displayName: Human-readable display name of this API key. Modifiable by
      user.
    keyString: Output only. An encrypted and signed value held by this API
      key. This field will only be accessed through limited methods.
    name: Output only. The resource name of the api key. Api key names have
      the form `projects/123/keys/abcd_123_dx`
    restrictions: Key restrictions.
    state: Whether Key is active or deleted.
    updateTime: Output only. A timestamp identifying the time this API key was
      last updated.
  """

  class StateValueValuesEnum(_messages.Enum):
    r"""Whether Key is active or deleted.

    Values:
      API_KEY_STATE_UNSPECIFIED: Default value indicates that the field is
        unset. It should never be used.
      ACTIVE: The key is in active state and can be used.
      DELETED: The key is marked as deleted. Deleted keys can be listed and
        undeleted within 30 days of deletion.
    """
    API_KEY_STATE_UNSPECIFIED = 0
    ACTIVE = 1
    DELETED = 2

  createTime = _messages.StringField(1)
  creator = _messages.StringField(2)
  displayName = _messages.StringField(3)
  keyString = _messages.StringField(4)
  name = _messages.StringField(5)
  restrictions = _messages.MessageField('V2alpha1Restrictions', 6)
  state = _messages.EnumField('StateValueValuesEnum', 7)
  updateTime = _messages.StringField(8)


class V2alpha1ApiTarget(_messages.Message):
  r"""A restriction for a specific service and optionally one or multiple
  specific methods. Both fields are not case sensitive.

  Fields:
    methods: Optional. List of one or more methods that can be called. If
      empty, all methods for the service are allowed. A wildcard (*) can be
      used as the last symbol. Valid examples:
      google.cloud.translate.v2.TranslateService.GetSupportedLanguage
      TranslateText Get* google.cloud.translate.v2.TranslateService.Get*
    service: The service for this restriction. It should be canonical One
      Platform service name, for example:
      google.cloud.translate.v2.TranslateService.
  """

  methods = _messages.StringField(1, repeated=True)
  service = _messages.StringField(2)


class V2alpha1BrowserKeyRestrictions(_messages.Message):
  r"""Key restrictions that are specific to browser keys.

  Fields:
    allowedReferrers: A list of regular expressions for the referrer URLs that
      are allowed when making an API call with this key.
  """

  allowedReferrers = _messages.StringField(1, repeated=True)


class V2alpha1CloneKeyRequest(_messages.Message):
  r"""Request message for `CloneKey` method."""


class V2alpha1GetKeyStringNameResponse(_messages.Message):
  r"""Response message for `GetKeyStringName` method.

  Fields:
    name: The resource name
    parent: The parent that owns the key with the value specified in the
      request.
  """

  name = _messages.StringField(1)
  parent = _messages.StringField(2)


class V2alpha1GetKeyStringResponse(_messages.Message):
  r"""Response message for `GetKeyString` method.

  Fields:
    keyString: An encrypted and signed value of the key.
  """

  keyString = _messages.StringField(1)


class V2alpha1IosKeyRestrictions(_messages.Message):
  r"""Key restrictions that are specific to iOS keys.

  Fields:
    allowedBundleIds: A list of bundle IDs that are allowed when making API
      calls with this key.
  """

  allowedBundleIds = _messages.StringField(1, repeated=True)


class V2alpha1ListKeysResponse(_messages.Message):
  r"""Response message for `ListKeys` method.

  Fields:
    keys: A list of API keys.
    nextPageToken: The pagination token for the next page of results.
  """

  keys = _messages.MessageField('V2alpha1ApiKey', 1, repeated=True)
  nextPageToken = _messages.StringField(2)


class V2alpha1Restrictions(_messages.Message):
  r"""Restrictions for all types of API Keys.

  Fields:
    androidKeyRestrictions: Key restriction that are specific to android keys.
      Android apps
    apiTargets: A restriction for a specific service and optionally one or
      multiple specific methods. Requests will be allowed if they match any of
      these restrictions. If no restrictions are specified, all targets are
      allowed.
    browserKeyRestrictions: Key restrictions that are specific to browser
      keys. Referer
    iosKeyRestrictions: Key restriction that are specific to iOS keys. IOS app
      id
    serverKeyRestrictions: Key restrictions that are specific to server keys.
      Allowed ips
  """

  androidKeyRestrictions = _messages.MessageField('V2alpha1AndroidKeyRestrictions', 1)
  apiTargets = _messages.MessageField('V2alpha1ApiTarget', 2, repeated=True)
  browserKeyRestrictions = _messages.MessageField('V2alpha1BrowserKeyRestrictions', 3)
  iosKeyRestrictions = _messages.MessageField('V2alpha1IosKeyRestrictions', 4)
  serverKeyRestrictions = _messages.MessageField('V2alpha1ServerKeyRestrictions', 5)


class V2alpha1ServerKeyRestrictions(_messages.Message):
  r"""Key restrictions that are specific to server keys.

  Fields:
    allowedIps: A list of the caller IP addresses that are allowed when making
      an API call with this key.
  """

  allowedIps = _messages.StringField(1, repeated=True)


class V2alpha1UndeleteKeyRequest(_messages.Message):
  r"""Request message for `UndeleteKey` method."""


encoding.AddCustomJsonFieldMapping(
    StandardQueryParameters, 'f__xgafv', '$.xgafv')
encoding.AddCustomJsonEnumMapping(
    StandardQueryParameters.FXgafvValueValuesEnum, '_1', '1')
encoding.AddCustomJsonEnumMapping(
    StandardQueryParameters.FXgafvValueValuesEnum, '_2', '2')
