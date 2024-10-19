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

// MailersSection with all it's children resources
type MailersSection struct {
	Name          string      `json:"name"`
	Timeout       int32       `json:"timeout,omitempty"`
	MailerEntries interface{} `json:"mailer_entries,omitempty"`
}
