## Java

These settings apply only when `--java` is specified on the command line.

``` yaml $(java)
title: DataReplicationMgmtClient
service-name: Recovery Services Data Replication

graal-vm-config: false
remove-operation-group: ReplicationExtensionOperationStatus
rename-model: ProtectedItemModelPropertiesLastFailedPlannedFailoverJob:LastFailedPlannedFailoverJob,ProtectedItemModelPropertiesLastFailedEnableProtectionJob:LastFailedEnableProtectionJob,HyperVToAzStackHciReplicationExtensionModelCustomProperties:HyperVToAzStackHciRepExtnCustomProps,VMwareToAzStackHciPlannedFailoverModelCustomProperties:VMwareToAzStackHciPlannedFailoverCustomProps,VMwareToAzStackHciReplicationExtensionModelCustomProperties:VMwareToAzStackHciRepExtnCustomProps,HyperVToAzStackHciRecoveryPointModelCustomProperties:HyperVToAzStackHciRecoveryPointCustomProps,HyperVToAzStackHciPlannedFailoverModelCustomProperties:HyperVToAzStackHciPlannedFailoverCustomProps,VMwareToAzStackHciProtectedItemModelCustomProperties:VMwareToAzStackHciProtectedItemCustomProps,HyperVToAzStackHciProtectedItemModelCustomProperties:HyperVToAzStackHciProtectedItemCustomProps
```
