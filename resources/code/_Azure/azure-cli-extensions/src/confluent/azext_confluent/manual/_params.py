# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#
# Code generated by Microsoft (R) AutoRest Code Generator.
# Changes may cause incorrect behavior and will be lost if the code is
# regenerated.
# --------------------------------------------------------------------------


def load_arguments(self, _):
    with self.argument_context('confluent organization create') as c:
        c.argument('publisher_id', type=str, help='Publisher Id', default='confluentinc', arg_group='Offer Detail')
        c.argument('offer_id', type=str, help='Offer Id', default='confluent-cloud-azure-prod',
                   arg_group='Offer Detail')
        c.argument('plan_id', type=str, help='Offer Plan Id', arg_group='Offer Detail')
        c.argument('plan_name', type=str, help='Offer Plan Name', arg_group='Offer Detail')
        c.argument('term_unit', type=str, help='Offer Plan Term unit', arg_group='Offer Detail')

    with self.argument_context('confluent organization delete') as c:
        c.argument('yes', options_list=['--yes', '-y'], action='store_true', help='Do not prompt for confirmation.')

    with self.argument_context('confluent offer-detail show') as c:
        c.argument('publisher_id', type=str, help='Publisher Id', default='confluentinc')
        c.argument('offer_id', type=str, help='Offer Id', default='confluent-cloud-azure-prod',
                   arg_group='Offer Detail')
