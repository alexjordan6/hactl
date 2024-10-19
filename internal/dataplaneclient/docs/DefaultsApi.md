# \DefaultsApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefaultsSection**](DefaultsApi.md#CreateDefaultsSection) | **Put** /services/haproxy/configuration/defaults | Add a defaults section
[**DeleteDefaultsSection**](DefaultsApi.md#DeleteDefaultsSection) | **Delete** /services/haproxy/configuration/defaults/{name} | Delete a defaults section
[**GetDefaultsSection**](DefaultsApi.md#GetDefaultsSection) | **Get** /services/haproxy/configuration/defaults/{name} | Return a defaults section
[**GetDefaultsSections**](DefaultsApi.md#GetDefaultsSections) | **Get** /services/haproxy/configuration/defaults | Return an array of defaults
[**ReplaceDefaultsSection**](DefaultsApi.md#ReplaceDefaultsSection) | **Put** /services/haproxy/configuration/defaults/{name} | Replace a defatults section


# **CreateDefaultsSection**
> Defaults CreateDefaultsSection(ctx, data, optional)
Add a defaults section

Adds a new defaults section to the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Defaults**](Defaults.md)|  | 
 **optional** | ***DefaultsApiCreateDefaultsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultsApiCreateDefaultsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**Defaults**](defaults.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDefaultsSection**
> DeleteDefaultsSection(ctx, name, optional)
Delete a defaults section

Deletes a defaults section from the configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Defaults name | 
 **optional** | ***DefaultsApiDeleteDefaultsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultsApiDeleteDefaultsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultsSection**
> Defaults GetDefaultsSection(ctx, name, optional)
Return a defaults section

Returns one defaults section configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Defaults name | 
 **optional** | ***DefaultsApiGetDefaultsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultsApiGetDefaultsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**Defaults**](defaults.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefaultsSections**
> DefaultsSections GetDefaultsSections(ctx, optional)
Return an array of defaults

Returns an array of all configured defaults.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultsApiGetDefaultsSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultsApiGetDefaultsSectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**DefaultsSections**](defaults_sections.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceDefaultsSection**
> Defaults ReplaceDefaultsSection(ctx, name, data, optional)
Replace a defatults section

Replaces a defatults section configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Defaults name | 
  **data** | [**Defaults**](Defaults.md)|  | 
 **optional** | ***DefaultsApiReplaceDefaultsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultsApiReplaceDefaultsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**Defaults**](defaults.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

