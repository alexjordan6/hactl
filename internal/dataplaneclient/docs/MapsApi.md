# \MapsApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMapEntry**](MapsApi.md#AddMapEntry) | **Post** /services/haproxy/runtime/maps/{parent_name}/entries | Adds an entry into the map file
[**AddPayloadRuntimeMap**](MapsApi.md#AddPayloadRuntimeMap) | **Put** /services/haproxy/runtime/maps/{name} | Add a new map payload
[**ClearRuntimeMap**](MapsApi.md#ClearRuntimeMap) | **Delete** /services/haproxy/runtime/maps/{name} | Remove all map entries from the map file
[**DeleteRuntimeMapEntry**](MapsApi.md#DeleteRuntimeMapEntry) | **Delete** /services/haproxy/runtime/maps/{parent_name}/entries/{id} | Deletes all the map entries from the map by its id
[**GetAllRuntimeMapFiles**](MapsApi.md#GetAllRuntimeMapFiles) | **Get** /services/haproxy/runtime/maps | Return runtime map files
[**GetOneRuntimeMap**](MapsApi.md#GetOneRuntimeMap) | **Get** /services/haproxy/runtime/maps/{name} | Return one runtime map file
[**GetRuntimeMapEntry**](MapsApi.md#GetRuntimeMapEntry) | **Get** /services/haproxy/runtime/maps/{parent_name}/entries/{id} | Return one map runtime setting
[**ReplaceRuntimeMapEntry**](MapsApi.md#ReplaceRuntimeMapEntry) | **Put** /services/haproxy/runtime/maps/{parent_name}/entries/{id} | Replace the value corresponding to each id in a map
[**ShowRuntimeMap**](MapsApi.md#ShowRuntimeMap) | **Get** /services/haproxy/runtime/maps/{parent_name}/entries | Return one map runtime entries


# **AddMapEntry**
> MapEntry AddMapEntry(ctx, parentName, data, optional)
Adds an entry into the map file

Adds an entry into the map file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
  **data** | [**MapEntry**](MapEntry.md)|  | 
 **optional** | ***MapsApiAddMapEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiAddMapEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceSync** | **optional.Bool**| If true, immediately syncs changes to disk | [default to false]

### Return type

[**MapEntry**](map_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPayloadRuntimeMap**
> MapEntries AddPayloadRuntimeMap(ctx, name, data, optional)
Add a new map payload

Adds a new map payload.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file name | 
  **data** | [**MapEntries**](MapEntries.md)|  | 
 **optional** | ***MapsApiAddPayloadRuntimeMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiAddPayloadRuntimeMapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceSync** | **optional.Bool**| If true, immediately syncs changes to disk | [default to false]

### Return type

[**MapEntries**](map_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearRuntimeMap**
> ClearRuntimeMap(ctx, name, optional)
Remove all map entries from the map file

Remove all map entries from the map file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file name | 
 **optional** | ***MapsApiClearRuntimeMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiClearRuntimeMapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **optional.Bool**| If true, deletes file from disk | 
 **forceSync** | **optional.Bool**| If true, immediately syncs changes to disk | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRuntimeMapEntry**
> DeleteRuntimeMapEntry(ctx, id, parentName, optional)
Deletes all the map entries from the map by its id

Delete all the map entries from the map by its id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Map id | 
  **parentName** | **string**| Parent name | 
 **optional** | ***MapsApiDeleteRuntimeMapEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiDeleteRuntimeMapEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceSync** | **optional.Bool**| If true, immediately syncs changes to disk | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRuntimeMapFiles**
> Maps GetAllRuntimeMapFiles(ctx, optional)
Return runtime map files

Returns runtime map files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MapsApiGetAllRuntimeMapFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiGetAllRuntimeMapFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUnmanaged** | **optional.Bool**| If true, also show unmanaged map files loaded in haproxy | [default to false]

### Return type

[**Maps**](maps.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneRuntimeMap**
> ModelMap GetOneRuntimeMap(ctx, name)
Return one runtime map file

Returns one runtime map file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file name | 

### Return type

[**ModelMap**](map.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeMapEntry**
> MapEntry GetRuntimeMapEntry(ctx, id, parentName)
Return one map runtime setting

Returns one map runtime setting by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Map id | 
  **parentName** | **string**| Parent name | 

### Return type

[**MapEntry**](map_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRuntimeMapEntry**
> MapEntry ReplaceRuntimeMapEntry(ctx, id, parentName, data, optional)
Replace the value corresponding to each id in a map

Replaces the value corresponding to each id in a map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Map id | 
  **parentName** | **string**| Parent name | 
  **data** | [**Data**](Data.md)|  | 
 **optional** | ***MapsApiReplaceRuntimeMapEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MapsApiReplaceRuntimeMapEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **forceSync** | **optional.Bool**| If true, immediately syncs changes to disk | [default to false]

### Return type

[**MapEntry**](map_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowRuntimeMap**
> MapEntries ShowRuntimeMap(ctx, parentName)
Return one map runtime entries

Returns an array of all entries in a given runtime map file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 

### Return type

[**MapEntries**](map_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

