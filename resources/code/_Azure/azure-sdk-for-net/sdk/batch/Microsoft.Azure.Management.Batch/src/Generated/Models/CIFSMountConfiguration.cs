// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.Management.Batch.Models
{
    using Microsoft.Rest;
    using Newtonsoft.Json;
    using System.Linq;

    /// <summary>
    /// Information used to connect to a CIFS file system.
    /// </summary>
    public partial class CIFSMountConfiguration
    {
        /// <summary>
        /// Initializes a new instance of the CIFSMountConfiguration class.
        /// </summary>
        public CIFSMountConfiguration()
        {
            CustomInit();
        }

        /// <summary>
        /// Initializes a new instance of the CIFSMountConfiguration class.
        /// </summary>
        /// <param name="userName">The user to use for authentication against
        /// the CIFS file system.</param>
        /// <param name="source">The URI of the file system to mount.</param>
        /// <param name="relativeMountPath">The relative path on the compute
        /// node where the file system will be mounted</param>
        /// <param name="password">The password to use for authentication
        /// against the CIFS file system.</param>
        /// <param name="mountOptions">Additional command line options to pass
        /// to the mount command.</param>
        public CIFSMountConfiguration(string userName, string source, string relativeMountPath, string password, string mountOptions = default(string))
        {
            UserName = userName;
            Source = source;
            RelativeMountPath = relativeMountPath;
            MountOptions = mountOptions;
            Password = password;
            CustomInit();
        }

        /// <summary>
        /// An initialization method that performs custom operations like setting defaults
        /// </summary>
        partial void CustomInit();

        /// <summary>
        /// Gets or sets the user to use for authentication against the CIFS
        /// file system.
        /// </summary>
        [JsonProperty(PropertyName = "userName")]
        public string UserName { get; set; }

        /// <summary>
        /// Gets or sets the URI of the file system to mount.
        /// </summary>
        [JsonProperty(PropertyName = "source")]
        public string Source { get; set; }

        /// <summary>
        /// Gets or sets the relative path on the compute node where the file
        /// system will be mounted
        /// </summary>
        /// <remarks>
        /// All file systems are mounted relative to the Batch mounts
        /// directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment
        /// variable.
        /// </remarks>
        [JsonProperty(PropertyName = "relativeMountPath")]
        public string RelativeMountPath { get; set; }

        /// <summary>
        /// Gets or sets additional command line options to pass to the mount
        /// command.
        /// </summary>
        /// <remarks>
        /// These are 'net use' options in Windows and 'mount' options in
        /// Linux.
        /// </remarks>
        [JsonProperty(PropertyName = "mountOptions")]
        public string MountOptions { get; set; }

        /// <summary>
        /// Gets or sets the password to use for authentication against the
        /// CIFS file system.
        /// </summary>
        [JsonProperty(PropertyName = "password")]
        public string Password { get; set; }

        /// <summary>
        /// Validate the object.
        /// </summary>
        /// <exception cref="ValidationException">
        /// Thrown if validation fails
        /// </exception>
        public virtual void Validate()
        {
            if (UserName == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "UserName");
            }
            if (Source == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "Source");
            }
            if (RelativeMountPath == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "RelativeMountPath");
            }
            if (Password == null)
            {
                throw new ValidationException(ValidationRules.CannotBeNull, "Password");
            }
        }
    }
}
