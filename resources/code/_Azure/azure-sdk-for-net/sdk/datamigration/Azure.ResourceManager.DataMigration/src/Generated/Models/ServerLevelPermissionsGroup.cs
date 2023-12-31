// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.DataMigration.Models
{
    /// <summary> Permission group for validations. These groups will run a set of permissions for validating user activity. Select the permission group for the activity that you are performing. </summary>
    public enum ServerLevelPermissionsGroup
    {
        /// <summary> Default. </summary>
        Default,
        /// <summary> MigrationFromSqlServerToAzureDB. </summary>
        MigrationFromSqlServerToAzureDB,
        /// <summary> MigrationFromSqlServerToAzureMI. </summary>
        MigrationFromSqlServerToAzureMI,
        /// <summary> MigrationFromMySQLToAzureDBForMySQL. </summary>
        MigrationFromMySqlToAzureDBForMySql,
        /// <summary> MigrationFromSqlServerToAzureVM. </summary>
        MigrationFromSqlServerToAzureVm
    }
}
