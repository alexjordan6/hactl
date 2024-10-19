# \DiscoveryApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAPIEndpoints**](DiscoveryApi.md#GetAPIEndpoints) | **Get** / | Return list of root endpoints
[**GetConfigurationEndpoints**](DiscoveryApi.md#GetConfigurationEndpoints) | **Get** /services/haproxy/configuration | Return list of HAProxy advanced configuration endpoints
[**GetHaproxyEndpoints**](DiscoveryApi.md#GetHaproxyEndpoints) | **Get** /services/haproxy | Return list of HAProxy related endpoints
[**GetRuntimeEndpoints**](DiscoveryApi.md#GetRuntimeEndpoints) | **Get** /services/haproxy/runtime | Return list of HAProxy advanced runtime endpoints
[**GetServicesEndpoints**](DiscoveryApi.md#GetServicesEndpoints) | **Get** /services | Return list of service endpoints
[**GetSpoeEndpoints**](DiscoveryApi.md#GetSpoeEndpoints) | **Get** /services/haproxy/spoe | Return list of HAProxy SPOE endpoints
[**GetStatsEndpoints**](DiscoveryApi.md#GetStatsEndpoints) | **Get** /services/haproxy/stats | Return list of HAProxy stats endpoints
[**GetStorageEndpoints**](DiscoveryApi.md#GetStorageEndpoints) | **Get** /services/haproxy/storage | Return list of HAProxy storage endpoints


# **GetAPIEndpoints**
> Endpoints GetAPIEndpoints(ctx, )
Return list of root endpoints

Returns a list of root endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigurationEndpoints**
> Endpoints GetConfigurationEndpoints(ctx, )
Return list of HAProxy advanced configuration endpoints

Returns a list of endpoints to be used for advanced configuration of HAProxy objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHaproxyEndpoints**
> Endpoints GetHaproxyEndpoints(ctx, )
Return list of HAProxy related endpoints

Returns a list of HAProxy related endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeEndpoints**
> Endpoints GetRuntimeEndpoints(ctx, )
Return list of HAProxy advanced runtime endpoints

Returns a list of endpoints to be used for advanced runtime settings of HAProxy objects.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicesEndpoints**
> Endpoints GetServicesEndpoints(ctx, )
Return list of service endpoints

Returns a list of API managed services endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpoeEndpoints**
> Endpoints GetSpoeEndpoints(ctx, )
Return list of HAProxy SPOE endpoints

Returns a list of endpoints to be used for SPOE settings of HAProxy.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatsEndpoints**
> Endpoints GetStatsEndpoints(ctx, )
Return list of HAProxy stats endpoints

Returns a list of HAProxy stats endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStorageEndpoints**
> Endpoints GetStorageEndpoints(ctx, )
Return list of HAProxy storage endpoints

Returns a list of endpoints that use HAProxy storage for persistency, e.g. maps, ssl certificates...

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Endpoints**](endpoints.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

