//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

import "io"

// DefenderSettingsClientCreateOrUpdateResponse contains the response from method DefenderSettingsClient.CreateOrUpdate.
type DefenderSettingsClientCreateOrUpdateResponse struct {
	// IoT Defender settings
	DefenderSettingsModel
}

// DefenderSettingsClientDeleteResponse contains the response from method DefenderSettingsClient.Delete.
type DefenderSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DefenderSettingsClientDownloadManagerActivationResponse contains the response from method DefenderSettingsClient.DownloadManagerActivation.
type DefenderSettingsClientDownloadManagerActivationResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// DefenderSettingsClientGetResponse contains the response from method DefenderSettingsClient.Get.
type DefenderSettingsClientGetResponse struct {
	// IoT Defender settings
	DefenderSettingsModel
}

// DefenderSettingsClientListResponse contains the response from method DefenderSettingsClient.List.
type DefenderSettingsClientListResponse struct {
	// List of IoT Defender settings
	DefenderSettingsList
}

// DefenderSettingsClientPackageDownloadsResponse contains the response from method DefenderSettingsClient.PackageDownloads.
type DefenderSettingsClientPackageDownloadsResponse struct {
	// Information about package downloads
	PackageDownloads
}

// DeviceGroupsClientCreateOrUpdateResponse contains the response from method DeviceGroupsClient.CreateOrUpdate.
type DeviceGroupsClientCreateOrUpdateResponse struct {
	// Device group
	DeviceGroupModel
}

// DeviceGroupsClientDeleteResponse contains the response from method DeviceGroupsClient.Delete.
type DeviceGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeviceGroupsClientGetResponse contains the response from method DeviceGroupsClient.Get.
type DeviceGroupsClientGetResponse struct {
	// Device group
	DeviceGroupModel
}

// DeviceGroupsClientListResponse contains the response from method DeviceGroupsClient.NewListPager.
type DeviceGroupsClientListResponse struct {
	// List of device groups
	DeviceGroupList
}

// DevicesClientGetResponse contains the response from method DevicesClient.Get.
type DevicesClientGetResponse struct {
	// Device
	DeviceModel
}

// DevicesClientListResponse contains the response from method DevicesClient.NewListPager.
type DevicesClientListResponse struct {
	// List of devices
	DeviceList
}

// LocationsClientGetResponse contains the response from method LocationsClient.Get.
type LocationsClientGetResponse struct {
	// IoT Defender location
	LocationModel
}

// LocationsClientListResponse contains the response from method LocationsClient.NewListPager.
type LocationsClientListResponse struct {
	// List of Defender for IoT locations
	LocationList
}

// OnPremiseSensorsClientCreateOrUpdateResponse contains the response from method OnPremiseSensorsClient.CreateOrUpdate.
type OnPremiseSensorsClientCreateOrUpdateResponse struct {
	// On-premise IoT sensor
	OnPremiseSensor
}

// OnPremiseSensorsClientDeleteResponse contains the response from method OnPremiseSensorsClient.Delete.
type OnPremiseSensorsClientDeleteResponse struct {
	// placeholder for future response values
}

// OnPremiseSensorsClientDownloadActivationResponse contains the response from method OnPremiseSensorsClient.DownloadActivation.
type OnPremiseSensorsClientDownloadActivationResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// OnPremiseSensorsClientDownloadResetPasswordResponse contains the response from method OnPremiseSensorsClient.DownloadResetPassword.
type OnPremiseSensorsClientDownloadResetPasswordResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// OnPremiseSensorsClientGetResponse contains the response from method OnPremiseSensorsClient.Get.
type OnPremiseSensorsClientGetResponse struct {
	// On-premise IoT sensor
	OnPremiseSensor
}

// OnPremiseSensorsClientListResponse contains the response from method OnPremiseSensorsClient.List.
type OnPremiseSensorsClientListResponse struct {
	// List of on-premise IoT sensors
	OnPremiseSensorsList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Paged list of operation resources
	OperationList
}

// SensorsClientCreateOrUpdateResponse contains the response from method SensorsClient.CreateOrUpdate.
type SensorsClientCreateOrUpdateResponse struct {
	// IoT sensor model
	SensorModel
}

// SensorsClientDeleteResponse contains the response from method SensorsClient.Delete.
type SensorsClientDeleteResponse struct {
	// placeholder for future response values
}

// SensorsClientDownloadActivationResponse contains the response from method SensorsClient.DownloadActivation.
type SensorsClientDownloadActivationResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// SensorsClientDownloadResetPasswordResponse contains the response from method SensorsClient.DownloadResetPassword.
type SensorsClientDownloadResetPasswordResponse struct {
	// Body contains the streaming response.
	Body io.ReadCloser
}

// SensorsClientGetResponse contains the response from method SensorsClient.Get.
type SensorsClientGetResponse struct {
	// IoT sensor model
	SensorModel
}

// SensorsClientListResponse contains the response from method SensorsClient.List.
type SensorsClientListResponse struct {
	// List of IoT sensors
	SensorsList
}

// SensorsClientTriggerTiPackageUpdateResponse contains the response from method SensorsClient.TriggerTiPackageUpdate.
type SensorsClientTriggerTiPackageUpdateResponse struct {
	// placeholder for future response values
}

// SitesClientCreateOrUpdateResponse contains the response from method SitesClient.CreateOrUpdate.
type SitesClientCreateOrUpdateResponse struct {
	// IoT site model
	SiteModel
}

// SitesClientDeleteResponse contains the response from method SitesClient.Delete.
type SitesClientDeleteResponse struct {
	// placeholder for future response values
}

// SitesClientGetResponse contains the response from method SitesClient.Get.
type SitesClientGetResponse struct {
	// IoT site model
	SiteModel
}

// SitesClientListResponse contains the response from method SitesClient.List.
type SitesClientListResponse struct {
	// List of IoT sites
	SitesList
}
