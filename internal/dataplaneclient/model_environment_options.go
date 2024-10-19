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

type EnvironmentOptions struct {
	Presetenv []EnvironmentOptionsPresetenv `json:"presetenv,omitempty"`
	Resetenv string `json:"resetenv,omitempty"`
	Setenv []EnvironmentOptionsSetenv `json:"setenv,omitempty"`
	Unsetenv string `json:"unsetenv,omitempty"`
}
