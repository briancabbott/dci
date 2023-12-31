# Compute Admin

> see https://aka.ms/autorest

This is the AutoRest configuration file for Compute Admin.

---
## Getting Started
To build the SDK for Compute Admin, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`
---

## Configuration


### Basic Information
These are the global settings for the Compute API.

``` yaml
title: ComputeAdminClient
description: Compute Admin Client
openapi-type: arm
tag: package-2021-09-01

directive:
  - where:
      - $.definitions.ScaleUnit
    suppress:
      - NestedResourcesMustHaveListOperation
    reason: 'CRP cannot support the list API for scale units due to the undesired load that would inflict on the system'
```

### Tag: package-2021-09-01

These settings apply only when `--tag=package-2021-09-01` is specified on the command line.

``` yaml $(tag) == 'package-2021-09-01'
input-file:
   - Microsoft.Compute.Admin/preview/2015-12-01-preview/Compute.json
   - Microsoft.Compute.Admin/preview/2015-12-01-preview/PlatformImages.json
   - Microsoft.Compute.Admin/preview/2015-12-01-preview/VMExtensions.json
   - Microsoft.Compute.Admin/stable/2020-11-01/Features.json
   - Microsoft.Compute.Admin/stable/2021-01-01/Quotas.json
   - Microsoft.Compute.Admin/stable/2021-03-30/ScaleUnits.json
   - Microsoft.Compute.Admin/stable/2021-09-01/Disks.json
   - Microsoft.Compute.Admin/stable/2021-09-01/DiskMigrationJobs.json
```

---
# Code Generation

## C#

``` yaml $(csharp)
csharp:
  azure-arm: true
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: Microsoft.AzureStack.Management.Compute.Admin
  payload-flattening-threshold: 2
  output-folder: $(csharp-sdks-folder)/Generated
  clear-output-folder: true
```

# AutoRest v3 generators

> see https://aka.ms/autorest

``` yaml
input-file:  
     - Microsoft.Compute.Admin/preview/2015-12-01-preview/Compute.json
     - Microsoft.Compute.Admin/preview/2015-12-01-preview/PlatformImages.json
     - Microsoft.Compute.Admin/preview/2015-12-01-preview/VMExtensions.json
     - Microsoft.Compute.Admin/stable/2020-11-01/Features.json
     - Microsoft.Compute.Admin/stable/2021-01-01/Quotas.json
     - Microsoft.Compute.Admin/stable/2021-03-30/ScaleUnits.json
     - Microsoft.Compute.Admin/stable/2021-09-01/Disks.json
     - Microsoft.Compute.Admin/stable/2021-09-01/DiskMigrationJobs.json
```

## Multi-API/Profile support for AutoRest v3 generators 

AutoRest V3 generators require the use of `--tag=all-api-versions` to select api files.

This block is updated by an automatic script. Edits may be lost!

``` yaml $(tag) == 'all-api-versions' /* autogenerated */
# include the azure profile definitions from the standard location
require: $(this-folder)/../../../../profiles/readme.md

# all the input files across all versions
input-file:
   - $(this-folder)/Microsoft.Compute.Admin/preview/2015-12-01-preview/Compute.json
   - $(this-folder)/Microsoft.Compute.Admin/preview/2015-12-01-preview/PlatformImages.json
   - $(this-folder)/Microsoft.Compute.Admin/preview/2015-12-01-preview/VMExtensions.json
   - $(this-folder)/Microsoft.Compute.Admin/preview/2018-02-09/Quotas.json
   - $(this-folder)/Microsoft.Compute.Admin/preview/2018-07-30-preview/Disks.json
   - $(this-folder)/Microsoft.Compute.Admin/preview/2018-07-30-preview/DiskMigrationJobs.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2020-11-01/Features.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-01-01/Quotas.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-03-30/ScaleUnits.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-04-01/Disks.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-04-01/DiskMigrationJobs.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-09-01/Disks.json
   - $(this-folder)/Microsoft.Compute.Admin/stable/2021-09-01/DiskMigrationJobs.json
```

If there are files that should not be in the `all-api-versions` set, 
uncomment the  `exclude-file` section below and add the file paths.

``` yaml $(tag) == 'all-api-versions'
#exclude-file: 
#  - $(this-folder)/Microsoft.Example/stable/2010-01-01/somefile.json
```

