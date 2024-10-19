# \CrtLoadApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrtLoad**](CrtLoadApi.md#CreateCrtLoad) | **Post** /services/haproxy/configuration/crt_loads | Add a new certificate to load
[**DeleteCrtLoad**](CrtLoadApi.md#DeleteCrtLoad) | **Delete** /services/haproxy/configuration/crt_loads/{certificate} | Delete a certificate load entry
[**GetCrtLoad**](CrtLoadApi.md#GetCrtLoad) | **Get** /services/haproxy/configuration/crt_loads/{certificate} | Return one certificate load entry
[**GetCrtLoads**](CrtLoadApi.md#GetCrtLoads) | **Get** /services/haproxy/configuration/crt_loads | Return an array of loaded certificates
[**ReplaceCrtLoad**](CrtLoadApi.md#ReplaceCrtLoad) | **Put** /services/haproxy/configuration/crt_loads/{certificate} | Replace a certificate load entry


# **CreateCrtLoad**
> CrtLoad CreateCrtLoad(ctx, crtStore, data, optional)
Add a new certificate to load

Adds a new load entry to the specified crt_store section in the configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crtStore** | **string**| Parent crt_store section name | 
  **data** | [**CrtLoad**](CrtLoad.md)|  | 
 **optional** | ***CrtLoadApiCreateCrtLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrtLoadApiCreateCrtLoadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**CrtLoad**](crt_load.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCrtLoad**
> DeleteCrtLoad(ctx, certificate, crtStore, optional)
Delete a certificate load entry

Deletes a load entry by its certificate name in the specified crt_store section

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Certificate filename | 
  **crtStore** | **string**| Parent crt_store section name | 
 **optional** | ***CrtLoadApiDeleteCrtLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrtLoadApiDeleteCrtLoadOpts struct

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

# **GetCrtLoad**
> CrtLoad GetCrtLoad(ctx, certificate, crtStore, optional)
Return one certificate load entry

Returns one load entry by its certificate name in the specified crt_store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Certificate filename | 
  **crtStore** | **string**| Parent crt_store name | 
 **optional** | ***CrtLoadApiGetCrtLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrtLoadApiGetCrtLoadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**CrtLoad**](crt_load.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCrtLoads**
> CrtLoads GetCrtLoads(ctx, crtStore, optional)
Return an array of loaded certificates

Returns the list of loaded certificates from the specified crt_store

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **crtStore** | **string**| Parent crt_store name | 
 **optional** | ***CrtLoadApiGetCrtLoadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrtLoadApiGetCrtLoadsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**CrtLoads**](crt_loads.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceCrtLoad**
> CrtLoad ReplaceCrtLoad(ctx, certificate, crtStore, data, optional)
Replace a certificate load entry

Replaces a load entry by its certificate name in the specified crt_store section

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| Certificate filename | 
  **crtStore** | **string**| Parent crt_store section name | 
  **data** | [**CrtLoad**](CrtLoad.md)|  | 
 **optional** | ***CrtLoadApiReplaceCrtLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrtLoadApiReplaceCrtLoadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**CrtLoad**](crt_load.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

