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

type Source struct {
	Address       string `json:"address"`
	AddressSecond string `json:"address_second,omitempty"`
	Hdr           string `json:"hdr,omitempty"`
	Interface_    string `json:"interface,omitempty"`
	Occ           string `json:"occ,omitempty"`
	Port          int32  `json:"port,omitempty"`
	PortSecond    int32  `json:"port_second,omitempty"`
	Usesrc        string `json:"usesrc,omitempty"`
}
