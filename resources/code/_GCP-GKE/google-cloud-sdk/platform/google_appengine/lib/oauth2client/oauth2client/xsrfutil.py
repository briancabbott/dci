#
# Copyright 2014 the Melange authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Helper methods for creating & verifying XSRF tokens."""

__authors__ = [
    '"Doug Coker" <dcoker@google.com>',
    '"Joe Gregorio" <jcgregorio@google.com>',
]


import base64
import hashlib
import hmac
import time

import six
from oauth2client import util


# Delimiter character
DELIMITER = b':'


# 1 hour in seconds
DEFAULT_TIMEOUT_SECS = 1*60*60


def _force_bytes(s):
  if isinstance(s, bytes):
    return s
  s = str(s)
  if isinstance(s, six.text_type):
    return s.encode('utf-8')
  return s


@util.positional(2)
def generate_token(key, user_id, action_id="", when=None):
  """Generates a URL-safe token for the given user, action, time tuple.

  Args:
    key: secret key to use.
    user_id: the user ID of the authenticated user.
    action_id: a string identifier of the action they requested
      authorization for.
    when: the time in seconds since the epoch at which the user was
      authorized for this action. If not set the current time is used.

  Returns:
    A string XSRF protection token.
  """
  when = _force_bytes(when or int(time.time()))
  digester = hmac.new(_force_bytes(key), digestmod=hashlib.md5)
  digester.update(_force_bytes(user_id))
  digester.update(DELIMITER)
  digester.update(_force_bytes(action_id))
  digester.update(DELIMITER)
  digester.update(when)
  digest = digester.digest()

  token = base64.urlsafe_b64encode(digest + DELIMITER + when)
  return token


@util.positional(3)
def validate_token(key, token, user_id, action_id="", current_time=None):
  """Validates that the given token authorizes the user for the action.

  Tokens are invalid if the time of issue is too old or if the token
  does not match what generateToken outputs (i.e. the token was forged).

  Args:
    key: secret key to use.
    token: a string of the token generated by generateToken.
    user_id: the user ID of the authenticated user.
    action_id: a string identifier of the action they requested
      authorization for.

  Returns:
    A boolean - True if the user is authorized for the action, False
    otherwise.
  """
  token = _force_bytes(token)
  if not token:
    return False
  try:
    decoded = base64.urlsafe_b64decode(token)
    token_time = int(decoded.split(DELIMITER)[-1])
  except (TypeError, ValueError):
    return False
  if current_time is None:
    current_time = time.time()
  # If the token is too old it's not valid.
  if current_time - token_time > DEFAULT_TIMEOUT_SECS:
    return False

  # The given token should match the generated one with the same time.
  expected_token = generate_token(key, user_id, action_id=action_id,
                                  when=token_time)
  if len(token) != len(expected_token):
    return False

  # Perform constant time comparison to avoid timing attacks
  different = 0
  for x, y in zip(bytearray(token), bytearray(expected_token)):
    different |= x ^ y
  return not different
