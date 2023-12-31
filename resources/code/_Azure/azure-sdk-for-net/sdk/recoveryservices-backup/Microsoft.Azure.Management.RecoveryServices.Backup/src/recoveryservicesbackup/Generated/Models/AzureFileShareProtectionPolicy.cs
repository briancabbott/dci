// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.RecoveryServices.Backup.Models
{
    using Newtonsoft.Json;
    using System.Collections;
    using System.Collections.Generic;
    using System.Linq;

    /// <summary>
    /// AzureStorage backup policy.
    /// </summary>
    [Newtonsoft.Json.JsonObject("AzureStorage")]
    public partial class AzureFileShareProtectionPolicy : ProtectionPolicy
    {
        /// <summary>
        /// Initializes a new instance of the AzureFileShareProtectionPolicy
        /// class.
        /// </summary>
        public AzureFileShareProtectionPolicy()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the AzureFileShareProtectionPolicy
        /// class.
        /// </summary>
        /// <param name="protectedItemsCount">Number of items associated with
        /// this policy.</param>
        /// <param name="resourceGuardOperationRequests">ResourceGuard
        /// Operation Requests</param>
        /// <param name="workLoadType">Type of workload for the backup
        /// management. Possible values include: 'Invalid', 'VM', 'FileFolder',
        /// 'AzureSqlDb', 'SQLDB', 'Exchange', 'Sharepoint', 'VMwareVM',
        /// 'SystemState', 'Client', 'GenericDataSource', 'SQLDataBase',
        /// 'AzureFileShare', 'SAPHanaDatabase', 'SAPAseDatabase',
        /// 'SAPHanaDBInstance'</param>
        /// <param name="schedulePolicy">Backup schedule specified as part of
        /// backup policy.</param>
        /// <param name="retentionPolicy">Retention policy with the details on
        /// backup copy retention ranges.</param>
        /// <param name="timeZone">TimeZone optional input as string. For
        /// example: TimeZone = "Pacific Standard Time".</param>
        public AzureFileShareProtectionPolicy(int? protectedItemsCount = default(int?), IList<string> resourceGuardOperationRequests = default(IList<string>), string workLoadType = default(string), SchedulePolicy schedulePolicy = default(SchedulePolicy), RetentionPolicy retentionPolicy = default(RetentionPolicy), string timeZone = default(string))
            : base(protectedItemsCount, resourceGuardOperationRequests)
        {
            WorkLoadType = workLoadType;
            SchedulePolicy = schedulePolicy;
            RetentionPolicy = retentionPolicy;
            TimeZone = timeZone;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets type of workload for the backup management. Possible
        /// values include: 'Invalid', 'VM', 'FileFolder', 'AzureSqlDb',
        /// 'SQLDB', 'Exchange', 'Sharepoint', 'VMwareVM', 'SystemState',
        /// 'Client', 'GenericDataSource', 'SQLDataBase', 'AzureFileShare',
        /// 'SAPHanaDatabase', 'SAPAseDatabase', 'SAPHanaDBInstance'
        /// </summary>
        [JsonProperty(PropertyName = "workLoadType")]
        public string WorkLoadType { get; set; }

        /// <summary>
        /// Gets or sets backup schedule specified as part of backup policy.
        /// </summary>
        [JsonProperty(PropertyName = "schedulePolicy")]
        public SchedulePolicy SchedulePolicy { get; set; }

        /// <summary>
        /// Gets or sets retention policy with the details on backup copy
        /// retention ranges.
        /// </summary>
        [JsonProperty(PropertyName = "retentionPolicy")]
        public RetentionPolicy RetentionPolicy { get; set; }

        /// <summary>
        /// Gets or sets timeZone optional input as string. For example:
        /// TimeZone = "Pacific Standard Time".
        /// </summary>
        [JsonProperty(PropertyName = "timeZone")]
        public string TimeZone { get; set; }

    }
}
