# Add a provider to an existing outpost

data "authentik_outpost" "embedded" {
  name = "authentik Embedded Outpost"
}

resource "authentik_outpost_provider" "proxy" {
  outpost  = data.authentik_outpost.embedded.id
  provider = authentik_provider_proxy.proxy.id
}
