# \StickTableApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStickTable**](StickTableApi.md#GetStickTable) | **Get** /services/haproxy/runtime/stick_tables/{name} | Return Stick Table
[**GetStickTableEntries**](StickTableApi.md#GetStickTableEntries) | **Get** /services/haproxy/runtime/stick_tables/{parent_name}/entries | Return Stick Table Entries
[**GetStickTables**](StickTableApi.md#GetStickTables) | **Get** /services/haproxy/runtime/stick_tables | Return Stick Tables
[**SetStickTableEntries**](StickTableApi.md#SetStickTableEntries) | **Post** /services/haproxy/runtime/stick_tables/{parent_name}/entries | Set Entry to Stick Table


# **GetStickTable**
> StickTable GetStickTable(ctx, name)
Return Stick Table

Returns one stick table from runtime.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Stick table name | 

### Return type

[**StickTable**](stick_table.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStickTableEntries**
> StickTableEntries GetStickTableEntries(ctx, parentName, optional)
Return Stick Table Entries

Returns an array of all entries in a given stick tables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***StickTableApiGetStickTableEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickTableApiGetStickTableEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| A list of filters in format data.&lt;type&gt; &lt;operator&gt; &lt;value&gt; separated by comma | 
 **key** | **optional.String**| Key which we want the entries for | 
 **count** | **optional.Int32**| Max number of entries to be returned for pagination | 
 **offset** | **optional.Int32**| Offset which indicates how many items we skip in pagination | 

### Return type

[**StickTableEntries**](stick_table_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStickTables**
> StickTables GetStickTables(ctx, )
Return Stick Tables

Returns an array of all stick tables.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StickTables**](stick_tables.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetStickTableEntries**
> SetStickTableEntries(ctx, parentName, optional)
Set Entry to Stick Table

Create or update a stick-table entry in the table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parentName** | **string**| Parent name | 
 **optional** | ***StickTableApiSetStickTableEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StickTableApiSetStickTableEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stickTableEntry** | [**optional.Interface of StickTableEntry1**](StickTableEntry1.md)| Stick table entry | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

