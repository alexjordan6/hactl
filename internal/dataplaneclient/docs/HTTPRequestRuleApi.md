# \HTTPRequestRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#CreateHTTPRequestRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/http_request_rules/{index} | Add a new HTTP Request Rule
[**CreateHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#CreateHTTPRequestRuleFrontend) | **Post** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules/{index} | Add a new HTTP Request Rule
[**DeleteHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#DeleteHTTPRequestRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/http_request_rules/{index} | Delete a HTTP Request Rule
[**DeleteHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#DeleteHTTPRequestRuleFrontend) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules/{index} | Delete a HTTP Request Rule
[**GetAllHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#GetAllHTTPRequestRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_request_rules | Return an array of all HTTP Request Rules
[**GetAllHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#GetAllHTTPRequestRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules | Return an array of all HTTP Request Rules
[**GetHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#GetHTTPRequestRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_request_rules/{index} | Return one HTTP Request Rule
[**GetHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#GetHTTPRequestRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules/{index} | Return one HTTP Request Rule
[**ReplaceAllHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#ReplaceAllHTTPRequestRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_request_rules | Replace an HTTP Request Rule list
[**ReplaceAllHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#ReplaceAllHTTPRequestRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules | Replace an HTTP Request Rule list
[**ReplaceHTTPRequestRuleBackend**](HTTPRequestRuleApi.md#ReplaceHTTPRequestRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_request_rules/{index} | Replace a HTTP Request Rule
[**ReplaceHTTPRequestRuleFrontend**](HTTPRequestRuleApi.md#ReplaceHTTPRequestRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_request_rules/{index} | Replace a HTTP Request Rule


# **CreateHTTPRequestRuleBackend**
> HttpRequestRule CreateHTTPRequestRuleBackend(ctx, index, parentName, data, optional)
Add a new HTTP Request Rule

Adds a new HTTP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiCreateHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiCreateHTTPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPRequestRuleFrontend**
> HttpRequestRule CreateHTTPRequestRuleFrontend(ctx, index, parentName, data, optional)
Add a new HTTP Request Rule

Adds a new HTTP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiCreateHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiCreateHTTPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPRequestRuleBackend**
> DeleteHTTPRequestRuleBackend(ctx, index, parentName, optional)
Delete a HTTP Request Rule

Deletes a HTTP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiDeleteHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiDeleteHTTPRequestRuleBackendOpts struct

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

# **DeleteHTTPRequestRuleFrontend**
> DeleteHTTPRequestRuleFrontend(ctx, index, parentName, optional)
Delete a HTTP Request Rule

Deletes a HTTP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiDeleteHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiDeleteHTTPRequestRuleFrontendOpts struct

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

# **GetAllHTTPRequestRuleBackend**
> HttpRequestRules GetAllHTTPRequestRuleBackend(ctx, parentName, optional)
Return an array of all HTTP Request Rules

Returns all HTTP Request Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiGetAllHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetAllHTTPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRules**](http_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPRequestRuleFrontend**
> HttpRequestRules GetAllHTTPRequestRuleFrontend(ctx, parentName, optional)
Return an array of all HTTP Request Rules

Returns all HTTP Request Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiGetAllHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetAllHTTPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRules**](http_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPRequestRuleBackend**
> HttpRequestRule GetHTTPRequestRuleBackend(ctx, index, parentName, optional)
Return one HTTP Request Rule

Returns one HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiGetHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetHTTPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPRequestRuleFrontend**
> HttpRequestRule GetHTTPRequestRuleFrontend(ctx, index, parentName, optional)
Return one HTTP Request Rule

Returns one HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPRequestRuleApiGetHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiGetHTTPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPRequestRuleBackend**
> HttpRequestRules ReplaceAllHTTPRequestRuleBackend(ctx, parentName, data, optional)
Replace an HTTP Request Rule list

Replaces a whole list of HTTP Request Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRules**](HttpRequestRules.md)|  | 
 **optional** | ***HTTPRequestRuleApiReplaceAllHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiReplaceAllHTTPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRules**](http_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPRequestRuleFrontend**
> HttpRequestRules ReplaceAllHTTPRequestRuleFrontend(ctx, parentName, data, optional)
Replace an HTTP Request Rule list

Replaces a whole list of HTTP Request Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRules**](HttpRequestRules.md)|  | 
 **optional** | ***HTTPRequestRuleApiReplaceAllHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiReplaceAllHTTPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRules**](http_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPRequestRuleBackend**
> HttpRequestRule ReplaceHTTPRequestRuleBackend(ctx, index, parentName, data, optional)
Replace a HTTP Request Rule

Replaces a HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiReplaceHTTPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiReplaceHTTPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPRequestRuleFrontend**
> HttpRequestRule ReplaceHTTPRequestRuleFrontend(ctx, index, parentName, data, optional)
Replace a HTTP Request Rule

Replaces a HTTP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpRequestRule**](HttpRequestRule.md)|  | 
 **optional** | ***HTTPRequestRuleApiReplaceHTTPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPRequestRuleApiReplaceHTTPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpRequestRule**](http_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

