# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------
# pylint: disable=super-with-arguments

from azure.cli.core import AzCommandsLoader

from azext_containerapp._help import helps  # pylint: disable=unused-import


class ContainerappCommandsLoader(AzCommandsLoader):

    def __init__(self, cli_ctx=None):
        from azure.cli.core.commands import CliCommandType
        containerapp_custom = CliCommandType(
            operations_tmpl='azext_containerapp.custom#{}',
            client_factory=None)
        super(ContainerappCommandsLoader, self).__init__(cli_ctx=cli_ctx,
                                                         custom_command_type=containerapp_custom)

    def load_command_table(self, args):
        from azext_containerapp.commands import load_command_table
        from azure.cli.core.aaz import load_aaz_command_table
        try:
            from . import aaz
        except ImportError:
            aaz = None
        if aaz:
            load_aaz_command_table(
                loader=self,
                aaz_pkg_name=aaz.__name__,
                args=args
            )
        load_command_table(self, args)
        return self.command_table

    def load_arguments(self, command):
        from azext_containerapp._params import load_arguments
        from azure.cli.command_modules.containerapp._params import load_arguments as load_arguments_from_core
        load_arguments_from_core(self, command)
        load_arguments(self, command)


COMMAND_LOADER_CLS = ContainerappCommandsLoader
