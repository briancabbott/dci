# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------
from pygments.lexer import RegexLexer, words  # pylint: disable=import-error
from pygments.token import Name, Keyword, Operator, Number  # pylint: disable=import-error


# pylint: disable=too-few-public-methods
def get_az_lexer(commands):

    class AzLexer(RegexLexer):
        """
        A custom lexer for Azure CLI
        """
        try:
            tokens = {
                'root': [
                    (words(tuple(kid for kid in commands.command_tree.children),
                           prefix=r'\b',
                           suffix=r'\b'),
                     Keyword),  # top level commands
                    (words(tuple(commands.get_all_subcommands()),
                           prefix=r'\b',
                           suffix=r'\b'),
                     Keyword.Declaration),  # all other commands
                    (words(tuple(param for param in commands.completable_param + commands.global_param),
                           prefix=r'\B',
                           suffix=r'\b'),
                     Name.Class),  # parameters
                    (r'.', Keyword),  # all else
                    (r' .', Keyword),
                ]
            }
        except IOError:  # if there is no cache
            pass
    return AzLexer if AzLexer.tokens else None


class ExampleLexer(RegexLexer):
    """ Lexer for the example description """
    tokens = {
        'root': [
            (r' .', Number),
            (r'.', Number),
        ]
    }


class ToolbarLexer(RegexLexer):
    """ Lexer for the the toolbar """
    tokens = {
        'root': [
            (r' .', Operator),
            (r'.', Operator),
        ]
    }


class ScenarioLexer(RegexLexer):
    """ Lexer for the recommended scenarios """
    # The purpose of this lexer is to tokenize and apply syntax highlighting to a text based on a set of regular expressions.
    tokens = {
        'root': [
            # (r' .', Number) matches any single character followed by a space and assigns it the token type Number.
            (r' .', Number),
            # (r'.', Number) matches any single character and assigns it the token type Number
            (r'.', Number),
        ]
    }
