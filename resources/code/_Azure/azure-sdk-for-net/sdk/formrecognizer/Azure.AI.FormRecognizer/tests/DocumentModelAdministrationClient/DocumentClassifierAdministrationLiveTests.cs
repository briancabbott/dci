﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Net;
using System.Threading.Tasks;
using Azure.Core.TestFramework;
using NUnit.Framework;

namespace Azure.AI.FormRecognizer.DocumentAnalysis.Tests
{
    /// <summary>
    /// The suite of tests for the document classifier methods in the <see cref="DocumentModelAdministrationClient"/> class.
    /// </summary>
    /// <remarks>
    /// These tests have a dependency on live Azure services and may incur costs for the associated
    /// Azure subscription.
    /// </remarks>
    [ServiceVersion(Min = DocumentAnalysisClientOptions.ServiceVersion.V2023_07_31)]
    public class DocumentClassifierAdministrationLiveTests : DocumentAnalysisLiveTestBase
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="DocumentClassifierAdministrationLiveTests"/> class.
        /// </summary>
        /// <param name="isAsync">A flag used by the Azure Core Test Framework to differentiate between tests for asynchronous and synchronous methods.</param>
        public DocumentClassifierAdministrationLiveTests(bool isAsync, DocumentAnalysisClientOptions.ServiceVersion serviceVersion)
            : base(isAsync, serviceVersion)
        {
        }

        #region Build

        [RecordedTest]
        public async Task BuildDocumentClassifierCanAuthenticateWithTokenCredential()
        {
            var client = CreateDocumentModelAdministrationClient(useTokenCredential: true);
            var classifierId = Recording.GenerateId();

            var trainingFilesUri = new Uri(TestEnvironment.ClassifierTrainingSasUrl);
            var sourceA = new BlobContentSource(trainingFilesUri) { Prefix = "IRS-1040-A/train" };
            var sourceB = new BlobContentSource(trainingFilesUri) { Prefix = "IRS-1040-B/train" };

            var documentTypes = new Dictionary<string, ClassifierDocumentTypeDetails>()
            {
                { "IRS-1040-A", new ClassifierDocumentTypeDetails(sourceA) },
                { "IRS-1040-B", new ClassifierDocumentTypeDetails(sourceB) }
            };

            BuildDocumentClassifierOperation operation = null;

            try
            {
                operation = await client.BuildDocumentClassifierAsync(WaitUntil.Completed, documentTypes, classifierId);
            }
            finally
            {
                if (operation != null && operation.HasValue)
                {
                    await client.DeleteDocumentClassifierAsync(classifierId);
                }
            }

            // Sanity check to make sure we got an actual response back from the service.
            Assert.AreEqual(classifierId, operation.Value.ClassifierId);
        }

        [RecordedTest]
        public async Task BuildDocumentClassifierWithAzureBlobContentSource()
        {
            var client = CreateDocumentModelAdministrationClient();
            var classifierId = Recording.GenerateId();
            var description = $"This classifier was generated by a .NET test: {nameof(BuildDocumentClassifierWithAzureBlobContentSource)}";
            var startTime = Recording.UtcNow;

            var trainingFilesUri = new Uri(TestEnvironment.ClassifierTrainingSasUrl);
            var sourceA = new BlobContentSource(trainingFilesUri) { Prefix = "IRS-1040-A/train" };
            var sourceB = new BlobContentSource(trainingFilesUri) { Prefix = "IRS-1040-B/train" };

            var documentTypes = new Dictionary<string, ClassifierDocumentTypeDetails>()
            {
                { "IRS-1040-A", new ClassifierDocumentTypeDetails(sourceA) },
                { "IRS-1040-B", new ClassifierDocumentTypeDetails(sourceB) }
            };

            BuildDocumentClassifierOperation operation = null;

            try
            {
                operation = await client.BuildDocumentClassifierAsync(WaitUntil.Completed, documentTypes, classifierId, description);
            }
            finally
            {
                if (operation != null && operation.HasValue)
                {
                    await client.DeleteDocumentClassifierAsync(classifierId);
                }
            }

            Assert.IsTrue(operation.HasValue);

            DocumentClassifierDetails classifier = operation.Value;

            Assert.AreEqual(classifierId, classifier.ClassifierId);
            Assert.AreEqual(description, classifier.Description);
            Assert.AreEqual(ServiceVersionString, classifier.ServiceVersion);
            Assert.Greater(classifier.CreatedOn, startTime);
            Assert.Greater(classifier.ExpiresOn, classifier.CreatedOn);

            DocumentAssert.AreEquivalent(documentTypes, classifier.DocumentTypes);
        }

        [RecordedTest]
        public async Task BuildDocumentClassifierWithAzureBlobFileListSource()
        {
            var client = CreateDocumentModelAdministrationClient();
            var classifierId = Recording.GenerateId();
            var description = $"This classifier was generated by a .NET test: {nameof(BuildDocumentClassifierWithAzureBlobFileListSource)}";
            var startTime = Recording.UtcNow;

            var trainingFilesUri = new Uri(TestEnvironment.ClassifierTrainingSasUrl);
            var sourceA = new BlobFileListContentSource(trainingFilesUri, "IRS-1040-A.jsonl");
            var sourceB = new BlobFileListContentSource(trainingFilesUri, "IRS-1040-B.jsonl");

            var documentTypes = new Dictionary<string, ClassifierDocumentTypeDetails>()
            {
                { "IRS-1040-A", new ClassifierDocumentTypeDetails(sourceA) },
                { "IRS-1040-B", new ClassifierDocumentTypeDetails(sourceB) }
            };

            BuildDocumentClassifierOperation operation = null;

            try
            {
                operation = await client.BuildDocumentClassifierAsync(WaitUntil.Completed, documentTypes, classifierId, description);
            }
            finally
            {
                if (operation != null && operation.HasValue)
                {
                    await client.DeleteDocumentClassifierAsync(classifierId);
                }
            }

            Assert.IsTrue(operation.HasValue);

            DocumentClassifierDetails classifier = operation.Value;

            Assert.AreEqual(classifierId, classifier.ClassifierId);
            Assert.AreEqual(description, classifier.Description);
            Assert.AreEqual(ServiceVersionString, classifier.ServiceVersion);
            Assert.Greater(classifier.CreatedOn, startTime);
            Assert.Greater(classifier.ExpiresOn, classifier.CreatedOn);

            DocumentAssert.AreEquivalent(documentTypes, classifier.DocumentTypes);
        }

        #endregion Build

        #region Get

        [RecordedTest]
        [TestCase(true)]
        [TestCase(false)]
        public async Task GetDocumentClassifier(bool useTokenCredential)
        {
            var client = CreateDocumentModelAdministrationClient(useTokenCredential);
            var classifierId = Recording.GenerateId();
            var description = $"This classifier was generated by a .NET test: {nameof(GetDocumentClassifier)}";
            await using var disposableClassifier = await BuildDisposableDocumentClassifierAsync(classifierId, description);

            DocumentClassifierDetails expected = disposableClassifier.Value;
            DocumentClassifierDetails classifier = await client.GetDocumentClassifierAsync(classifierId);

            DocumentAssert.AreEqual(expected, classifier);
        }

        [RecordedTest]
        public void GetDocumentClassifierThrowsWhenClassifierDoesNotExist()
        {
            var client = CreateDocumentModelAdministrationClient();
            var fakeId = "00000000-0000-0000-0000-000000000000";

            RequestFailedException ex = Assert.ThrowsAsync<RequestFailedException>(async () => await client.GetDocumentClassifierAsync(fakeId));
            Assert.AreEqual("NotFound", ex.ErrorCode);
        }

        #endregion Get

        #region List

        [RecordedTest]
        [TestCase(true)]
        [TestCase(false)]
        public async Task GetDocumentClassifiers(bool useTokenCredential)
        {
            var client = CreateDocumentModelAdministrationClient(useTokenCredential);
            var classifierIds = new string[] { Recording.GenerateId(), Recording.GenerateId() };
            var description = $"This classifier was generated by a .NET test: {nameof(GetDocumentClassifiers)}";

            await using var disposableClassifier0 = await BuildDisposableDocumentClassifierAsync(classifierIds[0], description);
            await using var disposableClassifier1 = await BuildDisposableDocumentClassifierAsync(classifierIds[1], description);

            var idMapping = new Dictionary<string, DocumentClassifierDetails>();
            var expectedIdMapping = new Dictionary<string, DocumentClassifierDetails>()
            {
                { classifierIds[0], disposableClassifier0.Value },
                { classifierIds[1], disposableClassifier1.Value }
            };

            await foreach (DocumentClassifierDetails classifier in client.GetDocumentClassifiersAsync())
            {
                if (expectedIdMapping.ContainsKey(classifier.ClassifierId))
                {
                    idMapping.Add(classifier.ClassifierId, classifier);
                }

                if (idMapping.Count == expectedIdMapping.Count)
                {
                    break;
                }
            }

            foreach (string id in expectedIdMapping.Keys)
            {
                Assert.True(idMapping.ContainsKey(id));

                DocumentClassifierDetails classifier = idMapping[id];
                DocumentClassifierDetails expected = expectedIdMapping[id];

                DocumentAssert.AreEqual(expected, classifier);
            }
        }

        #endregion

        #region Delete

        [RecordedTest]
        [TestCase(true)]
        [TestCase(false)]
        public async Task DeleteDocumentClassifier(bool useTokenCredential)
        {
            var client = CreateDocumentModelAdministrationClient(useTokenCredential);
            var classifierId = Recording.GenerateId();

            await BuildDisposableDocumentClassifierAsync(classifierId);

            var response = await client.DeleteDocumentClassifierAsync(classifierId);

            Assert.AreEqual((int)HttpStatusCode.NoContent, response.Status);
        }

        [RecordedTest]
        public void DeleteDocumentClassifierThrowsWhenClassifierDoesNotExist()
        {
            var client = CreateDocumentModelAdministrationClient();
            var fakeId = "00000000-0000-0000-0000-000000000000";

            RequestFailedException ex = Assert.ThrowsAsync<RequestFailedException>(async () => await client.DeleteDocumentClassifierAsync(fakeId));
            Assert.AreEqual("NotFound", ex.ErrorCode);
        }

        #endregion Delete
    }
}
