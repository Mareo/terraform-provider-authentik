package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceStageAuthenticatorTOTP(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceStageAuthenticatorTOTP(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("authentik_stage_authenticator_totp.name", "name", rName),
				),
			},
			{
				Config: testAccResourceStageAuthenticatorTOTP(rName + "test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("authentik_stage_authenticator_totp.name", "name", rName+"test"),
				),
			},
		},
	})
}

func testAccResourceStageAuthenticatorTOTP(name string) string {
	return fmt.Sprintf(`
resource "authentik_stage_authenticator_totp" "name" {
  name              = "%[1]s"
}
`, name)
}
