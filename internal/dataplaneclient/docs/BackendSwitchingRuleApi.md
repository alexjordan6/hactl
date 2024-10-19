# \BackendSwitchingRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackendSwitchingRule**](BackendSwitchingRuleApi.md#CreateBackendSwitchingRule) | **Post** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules/{index} | Add a new Backend Switching Rule
[**DeleteBackendSwitchingRule**](BackendSwitchingRuleApi.md#DeleteBackendSwitchingRule) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules/{index} | Delete a Backend Switching Rule
[**GetBackendSwitchingRule**](BackendSwitchingRuleApi.md#GetBackendSwitchingRule) | **Get** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules/{index} | Return one Backend Switching Rule
[**GetBackendSwitchingRules**](BackendSwitchingRuleApi.md#GetBackendSwitchingRules) | **Get** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules | Return an array of all Backend Switching Rules
[**ReplaceBackendSwitchingRule**](BackendSwitchingRuleApi.md#ReplaceBackendSwitchingRule) | **Put** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules/{index} | Replace a Backend Switching Rule
[**ReplaceBackendSwitchingRules**](BackendSwitchingRuleApi.md#ReplaceBackendSwitchingRules) | **Put** /services/haproxy/configuration/frontends/{parent_name}/backend_switching_rules | Replace an Backend Switching Rule list


# **CreateBackendSwitchingRule**
> BackendSwitchingRule CreateBackendSwitchingRule(ctx, index, parentName, data, optional)
Add a new Backend Switching Rule

Adds a new Backend Switching Rule of the specified type in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Switching Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**BackendSwitchingRule**](BackendSwitchingRule.md)|  | 
 **optional** | ***BackendSwitchingRuleApiCreateBackendSwitchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiCreateBackendSwitchingRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**BackendSwitchingRule**](backend_switching_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBackendSwitchingRule**
> DeleteBackendSwitchingRule(ctx, index, parentName, optional)
Delete a Backend Switching Rule

Deletes a Backend Switching Rule configuration by it's index from the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Switching Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***BackendSwitchingRuleApiDeleteBackendSwitchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiDeleteBackendSwitchingRuleOpts struct

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

# **GetBackendSwitchingRule**
> BackendSwitchingRule GetBackendSwitchingRule(ctx, index, parentName, optional)
Return one Backend Switching Rule

Returns one Backend Switching Rule configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Switching Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***BackendSwitchingRuleApiGetBackendSwitchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiGetBackendSwitchingRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**BackendSwitchingRule**](backend_switching_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackendSwitchingRules**
> BackendSwitchingRules GetBackendSwitchingRules(ctx, parentName, optional)
Return an array of all Backend Switching Rules

Returns all Backend Switching Rules that are configured in specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***BackendSwitchingRuleApiGetBackendSwitchingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiGetBackendSwitchingRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**BackendSwitchingRules**](backend_switching_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBackendSwitchingRule**
> BackendSwitchingRule ReplaceBackendSwitchingRule(ctx, index, parentName, data, optional)
Replace a Backend Switching Rule

Replaces a Backend Switching Rule configuration by it's index in the specified frontend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Switching Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**BackendSwitchingRule**](BackendSwitchingRule.md)|  | 
 **optional** | ***BackendSwitchingRuleApiReplaceBackendSwitchingRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiReplaceBackendSwitchingRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**BackendSwitchingRule**](backend_switching_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceBackendSwitchingRules**
> BackendSwitchingRules ReplaceBackendSwitchingRules(ctx, parentName, data, optional)
Replace an Backend Switching Rule list

Replaces a whole list of Backend Switching Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**BackendSwitchingRules**](BackendSwitchingRules.md)|  | 
 **optional** | ***BackendSwitchingRuleApiReplaceBackendSwitchingRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BackendSwitchingRuleApiReplaceBackendSwitchingRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**BackendSwitchingRules**](backend_switching_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

