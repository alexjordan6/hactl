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

// Sets a FastCGI parameter to pass to this application. Its value, defined by <format> can take a formatted string, the same as the log directive. Optionally, you can follow it with an ACL-based condition, in which case the FastCGI application evaluates it only if the condition is true.
type FcgiSetParam struct {
	Cond     string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Format   string `json:"format,omitempty"`
	Name     string `json:"name,omitempty"`
}
