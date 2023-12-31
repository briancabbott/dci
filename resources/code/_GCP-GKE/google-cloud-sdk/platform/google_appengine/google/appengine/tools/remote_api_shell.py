#!/usr/bin/env python
#
# Copyright 2007 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""An interactive python shell that uses remote_api.

Usage:
  %prog [-s HOSTNAME] [-p PATH] [--secure] [APPID]

If the -s HOSTNAME flag is not specified, the APPID must be specified.
"""





from google.appengine.tools import os_compat

import atexit
import code
import getpass
import optparse
import os
import sys



from google.appengine.api import memcache
from google.appengine.api import urlfetch
from google.appengine.api import users
from google.appengine.ext import db
from google.appengine.ext import ndb
from google.appengine.ext.remote_api import remote_api_stub

from google.appengine.tools import appengine_rpc

try:
  import readline
except ImportError:
  readline = None

HISTORY_PATH = os.path.expanduser('~/.remote_api_shell_history')
DEFAULT_PATH = '/_ah/remote_api'
BANNER = """App Engine remote_api shell
Python %s
The db, ndb, users, urlfetch, and memcache modules are imported.\
""" % sys.version


def auth_func():
  return (input('Email: '), getpass.getpass('Password: '))


def remote_api_shell(servername, appid, path, secure,
                     rpc_server_factory, oauth2=False):
  """Actually run the remote_api_shell."""


  if oauth2:
    remote_api_stub.ConfigureRemoteApiForOAuth(servername, path,
                                               secure=secure, app_id=appid)
  else:
    remote_api_stub.ConfigureRemoteApi(appid, path, auth_func,
                                       servername=servername,
                                       save_cookies=True, secure=secure,
                                       rpc_server_factory=rpc_server_factory)
  remote_api_stub.MaybeInvokeAuthentication()


  os.environ['SERVER_SOFTWARE'] = 'Development (remote_api_shell)/1.0'

  if not appid:

    appid = os.environ['APPLICATION_ID']
  sys.ps1 = '%s> ' % appid
  if readline is not None:

    readline.parse_and_bind('tab: complete')
    atexit.register(lambda: readline.write_history_file(HISTORY_PATH))
    if os.path.exists(HISTORY_PATH):
      readline.read_history_file(HISTORY_PATH)


  if '' not in sys.path:
    sys.path.insert(0, '')

  preimported_locals = {
      'memcache': memcache,
      'urlfetch': urlfetch,
      'users': users,
      'db': db,
      'ndb': ndb,
      }


  code.interact(banner=BANNER, local=preimported_locals)


def main(_):
  """Parse arguments and run shell."""
  parser = optparse.OptionParser(usage=__doc__)
  parser.add_option('-s', '--server', dest='server',
                    help=('The hostname your app is deployed on. '
                          'Defaults to <app_id>.appspot.com.'))
  parser.add_option('-p', '--path', dest='path',
                    help=('The path on the server to the remote_api handler. '
                          'Defaults to %s.' % DEFAULT_PATH))
  parser.add_option('--secure', dest='secure', action='store_true',
                    default=False, help=('Use HTTPS when communicating '
                                         'with the server.'))
  (options, args) = parser.parse_args()


  if ((not options.server and not args) or len(args) > 2
      or (options.path and len(args) > 1)):
    parser.print_usage(sys.stderr)
    if len(args) > 2:
      print('Unexpected arguments: %s' % args[2:], file=sys.stderr)
    elif options.path and len(args) > 1:
      print(sys.stderr, 'Path specified twice.', file=sys.stderr)
    sys.exit(1)


  servername = options.server
  appid = None
  path = options.path or DEFAULT_PATH
  if args:
    if servername:

      appid = args[0]
    else:

      servername = '%s.appspot.com' % args[0]
    if len(args) == 2:

      path = args[1]
  remote_api_shell(
      servername=servername,
      appid=appid,
      path=path,
      secure=options.secure,
      rpc_server_factory=appengine_rpc.HttpRpcServer,
      oauth2=True)


if __name__ == '__main__':
  main(sys.argv)
