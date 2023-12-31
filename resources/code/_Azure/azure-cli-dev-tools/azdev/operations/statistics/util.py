# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------

import copy

from knack.log import get_logger

from azdev.utilities import get_name_index


logger = get_logger(__name__)


def filter_modules(command_loader, modules=None, exclude=False, include_whl_extensions=False):
    modules = modules or []

    # command tables and help entries must be copied to allow for seperate linter scope
    command_table = command_loader.command_table.copy()
    command_group_table = command_loader.command_group_table.copy()
    command_loader = copy.copy(command_loader)
    command_loader.command_table = command_table
    command_loader.command_group_table = command_group_table
    name_index = get_name_index(include_whl_extensions=include_whl_extensions)

    for command_name in list(command_loader.command_table.keys()):
        try:
            source_name, _ = _get_command_source(command_name, command_loader.command_table)
        except ValueError as ex:
            # command is unrecognized
            logger.warning(ex)
            source_name = None

        try:
            long_name = name_index[source_name]
            is_specified = source_name in modules or long_name in modules
        except KeyError:
            is_specified = False
        if is_specified == exclude:
            # brute force method of ignoring commands from a module or extension
            command_loader.command_table.pop(command_name, None)

    # Remove unneeded command groups
    retained_command_groups = {' '.join(x.split(' ')[:-1]) for x in command_loader.command_table}
    excluded_command_groups = set(command_loader.command_group_table.keys()) - retained_command_groups

    for group_name in excluded_command_groups:
        command_loader.command_group_table.pop(group_name, None)

    return command_loader


def _get_command_source(command_name, command_table):
    from azure.cli.core.commands import ExtensionCommandSource  # pylint: disable=import-error
    command = command_table.get(command_name)
    # see if command is from an extension
    if isinstance(command.command_source, ExtensionCommandSource):
        return command.command_source.extension_name, True
    if command.command_source is None:
        raise ValueError('Command: `%s`, has no command source.' % command_name)
    # command is from module
    return command.command_source, False
