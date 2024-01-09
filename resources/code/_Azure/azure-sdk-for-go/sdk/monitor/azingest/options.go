//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azingest

// UploadOptions contains the optional parameters for the Client.Upload method.
type UploadOptions struct {
	// If the bytes of the "logs" parameter are already gzipped, set ContentEncoding to "gzip"
	ContentEncoding *string
}
