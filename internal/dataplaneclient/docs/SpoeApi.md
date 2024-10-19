# \SpoeApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpoe**](SpoeApi.md#CreateSpoe) | **Post** /services/haproxy/spoe/spoe_files | Creates SPOE file with its entries
[**CreateSpoeAgent**](SpoeApi.md#CreateSpoeAgent) | **Post** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/agents | Add a new spoe agent to scope
[**CreateSpoeGroup**](SpoeApi.md#CreateSpoeGroup) | **Post** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/groups | Add a new SPOE groups
[**CreateSpoeMessage**](SpoeApi.md#CreateSpoeMessage) | **Post** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/messages | Add a new spoe message to scope
[**CreateSpoeScope**](SpoeApi.md#CreateSpoeScope) | **Post** /services/haproxy/spoe/spoe_files/{parent_name}/scopes | Add a new spoe scope
[**DeleteSpoeAgent**](SpoeApi.md#DeleteSpoeAgent) | **Delete** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/agents/{name} | Delete a SPOE agent
[**DeleteSpoeFile**](SpoeApi.md#DeleteSpoeFile) | **Delete** /services/haproxy/spoe/spoe_files/{name} | Delete SPOE file
[**DeleteSpoeGroup**](SpoeApi.md#DeleteSpoeGroup) | **Delete** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/groups/{name} | Delete a SPOE groups
[**DeleteSpoeMessage**](SpoeApi.md#DeleteSpoeMessage) | **Delete** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/messages/{name} | Delete a spoe message
[**DeleteSpoeScope**](SpoeApi.md#DeleteSpoeScope) | **Delete** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{name} | Delete a SPOE scope
[**GetAllSpoeAgent**](SpoeApi.md#GetAllSpoeAgent) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/agents | Return an array of spoe agents in one scope
[**GetAllSpoeFiles**](SpoeApi.md#GetAllSpoeFiles) | **Get** /services/haproxy/spoe/spoe_files | Return all available SPOE files
[**GetAllSpoeGroup**](SpoeApi.md#GetAllSpoeGroup) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/groups | Return an array of SPOE groups
[**GetAllSpoeMessage**](SpoeApi.md#GetAllSpoeMessage) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/messages | Return an array of spoe messages in one scope
[**GetAllSpoeScope**](SpoeApi.md#GetAllSpoeScope) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes | Return an array of spoe scopes
[**GetOneSpoeFile**](SpoeApi.md#GetOneSpoeFile) | **Get** /services/haproxy/spoe/spoe_files/{name} | Return one SPOE file
[**GetSpoeAgent**](SpoeApi.md#GetSpoeAgent) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/agents/{name} | Return a spoe agent
[**GetSpoeConfigurationVersion**](SpoeApi.md#GetSpoeConfigurationVersion) | **Get** /services/haproxy/spoe/{parent_name}/version | Return a SPOE configuration version
[**GetSpoeGroup**](SpoeApi.md#GetSpoeGroup) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/groups/{name} | Return a SPOE groups
[**GetSpoeMessage**](SpoeApi.md#GetSpoeMessage) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/messages/{name} | Return a spoe message
[**GetSpoeScope**](SpoeApi.md#GetSpoeScope) | **Get** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{name} | Return one SPOE scope
[**ReplaceSpoeAgent**](SpoeApi.md#ReplaceSpoeAgent) | **Put** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/agents/{name} | Replace a SPOE agent
[**ReplaceSpoeGroup**](SpoeApi.md#ReplaceSpoeGroup) | **Put** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/groups/{name} | Replace a SPOE groups
[**ReplaceSpoeMessage**](SpoeApi.md#ReplaceSpoeMessage) | **Put** /services/haproxy/spoe/spoe_files/{parent_name}/scopes/{scope_name}/messages/{name} | Replace a spoe message


# **CreateSpoe**
> string CreateSpoe(ctx, optional)
Creates SPOE file with its entries

Creates SPOE file with its entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SpoeApiCreateSpoeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileUpload** | **optional.Interface of *os.File**| The spoe file to upload | 

### Return type

**string**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeAgent**
> SpoeAgent CreateSpoeAgent(ctx, parentName, scopeName, data, optional)
Add a new spoe agent to scope

Adds a new spoe agent to the spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **data** | [**SpoeAgent**](SpoeAgent.md)|  | 
 **optional** | ***SpoeApiCreateSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeGroup**
> SpoeGroup CreateSpoeGroup(ctx, parentName, scopeName, data, optional)
Add a new SPOE groups

Adds a new SPOE groups to the SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **data** | [**SpoeGroup**](SpoeGroup.md)|  | 
 **optional** | ***SpoeApiCreateSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeMessage**
> SpoeMessage CreateSpoeMessage(ctx, parentName, scopeName, data, optional)
Add a new spoe message to scope

Adds a new spoe message to the spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **data** | [**SpoeMessage**](SpoeMessage.md)|  | 
 **optional** | ***SpoeApiCreateSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSpoeScope**
> SpoeScope CreateSpoeScope(ctx, parentName, data, optional)
Add a new spoe scope

Adds a new spoe scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**SpoeScope**](SpoeScope.md)|  | 
 **optional** | ***SpoeApiCreateSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiCreateSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeScope**](spoe_scope.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeAgent**
> DeleteSpoeAgent(ctx, parentName, scopeName, name, optional)
Delete a SPOE agent

Deletes a SPOE agent from the configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe agent name | 
 **optional** | ***SpoeApiDeleteSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeFile**
> DeleteSpoeFile(ctx, name)
Delete SPOE file

Deletes SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SPOE file name | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeGroup**
> DeleteSpoeGroup(ctx, parentName, scopeName, name, optional)
Delete a SPOE groups

Deletes a SPOE groups from the one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe group name | 
 **optional** | ***SpoeApiDeleteSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeMessage**
> DeleteSpoeMessage(ctx, parentName, scopeName, name, optional)
Delete a spoe message

Deletes a spoe message from the SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe message name | 
 **optional** | ***SpoeApiDeleteSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSpoeScope**
> DeleteSpoeScope(ctx, parentName, name, optional)
Delete a SPOE scope

Deletes a SPOE scope from the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **name** | **string**| Spoe scope name | 
 **optional** | ***SpoeApiDeleteSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiDeleteSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeAgent**
> SpoeAgents GetAllSpoeAgent(ctx, parentName, scopeName, optional)
Return an array of spoe agents in one scope

Returns an array of all configured spoe agents in one scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
 **optional** | ***SpoeApiGetAllSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetAllSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeAgents**](spoe_agents.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeFiles**
> SpoeFiles GetAllSpoeFiles(ctx, )
Return all available SPOE files

Returns all available SPOE files.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SpoeFiles**](spoe_files.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeGroup**
> SpoeGroups GetAllSpoeGroup(ctx, parentName, scopeName, optional)
Return an array of SPOE groups

Returns an array of all configured SPOE groups in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
 **optional** | ***SpoeApiGetAllSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetAllSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeGroups**](spoe_groups.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeMessage**
> SpoeMessages GetAllSpoeMessage(ctx, parentName, scopeName, optional)
Return an array of spoe messages in one scope

Returns an array of all configured spoe messages in one scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
 **optional** | ***SpoeApiGetAllSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetAllSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeMessages**](spoe_messages.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSpoeScope**
> SpoeScopes GetAllSpoeScope(ctx, parentName, optional)
Return an array of spoe scopes

Returns an array of all configured spoe scopes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***SpoeApiGetAllSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetAllSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeScopes**](spoe_scopes.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneSpoeFile**
> InlineResponse200 GetOneSpoeFile(ctx, name)
Return one SPOE file

Returns one SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SPOE file name | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeAgent**
> SpoeAgent GetSpoeAgent(ctx, parentName, scopeName, name, optional)
Return a spoe agent

Returns one spoe agent configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe agent name | 
 **optional** | ***SpoeApiGetSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeConfigurationVersion**
> int32 GetSpoeConfigurationVersion(ctx, parentName, optional)
Return a SPOE configuration version

Returns SPOE configuration version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***SpoeApiGetSpoeConfigurationVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeConfigurationVersionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

**int32**

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeGroup**
> SpoeGroup GetSpoeGroup(ctx, parentName, scopeName, name, optional)
Return a SPOE groups

Returns one SPOE groups configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe group name | 
 **optional** | ***SpoeApiGetSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeMessage**
> SpoeMessage GetSpoeMessage(ctx, parentName, scopeName, name, optional)
Return a spoe message

Returns one spoe message configuration in SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe message name | 
 **optional** | ***SpoeApiGetSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeScope**
> SpoeScope GetSpoeScope(ctx, parentName, name, optional)
Return one SPOE scope

Returns one SPOE scope in one SPOE file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **name** | **string**| Spoe scope | 
 **optional** | ***SpoeApiGetSpoeScopeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiGetSpoeScopeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**SpoeScope**](spoe_scope.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeAgent**
> SpoeAgent ReplaceSpoeAgent(ctx, parentName, scopeName, name, data, optional)
Replace a SPOE agent

Replaces a SPOE agent configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe agent name | 
  **data** | [**SpoeAgent**](SpoeAgent.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeAgentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeAgent**](spoe_agent.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeGroup**
> SpoeGroup ReplaceSpoeGroup(ctx, parentName, scopeName, name, data, optional)
Replace a SPOE groups

Replaces a SPOE groups configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe group name | 
  **data** | [**SpoeGroup**](SpoeGroup.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeGroup**](spoe_group.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSpoeMessage**
> SpoeMessage ReplaceSpoeMessage(ctx, parentName, scopeName, name, data, optional)
Replace a spoe message

Replaces a spoe message configuration in one SPOE scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **scopeName** | **string**| Spoe scope name | 
  **name** | **string**| Spoe message name | 
  **data** | [**SpoeMessage**](SpoeMessage.md)|  | 
 **optional** | ***SpoeApiReplaceSpoeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SpoeApiReplaceSpoeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**SpoeMessage**](spoe_message.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

