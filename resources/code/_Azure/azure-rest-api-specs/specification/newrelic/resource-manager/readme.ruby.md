## Ruby

These settings apply only when `--ruby` is specified on the command line.

```yaml
package-name: azure_mgmt_newrelic
package-version: 2022-07-01-preview
azure-arm: true
```

### Tag: package-2022-07-01-preview and ruby

These settings apply only when `--tag=package-2022-07-01-preview --ruby` is specified on the command line.
Please also specify `--ruby-sdks-folder=<path to the root directory of your azure-sdk-for-ruby clone>`.

```yaml $(tag) == 'package-2022-07-01-preview' && $(ruby)
namespace: NewRelic.Observability
output-folder: $(ruby-sdks-folder)/newrelic
```
