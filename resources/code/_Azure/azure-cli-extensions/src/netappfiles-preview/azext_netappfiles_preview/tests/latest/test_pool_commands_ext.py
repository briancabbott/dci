# --------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for license information.
# --------------------------------------------------------------------------------------------

import unittest
from azure.cli.testsdk import ScenarioTest, ResourceGroupPreparer

POOL_DEFAULT = "--service-level 'Premium' --size 4398046511104"

# No tidy up of tests required. The resource group is automatically removed


class AzureNetAppFilesExtPoolServiceScenarioTest(ScenarioTest):
    @unittest.skip('This command is available from main module now')
    @ResourceGroupPreparer(name_prefix='cli_tests_rg')
    def test_ext_create_delete_pool(self):
        account_name = self.create_random_name(prefix='cli-acc-', length=24)
        pool_name = self.create_random_name(prefix='cli-pool-', length=24)
        tags = "Tag1=Value1 Tag2=Value2"

        self.cmd("az netappfiles account create --resource-group {rg} --account-name '%s' -l 'westus2'" % account_name).get_output_in_json()
        pool = self.cmd("az netappfiles pool create --resource-group {rg} --account-name %s --pool-name %s -l 'westus2' %s --tags '%s'" % (account_name, pool_name, POOL_DEFAULT, tags)).get_output_in_json()
        assert pool['name'] == account_name + '/' + pool_name
        assert pool['tags']['Tag1'] == 'Value1'
        assert pool['tags']['Tag2'] == 'Value2'

        pool_list = self.cmd("netappfiles pool list --resource-group {rg} --account-name %s" % account_name).get_output_in_json()
        assert len(pool_list) == 1

        self.cmd("az netappfiles pool delete --resource-group {rg} --account-name '%s' --pool-name '%s'" % (account_name, pool_name))
        pool_list = self.cmd("netappfiles pool list --resource-group {rg} --account-name %s" % account_name).get_output_in_json()
        assert len(pool_list) == 0

        # and again with short forms and also unquoted
        pool = self.cmd("az netappfiles pool create -g {rg} -a %s -p %s -l 'westus2' --service-level 'Premium' --size 4398046511104 --tags '%s'" % (account_name, pool_name, tags)).get_output_in_json()
        assert pool['name'] == account_name + '/' + pool_name
        assert pool['tags']['Tag1'] == 'Value1'
        assert pool['tags']['Tag2'] == 'Value2'

        self.cmd("az netappfiles pool delete --resource-group {rg} -a %s -p %s" % (account_name, pool_name))
        pool_list = self.cmd("netappfiles pool list --resource-group {rg} -a %s" % account_name).get_output_in_json()
        assert len(pool_list) == 0

    @unittest.skip('This command is available from main module now')
    @ResourceGroupPreparer(name_prefix='cli_tests_rg')
    def test_ext_list_pools(self):
        account_name = self.create_random_name(prefix='cli', length=24)
        pools = [self.create_random_name(prefix='cli', length=24), self.create_random_name(prefix='cli', length=24)]
        self.cmd("az netappfiles account create -g {rg} -a '%s' -l 'westus2'" % account_name).get_output_in_json()

        for pool_name in pools:
            self.cmd("az netappfiles pool create -g {rg} -a '%s' -p '%s' -l 'westus2' %s --tags 'Tag1=Value1'" % (account_name, pool_name, POOL_DEFAULT)).get_output_in_json()

        pool_list = self.cmd("netappfiles pool list -g {rg} -a '%s'" % account_name).get_output_in_json()
        assert len(pool_list) == 2

        for pool_name in pools:
            self.cmd("az netappfiles pool delete -g {rg} -a %s -p %s" % (account_name, pool_name))
        pool_list = self.cmd("netappfiles pool list --resource-group {rg} -a '%s'" % account_name).get_output_in_json()
        assert len(pool_list) == 0

    @unittest.skip('This command is available from main module now')
    @ResourceGroupPreparer(name_prefix='cli_tests_rg')
    def test_ext_get_pool_by_name(self):
        account_name = self.create_random_name(prefix='cli', length=24)
        pool_name = self.create_random_name(prefix='cli', length=24)
        self.cmd("az netappfiles account create -g {rg} -a '%s' -l 'westus2'" % account_name).get_output_in_json()
        self.cmd("az netappfiles pool create -g {rg} -a %s -p %s -l 'westus2' %s" % (account_name, pool_name, POOL_DEFAULT)).get_output_in_json()

        pool = self.cmd("az netappfiles pool show --resource-group {rg} -a %s -p %s" % (account_name, pool_name)).get_output_in_json()
        assert pool['name'] == account_name + '/' + pool_name
        pool_from_id = self.cmd("az netappfiles pool show --ids %s" % pool['id']).get_output_in_json()
        assert pool_from_id['name'] == account_name + '/' + pool_name

    @unittest.skip('This command is available from main module now')
    @ResourceGroupPreparer(name_prefix='cli_tests_rg')
    def test_ext_update_pool(self):
        account_name = self.create_random_name(prefix='cli-acc-', length=24)
        pool_name = self.create_random_name(prefix='cli-pool-', length=24)
        tag = "Tag1=Value1"

        self.cmd("az netappfiles account create -g {rg} -a %s -l 'westus2'" % account_name).get_output_in_json()

        pool = self.cmd("az netappfiles pool create -g {rg} -a %s -p %s -l 'westus2' %s" % (account_name, pool_name, POOL_DEFAULT)).get_output_in_json()

        assert pool['name'] == account_name + '/' + pool_name
        pool = self.cmd("az netappfiles pool update --resource-group {rg} -a %s -p %s --tags %s --service-level 'Standard'" % (account_name, pool_name, tag)).get_output_in_json()
        assert pool['name'] == account_name + '/' + pool_name
        assert pool['serviceLevel'] == "Standard"
        assert pool['tags']['Tag1'] == 'Value1'
