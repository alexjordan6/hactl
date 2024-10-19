# \DeclareCaptureApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeclareCapture**](DeclareCaptureApi.md#CreateDeclareCapture) | **Post** /services/haproxy/configuration/frontends/{parent_name}/captures/{index} | Add a new declare capture
[**DeleteDeclareCapture**](DeclareCaptureApi.md#DeleteDeclareCapture) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/captures/{index} | Delete a declare capture
[**GetDeclareCapture**](DeclareCaptureApi.md#GetDeclareCapture) | **Get** /services/haproxy/configuration/frontends/{parent_name}/captures/{index} | Return one declare capture
[**GetDeclareCaptures**](DeclareCaptureApi.md#GetDeclareCaptures) | **Get** /services/haproxy/configuration/frontends/{parent_name}/captures | Return an array of declare captures
[**ReplaceDeclareCapture**](DeclareCaptureApi.md#ReplaceDeclareCapture) | **Put** /services/haproxy/configuration/frontends/{parent_name}/captures/{index} | Replace a declare capture
[**ReplaceDeclareCaptures**](DeclareCaptureApi.md#ReplaceDeclareCaptures) | **Put** /services/haproxy/configuration/frontends/{parent_name}/captures | Replace a declare capture list


# **CreateDeclareCapture**
> Capture CreateDeclareCapture(ctx, index, parentName, data, optional)
Add a new declare capture

Adds a new declare capture in the specified frontend in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**Capture**](Capture.md)|  | 
 **optional** | ***DeclareCaptureApiCreateDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiCreateDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDeclareCapture**
> DeleteDeclareCapture(ctx, index, parentName, optional)
Delete a declare capture

Deletes a declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***DeclareCaptureApiDeleteDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiDeleteDeclareCaptureOpts struct

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

# **GetDeclareCapture**
> Capture GetDeclareCapture(ctx, index, parentName, optional)
Return one declare capture

Returns one declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***DeclareCaptureApiGetDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiGetDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDeclareCaptures**
> Captures GetDeclareCaptures(ctx, parentName, optional)
Return an array of declare captures

Returns an array of all declare capture records that are configured in specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***DeclareCaptureApiGetDeclareCapturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiGetDeclareCapturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**Captures**](captures.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDeclareCapture**
> Capture ReplaceDeclareCapture(ctx, index, parentName, data, optional)
Replace a declare capture

Replaces a declare capture configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Declare Capture Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**Capture**](Capture.md)|  | 
 **optional** | ***DeclareCaptureApiReplaceDeclareCaptureOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiReplaceDeclareCaptureOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Capture**](capture.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDeclareCaptures**
> Captures ReplaceDeclareCaptures(ctx, parentName, data, optional)
Replace a declare capture list

Replaces a whole list of declare capture with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**Captures**](Captures.md)|  | 
 **optional** | ***DeclareCaptureApiReplaceDeclareCapturesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeclareCaptureApiReplaceDeclareCapturesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**Captures**](captures.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

