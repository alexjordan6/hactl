/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs. 
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StatsOptions struct {
	StatsAdmin bool `json:"stats_admin,omitempty"`
	StatsAdminCond string `json:"stats_admin_cond,omitempty"`
	StatsAdminCondTest string `json:"stats_admin_cond_test,omitempty"`
	StatsAuths []StatsAuth `json:"stats_auths,omitempty"`
	StatsEnable bool `json:"stats_enable,omitempty"`
	StatsHideVersion bool `json:"stats_hide_version,omitempty"`
	StatsHttpRequests []StatsHttpRequest `json:"stats_http_requests,omitempty"`
	StatsMaxconn int32 `json:"stats_maxconn,omitempty"`
	StatsRealm bool `json:"stats_realm,omitempty"`
	StatsRealmRealm string `json:"stats_realm_realm,omitempty"`
	StatsRefreshDelay int32 `json:"stats_refresh_delay,omitempty"`
	StatsShowDesc string `json:"stats_show_desc,omitempty"`
	StatsShowLegends bool `json:"stats_show_legends,omitempty"`
	StatsShowModules bool `json:"stats_show_modules,omitempty"`
	StatsShowNodeName string `json:"stats_show_node_name,omitempty"`
	StatsUriPrefix string `json:"stats_uri_prefix,omitempty"`
}
