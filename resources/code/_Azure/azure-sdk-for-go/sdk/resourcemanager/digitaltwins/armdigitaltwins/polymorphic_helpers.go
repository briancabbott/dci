//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import "encoding/json"

func unmarshalEndpointResourcePropertiesClassification(rawMsg json.RawMessage) (EndpointResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b EndpointResourcePropertiesClassification
	switch m["endpointType"] {
	case string(EndpointTypeEventGrid):
		b = &EventGrid{}
	case string(EndpointTypeEventHub):
		b = &EventHub{}
	case string(EndpointTypeServiceBus):
		b = &ServiceBus{}
	default:
		b = &EndpointResourceProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTimeSeriesDatabaseConnectionPropertiesClassification(rawMsg json.RawMessage) (TimeSeriesDatabaseConnectionPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TimeSeriesDatabaseConnectionPropertiesClassification
	switch m["connectionType"] {
	case string(ConnectionTypeAzureDataExplorer):
		b = &AzureDataExplorerConnectionProperties{}
	default:
		b = &TimeSeriesDatabaseConnectionProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
