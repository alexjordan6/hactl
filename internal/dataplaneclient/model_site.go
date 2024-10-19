/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dataplaneclient

// Site configuration. Sites are considered as one service and all farms connected to that service. Farms are connected to service using use-backend and default_backend directives. Sites let you configure simple HAProxy configurations, for more advanced options use /haproxy/configuration endpoints.
type Site struct {
	Farms   []SiteFarms  `json:"farms,omitempty"`
	Name    string       `json:"name"`
	Service *SiteService `json:"service,omitempty"`
}
