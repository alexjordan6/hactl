# FcgiApp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docroot** | **string** | Defines the document root on the remote host. The parameter serves to build the default value of FastCGI parameters SCRIPT_FILENAME and PATH_TRANSLATED. It is a mandatory setting. | [default to null]
**GetValues** | **string** | Enables or disables the retrieval of variables related to connection management. | [optional] [default to null]
**Index** | **string** | Defines the script name to append after a URI that ends with a slash (\&quot;/\&quot;) to set the default value for the FastCGI parameter SCRIPT_NAME. It is an optional setting. | [optional] [default to null]
**KeepConn** | **string** | Tells the FastCGI application whether or not to keep the connection open after it sends a response. If disabled, the FastCGI application closes the connection after responding to this request. | [optional] [default to null]
**LogStderrs** | [**[]FcgiLogStderr**](fcgiLogStderr.md) |  | [optional] [default to null]
**MaxReqs** | **int32** | Defines the maximum number of concurrent requests this application can accept. If the FastCGI application retrieves the variable FCGI_MAX_REQS during connection establishment, it can override this option. Furthermore, if the application does not do multiplexing, it will ignore this option. | [optional] [default to null]
**MpxsConns** | **string** | Enables or disables the support of connection multiplexing. If the FastCGI application retrieves the variable FCGI_MPXS_CONNS during connection establishment, it can override this option. | [optional] [default to null]
**Name** | **string** | Declares a FastCGI application | [default to null]
**PassHeaders** | [**[]FcgiPassHeader**](fcgiPassHeader.md) |  | [optional] [default to null]
**PathInfo** | **string** | Defines a regular expression to extract the script-name and the path-info from the URI. Thus, &lt;regex&gt; must have two captures: the first to capture the script name, and the second to capture the path- info. If not defined, it does not perform matching on the URI, and does not fill the FastCGI parameters PATH_INFO and PATH_TRANSLATED. | [optional] [default to null]
**SetParams** | [**[]FcgiSetParam**](fcgiSetParam.md) |  | [optional] [default to null]
**AclList** | [***Acls**](acls.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


