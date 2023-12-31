// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.ResourceManager.Sql.Models
{
    /// <summary> Properties of the column in the table of database full schema. </summary>
    public partial class SyncFullSchemaTableColumn
    {
        /// <summary> Initializes a new instance of <see cref="SyncFullSchemaTableColumn"/>. </summary>
        internal SyncFullSchemaTableColumn()
        {
        }

        /// <summary> Initializes a new instance of <see cref="SyncFullSchemaTableColumn"/>. </summary>
        /// <param name="dataSize"> Data size of the column. </param>
        /// <param name="dataType"> Data type of the column. </param>
        /// <param name="errorId"> Error id of the column. </param>
        /// <param name="hasError"> If there is error in the table. </param>
        /// <param name="isPrimaryKey"> If it is the primary key of the table. </param>
        /// <param name="name"> Name of the column. </param>
        /// <param name="quotedName"> Quoted name of the column. </param>
        internal SyncFullSchemaTableColumn(string dataSize, string dataType, string errorId, bool? hasError, bool? isPrimaryKey, string name, string quotedName)
        {
            DataSize = dataSize;
            DataType = dataType;
            ErrorId = errorId;
            HasError = hasError;
            IsPrimaryKey = isPrimaryKey;
            Name = name;
            QuotedName = quotedName;
        }

        /// <summary> Data size of the column. </summary>
        public string DataSize { get; }
        /// <summary> Data type of the column. </summary>
        public string DataType { get; }
        /// <summary> Error id of the column. </summary>
        public string ErrorId { get; }
        /// <summary> If there is error in the table. </summary>
        public bool? HasError { get; }
        /// <summary> If it is the primary key of the table. </summary>
        public bool? IsPrimaryKey { get; }
        /// <summary> Name of the column. </summary>
        public string Name { get; }
        /// <summary> Quoted name of the column. </summary>
        public string QuotedName { get; }
    }
}
