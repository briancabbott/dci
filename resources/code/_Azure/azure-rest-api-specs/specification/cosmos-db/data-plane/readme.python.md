## Python

These settings apply only when `--python` is specified on the command line.
Please also specify `--python-sdks-folder=<path to the root directory of your azure-sdk-for-python clone>`.

``` yaml $(python)
python-mode: create
python:
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: azure.table
  package-name: azure-table
  clear-output-folder: true
  no-namespace-folders: true
  enable-xml: true
  vanilla: true
  output-folder: "$(python-sdks-folder)/sdk/tables/azure-data-tables/azure/data/tables/_generated"
  package-version: "1.0.0b1"
  models-mode: msrest
  azure-validator: false
```

### Use strings for dates when python doesn't have enough precision
``` yaml
directive:
- from: swagger-document
  where: $.definitions.AccessPolicy.properties
  transform: >
    $.Start.format = "str";
    $.Expiry.format = "str";
```

### SignedIdentifier shouldn't require an AccessPolicy, only ID
``` yaml
directive:
- from: swagger-document
  where: $.definitions.SignedIdentifier
  transform: >
    $.required = [ "Id" ];
```

