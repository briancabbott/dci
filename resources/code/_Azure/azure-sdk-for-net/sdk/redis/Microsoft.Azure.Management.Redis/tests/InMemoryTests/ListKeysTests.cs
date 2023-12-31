﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.

using Microsoft.Azure.Management.Redis;
using Microsoft.Azure.Management.Redis.Models;
using Microsoft.Rest;
using Microsoft.Rest.Azure;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace AzureRedisCache.Tests.InMemoryTests
{
    public class ListKeysTests
    {
        [Fact]
        public void ListKeys_Basic()
        {
            string responseString = (@"{
	            ""primaryKey"": ""aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa="",
	            ""secondaryKey"": ""bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb=""
            }
            ");
            string requestIdHeader = "0d33aff8-8a4e-4565-b893-a10e52260de0";
            RedisManagementClient client = Utility.GetRedisManagementClient(responseString, requestIdHeader, HttpStatusCode.OK);
            var response = client.Redis.ListKeys(resourceGroupName: "resource-group", name: "cachename");

            Assert.Equal("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa=", response.PrimaryKey);
            Assert.Equal("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb=", response.SecondaryKey);
        }

        [Fact]
        public void ListKeys_ParametersChecking()
        {
            RedisManagementClient client = Utility.GetRedisManagementClient(null, null, HttpStatusCode.NotFound);
            Exception e = Assert.Throws<ValidationException>(() => client.Redis.ListKeys(resourceGroupName: null, name: "cachename"));
            Assert.Contains("resourceGroupName", e.Message);
            e = Assert.Throws<ValidationException>(() => client.Redis.ListKeys(resourceGroupName: "resource-group", name: null));
            Assert.Contains("name", e.Message);
        }

        [Fact]
        public void ListKeys_404()
        {
            RedisManagementClient client = Utility.GetRedisManagementClient(null, null, HttpStatusCode.NotFound);
            Assert.Throws<ErrorResponseException>(() => client.Redis.ListKeys(resourceGroupName: "resource-group", name: "cachename"));
        }

        [Fact]
        public void ListKeys_InvalidJSONFromCSM()
        {
            string responseString = (@"Exception: Any exception from CSM");
            RedisManagementClient client = Utility.GetRedisManagementClient(responseString, null, HttpStatusCode.OK);
            Assert.Throws<SerializationException>(() => client.Redis.ListKeys(resourceGroupName: "resource-group", name: "cachename"));
        }

        [Fact]
        public void ListKeys_EmptyJSONFromCSM()
        {
            string responseString = (@"{}");
            RedisManagementClient client = Utility.GetRedisManagementClient(responseString, null, HttpStatusCode.OK);
            var response = client.Redis.ListKeys(resourceGroupName: "resource-group", name: "cachename");
            Assert.Null(response.PrimaryKey);
            Assert.Null(response.SecondaryKey);
        }
    }
}
