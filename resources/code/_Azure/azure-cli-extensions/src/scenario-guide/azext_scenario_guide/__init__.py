# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------

from azure.cli.core import AzCommandsLoader
from ._help import helps  # pylint: disable=unused-import


class ScenarioGuideCommandsLoader(AzCommandsLoader):

    def __init__(self, cli_ctx=None):
        from azure.cli.core.commands import CliCommandType
        scenario_guide_custom = CliCommandType(
            operations_tmpl='azext_scenario_guide.custom#{}')
        super().__init__(cli_ctx=cli_ctx, custom_command_type=scenario_guide_custom)

    def load_command_table(self, args):
        from .commands import load_command_table
        load_command_table(self, args)
        return self.command_table

    def load_arguments(self, command):
        from ._params import load_arguments
        load_arguments(self, command)


COMMAND_LOADER_CLS = ScenarioGuideCommandsLoader
