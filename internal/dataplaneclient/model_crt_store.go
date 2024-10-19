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

// Storage mechanism to load and store certificates used in the configuration
type CrtStore struct {
	// Default directory to fetch SSL certificates from
	CrtBase string `json:"crt_base,omitempty"`
	// Default directory to fetch SSL private keys from
	KeyBase string `json:"key_base,omitempty"`
	Loads *CrtLoads `json:"loads,omitempty"`
	Name string `json:"name"`
}
