# To get the ID of an outpost by managed value

data "authentik_group" "embedded" {
  managed = "goauthentik.io/outposts/embedded"
}

# Then use `data.authentik_group.embedded.id`
