# \ClusterApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCluster**](ClusterApi.md#DeleteCluster) | **Delete** /cluster | Delete cluster settings
[**EditCluster**](ClusterApi.md#EditCluster) | **Put** /cluster | Edit cluster settings
[**GetCluster**](ClusterApi.md#GetCluster) | **Get** /cluster | Return cluster data
[**InitiateCertificateRefresh**](ClusterApi.md#InitiateCertificateRefresh) | **Post** /cluster/certificate | Initiates a certificate refresh
[**PostCluster**](ClusterApi.md#PostCluster) | **Post** /cluster | Post cluster settings


# **DeleteCluster**
> DeleteCluster(ctx, optional)
Delete cluster settings

Delete cluster settings and move the node back to single mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterApiDeleteClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiDeleteClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configuration** | **optional.String**| In case of moving to single mode do we keep or clean configuration | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditCluster**
> ClusterSettings EditCluster(ctx, data, optional)
Edit cluster settings

Edit cluster settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ClusterSettings**](ClusterSettings.md)|  | 
 **optional** | ***ClusterApiEditClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiEditClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**ClusterSettings**](cluster_settings.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCluster**
> ClusterSettings GetCluster(ctx, )
Return cluster data

Returns cluster data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterSettings**](cluster_settings.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateCertificateRefresh**
> InitiateCertificateRefresh(ctx, )
Initiates a certificate refresh

Initiates a certificate refresh

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCluster**
> ClusterSettings PostCluster(ctx, data, optional)
Post cluster settings

Post cluster settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ClusterSettings**](ClusterSettings.md)|  | 
 **optional** | ***ClusterApiPostClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiPostClusterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | **optional.String**| In case of moving to single mode do we keep or clean configuration | 
 **advertisedAddress** | **optional.String**| Force the advertised address when joining a cluster | 
 **advertisedPort** | **optional.Int32**| Force the advertised port when joining a cluster | 
 **version** | **optional.Int32**| Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it&#39;s own version. | 

### Return type

[**ClusterSettings**](cluster_settings.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

