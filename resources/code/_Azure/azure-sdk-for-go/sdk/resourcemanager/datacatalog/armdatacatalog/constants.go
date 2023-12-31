//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatacatalog

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datacatalog/armdatacatalog"
	moduleVersion = "v1.2.0"
)

// SKUType - Azure data catalog SKU.
type SKUType string

const (
	SKUTypeFree     SKUType = "Free"
	SKUTypeStandard SKUType = "Standard"
)

// PossibleSKUTypeValues returns the possible values for the SKUType const type.
func PossibleSKUTypeValues() []SKUType {
	return []SKUType{
		SKUTypeFree,
		SKUTypeStandard,
	}
}
