---
page_title: "authentik_policy_expiry Resource - terraform-provider-authentik"
subcategory: "Customization"
description: |-
  
---

# authentik_policy_expiry (Resource)



## Example Usage

{{tffile "examples/resources/authentik_policy_expiry/resource.tf"}}

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `days` (Number)
- `name` (String)

### Optional

- `deny_only` (Boolean) Defaults to `false`.
- `execution_logging` (Boolean) Defaults to `false`.

### Read-Only

- `id` (String) The ID of this resource.


