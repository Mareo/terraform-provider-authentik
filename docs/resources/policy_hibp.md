---
page_title: "authentik_policy_hibp Resource - terraform-provider-authentik"
subcategory: "Customization"
description: |-
  
---

# authentik_policy_hibp (Resource)



## Example Usage

{{tffile "examples/resources/authentik_policy_hibp/resource.tf"}}

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `allowed_count` (Number) Defaults to `1`.
- `execution_logging` (Boolean) Defaults to `false`.
- `password_field` (String) Defaults to `password`.

### Read-Only

- `id` (String) The ID of this resource.


