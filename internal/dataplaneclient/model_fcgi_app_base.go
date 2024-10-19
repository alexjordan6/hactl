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

// HAProxy FastCGI application configuration
type FcgiAppBase struct {
	// Defines the document root on the remote host. The parameter serves to build the default value of FastCGI parameters SCRIPT_FILENAME and PATH_TRANSLATED. It is a mandatory setting.
	Docroot string `json:"docroot"`
	// Enables or disables the retrieval of variables related to connection management.
	GetValues string `json:"get_values,omitempty"`
	// Defines the script name to append after a URI that ends with a slash (\"/\") to set the default value for the FastCGI parameter SCRIPT_NAME. It is an optional setting.
	Index string `json:"index,omitempty"`
	// Tells the FastCGI application whether or not to keep the connection open after it sends a response. If disabled, the FastCGI application closes the connection after responding to this request.
	KeepConn string `json:"keep_conn,omitempty"`
	LogStderrs []FcgiLogStderr `json:"log_stderrs,omitempty"`
	// Defines the maximum number of concurrent requests this application can accept. If the FastCGI application retrieves the variable FCGI_MAX_REQS during connection establishment, it can override this option. Furthermore, if the application does not do multiplexing, it will ignore this option.
	MaxReqs int32 `json:"max_reqs,omitempty"`
	// Enables or disables the support of connection multiplexing. If the FastCGI application retrieves the variable FCGI_MPXS_CONNS during connection establishment, it can override this option.
	MpxsConns string `json:"mpxs_conns,omitempty"`
	// Declares a FastCGI application
	Name string `json:"name"`
	PassHeaders []FcgiPassHeader `json:"pass_headers,omitempty"`
	// Defines a regular expression to extract the script-name and the path-info from the URI. Thus, <regex> must have two captures: the first to capture the script name, and the second to capture the path- info. If not defined, it does not perform matching on the URI, and does not fill the FastCGI parameters PATH_INFO and PATH_TRANSLATED.
	PathInfo string `json:"path_info,omitempty"`
	SetParams []FcgiSetParam `json:"set_params,omitempty"`
}
