package ovh

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/ovh/go-ovh/ovh"
)

const testAccHostingWebBasic = `
data "ovh_order_cart" "mycart" {
  ovh_subsidiary = "fr"
  description    = "%s"
}
  
data "ovh_order_cart_product_plan" "database" {
  cart_id        = data.ovh_order_cart.mycart.id
  price_capacity = "renew"
  product        = "privateSQL"
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
`

func init() {
	resource.AddTestSweepers("ovh_hosting_web", &resource.Sweeper{
		Name:         "ovh_hosting_web",
		Dependencies: []string{"ovh_hosting_web"},
		F:            testSweepHostingPrivateDatabase,
	})
}

func testSweepHostingPrivateDatabase(region string) error {
	config, err := sharedConfigForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}

	serviceNames := make([]string, 0)
	if err := config.OVHClient.Get("/hosting/web", &serviceNames); err != nil {
		return fmt.Errorf("Error calling GET /hosting/web:\n\t %q", err)
	}

	if len(serviceNames) == 0 {
		log.Print("[DEBUG] No hosting web to sweep")
		return nil
	}

	for _, serviceName := range serviceNames {
		r := &HostingPrivateDatabase{}
		log.Printf("[DEBUG] Will get hosting web: %v", serviceName)
		endpoint := fmt.Sprintf(
			"/hosting/web/%s",
			url.PathEscape(serviceName),
		)

		if err := config.OVHClient.Get(endpoint, r); err != nil {
			return fmt.Errorf("calling Get %s:\n\t %q", endpoint, err)
		}

		log.Printf("[DEBUG] Will delete web: %v", serviceName)

		terminate := func() (string, error) {
			log.Printf("[DEBUG] Will terminate hosting web %s", serviceName)
			endpoint := fmt.Sprintf(
				"/hosting/web/%s/terminate",
				url.PathEscape(serviceName),
			)
			if err := config.OVHClient.Post(endpoint, nil, nil); err != nil {
				if errOvh, ok := err.(*ovh.APIError); ok && (errOvh.Code == 404 || errOvh.Code == 460) {
					return "", nil
				}
				return "", fmt.Errorf("calling Post %s:\n\t %q", endpoint, err)
			}
			return serviceName, nil
		}

		confirmTerminate := func(token string) error {
			log.Printf("[DEBUG] Will confirm termination of hosting web %s", serviceName)
			endpoint := fmt.Sprintf(
				"/hosting/web/%s/confirmTermination",
				url.PathEscape(serviceName),
			)
			if err := config.OVHClient.Post(endpoint, &HostingPrivateDatabaseConfirmTerminationOpts{Token: token}, nil); err != nil {
				return fmt.Errorf("calling Post %s:\n\t %q", endpoint, err)
			}
			return nil
		}

		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			if err := orderDelete(nil, config, terminate, confirmTerminate); err != nil {
				return resource.RetryableError(err)
			}

			// Successful delete
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func TestAccHostingweb_basic(t *testing.T) {
	desc := acctest.RandomWithPrefix(test_prefix)
	displayName := acctest.RandomWithPrefix(test_prefix)
	dc := os.Getenv("OVH_HOSTING_WEB_DC_TEST")
	engine := os.Getenv("OVH_HOSTING_WEB_ENGINE_TEST")

	config := fmt.Sprintf(
		testAccHostingWebBasic,
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
					resource.TestCheckResourceAttr(
						"ovh_hosting_web.web", "display_name", displayName),
					resource.TestCheckResourceAttr(
						"ovh_hosting_web.web", "datacenter", dc),
					resource.TestCheckResourceAttr(
						"ovh_hosting_web.web", "version", engine),
				),
			},
		},
	})
}
