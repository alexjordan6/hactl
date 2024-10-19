# \ConfigurationApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationVersion**](ConfigurationApi.md#GetConfigurationVersion) | **Get** /services/haproxy/configuration/version | Return a configuration version
[**GetHAProxyConfiguration**](ConfigurationApi.md#GetHAProxyConfiguration) | **Get** /services/haproxy/configuration/raw | Return HAProxy configuration
[**PostHAProxyConfiguration**](ConfigurationApi.md#PostHAProxyConfiguration) | **Post** /services/haproxy/configuration/raw | Push new haproxy configuration


# **GetConfigurationVersion**
> int32 GetConfigurationVersion(ctx, optional)
Return a configuration version

Returns configuration version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationApiGetConfigurationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationApiGetConfigurationVersionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

**int32**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHAProxyConfiguration**
> string GetHAProxyConfiguration(ctx, optional)
Return HAProxy configuration

Returns HAProxy configuration file in plain text

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationApiGetHAProxyConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationApiGetHAProxyConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

**string**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostHAProxyConfiguration**
> string PostHAProxyConfiguration(ctx, data, optional)
Push new haproxy configuration

Push a new haproxy configuration file in plain text

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 
 **optional** | ***ConfigurationApiPostHAProxyConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationApiPostHAProxyConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipVersion** | **optional.Bool**| If set, no version check will be done and the pushed config will be enforced | [default to false]
 **skipReload** | **optional.Bool**| If set, no reload will be initiated and runtime actions from X-Runtime-Actions will be applied | [default to false]
 **onlyValidate** | **optional.Bool**| If set, only validates configuration, without applying it | [default to false]
 **xRuntimeActions** | **optional.String**| List of Runtime API commands with parameters separated by &#39;;&#39; | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

**string**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

