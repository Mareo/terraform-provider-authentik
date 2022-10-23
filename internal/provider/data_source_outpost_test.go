package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOutpost(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOutpostManaged,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "name", "authentik Embedded Outpost"),
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "managed", "goauthentik.io/outposts/embedded"),
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "type", "proxy"),
				),
			},
			{
				Config: testAccDataSourceOutpostName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "name", "authentik Embedded Outpost"),
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "managed", "goauthentik.io/outposts/embedded"),
					resource.TestCheckResourceAttr("data.authentik_outpost.embedded", "type", "proxy"),
				),
			},
		},
	})
}

const testAccDataSourceOutpostManaged = `
data "authentik_group" "embedded" {
  managed = "goauthentik.io/outposts/embedded"
}
`

const testAccDataSourceOutpostManagedName = `
data "authentik_group" "embedded" {
  name = "authentik Embedded Outpost"
}
`
