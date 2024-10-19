# \TCPResponseRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPResponseRuleBackend**](TCPResponseRuleApi.md#CreateTCPResponseRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules/{index} | Add a new TCP Response Rule
[**DeleteTCPResponseRuleBackend**](TCPResponseRuleApi.md#DeleteTCPResponseRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules/{index} | Delete a TCP Response Rule
[**GetAllTCPResponseRuleBackend**](TCPResponseRuleApi.md#GetAllTCPResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules | Return an array of all TCP Response Rules
[**GetTCPResponseRuleBackend**](TCPResponseRuleApi.md#GetTCPResponseRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules/{index} | Return one TCP Response Rule
[**ReplaceAllTCPResponseRuleBackend**](TCPResponseRuleApi.md#ReplaceAllTCPResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules | Replace a TCP Response Rule list
[**ReplaceTCPResponseRuleBackend**](TCPResponseRuleApi.md#ReplaceTCPResponseRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_response_rules/{index} | Replace a TCP Response Rule


# **CreateTCPResponseRuleBackend**
> TcpResponseRule CreateTCPResponseRuleBackend(ctx, index, parentName, data, optional)
Add a new TCP Response Rule

Adds a new TCP Response Rule of the specified type in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpResponseRule**](TcpResponseRule.md)|  | 
 **optional** | ***TCPResponseRuleApiCreateTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiCreateTCPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPResponseRuleBackend**
> DeleteTCPResponseRuleBackend(ctx, index, parentName, optional)
Delete a TCP Response Rule

Deletes a TCP Response Rule configuration by it's index from the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPResponseRuleApiDeleteTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiDeleteTCPResponseRuleBackendOpts struct

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

# **GetAllTCPResponseRuleBackend**
> TcpResponseRules GetAllTCPResponseRuleBackend(ctx, parentName, optional)
Return an array of all TCP Response Rules

Returns all TCP Response Rules that are configured in specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPResponseRuleApiGetAllTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiGetAllTCPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpResponseRules**](tcp_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPResponseRuleBackend**
> TcpResponseRule GetTCPResponseRuleBackend(ctx, index, parentName, optional)
Return one TCP Response Rule

Returns one TCP Response Rule configuration by it's index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPResponseRuleApiGetTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiGetTCPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTCPResponseRuleBackend**
> TcpResponseRules ReplaceAllTCPResponseRuleBackend(ctx, parentName, data, optional)
Replace a TCP Response Rule list

Replaces a whole list of TCP Response Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**TcpResponseRules**](TcpResponseRules.md)|  | 
 **optional** | ***TCPResponseRuleApiReplaceAllTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiReplaceAllTCPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpResponseRules**](tcp_response_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPResponseRuleBackend**
> TcpResponseRule ReplaceTCPResponseRuleBackend(ctx, index, parentName, data, optional)
Replace a TCP Response Rule

Replaces a TCP Response Rule configuration by it's Index in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Response Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpResponseRule**](TcpResponseRule.md)|  | 
 **optional** | ***TCPResponseRuleApiReplaceTCPResponseRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPResponseRuleApiReplaceTCPResponseRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpResponseRule**](tcp_response_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

