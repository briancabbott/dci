## Java

These settings apply only when `--java` is specified on the command line.
Please also specify `--azure-libraries-for-java-folder=<path to the root directory of your azure-libraries-for-java clone>`.

``` yaml $(java)
azure-arm: true
fluent: true
namespace: com.microsoft.azure.management.recoveryservicesbackup
license-header: MICROSOFT_MIT_NO_CODEGEN
payload-flattening-threshold: 1
output-folder: $(azure-libraries-for-java-folder)/azure-mgmt-recoveryservicesbackup
directive:
  - from: bms.json
    where: "$.definitions.AzureVmWorkloadSAPHanaHSRProtectableItem"
    transform: >
      $["x-ms-client-name"] = "AzureVmWorkloadSapHanaHsr"
  - from: bms.json
    where: "$.definitions.ProtectedItem.properties.softDeleteRetentionPeriodInDays"
    transform: >
      $["x-ms-client-name"] = "softDeleteRetentionPeriod"
  - from: bms.json
    where: "$.definitions.PrivateLinkServiceConnectionState.properties.actionsRequired"
    transform: >
      $["x-ms-client-name"] = "actionRequired"
```

### Java multi-api

``` yaml $(java) && $(multiapi)
batch:
  - tag: package-2021-01
  - tag: package-2020-12
  - tag: package-2016-08
  - tag: package-2020-02
directive:
  - rename-operation:
      from: Operations_List
      to: Operation_List
```
### Tag: package-2021-01 and java

These settings apply only when `--tag=package-2021-01 --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2021-01' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.recoveryservices.backup.v2021_01_01
  output-folder: $(azure-libraries-for-java-folder)/sdk/recoveryservices.backup/mgmt-v2021_01_01
regenerate-manager: true
generate-interface: true
```

### Tag: package-2020-12 and java

These settings apply only when `--tag=package-2020-12 --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2020-12' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.recoveryservices.backup.v2020_12_01
  output-folder: $(azure-libraries-for-java-folder)/sdk/recoveryservices.backup/mgmt-v2020_12_01
regenerate-manager: true
generate-interface: true
```

### Tag: package-2020-02 and java

These settings apply only when `--tag=package-2020-02 --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2020-02' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.recoveryservices.backup.v2020_02_02
  output-folder: $(azure-libraries-for-java-folder)/sdk/recoveryservices.backup/mgmt-v2020_02_02
regenerate-manager: true
generate-interface: true
```

### Tag: package-2016-08 and java

These settings apply only when `--tag=package-2016-08 --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2016-08' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.recoveryservices.backup.v2016_08_10
  output-folder: $(azure-libraries-for-java-folder)/sdk/recoveryservices.backup/mgmt-v2016_08_10
regenerate-manager: true
generate-interface: true
```

### Tag: package-2016-06 and java

These settings apply only when `--tag=package-2016-06 --java` is specified on the command line.
Please also specify `--azure-libraries-for-java=<path to the root directory of your azure-sdk-for-java clone>`.

``` yaml $(tag) == 'package-2016-06' && $(java) && $(multiapi)
java:
  namespace: com.microsoft.azure.management.recoveryservices.backup.v2016_06_01
  output-folder: $(azure-libraries-for-java-folder)/sdk/recoveryservices.backup/mgmt-v2016_06_01
regenerate-manager: true
generate-interface: true
```
