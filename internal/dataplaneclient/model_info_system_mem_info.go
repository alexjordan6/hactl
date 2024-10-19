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

type InfoSystemMemInfo struct {
	DataplaneapiMemory int32 `json:"dataplaneapi_memory,omitempty"`
	FreeMemory         int32 `json:"free_memory,omitempty"`
	TotalMemory        int32 `json:"total_memory,omitempty"`
}
