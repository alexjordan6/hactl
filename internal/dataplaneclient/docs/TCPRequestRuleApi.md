# \TCPRequestRuleApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTCPRequestRuleBackend**](TCPRequestRuleApi.md#CreateTCPRequestRuleBackend) | **Post** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules/{index} | Add a new TCP Request Rule
[**CreateTCPRequestRuleFrontend**](TCPRequestRuleApi.md#CreateTCPRequestRuleFrontend) | **Post** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules/{index} | Add a new TCP Request Rule
[**DeleteTCPRequestRuleBackend**](TCPRequestRuleApi.md#DeleteTCPRequestRuleBackend) | **Delete** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules/{index} | Delete a TCP Request Rule
[**DeleteTCPRequestRuleFrontend**](TCPRequestRuleApi.md#DeleteTCPRequestRuleFrontend) | **Delete** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules/{index} | Delete a TCP Request Rule
[**GetAllTCPRequestRuleBackend**](TCPRequestRuleApi.md#GetAllTCPRequestRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules | Return an array of all TCP Request Rules
[**GetAllTCPRequestRuleFrontend**](TCPRequestRuleApi.md#GetAllTCPRequestRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules | Return an array of all TCP Request Rules
[**GetTCPRequestRuleBackend**](TCPRequestRuleApi.md#GetTCPRequestRuleBackend) | **Get** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules/{index} | Return one TCP Request Rule
[**GetTCPRequestRuleFrontend**](TCPRequestRuleApi.md#GetTCPRequestRuleFrontend) | **Get** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules/{index} | Return one TCP Request Rule
[**ReplaceAllTCPRequestRuleBackend**](TCPRequestRuleApi.md#ReplaceAllTCPRequestRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules | Replace an TCP Request Rule list
[**ReplaceAllTCPRequestRuleFrontend**](TCPRequestRuleApi.md#ReplaceAllTCPRequestRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules | Replace an TCP Request Rule list
[**ReplaceTCPRequestRuleBackend**](TCPRequestRuleApi.md#ReplaceTCPRequestRuleBackend) | **Put** /services/haproxy/configuration/backends/{parent_name}/tcp_request_rules/{index} | Replace a TCP Request Rule
[**ReplaceTCPRequestRuleFrontend**](TCPRequestRuleApi.md#ReplaceTCPRequestRuleFrontend) | **Put** /services/haproxy/configuration/frontends/{parent_name}/tcp_request_rules/{index} | Replace a TCP Request Rule


# **CreateTCPRequestRuleBackend**
> TcpRequestRule CreateTCPRequestRuleBackend(ctx, index, parentName, data, optional)
Add a new TCP Request Rule

Adds a new TCP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiCreateTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiCreateTCPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTCPRequestRuleFrontend**
> TcpRequestRule CreateTCPRequestRuleFrontend(ctx, index, parentName, data, optional)
Add a new TCP Request Rule

Adds a new TCP Request Rule of the specified type in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiCreateTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiCreateTCPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTCPRequestRuleBackend**
> DeleteTCPRequestRuleBackend(ctx, index, parentName, optional)
Delete a TCP Request Rule

Deletes a TCP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiDeleteTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiDeleteTCPRequestRuleBackendOpts struct

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

# **DeleteTCPRequestRuleFrontend**
> DeleteTCPRequestRuleFrontend(ctx, index, parentName, optional)
Delete a TCP Request Rule

Deletes a TCP Request Rule configuration by it's index from the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiDeleteTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiDeleteTCPRequestRuleFrontendOpts struct

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

# **GetAllTCPRequestRuleBackend**
> TcpRequestRules GetAllTCPRequestRuleBackend(ctx, parentName, optional)
Return an array of all TCP Request Rules

Returns all TCP Request Rules that are configured in specified parent and parent type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiGetAllTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetAllTCPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRules**](tcp_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTCPRequestRuleFrontend**
> TcpRequestRules GetAllTCPRequestRuleFrontend(ctx, parentName, optional)
Return an array of all TCP Request Rules

Returns all TCP Request Rules that are configured in specified parent and parent type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiGetAllTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetAllTCPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRules**](tcp_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPRequestRuleBackend**
> TcpRequestRule GetTCPRequestRuleBackend(ctx, index, parentName, optional)
Return one TCP Request Rule

Returns one TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiGetTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetTCPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTCPRequestRuleFrontend**
> TcpRequestRule GetTCPRequestRuleFrontend(ctx, index, parentName, optional)
Return one TCP Request Rule

Returns one TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
 **optional** | ***TCPRequestRuleApiGetTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiGetTCPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTCPRequestRuleBackend**
> TcpRequestRules ReplaceAllTCPRequestRuleBackend(ctx, parentName, data, optional)
Replace an TCP Request Rule list

Replaces a whole list of TCP Request Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRules**](TcpRequestRules.md)|  | 
 **optional** | ***TCPRequestRuleApiReplaceAllTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiReplaceAllTCPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRules**](tcp_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAllTCPRequestRuleFrontend**
> TcpRequestRules ReplaceAllTCPRequestRuleFrontend(ctx, parentName, data, optional)
Replace an TCP Request Rule list

Replaces a whole list of TCP Request Rules with the list given in parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRules**](TcpRequestRules.md)|  | 
 **optional** | ***TCPRequestRuleApiReplaceAllTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiReplaceAllTCPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRules**](tcp_request_rules.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPRequestRuleBackend**
> TcpRequestRule ReplaceTCPRequestRuleBackend(ctx, index, parentName, data, optional)
Replace a TCP Request Rule

Replaces a TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiReplaceTCPRequestRuleBackendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiReplaceTCPRequestRuleBackendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceTCPRequestRuleFrontend**
> TcpRequestRule ReplaceTCPRequestRuleFrontend(ctx, index, parentName, data, optional)
Replace a TCP Request Rule

Replaces a TCP Request Rule configuration by it's index in the specified parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **index** | **int32**| TCP Request Rule Index | 
  **parentName** | **string**| Parent name | 
  **data** | [**TcpRequestRule**](TcpRequestRule.md)|  | 
 **optional** | ***TCPRequestRuleApiReplaceTCPRequestRuleFrontendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TCPRequestRuleApiReplaceTCPRequestRuleFrontendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**TcpRequestRule**](tcp_request_rule.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

