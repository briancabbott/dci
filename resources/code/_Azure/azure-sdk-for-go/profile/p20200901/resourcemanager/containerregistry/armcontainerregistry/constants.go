//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

// Action - The action of virtual network rule.
type Action string

const (
	ActionAllow Action = "Allow"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionAllow,
	}
}

// DefaultAction - The default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow DefaultAction = "Allow"
	DefaultActionDeny  DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
	}
}

// ImportMode - When Force, any existing target tags will be overwritten. When NoForce, any existing target tags will fail
// the operation before any copying begins.
type ImportMode string

const (
	ImportModeForce   ImportMode = "Force"
	ImportModeNoForce ImportMode = "NoForce"
)

// PossibleImportModeValues returns the possible values for the ImportMode const type.
func PossibleImportModeValues() []ImportMode {
	return []ImportMode{
		ImportModeForce,
		ImportModeNoForce,
	}
}

// PasswordName - The password name.
type PasswordName string

const (
	PasswordNamePassword  PasswordName = "password"
	PasswordNamePassword2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns the possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{
		PasswordNamePassword,
		PasswordNamePassword2,
	}
}

// PolicyStatus - The value that indicates whether the policy is enabled or not.
type PolicyStatus string

const (
	PolicyStatusDisabled PolicyStatus = "disabled"
	PolicyStatusEnabled  PolicyStatus = "enabled"
)

// PossiblePolicyStatusValues returns the possible values for the PolicyStatus const type.
func PossiblePolicyStatusValues() []PolicyStatus {
	return []PolicyStatus{
		PolicyStatusDisabled,
		PolicyStatusEnabled,
	}
}

// ProvisioningState - The provisioning state of the container registry at the time the operation was called.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RegistryUsageUnit - The unit of measurement.
type RegistryUsageUnit string

const (
	RegistryUsageUnitBytes RegistryUsageUnit = "Bytes"
	RegistryUsageUnitCount RegistryUsageUnit = "Count"
)

// PossibleRegistryUsageUnitValues returns the possible values for the RegistryUsageUnit const type.
func PossibleRegistryUsageUnitValues() []RegistryUsageUnit {
	return []RegistryUsageUnit{
		RegistryUsageUnitBytes,
		RegistryUsageUnitCount,
	}
}

// SKUName - The SKU name of the container registry. Required for registry creation.
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNameClassic  SKUName = "Classic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNameClassic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// SKUTier - The SKU tier based on the SKU name.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierClassic  SKUTier = "Classic"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierClassic,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// TrustPolicyType - The type of trust policy.
type TrustPolicyType string

const (
	TrustPolicyTypeNotary TrustPolicyType = "Notary"
)

// PossibleTrustPolicyTypeValues returns the possible values for the TrustPolicyType const type.
func PossibleTrustPolicyTypeValues() []TrustPolicyType {
	return []TrustPolicyType{
		TrustPolicyTypeNotary,
	}
}

type WebhookAction string

const (
	WebhookActionChartDelete WebhookAction = "chart_delete"
	WebhookActionChartPush   WebhookAction = "chart_push"
	WebhookActionDelete      WebhookAction = "delete"
	WebhookActionPush        WebhookAction = "push"
	WebhookActionQuarantine  WebhookAction = "quarantine"
)

// PossibleWebhookActionValues returns the possible values for the WebhookAction const type.
func PossibleWebhookActionValues() []WebhookAction {
	return []WebhookAction{
		WebhookActionChartDelete,
		WebhookActionChartPush,
		WebhookActionDelete,
		WebhookActionPush,
		WebhookActionQuarantine,
	}
}

// WebhookStatus - The status of the webhook at the time the operation was called.
type WebhookStatus string

const (
	WebhookStatusDisabled WebhookStatus = "disabled"
	WebhookStatusEnabled  WebhookStatus = "enabled"
)

// PossibleWebhookStatusValues returns the possible values for the WebhookStatus const type.
func PossibleWebhookStatusValues() []WebhookStatus {
	return []WebhookStatus{
		WebhookStatusDisabled,
		WebhookStatusEnabled,
	}
}
