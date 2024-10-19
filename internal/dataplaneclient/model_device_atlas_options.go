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

type DeviceAtlasOptions struct {
	JsonFile         string `json:"json_file,omitempty"`
	LogLevel         string `json:"log_level,omitempty"`
	PropertiesCookie string `json:"properties_cookie,omitempty"`
	Separator        string `json:"separator,omitempty"`
}
