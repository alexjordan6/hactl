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

type HashType struct {
	Function string `json:"function,omitempty"`
	Method   string `json:"method,omitempty"`
	Modifier string `json:"modifier,omitempty"`
}
