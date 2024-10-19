# \PeerEntryApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePeerEntry**](PeerEntryApi.md#CreatePeerEntry) | **Post** /services/haproxy/configuration/peer_entries | Add a new peer_entry
[**DeletePeerEntry**](PeerEntryApi.md#DeletePeerEntry) | **Delete** /services/haproxy/configuration/peer_entries/{name} | Delete a peer_entry
[**GetPeerEntries**](PeerEntryApi.md#GetPeerEntries) | **Get** /services/haproxy/configuration/peer_entries | Return an array of peer_entries
[**GetPeerEntry**](PeerEntryApi.md#GetPeerEntry) | **Get** /services/haproxy/configuration/peer_entries/{name} | Return one peer_entry
[**ReplacePeerEntry**](PeerEntryApi.md#ReplacePeerEntry) | **Put** /services/haproxy/configuration/peer_entries/{name} | Replace a peer_entry


# **CreatePeerEntry**
> PeerEntry CreatePeerEntry(ctx, peerSection, data, optional)
Add a new peer_entry

Adds a new peer entry in the specified peer section in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerSection** | **string**| Parent peer section name | 
  **data** | [**PeerEntry**](PeerEntry.md)|  | 
 **optional** | ***PeerEntryApiCreatePeerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerEntryApiCreatePeerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**PeerEntry**](peer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePeerEntry**
> DeletePeerEntry(ctx, name, peerSection, optional)
Delete a peer_entry

Deletes a peer entry configuration by it's name in the specified peer section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| PeerEntry name | 
  **peerSection** | **string**| Parent peers name | 
 **optional** | ***PeerEntryApiDeletePeerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerEntryApiDeletePeerEntryOpts struct

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

# **GetPeerEntries**
> PeerEntries GetPeerEntries(ctx, peerSection, optional)
Return an array of peer_entries

Returns an array of all peer_entries that are configured in specified peer section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **peerSection** | **string**| Parent peer section name | 
 **optional** | ***PeerEntryApiGetPeerEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerEntryApiGetPeerEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**PeerEntries**](peer_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeerEntry**
> PeerEntry GetPeerEntry(ctx, name, peerSection, optional)
Return one peer_entry

Returns one peer_entry configuration by it's name in the specified peer section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| PeerEntry name | 
  **peerSection** | **string**| Parent peers name | 
 **optional** | ***PeerEntryApiGetPeerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerEntryApiGetPeerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**PeerEntry**](peer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplacePeerEntry**
> PeerEntry ReplacePeerEntry(ctx, name, peerSection, data, optional)
Replace a peer_entry

Replaces a peer entry configuration by it's name in the specified peer section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| PeerEntry name | 
  **peerSection** | **string**| Parent peers name | 
  **data** | [**PeerEntry**](PeerEntry.md)|  | 
 **optional** | ***PeerEntryApiReplacePeerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PeerEntryApiReplacePeerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**PeerEntry**](peer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

