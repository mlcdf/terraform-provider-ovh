package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataSourceHostingWebConfig = `
// data "ovh_order_cart" "mycart" {
// 	ovh_subsidiary = "fr"
// 	description    = "%s"
// }
	
// data "ovh_order_cart_product_plan" "web" {
// 	cart_id        = data.ovh_order_cart.mycart.id
// 	price_capacity = "renew"
// 	product        = "webHosting"
// 	plan_code      = "pro2014"
// }

// resource "ovh_hosting_web" "web" {
// 	ovh_subsidiary = data.ovh_order_cart.mycart.ovh_subsidiary
// 	payment_mean   = "ovh-account"
// 	display_name   = "%s"

// 	plan {
// 		duration     = "P1Y"
// 		plan_code    = data.ovh_order_cart_product_plan.web.plan_code
// 		pricing_mode = data.ovh_order_cart_product_plan.web.selected_price[0].pricing_mode

// 		configuration {
// 			label = "district"
// 			value = "%s"
// 		}

// 		configuration {
// 			label = "webhosting_domain"
// 			value = "**1"
// 		}
// 	}
// }

// data "ovh_hosting_web" "database" {
//   service_name = ovh_hosting_web.web.service_name
// }

data "ovh_hosting_web" "web" {
	service_name = "mlcdfoj.cluster030.hosting.ovh.net"
}
`

func TestAccDataSourceHostingweb_basic(t *testing.T) {

	desc := acctest.RandomWithPrefix(test_prefix)
	displayName := acctest.RandomWithPrefix(test_prefix)
	district := os.Getenv("OVH_HOSTING_WEB_DISTRICT_TEST")

	config := fmt.Sprintf(
		testAccDataSourceHostingWebConfig,
		desc,
		displayName,
		district,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckHostingWeb(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"cluster",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"cluster_ip",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"cluster_ipv6",
					),
					// resource.TestCheckResourceAttrSet(
					// 	"data.ovh_hosting_web.web",
					// 	"countries_ip",
					// ),
					resource.TestCheckResourceAttr(
						"data.ovh_hosting_web.web",
						"datacenter",
						district,
					),
					resource.TestCheckResourceAttr(
						"data.ovh_hosting_web.web",
						"offer",
						"pro2014",
					),
					resource.TestCheckResourceAttr(
						"data.ovh_hosting_web.web",
						"operation_system",
						"linux",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"home",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"has_cdn",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"has_hosted_ssl",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"last_ovh_config_scan",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"resource_type",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"state",
					),
					resource.TestCheckResourceAttrSet(
						"data.ovh_hosting_web.web",
						"token",
					),
				),
			},
		},
	})
}
