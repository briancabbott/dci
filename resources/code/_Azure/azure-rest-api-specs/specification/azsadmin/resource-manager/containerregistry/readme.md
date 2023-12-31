# ContainerRegistry Admin

> see https://aka.ms/autorest

This is the AutoRest configuration file for ContainerRegistry Admin.

---
## Getting Started
To build the SDK for ContainerRegistry Admin, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`
---

## Configuration

### Basic Information
These are the global settings for the ContainerRegistry API.

``` yaml
title: ContainerRegistryAdminClient
description: ContainerRegistry Admin Client
openapi-type: arm
tag: package-2019-11-01-preview
```

### Tag: package-2019-11-01-preview

These settings apply only when `--tag=package-2019-11-01-preview` is specified on the command line.

``` yaml $(tag) == 'package-2019-11-01-preview'
input-file:
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/capacity.json"
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/configuration.json"
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/operations.json"
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/quotas.json"
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/registries.json"
    - "Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/setup.json"
```

---
# Code Generation

## C#

These settings apply only when `--csharp` is specified on the command line.
Please also specify `--csharp-sdks-folder=<path to "SDKs" directory of your azure-sdk-for-net clone>`.

``` yaml $(csharp)
csharp:
  azure-arm: true
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: Microsoft.AzureStack.Management.ContainerRegistry.Admin
  payload-flattening-threshold: 1
  output-folder: $(csharp-sdks-folder)/Generated
  clear-output-folder: true
```

## Multi-API/Profile support for AutoRest v3 generators 

AutoRest V3 generators require the use of `--tag=all-api-versions` to select api files.

This block is updated by an automatic script. Edits may be lost!

``` yaml $(tag) == 'all-api-versions' /* autogenerated */
# include the azure profile definitions from the standard location
require: $(this-folder)/../../../../profiles/readme.md

# all the input files across all versions
input-file:
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/capacity.json
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/configuration.json
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/operations.json
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/quotas.json
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/registries.json
  - $(this-folder)/Microsoft.ContainerRegistry.Admin/preview/2019-11-01-preview/setup.json
```

If there are files that should not be in the `all-api-versions` set, 
uncomment the  `exclude-file` section below and add the file paths.

``` yaml $(tag) == 'all-api-versions'
#exclude-file: 
#  - $(this-folder)/Microsoft.Example/stable/2010-01-01/somefile.json
```
