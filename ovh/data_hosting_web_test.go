package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataSourceHostingWebConfig = `
data "ovh_order_cart" "mycart" {
	ovh_subsidiary = "fr"
	description    = "%s"
  }
	
  data "ovh_order_cart_product_plan" "" {
	cart_id        = data.ovh_order_cart.mycart.id
	price_capacity = "renew"
	product        = "pro2014"
	plan_code      = "private-sql-512-instance"
  }
	
  resource "ovh_hosting_web" "database" {
	ovh_subsidiary = data.ovh_order_cart.mycart.ovh_subsidiary
	payment_mean   = "ovh-account"
	display_name   = "%s"
  
	plan {
	  duration     = "P1M"
	  plan_code    = data.ovh_order_cart_product_plan.database.plan_code
	  pricing_mode = data.ovh_order_cart_product_plan.database.selected_price[0].pricing_mode
  
	  configuration {
		label = "dc"
		value = "%s"
	  }
  
	  configuration {
		label = "engine"
		value = "%s"
	  }
	}
  }

data "ovh_hosting_web" "database" {
  service_name = ovh_hosting_web.database.service_name
}
`

func TestAccDataSourceHostingweb_basic(t *testing.T) {

	desc := acctest.RandomWithPrefix(test_prefix)
	displayName := acctest.RandomWithPrefix(test_prefix)
	dc := os.Getenv("OVH_HOSTING_web_DC_TEST")
	engine := os.Getenv("OVH_HOSTING_web_ENGINE_TEST")

	config := fmt.Sprintf(
		testAccDataSourceHostingWebConfig,
		desc,
		displayName,
		dc,
		engine,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckHostingWeb(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.database",
						"cpu",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.database",
						"datacenter",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.database",
						"offer",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.database",
						"type",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.database",
						"version_number",
					),
				),
			},
		},
	})
}
