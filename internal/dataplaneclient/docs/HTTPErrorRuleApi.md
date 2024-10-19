# \HTTPErrorRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#CreateHTTPErrorRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/http_error_rules/{index} | Add a new HTTP Error Rule
[**CreateHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#CreateHTTPErrorRuleDefaults) | **Post** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules/{index} | Add a new HTTP Error Rule
[**CreateHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#CreateHTTPErrorRuleFrontend) | **Post** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules/{index} | Add a new HTTP Error Rule
[**DeleteHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#DeleteHTTPErrorRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/http_error_rules/{index} | Delete a HTTP Error Rule
[**DeleteHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#DeleteHTTPErrorRuleDefaults) | **Delete** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules/{index} | Delete a HTTP Error Rule
[**DeleteHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#DeleteHTTPErrorRuleFrontend) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules/{index} | Delete a HTTP Error Rule
[**GetAllHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#GetAllHTTPErrorRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_error_rules | Return an array of all HTTP Error Rules
[**GetAllHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#GetAllHTTPErrorRuleDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules | Return an array of all HTTP Error Rules
[**GetAllHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#GetAllHTTPErrorRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules | Return an array of all HTTP Error Rules
[**GetHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#GetHTTPErrorRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/http_error_rules/{index} | Return one HTTP Error Rule
[**GetHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#GetHTTPErrorRuleDefaults) | **Get** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules/{index} | Return one HTTP Error Rule
[**GetHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#GetHTTPErrorRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules/{index} | Return one HTTP Error Rule
[**ReplaceAllHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#ReplaceAllHTTPErrorRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_error_rules | Replace an HTTP Error Rules list
[**ReplaceAllHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#ReplaceAllHTTPErrorRuleDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules | Replace an HTTP Error Rules list
[**ReplaceAllHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#ReplaceAllHTTPErrorRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules | Replace an HTTP Error Rules list
[**ReplaceHTTPErrorRuleBackend**](HTTPErrorRuleApi.md#ReplaceHTTPErrorRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/http_error_rules/{index} | Replace a HTTP Error Rule
[**ReplaceHTTPErrorRuleDefaults**](HTTPErrorRuleApi.md#ReplaceHTTPErrorRuleDefaults) | **Put** /services/haproxy/configuration/defaults/{parent_name}/http_error_rules/{index} | Replace a HTTP Error Rule
[**ReplaceHTTPErrorRuleFrontend**](HTTPErrorRuleApi.md#ReplaceHTTPErrorRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/http_error_rules/{index} | Replace a HTTP Error Rule


# **CreateHTTPErrorRuleBackend**
> HttpErrorRule CreateHTTPErrorRuleBackend(ctx, index, parentName, data, optional)
Add a new HTTP Error Rule

Adds a new HTTP Error Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiCreateHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiCreateHTTPErrorRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPErrorRuleDefaults**
> HttpErrorRule CreateHTTPErrorRuleDefaults(ctx, index, parentName, data, optional)
Add a new HTTP Error Rule

Adds a new HTTP Error Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiCreateHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiCreateHTTPErrorRuleDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHTTPErrorRuleFrontend**
> HttpErrorRule CreateHTTPErrorRuleFrontend(ctx, index, parentName, data, optional)
Add a new HTTP Error Rule

Adds a new HTTP Error Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiCreateHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiCreateHTTPErrorRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHTTPErrorRuleBackend**
> DeleteHTTPErrorRuleBackend(ctx, index, parentName, optional)
Delete a HTTP Error Rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiDeleteHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiDeleteHTTPErrorRuleBackendOpts struct

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

# **DeleteHTTPErrorRuleDefaults**
> DeleteHTTPErrorRuleDefaults(ctx, index, parentName, optional)
Delete a HTTP Error Rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiDeleteHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiDeleteHTTPErrorRuleDefaultsOpts struct

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

# **DeleteHTTPErrorRuleFrontend**
> DeleteHTTPErrorRuleFrontend(ctx, index, parentName, optional)
Delete a HTTP Error Rule

Deletes a HTTP Error Rule configuration by its index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiDeleteHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiDeleteHTTPErrorRuleFrontendOpts struct

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

# **GetAllHTTPErrorRuleBackend**
> HttpErrorRules GetAllHTTPErrorRuleBackend(ctx, parentName, optional)
Return an array of all HTTP Error Rules

Returns all HTTP Error Rules that are configured in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetAllHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetAllHTTPErrorRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPErrorRuleDefaults**
> HttpErrorRules GetAllHTTPErrorRuleDefaults(ctx, parentName, optional)
Return an array of all HTTP Error Rules

Returns all HTTP Error Rules that are configured in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetAllHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetAllHTTPErrorRuleDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHTTPErrorRuleFrontend**
> HttpErrorRules GetAllHTTPErrorRuleFrontend(ctx, parentName, optional)
Return an array of all HTTP Error Rules

Returns all HTTP Error Rules that are configured in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetAllHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetAllHTTPErrorRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPErrorRuleBackend**
> HttpErrorRule GetHTTPErrorRuleBackend(ctx, index, parentName, optional)
Return one HTTP Error Rule

Returns one HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetHTTPErrorRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPErrorRuleDefaults**
> HttpErrorRule GetHTTPErrorRuleDefaults(ctx, index, parentName, optional)
Return one HTTP Error Rule

Returns one HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetHTTPErrorRuleDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHTTPErrorRuleFrontend**
> HttpErrorRule GetHTTPErrorRuleFrontend(ctx, index, parentName, optional)
Return one HTTP Error Rule

Returns one HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***HTTPErrorRuleApiGetHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiGetHTTPErrorRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPErrorRuleBackend**
> HttpErrorRules ReplaceAllHTTPErrorRuleBackend(ctx, parentName, data, optional)
Replace an HTTP Error Rules list

Replaces a whole list of HTTP Error Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRules**](HttpErrorRules.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceAllHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceAllHTTPErrorRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPErrorRuleDefaults**
> HttpErrorRules ReplaceAllHTTPErrorRuleDefaults(ctx, parentName, data, optional)
Replace an HTTP Error Rules list

Replaces a whole list of HTTP Error Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRules**](HttpErrorRules.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceAllHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceAllHTTPErrorRuleDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllHTTPErrorRuleFrontend**
> HttpErrorRules ReplaceAllHTTPErrorRuleFrontend(ctx, parentName, data, optional)
Replace an HTTP Error Rules list

Replaces a whole list of HTTP Error Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRules**](HttpErrorRules.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceAllHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceAllHTTPErrorRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRules**](http_error_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPErrorRuleBackend**
> HttpErrorRule ReplaceHTTPErrorRuleBackend(ctx, index, parentName, data, optional)
Replace a HTTP Error Rule

Replaces a HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceHTTPErrorRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceHTTPErrorRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPErrorRuleDefaults**
> HttpErrorRule ReplaceHTTPErrorRuleDefaults(ctx, index, parentName, data, optional)
Replace a HTTP Error Rule

Replaces a HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceHTTPErrorRuleDefaultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceHTTPErrorRuleDefaultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceHTTPErrorRuleFrontend**
> HttpErrorRule ReplaceHTTPErrorRuleFrontend(ctx, index, parentName, data, optional)
Replace a HTTP Error Rule

Replaces a HTTP Error Rule configuration by its index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| HTTP Error Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**HttpErrorRule**](HttpErrorRule.md)|  | 
 **optional** | ***HTTPErrorRuleApiReplaceHTTPErrorRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HTTPErrorRuleApiReplaceHTTPErrorRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**HttpErrorRule**](http_error_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

