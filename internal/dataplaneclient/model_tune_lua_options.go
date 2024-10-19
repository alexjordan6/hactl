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

type TuneLuaOptions struct {
	BurstTimeout   int32  `json:"burst_timeout,omitempty"`
	ForcedYield    int32  `json:"forced_yield,omitempty"`
	LogLoggers     string `json:"log_loggers,omitempty"`
	LogStderr      string `json:"log_stderr,omitempty"`
	Maxmem         int32  `json:"maxmem,omitempty"`
	ServiceTimeout int32  `json:"service_timeout,omitempty"`
	SessionTimeout int32  `json:"session_timeout,omitempty"`
	TaskTimeout    int32  `json:"task_timeout,omitempty"`
}
