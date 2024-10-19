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

// Peer Section with all it's children resources
type PeerSection struct {
	DefaultBind *DefaultBind `json:"default_bind,omitempty"`
	DefaultServer *DefaultServer `json:"default_server,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Name string `json:"name"`
	// In some configurations, one would like to distribute the stick-table contents to some peers in place of sending all the stick-table contents to each peer declared in the \"peers\" section. In such cases, \"shards\" specifies the number of peer involved in this stick-table contents distribution.
	Shards int32 `json:"shards,omitempty"`
	Binds interface{} `json:"binds,omitempty"`
	LogTargetList *LogTargets `json:"log_target_list,omitempty"`
	PeerEntries interface{} `json:"peer_entries,omitempty"`
	Servers interface{} `json:"servers,omitempty"`
}
