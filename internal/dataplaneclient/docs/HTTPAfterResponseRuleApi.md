# \HTTPAfterResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#CreateHTTPAfterResponseRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules/{index} | Add a new HTTP After Response Rule
[**CreateHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#CreateHTTPAfterResponseRuleFrontend) | **Post** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules/{index} | Add a new HTTP After Response Rule
[**DeleteHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#DeleteHTTPAfterResponseRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules/{index} | Delete a HTTP After Response Rule
[**DeleteHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#DeleteHTTPAfterResponseRuleFrontend) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules/{index} | Delete a HTTP After Response Rule
[**GetAllHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#GetAllHTTPAfterResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules | Return an array of all HTTP After Response Rules
[**GetAllHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#GetAllHTTPAfterResponseRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules | Return an array of all HTTP After Response Rules
[**GetHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#GetHTTPAfterResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules/{index} | Return one HTTP After Response Rule
[**GetHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#GetHTTPAfterResponseRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules/{index} | Return one HTTP After Response Rule
[**ReplaceAllHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#ReplaceAllHTTPAfterResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules | Replace an HTTP After Response Rules list
[**ReplaceAllHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#ReplaceAllHTTPAfterResponseRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules | Replace an HTTP After Response Rules list
[**ReplaceHTTPAfterResponseRuleBackend**](HTTPAfterResponseRuleApi.md#ReplaceHTTPAfterResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_after_response_rules/{index} | Replace a HTTP After Response Rule
[**ReplaceHTTPAfterResponseRuleFrontend**](HTTPAfterResponseRuleApi.md#ReplaceHTTPAfterResponseRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_after_response_rules/{index} | Replace a HTTP After Response Rule


# **CreateHTTPAfterResponseRuleBackend**
> HttpAfterResponseRule CreateHTTPAfterResponseRuleBackend(ctx, index, parentName, data, optional)
Add a new HTTP After Response Rule

Adds a new HTTP After Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPAfterResponseRuleFrontend**
> HttpAfterResponseRule CreateHTTPAfterResponseRuleFrontend(ctx, index, parentName, data, optional)
Add a new HTTP After Response Rule

Adds a new HTTP After Response Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiCreateHTTPAfterResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPAfterResponseRuleBackend**
> DeleteHTTPAfterResponseRuleBackend(ctx, index, parentName, optional)
Delete a HTTP After Response Rule

Deletes a HTTP After Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleBackendOpts struct

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

# **DeleteHTTPAfterResponseRuleFrontend**
> DeleteHTTPAfterResponseRuleFrontend(ctx, index, parentName, optional)
Delete a HTTP After Response Rule

Deletes a HTTP After Response Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiDeleteHTTPAfterResponseRuleFrontendOpts struct

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

# **GetAllHTTPAfterResponseRuleBackend**
> HttpAfterResponseRules GetAllHTTPAfterResponseRuleBackend(ctx, parentName, optional)
Return an array of all HTTP After Response Rules

Returns all HTTP After Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiGetAllHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetAllHTTPAfterResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRules**](http_after_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPAfterResponseRuleFrontend**
> HttpAfterResponseRules GetAllHTTPAfterResponseRuleFrontend(ctx, parentName, optional)
Return an array of all HTTP After Response Rules

Returns all HTTP After Response Rules that are configured in specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiGetAllHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetAllHTTPAfterResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRules**](http_after_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPAfterResponseRuleBackend**
> HttpAfterResponseRule GetHTTPAfterResponseRuleBackend(ctx, index, parentName, optional)
Return one HTTP After Response Rule

Returns one HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPAfterResponseRuleFrontend**
> HttpAfterResponseRule GetHTTPAfterResponseRuleFrontend(ctx, index, parentName, optional)
Return one HTTP After Response Rule

Returns one HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiGetHTTPAfterResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPAfterResponseRuleBackend**
> HttpAfterResponseRules ReplaceAllHTTPAfterResponseRuleBackend(ctx, parentName, data, optional)
Replace an HTTP After Response Rules list

Replaces a whole list of HTTP After Response Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRules**](HttpAfterResponseRules.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiReplaceAllHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiReplaceAllHTTPAfterResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRules**](http_after_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPAfterResponseRuleFrontend**
> HttpAfterResponseRules ReplaceAllHTTPAfterResponseRuleFrontend(ctx, parentName, data, optional)
Replace an HTTP After Response Rules list

Replaces a whole list of HTTP After Response Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRules**](HttpAfterResponseRules.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiReplaceAllHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiReplaceAllHTTPAfterResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRules**](http_after_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPAfterResponseRuleBackend**
> HttpAfterResponseRule ReplaceHTTPAfterResponseRuleBackend(ctx, index, parentName, data, optional)
Replace a HTTP After Response Rule

Replaces a HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPAfterResponseRuleFrontend**
> HttpAfterResponseRule ReplaceHTTPAfterResponseRuleFrontend(ctx, index, parentName, data, optional)
Replace a HTTP After Response Rule

Replaces a HTTP After Response Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP After Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpAfterResponseRule**](HttpAfterResponseRule.md)|  | 
 **optional** | ***HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPAfterResponseRuleApiReplaceHTTPAfterResponseRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpAfterResponseRule**](http_after_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

