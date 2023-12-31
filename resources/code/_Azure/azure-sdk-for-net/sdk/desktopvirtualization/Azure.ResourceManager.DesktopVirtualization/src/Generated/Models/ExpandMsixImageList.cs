// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

using System.Collections.Generic;
using Azure.Core;

namespace Azure.ResourceManager.DesktopVirtualization.Models
{
    /// <summary> List of MSIX package properties retrieved from MSIX Image expansion. </summary>
    internal partial class ExpandMsixImageList
    {
        /// <summary> Initializes a new instance of <see cref="ExpandMsixImageList"/>. </summary>
        internal ExpandMsixImageList()
        {
            Value = new ChangeTrackingList<ExpandMsixImage>();
        }

        /// <summary> Initializes a new instance of <see cref="ExpandMsixImageList"/>. </summary>
        /// <param name="value"> List of MSIX package properties from give MSIX Image. </param>
        /// <param name="nextLink"> Link to the next page of results. </param>
        internal ExpandMsixImageList(IReadOnlyList<ExpandMsixImage> value, string nextLink)
        {
            Value = value;
            NextLink = nextLink;
        }

        /// <summary> List of MSIX package properties from give MSIX Image. </summary>
        public IReadOnlyList<ExpandMsixImage> Value { get; }
        /// <summary> Link to the next page of results. </summary>
        public string NextLink { get; }
    }
}
