//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/timeseriesinsights/armtimeseriesinsights"
)

func unmarshalEnvironmentCreateOrUpdateParametersClassification(rawMsg json.RawMessage) (armtimeseriesinsights.EnvironmentCreateOrUpdateParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armtimeseriesinsights.EnvironmentCreateOrUpdateParametersClassification
	switch m["kind"] {
	case string(armtimeseriesinsights.EnvironmentKindGen1):
		b = &armtimeseriesinsights.Gen1EnvironmentCreateOrUpdateParameters{}
	case string(armtimeseriesinsights.EnvironmentKindGen2):
		b = &armtimeseriesinsights.Gen2EnvironmentCreateOrUpdateParameters{}
	default:
		b = &armtimeseriesinsights.EnvironmentCreateOrUpdateParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEnvironmentUpdateParametersClassification(rawMsg json.RawMessage) (armtimeseriesinsights.EnvironmentUpdateParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armtimeseriesinsights.EnvironmentUpdateParametersClassification
	switch m["kind"] {
	case string(armtimeseriesinsights.EnvironmentKindGen1):
		b = &armtimeseriesinsights.Gen1EnvironmentUpdateParameters{}
	case string(armtimeseriesinsights.EnvironmentKindGen2):
		b = &armtimeseriesinsights.Gen2EnvironmentUpdateParameters{}
	default:
		b = &armtimeseriesinsights.EnvironmentUpdateParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEventSourceCreateOrUpdateParametersClassification(rawMsg json.RawMessage) (armtimeseriesinsights.EventSourceCreateOrUpdateParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armtimeseriesinsights.EventSourceCreateOrUpdateParametersClassification
	switch m["kind"] {
	case string(armtimeseriesinsights.EventSourceKindMicrosoftEventHub):
		b = &armtimeseriesinsights.EventHubEventSourceCreateOrUpdateParameters{}
	case string(armtimeseriesinsights.EventSourceKindMicrosoftIoTHub):
		b = &armtimeseriesinsights.IoTHubEventSourceCreateOrUpdateParameters{}
	default:
		b = &armtimeseriesinsights.EventSourceCreateOrUpdateParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalEventSourceUpdateParametersClassification(rawMsg json.RawMessage) (armtimeseriesinsights.EventSourceUpdateParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armtimeseriesinsights.EventSourceUpdateParametersClassification
	switch m["kind"] {
	case string(armtimeseriesinsights.EventSourceKindMicrosoftEventHub):
		b = &armtimeseriesinsights.EventHubEventSourceUpdateParameters{}
	case string(armtimeseriesinsights.EventSourceKindMicrosoftIoTHub):
		b = &armtimeseriesinsights.IoTHubEventSourceUpdateParameters{}
	default:
		b = &armtimeseriesinsights.EventSourceUpdateParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
