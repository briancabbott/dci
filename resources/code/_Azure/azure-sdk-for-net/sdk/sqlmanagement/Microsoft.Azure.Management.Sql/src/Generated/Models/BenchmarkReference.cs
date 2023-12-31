// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Sql.Models
{
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// SQL Vulnerability Assessment benchmark reference
    /// </summary>
    public partial class BenchmarkReference
    {
        /// <summary>
        /// Initializes a new instance of the BenchmarkReference class.
        /// </summary>
        public BenchmarkReference()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the BenchmarkReference class.
        /// </summary>
        /// <param name="benchmark">SQL Vulnerability Assessment benchmark
        /// name</param>
        /// <param name="reference">SQL Vulnerability Assessment benchmark
        /// reference.</param>
        public BenchmarkReference(string benchmark = default(string), string reference = default(string))
        {
            Benchmark = benchmark;
            Reference = reference;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets SQL Vulnerability Assessment benchmark name
        /// </summary>
        [JsonProperty(PropertyName = "benchmark")]
        public string Benchmark { get; private set; }

        /// <summary>
        /// Gets SQL Vulnerability Assessment benchmark reference.
        /// </summary>
        [JsonProperty(PropertyName = "reference")]
        public string Reference { get; private set; }

    }
}
