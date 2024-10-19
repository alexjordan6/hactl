# \DgramBindApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDgramBind**](DgramBindApi.md#CreateDgramBind) | **Post** /services/haproxy/configuration/log_forwards/{parent_name}/dgram_binds | Add a new dgram bind
[**DeleteDgramBind**](DgramBindApi.md#DeleteDgramBind) | **Delete** /services/haproxy/configuration/log_forwards/{parent_name}/dgram_binds/{name} | Delete a dgram bind
[**GetDgramBind**](DgramBindApi.md#GetDgramBind) | **Get** /services/haproxy/configuration/log_forwards/{parent_name}/dgram_binds/{name} | Return one dgram bind
[**GetDgramBinds**](DgramBindApi.md#GetDgramBinds) | **Get** /services/haproxy/configuration/log_forwards/{parent_name}/dgram_binds | Return an array of dgram binds
[**ReplaceDgramBind**](DgramBindApi.md#ReplaceDgramBind) | **Put** /services/haproxy/configuration/log_forwards/{parent_name}/dgram_binds/{name} | Replace a dgram bind


# **CreateDgramBind**
> DgramBind CreateDgramBind(ctx, parentName, data, optional)
Add a new dgram bind

Adds a new dgram bind in the specified log forward in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**DgramBind**](DgramBind.md)|  | 
 **optional** | ***DgramBindApiCreateDgramBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DgramBindApiCreateDgramBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**DgramBind**](dgram_bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDgramBind**
> DeleteDgramBind(ctx, name, parentName, optional)
Delete a dgram bind

Deletes a dgram bind configuration by it's name in the specified log forward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
  **parentName** | **string**| Parent name | 
 **optional** | ***DgramBindApiDeleteDgramBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DgramBindApiDeleteDgramBindOpts struct

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

# **GetDgramBind**
> DgramBind GetDgramBind(ctx, name, parentName, optional)
Return one dgram bind

Returns one dgram bind configuration by it's name in the specified log forward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
  **parentName** | **string**| Parent name | 
 **optional** | ***DgramBindApiGetDgramBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DgramBindApiGetDgramBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**DgramBind**](dgram_bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDgramBinds**
> DgramBinds GetDgramBinds(ctx, parentName, optional)
Return an array of dgram binds

Returns an array of all dgram binds that are configured in specified log forward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***DgramBindApiGetDgramBindsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DgramBindApiGetDgramBindsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**DgramBinds**](dgram_binds.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDgramBind**
> DgramBind ReplaceDgramBind(ctx, name, parentName, data, optional)
Replace a dgram bind

Replaces a dgram bind configuration by it's name in the specified log forward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Bind name | 
  **parentName** | **string**| Parent name | 
  **data** | [**DgramBind**](DgramBind.md)|  | 
 **optional** | ***DgramBindApiReplaceDgramBindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DgramBindApiReplaceDgramBindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**DgramBind**](dgram_bind.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

