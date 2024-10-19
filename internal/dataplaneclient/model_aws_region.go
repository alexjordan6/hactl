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

// AWS region configuration
type AwsRegion struct {
	// AWS Access Key ID.
	AccessKeyId string `json:"access_key_id,omitempty"`
	// Specify the AWS filters used to filter the EC2 instances to add
	Allowlist []AwsFilters `json:"allowlist,omitempty"`
	// Specify the AWS filters used to filter the EC2 instances to ignore
	Denylist    []AwsFilters `json:"denylist,omitempty"`
	Description string       `json:"description,omitempty"`
	Enabled     bool         `json:"enabled"`
	// Auto generated ID.
	Id string `json:"id,omitempty"`
	// Select which IPv4 address the Service Discovery has to use for the backend server entry
	Ipv4Address string `json:"ipv4_address"`
	Name        string `json:"name"`
	Region      string `json:"region"`
	// Duration in seconds in-between data pulling requests to the AWS region
	RetryTimeout int32 `json:"retry_timeout"`
	// AWS Secret Access Key.
	SecretAccessKey            string `json:"secret_access_key,omitempty"`
	ServerSlotsBase            int32  `json:"server_slots_base,omitempty"`
	ServerSlotsGrowthIncrement int32  `json:"server_slots_growth_increment,omitempty"`
	ServerSlotsGrowthType      string `json:"server_slots_growth_type,omitempty"`
}
