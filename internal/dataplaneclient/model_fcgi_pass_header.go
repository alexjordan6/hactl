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

// Specifies the name of a request header to pass to the FastCGI application. Optionally, you can follow it with an ACL-based condition, in which case the FastCGI application evaluates it only if the condition is true. Most request headers are already available to the FastCGI application with the prefix \"HTTP\". Thus, you only need this directive to pass headers that are purposefully omitted. Currently, the headers \"Authorization\", \"Proxy-Authorization\", and hop-by-hop headers are omitted. Note that the headers \"Content-type\" and \"Content-length\" never pass to the FastCGI application because they are already converted into parameters.
type FcgiPassHeader struct {
	Cond     string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Name     string `json:"name,omitempty"`
}
