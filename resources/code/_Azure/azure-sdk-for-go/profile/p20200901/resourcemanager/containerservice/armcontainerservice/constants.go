//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

// AgentPoolMode - AgentPoolMode represents mode of an agent pool.
type AgentPoolMode string

const (
	AgentPoolModeSystem AgentPoolMode = "System"
	AgentPoolModeUser   AgentPoolMode = "User"
)

// PossibleAgentPoolModeValues returns the possible values for the AgentPoolMode const type.
func PossibleAgentPoolModeValues() []AgentPoolMode {
	return []AgentPoolMode{
		AgentPoolModeSystem,
		AgentPoolModeUser,
	}
}

// AgentPoolType - AgentPoolType represents types of an agent pool.
type AgentPoolType string

const (
	AgentPoolTypeAvailabilitySet         AgentPoolType = "AvailabilitySet"
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

// PossibleAgentPoolTypeValues returns the possible values for the AgentPoolType const type.
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return []AgentPoolType{
		AgentPoolTypeAvailabilitySet,
		AgentPoolTypeVirtualMachineScaleSets,
	}
}

// Code - Tells whether the cluster is Running or Stopped
type Code string

const (
	CodeRunning Code = "Running"
	CodeStopped Code = "Stopped"
)

// PossibleCodeValues returns the possible values for the Code const type.
func PossibleCodeValues() []Code {
	return []Code{
		CodeRunning,
		CodeStopped,
	}
}

// ConnectionStatus - The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

// PossibleConnectionStatusValues returns the possible values for the ConnectionStatus const type.
func PossibleConnectionStatusValues() []ConnectionStatus {
	return []ConnectionStatus{
		ConnectionStatusApproved,
		ConnectionStatusDisconnected,
		ConnectionStatusPending,
		ConnectionStatusRejected,
	}
}

// ContainerServiceOrchestratorTypes - The orchestrator to use to manage container service cluster resources. Valid values
// are Kubernetes, Swarm, DCOS, DockerCE and Custom.
type ContainerServiceOrchestratorTypes string

const (
	ContainerServiceOrchestratorTypesCustom     ContainerServiceOrchestratorTypes = "Custom"
	ContainerServiceOrchestratorTypesDCOS       ContainerServiceOrchestratorTypes = "DCOS"
	ContainerServiceOrchestratorTypesDockerCE   ContainerServiceOrchestratorTypes = "DockerCE"
	ContainerServiceOrchestratorTypesKubernetes ContainerServiceOrchestratorTypes = "Kubernetes"
	ContainerServiceOrchestratorTypesSwarm      ContainerServiceOrchestratorTypes = "Swarm"
)

// PossibleContainerServiceOrchestratorTypesValues returns the possible values for the ContainerServiceOrchestratorTypes const type.
func PossibleContainerServiceOrchestratorTypesValues() []ContainerServiceOrchestratorTypes {
	return []ContainerServiceOrchestratorTypes{
		ContainerServiceOrchestratorTypesCustom,
		ContainerServiceOrchestratorTypesDCOS,
		ContainerServiceOrchestratorTypesDockerCE,
		ContainerServiceOrchestratorTypesKubernetes,
		ContainerServiceOrchestratorTypesSwarm,
	}
}

// ContainerServiceStorageProfileTypes - Storage profile specifies what kind of storage used. Choose from StorageAccount and
// ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
type ContainerServiceStorageProfileTypes string

const (
	ContainerServiceStorageProfileTypesManagedDisks   ContainerServiceStorageProfileTypes = "ManagedDisks"
	ContainerServiceStorageProfileTypesStorageAccount ContainerServiceStorageProfileTypes = "StorageAccount"
)

// PossibleContainerServiceStorageProfileTypesValues returns the possible values for the ContainerServiceStorageProfileTypes const type.
func PossibleContainerServiceStorageProfileTypesValues() []ContainerServiceStorageProfileTypes {
	return []ContainerServiceStorageProfileTypes{
		ContainerServiceStorageProfileTypesManagedDisks,
		ContainerServiceStorageProfileTypesStorageAccount,
	}
}

// ContainerServiceVMSizeTypes - Size of agent VMs.
type ContainerServiceVMSizeTypes string

const (
	ContainerServiceVMSizeTypesStandardA1          ContainerServiceVMSizeTypes = "Standard_A1"
	ContainerServiceVMSizeTypesStandardA10         ContainerServiceVMSizeTypes = "Standard_A10"
	ContainerServiceVMSizeTypesStandardA11         ContainerServiceVMSizeTypes = "Standard_A11"
	ContainerServiceVMSizeTypesStandardA1V2        ContainerServiceVMSizeTypes = "Standard_A1_v2"
	ContainerServiceVMSizeTypesStandardA2          ContainerServiceVMSizeTypes = "Standard_A2"
	ContainerServiceVMSizeTypesStandardA2MV2       ContainerServiceVMSizeTypes = "Standard_A2m_v2"
	ContainerServiceVMSizeTypesStandardA2V2        ContainerServiceVMSizeTypes = "Standard_A2_v2"
	ContainerServiceVMSizeTypesStandardA3          ContainerServiceVMSizeTypes = "Standard_A3"
	ContainerServiceVMSizeTypesStandardA4          ContainerServiceVMSizeTypes = "Standard_A4"
	ContainerServiceVMSizeTypesStandardA4MV2       ContainerServiceVMSizeTypes = "Standard_A4m_v2"
	ContainerServiceVMSizeTypesStandardA4V2        ContainerServiceVMSizeTypes = "Standard_A4_v2"
	ContainerServiceVMSizeTypesStandardA5          ContainerServiceVMSizeTypes = "Standard_A5"
	ContainerServiceVMSizeTypesStandardA6          ContainerServiceVMSizeTypes = "Standard_A6"
	ContainerServiceVMSizeTypesStandardA7          ContainerServiceVMSizeTypes = "Standard_A7"
	ContainerServiceVMSizeTypesStandardA8          ContainerServiceVMSizeTypes = "Standard_A8"
	ContainerServiceVMSizeTypesStandardA8MV2       ContainerServiceVMSizeTypes = "Standard_A8m_v2"
	ContainerServiceVMSizeTypesStandardA8V2        ContainerServiceVMSizeTypes = "Standard_A8_v2"
	ContainerServiceVMSizeTypesStandardA9          ContainerServiceVMSizeTypes = "Standard_A9"
	ContainerServiceVMSizeTypesStandardB2Ms        ContainerServiceVMSizeTypes = "Standard_B2ms"
	ContainerServiceVMSizeTypesStandardB2S         ContainerServiceVMSizeTypes = "Standard_B2s"
	ContainerServiceVMSizeTypesStandardB4Ms        ContainerServiceVMSizeTypes = "Standard_B4ms"
	ContainerServiceVMSizeTypesStandardB8Ms        ContainerServiceVMSizeTypes = "Standard_B8ms"
	ContainerServiceVMSizeTypesStandardD1          ContainerServiceVMSizeTypes = "Standard_D1"
	ContainerServiceVMSizeTypesStandardD11         ContainerServiceVMSizeTypes = "Standard_D11"
	ContainerServiceVMSizeTypesStandardD11V2       ContainerServiceVMSizeTypes = "Standard_D11_v2"
	ContainerServiceVMSizeTypesStandardD11V2Promo  ContainerServiceVMSizeTypes = "Standard_D11_v2_Promo"
	ContainerServiceVMSizeTypesStandardD12         ContainerServiceVMSizeTypes = "Standard_D12"
	ContainerServiceVMSizeTypesStandardD12V2       ContainerServiceVMSizeTypes = "Standard_D12_v2"
	ContainerServiceVMSizeTypesStandardD12V2Promo  ContainerServiceVMSizeTypes = "Standard_D12_v2_Promo"
	ContainerServiceVMSizeTypesStandardD13         ContainerServiceVMSizeTypes = "Standard_D13"
	ContainerServiceVMSizeTypesStandardD13V2       ContainerServiceVMSizeTypes = "Standard_D13_v2"
	ContainerServiceVMSizeTypesStandardD13V2Promo  ContainerServiceVMSizeTypes = "Standard_D13_v2_Promo"
	ContainerServiceVMSizeTypesStandardD14         ContainerServiceVMSizeTypes = "Standard_D14"
	ContainerServiceVMSizeTypesStandardD14V2       ContainerServiceVMSizeTypes = "Standard_D14_v2"
	ContainerServiceVMSizeTypesStandardD14V2Promo  ContainerServiceVMSizeTypes = "Standard_D14_v2_Promo"
	ContainerServiceVMSizeTypesStandardD15V2       ContainerServiceVMSizeTypes = "Standard_D15_v2"
	ContainerServiceVMSizeTypesStandardD16SV3      ContainerServiceVMSizeTypes = "Standard_D16s_v3"
	ContainerServiceVMSizeTypesStandardD16V3       ContainerServiceVMSizeTypes = "Standard_D16_v3"
	ContainerServiceVMSizeTypesStandardD1V2        ContainerServiceVMSizeTypes = "Standard_D1_v2"
	ContainerServiceVMSizeTypesStandardD2          ContainerServiceVMSizeTypes = "Standard_D2"
	ContainerServiceVMSizeTypesStandardD2SV3       ContainerServiceVMSizeTypes = "Standard_D2s_v3"
	ContainerServiceVMSizeTypesStandardD2V2        ContainerServiceVMSizeTypes = "Standard_D2_v2"
	ContainerServiceVMSizeTypesStandardD2V2Promo   ContainerServiceVMSizeTypes = "Standard_D2_v2_Promo"
	ContainerServiceVMSizeTypesStandardD2V3        ContainerServiceVMSizeTypes = "Standard_D2_v3"
	ContainerServiceVMSizeTypesStandardD3          ContainerServiceVMSizeTypes = "Standard_D3"
	ContainerServiceVMSizeTypesStandardD32SV3      ContainerServiceVMSizeTypes = "Standard_D32s_v3"
	ContainerServiceVMSizeTypesStandardD32V3       ContainerServiceVMSizeTypes = "Standard_D32_v3"
	ContainerServiceVMSizeTypesStandardD3V2        ContainerServiceVMSizeTypes = "Standard_D3_v2"
	ContainerServiceVMSizeTypesStandardD3V2Promo   ContainerServiceVMSizeTypes = "Standard_D3_v2_Promo"
	ContainerServiceVMSizeTypesStandardD4          ContainerServiceVMSizeTypes = "Standard_D4"
	ContainerServiceVMSizeTypesStandardD4SV3       ContainerServiceVMSizeTypes = "Standard_D4s_v3"
	ContainerServiceVMSizeTypesStandardD4V2        ContainerServiceVMSizeTypes = "Standard_D4_v2"
	ContainerServiceVMSizeTypesStandardD4V2Promo   ContainerServiceVMSizeTypes = "Standard_D4_v2_Promo"
	ContainerServiceVMSizeTypesStandardD4V3        ContainerServiceVMSizeTypes = "Standard_D4_v3"
	ContainerServiceVMSizeTypesStandardD5V2        ContainerServiceVMSizeTypes = "Standard_D5_v2"
	ContainerServiceVMSizeTypesStandardD5V2Promo   ContainerServiceVMSizeTypes = "Standard_D5_v2_Promo"
	ContainerServiceVMSizeTypesStandardD64SV3      ContainerServiceVMSizeTypes = "Standard_D64s_v3"
	ContainerServiceVMSizeTypesStandardD64V3       ContainerServiceVMSizeTypes = "Standard_D64_v3"
	ContainerServiceVMSizeTypesStandardD8SV3       ContainerServiceVMSizeTypes = "Standard_D8s_v3"
	ContainerServiceVMSizeTypesStandardD8V3        ContainerServiceVMSizeTypes = "Standard_D8_v3"
	ContainerServiceVMSizeTypesStandardDS1         ContainerServiceVMSizeTypes = "Standard_DS1"
	ContainerServiceVMSizeTypesStandardDS11        ContainerServiceVMSizeTypes = "Standard_DS11"
	ContainerServiceVMSizeTypesStandardDS11V2      ContainerServiceVMSizeTypes = "Standard_DS11_v2"
	ContainerServiceVMSizeTypesStandardDS11V2Promo ContainerServiceVMSizeTypes = "Standard_DS11_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS12        ContainerServiceVMSizeTypes = "Standard_DS12"
	ContainerServiceVMSizeTypesStandardDS12V2      ContainerServiceVMSizeTypes = "Standard_DS12_v2"
	ContainerServiceVMSizeTypesStandardDS12V2Promo ContainerServiceVMSizeTypes = "Standard_DS12_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS13        ContainerServiceVMSizeTypes = "Standard_DS13"
	ContainerServiceVMSizeTypesStandardDS132V2     ContainerServiceVMSizeTypes = "Standard_DS13-2_v2"
	ContainerServiceVMSizeTypesStandardDS134V2     ContainerServiceVMSizeTypes = "Standard_DS13-4_v2"
	ContainerServiceVMSizeTypesStandardDS13V2      ContainerServiceVMSizeTypes = "Standard_DS13_v2"
	ContainerServiceVMSizeTypesStandardDS13V2Promo ContainerServiceVMSizeTypes = "Standard_DS13_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS14        ContainerServiceVMSizeTypes = "Standard_DS14"
	ContainerServiceVMSizeTypesStandardDS144V2     ContainerServiceVMSizeTypes = "Standard_DS14-4_v2"
	ContainerServiceVMSizeTypesStandardDS148V2     ContainerServiceVMSizeTypes = "Standard_DS14-8_v2"
	ContainerServiceVMSizeTypesStandardDS14V2      ContainerServiceVMSizeTypes = "Standard_DS14_v2"
	ContainerServiceVMSizeTypesStandardDS14V2Promo ContainerServiceVMSizeTypes = "Standard_DS14_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS15V2      ContainerServiceVMSizeTypes = "Standard_DS15_v2"
	ContainerServiceVMSizeTypesStandardDS1V2       ContainerServiceVMSizeTypes = "Standard_DS1_v2"
	ContainerServiceVMSizeTypesStandardDS2         ContainerServiceVMSizeTypes = "Standard_DS2"
	ContainerServiceVMSizeTypesStandardDS2V2       ContainerServiceVMSizeTypes = "Standard_DS2_v2"
	ContainerServiceVMSizeTypesStandardDS2V2Promo  ContainerServiceVMSizeTypes = "Standard_DS2_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS3         ContainerServiceVMSizeTypes = "Standard_DS3"
	ContainerServiceVMSizeTypesStandardDS3V2       ContainerServiceVMSizeTypes = "Standard_DS3_v2"
	ContainerServiceVMSizeTypesStandardDS3V2Promo  ContainerServiceVMSizeTypes = "Standard_DS3_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS4         ContainerServiceVMSizeTypes = "Standard_DS4"
	ContainerServiceVMSizeTypesStandardDS4V2       ContainerServiceVMSizeTypes = "Standard_DS4_v2"
	ContainerServiceVMSizeTypesStandardDS4V2Promo  ContainerServiceVMSizeTypes = "Standard_DS4_v2_Promo"
	ContainerServiceVMSizeTypesStandardDS5V2       ContainerServiceVMSizeTypes = "Standard_DS5_v2"
	ContainerServiceVMSizeTypesStandardDS5V2Promo  ContainerServiceVMSizeTypes = "Standard_DS5_v2_Promo"
	ContainerServiceVMSizeTypesStandardE16SV3      ContainerServiceVMSizeTypes = "Standard_E16s_v3"
	ContainerServiceVMSizeTypesStandardE16V3       ContainerServiceVMSizeTypes = "Standard_E16_v3"
	ContainerServiceVMSizeTypesStandardE2SV3       ContainerServiceVMSizeTypes = "Standard_E2s_v3"
	ContainerServiceVMSizeTypesStandardE2V3        ContainerServiceVMSizeTypes = "Standard_E2_v3"
	ContainerServiceVMSizeTypesStandardE3216SV3    ContainerServiceVMSizeTypes = "Standard_E32-16s_v3"
	ContainerServiceVMSizeTypesStandardE328SV3     ContainerServiceVMSizeTypes = "Standard_E32-8s_v3"
	ContainerServiceVMSizeTypesStandardE32SV3      ContainerServiceVMSizeTypes = "Standard_E32s_v3"
	ContainerServiceVMSizeTypesStandardE32V3       ContainerServiceVMSizeTypes = "Standard_E32_v3"
	ContainerServiceVMSizeTypesStandardE4SV3       ContainerServiceVMSizeTypes = "Standard_E4s_v3"
	ContainerServiceVMSizeTypesStandardE4V3        ContainerServiceVMSizeTypes = "Standard_E4_v3"
	ContainerServiceVMSizeTypesStandardE6416SV3    ContainerServiceVMSizeTypes = "Standard_E64-16s_v3"
	ContainerServiceVMSizeTypesStandardE6432SV3    ContainerServiceVMSizeTypes = "Standard_E64-32s_v3"
	ContainerServiceVMSizeTypesStandardE64SV3      ContainerServiceVMSizeTypes = "Standard_E64s_v3"
	ContainerServiceVMSizeTypesStandardE64V3       ContainerServiceVMSizeTypes = "Standard_E64_v3"
	ContainerServiceVMSizeTypesStandardE8SV3       ContainerServiceVMSizeTypes = "Standard_E8s_v3"
	ContainerServiceVMSizeTypesStandardE8V3        ContainerServiceVMSizeTypes = "Standard_E8_v3"
	ContainerServiceVMSizeTypesStandardF1          ContainerServiceVMSizeTypes = "Standard_F1"
	ContainerServiceVMSizeTypesStandardF16         ContainerServiceVMSizeTypes = "Standard_F16"
	ContainerServiceVMSizeTypesStandardF16S        ContainerServiceVMSizeTypes = "Standard_F16s"
	ContainerServiceVMSizeTypesStandardF16SV2      ContainerServiceVMSizeTypes = "Standard_F16s_v2"
	ContainerServiceVMSizeTypesStandardF1S         ContainerServiceVMSizeTypes = "Standard_F1s"
	ContainerServiceVMSizeTypesStandardF2          ContainerServiceVMSizeTypes = "Standard_F2"
	ContainerServiceVMSizeTypesStandardF2S         ContainerServiceVMSizeTypes = "Standard_F2s"
	ContainerServiceVMSizeTypesStandardF2SV2       ContainerServiceVMSizeTypes = "Standard_F2s_v2"
	ContainerServiceVMSizeTypesStandardF32SV2      ContainerServiceVMSizeTypes = "Standard_F32s_v2"
	ContainerServiceVMSizeTypesStandardF4          ContainerServiceVMSizeTypes = "Standard_F4"
	ContainerServiceVMSizeTypesStandardF4S         ContainerServiceVMSizeTypes = "Standard_F4s"
	ContainerServiceVMSizeTypesStandardF4SV2       ContainerServiceVMSizeTypes = "Standard_F4s_v2"
	ContainerServiceVMSizeTypesStandardF64SV2      ContainerServiceVMSizeTypes = "Standard_F64s_v2"
	ContainerServiceVMSizeTypesStandardF72SV2      ContainerServiceVMSizeTypes = "Standard_F72s_v2"
	ContainerServiceVMSizeTypesStandardF8          ContainerServiceVMSizeTypes = "Standard_F8"
	ContainerServiceVMSizeTypesStandardF8S         ContainerServiceVMSizeTypes = "Standard_F8s"
	ContainerServiceVMSizeTypesStandardF8SV2       ContainerServiceVMSizeTypes = "Standard_F8s_v2"
	ContainerServiceVMSizeTypesStandardG1          ContainerServiceVMSizeTypes = "Standard_G1"
	ContainerServiceVMSizeTypesStandardG2          ContainerServiceVMSizeTypes = "Standard_G2"
	ContainerServiceVMSizeTypesStandardG3          ContainerServiceVMSizeTypes = "Standard_G3"
	ContainerServiceVMSizeTypesStandardG4          ContainerServiceVMSizeTypes = "Standard_G4"
	ContainerServiceVMSizeTypesStandardG5          ContainerServiceVMSizeTypes = "Standard_G5"
	ContainerServiceVMSizeTypesStandardGS1         ContainerServiceVMSizeTypes = "Standard_GS1"
	ContainerServiceVMSizeTypesStandardGS2         ContainerServiceVMSizeTypes = "Standard_GS2"
	ContainerServiceVMSizeTypesStandardGS3         ContainerServiceVMSizeTypes = "Standard_GS3"
	ContainerServiceVMSizeTypesStandardGS4         ContainerServiceVMSizeTypes = "Standard_GS4"
	ContainerServiceVMSizeTypesStandardGS44        ContainerServiceVMSizeTypes = "Standard_GS4-4"
	ContainerServiceVMSizeTypesStandardGS48        ContainerServiceVMSizeTypes = "Standard_GS4-8"
	ContainerServiceVMSizeTypesStandardGS5         ContainerServiceVMSizeTypes = "Standard_GS5"
	ContainerServiceVMSizeTypesStandardGS516       ContainerServiceVMSizeTypes = "Standard_GS5-16"
	ContainerServiceVMSizeTypesStandardGS58        ContainerServiceVMSizeTypes = "Standard_GS5-8"
	ContainerServiceVMSizeTypesStandardH16         ContainerServiceVMSizeTypes = "Standard_H16"
	ContainerServiceVMSizeTypesStandardH16M        ContainerServiceVMSizeTypes = "Standard_H16m"
	ContainerServiceVMSizeTypesStandardH16Mr       ContainerServiceVMSizeTypes = "Standard_H16mr"
	ContainerServiceVMSizeTypesStandardH16R        ContainerServiceVMSizeTypes = "Standard_H16r"
	ContainerServiceVMSizeTypesStandardH8          ContainerServiceVMSizeTypes = "Standard_H8"
	ContainerServiceVMSizeTypesStandardH8M         ContainerServiceVMSizeTypes = "Standard_H8m"
	ContainerServiceVMSizeTypesStandardL16S        ContainerServiceVMSizeTypes = "Standard_L16s"
	ContainerServiceVMSizeTypesStandardL32S        ContainerServiceVMSizeTypes = "Standard_L32s"
	ContainerServiceVMSizeTypesStandardL4S         ContainerServiceVMSizeTypes = "Standard_L4s"
	ContainerServiceVMSizeTypesStandardL8S         ContainerServiceVMSizeTypes = "Standard_L8s"
	ContainerServiceVMSizeTypesStandardM12832Ms    ContainerServiceVMSizeTypes = "Standard_M128-32ms"
	ContainerServiceVMSizeTypesStandardM12864Ms    ContainerServiceVMSizeTypes = "Standard_M128-64ms"
	ContainerServiceVMSizeTypesStandardM128Ms      ContainerServiceVMSizeTypes = "Standard_M128ms"
	ContainerServiceVMSizeTypesStandardM128S       ContainerServiceVMSizeTypes = "Standard_M128s"
	ContainerServiceVMSizeTypesStandardM6416Ms     ContainerServiceVMSizeTypes = "Standard_M64-16ms"
	ContainerServiceVMSizeTypesStandardM6432Ms     ContainerServiceVMSizeTypes = "Standard_M64-32ms"
	ContainerServiceVMSizeTypesStandardM64Ms       ContainerServiceVMSizeTypes = "Standard_M64ms"
	ContainerServiceVMSizeTypesStandardM64S        ContainerServiceVMSizeTypes = "Standard_M64s"
	ContainerServiceVMSizeTypesStandardNC12        ContainerServiceVMSizeTypes = "Standard_NC12"
	ContainerServiceVMSizeTypesStandardNC12SV2     ContainerServiceVMSizeTypes = "Standard_NC12s_v2"
	ContainerServiceVMSizeTypesStandardNC12SV3     ContainerServiceVMSizeTypes = "Standard_NC12s_v3"
	ContainerServiceVMSizeTypesStandardNC24        ContainerServiceVMSizeTypes = "Standard_NC24"
	ContainerServiceVMSizeTypesStandardNC24R       ContainerServiceVMSizeTypes = "Standard_NC24r"
	ContainerServiceVMSizeTypesStandardNC24RsV2    ContainerServiceVMSizeTypes = "Standard_NC24rs_v2"
	ContainerServiceVMSizeTypesStandardNC24RsV3    ContainerServiceVMSizeTypes = "Standard_NC24rs_v3"
	ContainerServiceVMSizeTypesStandardNC24SV2     ContainerServiceVMSizeTypes = "Standard_NC24s_v2"
	ContainerServiceVMSizeTypesStandardNC24SV3     ContainerServiceVMSizeTypes = "Standard_NC24s_v3"
	ContainerServiceVMSizeTypesStandardNC6         ContainerServiceVMSizeTypes = "Standard_NC6"
	ContainerServiceVMSizeTypesStandardNC6SV2      ContainerServiceVMSizeTypes = "Standard_NC6s_v2"
	ContainerServiceVMSizeTypesStandardNC6SV3      ContainerServiceVMSizeTypes = "Standard_NC6s_v3"
	ContainerServiceVMSizeTypesStandardND12S       ContainerServiceVMSizeTypes = "Standard_ND12s"
	ContainerServiceVMSizeTypesStandardND24Rs      ContainerServiceVMSizeTypes = "Standard_ND24rs"
	ContainerServiceVMSizeTypesStandardND24S       ContainerServiceVMSizeTypes = "Standard_ND24s"
	ContainerServiceVMSizeTypesStandardND6S        ContainerServiceVMSizeTypes = "Standard_ND6s"
	ContainerServiceVMSizeTypesStandardNV12        ContainerServiceVMSizeTypes = "Standard_NV12"
	ContainerServiceVMSizeTypesStandardNV24        ContainerServiceVMSizeTypes = "Standard_NV24"
	ContainerServiceVMSizeTypesStandardNV6         ContainerServiceVMSizeTypes = "Standard_NV6"
)

// PossibleContainerServiceVMSizeTypesValues returns the possible values for the ContainerServiceVMSizeTypes const type.
func PossibleContainerServiceVMSizeTypesValues() []ContainerServiceVMSizeTypes {
	return []ContainerServiceVMSizeTypes{
		ContainerServiceVMSizeTypesStandardA1,
		ContainerServiceVMSizeTypesStandardA10,
		ContainerServiceVMSizeTypesStandardA11,
		ContainerServiceVMSizeTypesStandardA1V2,
		ContainerServiceVMSizeTypesStandardA2,
		ContainerServiceVMSizeTypesStandardA2MV2,
		ContainerServiceVMSizeTypesStandardA2V2,
		ContainerServiceVMSizeTypesStandardA3,
		ContainerServiceVMSizeTypesStandardA4,
		ContainerServiceVMSizeTypesStandardA4MV2,
		ContainerServiceVMSizeTypesStandardA4V2,
		ContainerServiceVMSizeTypesStandardA5,
		ContainerServiceVMSizeTypesStandardA6,
		ContainerServiceVMSizeTypesStandardA7,
		ContainerServiceVMSizeTypesStandardA8,
		ContainerServiceVMSizeTypesStandardA8MV2,
		ContainerServiceVMSizeTypesStandardA8V2,
		ContainerServiceVMSizeTypesStandardA9,
		ContainerServiceVMSizeTypesStandardB2Ms,
		ContainerServiceVMSizeTypesStandardB2S,
		ContainerServiceVMSizeTypesStandardB4Ms,
		ContainerServiceVMSizeTypesStandardB8Ms,
		ContainerServiceVMSizeTypesStandardD1,
		ContainerServiceVMSizeTypesStandardD11,
		ContainerServiceVMSizeTypesStandardD11V2,
		ContainerServiceVMSizeTypesStandardD11V2Promo,
		ContainerServiceVMSizeTypesStandardD12,
		ContainerServiceVMSizeTypesStandardD12V2,
		ContainerServiceVMSizeTypesStandardD12V2Promo,
		ContainerServiceVMSizeTypesStandardD13,
		ContainerServiceVMSizeTypesStandardD13V2,
		ContainerServiceVMSizeTypesStandardD13V2Promo,
		ContainerServiceVMSizeTypesStandardD14,
		ContainerServiceVMSizeTypesStandardD14V2,
		ContainerServiceVMSizeTypesStandardD14V2Promo,
		ContainerServiceVMSizeTypesStandardD15V2,
		ContainerServiceVMSizeTypesStandardD16SV3,
		ContainerServiceVMSizeTypesStandardD16V3,
		ContainerServiceVMSizeTypesStandardD1V2,
		ContainerServiceVMSizeTypesStandardD2,
		ContainerServiceVMSizeTypesStandardD2SV3,
		ContainerServiceVMSizeTypesStandardD2V2,
		ContainerServiceVMSizeTypesStandardD2V2Promo,
		ContainerServiceVMSizeTypesStandardD2V3,
		ContainerServiceVMSizeTypesStandardD3,
		ContainerServiceVMSizeTypesStandardD32SV3,
		ContainerServiceVMSizeTypesStandardD32V3,
		ContainerServiceVMSizeTypesStandardD3V2,
		ContainerServiceVMSizeTypesStandardD3V2Promo,
		ContainerServiceVMSizeTypesStandardD4,
		ContainerServiceVMSizeTypesStandardD4SV3,
		ContainerServiceVMSizeTypesStandardD4V2,
		ContainerServiceVMSizeTypesStandardD4V2Promo,
		ContainerServiceVMSizeTypesStandardD4V3,
		ContainerServiceVMSizeTypesStandardD5V2,
		ContainerServiceVMSizeTypesStandardD5V2Promo,
		ContainerServiceVMSizeTypesStandardD64SV3,
		ContainerServiceVMSizeTypesStandardD64V3,
		ContainerServiceVMSizeTypesStandardD8SV3,
		ContainerServiceVMSizeTypesStandardD8V3,
		ContainerServiceVMSizeTypesStandardDS1,
		ContainerServiceVMSizeTypesStandardDS11,
		ContainerServiceVMSizeTypesStandardDS11V2,
		ContainerServiceVMSizeTypesStandardDS11V2Promo,
		ContainerServiceVMSizeTypesStandardDS12,
		ContainerServiceVMSizeTypesStandardDS12V2,
		ContainerServiceVMSizeTypesStandardDS12V2Promo,
		ContainerServiceVMSizeTypesStandardDS13,
		ContainerServiceVMSizeTypesStandardDS132V2,
		ContainerServiceVMSizeTypesStandardDS134V2,
		ContainerServiceVMSizeTypesStandardDS13V2,
		ContainerServiceVMSizeTypesStandardDS13V2Promo,
		ContainerServiceVMSizeTypesStandardDS14,
		ContainerServiceVMSizeTypesStandardDS144V2,
		ContainerServiceVMSizeTypesStandardDS148V2,
		ContainerServiceVMSizeTypesStandardDS14V2,
		ContainerServiceVMSizeTypesStandardDS14V2Promo,
		ContainerServiceVMSizeTypesStandardDS15V2,
		ContainerServiceVMSizeTypesStandardDS1V2,
		ContainerServiceVMSizeTypesStandardDS2,
		ContainerServiceVMSizeTypesStandardDS2V2,
		ContainerServiceVMSizeTypesStandardDS2V2Promo,
		ContainerServiceVMSizeTypesStandardDS3,
		ContainerServiceVMSizeTypesStandardDS3V2,
		ContainerServiceVMSizeTypesStandardDS3V2Promo,
		ContainerServiceVMSizeTypesStandardDS4,
		ContainerServiceVMSizeTypesStandardDS4V2,
		ContainerServiceVMSizeTypesStandardDS4V2Promo,
		ContainerServiceVMSizeTypesStandardDS5V2,
		ContainerServiceVMSizeTypesStandardDS5V2Promo,
		ContainerServiceVMSizeTypesStandardE16SV3,
		ContainerServiceVMSizeTypesStandardE16V3,
		ContainerServiceVMSizeTypesStandardE2SV3,
		ContainerServiceVMSizeTypesStandardE2V3,
		ContainerServiceVMSizeTypesStandardE3216SV3,
		ContainerServiceVMSizeTypesStandardE328SV3,
		ContainerServiceVMSizeTypesStandardE32SV3,
		ContainerServiceVMSizeTypesStandardE32V3,
		ContainerServiceVMSizeTypesStandardE4SV3,
		ContainerServiceVMSizeTypesStandardE4V3,
		ContainerServiceVMSizeTypesStandardE6416SV3,
		ContainerServiceVMSizeTypesStandardE6432SV3,
		ContainerServiceVMSizeTypesStandardE64SV3,
		ContainerServiceVMSizeTypesStandardE64V3,
		ContainerServiceVMSizeTypesStandardE8SV3,
		ContainerServiceVMSizeTypesStandardE8V3,
		ContainerServiceVMSizeTypesStandardF1,
		ContainerServiceVMSizeTypesStandardF16,
		ContainerServiceVMSizeTypesStandardF16S,
		ContainerServiceVMSizeTypesStandardF16SV2,
		ContainerServiceVMSizeTypesStandardF1S,
		ContainerServiceVMSizeTypesStandardF2,
		ContainerServiceVMSizeTypesStandardF2S,
		ContainerServiceVMSizeTypesStandardF2SV2,
		ContainerServiceVMSizeTypesStandardF32SV2,
		ContainerServiceVMSizeTypesStandardF4,
		ContainerServiceVMSizeTypesStandardF4S,
		ContainerServiceVMSizeTypesStandardF4SV2,
		ContainerServiceVMSizeTypesStandardF64SV2,
		ContainerServiceVMSizeTypesStandardF72SV2,
		ContainerServiceVMSizeTypesStandardF8,
		ContainerServiceVMSizeTypesStandardF8S,
		ContainerServiceVMSizeTypesStandardF8SV2,
		ContainerServiceVMSizeTypesStandardG1,
		ContainerServiceVMSizeTypesStandardG2,
		ContainerServiceVMSizeTypesStandardG3,
		ContainerServiceVMSizeTypesStandardG4,
		ContainerServiceVMSizeTypesStandardG5,
		ContainerServiceVMSizeTypesStandardGS1,
		ContainerServiceVMSizeTypesStandardGS2,
		ContainerServiceVMSizeTypesStandardGS3,
		ContainerServiceVMSizeTypesStandardGS4,
		ContainerServiceVMSizeTypesStandardGS44,
		ContainerServiceVMSizeTypesStandardGS48,
		ContainerServiceVMSizeTypesStandardGS5,
		ContainerServiceVMSizeTypesStandardGS516,
		ContainerServiceVMSizeTypesStandardGS58,
		ContainerServiceVMSizeTypesStandardH16,
		ContainerServiceVMSizeTypesStandardH16M,
		ContainerServiceVMSizeTypesStandardH16Mr,
		ContainerServiceVMSizeTypesStandardH16R,
		ContainerServiceVMSizeTypesStandardH8,
		ContainerServiceVMSizeTypesStandardH8M,
		ContainerServiceVMSizeTypesStandardL16S,
		ContainerServiceVMSizeTypesStandardL32S,
		ContainerServiceVMSizeTypesStandardL4S,
		ContainerServiceVMSizeTypesStandardL8S,
		ContainerServiceVMSizeTypesStandardM12832Ms,
		ContainerServiceVMSizeTypesStandardM12864Ms,
		ContainerServiceVMSizeTypesStandardM128Ms,
		ContainerServiceVMSizeTypesStandardM128S,
		ContainerServiceVMSizeTypesStandardM6416Ms,
		ContainerServiceVMSizeTypesStandardM6432Ms,
		ContainerServiceVMSizeTypesStandardM64Ms,
		ContainerServiceVMSizeTypesStandardM64S,
		ContainerServiceVMSizeTypesStandardNC12,
		ContainerServiceVMSizeTypesStandardNC12SV2,
		ContainerServiceVMSizeTypesStandardNC12SV3,
		ContainerServiceVMSizeTypesStandardNC24,
		ContainerServiceVMSizeTypesStandardNC24R,
		ContainerServiceVMSizeTypesStandardNC24RsV2,
		ContainerServiceVMSizeTypesStandardNC24RsV3,
		ContainerServiceVMSizeTypesStandardNC24SV2,
		ContainerServiceVMSizeTypesStandardNC24SV3,
		ContainerServiceVMSizeTypesStandardNC6,
		ContainerServiceVMSizeTypesStandardNC6SV2,
		ContainerServiceVMSizeTypesStandardNC6SV3,
		ContainerServiceVMSizeTypesStandardND12S,
		ContainerServiceVMSizeTypesStandardND24Rs,
		ContainerServiceVMSizeTypesStandardND24S,
		ContainerServiceVMSizeTypesStandardND6S,
		ContainerServiceVMSizeTypesStandardNV12,
		ContainerServiceVMSizeTypesStandardNV24,
		ContainerServiceVMSizeTypesStandardNV6,
	}
}

// Count - Number of masters (VMs) in the container service cluster. Allowed values are 1, 3, and 5. The default value is
// 1.
type Count int32

const (
	CountOne   Count = 1
	CountThree Count = 3
	CountFive  Count = 5
)

// PossibleCountValues returns the possible values for the Count const type.
func PossibleCountValues() []Count {
	return []Count{
		CountOne,
		CountThree,
		CountFive,
	}
}

type Expander string

const (
	ExpanderLeastWaste Expander = "least-waste"
	ExpanderMostPods   Expander = "most-pods"
	ExpanderRandom     Expander = "random"
)

// PossibleExpanderValues returns the possible values for the Expander const type.
func PossibleExpanderValues() []Expander {
	return []Expander{
		ExpanderLeastWaste,
		ExpanderMostPods,
		ExpanderRandom,
	}
}

// LicenseType - The licenseType to use for Windows VMs. Windows_Server is used to enable Azure Hybrid User Benefits for Windows
// VMs.
type LicenseType string

const (
	LicenseTypeNone          LicenseType = "None"
	LicenseTypeWindowsServer LicenseType = "Windows_Server"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeNone,
		LicenseTypeWindowsServer,
	}
}

// LoadBalancerSKU - The load balancer sku for the managed cluster.
type LoadBalancerSKU string

const (
	LoadBalancerSKUBasic    LoadBalancerSKU = "basic"
	LoadBalancerSKUStandard LoadBalancerSKU = "standard"
)

// PossibleLoadBalancerSKUValues returns the possible values for the LoadBalancerSKU const type.
func PossibleLoadBalancerSKUValues() []LoadBalancerSKU {
	return []LoadBalancerSKU{
		LoadBalancerSKUBasic,
		LoadBalancerSKUStandard,
	}
}

// ManagedClusterPodIdentityProvisioningState - The current provisioning state of the pod identity.
type ManagedClusterPodIdentityProvisioningState string

const (
	ManagedClusterPodIdentityProvisioningStateAssigned ManagedClusterPodIdentityProvisioningState = "Assigned"
	ManagedClusterPodIdentityProvisioningStateDeleting ManagedClusterPodIdentityProvisioningState = "Deleting"
	ManagedClusterPodIdentityProvisioningStateFailed   ManagedClusterPodIdentityProvisioningState = "Failed"
	ManagedClusterPodIdentityProvisioningStateUpdating ManagedClusterPodIdentityProvisioningState = "Updating"
)

// PossibleManagedClusterPodIdentityProvisioningStateValues returns the possible values for the ManagedClusterPodIdentityProvisioningState const type.
func PossibleManagedClusterPodIdentityProvisioningStateValues() []ManagedClusterPodIdentityProvisioningState {
	return []ManagedClusterPodIdentityProvisioningState{
		ManagedClusterPodIdentityProvisioningStateAssigned,
		ManagedClusterPodIdentityProvisioningStateDeleting,
		ManagedClusterPodIdentityProvisioningStateFailed,
		ManagedClusterPodIdentityProvisioningStateUpdating,
	}
}

// ManagedClusterSKUName - Name of a managed cluster SKU.
type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = "Basic"
)

// PossibleManagedClusterSKUNameValues returns the possible values for the ManagedClusterSKUName const type.
func PossibleManagedClusterSKUNameValues() []ManagedClusterSKUName {
	return []ManagedClusterSKUName{
		ManagedClusterSKUNameBasic,
	}
}

// ManagedClusterSKUTier - Tier of a managed cluster SKU.
type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree ManagedClusterSKUTier = "Free"
	ManagedClusterSKUTierPaid ManagedClusterSKUTier = "Paid"
)

// PossibleManagedClusterSKUTierValues returns the possible values for the ManagedClusterSKUTier const type.
func PossibleManagedClusterSKUTierValues() []ManagedClusterSKUTier {
	return []ManagedClusterSKUTier{
		ManagedClusterSKUTierFree,
		ManagedClusterSKUTierPaid,
	}
}

// NetworkMode - Network mode used for building Kubernetes network.
type NetworkMode string

const (
	NetworkModeBridge      NetworkMode = "bridge"
	NetworkModeTransparent NetworkMode = "transparent"
)

// PossibleNetworkModeValues returns the possible values for the NetworkMode const type.
func PossibleNetworkModeValues() []NetworkMode {
	return []NetworkMode{
		NetworkModeBridge,
		NetworkModeTransparent,
	}
}

// NetworkPlugin - Network plugin used for building Kubernetes network.
type NetworkPlugin string

const (
	NetworkPluginAzure   NetworkPlugin = "azure"
	NetworkPluginKubenet NetworkPlugin = "kubenet"
)

// PossibleNetworkPluginValues returns the possible values for the NetworkPlugin const type.
func PossibleNetworkPluginValues() []NetworkPlugin {
	return []NetworkPlugin{
		NetworkPluginAzure,
		NetworkPluginKubenet,
	}
}

// NetworkPolicy - Network policy used for building Kubernetes network.
type NetworkPolicy string

const (
	NetworkPolicyAzure  NetworkPolicy = "azure"
	NetworkPolicyCalico NetworkPolicy = "calico"
)

// PossibleNetworkPolicyValues returns the possible values for the NetworkPolicy const type.
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return []NetworkPolicy{
		NetworkPolicyAzure,
		NetworkPolicyCalico,
	}
}

// OSDiskType - OSDiskType represents the type of an OS disk on an agent pool.
type OSDiskType string

const (
	OSDiskTypeEphemeral OSDiskType = "Ephemeral"
	OSDiskTypeManaged   OSDiskType = "Managed"
)

// PossibleOSDiskTypeValues returns the possible values for the OSDiskType const type.
func PossibleOSDiskTypeValues() []OSDiskType {
	return []OSDiskType{
		OSDiskTypeEphemeral,
		OSDiskTypeManaged,
	}
}

// OSType - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

// PossibleOSTypeValues returns the possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{
		OSTypeLinux,
		OSTypeWindows,
	}
}

// OutboundType - The outbound (egress) routing method.
type OutboundType string

const (
	OutboundTypeLoadBalancer       OutboundType = "loadBalancer"
	OutboundTypeUserDefinedRouting OutboundType = "userDefinedRouting"
)

// PossibleOutboundTypeValues returns the possible values for the OutboundType const type.
func PossibleOutboundTypeValues() []OutboundType {
	return []OutboundType{
		OutboundTypeLoadBalancer,
		OutboundTypeUserDefinedRouting,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// ResourceIdentityType - The type of identity used for the managed cluster. Type 'SystemAssigned' will use an implicitly
// created identity in master components and an auto-created user assigned identity in MC_ resource group
// in agent nodes. Type 'None' will not use MSI for the managed cluster, service principal will be used instead.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// ScaleSetEvictionPolicy - ScaleSetEvictionPolicy to be used to specify eviction policy for Spot virtual machine scale set.
// Default to Delete.
type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = "Deallocate"
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = "Delete"
)

// PossibleScaleSetEvictionPolicyValues returns the possible values for the ScaleSetEvictionPolicy const type.
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return []ScaleSetEvictionPolicy{
		ScaleSetEvictionPolicyDeallocate,
		ScaleSetEvictionPolicyDelete,
	}
}

// ScaleSetPriority - ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
type ScaleSetPriority string

const (
	ScaleSetPriorityRegular ScaleSetPriority = "Regular"
	ScaleSetPrioritySpot    ScaleSetPriority = "Spot"
)

// PossibleScaleSetPriorityValues returns the possible values for the ScaleSetPriority const type.
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return []ScaleSetPriority{
		ScaleSetPriorityRegular,
		ScaleSetPrioritySpot,
	}
}

// UpgradeChannel - upgrade channel for auto upgrade.
type UpgradeChannel string

const (
	UpgradeChannelNone   UpgradeChannel = "none"
	UpgradeChannelPatch  UpgradeChannel = "patch"
	UpgradeChannelRapid  UpgradeChannel = "rapid"
	UpgradeChannelStable UpgradeChannel = "stable"
)

// PossibleUpgradeChannelValues returns the possible values for the UpgradeChannel const type.
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return []UpgradeChannel{
		UpgradeChannelNone,
		UpgradeChannelPatch,
		UpgradeChannelRapid,
		UpgradeChannelStable,
	}
}
