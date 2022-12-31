package ovh

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

type WebHostingBooleanString string

const activeField = "active"
const noneField = "none"

type HostingWeb struct {
	ServiceName             string                            `json:"serviceName"`
	AvailableBoostOffer     []HostingWebAvailableBoostOffer   `json:"availableBoostOffer"`
	BoostOffer              *string                           `json:"boostOffer"`
	CountriesIP             []HostingWebCountryIp             `json:"countriesIp"`
	Cluster                 string                            `json:"cluster"`
	ClusterIP               string                            `json:"clusterIp"`
	ClusterIPv6             string                            `json:"clusterIpv6"`
	DisplayName             *string                           `json:"displayName"`
	Datacenter              string                            `json:"datacenter"`
	Filer                   string                            `json:"filer"`
	HasCDN                  bool                              `json:"hasCdn"`
	HasHostedSSL            bool                              `json:"hasHostedSsl"`
	Home                    string                            `json:"home"`
	HostingIP               string                            `json:"hostingIp"`
	HostingIPv6             string                            `json:"hostingIpv6"`
	LastOVHConfigScan       *time.Time                        `json:"lastOvhConfigScan"`
	Offer                   string                            `json:"offer"`
	OperationSystem         string                            `json:"operatingSystem"`
	PHPVersions             []HostingWebVersion               `json:"phpVersions"`
	PrimaryLogin            string                            `json:"primaryLogin"`
	QuotaSize               *UnitAndValue                     `json:"quotaSize"`
	QuotaUsed               *UnitAndValue                     `json:"quotaUsed"`
	RecommendedOffer        *string                           `json:"recommendedOffer"`
	ResourceType            string                            `json:"resourceType"`
	ServiceManagementAccess HostingWebServiceManagementAccess `json:"serviceManagementAccess"`
	State                   string                            `json:"state"`
	Token                   string                            `json:"token"`
	Updates                 []string                          `json:"updates"`
}

type HostingWebAvailableBoostOffer struct {
	Price OrderCartGenericProductPricePrice `json:"price"`
	Offer string                            `json:"offer"`
}

func (v HostingWebAvailableBoostOffer) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["price"] = v.Price.ToMap()
	obj["offer"] = v.Offer

	return obj
}

type HostingWebCountryIp struct {
	Country string `json:"country"`
	IP      string `json:"ip"`
	IPv6    string `json:"ipv6"`
}

func (v HostingWebCountryIp) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["country"] = v.Country
	obj["ip"] = v.IP
	obj["ipv6"] = v.IPv6

	return obj
}

type HostingWebServiceManagementAccess struct {
	SSH  HostingWebAddress `json:"ssh"`
	FTP  HostingWebAddress `json:"ftp"`
	HTTP HostingWebAddress `json:"http"`
}

type HostingWebAddress struct {
	Port int    `json:"port"`
	URL  string `json:"url"`
}

func (v HostingWebAddress) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["port"] = v.Port
	obj["url"] = v.URL

	return obj
}

type HostingWebVersion struct {
	Support string `json:"support"`
	Version string `json:"version"`
}

func (v HostingWebVersion) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["support"] = v.Support
	obj["version"] = v.Version

	return obj
}

func (v HostingWeb) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["service_name"] = v.ServiceName

	availableBoostOffer := []map[string]interface{}{}
	for _, v := range v.AvailableBoostOffer {
		availableBoostOffer = append(availableBoostOffer, v.ToMap())
	}

	obj["available_boost_offer"] = availableBoostOffer

	obj["boost_offer"] = v.BoostOffer
	obj["cluster"] = v.Cluster
	obj["cluster_ip"] = v.ClusterIP
	obj["cluster_ipv6"] = v.ClusterIPv6

	countriesIP := []map[string]interface{}{}
	for _, v := range v.CountriesIP {
		countriesIP = append(countriesIP, v.ToMap())
	}

	obj["countries_ip"] = countriesIP

	obj["datacenter"] = v.Datacenter
	obj["display_name"] = v.DisplayName
	obj["filer"] = v.Filer
	obj["has_cdn"] = v.HasCDN
	obj["has_hosted_ssl"] = v.HasHostedSSL
	obj["home"] = v.Home
	obj["hosting_ip"] = v.HostingIP
	obj["hosting_ipv6"] = v.HostingIPv6
	obj["last_ovh_config_scan"] = v.LastOVHConfigScan.String()
	obj["offer"] = v.Offer
	obj["operation_system"] = v.OperationSystem

	phpVersions := []map[string]interface{}{}
	for _, v := range v.PHPVersions {
		phpVersions = append(phpVersions, v.ToMap())
	}

	obj["php_versions"] = phpVersions

	obj["primary_login"] = v.PrimaryLogin
	obj["quota_size"] = v.QuotaSize.Value
	obj["quota_used"] = v.QuotaUsed.Value
	obj["recommended_offer"] = v.RecommendedOffer
	obj["resource_type"] = v.ResourceType
	obj["service_management_access_ftp_url"] = v.ServiceManagementAccess.FTP.URL
	obj["service_management_access_ftp_port"] = v.ServiceManagementAccess.FTP.Port
	obj["service_management_access_http_url"] = v.ServiceManagementAccess.HTTP.URL
	obj["service_management_access_http_port"] = v.ServiceManagementAccess.HTTP.Port
	obj["service_management_access_ssh_url"] = v.ServiceManagementAccess.SSH.URL
	obj["service_management_access_ssh_port"] = v.ServiceManagementAccess.SSH.Port
	obj["state"] = v.State
	obj["token"] = v.Token

	obj["updates"] = v.Updates

	return obj
}

type HostingWebOpts struct {
	DisplayName *string `json:"displayName"`
}

func (opts *HostingWebOpts) FromResource(d *schema.ResourceData) *HostingWebOpts {
	opts.DisplayName = helpers.GetNilStringPointerFromData(d, "display_name")

	return opts
}

type HostingWebConfirmTerminationOpts struct {
	Token string `json:"token"`
}

type HostingWebAttachedDomain struct {
	CDN         WebHostingBooleanString `json:"cdn"`
	Domain      string                  `json:"domain"`
	Firewall    WebHostingBooleanString `json:"firewall"`
	IPLocation  string                  `json:"ipLocation"`
	IsFlushable *string                 `json:"isFlushable"`
	OwnLog      string                  `json:"ownLog"`
	Path        string                  `json:"path"`
	RuntimeID   *string                 `json:"runtimeId"`
	SSL         bool                    `json:"ssl"`
	Status      string                  `json:"status"`
	TaskId      *int                    `json:"id"`
}

func (v HostingWebAttachedDomain) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["domain"] = v.Domain

	return obj
}

type DatasourceHostingWebAttachedDomain struct {
	CDN        WebHostingBooleanString `json:"cdn"`
	Domain     string                  `json:"domain"`
	Firewall   WebHostingBooleanString `json:"firewall"`
	IPLocation string                  `json:"ipLocation"`
	OwnLog     string                  `json:"ownLog"`
	Path       string                  `json:"path"`
	RuntimeID  *string                 `json:"runtimeId"`
	SSL        bool                    `json:"ssl"`
	Status     string                  `json:"status"`
}

func (v DatasourceHostingWebAttachedDomain) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["domain"] = v.Domain

	if v.CDN == activeField {
		obj["cdn"] = true
	} else {
		obj["cdn"] = false
	}

	if v.Firewall == activeField {
		obj["firewall"] = true
	} else {
		obj["firewall"] = false
	}

	obj["ip_location"] = v.IPLocation
	obj["own_log"] = v.OwnLog
	obj["path"] = v.Path
	obj["runtime_id"] = v.RuntimeID
	obj["ssl"] = v.SSL
	obj["status"] = v.Status

	return obj
}

type HostingWebAttachedDomainCreateOpts struct {
	Domain    string `json:"domain"`
	CDN       string `json:"cdn,omitempty"`
	Firewall  string `json:"firewall,omitempty"`
	OwnLog    string `json:"ownLog,omitempty"`
	Path      string `json:"path"`
	RuntimeID int    `json:"runtimeId,omitempty"`
	SSL       bool   `json:"ssl,omitempty"`
}

func (opts *HostingWebAttachedDomainCreateOpts) FromResource(d *schema.ResourceData) *HostingWebAttachedDomainCreateOpts {
	opts.Domain = d.Get("domain").(string)

	return opts
}
