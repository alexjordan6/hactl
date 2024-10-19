# \HTTPErrorsApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPErrorsSection**](HTTPErrorsApi.md#CreateHTTPErrorsSection) | **Post** /services/haproxy/configuration/http_errors_sections | Add a new http-error section
[**DeleteHTTPErrorsSection**](HTTPErrorsApi.md#DeleteHTTPErrorsSection) | **Delete** /services/haproxy/configuration/http_errors_sections/{name} | Delete a http-error section
[**GetHTTPErrorsSection**](HTTPErrorsApi.md#GetHTTPErrorsSection) | **Get** /services/haproxy/configuration/http_errors_sections/{name} | Return a http-error section
[**GetHTTPErrorsSections**](HTTPErrorsApi.md#GetHTTPErrorsSections) | **Get** /services/haproxy/configuration/http_errors_sections | Return an array of http-error sections
[**ReplaceHTTPErrorsSection**](HTTPErrorsApi.md#ReplaceHTTPErrorsSection) | **Put** /services/haproxy/configuration/http_errors_sections/{name} | Replace a http-error section


# **CreateHTTPErrorsSection**
> HttpErrorsSection CreateHTTPErrorsSection(ctx, data, optional)
Add a new http-error section

Adds a new http-error section to the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**HttpErrorsSection**](HttpErrorsSection.md)|  | 
 **optional** | ***HTTPErrorsApiCreateHTTPErrorsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorsApiCreateHTTPErrorsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorsSection**](http_errors_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPErrorsSection**
> DeleteHTTPErrorsSection(ctx, name, optional)
Delete a http-error section

Deletes a http-error section with a given name from the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| http-error section name | 
 **optional** | ***HTTPErrorsApiDeleteHTTPErrorsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorsApiDeleteHTTPErrorsSectionOpts struct

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

# **GetHTTPErrorsSection**
> HttpErrorsSection GetHTTPErrorsSection(ctx, name, optional)
Return a http-error section

Returns one http-error section with a given name from the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| http-error section name | 
 **optional** | ***HTTPErrorsApiGetHTTPErrorsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorsApiGetHTTPErrorsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorsSection**](http_errors_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPErrorsSections**
> HttpErrorsSections GetHTTPErrorsSections(ctx, optional)
Return an array of http-error sections

Returns an array of all configured http-error sections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HTTPErrorsApiGetHTTPErrorsSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorsApiGetHTTPErrorsSectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorsSections**](http_errors_sections.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPErrorsSection**
> HttpErrorsSection ReplaceHTTPErrorsSection(ctx, name, data, optional)
Replace a http-error section

Replaces a http-error section with a given name in the configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| http-error section name | 
  **data** | [**HttpErrorsSection**](HttpErrorsSection.md)|  | 
 **optional** | ***HTTPErrorsApiReplaceHTTPErrorsSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorsApiReplaceHTTPErrorsSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorsSection**](http_errors_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

