﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Diagnostics;
using System.Diagnostics.Tracing;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Core.Diagnostics;
using Azure.Core.Tests;
using NUnit.Framework;

namespace Azure.Identity.Tests
{
    [TestFixture(true)]
    [TestFixture(false)]
    public class ProcessRunnerTests
    {
        public ProcessRunnerTests(bool isAsync)
        {
            IsAsync = isAsync;
        }

        public bool IsAsync { get; }

        [Test]
        public async Task ProcessRunnerSucceeded([Values(true, false)] bool logPII)
        {
            var output = "Test output";
            var fileName = "someFileName.exe";
            var args = "arg1 arg2";
            var pi = new ProcessStartInfo(fileName, args);
            var process = new TestProcess { StartInfo = pi, Output = output };
            string result;
            string log = string.Empty;
            using AzureEventSourceListener listener = new AzureEventSourceListener((args, s) =>
            {
                log = $"{args.EventName} {s}";
            }, EventLevel.Verbose);
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), logPII, default);
            result = await Run(runner);

            Assert.AreEqual(output, result);
            if (logPII)
            {
                Assert.That(log, Contains.Substring(fileName));
                Assert.That(log, Contains.Substring(args));
                Assert.That(log, Contains.Substring(nameof(AzureIdentityEventSource.ProcessRunnerInformational)));
            }
            else
            {
                Assert.That(log, Does.Not.Contain(fileName));
                Assert.That(log, Does.Not.Contain(args));
                Assert.That(log, Does.Not.Contain(nameof(AzureIdentityEventSource.ProcessRunnerInformational)));
            }
        }

        [Test]
        public async Task ProcessRunnerRealProcessSucceeded()
        {
            var output = "Test output";
            var process = CreateRealProcess($"echo {output}", $"echo {output}");
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);
            var result = await Run(runner);

            Assert.AreEqual(output, result);
        }

        [Test]
        public async Task ProcessRunnerRealProcessLargeOutputSucceeded()
        {
            var output = string.Concat(Enumerable.Repeat("Test output", 500));
            var command = $"echo {output}";
            var process = CreateRealProcess(command, command);
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);
            var result = await Run(runner);

            Assert.AreEqual(output, result);
        }

        [Test]
        public async Task ProcessRunnerRealProcessMultipleLinesSucceeded()
        {
            var output = "Test output";
            var command = string.Join("&&", Enumerable.Repeat($"echo {output}", 100));
            var process = CreateRealProcess(command, command);
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);
            var result = await Run(runner);

            Assert.AreEqual(string.Join(Environment.NewLine, Enumerable.Repeat(output, 100)), result);
        }

        [Test]
        public void ProcessRunnerProcessFailsToStart()
        {
            var process = new TestProcess { FailedToStart = true };
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);

            Assert.CatchAsync<InvalidOperationException>(async () => await Run(runner));
        }

        [Test]
        public void ProcessRunnerCanceledByTimeout()
        {
            var cts = new CancellationTokenSource();
            var process = new TestProcess { Output = "Test output", Timeout = 5000 };
            using var runner = new ProcessRunner(process, TimeSpan.FromMilliseconds(100), false, cts.Token);

            Assert.CatchAsync<OperationCanceledException>(async () => await Run(runner));
        }

        [Test]
        public void ProcessRunnerCanceledByCancellationToken()
        {
            var cts = new CancellationTokenSource();
            var process = new TestProcess { Output = "Test output", Timeout = 5000 };
            using var runner = new ProcessRunner(process, TimeSpan.FromMilliseconds(5000), false, cts.Token);
            cts.CancelAfter(100);

            Assert.CatchAsync<OperationCanceledException>(async () => await Run(runner));
        }

        [Test]
        public void ProcessRunnerCreatedOnCanceled()
        {
            var process = new TestProcess { Output = "Test output", Timeout = 5000 };
            var cancellationToken = new CancellationToken(true);
            using var runner = new ProcessRunner(process, TimeSpan.FromMilliseconds(5000), false, cancellationToken);

            Assert.CatchAsync<OperationCanceledException>(async () => await Run(runner));
        }

        [Test]
        public void ProcessRunnerCanceledBeforeRun()
        {
            var cts = new CancellationTokenSource();
            var process = new TestProcess { Output = "Test output", Timeout = 5000 };
            using var runner = new ProcessRunner(process, TimeSpan.FromMilliseconds(5000), false, cts.Token);

            cts.Cancel();

            Assert.CatchAsync<OperationCanceledException>(async () => await Run(runner));
        }

        [Test]
        public async Task ProcessRunnerCanceledFinished()
        {
            var cts = new CancellationTokenSource();
            var process = new TestProcess { Output = "Test output" };
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(5000), false, cts.Token);
            await Run(runner);
            cts.Cancel();
        }

        [Test]
        public void ProcessRunnerFailedWithErrorMessage([Values(true, false)] bool logPII)
        {
            string log = string.Empty;
            using AzureEventSourceListener listener = new AzureEventSourceListener((args, s) =>
            {
                log = $"{args.EventName} {s}";
            }, EventLevel.Verbose);
            var error = "Test error";
            var process = new TestProcess { Error = error };
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), logPII, default);

            var exception = Assert.CatchAsync<InvalidOperationException>(async () => await Run(runner));
            Assert.AreEqual(error, exception.Message);
            if (logPII)
            {
                Assert.That(log, Contains.Substring(error));
                Assert.That(log, Contains.Substring(nameof(AzureIdentityEventSource.ProcessRunnerError)));
            }
            else
            {
                Assert.That(log, Does.Not.Contain(error));
                Assert.That(log, Does.Not.Contain(nameof(AzureIdentityEventSource.ProcessRunnerError)));
            }
        }

        [Test]
        public void ProcessRunnerRealProcessFailedWithErrorMessage()
        {
            var error = "Test error";
            var process = CreateRealProcess($"echo {error} 1>&2 & exit 1", $">&2 echo {error} & exit 1");
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);

            var exception = Assert.CatchAsync<InvalidOperationException>(async () => await Run(runner));
            Assert.AreEqual(error, exception.Message.Trim());
        }

        [Test]
        public void ProcessRunnerRealProcessFailedWithLargeErrorMessage()
        {
            var error = string.Concat(Enumerable.Repeat("Test error", 500));
            ;
            var process = CreateRealProcess($"echo {error} 1>&2 & exit 1", $">&2 echo {error} & exit 1");
            using var runner = new ProcessRunner(process, TimeSpan.FromSeconds(30), false, default);

            var exception = Assert.CatchAsync<InvalidOperationException>(async () => await Run(runner));
            Assert.AreEqual(error, exception.Message.Trim());
        }

        [Test]
        public void ProcessRunnerFailedOnKillProcess()
        {
            var output = "Test output";
            var process = new TestProcess { Output = output, ExceptionOnProcessKill = new Win32Exception(1), Timeout = 5000 };
            using var runner = new ProcessRunner(process, TimeSpan.FromMilliseconds(50), false, default);

            var exception = Assert.CatchAsync<Win32Exception>(async () => await Run(runner));
            Assert.AreEqual(1, exception.NativeErrorCode);
        }

        private async Task<string> Run(ProcessRunner runner) => IsAsync ? await runner.RunAsync() : runner.Run();

        private static IProcess CreateRealProcess(string windowsCommand, string nonWindowsCommand)
        {
            var isWindows = RuntimeInformation.IsOSPlatform(OSPlatform.Windows);
            var fileName = isWindows ? "cmd" : "sh";
            var arguments = isWindows ? $"/C \"{windowsCommand}\"" : $"-c \"{nonWindowsCommand}\"";

            return ProcessService.Default.Create(new ProcessStartInfo { FileName = fileName, Arguments = arguments, ErrorDialog = false, CreateNoWindow = true });
        }
    }
}
