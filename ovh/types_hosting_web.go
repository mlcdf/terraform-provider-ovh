package ovh

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

type HostingWeb struct {
	ServiceName             string                  `json:"serviceName"`
	AvailableBoostOffer     []AvailableBoostOffer   `json:"availableBoostOffer"`
	BoostOffer              *string                 `json:"boostOffer"`
	CountriesIp             []CountryIp             `json:"countriesIp"`
	Cluster                 string                  `json:"cluster"`
	ClusterIP               string                  `json:"clusterIp"`
	ClusterIPv6             string                  `json:"clusterIpv6"`
	DisplayName             *string                 `json:"displayName"`
	Datacenter              string                  `json:"datacenter"`
	Filer                   string                  `json:"filer"`
	HasCDN                  bool                    `json:"hasCdn"`
	HasHostedSSL            bool                    `json:"hasHostedSsl"`
	Home                    string                  `json:"home"`
	HostingIP               string                  `json:"hostingIp"`
	HostingIPv6             string                  `json:"hostingIpv6"`
	LastOVHConfigScan       *time.Time              `json:"lastOvhConfigScan"`
	Offer                   string                  `json:"offer"`
	OperationSystem         string                  `json:"operationSystem"`
	PHPVersions             []string                `json:"phpVersions"`
	PrimaryLogin            string                  `json:"primaryLogin"`
	QuotaSize               *UnitAndValue           `json:"quotasize"`
	QuotaUsed               *UnitAndValue           `json:"quotaUsed"`
	RecommendedOffer        *string                 `json:"recommendedOffer"`
	ResourceType            string                  `json:"resourceType"`
	ServiceManagementAccess ServiceManagementAccess `json:"serviceManagementAccess"`
	State                   string                  `json:"state"`
	Token                   string                  `json:"token"`
	TrafficQuotaSize        *UnitAndValue           `json:"trafficQuotaSize"`
	TrafficQuotaUsed        *UnitAndValue           `json:"trafficQuotaUsed"`
	Updates                 []string                `json:"updates"`
}

type AvailableBoostOffer struct {
}

type CountryIp struct {
	Country string `json:"country"`
	IP      string `json:"ip"`
	IPv6    string `json:"ipv6"`
}

type ServiceManagementAccess struct {
	SSH  Address `json:"ssh"`
	FTP  Address `json:"ftp"`
	HTTP Address `json:"http"`
}

type Address struct {
	Port int    `json:"port"`
	URL  string `json:"url"`
}

func (v HostingWeb) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["service_name"] = v.ServiceName

	obj["available_boost_offer"] = v.AvailableBoostOffer
	obj["boost_offer"] = v.BoostOffer
	obj["cluster"] = v.Cluster
	obj["cluster_ip"] = v.ClusterIP
	obj["cluster_ipv6"] = v.ClusterIPv6
	obj["countries_ip"] = v.CountriesIp
	obj["datacenter"] = v.Datacenter
	obj["display_name"] = v.DisplayName
	obj["filer"] = v.Filer
	obj["has_cdn"] = v.HasCDN
	obj["has_hosted_ssl"] = v.HasHostedSSL
	obj["home"] = v.Home
	obj["hosting_ip"] = v.HostingIP
	obj["hosting_ipv6"] = v.HostingIPv6
	obj["last_ovh_config_scan"] = v.LastOVHConfigScan
	obj["offer"] = v.Offer
	obj["operation_system"] = v.OperationSystem
	obj["php_version"] = v.PHPVersions
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
	obj["traffic_quota_size"] = v.TrafficQuotaSize.Value
	obj["traffic_quota_used"] = v.TrafficQuotaUsed.Value
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

type DataSourceHostingWebDatabase struct {
	BackupTime   string                               `json:"backupTime"`
	QuotaUsed    *UnitAndValue                        `json:"quotaUsed"`
	CreationDate string                               `json:"creationDate"`
	Users        []*DataSourceHostingWebDatabaseUsers `json:"users"`
}

func (v DataSourceHostingWebDatabase) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["backup_time"] = v.BackupTime
	obj["quota_used"] = v.QuotaUsed.Value
	obj["creation_date"] = v.CreationDate

	var users []map[string]interface{}
	for _, r := range v.Users {
		users = append(users, r.ToMap())
	}
	obj["users"] = users
	return obj
}

type DataSourceHostingWebDatabaseUsers struct {
	UserName  string `json:"userName"`
	GrantType string `json:"grantType"`
}

func (v DataSourceHostingWebDatabaseUsers) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["user_name"] = v.UserName
	obj["grant_type"] = v.GrantType
	return obj
}

type HostingWebDatabase struct {
	DoneDate     string `json:"doneDate"`
	LastUpdate   string `json:"lastUpdate"`
	UserName     string `json:"userName"`
	DumpId       string `json:"dumpId"`
	DatabaseName string `json:"databaseName"`
	TaskId       int    `json:"id"`
	StartDate    string `json:"startDate"`
	Status       string `json:"status"`
}

func (v HostingWebDatabase) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["database_name"] = v.DatabaseName
	return obj
}

type HostingWebDatabaseCreateOpts struct {
	DatabaseName string `json:"databaseName"`
}

func (opts *HostingWebDatabaseCreateOpts) FromResource(d *schema.ResourceData) *HostingWebDatabaseCreateOpts {
	opts.DatabaseName = d.Get("database_name").(string)

	return opts
}

type DataSourceHostingWebUser struct {
	CreationDate string                               `json:"creationDate"`
	Databases    []*DataSourceHostingWebUserDatabases `json:"databases"`
}

func (v DataSourceHostingWebUser) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["creation_date"] = v.CreationDate

	var databases []map[string]interface{}
	for _, r := range v.Databases {
		databases = append(databases, r.ToMap())
	}
	obj["databases"] = databases
	return obj
}

type DataSourceHostingWebUserDatabases struct {
	DatabaseName string `json:"databaseName"`
	GrantType    string `json:"grantType"`
}

func (v DataSourceHostingWebUserDatabases) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["database_name"] = v.DatabaseName
	obj["grant_type"] = v.GrantType
	return obj
}

type HostingWebUser struct {
	LastUpdate   string `json:"lastUpdate"`
	DoneDate     string `json:"doneDate"`
	Status       string `json:"status"`
	StartDate    string `json:"startDate"`
	DatabaseName string `json:"databaseName"`
	UserName     string `json:"userName"`
	TaskId       int    `json:"id"`
}

func (v HostingWebUser) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["user_name"] = v.UserName
	return obj
}

type HostingWebUserCreateOpts struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

func (opts *HostingWebUserCreateOpts) FromResource(d *schema.ResourceData) *HostingWebUserCreateOpts {
	opts.Password = d.Get("password").(string)
	opts.UserName = d.Get("user_name").(string)

	return opts
}

type HostingWebUserGrant struct {
	LastUpdate   string `json:"lastUpdate"`
	DoneDate     string `json:"doneDate"`
	Status       string `json:"status"`
	StartDate    string `json:"startDate"`
	DatabaseName string `json:"databaseName"`
	UserName     string `json:"userName"`
	TaskId       int    `json:"id"`
}

type HostingWebUserGrantCreateOpts struct {
	DatabaseName string `json:"databaseName"`
	Grant        string `json:"grant"`
}

func (opts *HostingWebUserGrantCreateOpts) FromResource(d *schema.ResourceData) *HostingWebUserGrantCreateOpts {
	opts.DatabaseName = d.Get("database_name").(string)
	opts.Grant = d.Get("grant").(string)

	return opts
}

func (v HostingWebUserGrantCreateOpts) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["grant"] = v.Grant
	obj["database_name"] = v.DatabaseName
	return obj
}

type DataSourceHostingWebUserGrant struct {
	CreationDate string `json:"creationDate"`
	Grant        string `json:"grant"`
}

func (v DataSourceHostingWebUserGrant) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["creation_date"] = v.CreationDate
	obj["grant"] = v.Grant

	return obj
}

type HostingWebWhitelist struct {
	CreationDate string `json:"creationDate"`
	LastUpdate   string `json:"lastUpdate"`
	Name         string `json:"name"`
	Service      bool   `json:"service"`
	Sftp         bool   `json:"sftp"`
	Status       string `json:"status"`
	TaskId       int    `json:"id"`
}

func (v HostingWebWhitelist) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["name"] = v.Name
	obj["service"] = v.Service
	obj["sftp"] = v.Sftp

	return obj
}

func (v HostingWebWhitelist) DataSourceToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["name"] = v.Name
	obj["service"] = v.Service
	obj["sftp"] = v.Sftp
	obj["creation_date"] = v.CreationDate
	obj["last_update"] = v.LastUpdate
	obj["status"] = v.Status

	return obj
}

type HostingWebWhitelistCreateOpts struct {
	Ip      string `json:"ip"`
	Name    string `json:"name"`
	Service bool   `json:"service"`
	Sftp    bool   `json:"sftp"`
}

func (opts *HostingWebWhitelistCreateOpts) FromResource(d *schema.ResourceData) *HostingWebWhitelistCreateOpts {
	opts.Ip = d.Get("ip").(string)
	opts.Name = d.Get("name").(string)
	opts.Service = d.Get("service").(bool)
	opts.Sftp = d.Get("sftp").(bool)

	return opts
}

type HostingWebWhitelistUpdateOpts struct {
	Name    string `json:"name"`
	Service bool   `json:"service"`
	Sftp    bool   `json:"sftp"`
}

func (opts *HostingWebWhitelistUpdateOpts) FromResource(d *schema.ResourceData) *HostingWebWhitelistUpdateOpts {
	opts.Name = d.Get("name").(string)
	opts.Service = d.Get("service").(bool)
	opts.Sftp = d.Get("sftp").(bool)

	return opts
}
