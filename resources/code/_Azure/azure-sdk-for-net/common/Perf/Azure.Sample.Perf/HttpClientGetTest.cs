﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using Azure.Test.Perf;
using CommandLine;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;

namespace Azure.Sample.Perf
{
    public class HttpClientGetTest : PerfTest<UriOptions>
    {
        private static HttpClient _httpClient;

        public HttpClientGetTest(UriOptions options) : base(options)
        {
        }

        public override Task GlobalSetupAsync()
        {
            if (Options.Insecure)
            {
                var httpClientHandler = new HttpClientHandler();
                httpClientHandler.ServerCertificateCustomValidationCallback = (message, cert, chain, errors) => true;
                _httpClient = new HttpClient(httpClientHandler);
            }
            else
            {
                _httpClient = new HttpClient();
            }

            return Task.CompletedTask;
        }

        public override void Run(CancellationToken cancellationToken)
        {
            _httpClient.GetStringAsync(Options.Uri).Wait();
        }

        public override async Task RunAsync(CancellationToken cancellationToken)
        {
            await _httpClient.GetStringAsync(Options.Uri);
        }
    }
}
