---
layout: ""
page_title: "authentik Provider"
description: |-
  Manage https://goauthentik.io with terraform
---

# authentik Provider

The authentik provider provides resources to interact with the authentik API.

## Example Usage

```terraform
provider "authentik" {
  url   = "https://beryjuorg-dev.my.goauthentik.io"
  token = "foo-bar"
  # Optionally set insecure to ignore TLS Certificates
  # insecure = true
}
```

### Configure provider with environment variables
It is optionally possible to configure the provider by passing environment variables to terraform
```bash
export AUTHENTIK_URL=https://...
export AUTHENTIK_TOKEN=<secret_token>
export AUTHENTIK_INSECURE=false
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `insecure` (Boolean) Whether to skip TLS verification, can optionally be passed as `AUTHENTIK_INSECURE` environmental variable
- `token` (String, Sensitive) The authentik API token, can optionally be passed as `AUTHENTIK_TOKEN` environmental variable
- `url` (String) The authentik API endpoint, can optionally be passed as `AUTHENTIK_URL` environmental variable
