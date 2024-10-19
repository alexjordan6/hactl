# \TCPCheckApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPCheckBackend**](TCPCheckApi.md#CreateTCPCheckBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/tcp_checks/{index} | Add a new TCP check
[**CreateTCPCheckDefaults**](TCPCheckApi.md#CreateTCPCheckDefaults) | **Post** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks/{index} | Add a new TCP check
[**DeleteTCPCheckBackend**](TCPCheckApi.md#DeleteTCPCheckBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/tcp_checks/{index} | Delete a TCP check
[**DeleteTCPCheckDefaults**](TCPCheckApi.md#DeleteTCPCheckDefaults) | **Delete** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks/{index} | Delete a TCP check
[**GetAllTCPCheckBackend**](TCPCheckApi.md#GetAllTCPCheckBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_checks | Return an array of TCP checks
[**GetAllTCPCheckDefaults**](TCPCheckApi.md#GetAllTCPCheckDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks | Return an array of TCP checks
[**GetTCPCheckBackend**](TCPCheckApi.md#GetTCPCheckBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_checks/{index} | Return one TCP check
[**GetTCPCheckDefaults**](TCPCheckApi.md#GetTCPCheckDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks/{index} | Return one TCP check
[**ReplaceAllTCPCheckBackend**](TCPCheckApi.md#ReplaceAllTCPCheckBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_checks | Replace an TCP Check list
[**ReplaceAllTCPCheckDefaults**](TCPCheckApi.md#ReplaceAllTCPCheckDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks | Replace an TCP Check list
[**ReplaceTCPCheckBackend**](TCPCheckApi.md#ReplaceTCPCheckBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_checks/{index} | Replace a TCP check
[**ReplaceTCPCheckDefaults**](TCPCheckApi.md#ReplaceTCPCheckDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/tcp_checks/{index} | Replace a TCP check


# **CreateTCPCheckBackend**
> TcpCheck CreateTCPCheckBackend(ctx, index, parentName, data, optional)
Add a new TCP check

Adds a new TCP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiCreateTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiCreateTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTCPCheckDefaults**
> TcpCheck CreateTCPCheckDefaults(ctx, index, parentName, data, optional)
Add a new TCP check

Adds a new TCP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiCreateTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiCreateTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPCheckBackend**
> DeleteTCPCheckBackend(ctx, index, parentName, optional)
Delete a TCP check

Deletes a TCP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiDeleteTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiDeleteTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPCheckDefaults**
> DeleteTCPCheckDefaults(ctx, index, parentName, optional)
Delete a TCP check

Deletes a TCP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiDeleteTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiDeleteTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTCPCheckBackend**
> TcpChecks GetAllTCPCheckBackend(ctx, parentName, optional)
Return an array of TCP checks

Returns all TCP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiGetAllTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetAllTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpChecks**](tcp_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTCPCheckDefaults**
> TcpChecks GetAllTCPCheckDefaults(ctx, parentName, optional)
Return an array of TCP checks

Returns all TCP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiGetAllTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetAllTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpChecks**](tcp_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPCheckBackend**
> TcpCheck GetTCPCheckBackend(ctx, index, parentName, optional)
Return one TCP check

Returns one TCP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiGetTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPCheckDefaults**
> TcpCheck GetTCPCheckDefaults(ctx, index, parentName, optional)
Return one TCP check

Returns one TCP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPCheckApiGetTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiGetTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTCPCheckBackend**
> TcpChecks ReplaceAllTCPCheckBackend(ctx, parentName, data, optional)
Replace an TCP Check list

Replaces a whole list of TCP Checks with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**TcpChecks**](TcpChecks.md)|  | 
 **optional** | ***TCPCheckApiReplaceAllTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiReplaceAllTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpChecks**](tcp_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTCPCheckDefaults**
> TcpChecks ReplaceAllTCPCheckDefaults(ctx, parentName, data, optional)
Replace an TCP Check list

Replaces a whole list of TCP Checks with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**TcpChecks**](TcpChecks.md)|  | 
 **optional** | ***TCPCheckApiReplaceAllTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiReplaceAllTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpChecks**](tcp_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPCheckBackend**
> TcpCheck ReplaceTCPCheckBackend(ctx, index, parentName, data, optional)
Replace a TCP check

Replaces a TCP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiReplaceTCPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiReplaceTCPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPCheckDefaults**
> TcpCheck ReplaceTCPCheckDefaults(ctx, index, parentName, data, optional)
Replace a TCP check

Replaces a TCP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpCheck**](TcpCheck.md)|  | 
 **optional** | ***TCPCheckApiReplaceTCPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPCheckApiReplaceTCPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpCheck**](tcp_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

