# \HTTPCheckApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPCheckBackend**](HTTPCheckApi.md#CreateHTTPCheckBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/http_checks/{index} | Add a new HTTP check
[**CreateHTTPCheckDefaults**](HTTPCheckApi.md#CreateHTTPCheckDefaults) | **Post** /services/haproxy/configuration/defaults/{parent_name}/http_checks/{index} | Add a new HTTP check
[**DeleteHTTPCheckBackend**](HTTPCheckApi.md#DeleteHTTPCheckBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/http_checks/{index} | Delete a HTTP check
[**DeleteHTTPCheckDefaults**](HTTPCheckApi.md#DeleteHTTPCheckDefaults) | **Delete** /services/haproxy/configuration/defaults/{parent_name}/http_checks/{index} | Delete a HTTP check
[**GetAllHTTPCheckBackend**](HTTPCheckApi.md#GetAllHTTPCheckBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_checks | Return an array of HTTP checks
[**GetAllHTTPCheckDefaults**](HTTPCheckApi.md#GetAllHTTPCheckDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/http_checks | Return an array of HTTP checks
[**GetHTTPCheckBackend**](HTTPCheckApi.md#GetHTTPCheckBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_checks/{index} | Return one HTTP check
[**GetHTTPCheckDefaults**](HTTPCheckApi.md#GetHTTPCheckDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/http_checks/{index} | Return one HTTP check
[**ReplaceAllHTTPCheckBackend**](HTTPCheckApi.md#ReplaceAllHTTPCheckBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_checks | Replace an HTTP checks list
[**ReplaceAllHTTPCheckDefaults**](HTTPCheckApi.md#ReplaceAllHTTPCheckDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/http_checks | Replace an HTTP checks list
[**ReplaceHTTPCheckBackend**](HTTPCheckApi.md#ReplaceHTTPCheckBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_checks/{index} | Replace a HTTP check
[**ReplaceHTTPCheckDefaults**](HTTPCheckApi.md#ReplaceHTTPCheckDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/http_checks/{index} | Replace a HTTP check


# **CreateHTTPCheckBackend**
> HttpCheck CreateHTTPCheckBackend(ctx, index, parentName, data, optional)
Add a new HTTP check

Adds a new HTTP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiCreateHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiCreateHTTPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPCheckDefaults**
> HttpCheck CreateHTTPCheckDefaults(ctx, index, parentName, data, optional)
Add a new HTTP check

Adds a new HTTP check of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiCreateHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiCreateHTTPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPCheckBackend**
> DeleteHTTPCheckBackend(ctx, index, parentName, optional)
Delete a HTTP check

Deletes a HTTP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiDeleteHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiDeleteHTTPCheckBackendOpts struct

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

# **DeleteHTTPCheckDefaults**
> DeleteHTTPCheckDefaults(ctx, index, parentName, optional)
Delete a HTTP check

Deletes a HTTP check configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiDeleteHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiDeleteHTTPCheckDefaultsOpts struct

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

# **GetAllHTTPCheckBackend**
> HttpChecks GetAllHTTPCheckBackend(ctx, parentName, optional)
Return an array of HTTP checks

Returns all HTTP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiGetAllHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetAllHTTPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpChecks**](http_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPCheckDefaults**
> HttpChecks GetAllHTTPCheckDefaults(ctx, parentName, optional)
Return an array of HTTP checks

Returns all HTTP checks that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiGetAllHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetAllHTTPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpChecks**](http_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPCheckBackend**
> HttpCheck GetHTTPCheckBackend(ctx, index, parentName, optional)
Return one HTTP check

Returns one HTTP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiGetHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetHTTPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPCheckDefaults**
> HttpCheck GetHTTPCheckDefaults(ctx, index, parentName, optional)
Return one HTTP check

Returns one HTTP check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPCheckApiGetHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiGetHTTPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPCheckBackend**
> HttpChecks ReplaceAllHTTPCheckBackend(ctx, parentName, data, optional)
Replace an HTTP checks list

Replaces a whole list of HTTP checks with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpChecks**](HttpChecks.md)|  | 
 **optional** | ***HTTPCheckApiReplaceAllHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiReplaceAllHTTPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpChecks**](http_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPCheckDefaults**
> HttpChecks ReplaceAllHTTPCheckDefaults(ctx, parentName, data, optional)
Replace an HTTP checks list

Replaces a whole list of HTTP checks with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpChecks**](HttpChecks.md)|  | 
 **optional** | ***HTTPCheckApiReplaceAllHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiReplaceAllHTTPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpChecks**](http_checks.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPCheckBackend**
> HttpCheck ReplaceHTTPCheckBackend(ctx, index, parentName, data, optional)
Replace a HTTP check

Replaces a HTTP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiReplaceHTTPCheckBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiReplaceHTTPCheckBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPCheckDefaults**
> HttpCheck ReplaceHTTPCheckDefaults(ctx, index, parentName, data, optional)
Replace a HTTP check

Replaces a HTTP Check configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Check Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpCheck**](HttpCheck.md)|  | 
 **optional** | ***HTTPCheckApiReplaceHTTPCheckDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPCheckApiReplaceHTTPCheckDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpCheck**](http_check.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

