﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Threading.Tasks;
using Azure.Core.TestFramework;
using Azure.Analytics.Synapse.Tests;
using Azure.Identity;
using NUnit.Framework;
using Azure.Analytics.Synapse.Artifacts.Models;

namespace Azure.Analytics.Synapse.Artifacts.Samples
{
    public partial class Sample4_HelloWorldDataFlow : SamplesBase<SynapseTestEnvironment>
    {
        [Test]
        public async Task DataFlowSample()
        {
            #region Snippet:CreateDataFlowClientPrep
#if SNIPPET
            // Replace the string below with your actual endpoint url.
            string endpoint = "<my-endpoint-url>";
#else
            string endpoint = TestEnvironment.EndpointUrl;
#endif
            string dataFlowName = "Test-DataFlow";
            #endregion

            #region Snippet:CreateDataFlowClient
            DataFlowClient client = new DataFlowClient(endpoint: new Uri(endpoint), credential: new DefaultAzureCredential());
            #endregion

            #region Snippet:CreateDataFlow
            DataFlowCreateOrUpdateDataFlowOperation operation = client.StartCreateOrUpdateDataFlow(dataFlowName, new DataFlowResource(new DataFlow()));
            Response<DataFlowResource> createdDataflow = await operation.WaitForCompletionAsync();
            #endregion

            #region Snippet:RetrieveDataFlow
            DataFlowResource retrievedDataflow = client.GetDataFlow(dataFlowName);
            #endregion

            #region Snippet:ListDataFlows
            Pageable<DataFlowResource> dataFlows = client.GetDataFlowsByWorkspace();
            foreach (DataFlowResource dataflow in dataFlows)
            {
                System.Console.WriteLine(dataflow.Name);
            }
            #endregion

            #region Snippet:DeleteDataFlow
            DataFlowDeleteDataFlowOperation deleteOperation = client.StartDeleteDataFlow(dataFlowName);
            await deleteOperation.WaitForCompletionResponseAsync();
            #endregion
        }
    }
}
