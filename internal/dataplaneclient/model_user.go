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

// HAProxy userlist user
type User struct {
	Groups         string `json:"groups,omitempty"`
	Password       string `json:"password"`
	SecurePassword bool   `json:"secure_password"`
	Username       string `json:"username"`
}
