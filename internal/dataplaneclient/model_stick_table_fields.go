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

type StickTableFields struct {
	Field string `json:"field,omitempty"`
	Period int32 `json:"period,omitempty"`
	Type_ string `json:"type,omitempty"`
}
