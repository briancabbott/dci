//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement

// ActionClassification provides polymorphic access to related types.
// Call the interface's GetAction() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Action, *AddActionGroups, *RemoveAllActionGroups
type ActionClassification interface {
	// GetAction returns the Action content of the underlying type.
	GetAction() *Action
}

// AlertsMetaDataPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetAlertsMetaDataProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AlertsMetaDataProperties, *MonitorServiceList
type AlertsMetaDataPropertiesClassification interface {
	// GetAlertsMetaDataProperties returns the AlertsMetaDataProperties content of the underlying type.
	GetAlertsMetaDataProperties() *AlertsMetaDataProperties
}

// RecurrenceClassification provides polymorphic access to related types.
// Call the interface's GetRecurrence() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DailyRecurrence, *MonthlyRecurrence, *Recurrence, *WeeklyRecurrence
type RecurrenceClassification interface {
	// GetRecurrence returns the Recurrence content of the underlying type.
	GetRecurrence() *Recurrence
}
