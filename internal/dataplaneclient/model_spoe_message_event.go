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

type SpoeMessageEvent struct {
	Cond string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Name string `json:"name"`
}
