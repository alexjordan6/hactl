# \StickRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStickRule**](StickRuleApi.md#CreateStickRule) | **Post** /services/haproxy/configuration/backends/{parent_name}/stick_rules/{index} | Add a new Stick Rule
[**DeleteStickRule**](StickRuleApi.md#DeleteStickRule) | **Delete** /services/haproxy/configuration/backends/{parent_name}/stick_rules/{index} | Delete a Stick Rule
[**GetStickRule**](StickRuleApi.md#GetStickRule) | **Get** /services/haproxy/configuration/backends/{parent_name}/stick_rules/{index} | Return one Stick Rule
[**GetStickRules**](StickRuleApi.md#GetStickRules) | **Get** /services/haproxy/configuration/backends/{parent_name}/stick_rules | Return an array of all Stick Rules
[**ReplaceStickRule**](StickRuleApi.md#ReplaceStickRule) | **Put** /services/haproxy/configuration/backends/{parent_name}/stick_rules/{index} | Replace a Stick Rule
[**ReplaceStickRules**](StickRuleApi.md#ReplaceStickRules) | **Put** /services/haproxy/configuration/backends/{parent_name}/stick_rules | Replace a Stick Rule list


# **CreateStickRule**
> StickRule CreateStickRule(ctx, index, parentName, data, optional)
Add a new Stick Rule

Adds a new Stick Rule of the specified type in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Stick Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**StickRule**](StickRule.md)|  | 
 **optional** | ***StickRuleApiCreateStickRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiCreateStickRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**StickRule**](stick_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStickRule**
> DeleteStickRule(ctx, index, parentName, optional)
Delete a Stick Rule

Deletes a Stick Rule configuration by it's index from the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Stick Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***StickRuleApiDeleteStickRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiDeleteStickRuleOpts struct

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

# **GetStickRule**
> StickRule GetStickRule(ctx, index, parentName, optional)
Return one Stick Rule

Returns one Stick Rule configuration by it's index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Stick Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***StickRuleApiGetStickRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiGetStickRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**StickRule**](stick_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStickRules**
> StickRules GetStickRules(ctx, parentName, optional)
Return an array of all Stick Rules

Returns all Stick Rules that are configured in specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***StickRuleApiGetStickRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiGetStickRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**StickRules**](stick_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStickRule**
> StickRule ReplaceStickRule(ctx, index, parentName, data, optional)
Replace a Stick Rule

Replaces a Stick Rule configuration by it's index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| Stick Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**StickRule**](StickRule.md)|  | 
 **optional** | ***StickRuleApiReplaceStickRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiReplaceStickRuleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**StickRule**](stick_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStickRules**
> StickRules ReplaceStickRules(ctx, parentName, data, optional)
Replace a Stick Rule list

Replaces a whole list of Stick Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**StickRules**](StickRules.md)|  | 
 **optional** | ***StickRuleApiReplaceStickRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickRuleApiReplaceStickRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**StickRules**](stick_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

