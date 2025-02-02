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

// HAProxy HTTP error rule configuration (corresponds to http-error directives)
type HttpErrorRule struct {
	ReturnContent       string         `json:"return_content,omitempty"`
	ReturnContentFormat string         `json:"return_content_format,omitempty"`
	ReturnContentType   string         `json:"return_content_type,omitempty"`
	ReturnHdrs          []ReturnHeader `json:"return_hdrs,omitempty"`
	Status              int32          `json:"status"`
	Type_               string         `json:"type"`
}
