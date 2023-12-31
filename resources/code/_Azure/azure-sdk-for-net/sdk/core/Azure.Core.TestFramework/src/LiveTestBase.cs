﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System.Threading.Tasks;
using NUnit.Framework;

namespace Azure.Core.TestFramework
{
    /// <summary>
    /// Base class for tests that need to run live but are not recorded and don't need sync/async testing support of <see cref="ClientTestBase"/>
    /// </summary>
    /// <typeparam name="TEnvironment">The <see cref="TestEnvironment"/> implementation to use.</typeparam>
    [LiveOnly]
    public class LiveTestBase<TEnvironment> where TEnvironment : TestEnvironment, new()
    {
        protected LiveTestBase()
        {
            TestEnvironment = new TEnvironment()
            {
                Mode = RecordedTestMode.Live
            };
        }

        protected TEnvironment TestEnvironment { get; }

        [OneTimeSetUp]
        public async ValueTask WaitForEnvironment()
        {
            await TestEnvironment.WaitForEnvironmentAsync();
        }
    }
}
