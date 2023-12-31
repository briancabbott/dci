// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

namespace Azure.ResourceManager.SecurityCenter.Samples
{
    public partial class Sample_IngestionSettingResource
    {
        // Get a ingestion setting on subscription
        [NUnit.Framework.Test]
        [NUnit.Framework.Ignore("Only verifying that the sample builds")]
        public async Task Get_GetAIngestionSettingOnSubscription()
        {
            // Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/GetIngestionSetting_example.json
            // this example is just showing the usage of "IngestionSettings_Get" operation, for the dependent resources, they will have to be created separately.

            // get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
            TokenCredential cred = new DefaultAzureCredential();
            // authenticate your client
            ArmClient client = new ArmClient(cred);

            // this example assumes you already have this IngestionSettingResource created on azure
            // for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
            string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
            string ingestionSettingName = "default";
            ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
            IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

            // invoke the operation
            IngestionSettingResource result = await ingestionSetting.GetAsync();

            // the variable result is a resource, you could call other operations on this instance as well
            // but just for demo, we get its data from this resource instance
            IngestionSettingData resourceData = result.Data;
            // for demo we just print out the id
            Console.WriteLine($"Succeeded on id: {resourceData.Id}");
        }

        // Create an ingestion setting for subscription
        [NUnit.Framework.Test]
        [NUnit.Framework.Ignore("Only verifying that the sample builds")]
        public async Task Update_CreateAnIngestionSettingForSubscription()
        {
            // Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/CreateIngestionSetting_example.json
            // this example is just showing the usage of "IngestionSettings_Create" operation, for the dependent resources, they will have to be created separately.

            // get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
            TokenCredential cred = new DefaultAzureCredential();
            // authenticate your client
            ArmClient client = new ArmClient(cred);

            // this example assumes you already have this IngestionSettingResource created on azure
            // for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
            string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
            string ingestionSettingName = "default";
            ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
            IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

            // invoke the operation
            IngestionSettingData data = new IngestionSettingData();
            ArmOperation<IngestionSettingResource> lro = await ingestionSetting.UpdateAsync(WaitUntil.Completed, data);
            IngestionSettingResource result = lro.Value;

            // the variable result is a resource, you could call other operations on this instance as well
            // but just for demo, we get its data from this resource instance
            IngestionSettingData resourceData = result.Data;
            // for demo we just print out the id
            Console.WriteLine($"Succeeded on id: {resourceData.Id}");
        }

        // Delete an ingestion setting for the subscription
        [NUnit.Framework.Test]
        [NUnit.Framework.Ignore("Only verifying that the sample builds")]
        public async Task Delete_DeleteAnIngestionSettingForTheSubscription()
        {
            // Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/DeleteIngestionSetting_example.json
            // this example is just showing the usage of "IngestionSettings_Delete" operation, for the dependent resources, they will have to be created separately.

            // get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
            TokenCredential cred = new DefaultAzureCredential();
            // authenticate your client
            ArmClient client = new ArmClient(cred);

            // this example assumes you already have this IngestionSettingResource created on azure
            // for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
            string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
            string ingestionSettingName = "default";
            ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
            IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

            // invoke the operation
            await ingestionSetting.DeleteAsync(WaitUntil.Completed);

            Console.WriteLine($"Succeeded");
        }

        // List ingestion setting tokens
        [NUnit.Framework.Test]
        [NUnit.Framework.Ignore("Only verifying that the sample builds")]
        public async Task GetTokens_ListIngestionSettingTokens()
        {
            // Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListTokensIngestionSetting_example.json
            // this example is just showing the usage of "IngestionSettings_ListTokens" operation, for the dependent resources, they will have to be created separately.

            // get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
            TokenCredential cred = new DefaultAzureCredential();
            // authenticate your client
            ArmClient client = new ArmClient(cred);

            // this example assumes you already have this IngestionSettingResource created on azure
            // for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
            string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
            string ingestionSettingName = "default";
            ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
            IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

            // invoke the operation
            IngestionSettingToken result = await ingestionSetting.GetTokensAsync();

            Console.WriteLine($"Succeeded: {result}");
        }

        // List connection strings for ingesting security data and logs
        [NUnit.Framework.Test]
        [NUnit.Framework.Ignore("Only verifying that the sample builds")]
        public async Task GetConnectionStrings_ListConnectionStringsForIngestingSecurityDataAndLogs()
        {
            // Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/ListConnectionStrings_example.json
            // this example is just showing the usage of "IngestionSettings_ListConnectionStrings" operation, for the dependent resources, they will have to be created separately.

            // get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
            TokenCredential cred = new DefaultAzureCredential();
            // authenticate your client
            ArmClient client = new ArmClient(cred);

            // this example assumes you already have this IngestionSettingResource created on azure
            // for more information of creating IngestionSettingResource, please refer to the document of IngestionSettingResource
            string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
            string ingestionSettingName = "default";
            ResourceIdentifier ingestionSettingResourceId = IngestionSettingResource.CreateResourceIdentifier(subscriptionId, ingestionSettingName);
            IngestionSettingResource ingestionSetting = client.GetIngestionSettingResource(ingestionSettingResourceId);

            // invoke the operation and iterate over the result
            await foreach (IngestionConnectionString item in ingestionSetting.GetConnectionStringsAsync())
            {
                Console.WriteLine($"Succeeded: {item}");
            }

            Console.WriteLine($"Succeeded");
        }
    }
}
