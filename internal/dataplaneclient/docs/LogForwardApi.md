# \LogForwardApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLogForward**](LogForwardApi.md#CreateLogForward) | **Post** /services/haproxy/configuration/log_forwards | Add a log forward
[**DeleteLogForward**](LogForwardApi.md#DeleteLogForward) | **Delete** /services/haproxy/configuration/log_forwards/{name} | Delete a log forward
[**GetLogForward**](LogForwardApi.md#GetLogForward) | **Get** /services/haproxy/configuration/log_forwards/{name} | Return a log forward
[**GetLogForwards**](LogForwardApi.md#GetLogForwards) | **Get** /services/haproxy/configuration/log_forwards | Return an array of log forwards
[**ReplaceLogForward**](LogForwardApi.md#ReplaceLogForward) | **Put** /services/haproxy/configuration/log_forwards/{name} | Replace a log forward


# **CreateLogForward**
> LogForward CreateLogForward(ctx, data, optional)
Add a log forward

Adds a new log_forward to the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LogForward**](LogForward.md)|  | 
 **optional** | ***LogForwardApiCreateLogForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogForwardApiCreateLogForwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**LogForward**](log_forward.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogForward**
> DeleteLogForward(ctx, name, optional)
Delete a log forward

Deletes a log forward from the configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Log Forward name | 
 **optional** | ***LogForwardApiDeleteLogForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogForwardApiDeleteLogForwardOpts struct

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

# **GetLogForward**
> LogForward GetLogForward(ctx, name, optional)
Return a log forward

Returns one log forward configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Log Forward name | 
 **optional** | ***LogForwardApiGetLogForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogForwardApiGetLogForwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**LogForward**](log_forward.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogForwards**
> LogForwards GetLogForwards(ctx, optional)
Return an array of log forwards

Returns an array of all configured log forwards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogForwardApiGetLogForwardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogForwardApiGetLogForwardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**LogForwards**](log_forwards.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceLogForward**
> LogForward ReplaceLogForward(ctx, name, data, optional)
Replace a log forward

Replaces a log forward configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Log Forward name | 
  **data** | [**LogForward**](LogForward.md)|  | 
 **optional** | ***LogForwardApiReplaceLogForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogForwardApiReplaceLogForwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**LogForward**](log_forward.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

