# To get the complete outposts list

data "authentik_outposts" "all" {
}

# Then use `data.authentik_outposts.all.outposts`

# Or, to filter according to a specific field

data "authentik_outposts" "proxies" {
  type = "proxy"
}

# Then use `data.authentik_groups.proxies.outposts`
