package ovh

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

func resourceHostingWebAttachedDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceHostingWebAttachedDomainCreate,
		Read:   resourceHostingWebAttachedDomainRead,
		Delete: resourceHostingWebAttachedDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceHostingWebAttachedDomainImportState,
		},
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Description: "The internal name of your private database",
			},
			"cdn": {
				Type:        schema.TypeString,
				Description: "The internal name of your private database",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Linked domain (fqdn)",
			},
			"firewall": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Firewall state for this path",
			},
			"own_log": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Firewall state for this path",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Firewall state for this path",
			},
		},
	}
}

func resourceHostingWebAttachedDomainCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	databaseName := d.Get("database_name").(string)

	opts := (&HostingWebAttachedDomainCreateOpts{}).FromResource(d)
	ds := &HostingWebAttachedDomain{}

	log.Printf("[DEBUG][Create] HostingWebAttachedDomain")
	endpoint := fmt.Sprintf("/hosting/privateDatabase/%s/database", url.PathEscape(serviceName))
	err := config.OVHClient.Post(endpoint, opts, &ds)
	if err != nil {
		return fmt.Errorf("failed to create database: %s", err)
	}

	log.Printf("[DEBUG][Create][WaitForArchived] HostingWebAttachedDomain")
	endpoint = fmt.Sprintf("/hosting/privateDatabase/%s/tasks/%d", url.PathEscape(serviceName), ds.TaskId)
	err = WaitArchivedHostingTask(config.OVHClient, endpoint, 2*time.Minute)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceName, databaseName))
	return resourceHostingWebAttachedDomainRead(d, meta)
}

func resourceHostingWebAttachedDomainRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	databaseName := d.Get("database_name").(string)

	ds := &HostingWebAttachedDomain{}

	log.Printf("[DEBUG][Read] HostingWebAttachedDomain")
	endpoint := fmt.Sprintf("/hosting/privateDatabase/%s/database/%s", url.PathEscape(serviceName), url.PathEscape(databaseName))
	if err := config.OVHClient.Get(endpoint, &ds); err != nil {
		return helpers.CheckDeleted(d, err, endpoint)
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceName, databaseName))
	for k, v := range ds.ToMap() {
		d.Set(k, v)
	}

	return nil
}

func resourceHostingWebAttachedDomainDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	databaseName := d.Get("database_name").(string)

	ds := &HostingWebAttachedDomain{}

	log.Printf("[DEBUG][Delete] HostingWebAttachedDomain")
	endpoint := fmt.Sprintf("/hosting/privateDatabase/%s/database/%s", url.PathEscape(serviceName), url.PathEscape(databaseName))
	if err := config.OVHClient.Delete(endpoint, ds); err != nil {
		return helpers.CheckDeleted(d, err, endpoint)
	}

	log.Printf("[DEBUG][Delete][WaitForArchived] HostingWebAttachedDomain")
	endpoint = fmt.Sprintf("/hosting/privateDatabase/%s/tasks/%d", url.PathEscape(serviceName), ds.TaskId)
	err := WaitArchivedHostingTask(config.OVHClient, endpoint, 2*time.Minute)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceHostingWebAttachedDomainImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	givenId := d.Id()
	splitId := strings.SplitN(givenId, "/", 2)

	log.Printf("[DEBUG][Import] HostingWebAttachedDomain givenId: %s", givenId)

	if len(splitId) != 2 {
		return nil, fmt.Errorf("import Id is not SERVICE_NAME/DATABASE_NAME formatted")
	}
	d.SetId(splitId[0])
	d.Set("service_name", splitId[0])
	d.Set("attached_domain", splitId[1])
	results := make([]*schema.ResourceData, 1)
	results[0] = d
	return results, nil
}
