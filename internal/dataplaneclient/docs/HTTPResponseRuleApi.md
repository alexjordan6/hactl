# \HTTPResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#CreateHTTPResponseRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/http_response_rules/{index} | Add a new HTTP Response Rule
[**CreateHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#CreateHTTPResponseRuleFrontend) | **Post** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules/{index} | Add a new HTTP Response Rule
[**DeleteHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#DeleteHTTPResponseRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/http_response_rules/{index} | Delete a HTTP Response Rule
[**DeleteHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#DeleteHTTPResponseRuleFrontend) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules/{index} | Delete a HTTP Response Rule
[**GetAllHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#GetAllHTTPResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_response_rules | Return an array of all HTTP Response Rules
[**GetAllHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#GetAllHTTPResponseRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules | Return an array of all HTTP Response Rules
[**GetHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#GetHTTPResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_response_rules/{index} | Return one HTTP Response Rule
[**GetHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#GetHTTPResponseRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules/{index} | Return one HTTP Response Rule
[**ReplaceAllHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#ReplaceAllHTTPResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_response_rules | Replace an HTTP Response Rule list
[**ReplaceAllHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#ReplaceAllHTTPResponseRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules | Replace an HTTP Response Rule list
[**ReplaceHTTPResponseRuleBackend**](HTTPResponseRuleApi.md#ReplaceHTTPResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_response_rules/{index} | Replace a HTTP Response Rule
[**ReplaceHTTPResponseRuleFrontend**](HTTPResponseRuleApi.md#ReplaceHTTPResponseRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_response_rules/{index} | Replace a HTTP Response Rule


# **CreateHTTPResponseRuleBackend**
> HttpResponseRule CreateHTTPResponseRuleBackend(ctx, index, parentName, data, optional)
Add a new HTTP Response Rule

Adds a new HTTP Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiCreateHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiCreateHTTPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPResponseRuleFrontend**
> HttpResponseRule CreateHTTPResponseRuleFrontend(ctx, index, parentName, data, optional)
Add a new HTTP Response Rule

Adds a new HTTP Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiCreateHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiCreateHTTPResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPResponseRuleBackend**
> DeleteHTTPResponseRuleBackend(ctx, index, parentName, optional)
Delete a HTTP Response Rule

Deletes a HTTP Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiDeleteHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiDeleteHTTPResponseRuleBackendOpts struct

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

# **DeleteHTTPResponseRuleFrontend**
> DeleteHTTPResponseRuleFrontend(ctx, index, parentName, optional)
Delete a HTTP Response Rule

Deletes a HTTP Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiDeleteHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiDeleteHTTPResponseRuleFrontendOpts struct

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

# **GetAllHTTPResponseRuleBackend**
> HttpResponseRules GetAllHTTPResponseRuleBackend(ctx, parentName, optional)
Return an array of all HTTP Response Rules

Returns all HTTP Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiGetAllHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetAllHTTPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRules**](http_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPResponseRuleFrontend**
> HttpResponseRules GetAllHTTPResponseRuleFrontend(ctx, parentName, optional)
Return an array of all HTTP Response Rules

Returns all HTTP Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiGetAllHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetAllHTTPResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRules**](http_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPResponseRuleBackend**
> HttpResponseRule GetHTTPResponseRuleBackend(ctx, index, parentName, optional)
Return one HTTP Response Rule

Returns one HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiGetHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetHTTPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPResponseRuleFrontend**
> HttpResponseRule GetHTTPResponseRuleFrontend(ctx, index, parentName, optional)
Return one HTTP Response Rule

Returns one HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPResponseRuleApiGetHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiGetHTTPResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPResponseRuleBackend**
> HttpResponseRules ReplaceAllHTTPResponseRuleBackend(ctx, parentName, data, optional)
Replace an HTTP Response Rule list

Replaces a whole list of HTTP Response Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRules**](HttpResponseRules.md)|  | 
 **optional** | ***HTTPResponseRuleApiReplaceAllHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiReplaceAllHTTPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRules**](http_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPResponseRuleFrontend**
> HttpResponseRules ReplaceAllHTTPResponseRuleFrontend(ctx, parentName, data, optional)
Replace an HTTP Response Rule list

Replaces a whole list of HTTP Response Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRules**](HttpResponseRules.md)|  | 
 **optional** | ***HTTPResponseRuleApiReplaceAllHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiReplaceAllHTTPResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRules**](http_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPResponseRuleBackend**
> HttpResponseRule ReplaceHTTPResponseRuleBackend(ctx, index, parentName, data, optional)
Replace a HTTP Response Rule

Replaces a HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiReplaceHTTPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiReplaceHTTPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPResponseRuleFrontend**
> HttpResponseRule ReplaceHTTPResponseRuleFrontend(ctx, index, parentName, data, optional)
Replace a HTTP Response Rule

Replaces a HTTP Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpResponseRule**](HttpResponseRule.md)|  | 
 **optional** | ***HTTPResponseRuleApiReplaceHTTPResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPResponseRuleApiReplaceHTTPResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpResponseRule**](http_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

