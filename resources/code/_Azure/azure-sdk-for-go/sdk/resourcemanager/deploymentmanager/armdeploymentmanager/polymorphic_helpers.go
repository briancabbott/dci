//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

import "encoding/json"

func unmarshalAuthenticationClassification(rawMsg json.RawMessage) (AuthenticationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AuthenticationClassification
	switch m["type"] {
	case "Sas":
		b = &SasAuthentication{}
	default:
		b = &Authentication{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalHealthCheckStepAttributesClassification(rawMsg json.RawMessage) (HealthCheckStepAttributesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b HealthCheckStepAttributesClassification
	switch m["type"] {
	case "REST":
		b = &RestHealthCheckStepAttributes{}
	default:
		b = &HealthCheckStepAttributes{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalRestRequestAuthenticationClassification(rawMsg json.RawMessage) (RestRequestAuthenticationClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b RestRequestAuthenticationClassification
	switch m["type"] {
	case string(RestAuthTypeAPIKey):
		b = &APIKeyAuthentication{}
	case string(RestAuthTypeRolloutIdentity):
		b = &RolloutIdentityAuthentication{}
	default:
		b = &RestRequestAuthentication{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalStepPropertiesClassification(rawMsg json.RawMessage) (StepPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b StepPropertiesClassification
	switch m["stepType"] {
	case string(StepTypeHealthCheck):
		b = &HealthCheckStepProperties{}
	case string(StepTypeWait):
		b = &WaitStepProperties{}
	default:
		b = &StepProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
