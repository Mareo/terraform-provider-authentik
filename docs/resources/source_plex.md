---
page_title: "authentik_source_plex Resource - terraform-provider-authentik"
subcategory: "Directory"
description: |-
  
---

# authentik_source_plex (Resource)



## Example Usage

{{tffile "examples/resources/authentik_source_plex/resource.tf"}}

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `authentication_flow` (String)
- `client_id` (String)
- `enrollment_flow` (String)
- `name` (String)
- `plex_token` (String, Sensitive)
- `slug` (String)

### Optional

- `allow_friends` (Boolean) Defaults to `true`.
- `allowed_servers` (List of String)
- `enabled` (Boolean) Defaults to `true`.
- `policy_engine_mode` (String) Defaults to `any`.
- `user_matching_mode` (String) Defaults to `identifier`.
- `user_path_template` (String) Defaults to `goauthentik.io/sources/%(slug)s`.
- `uuid` (String)

### Read-Only

- `id` (String) The ID of this resource.


