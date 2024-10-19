# \MailerEntryApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailerEntry**](MailerEntryApi.md#CreateMailerEntry) | **Post** /services/haproxy/configuration/mailer_entries | Add a new mailer_entry
[**DeleteMailerEntry**](MailerEntryApi.md#DeleteMailerEntry) | **Delete** /services/haproxy/configuration/mailer_entries/{name} | Delete a mailer_entry
[**GetMailerEntries**](MailerEntryApi.md#GetMailerEntries) | **Get** /services/haproxy/configuration/mailer_entries | Return an array of mailer_entries
[**GetMailerEntry**](MailerEntryApi.md#GetMailerEntry) | **Get** /services/haproxy/configuration/mailer_entries/{name} | Return one mailer_entry
[**ReplaceMailerEntry**](MailerEntryApi.md#ReplaceMailerEntry) | **Put** /services/haproxy/configuration/mailer_entries/{name} | Replace a mailer_entry


# **CreateMailerEntry**
> MailerEntry CreateMailerEntry(ctx, mailersSection, data, optional)
Add a new mailer_entry

Adds a new mailer entry to the specified mailers section in the configuration file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mailersSection** | **string**| Parent mailers section name | 
  **data** | [**MailerEntry**](MailerEntry.md)|  | 
 **optional** | ***MailerEntryApiCreateMailerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailerEntryApiCreateMailerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**MailerEntry**](mailer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMailerEntry**
> DeleteMailerEntry(ctx, name, mailersSection, optional)
Delete a mailer_entry

Deletes a mailer entry configuration by it's name in the specified mailers section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| MailerEntry name | 
  **mailersSection** | **string**| Parent mailers section name | 
 **optional** | ***MailerEntryApiDeleteMailerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailerEntryApiDeleteMailerEntryOpts struct

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

# **GetMailerEntries**
> MailerEntries GetMailerEntries(ctx, mailersSection, optional)
Return an array of mailer_entries

Returns an array of all the mailer_entries configured in the specified mailers section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mailersSection** | **string**| Parent mailers section name | 
 **optional** | ***MailerEntryApiGetMailerEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailerEntryApiGetMailerEntriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**MailerEntries**](mailer_entries.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMailerEntry**
> MailerEntry GetMailerEntry(ctx, name, mailersSection, optional)
Return one mailer_entry

Returns one mailer_entry configuration by it's name in the specified mailers section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| MailerEntry name | 
  **mailersSection** | **string**| Parent mailers name | 
 **optional** | ***MailerEntryApiGetMailerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailerEntryApiGetMailerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 

### Return type

[**MailerEntry**](mailer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceMailerEntry**
> MailerEntry ReplaceMailerEntry(ctx, name, mailersSection, data, optional)
Replace a mailer_entry

Replaces a mailer entry configuration by it's name in the specified mailers section.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| MailerEntry name | 
  **mailersSection** | **string**| Parent mailers section name | 
  **data** | [**MailerEntry**](MailerEntry.md)|  | 
 **optional** | ***MailerEntryApiReplaceMailerEntryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailerEntryApiReplaceMailerEntryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**MailerEntry**](mailer_entry.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

