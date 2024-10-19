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

type AwsFilters struct {
	// Key to use as filter, using the format specified at https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html#options
	Key string `json:"key"`
	// Value of the filter to use
	Value string `json:"value"`
}
