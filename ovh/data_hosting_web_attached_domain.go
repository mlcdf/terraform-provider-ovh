package ovh

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHostingWebAttachedDomain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostingWebAttachedDomainRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"available_boost_offer": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Available offers for boost option",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"offer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"price_currency_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"price_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"price_value": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"boost_offer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current boost offer",
			},
			"cluster": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster name",
			},
			"cluster_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ip of the cluster",
			},
			"cluster_ipv6": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ipv6 of the cluster",
			},
			"countries_ip": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Available clusterIp by countries",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"datacenter": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Datacenter where this service is located",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name displayed in customer panel for your web hosting",
			},
			"filer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Filer name",
			},
			"has_cdn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Has a CDN service linked on the hosting",
			},
			"has_hosted_ssl": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Has a HostedSSL service linked on the hosting",
			},
			"home": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Path of the home directory",
			},
			"hosting_ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP of the hosting",
			},
			"hosting_ipv6": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 of the hosting",
			},
			"last_ovh_config_scan": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Date of last ovhConfig scan",
			},
			"offer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hosting's offer",
			},
			"operation_system": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hosting's operating system",
			},
			"php_versions": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of availables PHP versions for this hosting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"support": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"primary_login": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hosting's main login",
			},
			"quota_size": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Space used",
			},
			"quota_used": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Space allowed",
			},
			"recommended_offer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If your offer is old, return a recommended offer to migrate on",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hosting resource type",
			},
			"service_management_access_ftp_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to use to access your hosting via FTP",
			},
			"service_management_access_ftp_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "PORT to use to access your hosting via FTP",
			},
			"service_management_access_http_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to use to see your hosting without a domain",
			},
			"service_management_access_http_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port to use to see your hosting without a domain",
			},
			"service_management_access_ssh_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to use to access your webhosting via SSH",
			},
			"service_management_access_ssh_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port to use to access your webhosting via SSH",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of your hosting",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Use to link an external domain. (This token has to be insert into a TXT field on your dns zone with ovhcontrol subdomain)",
			},
			"updates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of updates on your hosting",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceHostingWebAttachedDomainRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)

	ds := &HostingWeb{}
	err := config.OVHClient.Get(
		fmt.Sprintf(
			"/hosting/web/%s",
			url.PathEscape(serviceName),
		),
		&ds,
	)

	if err != nil {
		return fmt.Errorf(
			"Error calling hosting/web/%s:\n\t %q",
			serviceName,
			err,
		)
	}

	for k, v := range ds.ToMap() {
		if k != "service_name" {
			d.Set(k, v)
		}
	}
	d.SetId(ds.ServiceName)

	return nil
}
