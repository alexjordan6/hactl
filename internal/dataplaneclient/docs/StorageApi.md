# \StorageApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorageGeneralFile**](StorageApi.md#CreateStorageGeneralFile) | **Post** /services/haproxy/storage/general | Creates a managed storage general use file with contents
[**CreateStorageMapFile**](StorageApi.md#CreateStorageMapFile) | **Post** /services/haproxy/storage/maps | Creates a managed storage map file with its entries
[**CreateStorageSSLCertificate**](StorageApi.md#CreateStorageSSLCertificate) | **Post** /services/haproxy/storage/ssl_certificates | Create SSL certificate
[**DeleteStorageGeneralFile**](StorageApi.md#DeleteStorageGeneralFile) | **Delete** /services/haproxy/storage/general/{name} | Deletes a managed general use file from disk
[**DeleteStorageMap**](StorageApi.md#DeleteStorageMap) | **Delete** /services/haproxy/storage/maps/{name} | Deletes a managed map file from disk
[**DeleteStorageSSLCertificate**](StorageApi.md#DeleteStorageSSLCertificate) | **Delete** /services/haproxy/storage/ssl_certificates/{name} | Delete SSL certificate from disk
[**GetAllStorageGeneralFiles**](StorageApi.md#GetAllStorageGeneralFiles) | **Get** /services/haproxy/storage/general | Return a list of all managed general use files
[**GetAllStorageMapFiles**](StorageApi.md#GetAllStorageMapFiles) | **Get** /services/haproxy/storage/maps | Return a list of all managed map files
[**GetAllStorageSSLCertificates**](StorageApi.md#GetAllStorageSSLCertificates) | **Get** /services/haproxy/storage/ssl_certificates | Return all available SSL certificates on disk
[**GetOneStorageGeneralFile**](StorageApi.md#GetOneStorageGeneralFile) | **Get** /services/haproxy/storage/general/{name} | Return the contents of one managed general use file from disk
[**GetOneStorageMap**](StorageApi.md#GetOneStorageMap) | **Get** /services/haproxy/storage/maps/{name} | Return the contents of one managed map file from disk
[**GetOneStorageSSLCertificate**](StorageApi.md#GetOneStorageSSLCertificate) | **Get** /services/haproxy/storage/ssl_certificates/{name} | Return one SSL certificate from disk
[**ReplaceStorageGeneralFile**](StorageApi.md#ReplaceStorageGeneralFile) | **Put** /services/haproxy/storage/general/{name} | Replace contents of a managed general use file on disk
[**ReplaceStorageMapFile**](StorageApi.md#ReplaceStorageMapFile) | **Put** /services/haproxy/storage/maps/{name} | Replace contents of a managed map file on disk
[**ReplaceStorageSSLCertificate**](StorageApi.md#ReplaceStorageSSLCertificate) | **Put** /services/haproxy/storage/ssl_certificates/{name} | Replace SSL certificates on disk


# **CreateStorageGeneralFile**
> GeneralFile CreateStorageGeneralFile(ctx, optional)
Creates a managed storage general use file with contents

Creates a managed storage general use file with contents.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiCreateStorageGeneralFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiCreateStorageGeneralFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileUpload** | **optional.Interface of *os.File**| General use file content | 

### Return type

[**GeneralFile**](general_file.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStorageMapFile**
> ModelMap CreateStorageMapFile(ctx, optional)
Creates a managed storage map file with its entries

Creates a managed storage map file with its entries.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiCreateStorageMapFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiCreateStorageMapFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileUpload** | **optional.Interface of *os.File**| The map file contents | 

### Return type

[**ModelMap**](map.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStorageSSLCertificate**
> SslCertificate CreateStorageSSLCertificate(ctx, optional)
Create SSL certificate

Creates SSL certificate.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StorageApiCreateStorageSSLCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiCreateStorageSSLCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileUpload** | **optional.Interface of *os.File**| The SSL certificate to upload | 
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**SslCertificate**](ssl_certificate.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageGeneralFile**
> DeleteStorageGeneralFile(ctx, name)
Deletes a managed general use file from disk

Deletes a managed general use file from disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| General use file storage_name | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageMap**
> DeleteStorageMap(ctx, name)
Deletes a managed map file from disk

Deletes a managed map file from disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file storage_name | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStorageSSLCertificate**
> DeleteStorageSSLCertificate(ctx, name, optional)
Delete SSL certificate from disk

Deletes SSL certificate from disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SSL certificate name | 
 **optional** | ***StorageApiDeleteStorageSSLCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiDeleteStorageSSLCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipReload** | **optional.Bool**| If set, no reload will be initiated after update | [default to false]
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStorageGeneralFiles**
> GeneralFiles GetAllStorageGeneralFiles(ctx, )
Return a list of all managed general use files

Returns a list of all managed general use files

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GeneralFiles**](general_files.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStorageMapFiles**
> Maps GetAllStorageMapFiles(ctx, )
Return a list of all managed map files

Returns a list of all managed map files

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Maps**](maps.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStorageSSLCertificates**
> SslCertificates GetAllStorageSSLCertificates(ctx, )
Return all available SSL certificates on disk

Returns all available SSL certificates on disk.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SslCertificates**](ssl_certificates.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneStorageGeneralFile**
> *os.File GetOneStorageGeneralFile(ctx, name)
Return the contents of one managed general use file from disk

Returns the contents of one managed general use file from disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| General use file storage_name | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneStorageMap**
> *os.File GetOneStorageMap(ctx, name)
Return the contents of one managed map file from disk

Returns the contents of one managed map file from disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file storage_name | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOneStorageSSLCertificate**
> SslCertificate GetOneStorageSSLCertificate(ctx, name)
Return one SSL certificate from disk

Returns one SSL certificate from disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SSL certificate name | 

### Return type

[**SslCertificate**](ssl_certificate.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStorageGeneralFile**
> ReplaceStorageGeneralFile(ctx, name, optional)
Replace contents of a managed general use file on disk

Replaces the contents of a managed general use file on disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| General use file storage_name | 
 **optional** | ***StorageApiReplaceStorageGeneralFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiReplaceStorageGeneralFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileUpload** | **optional.Interface of *os.File**| General use file content | 
 **skipReload** | **optional.Bool**| If set, no reload will be initiated after update | [default to false]
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStorageMapFile**
> ReplaceStorageMapFile(ctx, name, data, optional)
Replace contents of a managed map file on disk

Replaces the contents of a managed map file on disk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Map file storage_name | 
  **data** | **string**|  | 
 **optional** | ***StorageApiReplaceStorageMapFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiReplaceStorageMapFileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipReload** | **optional.Bool**| If set, no reload will be initiated after update | [default to false]
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceStorageSSLCertificate**
> SslCertificate ReplaceStorageSSLCertificate(ctx, name, data, optional)
Replace SSL certificates on disk

Replaces SSL certificate on disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| SSL certificate name | 
  **data** | **string**|  | 
 **optional** | ***StorageApiReplaceStorageSSLCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StorageApiReplaceStorageSSLCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipReload** | **optional.Bool**| If set, no reload will be initiated after update | [default to false]
 **forceReload** | **optional.Bool**| If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration. | [default to false]

### Return type

[**SslCertificate**](ssl_certificate.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

