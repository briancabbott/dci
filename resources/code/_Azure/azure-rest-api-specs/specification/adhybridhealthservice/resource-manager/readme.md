# ADHybridHealthService

> see https://aka.ms/autorest

This is the AutoRest configuration file for ADHybridHealthService.



---
## Getting Started
To build the SDK for ADHybridHealthService, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`
---

## Configuration



### Basic Information
These are the global settings for the ADHybridHealthService API.

``` yaml
openapi-type: arm
tag: package-2014-01
```


### Tag: package-2014-01

These settings apply only when `--tag=package-2014-01` is specified on the command line.

``` yaml $(tag) == 'package-2014-01'
input-file:
- Microsoft.ADHybridHealthService/stable/2014-01-01/ADHybridHealthService.json
```

---
# Code Generation


## Swagger to SDK

This section describes what SDK should be generated by the automatic system.
This is not used by Autorest itself.

``` yaml $(swagger-to-sdk)
swagger-to-sdk:
  - repo: azure-sdk-for-python
  - repo: azure-sdk-for-java
  - repo: azure-sdk-for-go
  - repo: azure-sdk-for-node
  - repo: azure-resource-manager-schemas
  - repo: azure-powershell
```


## Python

These settings apply only when `--python` is specified on the command line.
Please also specify `--python-sdks-folder=<path to the root directory of your azure-sdk-for-python clone>`.
Use `--python-mode=update` if you already have a setup.py and just want to update the code itself.

``` yaml $(python)
python-mode: create
python:
  azure-arm: true
  license-header: MICROSOFT_MIT_NO_VERSION
  payload-flattening-threshold: 2
  namespace: azure.mgmt.adhybridhealthservice
  package-name: azure-mgmt-adhybridhealthservice
  package-version: 1.0.1
  clear-output-folder: true
```
``` yaml $(python) && $(python-mode) == 'update'
python:
  no-namespace-folders: true
  output-folder: $(python-sdks-folder)/azure-mgmt-adhybridhealthservice/azure/mgmt/adhybridhealthservice
```
``` yaml $(python) && $(python-mode) == 'create'
python:
  basic-setup-py: true
  output-folder: $(python-sdks-folder)/azure-mgmt-adhybridhealthservice
```


## Go

See configuration in [readme.go.md](./readme.go.md)

## Java

See configuration in [readme.java.md](./readme.java.md)



