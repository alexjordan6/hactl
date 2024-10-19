# \MailersApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailersSection**](MailersApi.md#CreateMailersSection) | **Post** /services/haproxy/configuration/mailers_section | Add a new Mailers section
[**DeleteMailersSection**](MailersApi.md#DeleteMailersSection) | **Delete** /services/haproxy/configuration/mailers_section/{name} | Delete a Mailers section
[**EditMailersSection**](MailersApi.md#EditMailersSection) | **Put** /services/haproxy/configuration/mailers_section/{name} | Modify a Mailers section
[**GetMailersSection**](MailersApi.md#GetMailersSection) | **Get** /services/haproxy/configuration/mailers_section/{name} | Return a Mailers section
[**GetMailersSections**](MailersApi.md#GetMailersSections) | **Get** /services/haproxy/configuration/mailers_section | Return an array of mailers sections


# **CreateMailersSection**
> MailersSection CreateMailersSection(ctx, data, optional)
Add a new Mailers section

Creates a new empty Mailers section

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MailersSection**](MailersSection.md)|  | 
 **optional** | ***MailersApiCreateMailersSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailersApiCreateMailersSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**MailersSection**](mailers_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMailersSection**
> DeleteMailersSection(ctx, name, optional)
Delete a Mailers section

Deletes a mailers from the configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mailers name | 
 **optional** | ***MailersApiDeleteMailersSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailersApiDeleteMailersSectionOpts struct

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

# **EditMailersSection**
> MailersSection EditMailersSection(ctx, name, data, optional)
Modify a Mailers section

Modifies a mailers configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mailers name | 
  **data** | [**MailersSection**](MailersSection.md)|  | 
 **optional** | ***MailersApiEditMailersSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailersApiEditMailersSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**MailersSection**](mailers_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMailersSection**
> MailersSection GetMailersSection(ctx, name, optional)
Return a Mailers section

Returns one mailers configuration by it's name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Mailers name | 
 **optional** | ***MailersApiGetMailersSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailersApiGetMailersSectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**MailersSection**](mailers_section.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMailersSections**
> MailersSections GetMailersSections(ctx, optional)
Return an array of mailers sections

Returns an array of all the configured mailers in HAProxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MailersApiGetMailersSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MailersApiGetMailersSectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**| ID of the transaction where we want to add the operation. Cannot be used when version is specified. | 
 **fullSection** | **optional.Bool**| Indicates if the action affects the specified child resources as well | [default to false]

### Return type

[**MailersSections**](mailers_sections.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

