# \ServerTemplateApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServerTemplate**](ServerTemplateApi.md#CreateServerTemplate) | **Post** /services/haproxy/configuration/backends/{parent_name}/server_templates | Add a new server template
[**DeleteServerTemplate**](ServerTemplateApi.md#DeleteServerTemplate) | **Delete** /services/haproxy/configuration/backends/{parent_name}/server_templates/{prefix} | Delete a server template
[**GetServerTemplate**](ServerTemplateApi.md#GetServerTemplate) | **Get** /services/haproxy/configuration/backends/{parent_name}/server_templates/{prefix} | Return one server template
[**GetServerTemplates**](ServerTemplateApi.md#GetServerTemplates) | **Get** /services/haproxy/configuration/backends/{parent_name}/server_templates | Return an array of server templates
[**ReplaceServerTemplate**](ServerTemplateApi.md#ReplaceServerTemplate) | **Put** /services/haproxy/configuration/backends/{parent_name}/server_templates/{prefix} | Replace a server template


# **CreateServerTemplate**
> ServerTemplate CreateServerTemplate(ctx, parentName, data, optional)
Add a new server template

Adds a new server template in the specified backend in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**ServerTemplate**](ServerTemplate.md)|  | 
 **optional** | ***ServerTemplateApiCreateServerTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerTemplateApiCreateServerTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**ServerTemplate**](server_template.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServerTemplate**
> DeleteServerTemplate(ctx, prefix, parentName, optional)
Delete a server template

Deletes a server template configuration by it's prefix in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefix** | **string**| Server template prefix | 
  **parentName** | **string**| Parent name | 
 **optional** | ***ServerTemplateApiDeleteServerTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerTemplateApiDeleteServerTemplateOpts struct

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

# **GetServerTemplate**
> ServerTemplate GetServerTemplate(ctx, prefix, parentName, optional)
Return one server template

Returns one server template configuration by it's prefix in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefix** | **string**| Server template prefix | 
  **parentName** | **string**| Parent name | 
 **optional** | ***ServerTemplateApiGetServerTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerTemplateApiGetServerTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**ServerTemplate**](server_template.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerTemplates**
> ServerTemplates GetServerTemplates(ctx, parentName, optional)
Return an array of server templates

Returns an array of all server templates that are configured in specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***ServerTemplateApiGetServerTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerTemplateApiGetServerTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**ServerTemplates**](server_templates.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceServerTemplate**
> ServerTemplate ReplaceServerTemplate(ctx, prefix, parentName, data, optional)
Replace a server template

Replaces a server template configuration by it's prefix in the specified backend.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefix** | **string**| Server template prefix | 
  **parentName** | **string**| Parent name | 
  **data** | [**ServerTemplate**](ServerTemplate.md)|  | 
 **optional** | ***ServerTemplateApiReplaceServerTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServerTemplateApiReplaceServerTemplateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**ServerTemplate**](server_template.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

